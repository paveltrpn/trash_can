package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"skillfactory-go-developer/pkg/user"
)

var (
	reqPort int = 3333
)

func makeAndSendReq(data []byte, method string, port int, url string) error {
	reqUrl := fmt.Sprintf("http://localhost:%d%v", port, url)
	req, err := http.NewRequest(method, reqUrl, bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer req.Body.Close()

	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Charset":      {"utf-8"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Println(res.Status)
	resBody, _ := io.ReadAll(res.Body)
	fmt.Println(string(resBody))

	return nil
}

func requestUserCreate(usr user.User, port int) {
	reqJson, err := json.Marshal(usr)
	if err != nil {
		log.Fatalf("requestUserCreate(): error marshalling json - %v\n", err)
	}

	err = makeAndSendReq(reqJson, http.MethodPost, port, "/create")
	if err != nil {
		log.Fatalf("requestUserCreate(): error sending request - %v\n", err)
	}
}

func requestMakeFriends(srcId, tgtId int, port int) {
	reqJson, err := json.Marshal(map[string]any{"source_id": srcId,
		"target_id": tgtId,
	})
	if err != nil {
		log.Fatalf("requestMakeFriends(): error marshalling json - %v\n", err)
	}

	err = makeAndSendReq(reqJson, http.MethodPost, port, "/make_friends")
	if err != nil {
		log.Fatalf("requestMakeFriends(): error sending request - %v\n", err)
	}
}

func requestUserDelete(tgtId int, port int) {
	reqJson, err := json.Marshal(map[string]any{"target_id": tgtId})
	if err != nil {
		log.Fatalf("requestUserDelete(): error marshalling json - %v\n", err)
	}

	err = makeAndSendReq(reqJson, http.MethodDelete, port, "/user")
	if err != nil {
		log.Fatalf("requestUserDelete(): error sending request - %v\n", err)
	}
}

func requsetShowFriends(user_id int, port int) {
	err := makeAndSendReq([]byte{}, http.MethodGet, port, fmt.Sprintf("/friends/%v", user_id))

	if err != nil {
		log.Fatalf("requestUpdateAge(): error sending request - %v\n", err)
	}
}

func requestUpdateAge(target_id, newAge int, port int) {
	reqJson, err := json.Marshal(map[string]int{"new age": newAge})
	if err != nil {
		log.Fatalf("requestUpdateAge(): error marshalling json - %v\n", err)
	}

	err = makeAndSendReq(reqJson, http.MethodPut, port, fmt.Sprintf("/%v", target_id))
	if err != nil {
		log.Fatalf("requestUpdateAge(): error sending request - %v\n", err)
	}
}

func requestShowUsers(port int) {
	err := makeAndSendReq([]byte{}, http.MethodGet, port, "/show")
	if err != nil {
		log.Fatalf("requestUpdateAge(): error sending request - %v\n", err)
	}
}

func main() {
	var (
		action string
	)

	flag.IntVar(&reqPort, "port", 3333, "enter port")
	flag.Parse()

	fmt.Printf("request sending to port - %v\n", reqPort)

	for {
		fmt.Printf("choose action - create make-friends delete show-friends update-age show\n")
		fmt.Scan(&action)

		switch action {
		case "create":
			fmt.Printf("enter Name (string) and Age (number)\n")

			var (
				name    string
				age     int
				friends []int
			)

			fmt.Scanf("%v %v", &name, &age)
			friends = make([]int, 0)

			requestUserCreate(user.User{Name: name, Age: age, Friends: friends}, reqPort)

		case "make-friends":
			fmt.Printf("enter user id's - source_id (number) and target_id (number)\n")

			var source_id, target_id int

			fmt.Scanf("%v %v", &source_id, &target_id)
			requestMakeFriends(source_id, target_id, reqPort)

		case "delete":
			fmt.Printf("enter user id (number)\n")

			var target_id, newAge int

			fmt.Scanf("%v %v", &target_id, &newAge)
			requestUserDelete(target_id, reqPort)

		case "show-friends":
			fmt.Printf("enter user id (number)\n")

			var user_id int

			fmt.Scanf("%v", &user_id)
			requsetShowFriends(user_id, reqPort)

		case "update-age":
			fmt.Printf("enter user id (number) and new age (number)\n")

			var target_id, newAge int

			fmt.Scanf("%v %v", &target_id, &newAge)
			requestUpdateAge(target_id, newAge, reqPort)
		case "show":
			requestShowUsers(reqPort)

		default:
			fmt.Printf("unknown action - %v\n", action)
		}
	}
}
