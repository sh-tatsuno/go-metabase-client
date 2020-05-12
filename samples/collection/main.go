package main

import (
	"flag"
	"fmt"

	metabase "github.com/sh-tatsuno/go-metabase-client/metabase"
)

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

	// Get Collection
	cs, err := c.GetCollections()
	if err != nil {
		fmt.Printf("err in Get Collections: %v\n", err)
		return
	}
	fmt.Printf("Get Collections:\n %+v\n\n", cs)

	cs2, err := c.GetCollection(2)
	if err != nil {
		fmt.Printf("err in Get Collection: %v\n", err)
		return
	}
	fmt.Printf("Get Collection:\n %+v\n\n", cs2)

	csr, err := c.GetRootCollection()
	if err != nil {
		fmt.Printf("err in Get Root Collections: %v\n", err)
		return
	}
	fmt.Printf("Get Root Collections:\n %+v\n\n", csr)

	csgp, err := c.GetCollectionGraphPermission()
	if err != nil {
		fmt.Printf("err in Get Collection Graph Permission: %v\n", err)
		return
	}
	fmt.Printf("Get Collection Graph Permission:\n %+v\n\n", csgp)

	// change struct
	if csgp.Groups["1"].Root == "write" {
		csgp.Groups["1"] = metabase.CollectionGraphPermissionGroup{Root: "none"}
	} else {
		csgp.Groups["1"] = metabase.CollectionGraphPermissionGroup{Root: "write"}
	}

	csgp2, err := c.UpdateCollectionGraphPermission(*csgp)
	if err != nil {
		fmt.Printf("err in Update Collection Graph Permission: %v\n", err)
		return
	}
	fmt.Printf("Update Collection Graph Permission:\n %+v\n\n", csgp2)

	cr := metabase.CollectionRequest{
		Name:        "test name",
		Description: "test collection",
		Color:       "#AAAAAA",
	}
	cs3, err := c.CreateCollection(cr)
	if err != nil {
		fmt.Printf("err in Create Collection: %v\n", err)
		return
	}
	fmt.Printf("Create Collection:\n %+v\n\n", cs3)

	cu := metabase.CollectionPatch{
		ID:    cs3.ID,
		Name:  cs3.Name,
		Color: "#BBBBBB",
	}
	cs4, err := c.UpdateCollection(cu)
	if err != nil {
		fmt.Printf("err in Update Collection: %v\n", err)
		return
	}
	fmt.Printf("Update Collection:\n %+v\n\n", cs4)

}
