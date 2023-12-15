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
