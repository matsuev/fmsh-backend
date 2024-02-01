package models

import "github.com/google/uuid"

// Account ...
type Account struct {
	ID                uuid.UUID
	UserName          string
	EncryptedPassword []byte
	name              string
}
