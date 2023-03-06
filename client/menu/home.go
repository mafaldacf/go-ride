package menu

import (
	"bufio"
	"context"
	"fmt"
	"os"

	pb "go-ride/proto"
)

func ShowClientMenu(client *pb.RideServiceClient, username string, authToken string) {
	fmt.Printf("\nWelcome %s!\n", username)

	for {
		var option string

		fmt.Println()
		fmt.Println("------- HOME MENU -------")
		fmt.Println("Select one option:")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Add Funds")
		fmt.Println("0. Logout")
		fmt.Println("-------------------------")
		fmt.Print("\n> ")

		fmt.Scan(&option)
		switch option {
		case "1":
			checkBalance(client, authToken)
		case "2":
			addFundsHelper(client, authToken)
		case "0":
			fmt.Printf("\nGoodbye %s!\n", username)
			logout(client, authToken)
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}

func addFundsHelper(client *pb.RideServiceClient, authToken string) {
	var amount uint32

	fmt.Print("Amount: ")
	_, err := fmt.Scanf("%d", &amount)

	if err != nil {
		fmt.Println("ERROR: Invalid Amount")
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadString('\n') // discard any remaining input
	} else {
		addFunds(client, authToken, amount)
	}
}

func checkBalance(client *pb.RideServiceClient, authToken string) {
	ctx := context.Background()

	request := &pb.CheckBalanceRequest{
		AuthToken: authToken,
	}

	resp, err := (*client).CheckBalance(ctx, request)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("[Check Balance] balance:%v\u20AC\n", resp.Balance) //print balance + euro symbol
}

func addFunds(client *pb.RideServiceClient, authToken string, amount uint32) {
	ctx := context.Background()

	request := &pb.AddFundsRequest{
		AuthToken: authToken,
		Amount:    amount,
	}

	resp, err := (*client).AddFunds(ctx, request)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("[Add Funds] balance:%v\u20AC\n", resp.Balance) //print balance + euro symbol
}

func logout(client *pb.RideServiceClient, authToken string) {
	ctx := context.Background()

	request := &pb.LogoutRequest{
		AuthToken: authToken,
	}

	(*client).Logout(ctx, request)
}
