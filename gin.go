package go_redoc

import (
	"bytes"
	_ "embed"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

type Redoc struct {
	SwaggerJsonPath string `example:"/swagger.json"` // 设置 swagger.json 路由路径
	JsonData        string // swag 生成的 redoc.go 文件中 SwaggerInfo.ReadDoc()
	RedocPath       string `example:"/redoc"` // 设置 redoc 文档的路由路径
	Title           string // redoc 文档标题
	Description     string // redoc 文档描述
}

func GinHandlerDocs(g *gin.RouterGroup, r *Redoc) {
	g.GET(r.SwaggerJsonPath, func(c *gin.Context) {
		c.Writer.Header().Set("content-type", "application/json")
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.WriteString(r.JsonData)
	})
	g.GET(r.RedocPath, func(c *gin.Context) {
		p := g.BasePath() + r.SwaggerJsonPath
		buf := bytes.NewBuffer(nil)
		tmpl := template.Must(template.New("redoc").Parse(htmlData))
		if err := tmpl.Execute(buf, map[string]string{
			"body":        javaScriptData,
			"title":       r.Title,
			"description": r.Description,
			"url":         p,
		}); err != nil {
			panic(err)
		}

		c.Writer.Header().Set("Content-Type", "text/html")
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write(buf.Bytes())
	})
}
