package database

import (
	"Mereb-V1/models"
	"sync"
)

//

var persons = make(map[string]models.Person)

//    For Thread Safty

var mutex = &sync.Mutex{}

func CreatePerson(person models.Person) {
	mutex.Lock()

	//  place current user using id as key

	persons[person.ID.String()] = person

	mutex.Unlock()

}

func GetPersons() []models.Person {
	mutex.Lock()
	allPersons := make([]models.Person, 0, len(persons))
	for _, currentPerson := range persons {
		allPersons = append(allPersons, currentPerson)
	}

	mutex.Unlock()
	return allPersons
}
func GetPerson(id string) (models.Person, bool) {
	mutex.Lock()
	foundedUser, isExist := persons[id]
	mutex.Unlock()
	return foundedUser, isExist
}

func DeletePerson(id string) {
	mutex.Lock()

	delete(persons, id)
	mutex.Unlock()
}

func UpdatePerson(person models.Person) {
	mutex.Lock()
	persons[person.ID.String()] = person
	mutex.Unlock()
}
