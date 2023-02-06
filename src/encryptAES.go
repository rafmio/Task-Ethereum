package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func EncryptAES(password string, accAddr string) {
	fmt.Println("Let's encrypt AES-password for wallet's address: ", accAddr)
	ciphBlock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("creating the cipher block: ", err)
	}
	gcm, err := cipher.NewGCM(ciphBlock)
	if err != nil {
		fmt.Println("wrape the AES in Galios Counter Mode: ", err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("reading bytes: ", err)
	}

	encryptedAESPassword := gcm.Seal(nonce, nonce, []byte(password), nil)
	fmt.Println("Password was encrypted to AES")

	EncodeJSON(encryptedAESPassword, accAddr)
}

func EncodeJSON(encryptedAESPassword []byte, accAddr string) {
	fmt.Println("Now let's encode encrypted password to JSON and write to file")
	acc := AccountEntry{Wallet: accAddr, Password: encryptedAESPassword}
	Entries := []AccountEntry{acc}
	json_data, err := json.Marshal(Entries)
	if err != nil {
		fmt.Println("marshalling data: ", err.Error())
	}
	file, err := os.ReadFile(passwordsFile)
	if err != nil {
		fmt.Println("reading file", err.Error())
		fmt.Printf("Creting a new %s file\n", passwordsFile)
		err = os.WriteFile(passwordsFile, json_data, 0644)
		if err != nil {
			fmt.Println("writing to file: ", err.Error())
		}
	} else {
		fmt.Printf("Success: file %s is exist\n", passwordsFile)

		var unmarshalEntries []AccountEntry
		err = json.Unmarshal(file, &unmarshalEntries)
		if err != nil {
			fmt.Println("unmarshaling file: ", err.Error())
		}

		if len(unmarshalEntries) > 0 {
			unmarshalEntries = append(unmarshalEntries, acc)
		}

		jsonDataForWrite, err := json.Marshal(unmarshalEntries)
		if err != nil {
			fmt.Println("marshalling data: ", err.Error())
		}

		err = os.WriteFile(passwordsFile, jsonDataForWrite, 0644)
		if err != nil {
			fmt.Println("writing to file: ", err.Error())
		} else {
			fmt.Println("JSON entry successfully completed")
		}
	}
}
