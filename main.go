package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"test-binding/binding"
	"time"
)

type User struct {
	ID     int64
	Name   string
	Age    int
	Gender string
	Email  string
	DOB    time.Time
}

func (u *User) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&u.ID:     "id",
		&u.Name:   "name",
		&u.Age:    "age",
		&u.Gender: "gender",
		&u.Email:  "email",
		&u.DOB: binding.Field{
			Form:       "dob",
			TimeFormat: time.DateOnly,
		},
	}
}

func main() {
	run1()
}

func run1() {
	user := &User{}
	req := httptest.NewRequest("GET", "http://example.com/?name=mana&age=63&gender=male&dob=1959-11-23", nil)
	errs := binding.Bind(req, user)
	if errs != nil {
		log.Fatal(errs)
	} else {
		fmt.Printf("%+v\n", user)
	}
}

func run2() {
	t := "1959-11-23 15:04"
	tf := time.DateOnly
	v, err := time.Parse(tf, t)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(v)
	}
}
