.PHONY: tomato
tomato_test:
	helm template tomato . --set nodePort=30007 --namespace tomato

.PHONY: tomato
tomato:
	helm install tomato . --set nodePort=30007 --namespace tomato

remove:
	helm delete tomato --namespace tomato
