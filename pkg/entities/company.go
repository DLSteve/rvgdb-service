package entities

import "github.com/google/uuid"

type Company struct {
	Id              uuid.UUID  `json:"id"`
	Name            string     `json:"name"`
	ShortName       string     `json:"short_name"`
	HqCity          string     `json:"hq_city"`
	HqCountry       string     `json:"hq_country"`
	ParentCompanyId *uuid.UUID `json:"parent_company_id,omitempty"`
}
