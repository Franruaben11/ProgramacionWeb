-- name: GetActividad :one
SELECT * FROM Actividades WHERE id_actividad = $1;

-- name: ListActividades :many
SELECT * FROM Actividades ORDER BY nombre_actividad;

-- name: CreateActividad :one
INSERT INTO Actividades (nombre_actividad, descripcion) VALUES ($1, $2) RETURNING *;

-- name: UpdateActividad :exec
UPDATE Actividades SET nombre_actividad = $2 WHERE id_actividad = $1;

-- name: DeleteActividad :exec
DELETE FROM Actividades WHERE id_actividad = $1;

-- name: ListActividadesPorPaciente :many
SELECT a.* FROM Actividades a 
JOIN Paciente_Actividad pa ON a.id_actividad = pa.id_actividad 
WHERE pa.id_paciente = $1;
