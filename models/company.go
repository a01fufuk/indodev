package models

type Division struct {
	ID                 int    `json:"id"`
	DivisionCode       string `json:"divisionCode"`
	DivisionName       string `json:"name"`
	DivisionHierarchy  string `json:"divisionHierarchy"`
	DivisionParentCode string `json:"divisionParentCode"`
}

type ResponseDivision struct {
	DivisionCode       string              `json:"divisionCode"`
	DivisionParentCode string              `json:"divisionParentCode"`
	DivisionName       string              `json:"name"`
	Children           []*ResponseDivision `json:"children"`
}

type RequestCreateDivision struct {
	DivisionCode       string `json:"divisionCode"`
	DivisionName       string `json:"name"`
	DivisionParentCode string `json:"divisionParentCode"`
}

type RequestEditDivision struct {
	DivisionCode string `json:"divisionCode"`
	DivisionName string `json:"name"`
}

type RequestDeleteDivision struct {
	DivisionCode string `json:"divisionCode"`
}
