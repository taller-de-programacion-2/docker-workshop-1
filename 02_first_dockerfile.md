# Paso 2: Primer Dockerfile

## Nuestro primer Dockerfile

Lo que vimos hasta ahora nos sirve para poder levantar un container usando una imagen determinada, pero si quisieramos hacer nuestra propia imagen deberíamos hacer un `Dockerfile`.

Algunos comandos de `Dockerfile`:

- `FROM` -- set base image
- `RUN` -- execute command in container

Por ejemplo, el siguiente podría ser el formato de un `Dockerfile` muy simple:

```
FROM <IMAGEN_BASE>

RUN <COMANDO>
```

### Ejercicio 1

1. Crear un Dockerfile que tenga como imagen base `ubuntu` y que instale `Python 3`.
    
- Los comandos para instalar `python3` en ubuntu son:

    ```
    apt-get -y update
    apt-get -y install python3
    ```

- Pensar por qué es necesario el flag `-y`.


2. Buildear la imagen con:

```
docker build . -t python-image
```

Y luego verificar que se pudo buildear correctamente con:

```
docker image ls
```

3. Levantar un container que corra `python3` usando la imagen buildeada y lo aprendido en el [Hello World].

### Por puntos inexistentes extra:

4. Cambiar el dockerfile para que solamente use python sin tener que pullear una imagen de ubuntu

[< Hello World](01_hello_world.md) | [Servicio Node >](03_node_service.md)
