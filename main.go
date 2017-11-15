package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	user string
)

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		printUsage()
	}

	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	for _, u := range users {
		result := getUsers(u)
		fmt.Println(`Username: `, result.Login)
		fmt.Println(`Name: `, result.Name)
		fmt.Println(`Email: `, result.Email)
		fmt.Println(`Bio: `, result.Bio)
		fmt.Println("")
	}
}

func init() {
	flag.StringVar(&user, "user", " ", "Search Users")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
