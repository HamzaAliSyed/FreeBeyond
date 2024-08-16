package main

import (
	"fmt"
	"log"
	"net/http"
)

func genericOk(reponse http.ResponseWriter, request *http.Request) {
	fmt.Println("Received a request:", request.Method, request.URL.Path)

	fileServer := http.FileServer(http.Dir("."))
	fileServer.ServeHTTP(reponse, request)

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
	backend.HandleFunc("/api/accounts/signinrequest", HandleSignAccountRequest)
	backend.HandleFunc("/api/accounts/signin", HandleSignInFormRequest)
	backend.HandleFunc("/api/accounts/createacharacter", HandleCharacterCreation)
	backend.HandleFunc("/api/accounts/character/addcharactername", AddCharacterName)
	backend.HandleFunc("/api/accounts/character/addattributes", AddAttributes)
	backend.HandleFunc("/api/charactergeneration/skills/", HandleSkillsFactory)
	backend.HandleFunc("/api/charactergeneration/addcharactermotives", HandleAddCharacterMotives)
	log.Fatal(http.ListenAndServe(":"+port, backend))
}
