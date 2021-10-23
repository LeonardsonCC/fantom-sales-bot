package main

import (
  "fmt"
  "os"
)

func main() {
  fmt.Println("Starting...")

  consumer_key := os.Getenv("CONSUMER_KEY")
  consumer_secret := os.Getenv("CONSUMER_SECRET")
  access_token := os.Getenv("ACCESS_TOKEN")
  access_secret := os.Getenv("ACCESS_SECRET")

  fmt.Println(fmt.Sprintf("%v | %v | %v | %v", consumer_key, consumer_secret, access_token, access_secret))
}
