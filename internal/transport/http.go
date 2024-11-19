package transport

import (
	"encoding/json"
	"log"
	"my-first-api/internal/todo"
	"net/http"
)

type TodoItem struct {
	Item string `json:"item"`
}

type Server struct {
	mux *http.ServeMux
}

func NewServer(todoSvc *todo.Service) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		todoItems, err := todoSvc.GetAll() // Fixed: Added := for declaration
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json") // Added content type header
		b, err := json.Marshal(todoItems)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	})

	mux.HandleFunc("POST /todo", func(writer http.ResponseWriter, request *http.Request) {
		var t TodoItem
		err := json.NewDecoder(request.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return // Added missing return
		}
		err = todoSvc.Add(t.Item)
		if err != nil {
			log.Println(err) // Added error logging
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("GET /search", func(writer http.ResponseWriter, request *http.Request) {
		query := request.URL.Query().Get("q")
		if query == "" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		results, err := todoSvc.Search(query) // Fixed: Added error return
		if err != nil {
			log.Println(err) // Fixed: Changed Fatal to Println
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json") // Added content type header
		b, err := json.Marshal(results)
		if err != nil {
			log.Println(err) // Added error logging
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = writer.Write(b)
		if err != nil {
			log.Println(err)
		}
	})

	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
