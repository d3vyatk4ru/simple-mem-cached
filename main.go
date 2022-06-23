package main

import (
	"errors"
	"fmt"
)

type employee struct {
	id     uint
	age    uint8
	name   string
	salary float32
}

// set interfase for memory storage.
// the main idea is behavior of our model,
// but the realization is not main target
type storage interface {
	set(e employee) error
	get(id uint) (employee, error)
	delete(id uint) error
}

type memoryStorage struct {
	data map[uint]employee
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[uint]employee),
	}
}

func (m *memoryStorage) set(e employee) error {

	m.data[e.id] = e

	return nil
}

func (m *memoryStorage) get(id uint) (employee, error) {

	emp, exists := m.data[id]

	if !exists {
		return employee{}, errors.New("Employee id not exist!")
	}

	return emp, nil
}

func (m *memoryStorage) delete(id uint) error {
	delete(m.data, id)
	return nil
}

func main() {

	var s storage

	s = newMemoryStorage()

	emp := employee{
		1,
		18,
		"Vasya",
		100000,
	}

	s.set(emp)

	fmt.Println(s.get(1))
}
