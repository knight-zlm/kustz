package kubeutils

import (
	"sigs.k8s.io/yaml"
)

func YAMLMarshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func YAMLUnMarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
