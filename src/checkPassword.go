package main

import (
	"fmt"
)

func CheckPassword(password string, walletFileName string) bool {
	fmt.Println("Checking password...")

	addr := walletFileName[(len(walletFileName) - 40):]

	fmt.Println("Decrypting AES...")
	decryptedPassword := DecryptAES(addr)

	if password == decryptedPassword {
		return true
	} else {
		return false
	}
}
