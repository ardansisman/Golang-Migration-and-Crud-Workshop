package main

import "github.com/joho/godotenv"

func main() {
	if err := godotenv.Load("local.env"); err != nil {
		panic("Error loading .env file")
	}
}
