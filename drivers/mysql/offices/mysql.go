package offices

import (
	"backend/businesses/offices"

	"gorm.io/gorm"
)

type officeRepository struct {
	conn *gorm.DB
}

func OfficeMySQLRepository(conn *gorm.DB) offices.Repository {
	return &officeRepository{
		conn: conn,
	}
}

func (or *officeRepository) Create(officeDomain *offices.Domain) offices.Domain {

	rec := FromDomain(officeDomain)

	rec.Title = ""
	rec.Description = ""
	rec.City = ""
	rec.Rate = ""

	result := or.conn.Create(&rec)
	result.Last(&rec)

	return rec.ToDomain()
}

func (or *officeRepository) GetAll() []offices.Domain {
	var rec []Office

	or.conn.Find(&rec)

	officeDomain := []offices.Domain{}

	for _, office := range rec {
		officeDomain = append(officeDomain, office.ToDomain())
	}

	return officeDomain
}

func (or *officeRepository) GetByID(id string) offices.Domain {
	var office Office

	or.conn.First(&office, "id = ?", id)

	return office.ToDomain()
}

func (or *officeRepository) Delete(id string) bool {
	var office offices.Domain = or.GetByID(id)

	deletedOffice := FromDomain(&office)

	result := or.conn.Delete(&deletedOffice)

	return result.RowsAffected != 0
}

func (or *officeRepository) SearchByCity(city string) []offices.Domain {
	var search []offices.Domain

	or.conn.Find(&search, "city = ?", city)

	return search
}

func (or *officeRepository) SearchByRate(rate string) []offices.Domain {
	var search []offices.Domain

	or.conn.Find(&search, "rate = ?", rate)

	return search
}
