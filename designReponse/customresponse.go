package designreponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomResponse struct {
	Data   string      `json:"data"`
	Id     int         `json:"id"`
	Detail interface{} `json:"detail"`
}

/*
func (c *CustomResponse) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"newData":   c.Data,
		"newId":     c.Id,
		"newDetail": c.Detail,
	}
	return json.Marshal(m)
}
func (c *CustomResponse) Error() string {
	return "Failed to convert CustomResponse to json"
}
*/

func myfunc(c *gin.Context) {
	r := CustomResponse{
		Data:   "my data",
		Id:     10,
		Detail: map[string]string{"a": "a content", "b": "b_content"},
	}
	c.JSON(http.StatusOK, &r)

}

// This demo shows how to response custom object in the form of json. I show 2 ways of using json parser
// 1. Using the default json parser(You need to comment out  MarshalJSON()  and Error())
// 2. Using custom json parser.

// http://127.0.0.1:8080/test
// if you comment out func (c *CustomResponse) MarshalJSON() ([]byte, error) and func (c *CustomResponse) Error() string {
// result is : {"data":"my data","id":10,"detail":{"a":"a content","b":"b_content"}}

// if you don't comment out func (c *CustomResponse) MarshalJSON() ([]byte, error) and func (c *CustomResponse) Error() string {
// result is: {"newData":"my data","newDetail":{"a":"a content","b":"b_content"},"newId":10}
func Demo() {
	r := gin.Default()
	r.GET("/test", myfunc)
	r.Run(":8080")
}
