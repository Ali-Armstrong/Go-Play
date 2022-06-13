package main

import "fmt"

type contactInfo struct{
	email string
	zipCode string
}

type person struct{
	firstName string
	lastName string
	contact contactInfo
}

func (p person) print(){
	fmt.Println()
	fmt.Printf("%+v",p)
}

func (p person) updateName(newName string){
	p.firstName = newName
}

//*person is telling that we are working with address of type person
//*pointer is a person variable and give actual value
func (pointer *person) updateFirstName(newName string){
	(*pointer).firstName = newName
}

func main(){
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var bob person
	fmt.Println(bob)

	bob.firstName = "Bob"
	fmt.Println(bob)

	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email: "jim@gmail.com",
			zipCode: "900",
		},
	}

	jim.updateName("Jimmy")
	jim.print()

	jimPoint := &jim
	jimPoint.updateFirstName("Jimmy")
	jim.print()
	//or we can either call directly
	//Go do it automatically for us
	jim.updateFirstName("Tommy")
	jim.print()

}