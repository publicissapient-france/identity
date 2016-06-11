# Identity

## Build application

```shell
$ make
```

If you only want to compile the application

```shell
$ make compile
```

If you only want to build the Docker image

```shell
$ make build-image
```

Push the Docker image to DockerHub

```shell
$ make push-image
```

## Run application

Run the application

```shell
$ docker run -d -p 8080:8080 jlrigau/identity
```

Run the application with a specific name

```shell
$ docker run -d -e NAME=Unicorn -p 8080:8080 jlrigau/identity
```

Run the application with a specific filename

**Note** The corresponding file must be present in the `static/img` directory

```shell
$ docker run -d -e NAME=Unicorn -e FILENAME=unicorn.jpg \
				-p 8080:8080 jlrigau/identity
```

Run the application by indicating the url of the image to use

**Note** If the downloaded image has a different filename from `identity.png`,
you have to indicate a specific filename according to the file

```shell
$ docker run -d -e NAME="Unicorn" -e FILENAME=unicorn.jpeg \
				-e URL=https://index.co/uploads/lists/a981c586ee454b2f0210d64d013870dab46332c8.jpeg \
				-p 8080:8080 jlrigau/identity
```

## Run application with Docker Compose

Run the application with a Redis instance

```shell
$ docker-compose up -d
```

**Warning** Due to an issue with the Docker internal DNS on VirtualBox, you cannot run the application
on local environment using Docker Machine

Use the Redis `MONITOR` command for listening for all requests received by the server in real time

```shell
$ docker-compose run --rm redis redis-cli -h redis
redis:6379> MONITOR
OK
```

Then go to the application on `http://<docker_host>:8080/identity` and click on the `Hit` button.

Verify that a new message has been published on `service.hit` channel on Redis

```shell
1465686292.798837 [0 172.19.0.3:44510] "PUBLISH" "service.hit" "{\"name\":\"Unicorn\",\"filename\":\"unicorn.jpg\",\"hostname\":\"e8db00fd2b2e\",\"hits\":2}"
```