package main

var menu [5]string
var path string = "./wallet"
var key []byte = []byte("TZPtSIacEJG18IpqQSkTE6luYmnCNKgR")

type AccountEntry struct {
	Wallet   string
	Password []byte
}

var Entries []AccountEntry
var passwordsFile string = "passwords.json"


func main() {
	printMenu()
	chooseAction()
}
