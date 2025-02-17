package divisionService

import (
	"indodev/constans"
	"indodev/helpers"
	"indodev/models"
	"indodev/services"
	"indodev/utils"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type divisionService struct {
	Service services.UsecaseService
}

func NewDivisionService(service services.UsecaseService) divisionService {
	return divisionService{
		Service: service,
	}
}

func (svc divisionService) InsertDivision(ctx echo.Context) error {
	var result models.Response

	request := new(models.RequestCreateDivision)
	if err := helpers.BindValidateStruct(ctx, request); err != nil {
		log.Println("Error Validate InsertDivision : ", err.Error())
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	// CHECK DUPLICATE DIVISION CODE
	_, isExist, err := svc.Service.DivisionRepo.IsDivisionExistsByIndex(models.Division{
		DivisionCode: request.DivisionCode,
	})
	if err != nil {
		log.Println("EditDivision - Error IsDivisionExistsByIndex : ", err.Error())
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if isExist {
		log.Println("EditDivision - Error IsDivisionExistsByIndex : ", "Duplicate Division Code")
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, "Duplicate Division Code", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	// FIND PARENT DIV
	divisionHierarchy := ""
	if request.DivisionParentCode != constans.EMPTY_VALUE {
		parentDiv, isExist, err := svc.Service.DivisionRepo.IsDivisionExistsByIndex(models.Division{
			DivisionCode: request.DivisionParentCode,
		})
		if err != nil {
			log.Println("InsertDivision - Error IsDivisionExistsByIndex : ", err.Error())
			result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
			return ctx.JSON(http.StatusBadRequest, result)
		}

		if !isExist {
			log.Println("InsertDivision - Error IsDivisionExistsByIndex : ", "Parent Division Not Found")
			result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, "Parent Division Not Found", nil)
			return ctx.JSON(http.StatusBadRequest, result)
		}
		divisionHierarchy = parentDiv.DivisionHierarchy + "/" + request.DivisionCode
	} else {
		divisionHierarchy = request.DivisionCode
	}

	division := models.Division{
		DivisionCode:       request.DivisionCode,
		DivisionName:       request.DivisionName,
		DivisionHierarchy:  divisionHierarchy,
		DivisionParentCode: request.DivisionParentCode,
	}

	_, err = svc.Service.DivisionRepo.AddDivision(division)
	if err != nil {
		log.Println("InsertDivision - Error AddDivision : ", err.Error())
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Insert Request Has Been Submited Successfully")
	result = helpers.ResponseJSON(constans.TRUE_VALUE, constans.SUCCESS_CODE, constans.EMPTY_VALUE, nil)
	return ctx.JSON(http.StatusOK, result)
}

func (svc divisionService) EditDivision(ctx echo.Context) error {
	var result models.Response

	request := new(models.RequestEditDivision)
	if err := helpers.BindValidateStruct(ctx, request); err != nil {
		log.Println("Error Validate EditDivision : ", err.Error())
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	// CHECK IF EXIST
	divExist, isExist, err := svc.Service.DivisionRepo.IsDivisionExistsByIndex(models.Division{
		DivisionCode: request.DivisionCode,
	})
	if err != nil {
		log.Println("EditDivision - Error IsDivisionExistsByIndex : ", err.Error())
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !isExist {
		log.Println("EditDivision - Error IsDivisionExistsByIndex : ", "Division Not Exist")
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, "Division Not Exist", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	// FIND PARENT DIV

	division := models.Division{
		ID:                 divExist.ID,
		DivisionCode:       divExist.DivisionCode,
		DivisionName:       request.DivisionName,
		DivisionHierarchy:  divExist.DivisionHierarchy,
		DivisionParentCode: divExist.DivisionParentCode,
	}

	err = svc.Service.DivisionRepo.EditDivision(division)
	if err != nil {
		log.Println("EditDivision - Error AddDivision : ", err.Error())
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Edit Has Been Submited Successfully")
	result = helpers.ResponseJSON(constans.TRUE_VALUE, constans.SUCCESS_CODE, constans.EMPTY_VALUE, nil)
	return ctx.JSON(http.StatusOK, result)
}

func (svc divisionService) DeleteDivision(ctx echo.Context) error {
	var result models.Response

	request := new(models.RequestDeleteDivision)
	if err := helpers.BindValidateStruct(ctx, request); err != nil {
		log.Println("Error Validate DeleteDivision : ", err.Error())
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}

	// FIND DIVISION
	division, isExist, err := svc.Service.DivisionRepo.IsDivisionExistsByIndex(models.Division{
		DivisionCode: request.DivisionCode,
	})
	if err != nil {
		log.Println("DeleteDivision - Error IsDivisionExistsByIndex : ", err.Error())
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !isExist {
		log.Println("DeleteDivision - Error IsDivisionExistsByIndex : ", "Division Not Found")
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, "Division Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	fDiv, err := svc.Service.DivisionRepo.FindDivisionByHirarki(models.Division{
		DivisionHierarchy: division.DivisionHierarchy,
	})
	if err != nil {
		log.Println("DeleteDivision - Error FindDivisionByHirarki : ", err.Error())
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if len(fDiv) > 1 {
		log.Println("DeleteDivision - Error FindDivisionByHirarki : ", "Division Has Child")
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, "Division Has Child", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	err = svc.Service.DivisionRepo.RemoveDivision(division.ID)
	if err != nil {
		log.Println("DeleteDivision - Error AddDivision : ", err.Error())
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Delete Has Been Submited Successfully")
	result = helpers.ResponseJSON(constans.TRUE_VALUE, constans.SUCCESS_CODE, constans.EMPTY_VALUE, nil)
	return ctx.JSON(http.StatusOK, result)
}

func (svc divisionService) ListDivision(ctx echo.Context) error {
	var result models.Response

	// FIND DIVISION

	div, err := svc.Service.DivisionRepo.FindDivisionList()
	if err != nil {
		log.Println("DeleteDivision - Error AddDivision : ", err.Error())
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Delete Has Been Submited Successfully")
	result = helpers.ResponseJSON(constans.TRUE_VALUE, constans.SUCCESS_CODE, constans.EMPTY_VALUE, div)
	return ctx.JSON(http.StatusOK, result)
}

func (svc divisionService) ListDivisionHiearchy(ctx echo.Context) error {
	var result models.Response

	// FIND DIVISION
	div, err := svc.Service.DivisionRepo.FindDivisionList()
	if err != nil {
		log.Println("ListDivisionHiearchy - Error AddDivision : ", err.Error())
		result = helpers.ResponseJSON(constans.FALSE_VALUE, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var tmpDivision []models.ResponseDivision
	for _, dataDiv := range div {
		div := models.ResponseDivision{
			DivisionCode:       dataDiv.DivisionCode,
			DivisionName:       dataDiv.DivisionName,
			DivisionParentCode: dataDiv.DivisionParentCode,
			Children:           []*models.ResponseDivision{},
		}
		tmpDivision = append(tmpDivision, div)
	}

	resDiv := utils.BuildHierarchy(tmpDivision)
	log.Println("List Has Been Submited Successfully")
	result = helpers.ResponseJSON(constans.TRUE_VALUE, constans.SUCCESS_CODE, constans.EMPTY_VALUE, resDiv)
	return ctx.JSON(http.StatusOK, result)
}
