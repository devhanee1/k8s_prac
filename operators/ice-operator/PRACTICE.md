# Kubebuilder init 
```
kubebuilder init --domain my.domain --repo my.domain/guestbook
```

# Kubebuilder create api/controller
```
kubebuilder create api --group webapp --version v1 --kind Guestbook
```

# apiVersion/Kind Summary
- apiVersion is determined with group info + domain info + "/" + version info
- domain : my.domain
- group  : webapp
- version : v1
- So, the apiVersion : webapp.my.domain/v1

- Kind: kind 

# Rule for Import in the golang code (e.g : Controller)
```
import (
	webappv1 "my.domain/guestbook/api/v1"
)
```
- Prefix for import is started with "repo" 
- repo is my.domain/guestbook, so import statement for CR is "my.domain/guestbook/api/v1" 


# Sample of CR (Custom Resource)
```
apiVersion: webapp.my.domain/v1
kind: Guestbook
metadata:
  name: gestbook-sample
spec:
  size: 500
  nickname: duck-swim
```

# Install CRD into the Cluster
```
make install
```

# Run controller
```
make run
```

# Apply sample
```
k apply -f config/samples/
```

# Build and Push Controller & Deploy controller to Cluster
```
make docker-build docker-push IMG=<some-registry>/<project-name>:tag
make deploy IMG=<some-registry>/<project-name>:tag
```

# Test point of controller
## Reconcile Loop

```
 54     // TODO(user): your logic here
 55     logger := log.FromContext(ctx).WithValues("req.Namespace", req.Namespace, "req.Name", req.Name)
 56     logger.Info("*****  Reconcile Loop  ******************\n " + req.Name)
 57
 58     gbook := &webappv1.Guestbook{}
 59     r.Client.Get(context.TODO(), req.NamespacedName, gbook)
 60     logger.Info("***** CR Info ********************\n>> size: " + strconv.Itoa(int(gbook.Spec.Size)) + "\n>> NickName: " + gbook.Spec.NickName + "\n")
```
- Make Log
	- Pass context and get instance
	- Filter with namespace and name
```
log.FromContext(ctx).WithValues(...)
```
- Make empty CR instance & set using GET function
	- Pass context, namespace and CR instance pointer
```
gbook := &webappv1.Guestbook{}
r.Client.Get(context.TODO(), req.NamespaceName, gbook)
```

