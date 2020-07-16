package models

//LugarDepartamento .
type LugarDepartamento struct {
	tableName struct{} `pg:"lugares_departamento"`
	NumDep    int      `pg:"num_dep" json:"num_dep" binding:"required"`
	LugarDep  string   `pg:"lugar_dep" json:"lugar_dep" binding:"required"`
}
