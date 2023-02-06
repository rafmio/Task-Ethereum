package main

import (
	"fmt"
	"os"
)

func CheckForKeystore() {

	correctScanMarker := 1
	var answerYN string

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("reading directory: ", err)
		fmt.Println("It seems to be no keystore yet. Do you want to create a keystore? (y/n): ")
		checkAnswer, _ := fmt.Scanf("%s", &answerYN)
		if checkAnswer != correctScanMarker {
			fmt.Println("Incorrect input. Please check the action")
			printMenu()
			chooseAction()
		} else if answerYN == "n" || answerYN == "N" {
			printMenu()
			chooseAction()
		} else if answerYN == "y" || answerYN == "Y" {
			CreateKeyStore()
		} else {
			printMenu()
			chooseAction()
		}
	} else {
		CheckForWallets(files)
	}
}
