# Breaching Bare Metal Kubernetes Clusters -- Red Team Village

The only pre-reqs for these labs are to have access to a Kubernetes environment. Minikube will be just fine.

Once minikube is installed, and running, navigate to each folder and run the `lab.sh` script.

```bash
minikube start --cpus 2 --memory 8000 --nodes 3
```

# Order 

The intended order is as follows, but they can be run in any order as long as you run the clean up script at the end.

 - Namespaces
	 - Scanning without services Lab
	 - Scanning with services Lab
 - Secrets
	 - Mounted Secrets Lab
	 - steal_etcd Lab
 - RBAC
	 -  secret reader lab
	 - create_pod
- Pod Specs
	 - pod_breakout
