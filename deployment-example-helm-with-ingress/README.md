# Helm Usage 
## Start & Remove 
- helm install [release-name] [chart-location]
```
$ helm list
$ helm install release_name .
$ helm delete release_name
```

## With namespace
- helm install --namespace [Namespace-name] [release-name] [chart-location]
```
$ helm install --namespace orange b1 .
$ helm delete --namespace orange b1
$ helm list --namespace orange
```

## With param
- helm install --namespace [namespace-name] [release-name] [char-location] --set param=value
```
$ helm install apple . --set nodePort=30009
```

# Check config map
```
k get cm nginx-config -o yaml
```

## Setting host file 
- /private/etc/hosts
- add 127.0.0.1 ingress.example.local


## Check ports
- Check opened port
```
sudo lsof -PiTCP -sTCP:LISTEN
```
- Check PID by port
```
sudo lsof -i :3000
```

## Precondition for using ingress
- Create ingressClass
- refer : https://sam-thomas.medium.com/kubernetes-ingressclass-error-ingress-does-not-contain-a-valid-ingressclass-78aab72c15a6
