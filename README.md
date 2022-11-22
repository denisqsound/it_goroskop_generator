docker buildx build --platform=linux/arm64 .
docker buildx build --platform=linux/arm64 -t it_horoscope:0.0.4 .
docker buildx build --platform=linux/arm64  -t denisqsound/it_horoscope:0.0.5 --build-arg ARCH=amd64 .
