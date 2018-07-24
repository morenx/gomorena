clean:
	rm gomorena

deps:
	dep ensure

build:
	go build -o gomorena main.go

lint:
	golint

test: lint
	go test -cover -v ./...

run: build
	@echo 'Server running on localhost:8080'
	./gomorena

db:
	@docker run \
		--name morenadb \
		-e MYSQL_ROOT_PASSWORD=${LOCAL_MYSQL_PW} \
		-p3306:3306 \
		-d \
		mariadb:latest

dbrm:
	docker rm -f morenadb

dblogin:
	mysql mysql -h 127.0.0.1 -P 3306 -u root -p

dbloginMac:
	@echo 'Use command: mysql mysql -h 127.0.0.1 -P 3306 -u root -p'
	docker exec -it morenadb bash
