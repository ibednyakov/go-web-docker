// main.go

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/astaxie/beego"
)

const (
	Users     = "Users"
	User      = "User"
	UserId    = "UserId"
	Operation = ":operation"
)

func main() {
	/* This would match routes like the following:
	   /sum/3/5
	   /product/6/23
	   ...
	*/
	beego.Router("/:operation/:num1:int/:num2:int", &mainController{})
	beego.Run()
}

type mainController struct {
	beego.Controller
}

type Person struct {
	Fn string
	Ln string
}
type ColorGroup struct {
	ID1    int
	ID2    int
	Name   string
	Colors []string
	P      Person `json:"Person"`
}

func (c *mainController) Get() {

	//Obtain the values of the route parameters defined in the route above
	operation := c.Ctx.Input.Param(Operation)
	num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
	num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))

	//Set the values for use in the template
	c.Data["operation"] = operation
	c.Data["num1"] = num1
	c.Data["num2"] = num2
	c.TplName = "result.html"

	// Perform the calculation depending on the 'operation' route parameter
	switch operation {
	case "sum":
		c.Data["result"] = add(num1, num2)
	case "product":
		c.Data["result"] = multiply(num1, num2)
	case "json":
		per := Person{Fn: "John",
			Ln: "Doe",
		}
		group := ColorGroup{
			ID1:    num1,
			ID2:    num2,
			Name:   "Reds",
			Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
			P:      per,
		}

		if b, err := json.Marshal(group); err != nil {
			fmt.Println("error:", err)
		} else {
			c.Data["json"] = &group
			c.ServeJSON()
			os.Stdout.Write(b)
		}
	default:
		c.TplName = "invalid-route.html"
	}
}

func (c *mainController) Put() {
	operation := c.Ctx.Input.Param(Operation)
	switch operation {
	case Users:
	default:
		fmt.Println(operation)
	}
}

func (c *mainController) Post() {
	operation := c.Ctx.Input.Param(Operation)
	switch operation {
	default:
		fmt.Println(operation)
	}
}

func add(n1, n2 int) int {
	return n1 + n2
}

func multiply(n1, n2 int) int {
	return n1 * n2
}
