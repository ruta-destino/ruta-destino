# Ruta destino
Ruta destino es un proyecto que permite tener un control de los horarios que
realizan los buses en diferentes sectores. Es de libre navegación, por lo que
se puede conocer la información de cualquier lugar. La información viene
directamente de parte de las empresas de buses.

El proyecto está dividido en al menos 4 unidades lógicas. Por una parte está el
backend, que contiene todo lo relacionado a las comunicaciones con la base de
datos. El backend fue escrito usando el framework Fiber y el lenguaje de
programación Go. La elección de la base de datos fue PostgreSQL.

Como infraestructura la solución incluye información como lo son las
ubicaciones que luego serán usadas por las empresas. Esto incluye las regiones,
provincias, ciudades y además los terminales que hay por ciudad. Para manejar
esta información se creó un frontend dedicado a los administradores de Ruta
Destino. El frontend fue escrito con Svelte/SvelteKit.

Las empresas tendrán su propia interfaz dedicada para la administración de
datos. Prinpalmente lo que van a gestionar son sus recorridos, pero además
pueden almacenar información extra como lo pueden ser los tipos de pasajeros,
las paradas que tienen sus recorridos, las tarifas, etc. Por supuesto cada
empresa decide que información publicará y qué no.

Finalmente habrá una interfaz de exploración para los usuarios. Aquí se podrá
"visitar" los diferentes terminales para encontrar los recorridos que ocurren
cuando se revise la información, buscar rutas a ubicaciones en específico,
consultar información referente a tarifas y tiempos de viaje, entre otras
acciones.

# Situación actual
Ahora mismo el proyecto está solo comenzando, es por ello que la mayoría de las
cosas no estarán disponibles, sino que el avance que se realice va a funcionar
como un prototipo. La idea es continuar trabajando con este proyecto, ya que
parece una solución que puede aportar un valor a las personas que frecuentan
los buses para realizar sus viajes.

# Instrucciones de uso
El proyecto, al menos el prototipo, está organizado de tal modo que todos los
servicios se puedan ejecutar mediante Docker. Es por ello que es necesario
tener Docker instalado en nuestro sistema para poder ejecutar el proyecto,
además que debe ser un equipo Linux. En el caso de windows se puede usar docker
mediante wsl2. Con esas condiciones ya cumplidas, estos son los pasos
necesarios:

## Obtener el código del proyecto
Si solo se desea clonar la versión actual y ejecutar, se puede usar este comando.
```sh
git clone https://github.com/ruta-destino/ruta-destino.git
```
Si en caso contrario, se va a seguir desarrollando, hay que clonar usando
[ssh](https://docs.github.com/en/authentication/connecting-to-github-with-ssh).

```sh
git clone git@github.com:ruta-destino/ruta-destino.git
```

## Realizar configuraciones
Dentro de la carpeta del proyeto habrá un archivo en la raíz,
[`.env.sample`](.env.sample). Es necesario modificar ese archivo con los valores
apropiados, luego hay que guardarlo en la misma ubicación como `.env`.

## Ejecutar los servicios
Vienen incluidos dos archivos, [`docker-compose.yml`](docker-compose.yml) y
[`docker-compose-deploy.yml`](docker-compose-deploy.yml). El archivo deploy
está diseñado para usarse en producción, mientras que el otro archivo está
configurado para usarse en la etapa de desarrollo. La mayor parte del tiempo se
usará la versión de desarrollo, siempre se puede usar, incluso en un servidor.
Podemos ejecutar con el siguiente comando.
```sh
docker compose up
```
Si realmente necesitamos la versión de producción se puede usar esta versión.
```sh
docker compose --file docker-compose-deploy.yml up
```
Ambos comandos van a correr los servicios mientras el comando se ejecute, pero
al cancelar se detendrán y se apagarán. Para mantenerlos encendidos, se puede
agregar la opción -d al final del comando para que se ejecuten en segundo
plano.

## Actualizar el esquema de base de datos
Los comandos anteriores ya habrán descargado la versión correcta de PostgreSQL
y arrancado el servicio del backend. Sin embargo la base de datos no tendrá
ninguna de las tablas necesarias para la aplicación. Dentro de los servicios
hay una aplicación para aplicar migraciones de base de datos. Ya está vinculado
con las migraciones del proyecto, por lo que con este comando se puede aplicar
el esquema.
```sh
docker compose run --rm migrate -path=/migrations -database "postgres://rutadestino:password@db/rutadestino?sslmode=disable" up
```
Si se cambiaron las configuraciones en `.env` hay que usar esos valores para
conectarse a la base datos.

## Iniciar el frontend
Por el momento el frontend se debe iniciar por separado. Cuando se incorpore a
Docker estas instrucciones van a cambiar. Hay que entrar a la carpeta del
frontend que se desee ejecutar. Dentro de la carpeta tendremos que modificar el
archivo `.env.sample` con la ruta de la api. Si no se cambió esa información en
el archivo `docker-compose.yml`, la variable que está en el archivo debería ser
la correcta. Solo hay que guardar el archivo en esa ruta como `.env`.

Después es necesario instalar las dependencias que usa el frontend. Para eso
usamos el siguiente comando:
```sh
npm install
```
Luego para ejecutar el servidor de desarrollo del frontend usamos este comando.
```sh
npm run dev
```
Si queremos usar un puerto diferente, lo hacemos así.
```sh
npm run dev -- --port 9000
```
