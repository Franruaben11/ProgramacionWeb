-- getters

-- name: GetEnfermero :one
SELECT * FROM Enfermeros WHERE id_enfermero = $1;

-- name: GetPaciente :one
SELECT * FROM Pacientes WHERE id_paciente = $1;

-- name: GetFamiliar :one
SELECT * FROM Familiares WHERE id_familiar = $1;

-- name: GetActividad :one
SELECT * FROM Actividades WHERE id_actividad = $1;

-- name: GetAviso :one
SELECT * FROM Avisos WHERE id_aviso = $1;


-- listers básicos

-- name: ListEnfermeros :many
SELECT * FROM Enfermeros ORDER BY nombre;

-- name: ListPacientes :many
SELECT * FROM Pacientes ORDER BY nombre;

-- name: ListFamiliares :many
SELECT * FROM Familiares ORDER BY nombre;

-- name: ListActividades :many
SELECT * FROM Actividades ORDER BY nombre_actividad;

-- name: ListAvisos :many
SELECT * FROM Avisos ORDER BY nombre;


-- listers relacionales (Acá es donde usamos los JOIN con las tablas puente)

-- name: ListActividadesPorPaciente :many
SELECT a.* FROM Actividades a 
JOIN Paciente_Actividad pa ON a.id_actividad = pa.id_actividad 
WHERE pa.id_paciente = $1;

-- name: ListPacientesPorEnfermero :many
SELECT p.* FROM Pacientes p
JOIN Enfermero_Paciente ep ON p.id_paciente = ep.id_paciente
WHERE ep.id_enfermero = $1;

-- name: ListPacientesPorFamiliar :many
SELECT p.* FROM Pacientes p
JOIN Familiar_Paciente fp ON p.id_paciente = fp.id_paciente
WHERE fp.id_familiar = $1;


-- create

-- name: CreateEnfermero :one
INSERT INTO Enfermeros (nombre, contrasena) VALUES ($1, $2) RETURNING *;

-- name: CreatePaciente :one
INSERT INTO Pacientes (nombre) VALUES ($1) RETURNING *;

-- name: CreateFamiliar :one
INSERT INTO Familiares (nombre, contrasena) VALUES ($1, $2) RETURNING *;

-- name: CreateActividad :one
INSERT INTO Actividades (nombre_actividad) VALUES ($1) RETURNING *;

-- name: CreateAviso :one
INSERT INTO Avisos (nombre, descripcion, id_paciente, id_enfermero) 
VALUES ($1, $2, $3, $4) RETURNING *;


-- creates de relaciones (¡NUEVOS! Para llenar las tablas puente)

-- name: AsignarActividadAPaciente :exec
INSERT INTO Paciente_Actividad (id_paciente, id_actividad) VALUES ($1, $2);

-- name: AsignarPacienteAEnfermero :exec
INSERT INTO Enfermero_Paciente (id_enfermero, id_paciente) VALUES ($1, $2);

-- name: AsignarPacienteAFamiliar :exec
INSERT INTO Familiar_Paciente (id_familiar, id_paciente) VALUES ($1, $2);


-- update

-- name: UpdateEnfermero :exec
UPDATE Enfermeros SET nombre = $2, contrasena = $3 WHERE id_enfermero = $1;

-- name: UpdatePaciente :exec
UPDATE Pacientes SET nombre = $2 WHERE id_paciente = $1;

-- name: UpdateFamiliar :exec
UPDATE Familiares SET nombre = $2, contrasena = $3 WHERE id_familiar = $1;

-- name: UpdateActividad :exec
UPDATE Actividades SET nombre_actividad = $2 WHERE id_actividad = $1;

-- name: UpdateAviso :exec
UPDATE Avisos SET nombre = $2, descripcion = $3, id_paciente = $4, id_enfermero = $5 WHERE id_aviso = $1;


-- delete

-- name: DeleteEnfermero :exec
DELETE FROM Enfermeros WHERE id_enfermero = $1;

-- name: DeletePaciente :exec
DELETE FROM Pacientes WHERE id_paciente = $1;

-- name: DeleteFamiliar :exec
DELETE FROM Familiares WHERE id_familiar = $1;

-- name: DeleteActividad :exec
DELETE FROM Actividades WHERE id_actividad = $1;

-- name: DeleteAviso :exec
DELETE FROM Avisos WHERE id_aviso = $1;