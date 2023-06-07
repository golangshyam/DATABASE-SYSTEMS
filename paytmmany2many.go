package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

type Paytm struct {
	Moneytransfer     string
	Billpayments      string
	Loans_Creditcards string
	Ticketbooking     string
}

type Movietickets struct {
	gorm.Model
	Moviename string
	Seatno    string
	City      []Location `json:"city" gorm:"foreignkey:Refer"`
}

type Location struct {
	gorm.Model
	Refer  uint
	Name   string
	Street string
	Amount []Movietickets `gorm:"many2many:movietickets_location"`
}

type MovieResponse struct {
	Status string `json:"status"`
}

type Entertainment interface {
	Funreq(c *gin.Context)
}

type Cinema struct {
	db *gorm.DB
}

// 3. connect to db
// 4. create table

func Largescreen() Entertainment {

	db, err := gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/paytm?charset=utf8&parseTime=True")

	if err != nil {

		fmt.Print("jai balayya", err)

		panic("db not connected")
	}

	db.AutoMigrate(&Movietickets{}, &Location{})

	return &Cinema{db}

}

func Titles(req *gin.Context) *Paytm {

	var inputs Paytm

	req.ShouldBind(&inputs)

	return &inputs

}

func Interval(inputs *Paytm, req *Cinema) bool {

	fmt.Println(inputs)

	req.db.Create(&Movietickets{Moviename: "bumchick", Seatno: "44 to 45", City: []Location{{Name: "hyderabad", Street: "satyam"}}})

	return true

}

func (req *Cinema) Funreq(c *gin.Context) {

	//  Inside api
	// 1. validation
	// 2. query
	// 3. logic

	inputs := Titles(c)

	if inputs == nil {
		c.JSON(201, "invalid input")
		return
	}

	fmt.Println(inputs)

	resp := Interval(inputs, req)

	if resp == false {
		c.JSON(201, "tickets not avaliable house full ")
		return
	}

	var Response MovieResponse = MovieResponse{"box office bumper hit"}

	c.AsciiJSON(200, Response)
}

func main() {

	//  1. server
	//  2. register api to server

	r := gin.Default()

	obj := Largescreen()

	v1 := r.Group("/jai")
	{
		v1.GET("/balayya", obj.Funreq)
	}

	r.Run(":9090")

}
