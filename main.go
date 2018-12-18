package main

import (
	"MyProjects/GolangProjects/GoCollections/Collections"
	"fmt"
)

func main() {
	myList := Collections.NewArrayList()
	myList.AddItem(10)
	myList.AddItem(20)
	myList.AddItem(30)
	myList.AddItem(40)
	myList.AddItem(50)
	myList.AddItem(60)
	myList.RemoveItem(10)
	myList.RemoveItemAtIndex(0)
	// myList.AddItem(types.NewPersonWithValues("Lovelesh", "M", 20))
	// myList.AddItem(types.NewPersonWithValues("Ankita", "F", 18))

	fmt.Println("ArrayList: ", myList.ToString())
	fmt.Println("ArrayList Size: ", myList.Size())
	fmt.Println("ArrayList Capacity: ", myList.Capacity())
	// fmt.Println(myList.Contains(types.Person{"Lovelesh", "M", 20}))
}
