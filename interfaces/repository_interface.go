package interfaces

import (
	"accelone-challenge/model"
)

type ContactRepository interface {
	AddContact(contact model.Contact) (model.Contact, error)
	GetContact(contactID string) (model.Contact, error)
	DeleteContact(contactID string) error
	PutContact(contact model.Contact) (model.Contact, error)
}
