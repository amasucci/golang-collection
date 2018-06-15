package main
import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "github.com/gin-gonic/gin"
)
type PASSWORD struct {
    HASH string `json:"hash" binding:"required"`
}
func main() {
    r := gin.Default()
    r.POST("/bcrypt", func(c *gin.Context) {
        var hash PASSWORD
        c.BindJSON(&hash)
        fmt.Printf("Value: %v\n", hash)
        hashed, _ := bcrypt.GenerateFromPassword([]byte(hash.HASH), 13)
        fmt.Printf("Hashed value: %v\n", string(hashed))
        c.String(201, string(hashed))
    })
    r.Run(":8080") // listen an
}
