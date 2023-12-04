.PHONY: lion-test
lion-test:
	helm template lion .  --set message="events { } http { server { listen 80; location / { return 200 'Hello\, Lion!!\, version 0.8 '; } } }" --namespace lion
.PHONY: tiger-test
tiger-test:
	helm template tiger . --set message="events { } http { server { listen 80; location / { return 200 'Hello\, Tiger!!\, version 0.8 '; } } }" --namespace tiger

all_test: lion_test tiger_test


.PHONY: lion
lion:
	helm install lion . --set message="events { } http { server { listen 80; location / { return 200 'Hello\, Lion!!\, version 0.8 '; } } }" --namespace lion

.PHONY: tiger
tiger:
	helm install tiger . --set message="events { } http { server { listen 80; location / { return 200 'Hello\, Tiger!!\, version 0.8 '; } } }" --namespace tiger

all: lion tiger

.PHONY: tiger-remove
tiger-remove:
	helm delete tiger --namespace tiger


.PHONY: lion-remove
lion-remove:
	helm delete lion --namespace lion

remove:
	helm delete lion --namespace lion
	helm delete tiger --namespace tiger

