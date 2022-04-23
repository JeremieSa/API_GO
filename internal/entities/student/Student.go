package entities

type Student struct {
	Id           int
	FirstName    string
	LastName     string
	Age          int
	LanguageCode string
}

type studentList []Student

func NewStudent(id int, firstName string, lastName string, age int, languageCode string) Student {
	return Student{
		Id:           id,
		FirstName:    firstName,
		LastName:     lastName,
		Age:          age,
		LanguageCode: languageCode,
	}
}

func (l studentList) Len() int {
	return len(l)
}

func (l studentList) Less(i, j int) bool {
	return l[i].Id > l[j].Id
}

func (l studentList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func StudentString(student Student) string {
	return "{" +
		"	Etudiant number:" + string(rune(student.Id)) +
		"	First name: " + student.FirstName +
		"	Last name: " + student.LastName +
		"	Age: " + string(rune(student.Age)) +
		"	Language code: " + student.LanguageCode +
		"}"
}
