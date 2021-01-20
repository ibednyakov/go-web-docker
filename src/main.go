// main.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"database/sql"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

const (
	Users     = "Users"
	User      = "User"
	UserId    = "UserId"
	Operation = ":operation"
	ConnStr   = "user=dbUser password='mysecretpassword' host='db' sslmode=disable port=5432 connect_timeout=3 dbname=InfoTestDb"
)

func main() {

	beego.Router("/:operation/:tblname/:num1:int/:num2:int", &mainController{})
	beego.Run()
}

type mainController struct {
	beego.Controller
}

func (c *mainController) Get() {

	//Obtain the values of the route parameters defined in the route above
	operation := c.Ctx.Input.Param(Operation)
	table := c.Ctx.Input.Param(":tblname")
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
		db, err := sql.Open("postgres", ConnStr)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		defer db.Close()
		// orderInfo := OrderWithId{id: num1}
		// dbErr := db.QueryRow("SELECT * FROM $1 WHERE id = $2", table, num1).Scan(&orderInfo.id, &orderInfo.order)

		// if dbErr != nil {
		// 	log.Fatal(dbErr)
		// }
		// if b, err := json.Marshal(orderInfo); err != nil {
		// 	fmt.Println("error:", err)
		// } else {
		// 	c.Data["json"] = &orderInfo
		// 	c.ServeJSON()
		// 	os.Stdout.Write(b)
		// }
		// per := Person{Fn: "John",
		// 	Ln: "Doe",
		// }
		// group := ColorGroup{
		// 	ID1:    num1,
		// 	ID2:    num2,
		// 	Name:   "Reds",
		// 	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		// 	P:      per,
		// }

		log.Printf("Tablename: " + table + " id=" + string(num1))
		groupId := ColorGroupWithId{id: num1}
		rows, dbErr := db.Query("SELECT info FROM tmp_data WHERE id = $1", num1)
		if dbErr != nil {
			log.Fatal(dbErr)
		}
		defer rows.Close()

		data := make(ColorGroupMap)
		for rows.Next() {
			//err = rows.Scan(&groupId.group)
			err = rows.Scan(&data)
			if err != nil {
				panic(fmt.Sprintf("rows.Scan: %v", err))
			}
			fmt.Println("Data:")
			fmt.Println(data)

			fmt.Println("Assigning:")
			person := data["Person"].(map[string]interface{})
			groupId.group.P.Fn = person["fn"].(string)
			groupId.group.P.Ln = person["ln"].(string)

			groupId.group.ID1 = int(data["ID1"].(float64))
			groupId.group.ID2 = int(data["ID2"].(float64))
			if val, ok := data["Name"]; ok {
				groupId.group.Name = val.(string)
			}
			fmt.Println("done:")
		}

		fmt.Println("DONE:")
		//rows.Scan(&groupId.group)
		// for rows.Next() {
		//
		// 	log.Println(row)
		// 	var i interface{}
		// 	err = json.Unmarshal(row, &i)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	groupId.group = i.(map[string]interface{})
		// }

		if b, err := json.Marshal(groupId.group); err != nil {
			fmt.Println("error:", err)
		} else {
			c.Data["json"] = &groupId.group
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
