package web_project

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Server struct {
	Router      *gin.Engine
	BaseAddress string
	Database    *Database
	Config      *TomlConf
}

/*
	func (s *Server) checkCredentials(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if ok {
			ok = s.Database.CheckAccount(username, password)
			if ok {
				return
			}
		}
		c.Writer.Header().Set(
			"WWW-Authenticate",
			`Basic realm="restricted", charset="UTF-8"`)
	}
*/
func (s *Server) initGetMethods() {
	s.Router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	s.Router.GET("/main_page", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main_page.html", nil)
	})
}

func (s *Server) initPostMethods() {
	s.Router.POST("/", func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		var jsonData map[string]string
		c.Bind(jsonData)

		fmt.Println(jsonData, string(body), c.ContentType())
		c.Redirect(301, "http://localhost:8080/main_page")
	})
}

func InitServer() *Server {
	server := new(Server)
	server.Database = InitDb()
	server.Router = gin.New()
	server.Config = NewConfig()
	server.BaseAddress = server.Config.WebServer.Config.IP + ":" + server.Config.WebServer.Config.Port

	server.Router.Static("/assets", "front/assets/")
	server.Router.LoadHTMLGlob("front/*.html")
	server.initPostMethods()
	server.initGetMethods()
	server.listen()
	return server
}

func (s *Server) listen() {
	_ = s.Router.Run(s.BaseAddress)
}
