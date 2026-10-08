-- name: GetPaciente :one
SELECT * FROM Pacientes WHERE id_paciente = $1;


-- name: ListPacientes :many
SELECT * FROM Pacientes ORDER BY nombre;


-- name: CreatePaciente :one
INSERT INTO Pacientes (nombre) VALUES ($1) RETURNING *;


-- name: UpdatePaciente :exec
UPDATE Pacientes SET nombre = $2 WHERE id_paciente = $1;


-- name: DeletePaciente :exec
DELETE FROM Pacientes WHERE id_paciente = $1;


-- name: ListPacientesPorEnfermero :many
SELECT p.* FROM Pacientes p
JOIN Enfermeros ep ON p.id_paciente = ep.id_paciente
WHERE ep.id_enfermero = $1;


-- name: ListPacientesPorFamiliar :many
SELECT p.* FROM Pacientes p
JOIN Familiares fp ON p.id_paciente = fp.id_paciente
WHERE fp.id_familiar = $1;
