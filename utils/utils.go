package utils

import (
	"database/sql"
	"encoding/json"
	"indodev/models"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

func DBTransaction(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Rollback Panic
		} else if err != nil {
			tx.Rollback() // err is not nill
		} else {
			err = tx.Commit() // err is nil
		}
	}()
	err = txFunc(tx)
	return err
}

func Stringify(input interface{}) string {
	bytes, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	strings := string(bytes)
	bytes, err = json.Marshal(strings)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func JSONPrettyfy(data interface{}) string {
	bytesData, _ := json.MarshalIndent(data, "", "  ")
	return string(bytesData)
}

func ToString(i interface{}) string {
	log, _ := json.Marshal(i)
	logString := string(log)

	return logString
}

func ResponseJSON(success bool, code string, msg string, result interface{}) models.Response {
	tm := time.Now()
	response := models.Response{
		Success:          success,
		StatusCode:       code,
		Result:           result,
		Message:          msg,
		ResponseDatetime: tm,
	}

	return response
}

func BindValidateStruct(ctx echo.Context, i interface{}) error {
	if err := ctx.Bind(i); err != nil {
		return err
	}

	if err := ctx.Validate(i); err != nil {
		return err
	}

	return nil
}

func TimeStampNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func DateNow() string {
	return time.Now().Format("2006-01-02")
}

func BuildHierarchy(flatData []models.ResponseDivision) []*models.ResponseDivision {
	divisionMap := make(map[string]*models.ResponseDivision)
	var roots []*models.ResponseDivision

	for i := range flatData {
		division := &flatData[i]
		division.Children = []*models.ResponseDivision{}
		divisionMap[division.DivisionCode] = division
	}

	for _, division := range flatData {
		if division.DivisionParentCode == "" {
			roots = append(roots, divisionMap[division.DivisionCode])
		} else {
			if parent, exists := divisionMap[division.DivisionParentCode]; exists {
				parent.Children = append(parent.Children, divisionMap[division.DivisionCode])
			}
		}
	}

	return roots
}
