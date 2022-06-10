## docker build
docker build --build-arg custom_port=18888 -t bluesgao/http-echo:v1 .
docker build -t bluesgao/http-echo:v1 .

## docker run
docker run --name http_echo_1 -e HTTP_ECHO_CUSTOM_PORT=19999 -p 19999:19999 -di bluesgao/http-echo:v1
docker run --name http_echo_2 -e HTTP_ECHO_CUSTOM_PORT=29999 -p 29999:29999 -di bluesgao/http-echo:v1

## docker push(前置条件：docker login)
docker push bluesgao/http-echo:v4 