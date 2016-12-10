// UseMongoDB project main.go
package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	checkErr(err)
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"}, &Person{"Cla", "+55 53 8402 8510"})
	checkErr(err)

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	checkErr(err)

	fmt.Println("Phone:", result.Phone)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
