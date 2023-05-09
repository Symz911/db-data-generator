package helper

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	// Menjalankan perintah "cls" pada terminal
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func GetInput(prompt string) int {
    var input int

    for {
        fmt.Print(prompt)
        _, err := fmt.Scan(&input)
        if err != nil {
            fmt.Println("Invalid input, please try again!")
            continue
        }
        break
    }

    return input
}

func MainMenu() int8 {
	var option int8

	for {
		fmt.Println("Generate Dummy Data")
		fmt.Println("1. Users data")
		fmt.Println("2. Guilds data")
		fmt.Println("3. Channels data")
		fmt.Println("4. Members data")
		fmt.Println("5. Messages data")
		fmt.Println("6. Exit")
		fmt.Print("Select option: ")

		_, err := fmt.Scan(&option)
		if err != nil {
			ClearScreen()
			fmt.Println("Invalid input, please try again!")
			continue
		}

		break
	}

	return option
}

func GenerateUserMenu() int {
	count := GetInput("How many Users you want to generate: ")
	return count
}

func GenerateGuildMenu() int {
	count := GetInput("How many Guilds you want to generate: ")
	return count
}

func GenerateChannelMenu() int {
	count := GetInput("How many Channels you want to generate: ")
	return count
}

func GenerateMemberMenu() int {
	count := GetInput("How many Members you want to generate: ")
	return count
}

func GenerateMessageMenu() int {
	count := GetInput("How many Messages you want to generate: ")
	return count
}