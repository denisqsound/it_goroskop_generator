
# СБИЛДИТЬ ПРОЕКТ И ЗАПУШИТЬ
.PHONY: dc
dc: dc-build dc-push

.PHONY: dc-up
dc-up:
	docker rm horoscope ; docker run -d -p 127.0.0.1:3000:3000/tcp --name horoscope denisqsound/it_horoscope:latest

.PHONY: dc-build
dc-build:
	docker buildx build --platform=linux/amd64  -t denisqsound/it_horoscope:latest --build-arg ARCH=amd64 .

.PHONY: dc-push
dc-push:
	docker push denisqsound/it_horoscope:latest


.PHONY: dc-clean
dc-clean:
	docker stop horoscope; docker rmi denisqsound/it_horoscope:latest --force


.PHONY: dc-upgrade
dc-upgrade: dc-clean dc-up