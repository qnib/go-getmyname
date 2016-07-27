# go-getmyname

Golang tool to extract container-name from within a container.

Either the container finds a socket under `/var/run/docker.sock`...
```
$ docker run -ti --rm -v /var/run/docker.sock:/var/run/docker.sock:ro --name myname \
                 -v $(pwd):/go-getmyname/ qnib/alpn-consul /usr/local/bin/go-getmyname
myname
```
... or he is able to create a client out of the environment (DOCKER_HOST).
For now I do not support TLS encrypted clients. PRs are welcome... :)

```
$ docker run -ti --rm --name myname -e DOCKER_HOST=tcp://172.17.0.1:2376 \
                 qnib/alpn-consul /usr/local/bin/go-getmyname
myname
```
