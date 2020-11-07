# hello-world
HelloWorld project

## Description
Initial trial lambda function in GoLang for handling Hello World requests

## Build 
```bash
docker run --rm -v "$PWD":/go/src/handler:Z lambci/lambda:build-go1.x sh -c 'go build main.go'
```

## Run locally
```bash
docker run --rm -v "$PWD":/var/task:Z lambci/lambda:go1.x main
```

## Developing code
### Running, keep-alive, live reload on code change
```bash
docker run --restart on-failure --rm \
  -e DOCKER_LAMBDA_WATCH=1 -e DOCKER_LAMBDA_STAY_OPEN=1 -p 9001:9001 \
  -v "$PWD":/var/task:Z \
  lambci/lambda:go1.x main
```

### Testing lambda function locally in dev
```bash
curl -d '{}' http://localhost:9001/2015-03-31/functions/myfunction/invocations
```
