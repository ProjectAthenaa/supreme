build:
	docker build --build-arg GH_TOKEN=$(token)  -t registry.digitalocean.com/athenabot/modules/supreme:latest .