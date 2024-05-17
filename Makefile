build:
	docker build --build-arg GITHUB_USER=${TR_GIT_USER} --build-arg GITHUB_TOKEN=${TR_GIT_TOKEN} -t github.com/turistikrota/service.business . 

run:
	docker service create --name business-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --env-file .env --publish 6021:6021 --publish 7021:7021 github.com/turistikrota/service.business:latest

remove:
	docker service rm business-api-turistikrota-com

stop:
	docker service scale business-api-turistikrota-com=0

start:
	docker service scale business-api-turistikrota-com=1

restart: remove build run
	