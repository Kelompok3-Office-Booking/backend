package drivers

import (
	facilityDomain "backend/businesses/facilities"
	officeImageDomain "backend/businesses/office_images"
	officeDomain "backend/businesses/offices"
	userDomain "backend/businesses/users"

	facilityDB "backend/drivers/mysql/facilities"
	officeImageDB "backend/drivers/mysql/office_images"
	officeDB "backend/drivers/mysql/offices"
	userDB "backend/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

func NewOfficeRepository(conn *gorm.DB) officeDomain.Repository {
	return officeDB.NewMySQLRepository(conn)
}

func NewOfficeImageRepository(conn *gorm.DB) officeImageDomain.Repository {
	return officeImageDB.NewMySQLRepository(conn)
}

func NewFacilityRepository(conn *gorm.DB) facilityDomain.Repository {
	return facilityDB.NewMySQLRepository(conn)
}
