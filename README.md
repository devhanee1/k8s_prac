# K8s prac 


## tips

### Setting host file 
- /private/etc/hosts
- add 127.0.0.1 ingress.example.local


### Check ports
- Check opened port
```
sudo lsof -PiTCP -sTCP:LISTEN
```
- Check PID by port
```
sudo lsof -i :3000
```

### Port forwarding
[TBD]
