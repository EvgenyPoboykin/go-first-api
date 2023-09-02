run:
	go run .

build:
	go build

d-build:
	docker build -t first-api-app .

d-run:
	docker run -p 80:80 -it --rm --ip 0.0.0.0 --name first-api-running-app first-api-app 
