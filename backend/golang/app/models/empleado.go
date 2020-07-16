package models

import "time"

//Empleado .
type Empleado struct {
	tableName struct{}  `pg:"empleados"`
	Nombre    string    `pg:"nombre, notnull" json:"nombre" binding:"required"`
	Apellido  string    `pg:"apellido, notnull" json:"apellido" binding:"required"`
	Nss       int64     `pg:",pk" json:"nss" binding:"required"`
	FechaNac  time.Time `pg:"fecha_nac, notnull" json:"fecha_nac" binding:"required"`
	Direccion string    `pg:"direccion, notnull" json:"direccion" binding:"required"`
	Sexo      string    `pg:"sexo, notnull" json:"sexo" binding:"required"`
	Salario   int       `pg:"salario, notnull" json:"salario" binding:"required"`
	NssSup    int       `pg:"nss_sup" json:"nss_sup"`
	NumDep    int       `pg:"num_dep, notnull" json:"num_dep" binding:"required"`
}
