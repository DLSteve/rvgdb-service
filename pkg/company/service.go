package company

import (
	"database/sql"
	"log"
	"rvgdb-service/pkg/entities"
)

type Service interface {
	GetCompanyList() ([]entities.Company, error)
}

type service struct {
	db *sql.DB
}

func NewCompanyService(db *sql.DB) Service {
	return &service{
		db,
	}
}

func (s *service) GetCompanyList() ([]entities.Company, error) {
	var compList []entities.Company

	rows, err := s.db.Query("SELECT * FROM company")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comp entities.Company
		err = rows.Scan(&comp.Id, &comp.Name, &comp.ShortName, &comp.HqCity, &comp.HqCountry, &comp.ParentCompanyId)
		if err != nil {
			log.Println(err.Error())
		}

		compList = append(compList, comp)
	}

	return compList, nil
}
