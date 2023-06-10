## go-redoc
redoc for go

## Usage
 - `gin`
```
import (
    "github.com/gin-gonic/gin"
    go_redoc "github.com/zhangxiaofeng05/go-redoc"
)

...

r := gin.Default()
v1 := r.Group("/api/v1")

gr := &go_redoc.Redoc{
	SwaggerJsonPath: "/swagger.json",
	JsonData:        docs.SwaggerInfo.ReadDoc(),
	RedocPath:       "/redoc",
	Title:           docs.SwaggerInfo.Title,
	Description:     docs.SwaggerInfo.Description,
}
go_redoc.GinHandlerDocs(v1, gr)
```
http://127.0.0.1:8080/api/v1/redoc