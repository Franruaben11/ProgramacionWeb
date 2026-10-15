-- name: GetAviso :one
SELECT * FROM Avisos WHERE id_aviso = $1;

-- name: ListAvisos :many
SELECT * FROM Avisos ORDER BY nombre;

-- name: CreateAviso :one
INSERT INTO Avisos (nombre, descripcion, id_paciente, id_enfermero) 
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateAviso :exec
UPDATE Avisos SET nombre = $2, descripcion = $3, id_paciente = $4, id_enfermero = $5 WHERE id_aviso = $1;

-- name: DeleteAviso :exec
DELETE FROM Avisos WHERE id_aviso = $1;