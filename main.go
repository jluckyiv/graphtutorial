package main

import (
	"fmt"
	"log"
	// "time"

	"github.com/jluckyiv/graphtutorial/graphhelper"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Go Graph Tutorial")
	fmt.Println()

	// Load .env files
	// .env.local takes precedence (if present)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	graphHelper := graphhelper.NewGraphHelper()
	initializeGraph(graphHelper)

	greetUser(graphHelper)

	var choice int64 = -1

	for {
		fmt.Println("Please choose an option:")
		fmt.Println("0. Exit")
		fmt.Println("1. Display access token")
		fmt.Println("2. List my inbox")
		fmt.Println("3. Send mail")
		fmt.Println("4. List my calendar")
		fmt.Println("5. Make a graph call")

		_, err = fmt.Scanf("%d", &choice)
		if err != nil {
			choice = -1
		}

		switch choice {
		case 0:
			// Exit
			fmt.Println("Goodbye.")
		case 1:
			// Display access token
			displayAccessToken(graphHelper)
		case 2:
			// List emails from user's inbox
			listInbox(graphHelper)
		case 3:
			// Send an email
			sendMail(graphHelper)
		case 4:
			// List calendar
			listCalendar(graphHelper)
		case 5:
			// Run any Graph code
			makeGraphCall(graphHelper)
		default:
			fmt.Println("Invalid choice!")
		}
		if choice == 0 {
			break
		}
	}
}

func initializeGraph(graphHelper *graphhelper.GraphHelper) {
	err := graphHelper.InitializeGraphForUserAuth()
	if err != nil {
		log.Panicf("Error initializing Graph for user auth: %v\n", err)
	}
}

func greetUser(graphHelper *graphhelper.GraphHelper) {
	user, err := graphHelper.GetUser()
	if err != nil {
		log.Panicf("Error getting user: %v\n", err)
	}

	fmt.Printf("Hello, %s!\n", *user.GetDisplayName())

	email := user.GetMail()
	if email == nil {
		email = user.GetUserPrincipalName()
	}

	fmt.Printf("Email: %s\n", *email)
	fmt.Println()
}

func displayAccessToken(graphHelper *graphhelper.GraphHelper) {
	token, err := graphHelper.GetUserToken()
	if err != nil {
		log.Panicf("Error getting user token: %v\n", err)
	}

	fmt.Printf("User token: %s", *token)
	fmt.Println()
}

func listInbox(graphHelper *graphhelper.GraphHelper) {
	// TODO
}

func sendMail(graphHelper *graphhelper.GraphHelper) {
	// TODO
}

func listCalendar(graphHelper *graphhelper.GraphHelper) {
	// TODO
}

func makeGraphCall(graphHelper *graphhelper.GraphHelper) {
	// TODO
}
