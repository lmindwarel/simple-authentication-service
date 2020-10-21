package models

import (
	"time"

	models "github.com/lmindwarel/quizzbox-models"
)

// CollLogin is the collection name for logins
const CollLogin = "login"

// Login is strut holding account identifiers in order to log user in
type Login struct {
	ID       models.UUID `json:"id" bson:"_id"`
	Email    string      `json:"email" bson:"email"`
	Password string      `json:"password" bson:"password"`
}

// CollAccess is the collection name for access
const CollAccess = "access"

// Access store account authenticated access data
type Access struct {
	ID        models.UUID `json:"id" bson:"_id"`
	Token     string      `json:"token" bson:"token"`
	StartDate time.Time   `json:"startDate" bson:"startDate"`
	EndDate   time.Time   `json:"endDate" bson:"endDate"`
}
