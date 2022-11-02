package server

type methodHandler struct {
	Method  string `yaml:"method"`
	Handler string `yaml:"handler"`
}

type endpoint struct {
	Path    string          `yaml:"path"`
	Methods []methodHandler `yaml:"methods"`
}
