package main

import (
	"fmt"
	"os"
)

func printMenu() {
	menu = [5]string{
		"1 - Check for wallets",
		"2 - Create a new wallet",
		"3 - Access to a wallet",
		"4 - Delete wallet",
		"0 - Exit the program",
	}

	fmt.Printf("\nMENU\n")
	for _, value := range menu {
		fmt.Println(value)
	}
	fmt.Println()
}

func chooseAction() {
	action := -1
	checkAction := -1
	correctScanMarker := 1

	fmt.Printf("Choose action: ")

	for {
		checkAction, _ = fmt.Scanf("%d\n", &action)
		if checkAction != correctScanMarker || action > len(menu)-1 || action < 0 {
			printMenu()
			fmt.Println("Incorrect input. Pleace enter the menu's item")
		} else {
			switch action {
			case 1:
				CheckForKeystore()
			case 2:
				CreateKeyStore()
			case 3:
				CheckForKeystore()
			case 4:
				DeleteWallet()
			case 0:
				fmt.Println("Exiting the program. Bye!")
				os.Exit(0)
			}
			break
		}
	}
}
