package main

// Iterator 迭代器接口
type Iterator interface {
	hasNext() bool
	getNext() *User
}

// UserIterator 具体要迭代的迭代器对象，相当于上层要初始化这个对象，进行具体操作
type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) getNext() *User {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
