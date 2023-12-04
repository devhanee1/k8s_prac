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

