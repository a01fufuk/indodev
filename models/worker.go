package models

type RequestWorkerPelni struct {
	Url string `json:"url"`
}

type RequestWorkerGetOrgDes struct {
	Token   string `json:"token" validate:"required"`
	OrgCode string `json:"orgCode" validate:"required"`
}

type ResultWorkerIlcsManifest struct {
	Status              int64  `json:"status" bson:"status"`
	Message             string `json:"message" bson:"message"`
	NewPenumpang        string `json:"newPenumpang" bson:"newPenumpang"`
	OldPenumpang        string `json:"oldPenumpang" bson:"oldPenumpang"`
	StatusSendToGate    string `json:"statusSendToGate" bson:"statusSendToGate"`
	ResponseSendToGate  string `json:"responseSendToGate" bson:"responseSendToGate"`
	StatusSendToPtosR   string `json:"statusSendToPtosR" bson:"statusSendToPtosR"`
	ResponseSendToPtosR string `json:"responseSendToPtosR" bson:"responseSendToPtosR"`
}
