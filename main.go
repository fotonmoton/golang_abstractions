package main

import (
	"abstractions/controllers"
	"abstractions/db"
	"abstractions/models"
	"abstractions/search"
	"abstractions/serializers"
	"net/http"

	"github.com/gorilla/mux"
)

/*
1. Why abstractions and patterns (decomposition and building up)
2. What is abstraction (abstraction is the process of generalizing concrete details, focus attention on details of greater importance)
2. What types of abstractions (functions, functions as arguments, goroutines, language features, interfaces, new languages, packages, modules, data abstractions)
3. Design patterns as a form of abstraction. (MVC, Adapter)
4. Layered architecture as a form of abstraction
5. Coupling and Cohesion in a program (reducing maintenance and modification costs)

5.1 A change in one module usually forces a ripple effect of changes in other modules.
5.2 Assembly of modules might require more effort and/or time due to the increased inter-module dependency.
5.3 A particular module might be harder to reuse and/or test because dependent modules must be included.
*/

func main() {

	var greg = models.User{"Greg", 21}

	var conn = db.Connection{[]string{"prod1", "prod2"}}

	var search = search.NewSearch(&conn, &serializers.StringSerializer{})

	r := mux.NewRouter()
	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) { controllers.ShowUserController(w, r, &greg) }).Methods("GET")
	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) { controllers.ChangeUserAgeController(w, r, &greg, 28) }).Methods("POST")
	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		controllers.SearchController(w, r, r.URL.Query().Get("searchTerm"), search)
	}).Methods("POST")

	http.ListenAndServe("localhost:8080", r)
}
