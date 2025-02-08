# Prereqs

```bash
minikube start --cpus 2 --memory 8000 --nodes 3
```

# Scanning without services Lab 
- In this lab we find a pod in a different namespace
```bash
# Start lab
./scanning_no_service.sh
```
1. Set a sane prompt
```bash
# Check shell
echo $SHELL
# or
cat /etc/shells

PS1='\[\e[31m\]\u\[\e[96m\]@\[\e[35m\]\H\[\e[0m\]:\[\e[93m\]\w\[\e[0m\]\$ '


echo "Useful tools:"; missing=""; for i in kubectl hostname perl python python3 dpkg bash sh yq jq nmap curl wget ping apt apk openssl nc netcat sed vim vi nano base64 tar; do command -v "$i" >/dev/null 2>&1 && echo "$i" || missing="$missing $i"; done; if [ -n "$missing" ]; then echo  "Missing tools: $(echo "$missing" | sort)"; fi

```
2. We can install things!
```bash
apt update ; apt install libpcap-dev masscan -y
```
3. Run mass scan
```bash

LOS_IP=$(hostname -i)

masscan -p80,443,1337 $LOS_IP/16 --rate 25000
```
4. Curl the ips returned to show we can reach them
5. Exit container and look at the cluster from the admin view and note they're on different nodes
```bash
kubectl get pods --namespace eng1 --output wide
kubectl get pods --namespace eng2 --output wide
```

6. Delete the lab
```bash
kubectl delete -f ./scanning_no_service.yaml --force
```
