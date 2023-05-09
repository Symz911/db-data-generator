package main

import (
	"example/go_fakeit/helper"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	// Connect to PostgreSQL
	DB := helper.ConnectToDatabase("host=localhost user=postgres password=Mhmmds911 dbname=biscord port=5432 sslmode=disable")

	// Gofakeit seed init (0 = use crypto/rand)
	gofakeit.Seed(0)

	// Main loop
	for {
		// Main menu
		mainMenuOption := helper.MainMenu()

		switch mainMenuOption {
		// Generate dummy Users data
		case 1:
			count := helper.GenerateUserMenu()
			helper.GenerateUserData(DB, &count)

		// Generate dummy Guilds data
		case 2:
			count := helper.GenerateGuildMenu()
			helper.GenerateGuildData(DB, &count)

		// Generate dummy Channels data for every guild
		case 3:
			count := helper.GenerateChannelMenu()
			helper.GenerateChannelData(DB, &count)

		// Generate dummy Members data
		case 4:
			count := helper.GenerateMemberMenu()
			helper.GenerateMemberData(DB, &count)

		// Generate dummy Messages data
		case 5:
			count := helper.GenerateMessageMenu()
			helper.GenerateMessageData(DB, &count)

		// Exit from program
		case 6:
			helper.ClearScreen()
			fmt.Println("Program Exitted!")
		default:
			helper.ClearScreen()
			fmt.Println("Invalid option, please check your option!")
			continue
		}

		// Check for exit program
		if mainMenuOption == 6 {
			break
		}
	}
}
