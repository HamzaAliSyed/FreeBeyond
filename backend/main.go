package main

import (
	"fmt"
	"log"
	"net/http"
)

func genericOk(reponse http.ResponseWriter, request *http.Request) {
	fmt.Println("Received a request:", request.Method, request.URL.Path)

	// Serve files from the current directory if needed
	fileServer := http.FileServer(http.Dir("."))
	fileServer.ServeHTTP(reponse, request)

	// Respond with status OK
	reponse.WriteHeader(http.StatusOK)
}

func main() {
	ConnectToMongo()
	const port = "2712"
	backend := http.NewServeMux()
	backend.HandleFunc("/", genericOk)
	backend.HandleFunc("/api/accounts/createrequest", HandleCreateAccountRequest)
	backend.HandleFunc("/api/accounts/createaccount", HandleCreateAccountUserFormRequest)
	backend.HandleFunc("/api/accounts/usernamematchrequest", HandleUserNameMatchRequest)
	log.Fatal(http.ListenAndServe(":"+port, backend))
}
