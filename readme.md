## docker build
docker build --build-arg port=18888 -t bluesgao/http-echo:v1 .
docker build -t bluesgao/http-echo:v1 .

## docker run
docker run --name http_echo_1 -e HTTP_ECHO_PORT=19999 -p 19999:19999 -di bluesgao/http-echo:v1
docker run --name http_echo_2 -e HTTP_ECHO_PORT=29999 -p 29999:29999 -di bluesgao/http-echo:v1
