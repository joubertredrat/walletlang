package main

import (
	"github.com/joho/godotenv"
	"github.com/joubertredrat/walletlang/cmd"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
