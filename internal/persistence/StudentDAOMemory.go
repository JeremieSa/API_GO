package persistence

import (
	. "internal/entities/student"
)

var studentMap = make(map[int]Student)

type StudentDAOMemory struct{}

var _ StudentDAOInterface = (*StudentDAOMemory)(nil)

// (s StudentDAOMemory)
func (s StudentDAOMemory) FindAll() []Student {
	var result []Student

	for _, element := range studentMap {
		result = append(result, element)
	}

	return result
}

func (s StudentDAOMemory) Find(id int) *Student {
	for key, element := range studentMap {
		if key == id {
			return &element
		}
	}
	return nil
}

func (s StudentDAOMemory) Exist(id int) bool {
	_, ok := studentMap[id]
	return ok
}

func (s StudentDAOMemory) Delete(id int) bool {
	if !s.Exist(id) {
		return false
	}

	delete(studentMap, id)
	return true
}

func (s StudentDAOMemory) Create(student Student) bool {
	if !s.Exist(student.Id) {
		studentMap[student.Id] = student
		return true
	}

	return false
}

func (s StudentDAOMemory) Update(student Student) bool {
	if !s.Exist(student.Id) {
		return false
	}

	studentMap[student.Id] = student
	return true
}
