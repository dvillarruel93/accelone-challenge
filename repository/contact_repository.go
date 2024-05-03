package repository

import (
	uuid "accelone-challenge/helper"
	"accelone-challenge/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type ContactRepository struct {
	memoryDb map[string]interface{}
}

func NewContactRepository(db map[string]interface{}) *ContactRepository {
	return &ContactRepository{
		memoryDb: db,
	}
}

func (r *ContactRepository) AddContact(contact model.Contact) (model.Contact, error) {
	// generates the universal unique ID
	contactID := uuid.GenerateUUID()

	// check if uuid is already used in memoryDb (this is a non-possible case but is better to check it)
	if _, ok := r.memoryDb[contactID]; ok {
		return model.Contact{}, echo.NewHTTPError(http.StatusInternalServerError, "contact already added")
	}

	// if is not in memory, it sets the uuid and saves in memoryDB
	contact.ID = contactID
	r.memoryDb[contact.ID] = contact

	return contact, nil
}

func (r *ContactRepository) GetContact(contactID string) (model.Contact, error) {
	// if it's not in memoryDB it returns an empty struct without an error, but it has a log to let devs know that
	// there is no register in memoryDB with the received contactID
	if _, ok := r.memoryDb[contactID]; !ok {
		log.Infof("there is no contact in db with id %s", contactID)
		return model.Contact{}, nil
	}

	// It has to convert interface{} from map to model.Contact
	contact := r.memoryDb[contactID].(model.Contact)

	return contact, nil
}

func (r *ContactRepository) DeleteContact(contactID string) error {
	delete(r.memoryDb, contactID)

	return nil
}

func (r *ContactRepository) PutContact(contact model.Contact) (model.Contact, error) {
	// get contact from memoryDB, if is not present it returns an error
	if _, ok := r.memoryDb[contact.ID]; !ok {
		return model.Contact{}, echo.NewHTTPError(http.StatusCreated,
			fmt.Sprintf("there is no contact in db with id %s", contact.ID))
	}

	// save changed contact in memoryDB
	r.memoryDb[contact.ID] = contact

	return contact, nil
}
