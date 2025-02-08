# Scanning with Services 
- A more realistic way of looking for pods
- View `dnscan.go`
- `./scanning_with_service.sh`

1. Set a sane prompt
```bash
# Check shell
echo $SHELL
# or
cat /etc/shells

PS1='\[\e[31m\]\u\[\e[96m\]@\[\e[35m\]\H\[\e[0m\]:\[\e[93m\]\w\[\e[0m\]\$ '

echo "Useful tools:"; missing=""; for i in kubectl hostname perl python python3 dpkg bash sh yq jq nmap curl wget ping apt apk openssl nc netcat sed vim vi nano base64 tar; do command -v "$i" >/dev/null 2>&1 && echo "$i" || missing="$missing $i"; done; if [ -n "$missing" ]; then echo  "Missing tools: $(echo "$missing" | sort)"; fi

```
2. Downlaod dnscan
```bash
curl -L https://github.com/LowOrbitSecurity/dnscan/releases/download/latest/dnscan > dnscan && chmod +x dnscan && LOS_IP=$(hostname -i); echo $LOS_IP

./dnscan -subnet $LOS_IP/16

```
3. Delete the lab
```bash
kubectl delete -f ./scanning_with_service.yaml --force
```
