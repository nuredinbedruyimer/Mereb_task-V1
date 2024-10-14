package services

import (
	"Mereb-V1/constants"
	"Mereb-V1/database"
	"Mereb-V1/models"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

//  Crete New Person After Check Weather it Register Already or not

func CreatePersonService(person models.Person) error {
	//  Because We have  User with Out Id We have to create out own (UUID)

	person.ID = uuid.New()

	// Check The existance of other user with same id (same user twice) !!

	if _, isExist := database.GetPerson(person.ID.String()); isExist {
		return errors.New(constants.ERROR_USER_ALREADY_EXIST)
	}
	//  Use The The Creation At The DataBase Label Bay Loking and Unlocking With Muxes for Good resource Sharing

	database.CreatePerson(person)

	return nil
}

//  Fetch All Person

func FetchAllPersonsService() []models.Person {
	return database.GetPersons()
}

//  Fetch Single Person Using ID

func FetchPersonService(personID string) (*models.Person, error) {
	person, isExist := database.GetPerson(personID)

	if !isExist {
		return nil, errors.New(constants.ERROR_USER_NOT_EXIST)
	}
	return &person, nil

}

func DeletePersonService(personID string) (string, error) {
	_, isExist := database.GetPerson(personID)
	fmt.Println("isExist : ", isExist)

	if !isExist {
		return constants.USER_NOT_EXIST, errors.New(constants.ERROR_USER_NOT_EXIST)
	}
	database.DeletePerson(personID)
	return constants.USER_DELETED, nil
}

func UpdatePersonService(id string, updatedPerson models.Person) error {
	existingPerson, isExist := database.GetPerson(id)
	if !isExist {
		return errors.New(constants.ERROR_USER_ALREADY_EXIST)
	}

	existingPerson.Name = updatedPerson.Name
	existingPerson.Age = updatedPerson.Age
	existingPerson.Hobbies = updatedPerson.Hobbies

	database.UpdatePerson(existingPerson)
	return nil
}
