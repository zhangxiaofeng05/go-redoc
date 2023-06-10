package main

import (
	"log"

	"github.com/gin-gonic/gin"
	go_redoc "github.com/zhangxiaofeng05/go-redoc"
	"github.com/zhangxiaofeng05/go-redoc/_example/gin/controller"
	"github.com/zhangxiaofeng05/go-redoc/_example/gin/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", c.ShowAccount)
			accounts.GET("", c.ListAccounts)
			//accounts.POST("", c.AddAccount)
			//accounts.DELETE(":id", c.DeleteAccount)
			//accounts.PATCH(":id", c.UpdateAccount)
			//accounts.POST(":id/images", c.UploadAccountImage)
		}
		//...
	}

	//docs.HandlerDocs(v1)
	gr := &go_redoc.Redoc{
		SwaggerJsonPath: "/swagger.json",
		JsonData:        docs.SwaggerInfo.ReadDoc(),
		RedocPath:       "/redoc",
		Title:           docs.SwaggerInfo.Title,
		Description:     docs.SwaggerInfo.Description,
	}
	go_redoc.GinHandlerDocs(v1, gr)

	log.Println("redoc: ", "http://127.0.0.1:8080/api/v1/redoc")
	r.Run(":8080")
}

//...
