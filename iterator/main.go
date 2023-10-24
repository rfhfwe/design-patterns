package main

import "fmt"

var (
	users []*User
)

func Init() {
	users = append(users, &User{name: "a", age: 12})
	users = append(users, &User{name: "b", age: 20})
	users = append(users, &User{name: "c", age: 30})
}

func main() {

	Init()

	userCollection := &UserCollection{
		users: users,
	}
	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
