package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Platform struct {
	Name string
	Company string
}

func main() {
	// Connect to MongoDB server to get a session:
	session, err := mgo.Dial("localhost:27017")
	checkErr(err)
	defer session.Close()
	//session.SetMode(mgo.Monotonic, true)
	
	// Select DB and collection:
	collection := session.DB("test").C("platform")
	
	// Insert values:
	err = collection.Insert(&Platform{"Android", "Google"}, &Platform{"iOS", "Apple"})
	checkErr(err)
	
	// Query:
	result := Platform{}
	err = collection.Find(bson.M{"name": "Android"}).One(&result)
	checkErr(err)
	fmt.Println("Company:", result.Company)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}