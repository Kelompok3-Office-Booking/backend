package users

import (
	// "backend/businesses/offices"
	// "fmt"

	// "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type officeRepository struct {
	conn *gorm.DB
}
