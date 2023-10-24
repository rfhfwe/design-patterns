package main

// Collection 集合接口，目的是创建迭代器对象
type Collection interface {
	createIterator() Iterator
}

type User struct {
	name string
	age  int
}

// UserCollection 要迭代的目标对象
type UserCollection struct {
	users []*User
}

func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}
