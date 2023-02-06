package main

import (
	"fmt"
	"os"
	"time"
)

func CheckForWallets(files []os.DirEntry) {
	var fileIndex int
	correctScanMarker := 1
	minimalRangeValue := 1

	fmt.Println("Searching for files...")
	time.Sleep(time.Second * 1)
	if len(files) == 0 { // if directory is empty
		fmt.Println("Directory is empty")
		printMenu()
		chooseAction()
	} else {
		fmt.Println("The following files were found:")
		printMenu()
		for index, value := range files {
			fmt.Printf("index: %d - name: %s\n", index+1, value.Name())
		}
	}

	fmt.Printf("Please choose the wallet file's index (1 - %d): ", len(files))

	for {
		checkAnswer, _ := fmt.Scanf("%d", &fileIndex)
		if checkAnswer != correctScanMarker || fileIndex < minimalRangeValue || fileIndex > len(files) {
			fmt.Printf("Incorrect input. Please choose the wallet file's index (1 - %d): ", len(files))
		} else {
			walletFileName := path + "/" + files[fileIndex-1].Name()
			AccessToWallet(walletFileName)
			
			printMenu()
			chooseAction()
		}
	}
}
