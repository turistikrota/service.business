# Run Project

How to run project

## 1. add swarm network

```bash
docker network create --driver overlay --attachable turistikrota

```

## 2. add secrets

```bash
docker secret create jwt_private_key ./jwtRS256.key
docker secret create jwt_public_key ./jwtRS256.key.pub

```

## 3. build image

```bash
docker build --build-arg GITHUB_USER=<USER_NAME> --build-arg GITHUB_TOKEN=<ACCESS_TOKEN> -t github.com/turistikrota/service.owner .  
```

## 4. run container

```bash
docker service create --name owner-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --env-file .env --publish 6015:6015 github.com/turistikrota/service.owner:latest
```
