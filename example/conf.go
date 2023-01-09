// Automatic generate. Do not modify this file.
package example

type Default struct {
	Log struct {
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
	Port int `yaml:"port"`
}
