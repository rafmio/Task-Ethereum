package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)


func AccessToWallet(walletFileName string) {

	fmt.Println("Opening the wallet...")
	fmt.Println("Wallet's file name: ", walletFileName)
	walletFile, err := os.ReadFile(walletFileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The wallet file was read successfuly")
	fmt.Printf("Please enter wallet's password: ")

	// Password request
	var password string
	fmt.Scanf("%s\n", &password)

	checkPasswordResult := CheckPassword(password, walletFileName)
	if checkPasswordResult == false {
		fmt.Println("The password is incorrect")
		printMenu()
		chooseAction()
	}

	key, err := keystore.DecryptKey(walletFile, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("Access to wallet...")
	privData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Private:", hexutil.Encode(privData))

	pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Public: ", hexutil.Encode(pubData))

	addr := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	fmt.Println("Add: ", addr)
}
