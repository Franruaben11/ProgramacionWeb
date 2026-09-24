-- name: GetEnfermero :one
SELECT * FROM Enfermeros WHERE id_enfermero = $1;


-- name: ListEnfermeros :many
SELECT * FROM Enfermeros ORDER BY nombre;


-- name: CreateEnfermero :one
INSERT INTO Enfermeros (nombre, contrasena) VALUES ($1, $2) RETURNING *;


-- name: UpdateEnfermero :exec
UPDATE Enfermeros SET nombre = $2, contrasena = $3 WHERE id_enfermero = $1;


-- name: DeleteEnfermero :exec
DELETE FROM Enfermeros WHERE id_enfermero = $1;
