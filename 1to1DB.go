package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

type TrainingRequest struct {
	Company      string
	Technology   string
	Participants string
}

type Cliento struct {
	gorm.Model
	Company   string
	Trainings Trainingo `json:"trainings" gorm:"foreignkey:Refer"`
	//`gorm:"foreignKey:UserRefer ;references:Company"`
}

type Trainingo struct {
	gorm.Model
	Refer        uint
	Technology   string
	Participants string
}

type TrainingResponse struct {
	Status string `json:"email"`
}

type Technology interface {
	TrainingReq(c *gin.Context)
}

type Trainings struct {
	db *gorm.DB
}

// 3. connect to db
// 4. create table

func Constuctor() Technology {

	db, err := gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/1to1db?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")
	}

	db.AutoMigrate(&Cliento{}, &Trainingo{})

	return &Trainings{db}

}

func Inputvalidation(req *gin.Context) *TrainingRequest {

	var inputs TrainingRequest

	req.ShouldBind(&inputs)

	/*

		if len(inputs.Passsword) < 9 {
				    req.JSON(201,"invalid password")
				    return nil
		}

	*/

	return &inputs

}

func APIvalidation(inputs *TrainingRequest, req *Trainings) bool {

	//        var responses TrainingResponse

	fmt.Println(inputs)

	req.db.Create(&Cliento{Company: "jvt technologies", Trainings: Trainingo{Technology: "golang", Participants: "freshers"}})

	//       req.db.Where("email = ? AND passsword = ?",inputs.Email ,inputs.Passsword).Find(&responses)

	//     db.Where("email = ?", inputs.Email).First(&responses)

	/*
	          if   responses.Name == "" {
	   	 return false
	          }
	*/

	return true

}

func (req *Trainings) TrainingReq(c *gin.Context) {

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
		c.JSON(201, "invalid email or password ")
		return
	}

	var Response TrainingResponse = TrainingResponse{"training  registred"}

	c.AsciiJSON(200, Response)
}

func main() {

	//  1. server
	//  2. register api to server

	r := gin.Default()

	obj := Constuctor()

	v1 := r.Group("/ttd")
	{
		v1.GET("/login", obj.TrainingReq)
	}

	r.Run(":9090")

}
