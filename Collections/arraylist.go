package Collections

import (
	"errors"
	"reflect"
)

// List interface defines the contract between various types of Lists
type List interface {
	AddItem(interface{}) error
	RemoveItem(interface{})
	RemoveItemAtIndex(int) error
}

// ArrayList defines the structure of a simple ArrayList by using a slice
type ArrayList struct {
	slice       []interface{}
	elementType string
}

// NewArrayList returns an empty ArrayList
func NewArrayList() ArrayList {
	arrList := ArrayList{}
	arrList.Initialize()
	return arrList
}

// Initialize initialized the ArrayList with the empty slice
func (al *ArrayList) Initialize() []interface{} {
	al.slice = make([]interface{}, 0, 5)
	return al.slice
}

// AddItem adds an element into the ArrayList
func (al *ArrayList) AddItem(item interface{}) error {
	if al.elementType == "" {
		// storing the type of the first element added into the ArrayList
		al.elementType = reflect.TypeOf(item).String()
	}
	if al.elementType == reflect.TypeOf(item).String() {
		al.slice = append(al.slice, item)
	} else {
		errString := "Type Mismatch. Cannot insert item of type " + reflect.TypeOf(item).String() + " into slice of type " + reflect.TypeOf(al.slice).String()
		return errors.New(errString)
	}
	return nil
}

// RemoveItemAtIndex removes an element at specified index from the ArrayList
func (al *ArrayList) RemoveItemAtIndex(index int) error {
	if index >= al.Size() {
		return errors.New("Cannot remove item at index " + string(index) + ". Index is out of bounds!")
	}
	al.slice = append(al.slice[:index], al.slice[index+1:]...)
	return nil
}

// RemoveItem removes the specified element from the ArrayList
func (al *ArrayList) RemoveItem(value interface{}) {
	for i := 0; i < al.Size(); i++ {
		if al.slice[i] == value {
			al.RemoveItemAtIndex(i)
		}
	}
}

// Contains returns true if the slice contains the specified element
func (al *ArrayList) Contains(value interface{}) bool {
	var flag bool
	for i := 0; i < al.Size(); i++ {
		if al.slice[i] == value {
			flag = true
			break
		}
	}
	return flag
}

// Size returns the number of elements stored in the ArrayList
func (al *ArrayList) Size() int {
	return len(al.slice)
}

// Capacity returns the number of elements can be accomodated in the ArrayList
func (al *ArrayList) Capacity() int {
	return cap(al.slice)
}

// ToString returns the complete ArrayList in string form
func (al *ArrayList) ToString() interface{} {
	return al.slice
}
