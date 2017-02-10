# drone-webhook

## build
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
    docker build -t zyclonite/drone-webhook .

## release
    docker tag zyclonite/drone-webhook:latest registry.hub.docker.com/zyclonite/drone-webhook:latest
    docker push registry.hub.docker.com/zyclonite/drone-webhook:latest
