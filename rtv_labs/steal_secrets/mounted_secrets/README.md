# View Secrets in container
1. Start lab
```bash
./mounted_secrets.sh
```

```bash

echo $SHELL

PS1='\[\e[31m\]\u\[\e[96m\]@\[\e[35m\]\H\[\e[0m\]:\[\e[93m\]\w\[\e[0m\]\$ '

env
# Or
export


# Additionally check for secrets in the filesystem

# Install tree
apt update && apt install tree

 cd / ; tree -L 2

/etc/secrets/ssh-key
```

