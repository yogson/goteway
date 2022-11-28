package server


type methodHandler struct {
	Method  string `yaml:"method"`
	Handler string `yaml:"handler"`
	Middleware []string `yaml:"middleware"`
	Params map[string]any `yaml:"params"`
}

type endpoint struct {
	Path    string          `yaml:"path"`
	UseBase string			`yaml:"use-base,omitempty"`
	Methods []methodHandler	`yaml:"methods"`
}
