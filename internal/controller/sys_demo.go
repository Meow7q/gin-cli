package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type DemoController struct {
	Api
}

func NewDemoController() *DemoController {
	return &DemoController{}
}

func (api *DemoController) HelloWorld(c *gin.Context) {
	str, ok := c.GetQuery("name")
	if !ok {
		str = "gin-layout"
	}

	api.Success(c, fmt.Sprintf("hello %s", str))
}
