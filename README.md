# drone-webhook

## build
    go build
    docker build -t zyclonite/drone-webhook .

## release
    docker tag zyclonite/drone-webhook:latest registry.hub.docker.com/zyclonite/drone-webhook:latest
    docker push registry.hub.docker.com/zyclonite/drone-webhook:latest
