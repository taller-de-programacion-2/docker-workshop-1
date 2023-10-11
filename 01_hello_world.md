# Paso 1: Hello World

## Descargar una imagen de Docker

Para descargar una imagen de [Docker Hub](https://hub.docker.com/), por ejemplo: https://hub.docker.com/_/ubuntu

```
docker pull ubuntu
```

Al ejecutar ese comando, deberían obtener un output similar a esto:

```
Using default tag: latest
latest: Pulling from library/ubuntu
5bed26d33875: Pull complete
f11b29a9c730: Pull complete
930bda195c84: Pull complete
78bf9a5ad49e: Pull complete
Digest: sha256:bec5a2727be7fff3d308193cfde3491f8fba1a2ba392b7546b43a051853a341d
Status: Downloaded newer image for ubuntu:latest
docker.io/library/ubuntu:latest
```

## Hello World y otros

Si quisieramos levantar un container que imprima "Hello World" podríamos hacer:

```
docker run ubuntu /bin/echo 'Hello world'
```

Y si quisieramos en lugar de imprimir "Hello World", levantar una terminal en bash? Podríamos hacer:

```
docker run ubuntu /bin/bash
```

Pero con esto no alcanza, ya que simplemente va a correr /bin/bash pero no vamos a poder acceder a ese proceso. Para eso deberíamos hacer:

```
docker run -i -t ubuntu /bin/bash
```

De esta forma, deberían ver un prompt similar al siguiente:

```
> root@19d8ebf25422:/# 
```

Veamos que hacen los flags esos:

```
docker run --help
```

```
Usage:	docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

Run a command in a new container

Options:
  ...
  -i, --interactive                    Keep STDIN open even if not attached
  ...
  -t, --tty                            Allocate a pseudo-TTY
```

Pero por default los containers no se eliminan, sino que se ponen en `stop`. Si hacemos:

```
docker ps -a
```

Vemos todos los containers que estan en `stop` y no se eliminaron:

```
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS                          PORTS               NAMES
19d8ebf25422        ubuntu              "/bin/bash"              3 minutes ago       Exited (0) About a minute ago                       infallible_ritchie
838aec57698a        ubuntu              "/bin/bash"              6 minutes ago       Exited (0) 6 minutes ago                            suspicious_edison
5ea22c610847        ubuntu              "/bin/echo 'Hello wo…"   9 minutes ago       Exited (0) 9 minutes ago                            keen_yonath
```

Si quisieramos podemos hacer:

- `docker start <CONTAINER ID>`: para poder restartear un container
- `docker stop <CONTAINER ID>`: para poder stopear un container

Una solucion un poco mas fuerte pero para cuando queremos empezar en serio de 0 es hacer:

- `docker system prune -a`: para eliminar todas las imagenes, containers y volumenes que no se esten usando actualmente

[< Introducción](00_introduction.md) | [Primer Dockerfile >](02_first_dockerfile.md)
