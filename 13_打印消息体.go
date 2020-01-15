package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("object/BUSINESS/instance/_search", events)
	router.Run(":8089")
}

func events(c *gin.Context) {
	buf := make([]byte, 102400000)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println("消息体：", string(buf[0:n]))
	for k, v := range c.Request.Header {
		fmt.Println(k, v)
	}
	resp := map[string]string{"hello": "world"}
	c.JSON(http.StatusOK, resp)
	/*post_gwid := c.PostForm("name")
	fmt.Println(post_gwid)*/
}
