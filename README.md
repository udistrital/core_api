# core_api

Descripción

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
CORE_API_DB_USER=[usuario]
CORE_API_DB_PASS=[password del usuario]
CORE_API_DB_URL=[url de bd]
CORE_API_DB_NAME=[nombre de bd]
CORE_API_DB_SCHEMA=[esquema de bd]
CORE_API_HTTP_PORT=[puerto]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con CORE_API_DB_...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/core_api

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/core_api

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
CORE_API_HTTP_PORT=8080 CORE_API_DB_URL=127.0.0.1:27017 CORE_API_DB_SOME_VARIABLE=some_value bee run
```
### Ejecución Dockerfile
```shell
# docker build --tag=core_api . --no-cache
# docker run -p 80:80 core_api
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/core_api

#2. Moverse a la carpeta del repositorio
cd core_api

#3. Crear un fichero con el nombre **custom.env**
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# En Proceso
```
## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/core_api/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/core_api) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/core_api/status.svg?ref=refs/heads/release)](https://hubci.portaloas.udistrital.edu.co/udistrital/core_api) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/core_api/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/core_api) |


## Licencia

This file is part of core_api.

core_api is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

core_api is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with core_api. If not, see https://www.gnu.org/licenses/.

