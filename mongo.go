package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"fmt"
)

type Url struct {
	Url string		`bson:"url"`
	Success bool	`bson:"success"`
}

// book of urls
type Document struct {
	Book []Url `bson:"book"`
}

func main() {
	// establish db connection
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("urls")

	// simple insert, flat schema
	err = c.Insert(&Url{"http://github.com", true})
	if err != nil {
		log.Fatal(err)
	}

	// generate data for more complex schema insert
	u1 := Url{"http://www.google.com", true}
	u2 := Url{"http://www.yahoo.com", true}
	u3 := Url{"http://www.youtube.com", false}

	doc := Document{}
	doc.Book = append(doc.Book, u1)
	doc.Book = append(doc.Book, u2)
	doc.Book = append(doc.Book, u3)

	doc2 := Document{}
	doc2.Book = append(doc2.Book, u1)

	// insert to db collection
	err = c.Insert(doc)
	if err != nil {
		log.Fatal(err)
	}
	err = c.Insert(doc2)
	if err != nil {
		log.Fatal(err)
	}

	// query db
	result := []Document{}
	err = c.Find(bson.M{"book.url":"http://www.google.com"}).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	// print query results
	for _, doc := range result {
		for _, url := range doc.Book {
			fmt.Printf("url: %v\n", url)
		}
	}
}
