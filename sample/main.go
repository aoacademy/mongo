/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 17-04-2018
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

func main() {
	fmt.Println("vim-go")

	client, err := mgo.Dial("mongodb://127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019/?replicaSet=shamin")
	if err != nil {
		log.Fatal(err)
	}

	db := client.DB("test")
	c := db.C("nice")

	// Insert
	s := struct {
		Name   string
		Family string
		Time   time.Time
	}{
		Name:   "Parham",
		Family: "Alvani",
		Time:   time.Now(),
	}
	fmt.Println(s)

	if err := c.Insert(s); err != nil {
		log.Fatal(err)
	}
}
