package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateKeyStore() {
	var password string
	var accAddr string

	fmt.Println("Creating a keystore...")
	time.Sleep(time.Second * 1)
	key := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	fmt.Printf("The keystore '%s' has just been created\n", strings.Split(path, "/")[1])

	fmt.Println("And now let's create an account (wallet)")
	fmt.Printf("Please enter your password: ")
	fmt.Scanf("%s\n", &password)

	account, err := key.NewAccount(password)
	if err != nil {
		fmt.Println("creating a new account: ", err.Error())
	} else {
		fmt.Println()
		fmt.Println("The new account was created")
		fmt.Println("Account address: ", account.Address)
		accAddr = account.Address.String()
		fmt.Println()
		
		EncryptAES(password, accAddr)

	}

	printMenu()
	chooseAction()
}
