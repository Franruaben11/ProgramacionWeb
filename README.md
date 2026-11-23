# Trabajo Web - La Casa del Ángel

## Descripción del proyecto

La Casa del Ángel es una aplicación web para gestionar y comunicar la
información cotidiana de una residencia geriátrica.

El sistema centraliza los datos de pacientes, enfermeros, familiares y avisos.
Los enfermeros pueden registrar actividades asociadas a un paciente, incluyendo
su descripción y fecha y hora de realización. Los familiares pueden consultar
esta información para realizar el seguimiento diario de cada paciente.

El proyecto utiliza Go, plantillas `templ`, PostgreSQL y sqlc.

## Tecnologías utilizadas

- Go
- Templ
- HTMX
- PostgreSQL
- sqlc
- Docker y Docker Compose
- GNU Make
- Bash y curl

## Requisitos

Para ejecutar el proyecto se necesitan:

- Git
- Go
- Docker y Docker Compose
- GNU Make
- Bash
- curl

## Clonar el repositorio

```bash
git clone https://github.com/Franruaben11/ProgramacionWeb.git
cd ProgramacionWeb
git checkout tp6
```

## Ejecutar la aplicación

El proyecto puede iniciarse con:

```bash
make run-app
```

También puede ejecutarse directamente:

```bash
./run-app.sh
```

El script realiza automáticamente las siguientes tareas:

1. Genera las vistas `templ`.
2. Ejecuta sqlc para generar el código de acceso a la base de datos.
3. Compila el proyecto.
4. Elimina contenedores y volúmenes anteriores.
5. Levanta PostgreSQL y la aplicación con Docker Compose.
6. Espera a que PostgreSQL y el servidor estén disponibles.
7. Mantiene la aplicación ejecutándose hasta que se presione ENTER.
8. Detiene los contenedores y elimina los volúmenes al finalizar.

Una vez iniciada, la aplicación estará disponible en:

- http://localhost:8080
- http://localhost:8080/pacientes
- http://localhost:8080/enfermeros
- http://localhost:8080/familiares
- http://localhost:8080/avisos

## Comandos disponibles

```bash
make run-app        # Ejecuta todo el flujo de inicio y limpieza
make generate       # Genera las consultas sqlc y las vistas templ
make generate-templ # Genera las vistas templ
make up             # Levanta PostgreSQL y la aplicación
make down           # Detiene los contenedores
make logs           # Muestra los logs de Docker Compose
make db-up          # Levanta solamente PostgreSQL
make db-reset       # Recrea PostgreSQL y elimina sus datos
make build          # Compila el proyecto Go
make test-api       # Ejecuta las pruebas de la API y limpia el entorno
make test           # Ejecuta el flujo completo de pruebas
make test-reset     # Recrea PostgreSQL y ejecuta todas las pruebas
make fmt            # Formatea los archivos Go
```

## Interfaces dinámicas con HTMX

La interfaz utiliza HTMX para crear y eliminar registros sin recargar
completamente la página.

### Crear registros

Cada formulario utiliza:

- `hx-post` para enviar el formulario al endpoint correspondiente.
- `hx-target` para indicar el contenedor que debe actualizarse.
- `hx-swap="outerHTML"` para reemplazar completamente la lista existente.

Después de crear el registro, el servidor vuelve a consultar la base de datos
y devuelve solamente el componente actualizado de la lista.

Las operaciones están disponibles en:

- `POST /pacientes`
- `POST /enfermeros`
- `POST /familiares`
- `POST /avisos`

### Eliminar registros

Cada elemento de las listas tiene un botón con:

- `hx-delete` para solicitar la eliminación.
- `hx-target="closest tr"` para seleccionar la fila correspondiente.
- `hx-swap="outerHTML"` para eliminar la fila.
- `hx-confirm` para solicitar confirmación al usuario.

Las rutas de eliminación son:

- `DELETE /pacientes/{id}`
- `DELETE /enfermeros/{id}`
- `DELETE /familiares/{id}`
- `DELETE /avisos/{id}`

Los handlers de eliminación llaman al método correspondiente generado por
sqlc y devuelven una respuesta vacía con estado `200 OK`.

HTMX carga el contenido generado por las plantillas `templ` desde el servidor.
No utiliza JavaScript escrito en el proyecto ni realiza recargas completas al
crear o eliminar registros.

## Arquitectura y rutas principales

### Interfaz SSR con HTMX

- `/`: página de inicio.
- `/pacientes`: listado y formulario de pacientes.
- `/enfermeros`: listado y formulario de enfermeros.
- `/familiares`: listado y formulario de familiares.
- `/avisos`: listado y formulario de avisos.

### API JSON

- `/api/pacientes`
- `/api/enfermeros`
- `/api/familiares`
- `/api/avisos`

Las rutas bajo `/api` permiten realizar operaciones CRUD mediante respuestas
JSON y se mantienen como complemento de la interfaz SSR.

Las rutas JSON reciben y devuelven datos en formato JSON, mientras que las
rutas SSR devuelven fragmentos HTML utilizados por HTMX o páginas completas
generadas con `templ`.

## Base de datos

PostgreSQL se ejecuta en Docker Compose. Las consultas se escriben en archivos
SQL dentro de `db/queries/` y sqlc genera las operaciones de Go en
`db/sqlc/`.

El esquema se encuentra en:

```text
db/schema/schema.sql
```

Las credenciales locales utilizadas por Docker Compose son:

- Usuario: `user`
- Contraseña: `password`
- Base de datos: `mydatabase`
- Puerto PostgreSQL: `5432`

> El volumen de PostgreSQL se elimina al ejecutar `docker compose down -v`.
> Esto elimina todos los datos persistidos.

## Ejecutar sin Docker Compose para desarrollo

Es posible levantar primero PostgreSQL:

```bash
make db-up
```

Luego iniciar la aplicación desde el directorio del proyecto:

```bash
go run .
```

La aplicación estará disponible en http://localhost:8080.

## Pruebas

El comando:

```bash
make test
```

realiza automáticamente el siguiente flujo:

1. Genera las vistas y el código sqlc.
2. Compila el proyecto.
3. Elimina contenedores y volúmenes anteriores.
4. Levanta PostgreSQL y la aplicación.
5. Espera a que ambos servicios estén disponibles.
6. Ejecuta las pruebas HTTP definidas en `requests.sh`.
7. Elimina contenedores y volúmenes al finalizar.

Para ejecutar solamente las pruebas de la API:

```bash
make test-api
```

## Entidades principales

- **Paciente:** persona residente de la institución.
- **Enfermero:** integrante del personal de enfermería responsable del cuidado.
- **Familiar:** persona vinculada a un paciente.
- **Aviso:** actividad o comunicación registrada por un enfermero para un
  paciente. Incluye título, descripción, fecha y hora, paciente asociado y
  enfermero responsable.