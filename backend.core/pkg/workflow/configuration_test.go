package workflow

import "testing"
import "github.com/go-yaml/yaml"
import _ "embed"

//go:embed workflows.yaml
var b []byte

func TestParser(t *testing.T) {
	configuration := &Configuration{}
	err := yaml.Unmarshal(b, &configuration)
	if err != nil {
		panic(err)
	}
	err = configuration.Validate()
	if err != nil {
		panic(err)
	}
	print(configuration)
}
