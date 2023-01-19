package service

import (
	"unitable/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type profileStorage interface {
	SaveUser(*domain.User) error
}

type profileService struct {
	repo profileStorage
}

func NewProfileService(storage profileStorage) *profileService {
	return &profileService{repo: storage}
}

func (s *profileService) EditProfile(user *domain.User, profile *domain.UserProfile) error {
	user.Profile = *profile

	return s.repo.SaveUser(user)
}

func (s *profileService) AppendContact(user *domain.User, contactName string, contactValue string) error {

	user.Profile.Contacts = append(user.Profile.Contacts, *domain.NewProfileContact(contactName, contactValue))

	return s.repo.SaveUser(user)
}

func (s *profileService) EditContact(user *domain.User, contactID primitive.ObjectID, newName string, newValue string) error {

	var contacts []domain.ProfileContact

	for _, contact := range user.Profile.Contacts {
		if contact.ID == contactID {
			contact.Name = newName
			contact.Value = newValue
		}
		contacts = append(contacts, contact)
	}

	user.Profile.Contacts = contacts

	return s.repo.SaveUser(user)
}

func removeContact(slice []domain.ProfileContact, index int) []domain.ProfileContact {
	slice[index] = slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return slice
}

func (s *profileService) DeleteContacts(user *domain.User, contactIDs []string) error {

	for _, ID := range contactIDs {
		for index, userContact := range user.Profile.Contacts {
			if userContact.ID.Hex() == ID {
				user.Profile.Contacts = removeContact(user.Profile.Contacts, index)
			}
		}
	}

	return s.repo.SaveUser(user)
}
