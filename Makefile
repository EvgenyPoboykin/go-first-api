run:
	go run .

build:
	go build

d-build:
	docker build -t first-api-app .

d-run:
	docker run -p 7878:7878 -it --rm --name first-api-running-app first-api-app 
