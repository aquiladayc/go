package main

import "fmt"

type contactInfo struct {
    email string
    zipCode int
}


type person struct {
    firstName string
    lastName string
    contact contactInfo
}

func main(){
    jim := person{
        firstName: "Jim",
	lastName: "Hall",
	contact: contactInfo{
	    email: "jim@gmail.com",
	    zipCode: 1230000,
	},
    }
    jim.updateName("Jimmy")
    jim.print()
}

func (p person) print() {
    fmt.Printf("%+v", p)
}

func (pToPerson *person) updateName (newFirstName string) {
    pToPerson.firstName = newFirstName
}
