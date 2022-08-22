package maps

import "errors"

type Person struct {
	Id        int
	FirstName string
	LastName  string
}

var cache = make(map[int]*Person)

var ErrIdAlreadyKnown = errors.New("id is already known")
var ErrUnknownID = errors.New("id is unknown")

func AddPerson(person *Person) (*Person, error) {
	if cache[person.Id] != nil {
		return &Person{}, ErrIdAlreadyKnown
	}

	cache[person.Id] = person
	return cache[person.Id], nil
}

func GetPerson(id int) (*Person, error) {
	ret := cache[id]
	if ret == nil {
		return &Person{}, ErrUnknownID
	}

	return ret, nil
}

func ClearCache() {
	cache = make(map[int]*Person)
}

func RemovePerson(id int) {
	delete(cache, id)
}
