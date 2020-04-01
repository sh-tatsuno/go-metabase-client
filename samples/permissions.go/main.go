package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	metabase "github.com/sh-tatsuno/go-metabase-client/metabase"
)

func randomGenerator(n int) string {
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

	// Get User
	x, err := c.GetPermissionsMemberships()
	if err != nil {
		fmt.Printf("err in Get Permission: %v\n", err)
	}
	fmt.Printf("Get Graph:\n %+v\n\n", x)

	r, err := c.GetPermissionsGraphs()
	if err != nil {
		fmt.Printf("err in Get Permission: %v\n", err)
	}
	fmt.Printf("Get Graph:\n %+v\n\n", r)

}
