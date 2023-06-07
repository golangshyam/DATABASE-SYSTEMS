package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

type Onlineshopping struct {
	Amazon string
	Huddie string
	Tshirt string
}

type Offlineshopping struct {
	gorm.Model
	Peterengland string
	Trends       string
	Zudio        []Maxshopping `json:"zudio" gorm:"foreignkey:Refer"`
	//`gorm:"foreignKey:UserRefer ;references:Company"`
}

type Maxshopping struct {
	gorm.Model
	Refer  uint
	Shirts string
	Jeans  string
}

type Dressing struct {
	Status string `json:"status"`
}

type Festival interface {
	DressReq(c *gin.Context)
}

type Shopping struct {
	db *gorm.DB
}

// 3. connect to db
// 4. create table

func Constuctor() Festival {

	//db, err := gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/shopping?charset=utf8&parseTime=True")
	db, err := gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/shopping?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")
	}

	db.AutoMigrate(&Offlineshopping{}, &Maxshopping{})

	return &Shopping{db}

}

func Inputvalidation(req *gin.Context) *Onlineshopping {

	var inputs Onlineshopping

	req.ShouldBind(&inputs)

	return &inputs

}

func APIvalidation(inputs *Onlineshopping, req *Shopping) bool {

	//        var responses TrainingResponse

	fmt.Println(inputs)

	req.db.Create(&Offlineshopping{Peterengland: "shirt", Trends: "belt", Zudio: []Maxshopping{{Shirts: "small size", Jeans: "28 cm"}, {Shirts: "slim fit", Jeans: "skinny fit"}}})

	return true

}

func (req *Shopping) DressReq(c *gin.Context) {

	// 5. Inside api
	// above 3 steps

	inputs := Inputvalidation(c)

	if inputs == nil {
		c.JSON(201, "invalid input")
		return
	}

	fmt.Println(inputs)

	resp := APIvalidation(inputs, req)

	if resp == false {
		c.JSON(201, "invalid inputs arguments ")
		return
	}

	var Response Dressing = Dressing{"crazt and trendy style"}

	c.AsciiJSON(200, Response)
}

func main() {

	//  1. server
	//  2. register api to server

	r := gin.Default()

	obj := Constuctor()

	v1 := r.Group("/purchase")
	{
		v1.GET("/debits", obj.DressReq)
	}

	r.Run(":9090")

}
