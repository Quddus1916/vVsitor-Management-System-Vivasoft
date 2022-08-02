package model

type Setting struct {
	Id        int  `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CompanyId int  `json:"company_id"`
	Email     bool `json:"email" validate:"required,eq=true|eq=false"`
	Image     bool `json:"image" validate:"required,eq=true|eq=false"`
}
