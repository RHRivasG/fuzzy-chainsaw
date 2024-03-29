package services

import (
	"test-api/app/models"

	"github.com/go-pg/pg"
)

//EmpleadoServiceI .
type EmpleadoServiceI interface {
	FindAll(db *pg.DB) ([]models.Empleado, error)
}

//EmpleadoService .
type EmpleadoService struct{}

//FindAll empleados
func (s *EmpleadoService) FindAll(db *pg.DB) ([]models.Empleado, error) {
	var empleados []models.Empleado
	if err := db.Model(&models.Empleado{}).Select(&empleados); err != nil {
		return nil, err
	}
	return empleados, nil
}

//Create empleado
func (s *EmpleadoService) Create(empleado models.Empleado, db *pg.DB) error {
	return db.Insert(&empleado)
}

//Update empleado
func (s *EmpleadoService) Update(empleado models.Empleado, db *pg.DB) error {
	_, err := db.Model(&empleado).Where("nss = ?", &empleado.Nss).Update()
	return err
}
