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
	// dashboard should "all", "archived", "mine"
	// default "all"
	ds, err := c.GetDashboards("all")
	if err != nil {
		fmt.Printf("err in Get Dashboards: %v\n", err)
		return
	}
	fmt.Printf("Get Dashboards:\n %+v\n\n", ds)

	d, err := c.GetDashboard(1)
	if err != nil {
		fmt.Printf("err in Get Dashboard: %v\n", err)
		return
	}
	fmt.Printf("Get Dashboard:\n %+v\n\n", d)

	d2, err := c.GetDashboardRevisions(1)
	if err != nil {
		fmt.Printf("err in Get Dashboard Revisions: %v\n", err)
		return
	}
	fmt.Printf("Get Dashboard Revisions:\n %+v\n\n", d2)

}
