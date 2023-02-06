package main

import (
	"fmt"
	"os"
)

func DeleteWallet() {
	var fileIndex int
	var addr string
	correctScanMarker := 1
	minimalRangeValue := 1

	fmt.Println()
	fmt.Println("Let's delete the wallet")
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("reading directory: ", err)
		fmt.Println("It seems you have no wallets at all. It's nothing to delete")
		printMenu()
		chooseAction()
	} else if len(files) == 0 {
		fmt.Println("It seems you have no wallets at all. It's nothing to delete")
		printMenu()
		chooseAction()
	} else {
		fmt.Println("Pleace choose wallet to delete")
		fmt.Printf("Index N  Address\n")
		for index, value := range files {
			fileName := value.Name()
			addr = fileName[(len(fileName) - 40):]
			fmt.Printf("%d \t %s\n", index+1, addr)
		}

		fmt.Printf("Please choose the wallet file's index to delete (1 - %d): ", len(files))

		for {
			checkAnswer, _ := fmt.Scanf("%d", &fileIndex)
			if checkAnswer != correctScanMarker || fileIndex < minimalRangeValue || fileIndex > len(files) {
				fmt.Printf("Incorrect input. Please choose the wallet file's index (1 - %d): ", len(files))
			} else {
				walletFileName := path + "/" + files[fileIndex-1].Name()
				fmt.Println("Deleting: ", walletFileName)
				err := os.Remove(walletFileName)
				if err != nil {
					fmt.Println("removing file: ", err)
					break
				} else {
					fmt.Printf("The wallet's file was deleted: %s", walletFileName)
					break
				}
			}
		}
		printMenu()
		chooseAction()
	}
}
