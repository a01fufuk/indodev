package repositories

import "indodev/models"

type DivisionRepository interface {
	FindDivisionList() ([]models.Division, error)
	IsDivisionExistsByIndex(div models.Division) (models.Division, bool, error)
	FindDivisionByHirarki(div models.Division) ([]models.Division, error)
	AddDivision(div models.Division) (ID int64, err error)
	EditDivision(div models.Division) error
	RemoveDivision(id int) error
}
