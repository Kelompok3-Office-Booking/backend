package users

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	FullName  string
	Gender    string
	Email     string
	Password  string
	Photo     string
	Roles     string
}

type LoginDomain struct {
	Email    string
	Password string
}

type PhotoDomain struct {
	Photo string
}

type Usecase interface {
	Register(userDomain *Domain) Domain
	Login(userDomain *LoginDomain) map[string]string
	Token(userId string, roles string) map[string]string
	GetAll() []Domain
	GetByID(id string) Domain
	GetProfile(id string) Domain
	Delete(id string) bool
	UpdateProfileData(id string, userDomain *Domain) Domain
	UpdateProfilePhoto(id string, userDomain *PhotoDomain) bool
	SearchByEmail(email string) Domain
}

type Repository interface {
	Register(userDomain *Domain) Domain
	GetByEmail(userDomain *LoginDomain) Domain
	CheckUserByEmailOnly(id string, email string) bool
	GetAll() []Domain
	GetByID(id string) Domain
	Delete(id string) bool
	UpdateProfileData(id string, userDomain *Domain) Domain
	InsertURLtoUser(id string, userDomain *PhotoDomain) bool
	SearchByEmail(email string) Domain
}