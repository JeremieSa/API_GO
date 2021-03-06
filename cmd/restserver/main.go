package main

import (
	"fmt"
	. "internal/persistence"
	. "internal/web/rest"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// @title           API Go
// @version         1.0
// @description     This API brings CRUD operations about students and languages
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	fmt.Println("Meilleur API au monde, réalisé avec Gorilla Mux Routers !")
	InitialiseDB()
	defer CloseDatabase()

	// Instance of mux router
	r := mux.NewRouter()

	// Define routes and handlers
	// Languages
	r.HandleFunc("/rest/languages/{code}", GetOneLanguage).Methods("GET")
	r.HandleFunc("/rest/languages", GetAllLanguage).Methods("GET")
	r.HandleFunc("/rest/languages", CreateLanguageHandler).Methods("POST")
	r.HandleFunc("/rest/languages", UpdateLanguageHandler).Methods("PUT")
	r.HandleFunc("/rest/languages/{code}", DeleteLanguageByIdHandler).Methods("DELETE")

	// Students
	r.HandleFunc("/rest/students/{id}", GetOneStudent).Methods("GET")
	r.HandleFunc("/rest/students", GetAllStudent).Methods("GET")
	r.HandleFunc("/rest/students", CreateStudentHandler).Methods("POST")
	r.HandleFunc("/rest/students", UpdateStudentHandler).Methods("PUT")
	r.HandleFunc("/rest/students/{id}", DeleteStudentByIdHandler).Methods("DELETE")

	http.Handle("/", r)

	// Error handling
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
