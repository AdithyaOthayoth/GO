package main //main tells go that main package is the main entry point
import (
	"example.com/first-app/db"
	"example.com/first-app/routes"
	"github.com/gin-gonic/gin"
)

// uint - non negative integer data type, int=0, string="", float64=0.0, bool= false, int32
func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost 8080
}
