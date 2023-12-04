.PHONY: lion_test
lion_test:
	helm template lion . --set nodePort=30007 --set message="events { } http { server { listen 80; location / { return 200 'Hello\, Lion!!\, version 0.8 '; } } }" --namespace lion
.PHONY: tiger_test
tiger_test:
	helm template tiger . --set nodePort=30008 --set message="events { } http { server { listen 80; location / { return 200 'Hello\, Tiger!!\, version 0.8 '; } } }" --namespace tiger

all_test: lion_test tiger_test


.PHONY: lion
lion:
	helm install lion . --set nodePort=30007 --set message="events { } http { server { listen 80; location / { return 200 'Hello\, Lion!!\, version 0.8 '; } } }" --namespace lion

.PHONY: tiger
tiger:
	helm install tiger . --set nodePort=30008 --set message="events { } http { server { listen 80; location / { return 200 'Hello\, Tiger!!\, version 0.8 '; } } }" --namespace tiger

all: lion tiger

remove:
	helm delete lion --namespace lion
	helm delete tiger --namespace tiger

