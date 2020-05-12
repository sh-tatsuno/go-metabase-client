package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	metabase "github.com/sh-tatsuno/go-metabase-client/metabase"
)

func randomGenerator(n int64) string {
	rand.Seed(time.Now().UnixNano())
	numPatterns := []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = numPatterns[rand.Intn(len(numPatterns))]
	}
	return string(b)
}

func main() {

	h := flag.String("host", "http://localhost:3000", "host address")
	u := flag.String("user", "sample@example.com", "user name")
	p := flag.String("pass", "passw0rd!", "password")

	flag.Parse()

	// Open connection
	c, err := metabase.NewClient(*h, *u, *p)
	if err != nil {
		fmt.Printf("err in NewClient: %+v\n", err)
		return
	}

	// Create User
	email := "create_sample_" + randomGenerator(10) + "@example.com"
	user := metabase.UserRequest{
		FirstName: "first",
		LastName:  "last",
		Email:     email,
		Password:  "passw0rd!",
	}
	u1, err := c.CreateUser(user)
	if err != nil {
		fmt.Printf("err in Create User: %v\n", err)
		return
	}
	fmt.Printf("Created User:\n %+v\n\n", u1)

	// Get Users
	u2, err := c.GetUsers(true)
	if err != nil {
		fmt.Printf("err in Get Users: %v\n", err)
		return
	}
	fmt.Printf("All Users:\n %+v\n\n", u2)

	// Get User
	u3, err := c.GetUser(u1.ID)
	if err != nil {
		fmt.Printf("err in Get User: %v\n", err)
		return
	}
	fmt.Printf("Get User:\n %+v\n\n", u3)

	// Get Current User
	u4, err := c.GetCurrentUser()
	if err != nil {
		fmt.Printf("err in Get Current User: %v\n", err)
		return
	}
	fmt.Printf("Get Current User:\n %+v\n\n", u4)

	// Delete User
	err = c.DeleteUser(u1.ID)
	if err != nil {
		fmt.Printf("err in Delete User: %v\n", err)
		return
	}
	fmt.Println("deleted")

}
