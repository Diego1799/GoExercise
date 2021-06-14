package main

import (
	"encoding/json"
	"net/http"
)

type Curso struct {
	Title string
	Num   int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursos := Cursos{
			Curso{"Curso de go", 30},
			Curso{"Curso de go", 25},
		}
		json.NewEncoder(w).Encode(cursos)
	})

	http.ListenAndServe(":8000", nil)
}
