package offices

import (
	"backend/businesses/offices"
	"strconv"

	"gorm.io/gorm"
)

type officeRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) offices.Repository {
	return &officeRepository{
		conn: conn,
	}
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

func (or *officeRepository) Create(officeDomain *offices.Domain) offices.Domain {
	rec := FromDomain(officeDomain)

	result := or.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (or *officeRepository) Update(id string, officeDomain *offices.Domain) offices.Domain {
	var office offices.Domain = or.GetByID(id)

	updatedOffice := FromDomain(&office)

	updatedOffice.Title = officeDomain.Title
	updatedOffice.Description = officeDomain.Description
	updatedOffice.City = officeDomain.City

	or.conn.Save(&updatedOffice)

	return updatedOffice.ToDomain()
}

func (or *officeRepository) Delete(id string) bool {
	var office offices.Domain = or.GetByID(id)

	deletedOffice := FromDomain(&office)

	result := or.conn.Delete(&deletedOffice)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}

func (or *officeRepository) SearchByCity(city string) []offices.Domain {
	var rec []Office

	or.conn.Find(&rec, "city = ?", city)

	officeDomain := []offices.Domain{}

	for _, office := range rec {
		officeDomain = append(officeDomain, office.ToDomain())
	}

	return officeDomain
}

func (or *officeRepository) SearchByRate(rate string) []offices.Domain {
	var rec []Office

	intRate, _ := strconv.Atoi(rate)

	if intRate == 5 {
		or.conn.Find(&rec, "rate = ?", rate)
	} else {
		or.conn.Where("rate >= ? AND rate < ?", rate, intRate + 1).Order("rate desc, title").Find(&rec)
	}

	officeDomain := []offices.Domain{}

	for _, office := range rec {
		officeDomain = append(officeDomain, office.ToDomain())
	}

	return officeDomain
}
