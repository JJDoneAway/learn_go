package maps

import (
	"reflect"
	"testing"
)

func TestAddPerson(t *testing.T) {
	ClearCache()
	t.Run("Simple test", func(t *testing.T) {
		person := &Person{1, "1_V", "1_N"}

		got, _ := AddPerson(person)

		// we are dealing with pointers to the same address
		if person != got && !reflect.DeepEqual(*got, *person) {
			t.Errorf("Got %p got expect %p", got, person)
		}

		// so if we change an value via one pointer it must be also valid for the second one
		person.FirstName = "Pimmel"

		if person != got && !reflect.DeepEqual(*got, *person) {
			t.Errorf("Got %p got expect %p", got, person)
		}

		if len(cache) != 1 {
			t.Errorf("Cache size must be one but was %d", len(cache))
		}

	})

	t.Run("Try to add with already used ID", func(t *testing.T) {
		person1 := &Person{1, "1_V", "1_N"}
		person2 := &Person{1, "2_V", "2_N"}

		AddPerson(person1)
		got2, err := AddPerson(person2)

		// we are dealing with pointers to the same address
		if person2 == got2 || reflect.DeepEqual(*person2, *got2) {
			t.Errorf("must not be possible to overwrite")
		}

		if err == nil {
			t.Errorf("We need an error here but got nill")
		}

	})

}

func TestGetPerson(t *testing.T) {
	ClearCache()
	person := Person{1, "1_v", "1_n"}

	t.Run("Get nil test", func(t *testing.T) {
		_, err := GetPerson(person.Id)

		if err != ErrUnknownID {
			t.Errorf("Need error %v here but got nil", ErrUnknownID)
		}

	})

	t.Run("Get true person test", func(t *testing.T) {
		AddPerson(&person)
		got, err := GetPerson(person.Id)

		if err != nil {
			t.Errorf("Error must not be nil")
		}

		if got != &person {
			t.Errorf("Got %v but expect %v", got, person)
		}

	})

}

func TestRemoveFromCache(t *testing.T) {
	amount := 100
	for i := 1; i <= amount; i++ {
		AddPerson(&Person{i, "pimmel", "pummel"})
	}

	for i := 1; i <= amount; i++ {
		got, err := GetPerson(i)
		if got.Id != i || err != nil {
			t.Errorf("Expect Person with ID: %d but got rubbish or an error %v", i, err)
		}
	}

	RemovePerson(amount + 1)

	for i := 1; i <= amount; i++ {
		got, err := GetPerson(i)
		if got.Id != i || err != nil {
			t.Errorf("Expect Person with ID: %d but got rubbish or an error %v", i, err)
		}
	}

	RemovePerson(5)
	got, err := GetPerson(5)
	if err != ErrUnknownID {
		t.Errorf("%v must be cleaned but got one", got)
	}

}
