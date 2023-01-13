# Sistema basico puntos usuario exagonal y cqrs

> Implementación de un ejemplo de arquitectura exagonal y eventos con nats cqrs.

## Instalacion

OS X & Linux:

```sh
git clone https://github.com/ANDERSON1808/sistema_puntos_exagonal_cqrs.git
```

```sh
cd microConsultaPuntos
go mod tidy
```

```sh
cd microMutacionPuntos
go mod tidy
```

## Correrlo local

OS X & Linux:

- Tener en cuenta las url de nats y dynamodb

```sh
cd microMutacionPuntos
go run main.go
```

```sh
cd microConsultaPuntos
go run main.go
```

## Compilar y correr con docker-compose

```sh
docker-compose up -d  --build
```

## API Endpoint

|              Endpoint               | HTTP Method |     Description      |
|:-----------------------------------:|:-----------:|:--------------------:|
|    `/api/v1/puntos/acumulacion`     |   `POST`    |      `Created`       |
| `/api/v1/puntos/consulta/{:userId}` |    `GET`    |  `Find user by ID`   |
|      `/api/v1/puntos/redimir`       |   `POST`    | `Created` |

## Modelo Json acumular puntos

```json
{
  "NombreUsuario": "Nombre usuario prueba",
  "PuntoId": 1,
  "UsuarioId": 1,
  "Punto": 10,
  "DetalleMovimiento": "detalle movimiento"
}
```

## Probar endpoints API using curl

`Request`

```bash
curl --location --request POST 'http://127.0.0.1:34963/api/v1/puntos/acumulacion' \
--header 'Content-Type: application/json' \
--data-raw '{
    "NombreUsuario":"Nombre usuario prueba",
    "PuntoId":2,
    "UsuarioId":1,
    "Punto":10,
    "DetalleMovimiento":"compra de algo"
}'
```

`Response`

```bash
HTTP/1.1 201 Created
Content-Type: application/json
Content-Length: 302
```

- #### Consulta puntos por usuarioId

`Request`

```bash
curl --location --request GET 'http://127.0.0.1:8874/api/v1/puntos/consulta?q=1'
```

`Response`

```bash
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 279
```

```json
{
  "UsuarioId": 1,
  "NombreUsuario": "Nombre usuario prueba",
  "TotalPuntos": "20"
}
```

## Modelo Json redimir puntos

```json
{
  "UsuarioId":10,
  "PuntoRedimir":5
}
```

## Probar endpoints API using curl

`Request`

```bash
curl --location --request POST 'http://127.0.0.1:34963/api/v1/puntos/redimir' \
--header 'Content-Type: application/json' \
--data-raw '{
   "UsuarioId":10,
   "PuntoRedimir":5
}'
```

`Response`

```bash
HTTP/1.1 201 Created
Content-Type: application/json
Content-Length: 302
```
