package routers

import (
	"github.com/gin-gonic/gin"
)

//CreateRouter defines the router and sets the URL mappings.
func CreateRouter(m Mapper) *gin.Engine {
	router := gin.Default()

	m.configureMappings(router)

	return router
}
