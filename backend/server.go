package backend

import (
	"aurafiles/backend/data"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//TODO: RateLimit

// StartServer initialises the Mux router, setups the routes and starts the http server
func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", data.DefaultRoute).Methods("GET")
	r.HandleFunc("/", data.DefaultRoute).Methods("POST")

	r.HandleFunc("/upload", data.UploadFileRoute).Methods("POST")
	r.HandleFunc("/fileInfo/{id}", data.GetFileInfoRoute).Methods("GET")
	r.HandleFunc("/{id}", data.GetFileRoute).Methods("GET")
	r.HandleFunc("/addDownload/{id}", data.AddDownloadRoute).Methods("POST")

	fmt.Println("Server is listening on port 8000!")
	http.ListenAndServe(":8000", r)
}
