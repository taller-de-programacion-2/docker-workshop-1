# Paso 5: Usamos volumenes con docker-compose

## Usar networks entre microservicios

Con lo aprendido podemos solucionar como levantar multiples containers en una misma red haciendo que todos vivan y colaboren de manera simple y levantando en un solo comando de ```docker-compose up```.
Pero ahora surge la pregunta, si tengo un conjunto de microservicios que se comuniquen entre ellos, y no quiero exponer los puertos de mi computadora local y quedarme yo sin ellos, como puedo hacerlo?

Para esto existe la feature de networks, que si quieren profundizar mas sobre como funciona, siempre recomendamos ir a la documentacion oficial que esta ubicada en [su pagina](https://docs.docker.com/network/). Pero para aprender mejor ahora proponemos el siguiente ejercicio:

### Ejercicio 4

Tenemos los proyectos nodeApp y goApp, siendo nodeApp expuesto a la red de la computadora y goApp viviendo en una red interna de docker.

1) Buildear ambos proyectos con ```docker-compose up``` en cada uno.
2) Verificar que los logs de nodeApp indican que no nos podemos conectar a goApp.
3) Modificar los docker-compose correspondientes para que ambos proyectos se puedan ver.
4) Verificar que nodeApp se puede comunicar con goApp en sus logs.



[< Agregamos una DB](04_database.md) | [Inicio >](README.md)
