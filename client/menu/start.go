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
			username, authToken, ok = registerHelper(client)
			if ok {
				ShowClientMenu(client, username, authToken)
			}
		case "2":
			username, authToken, ok = loginHelper(client)
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

func registerHelper(client *pb.RideServiceClient) (string, string, bool) {
	var username, password string

	fmt.Print("Username: ")
	fmt.Scanf("%s", &username)

	fmt.Print("Password: ")
	fmt.Scanf("%s", &password)

	authToken, ok := register(client, username, password)
	return username, authToken, ok
}

func loginHelper(client *pb.RideServiceClient) (string, string, bool) {
	var username, password string

	fmt.Print("Username: ")
	fmt.Scanf("%s", &username)
	fmt.Print("Password: ")
	fmt.Scanf("%s", &password)

	authToken, ok := login(client, username, password)
	return username, authToken, ok
}

func register(client *pb.RideServiceClient, username string, password string) (string, bool) {
	ctx := context.Background()

	request := &pb.RegisterRequest{
		Username: username,
		Password: password,
	}

	resp, err := (*client).Register(ctx, request)

	if err != nil {
		fmt.Println(err)
		return "", false
	}
	return resp.AuthToken, true
}

func login(client *pb.RideServiceClient, username string, password string) (string, bool) {
	ctx := context.Background()

	request := &pb.LoginRequest{
		Username: username,
		Password: password,
	}

	resp, err := (*client).Login(ctx, request)

	if err != nil {
		fmt.Println(err)
		return "", false
	}
	return resp.AuthToken, true
}
