-- 1. Tablas Principales (Entidades independientes)

CREATE TABLE Enfermeros (
    id_enfermero SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    contrasena VARCHAR(255) NOT NULL
);

CREATE TABLE Familiares (
    id_familiar SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    contrasena VARCHAR(255) NOT NULL
);

CREATE TABLE Pacientes (
    id_paciente SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL
);

-- 2. Avisos (Conecta directamente un enfermero con un paciente)

CREATE TABLE Avisos (
    id_aviso SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion TEXT NOT NULL,
    fecha_hora_inicio TIMESTAMP NOT NULL,
    fecha_hora_fin TIMESTAMP NOT NULL,
    id_paciente INT NOT NULL REFERENCES Pacientes(id_paciente) ON DELETE CASCADE,
    id_enfermero INT NOT NULL REFERENCES Enfermeros(id_enfermero) ON DELETE CASCADE
);