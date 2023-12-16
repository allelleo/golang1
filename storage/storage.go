package storage

import (
	"errors"
	"fmt"
)

type User struct {
	id         int
	username   string
	first_name string
	last_name  string
	password   string
}

type storage interface {
	create(user User) error
	get(id int) (User, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]User
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]User),
	}
}

func (box *memoryStorage) create(user User) error {
	box.data[user.id] = user
	return nil
}

func (box *memoryStorage) get(id int) (User, error) {
	user, exists := box.data[id]
	if !exists {
		return User{}, errors.New("User not found")
	}
	return user, nil
}

func (box *memoryStorage) delete(id int) error {
	delete(box.data, id)
	return nil
}

func main() {
	var s storage = newMemoryStorage()
	fmt.Println(s)

	var user1 User = User{
		id:         1,
		username:   "user1",
		first_name: "user1",
		last_name:  "user1",
		password:   "user1",
	}

	s.create(user1)

	fmt.Println(s)

	s.get(user1.id)
	s.delete(user1.id)

	fmt.Println(s)
}
