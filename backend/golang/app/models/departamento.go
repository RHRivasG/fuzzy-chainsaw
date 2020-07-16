package models

import "time"

//Departamento .
type Departamento struct {
	tableName   struct{}  `pg:"departamentos"`
	NomDep      string    `pg:"nom_dep, notnull" json:"nom_dep" binding:"required"`
	NumDep      int       `pg:"num_dep,pk" json:"num_dep" binding:"required"`
	NssGte      int       `pg:"nss_gte" json:"nss_gte"`
	FechaIniGte time.Time `pg:"fecha_ini_gte, notnull" json:"fecha_ini_gte" binding:"required"`
}
