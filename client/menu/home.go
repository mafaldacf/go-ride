package menu

import (
	"bufio"
	"context"
	"fmt"
	"os"

	pb "go-ride/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ShowClientMenu(client *pb.RideServiceClient, username string, authToken string) {
	fmt.Printf("\nWelcome %s!\n", username)
	show_options := true

	for {
		var option string
		var err error

		if show_options {
			fmt.Println()
			fmt.Println("------- HOME MENU -------")
			fmt.Println("[INFO] Zones: Alpha <--> Bravo <--> Charlie <--> Delta <--> Echo <--> Foxtrot")
			fmt.Println("[INFO] To pickup any bike you need at least 5\u20AC")
			fmt.Println("[INFO] Each step costs 1\u20AC")
			fmt.Println()
			fmt.Println("Select one option:")
			fmt.Println("1. Check Balance")
			fmt.Println("2. Add Funds")
			fmt.Println("3. Pickup Bike")
			fmt.Println("4. Drop Bike")
			fmt.Println("5. Move")
			fmt.Println("6. Info Client")
			fmt.Println("7. Info Zones")
			fmt.Println("0. Logout")
			fmt.Println("-------------------------")
			fmt.Print("\n> ")
			show_options = false
		} else {
			fmt.Println("Select one option or type 'help' to show options:")
			fmt.Print("> ")
		}

		fmt.Scan(&option)
		switch option {
		case "1":
			err = checkBalance(client, authToken)
		case "2":
			err = addFunds(client, authToken)
		case "3":
			err = pickupBike(client, authToken)
		case "4":
			err = dropBike(client, authToken)
		case "5":
			err = move(client, authToken)
		case "6":
			err = infoClient(client, authToken)
		case "7":
			err = infoZones(client, authToken)
		case "0":
			fmt.Printf("\nGoodbye %s!\n", username)
			logout(client, authToken)
			return
		case "help":
			show_options = true
		default:
			fmt.Println("Invalid option!")
		}

		if err != nil {
			st, ok := status.FromError(err)
			if ok && st.Code() == codes.Unauthenticated {
				return // client is unauthenticated (auth token may have expired) and needs to login again
			}
		}

	}
}

/* --------

Helpers

---------- */

func readZone() (string, bool) {
	zone := ""

	fmt.Print("Select your desired zone (Alpha, Bravo, Charlie, Delta, Echo or Foxtrot): ")
	_, err := fmt.Scanf("%s", &zone)

	if err != nil {
		fmt.Println("ERROR: Invalid zone format")
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadString('\n') // discard any remaining input
		return zone, false
	}

	// cannot do this in a cleaner way :(
	if zone != "Alpha" && zone != "Bravo" && zone != "Charlie" && zone != "Delta" && zone != "Echo" && zone != "Foxtrot" {
		fmt.Println("ERROR: Invalid zone")
		return zone, false
	}

	return zone, true
}

/* --------

API calls

---------- */

func checkBalance(client *pb.RideServiceClient, authToken string) error {
	request := &pb.AuthRequest{
		AuthToken: authToken,
	}

	resp, err := (*client).CheckBalance(context.Background(), request)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("[Check Balance] balance:%v\u20AC\n", resp.Balance) //print balance + euro symbol
	return nil
}

func addFunds(client *pb.RideServiceClient, authToken string) error {
	var amount uint

	fmt.Print("[Add Funds] Amount: ")
	_, err := fmt.Scanf("%d", &amount)

	if err != nil {
		fmt.Println("ERROR: Invalid amount format")
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadString('\n') // discard any remaining input
		return nil
	}

	request := &pb.AddFundsRequest{
		AuthToken: authToken,
		Amount:    uint32(amount),
	}

	resp, err := (*client).AddFunds(context.Background(), request)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("[Add Funds] balance:%v\u20AC\n", resp.Balance) //print balance + euro symbol
	return nil
}

func infoClient(client *pb.RideServiceClient, authToken string) error {
	request := &pb.AuthRequest{
		AuthToken: authToken,
	}

	resp, err := (*client).InfoClient(context.Background(), request)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("[Info Client] onBike:%v, atZone:'%v'", resp.OnBike, resp.AtZone)
	return nil
}

func infoZones(client *pb.RideServiceClient, authToken string) error {
	request := &pb.AuthRequest{
		AuthToken: authToken,
	}

	resp, err := (*client).InfoZones(context.Background(), request)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("[Info Zones]")
	for _, zone := range resp.Zones {
		fmt.Printf("Zone:'%v' \t Bikes:%v \t Capacity %v\n", zone.Name, zone.Bikes, zone.Capacity)
	}

	return nil
}

func pickupBike(client *pb.RideServiceClient, authToken string) error {
	fmt.Print("[Pickup Bike] ")
	zone, ok := readZone()

	if !ok {
		return nil
	}

	request := &pb.PickupBikeRequest{
		AuthToken: authToken,
		Zone:      zone,
	}

	_, err := (*client).PickupBike(context.Background(), request)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("[Pickup Bike] Success! You are now riding a bike at %s\n", zone)
	return nil
}

func dropBike(client *pb.RideServiceClient, authToken string) error {
	request := &pb.AuthRequest{
		AuthToken: authToken,
	}

	_, err := (*client).DropBike(context.Background(), request)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("[Drop Bike] Success!")
	return nil
}

func move(client *pb.RideServiceClient, authToken string) error {
	fmt.Print("[Move] ")
	zone, ok := readZone()

	if !ok {
		return nil
	}

	request := &pb.MoveRequest{
		AuthToken: authToken,
		Zone:      zone,
	}

	_, err := (*client).Move(context.Background(), request)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("[Move] Success! You are now at %s\n", zone)
	return nil
}

func logout(client *pb.RideServiceClient, authToken string) {
	request := &pb.AuthRequest{
		AuthToken: authToken,
	}

	(*client).Logout(context.Background(), request)
}
