# Milestone 1

CLI for the HTTP client for the secret sharing web application will be:

```bash
secret-share <verb> [flags]
```

We only have 2 possible commands (actions):

1. `view` --url=url-of-server --id=id
2. `create` --url=url-of-server --data=some-secret-text

So, some examples:

```bash
# Create a new secret
secret-share create --url=localhost:8080/ --data="super-secret-colour"
# output: id=<some-id>

# View a created secret
secret-share view --url=localhost:8080/ --id=secret-colour-hashed-id
# output: data=<super-secret-colour>
```

## Structure of CLI: CLI Command Structure Reference

### Noun-First (`tool <noun> <verb> [flags]`)

Group by resource. All actions on a resource live together.

```
git remote add
git remote remove
git remote list

docker container ls
docker container start
docker container stop

aws s3 cp
aws s3 ls
aws s3 rm
```

**Good for:** tools with a fixed set of resources and many operations per resource.

---

### Verb-First (`tool <verb> <noun> [flags]`)

Group by action. One verb applies across many resource types.

```
kubectl get pods
kubectl get services
kubectl delete pods
kubectl describe nodes

systemctl start nginx
systemctl stop nginx
systemctl status nginx
```

**Good for:** tools with a small, stable set of actions applied broadly.

---

### Flags = Adjectives

Modify the noun/verb without changing the sentence shape. Keep meaning consistent across all subcommands.

```
kubectl get pods --namespace=prod --output=json
git branch list --merged --sort=-committerdate
```