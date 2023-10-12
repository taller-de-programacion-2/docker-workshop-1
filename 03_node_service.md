# Paso 3: Servicio Node

## Levantar un Servicio Node

Ahora vamos a usar algunos comandos nuevos para crear un `Dockerfile` para buildear una imagen que permita levantar el servicio que se encuentra en `/resources/03_node_service`.

Comandos nuevos de `Dockerfile`:

- `COPY` -- copies files or directories from source and adds them to the filesystem of the container at destination
- `WORKDIR` -- set working directory
- `EXPOSE` -- expose port
- `CMD` -- set executable for container

Por ejemplo, el siguiente podría ser el formato de un `Dockerfile` muy simple:

```
FROM <IMAGEN_BASE>

WORKDIR <WORKDIR_FOLDER>

COPY <SRC> <DST>

RUN <COMANDO>

RUN <OTRO_COMANDO>

EXPOSE <PUERTO>

ENV <ENV_VAR_NOMBRE>=<ENV_VAR_VALOR>

CMD <COMANDO PARA INICIAR PROGRAMA>
```

### Ejercicio 2

1. Crear un Dockerfile que corra el servicio node (ubicado en `/resources/03_node_service`) y cumpla las siguientes condiciones:

- Imagen Base: `node:18` de https://hub.docker.com/_/node
- Working Directory: `/app`
- Exponga el puerto 3000, donde la aplicación escuchara las requests.
- Setee la variable de ambiente PORT en el valor 3000.
- Instale las dependencias del servicio Node con el comando `npm i`.
- Inicie el servicio Node con el comando `npm start`.

2. Buildear la imagen con:

```
docker build . -t node-service
```

Y luego verificar que se pudo buildear correctamente con:

```
docker image ls
```

3. Levantar un container usando la imagen buildeada y los siguientes flags:

- `-e <ENV_VAR>=<ENV_VAR_VALUE>`: Para definir/sobreescribir una variable de entorno, en este caso se puede usar para sobreescribir el puerto donde correra internamente la aplicación.

- `-p <HOST_PORT>:<CONTAINER_PORT>`: Para publicar un puerto expuesto al host. (Opcional: ¿Cual es la diferencia con -P?)


4. Verificar que el servicio esta levantado y corriendo. Correr el comando: 

```
curl http://localhost:<HOST_PORT>/ping
```
Y luego verificar que la respuesta sea:
```
Pong
```

#### Notas

- Para instalar las dependencias especificadas en el [`package.json`](/resources/node_service/package.json) se debe correr:

```
npm i
```

### Preguntas bonus:

- Por que no funciona ctrl c bien?
- Cual es la diferencia entre docker stop _CONTAINER_ID_ y docker kill _CONTAINER_ID_?
- Cual es la diferencia entre CMD y ENTRYPOINT?
[< Primer Dockerfile](02_first_dockerfile.md) | [ Agregamos una DB>](04_database.md)
