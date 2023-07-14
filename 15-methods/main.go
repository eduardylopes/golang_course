package main

import "fmt"

type User struct {
	name string
	age  uint8
}

func (user User) save() {
	fmt.Printf("Saving the user %s in database\n", user.name)
}

func (user User) checkLegalAge() {
	if user.age >= 18 {
		fmt.Printf("%s is of age\n", user.name)
		return
	}

	fmt.Printf("%s is not of age\n", user.name)
}

// This method changes the user's age value
func (user *User) haveBirthday() {
	fmt.Println(user)     // This is a memory reference
	fmt.Println(user.age) // This is a value

	user.age++
}

func main() {
	user := User{"Eduardy", 28}
	fmt.Println(user)
	user.checkLegalAge()
	user.save()

	user2 := User{"Ana", 27}
	fmt.Println(user2)
	user2.checkLegalAge()
	user2.save()

	user3 := User{"Jonas", 16}
	fmt.Println(user3)
	user3.checkLegalAge()
	user3.save()

	user3.haveBirthday()
	fmt.Println(user3)
}
