package router

import (
	"net/http"

	docs "github.com/centraldigital/cfw-cms-bff/docs"
	"github.com/centraldigital/cfw-cms-bff/internal/handler"
	"github.com/centraldigital/cfw-core-lib/pkg/util/ginutil"

	"github.com/centraldigital/cfw-cms-bff/property"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const serviceBaseURL = "/api"

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-KEY
// @description API key

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description JWT
func SetupRouter(r *gin.Engine, h handler.Handler) {
	docs.SwaggerInfo.Title = property.Get().Server.ServiceName
	docs.SwaggerInfo.Description = property.Get().Server.ServiceDescription
	docs.SwaggerInfo.Version = property.Get().Server.ApiDocsVersion
	docs.SwaggerInfo.Host = property.Get().Server.ApiDocsHost
	docs.SwaggerInfo.Schemes = []string{property.Get().Server.ApiDocsSchema}
	docs.SwaggerInfo.BasePath = serviceBaseURL

	// Add API doc endpoint to router
	if property.Get().Server.ApiDocs {
		docsPath := "/docs" + serviceBaseURL
		r.GET(docsPath, func(ctx *gin.Context) { ctx.Redirect(http.StatusTemporaryRedirect, docsPath+"/swagger/index.html") })
		r.GET(docsPath+"/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	basePath := r.Group(serviceBaseURL)
	baseRouterV1 := basePath.Group("/v1")

	customer := baseRouterV1.Group("/customer")
	{
		customer.GET("/:customer-id", func(c *gin.Context) {
			h.HandlerCustomerIdGet(c, c.Param("customer-id"))
		})

		customer.GET("", h.HandlerCustomerGetAll)

		customer.POST("", h.HandlerCustomerPost)

		customer.PUT("/:customer-id", func(c *gin.Context) {
			h.HandlerCustomerPut(c, c.Param("customer-id"))
		})

		customer.DELETE("/:customer-id", func(c *gin.Context) {
			h.HandlerCustomerDelete(c, c.Param("customer-id"))
		})
	}

	dog := baseRouterV1.Group("/dog")
	{
		dog.POST("/:breed", func(c *gin.Context) {
			ginutil.BindReqJson204Resp(c, h.HandlerDogImagePost)
		})

	}

}
