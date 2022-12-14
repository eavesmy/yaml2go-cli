package example

type Default struct {
	Port int `yaml:"port"`
	Log  struct {
		Level string `yaml:"level"`
		Path  string `yaml:"path"`
	}
	Db struct {
		Path string `yaml:"path"`
		Name string `yaml:"name"`
	}
	Bind struct {
		Hosts []string `yaml:"hosts"`
	}
}
