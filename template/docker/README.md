# Build
```docker
docker build \           
-f deploy/Dockerfile \
-t lynam/go-oauth \
--build-arg GITHUB_TOKEN=d7519779515245e32e346444a458377f501b70b0 \
.
```

# Run service
#### Development Mode
```docker
docker-compose -f deploy/docker-compose.yml \
    -f deploy/docker-compose.override.yml \
    run --name go-oauth -d go-oauth
```

#### Production Mode
```docker
docker-compose -f deploy/docker-compose.yml \
    -f deploy/docker-compose.prod.yml \
    run --name go-oauth -d go-oauth
```
