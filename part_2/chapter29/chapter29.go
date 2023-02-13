package chapter29

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Holymoly")
	})
	mux.HandleFunc("/bar", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "BarBarrrr")
	})
	mux.HandleFunc("/student", func(writer http.ResponseWriter, request *http.Request) {
		var stud = Student{"Park", 31, 47}
		data, _ := json.Marshal(stud)
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusOK)

		fmt.Fprint(writer, string(data))
	})
	return mux
}

func Ex01() {
	http.ListenAndServe(":8080", MakeWebHandler())
}

type Student struct {
	Name  string
	Age   int
	Score int
}
