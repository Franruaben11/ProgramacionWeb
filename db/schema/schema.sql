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
    nombre_actividad VARCHAR(100) NOT NULL
);

-- 2. Avisos (Conecta directamente un enfermero con un paciente)

CREATE TABLE Avisos (
    id_aviso SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion TEXT,
    id_paciente INT REFERENCES Pacientes(id_paciente) ON DELETE CASCADE,
    id_enfermero INT REFERENCES Enfermeros(id_enfermero) ON DELETE CASCADE
);

-- 3. Tablas Intermedias (Para resolver las "listas")

-- Relación: Enfermeros atienden a varios pacientes
CREATE TABLE Enfermero_Paciente (
    id_enfermero INT REFERENCES Enfermeros(id_enfermero) ON DELETE CASCADE,
    id_paciente INT REFERENCES Pacientes(id_paciente) ON DELETE CASCADE,
    PRIMARY KEY (id_enfermero, id_paciente)
);

-- Relación: Familiares tienen a cargo varios pacientes (o un paciente tiene varios familiares)
CREATE TABLE Familiar_Paciente (
    id_familiar INT REFERENCES Familiares(id_familiar) ON DELETE CASCADE,
    id_paciente INT REFERENCES Pacientes(id_paciente) ON DELETE CASCADE,
    PRIMARY KEY (id_familiar, id_paciente)
);

-- Relación: Pacientes tienen asignadas varias actividades
CREATE TABLE Paciente_Actividad (
    id_paciente INT REFERENCES Pacientes(id_paciente) ON DELETE CASCADE,
    id_actividad INT REFERENCES Actividades(id_actividad) ON DELETE CASCADE,
    PRIMARY KEY (id_paciente, id_actividad)
);