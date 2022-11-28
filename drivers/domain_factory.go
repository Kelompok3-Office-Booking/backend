package drivers

import (
	userDomain "backend/businesses/users"
	// officeDomain "backend/businesses/offices"

	userDB "backend/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

// func OfficeRepository(conn *gorm.DB) officeDomain.OfficeRepository {

// }
