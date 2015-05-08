# Simple microservice dummy app
used for testing CI and minimal docker images.
See: https://github.com/alpe/ci-server

## Get started
~~~
go test -v ./...
~~~


### Docker
build image local:
~~~
docker run --rm \
  -v $(pwd):/src \
  -v /var/run/docker.sock:/var/run/docker.sock \
  $(boot2docker ip):5000/go_builder_image \
  ci-example-project
~~~

Start container
~~~
EXT_PORT=8090
CONTAINER_ID=$(docker run -d -p 127.0.0.1:$EXT_PORT:8080 ci-example-project)
~~~
Open http://127.0.0.1:$EXT_PORT

stop + cleanup
~~~
docker stop $CONTAINER_ID
docker rm -v $CONTAINER_ID
~~~
