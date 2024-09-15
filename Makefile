NAME := skynet

all: bin docker

bin:
	GOOS=linux GOARCH=amd64 go build -a -o $(NAME)-amd64-linux .
	GOOS=linux GOARCH=arm64 go build -a -o $(NAME)-arm64-linux .

docker:
	docker buildx build --platform linux/amd64,linux/arm64 -t tenox7/$(NAME):latest --load .

docker-push:
	docker buildx build --platform linux/amd64,linux/arm64 -t tenox7/$(NAME):latest --push .

clean:
	rm -f $(NAME) $(NAME)-*
