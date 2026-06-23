// CWE-502: Insecure YAML deserialization (go-yaml v2)
package main
import "gopkg.in/yaml.v2"
func loadConfig(data []byte) (map[string]interface{}, error) {
	var cfg map[string]interface{}
	err := yaml.Unmarshal(data, &cfg)
	return cfg, err
}
