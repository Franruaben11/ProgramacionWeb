# Trabajo Web - La Casa del Ángel

## Descripción del proyecto

La Casa del Ángel es una aplicación web para gestionar y comunicar la
información cotidiana de una residencia geriátrica.

El sistema está pensado para el personal de enfermería y los familiares de los
pacientes. En una residencia, la información sobre las actividades diarias
puede quedar distribuida en distintos medios o comunicarse de manera informal.
Esto dificulta el seguimiento de cada paciente y puede provocar confusiones
entre las personas responsables de su cuidado y sus familiares.

La aplicación centraliza en un único lugar los datos de pacientes, enfermeros,
familiares y avisos. Los enfermeros pueden registrar actividades con su
descripción y horario, asociándolas al paciente correspondiente. Los familiares
pueden consultar esa información para conocer las actividades realizadas
durante el día.

De esta manera, el proyecto busca mejorar la organización de la residencia,
facilitar el seguimiento de cada paciente y favorecer una comunicación más
clara entre el personal de cuidado y sus familiares.

## Requisitos

- Git
- Docker y Docker Compose
- Go
- GNU Make
- Bash y curl

## Preparar el proyecto

Clonar el repositorio y entrar en la carpeta del proyecto:

```bash
git clone https://github.com/Franruaben11/ProgramacionWeb.git
cd ProgramacionWeb
```

Para actualizar una copia existente:

```bash
git pull
```

## Comandos disponibles

```bash
make generate    # Genera el código Go de sqlc usando Docker
make generate-templ # Genera las vistas SSR a partir de archivos .templ
make up          # Genera sqlc, construye y levanta la aplicación
make down        # Detiene y elimina los contenedores
make logs        # Muestra los logs de los contenedores
make db-up       # Levanta solamente PostgreSQL
make build       # Compila el proyecto Go
make test-api    # Ejecuta las pruebas curl de la API y limpia el entorno
make test        # Ejecuta el flujo completo de pruebas de la API
make test-reset  # Recrea PostgreSQL y ejecuta el flujo completo de pruebas
make fmt         # Formatea los archivos Go
```

## Pruebas completas

El comando `make test` realiza automáticamente todas las tareas necesarias:

1. Genera el código de `sqlc`.
2. Compila el proyecto.
3. Elimina contenedores, redes y volúmenes anteriores.
4. Levanta PostgreSQL y la aplicación con Docker Compose.
5. Espera a que PostgreSQL y la API estén disponibles.
6. Ejecuta las pruebas de los endpoints mediante `requests.sh` y curl.
7. Elimina contenedores, redes y volúmenes al finalizar, incluso si una prueba falla.

Las pruebas cubren las operaciones válidas disponibles para:

- Enfermeros: POST, GET general, GET por ID, PUT y DELETE.
- Pacientes: POST, GET general, GET por ID, PUT y DELETE.
- Familiares: POST, GET general, GET por ID, PUT y DELETE.
- Avisos: POST, GET general, GET por ID y DELETE.

Además de comprobar que las peticiones finalicen correctamente, `requests.sh`
valida explícitamente:

- el código HTTP esperado para cada operación;
- atributos relevantes del JSON recibido;
- los identificadores devueltos por las operaciones de creación;
- que las actualizaciones queden persistidas al consultar nuevamente el recurso;
- que una eliminación devuelva `204` y que el recurso deje de estar disponible;
- errores por JSON inválido, campos obligatorios faltantes, fechas inválidas,
  métodos no permitidos e identificadores inexistentes.

Los identificadores se obtienen de las respuestas de creación, por lo que las
pruebas no dependen de que la base de datos comience a numerar registros desde
`1`.

Para ejecutar solamente el flujo de pruebas de la API:

```bash
make test-api
```

> Estos comandos eliminan el volumen de PostgreSQL y todos los datos guardados para comenzar con una base limpia.

## Ejecutar para desarrollo

Para levantar el entorno y dejarlo ejecutándose:

```bash
make up
```

La aplicación queda disponible en [http://localhost:8080](http://localhost:8080) y PostgreSQL en el puerto `5432`.

La Fase 2 agrega vistas SSR con templ bajo `/web` (inicio, listados y
formularios de pacientes, enfermeros, familiares y avisos). La API JSON se
mantiene separada bajo `/api/...` y el frontend TP4 continúa disponible en
`/static/index.html` (servido como `/`).

## Arquitectura y rutas

- **API JSON:** `/api/pacientes`, `/api/enfermeros`, `/api/familiares` y
  `/api/avisos`, incluyendo sus rutas por ID. El frontend TP4 consume estas
  rutas con `fetch`.
- **Compatibilidad:** las rutas antiguas sin `/api` continúan disponibles
  durante la transición.
- **Interfaz SSR:** `/web`, `/web/pacientes`, `/web/enfermeros`,
  `/web/familiares` y `/web/avisos`. Los formularios de alta se encuentran en
  `/web/<entidad>/nuevo` y redirigen con `303 See Other` después de crear.
- **Frontend TP4:** `/` sirve el panel estático que carga los registros en el
  navegador y permite crear y eliminar datos sin recargar la página.

La API y las páginas SSR comparten la misma capa de datos generada por sqlc.
Las vistas SSR se generan desde archivos `.templ` con:

```bash
make generate-templ
```

La autenticación no forma parte de esta versión; el objetivo actual es
demostrar persistencia, consumo de API, manipulación del DOM y renderizado
server-side.

También se puede iniciar la aplicación directamente con Go desde la raíz del proyecto:

```bash
go run .
```

En ese caso, PostgreSQL debe estar disponible localmente, por ejemplo mediante:

```bash
make db-up
```

## Funcionamiento principal

El flujo de uso de la aplicación es el siguiente:

1. Se registran los pacientes que viven en la residencia.
2. Se registran los enfermeros responsables de su cuidado.
3. Se registran los familiares vinculados a los pacientes.
4. Un enfermero crea un aviso para registrar una actividad o situación
   relacionada con un paciente.
5. El aviso incluye una descripción, un horario, el paciente involucrado y el
   enfermero responsable.
6. La información puede consultarse posteriormente para realizar el
   seguimiento diario del paciente.

## Entidades principales

- **Paciente:** persona residente de la institución. Se registra su nombre y se
  utiliza como referencia para asociar actividades y avisos.
- **Enfermero:** integrante del personal de enfermería que registra
  información sobre las actividades realizadas para un paciente.
- **Familiar:** persona vinculada a un paciente que necesita consultar
  información sobre sus actividades diarias.
- **Aviso:** actividad o comunicación asociada a un paciente y registrada por un
  enfermero. Incluye nombre, descripción, fecha y horario, paciente relacionado
  y enfermero responsable.
