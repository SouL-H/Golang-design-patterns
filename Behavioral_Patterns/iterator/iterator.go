package main

import (
	"fmt"
)

// ########## COLLECTION ##########

type collection interface {
	createIterator() iterator
}

// ########## USER COLLECTION ##########

type userCollection struct {
	users []*user
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}

// ########## ITERATOR ##########

type iterator interface {
	hasNext() bool
	getNext() *user
}

// ########## USER ITERATOR ##########

type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false

}

func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

// ########## USER ##########

type user struct {
	name string
	age  int
}

// ########## MAIN ##########

func main() {

	user1 := &user{
		name: "a",
		age:  30,
	}
	user2 := &user{
		name: "b",
		age:  20,
	}
	user3 := &user{
		name: "c",
		age:  45,
	}

	userCollection := &userCollection{
		users: []*user{user1, user2, user3},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
