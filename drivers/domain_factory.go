package drivers

import (
	officeDomain "backend/businesses/offices"
	userDomain "backend/businesses/users"
	officeImageDomain "backend/businesses/office_images"

	officeDB "backend/drivers/mysql/offices"
	userDB "backend/drivers/mysql/users"
	officeImageDB "backend/drivers/mysql/office_images"

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