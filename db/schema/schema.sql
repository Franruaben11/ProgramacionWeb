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

-- Tabla de catálogo para las actividades
CREATE TABLE Actividades (
    id_actividad SERIAL PRIMARY KEY,
    nombre_actividad VARCHAR(100) NOT NULL,
    descripcion TEXT
);

-- 2. Avisos (Conecta directamente un enfermero con un paciente)

CREATE TABLE Avisos (
    id_aviso SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion TEXT,
    id_actividad INT REFERENCES Actividades(id_actividad) ON DELETE CASCADE,
    id_paciente INT REFERENCES Pacientes(id_paciente) ON DELETE CASCADE,
    id_enfermero INT REFERENCES Enfermeros(id_enfermero) ON DELETE CASCADE
);

