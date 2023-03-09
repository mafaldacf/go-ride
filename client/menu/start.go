package menu

import (
	"context"
	"fmt"

	pb "go-ride/proto"
)

func ShowStartMenu(client *pb.RideServiceClient) {
	var username, authToken string
	var ok bool

	for {
		var option string

		fmt.Println()
		fmt.Println("------- START MENU -------")
		fmt.Println("Select one option:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("0. Exit")
		fmt.Println("--------------------------")
		fmt.Print("\n> ")

		fmt.Scanf("%s", &option)
		switch option {
		case "1":
			username, authToken, ok = register(client)
			if ok {
				ShowClientMenu(client, username, authToken)
			}
		case "2":
			username, authToken, ok = login(client)
			if ok {
				ShowClientMenu(client, username, authToken)
			}
		case "0":
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}

/* --------

Helpers

---------- */

func readUser() (string, string) {
	var username, password string

	fmt.Print("Username: ")
	fmt.Scanf("%s", &username)
	fmt.Print("Password: ")
	fmt.Scanf("%s", &password)

	return username, password
}

/* --------

API calls

---------- */

func register(client *pb.RideServiceClient) (string, string, bool) {
	username, password := readUser()

	request := &pb.RegisterRequest{
		Username: username,
		Password: password,
	}

	resp, err := (*client).Register(context.Background(), request)

	if err != nil {
		fmt.Println(err)
		return "", "", false
	}
	return username, resp.AuthToken, true
}

func login(client *pb.RideServiceClient) (string, string, bool) {
	username, password := readUser()

	request := &pb.LoginRequest{
		Username: username,
		Password: password,
	}

	resp, err := (*client).Login(context.Background(), request)

	if err != nil {
		fmt.Println(err)
		return "", "", false
	}
	return username, resp.AuthToken, true
}
