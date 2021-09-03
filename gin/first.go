package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
type story struct {
	book string `json:"book"`
	bookType string `json:"bookType"`
}
func writing(){


}

func main() {
r := gin.Default()
fmt.Println("gin")
r.GET("/api", writing)
	r.Run(":5000")

}
