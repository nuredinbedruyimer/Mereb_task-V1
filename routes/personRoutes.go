package routes

import (
	"Mereb-V1/controllers"

	"github.com/gin-gonic/gin"
)

func PersonRoutes(personRoutes *gin.Engine) *gin.Engine {

	// Person routes
	personRoutes.POST("/persons", controllers.CreatePersonController)
	personRoutes.GET("/persons", controllers.GetAllPersonController)
	personRoutes.GET("/persons/:id", controllers.GetPersonController)
	personRoutes.DELETE("/persons/:id", controllers.DeletePersonController)
	personRoutes.PUT("/persons/:id", controllers.UpdatePersonController)

	// Handle 404(Not Found)
	personRoutes.NoRoute(controllers.RouteNotFoundController)

	return personRoutes
}
