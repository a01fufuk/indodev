package divisionRepository

import (
	"database/sql"
	"errors"
	"indodev/constans"
	"indodev/models"
	"indodev/repositories"
	"indodev/utils"
	"log"
)

var defineColumn = ` division_code, division_name, division_hierarchy, division_parent_code `

type divisionRepository struct {
	RepoDB repositories.Repository
}

// NewTcsRepository
func NewDivisionRepository(repoDB repositories.Repository) divisionRepository {
	return divisionRepository{
		RepoDB: repoDB,
	}
}

func (ctx divisionRepository) FindDivisionList() ([]models.Division, error) {
	var result []models.Division
	var query = `SELECT id, ` + defineColumn + ` FROM division WHERE TRUE  ORDER BY division_hierarchy ASC`

	rows, err := ctx.RepoDB.DB.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	data, _ := divisionDto(rows)
	if len(data) == constans.EMPTY_VALUE_INT {
		return result, errors.New("tcs not found")
	}

	return data, nil
}

func (ctx divisionRepository) IsDivisionExistsByIndex(div models.Division) (models.Division, bool, error) {
	var query = `SELECT id, ` + defineColumn + ` FROM division WHERE TRUE `
	var args []interface{}
	var result models.Division

	if div.DivisionCode != constans.EMPTY_VALUE {
		query += ` AND division_code = ? `
		args = append(args, div.DivisionCode)
	}

	query += ` ORDER BY division_hierarchy ASC`
	query = utils.ReplaceSQL(query, "?")
	rows, err := ctx.RepoDB.DB.Query(query, args...)
	if err != nil {
		return result, constans.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := divisionDto(rows)
	if err != nil {
		return result, constans.FALSE_VALUE, err
	}

	if len(data) == constans.EMPTY_VALUE_INT {
		return result, constans.FALSE_VALUE, nil
	}

	return data[0], constans.TRUE_VALUE, nil
}

func (ctx divisionRepository) AddDivision(div models.Division) (ID int64, err error) {
	var args []interface{}
	args = append(args, div.DivisionCode, div.DivisionName, div.DivisionHierarchy, div.DivisionParentCode)

	query := `INSERT INTO division ( division_code, division_name, division_hierarchy, division_parent_code) VALUES (
			?, ?, ?, ?
		) RETURNING id`

	query = utils.ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, args...).Scan(&ID)

	if err != nil {
		return ID, err
	}

	return ID, nil
}

func (ctx divisionRepository) EditDivision(div models.Division) error {
	var err error

	log.Println(div.DivisionName)

	var query = `
	UPDATE division set  division_code = ?, division_name = ?, division_hierarchy = ?, division_parent_code = ?
	WHERE id = ?`

	query = utils.ReplaceSQL(query, "?")

	_, err = ctx.RepoDB.DB.Exec(query, div.DivisionCode, div.DivisionName, div.DivisionHierarchy, div.DivisionParentCode, div.ID)

	if err != nil {
		return err
	}

	return nil
}

func (ctx divisionRepository) RemoveDivision(id int) error {
	var err error

	_, err = ctx.RepoDB.DB.Exec("DELETE FROM division WHERE id = $1", id)

	if err != nil {
		return err
	}

	return err
}

func (ctx divisionRepository) FindDivisionByHirarki(div models.Division) ([]models.Division, error) {
	var args []interface{}
	var result []models.Division

	var query = `SELECT id, ` + defineColumn + ` FROM division WHERE TRUE `

	if div.DivisionHierarchy != constans.EMPTY_VALUE {
		query += ` AND division_hierarchy ILIKE '%'  || ? || '%' `
		args = append(args, div.DivisionHierarchy)
	}

	query = utils.ReplaceSQL(query, "?")
	rows, err := ctx.RepoDB.DB.Query(query, args...)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	data, _ := divisionDto(rows)
	if len(data) == constans.EMPTY_VALUE_INT {
		return result, nil
	}

	return data, nil
}

func divisionDto(rows *sql.Rows) ([]models.Division, error) {
	var result []models.Division
	for rows.Next() {
		var val models.Division
		err := rows.Scan(&val.ID, &val.DivisionCode, &val.DivisionName, &val.DivisionHierarchy, &val.DivisionParentCode)

		if err != nil {
			return result, err
		}
		result = append(result, val)
	}

	return result, nil
}
