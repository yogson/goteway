package server

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	ginzap "github.com/gin-contrib/zap"

	"github.com/yogson/goteway/sdk"
	"github.com/yogson/goteway/utils"
)


type version struct {
	UseVersion bool   `yaml:"use"`
	Number     string `yaml:"num"`
}

type Server struct {
	Config struct {
		Host       	string     	`yaml:"Host"`
		Port       	int        	`yaml:"Port"`
		URIBase    	string     	`yaml:"Base"`
		APIVersion 	version    	`yaml:"Version"`
		Endpoints  	[]endpoint 	`yaml:"Endpoints"`
		Debug		bool		`yaml:"Debug"`
	}

	router *gin.Engine
}

func (s *Server) apiUri() string {
	uri := s.Config.URIBase
	if uri[len(uri)-1] != '/' {
		uri += "/"
	}
	if s.Config.APIVersion.UseVersion {
		uri += s.Config.APIVersion.Number + "/"
	}
	return uri
}

func New() *Server {
	server := Server{}
	server.router = gin.Default()
	return &server
}

func (s *Server) Run() {
	connStr := fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port)
	err := s.router.Run(connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) WithYamlConfig(config []byte) *Server {
	err := yaml.Unmarshal(config, &s.Config)
	if err != nil {
		log.Fatal("Unable to load config: ", err)
	}
	if !s.Config.Debug {
		s.router = s.getRouter()
	}
	s.stuffRouter()

	return s
}

func (s *Server) getRouter() *gin.Engine {
	router := gin.New()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		return gin.Default()
	}
	logger := utils.Logger
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	return router
}

func (s *Server) stuffRouter() {
	basePath := s.apiUri()
	for _, ep := range s.Config.Endpoints {
		base := basePath
		if strings.ToLower(ep.UseBase) == "false" {
			base = ""
		}
		for _, methodHandler := range ep.Methods {
			var handlers []gin.HandlerFunc
			for _, middlewareName := range methodHandler.Middleware {
				middleware := HandlerFunctions[middlewareName]
				if middleware == nil {continue}
				handlers = append(handlers, middleware.Get())
			}
			handler := HandlerFunctions[methodHandler.Handler]
			if handler == nil {continue}
			params := methodHandler.Params
			handlers = append(handlers, handler.Get(params))
			s.router.Handle(
				strings.ToUpper(methodHandler.Method), base+ep.Path, handlers...)
		}
	}
}

func (s *Server) getHandlersChain(handler sdk.IHandler, params map[string]any) []gin.HandlerFunc {
	var chain []gin.HandlerFunc
	return append(chain, handler.Get(params))
}