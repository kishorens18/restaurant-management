package main

import (
	"os"
	"restaurant/database"
	"restaurant/middleware"
	"restaurant/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Authentication())

	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.TableRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}

/*

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
This line creates a variable called foodCollection and initializes it with a pointer to the food collection in the MongoDB database.

The var keyword is used to declare a variable. The variable name is foodCollection. The type of the variable is a pointer to a mongo.Collection. The * symbol indicates that the variable is a pointer.

The database.OpenCollection() function is used to open a collection in the MongoDB database. The first parameter to the function is the database.Client variable, which is a connection to the MongoDB database. The second parameter is the name of the collection to open.

The * symbol in front of the mongo.Collection type indicates that the variable is a pointer. A pointer is a variable that stores the address of another variable. In this case, the foodCollection variable stores the address of the food collection in the MongoDB database.

*/

/*

router := gin.New()
router.Use(gin.Logger())
router.Use(middleware.Authentication())
This line creates a new gin router and registers two middlewares: gin.Logger() and middleware.Authentication().

The gin.Logger() middleware logs all HTTP requests and responses. This can be helpful for debugging and monitoring the application.
The middleware.Authentication() middleware ensures that all requests are authenticated. This can be used to protect the application from unauthorized access.
Middleware is a Go concept that allows you to insert custom code into the request handling pipeline. This can be used to perform tasks such as logging, authentication, and rate limiting.



*/
