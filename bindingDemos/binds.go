package bindingdemos

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

func bindPerson(c *gin.Context) {
	var person Person
	if c.Bind(&person) == nil {
		log.Println("binding by Bind")
		log.Println("person.name=", person.Name)
		log.Println("person.address=", person.Address)
	}
	// similar to c.Bind
	if c.BindQuery(&person) == nil {
		log.Println("binding by BindQuery")
		log.Println("person.name=", person.Name)
		log.Println("person.address=", person.Address)
	}
	if c.BindJSON(&person) == nil {
		log.Println("binding by BindJSON")
		log.Println("person.name=", person.Name)
		log.Println("person.address=", person.Address)
	}
	c.String(http.StatusOK, "success")
}

// bind by query: http://127.0.0.1:8080/test?name=Jane&address=xian
// bing by json: curl -X GET localhost:8080/test --data '{"name":"Jane", "address":"xian"}' -H "Content-Type:application/json"
func TestBindingViaFormAndJson() {
	r := gin.Default()
	r.GET("/test", bindPerson)
	r.Run(":8080")
}
