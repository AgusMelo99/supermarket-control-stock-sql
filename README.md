# Control de pedidos y stock para un supermercado online (PostgreSQL + Go)

Trabajo práctico grupal de Bases de Datos (UNGS) para el supermercado ficticio *ChangoOnline*. El sistema administra el ingreso, la cancelación y la entrega de pedidos, controla el stock de los productos y notifica a los clientes por mail. La lógica de negocio vive en la base de datos (stored procedures y triggers de PostgreSQL) y una CLI en Go se encarga de armar el esquema, cargar datos y ejecutar las operaciones.

> ✏️ **Completar:** mi aporte dentro del equipo (por ejemplo, qué stored procedures, qué archivos Go o qué parte del informe desarrollé yo).

## Qué incluye

- **Stored procedures** para el ciclo de vida del pedido: creación, agregado de productos, cierre, cancelación y entrega. Cada uno valida sus condiciones y, si falla, registra el motivo en una tabla de errores.
- **Reposición de stock:** un procedimiento diario genera solicitudes para los productos que están en o por debajo de su punto de reposición.
- **Triggers** que generan los mails al cliente cuando un pedido se acepta, se cancela o se entrega, y un mail de encuesta de satisfacción.
- **Concurrencia:** las operaciones que modifican stock corren dentro de transacciones. Se probó con `pg_sleep` y dos sesiones en paralelo.
- **Comparación con NoSQL:** una versión de los mismos datos en BoltDB (clave-valor con buckets).

## Decisiones técnicas

- **Nivel de aislamiento:** probamos `SERIALIZABLE`, pero daba resultados inconsistentes en las pruebas y agrega riesgo de deadlocks. Nos quedamos con el `READ COMMITTED` de PostgreSQL, que alcanzó para evitar condiciones de carrera en este caso.
- **Valores nulos en Go:** los structs usan `sql.NullString`, `sql.NullInt64` y `sql.NullFloat64`, porque varias columnas admiten `NULL`.
- **Un archivo por procedimiento:** cada stored procedure y trigger está en su propio `.sql` para mantener el código ordenado.

## Estructura

| Ruta | Contenido |
|---|---|
| `main.go` | Menú de la CLI |
| `bd/` | Creación de la base, tablas, claves primarias y foráneas |
| `datajson/` | Lectura de los JSON y carga de datos y operaciones |
| `stored-procedures/` | Stored procedures y triggers en SQL |
| `ver_tablas/` | Visualización de tablas por consola |
| `bolt-db/` | Versión NoSQL con BoltDB |

## Cómo ejecutarlo

**Requisitos:** Go y PostgreSQL.

1. Ajustar los datos de conexión a PostgreSQL en `main.go`. ✏️ *verificar dónde están definidos*
2. Correr la CLI:
   ```bash
   go run .
   ```
3. Elegir las opciones del menú en este orden:
   1. Crear la base de datos
   2. Crear las tablas
   3. Crear las claves primarias y foráneas
   4. Cargar los stored procedures (opción 5)
   5. Cargar los datos desde los JSON (opción 6)
   6. Ejecutar las operaciones sobre la base (opción 14)
   7. Generar las solicitudes de reposición (opción 15)

   Las opciones 7 a 13 muestran el contenido de las tablas.

## Informe completo

El informe del trabajo (diseño, dificultades y conclusiones) está en [`informe.adoc`](informe.adoc). ✏️ *Renombrar el `README.adoc` actual a `informe.adoc`*
