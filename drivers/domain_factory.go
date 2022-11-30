package drivers

import (
	officeDomain "backend/businesses/offices"
	userDomain "backend/businesses/users"

	officeDB "backend/drivers/mysql/offices"
	userDB "backend/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

func NewOfficeRepository(conn *gorm.DB) officeDomain.Repository {
	return officeDB.OfficeMySQLRepository(conn)
}
