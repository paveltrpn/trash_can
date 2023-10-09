package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"skillfactory-go-developer/pkg/user"
	"strconv"
	"strings"

	chi "github.com/go-chi/chi/v5"
)

var userDbInstance user.UserDB

func handleUserCreate(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("get create\n")

	reqBody, _ := io.ReadAll(request.Body)

	var u user.User
	json.Unmarshal(reqBody, &u)

	id := userDbInstance.AddUser(u)

	writer.WriteHeader(201)
	_, err := writer.Write([]byte(fmt.Sprintf("new user id - %v", strconv.Itoa(id))))
	if err != nil {
		log.Println(err)
	}
}

func handleMakeFriends(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("get make-friends\n")

	reqBody, _ := io.ReadAll(request.Body)
	var reqMap map[string]int
	json.Unmarshal(reqBody, &reqMap)

	source, _ := reqMap["source_id"]
	target, _ := reqMap["target_id"]

	if err := userDbInstance.MakeFriends(source, target); err != nil {
		_, err := writer.Write([]byte(fmt.Sprintf("unable to make friends!")))
		if err != nil {
			log.Println(err)
		}
		return
	}

	_, err := writer.Write([]byte(fmt.Sprintf("username - %v and username - %v now are friends", userDbInstance.GetUserName(source), userDbInstance.GetUserName(target))))
	if err != nil {
		log.Println(err)
	}
}

func handleUserDelete(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("get delete\n")

	reqBody, _ := io.ReadAll(request.Body)
	var reqMap map[string]int
	json.Unmarshal(reqBody, &reqMap)

	target_id := reqMap["target_id"]

	var userName string

	if err := userDbInstance.CheckUser(target_id); err == nil {
		userName = userDbInstance.GetUserName(target_id)
		userDbInstance.DeleteUser(target_id)
	} else {
		_, err := writer.Write([]byte(fmt.Sprintf("error! user not exist - %v", target_id)))
		if err != nil {
			log.Println(err)
		}
		return
	}

	_, err := writer.Write([]byte(fmt.Sprintf("user - %v has been deleted!", userName)))
	if err != nil {
		log.Println(err)
	}

}

func hanldeShowFriends(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("get show-friends\n")

	user_id, _ := strconv.Atoi(chi.URLParam(request, "user_id"))

	if err := userDbInstance.CheckUser(user_id); err == nil {
		_, err := writer.Write([]byte(fmt.Sprintf("User - %v friends - %v\n", userDbInstance.GetUserName(user_id), userDbInstance.GetUserFriends(user_id))))
		if err != nil {
			log.Println(err)
		}
	} else {
		_, err := writer.Write([]byte(fmt.Sprintf("error! user not find - %v", user_id)))
		if err != nil {
			log.Println(err)
		}
	}
}

func handleUpdateAge(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("get update-age\n")

	reqBody, _ := io.ReadAll(request.Body)
	var newAge map[string]int
	json.Unmarshal(reqBody, &newAge)

	user_id, _ := strconv.Atoi(chi.URLParam(request, "user_id"))

	if err := userDbInstance.UpdateUserAge(user_id, newAge["new age"]); err != nil {
		_, err := writer.Write([]byte(fmt.Sprintf("error! user not exist - %v", user_id)))
		if err != nil {
			log.Println(err)
		}
		return
	}

	_, err := writer.Write([]byte(fmt.Sprintf("user - %v age has been updated", user_id)))
	if err != nil {
		log.Println(err)
	}
}

func handleShowUsers(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("get show\n")

	var rt []string
	for id, u := range userDbInstance.Db {
		rt = append(rt, fmt.Sprintf("id - %v name - %v, age - %v, friends - %v\n", id, u.Name, u.Age, u.Friends))
	}

	_, err := writer.Write([]byte(strings.Join(rt, "")))
	if err != nil {
		log.Println(err)
	}

}

func handleWannaBeTested(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("test me!\n")

	writer.WriteHeader(200)
	_, err := writer.Write([]byte("test me completely"))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	var (
		port int
	)

	flag.IntVar(&port, "port", 3333, "enter port")
	flag.Parse()

	userDbInstance.Init("db.json")
	defer userDbInstance.Close()

	router := chi.NewRouter()

	router.Post("/create", handleUserCreate)
	router.Post("/make_friends", handleMakeFriends)
	router.Delete("/user", handleUserDelete)
	router.Get("/friends/{user_id}", hanldeShowFriends)
	router.Put("/{user_id}", handleUpdateAge)
	router.Get("/show", handleShowUsers)
	router.Get("/test", handleWannaBeTested)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	fmt.Printf("Server listening at port - %v\n", port)
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error running http server: %s\n", err)
		}
	}
}
