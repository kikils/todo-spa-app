package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kikils/golang-todo/interfaces/controllers"
)

func SetUpRouting() *http.ServeMux {
	mux := http.NewServeMux()

	sqlhandler := NewSqlhandler()
	err := CreateTable(sqlhandler)
	if err != nil {
		log.Println(err.Error())
	}
	todoController := controllers.NewTodoController(sqlhandler)

	mux.HandleFunc("/todo/create", AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Create(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	}))
	mux.HandleFunc("/todo/update", AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Update(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	}))
	mux.HandleFunc("/todo/delete", AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Delete(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	}))
	mux.HandleFunc("/todo/get", AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoController.FindByUserID(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	}))
	return mux
}

func LogHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rAddr := r.RemoteAddr
		method := r.Method
		path := r.URL.Path
		fmt.Printf("Remote: %s [%s] %s\n", rAddr, method, path)
		h.ServeHTTP(w, r)
	})
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	body := map[string]string{
		"error": message,
	}
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
