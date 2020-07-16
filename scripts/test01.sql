CREATE TABLE empleados(
    nombre VARCHAR(10)	NOT NULL,
    apellido VARCHAR(10) NOT NULL,
    nss INT NOT NULL PRIMARY KEY,
    fecha_nac DATE NOT NULL,
    direccion VARCHAR NOT NULL,
    sexo VARCHAR(2) CHECK(SEXO='F' OR sexo='M'), 
    salario NUMERIC NOT NULL,
    nss_sup INT REFERENCES empleados(nss),
    num_dep SMALLINT NOT NULL 
);

CREATE TABLE departamentos(
    nom_dep VARCHAR(20) NOT NULL,
    num_dep SMALLINT NOT NULL PRIMARY KEY,
    nss_gte INT REFERENCES empleados(nss),
    fecha_ini_gte DATE NOT NULL
);

CREATE TABLE lugares_departamento(
    num_dep SMALLINT NOT NULL REFERENCES departamentos(num_dep),
    lugar_dep VARCHAR NOT NULL,
    PRIMARY KEY(num_dep, lugar_dep)
);

ALTER TABLE empleados
    ADD CONSTRAINT constraint_fk
    FOREIGN KEY (num_dep) REFERENCES departamentos(num_dep);
