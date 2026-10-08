-- name: GetFamiliar :one
SELECT * FROM Familiares WHERE id_familiar = $1;


-- name: ListFamiliares :many
SELECT * FROM Familiares ORDER BY nombre;


-- name: CreateFamiliar :one
INSERT INTO Familiares (nombre, contrasena) VALUES ($1, $2) RETURNING *;


-- name: UpdateFamiliar :exec
UPDATE Familiares SET nombre = $2, contrasena = $3 WHERE id_familiar = $1;


-- name: DeleteFamiliar :exec
DELETE FROM Familiares WHERE id_familiar = $1;
