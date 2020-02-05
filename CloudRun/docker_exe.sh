IMAGE_NAME=go-lang-cloud-run-test:latest

# docker rmi ${IMAGE_NAME}

docker build -t ${IMAGE_NAME} -f ./Dockerfile .
docker run --rm -it -p 8080:8080 ${IMAGE_NAME} /bin/bash

# docker rmi ${IMAGE_NAME}
