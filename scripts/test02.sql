INSERT INTO departamentos(nom_dep, num_dep, fecha_ini_gte)
    VALUES('investigacion', 5, '1978-05-22');
INSERT INTO departamentos(nom_dep, num_dep, fecha_ini_gte)
    VALUES('administracion',4,'1985-01-01');
INSERT INTO departamentos(nom_dep, num_dep, fecha_ini_gte)
    VALUES('direccion',1,'1971-06-19');

INSERT INTO empleados(nombre, apellido, nss, fecha_nac, direccion, sexo, salario, num_dep)
    VALUES ('Jaime', 'Botello', 888665555, '1927-11-10','Boleita','M',55000,1);
INSERT INTO empleados
    VALUES ('Jazmin', 'Valdes', 987654321, '1931-06-20','Tazon','F',43000,888665555,4);
INSERT INTO empleados
    VALUES ('Federico', 'Vizcarra', 333445555, '1945-12-08','Carrizal','M',40000,888665555,5);

UPDATE departamentos
    SET nss_gte= 333445555
    WHERE num_dep=5;
UPDATE departamentos
    SET nss_gte= 987654321
    WHERE num_dep=4;
UPDATE departamentos
    SET nss_gte= 888665555
    WHERE num_dep=1;

INSERT INTO lugares_departamento
    VALUES(1,'El Marques');
INSERT INTO lugares_departamento
    VALUES(4, 'Carrizal');
INSERT INTO lugares_departamento
    VALUES(5,'Los Chorros');
INSERT INTO lugares_departamento
    VALUES(5, 'Boleita');
INSERT INTO lugares_departamento
    VALUES(5,'El Marques');

