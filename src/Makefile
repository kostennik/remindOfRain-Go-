build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
	docker build -t weather:latest .
	rm ./main
run:
	docker run -it  --network host --rm weather:latest
push:
	docker push weather:latest

run-daemon:
	docker run -d --network host --name weather --restart always weather