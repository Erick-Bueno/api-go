import (
	"github.com/google/uuid"
)
package model

type User struct{
	ID int `json:"id"`
	ExternalId uuid.UUID `json:"external_id"`
	Name string `json:"name"`
	Age int `json:"age"`
}