package model
import (
	"github.com/google/uuid"
)

type User struct{
	ID int `json:"id"`
	ExternalId uuid.UUID `json:"external_id"`
	Name string `json:"name"`
	Age int `json:"age"`
}