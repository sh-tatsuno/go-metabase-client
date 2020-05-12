package main

import (
	"flag"
	"fmt"
	"math/rand"
	"reflect"
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

	// Get User
	gpm, err := c.GetPermissionsMemberships()
	if err != nil {
		fmt.Printf("err in Get Permission Membership: %v\n", err)
		return
	}
	fmt.Printf("Get Permission Membership:\n %+v\n\n", gpm)

	// Get User Permission
	gpg, err := c.GetPermissionsGraphs()
	if err != nil {
		fmt.Printf("err in Get Permission Graph: %v\n", err)
		return
	}
	fmt.Printf("Get Permission Graph:\n %+v\n\n", gpg)

	// update status
	// switch allow <-> deny permission of group 1 for graph 1
	allowPermission := metabase.BulkPermission{
		Native:  "write",
		Schemas: "all",
	}

	denyPermission := metabase.BulkPermission{
		Native: "none",
	}

	if reflect.DeepEqual(gpg.Groups["1"]["1"], &allowPermission) {
		gpg.Groups["1"]["1"] = &denyPermission
	} else {
		gpg.Groups["1"]["1"] = &allowPermission
	}

	upg, err := c.UpdatePermissionsGraph(*gpg)
	if err != nil {
		fmt.Printf("err in Update Permission: %v\n", err)
		return
	}
	fmt.Printf("Update Permission:\n %+v\n\n", upg)

}
