package usecase

import (
	"accelone-challenge/interfaces"
	"accelone-challenge/model"
)

type ContactUsecase struct {
	ContactRepository interfaces.ContactRepository
}

func NewContactUsecase(contactRepository interfaces.ContactRepository) *ContactUsecase {
	return &ContactUsecase{
		ContactRepository: contactRepository,
	}
}

func (u *ContactUsecase) AddContact(contact model.Contact) (model.Contact, error) {
	return u.ContactRepository.AddContact(contact)
}

func (u *ContactUsecase) GetContact(contactID string) (model.Contact, error) {
	return u.ContactRepository.GetContact(contactID)
}

func (u *ContactUsecase) DeleteContact(contactID string) error {
	return u.ContactRepository.DeleteContact(contactID)
}

func (u *ContactUsecase) PutContact(contact model.Contact) (model.Contact, error) {
	return u.ContactRepository.PutContact(contact)
}
