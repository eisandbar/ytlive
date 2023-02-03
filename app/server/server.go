package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/eisandbar/ytlive/app/store"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewServer(store store.Store) Server {
	server := Server{}
	server.store = store
	return server
}

func StartServer(server Server, port string) {
	// Start http server
	router := mux.NewRouter()
	router.HandleFunc("/streams", server.getStreams).Methods("GET")
	router.HandleFunc("/categories", server.getCategories).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://ytlive.online", "http://localhost:5173"},
		AllowCredentials: true,
	}).Handler(router)

	fmt.Println("Listening on port:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
}

type Server struct {
	store store.Store
}

func (s *Server) getStreams(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getStreams request")

	opts, err := getOptions(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := s.store.List(opts...)

	body, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		log.Printf("Failed to write response, %s\n", err)
	}
}

func (s *Server) getCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getCategories request")

	// If we only want the game categories
	gaming := r.URL.Query().Has("gaming") && r.URL.Query().Get("gaming") == "true"

	result := s.store.Categories(gaming)

	body, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		log.Printf("Failed to write response, %s\n", err)
	}
}

func getOptions(r *http.Request) ([]store.ListOption, error) {
	opts := make([]store.ListOption, 0)
	opts = append(opts, store.WithLive(true))

	query := r.URL.Query()

	if query.Has("maxResults") {
		value, err := strconv.Atoi(query.Get("maxResults"))
		if err != nil {
			return nil, err
		}
		opts = append(opts, store.WithMaxResults(value))
	}

	if query.Has("offset") {
		value, err := strconv.Atoi(query.Get("offset"))
		if err != nil {
			return nil, err
		}
		opts = append(opts, store.WithOffset(value))
	}

	if query.Has("gaming") {
		value, err := strconv.ParseBool(query.Get("gaming"))
		if err != nil {
			return nil, err
		}
		opts = append(opts, store.WithGaming(value))
	}

	if query.Has("filters") {
		filters := make([]string, 0, len(query["filters"]))

		for _, filter := range query["filters"] {
			decodedValue, err := url.QueryUnescape(filter)
			if err != nil {
				return nil, err
			}
			filters = append(filters, decodedValue)
		}
		opts = append(opts, store.WithFilters(filters))
	}

	return opts, nil

}
