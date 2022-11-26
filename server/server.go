package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"goteway/handlers"
	"strings"
)

type version struct {
	UseVersion bool   `yaml:"use"`
	Number     string `yaml:"num"`
}

type Server struct {
	Config struct {
		Host       string     `yaml:"Host"`
		Port       int        `yaml:"Port"`
		URIBase    string     `yaml:"Base"`
		APIVersion version    `yaml:"Version"`
		Endpoints  []endpoint `yaml:"Endpoints"`
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
	err := s.router.Run(fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) WithYamlConfig(config []byte) *Server {
	err := yaml.Unmarshal(config, &s.Config)
	if err != nil {
		log.Fatal(err)
	}
	s.stuffRouter()
	return s
}

func (s *Server) stuffRouter() {
	basePath := s.apiUri()
	for _, ep := range s.Config.Endpoints {
		for _, methodHandler := range ep.Methods {
			handler := handlers.HandlerFunctions[methodHandler.Handler]
			s.router.Handle(
				strings.ToUpper(methodHandler.Method), basePath+ep.Path, handler)
		}
	}
}
