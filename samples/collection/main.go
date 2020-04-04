package main

import (
	"flag"
	"fmt"

	metabase "github.com/sh-tatsuno/go-metabase-client/metabase"
)

func newStringPointer(s string) *string {
	return &s
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
	cs, err := c.GetCollenctions()
	if err != nil {
		fmt.Printf("err in Get Collenctions: %v\n", err)
	}
	fmt.Printf("Get Collenctions:\n %+v\n\n", cs)

	cs2, err := c.GetCollenction(2)
	if err != nil {
		fmt.Printf("err in Get Collenction: %v\n", err)
	}
	fmt.Printf("Get Collenction:\n %+v\n\n", cs2)

	csr, err := c.GetRootCollenction()
	if err != nil {
		fmt.Printf("err in Get Root Collenctions: %v\n", err)
	}
	fmt.Printf("Get Root Collenctions:\n %+v\n\n", csr)

	csgp, err := c.GetCollenctionGraphPermission()
	if err != nil {
		fmt.Printf("err in Get Collenction Graph: %v\n", err)
	}
	fmt.Printf("Get Collenction Graph:\n %+v\n\n", csgp)

	// change struct
	if csgp.Groups["1"].Root == "write" {
		csgp.Groups["1"] = metabase.CollectionGraphPermissionGroup{Root: "none"}
	} else {
		csgp.Groups["1"] = metabase.CollectionGraphPermissionGroup{Root: "write"}
	}

	csgp2, err := c.UpdateCollenctionGraphPermission(*csgp)
	if err != nil {
		fmt.Printf("err in Update Collenction Graph: %v\n", err)
	}
	fmt.Printf("Update Collenction Graph:\n %+v\n\n", csgp2)

	cr := metabase.CollectionRequest{
		Name:        "test name",
		Description: newStringPointer("test collection"),
		Color:       newStringPointer("#AAAAAA"),
		ParentID:    nil,
	}
	cs3, err := c.CreateCollenction(cr)
	if err != nil {
		fmt.Printf("err in Create Collenction: %v\n", err)
	}
	fmt.Printf("Create Collenction:\n %+v\n\n", cs3)

	cu := metabase.CollectionPatch{
		ID:       cs3.ID,
		Name:     &cs3.Name,
		Color:    newStringPointer("#BBBBBB"),
		ParentID: nil,
	}
	cs4, err := c.UpdateCollenction(cu)
	if err != nil {
		fmt.Printf("err in Update Collenction: %v\n", err)
	}
	fmt.Printf("Update Collenction:\n %+v\n\n", cs4)

}
