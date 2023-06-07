package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

type Credapp struct {
	Creditcard string
	Rentpay    string
	Shopping   string
}

type Billpay struct {
	gorm.Model
	Postpaidbill string
	Creditcards  []Cards `json:"creditcards" gorm:"foreignkey:Refer"`
}

type Cards struct {
	gorm.Model
	Refer    uint
	Name     string
	Amount   string
	Shopping string
}

type Payment_Response struct {
	Status string `json:"status"`
}

type Payment interface {
	Paymentreq(c *gin.Context)
}

type Bills struct {
	db *gorm.DB
}

// 3. connect to db
// 4. create table

func Construct() Payment {

	db, err := gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/cred?charset=utf8&parseTime=True")

	if err != nil {

		fmt.Print("shyam", err)

		panic("db not connected")
	}

	db.AutoMigrate(&Billpay{}, &Cards{})

	return &Bills{db}

}

func Inputvalid(req *gin.Context) *Credapp {

	var inputs Credapp

	req.ShouldBind(&inputs)

	return &inputs

}

func APIvalidat(inputs *Credapp, req *Bills) bool {

	fmt.Println(inputs)

	req.db.Create(&Billpay{Postpaidbill: "jio", Creditcards: []Cards{{Name: "HDFC", Amount: "36000RS", Shopping: "washing machine"}, {Name: "SBI", Amount: "16000RS", Shopping: "mobile phone"}}})

	return true

}

func (req *Bills) Paymentreq(c *gin.Context) {

	// 5. Inside api
	// above 3 steps

	inputs := Inputvalid(c)

	if inputs == nil {
		c.JSON(201, "invalid input")
		return
	}

	fmt.Println(inputs)

	resp := APIvalidat(inputs, req)

	if resp == false {
		c.JSON(201, "invalid input details ")
		return
	}

	var Response Payment_Response = Payment_Response{"payment status"}

	c.AsciiJSON(200, Response)
}

func main() {

	//  1. server
	//  2. register api to server

	r := gin.Default()

	obj := Construct()

	v1 := r.Group("/cred")
	{
		v1.GET("/pay", obj.Paymentreq)
	}

	r.Run(":9090")

}
