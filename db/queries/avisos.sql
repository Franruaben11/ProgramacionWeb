-- name: GetAviso :one
SELECT * FROM Avisos WHERE id_aviso = $1;

-- name: ListAvisos :many
SELECT * FROM Avisos ORDER BY nombre;

-- name: CreateAviso :one
INSERT INTO avisos (nombre, descripcion, fecha_hora_inicio, id_paciente, id_enfermero)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: DeleteAviso :exec
DELETE FROM Avisos WHERE id_aviso = $1;