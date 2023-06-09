package kustz

import (
	"os"

	"github.com/knight-zlm/kustz/pkg/kubeutils"
)

type Config struct {
	Name      string  `json:"name"`
	Namespace string  `json:"namespace"`
	Service   Service `json:"service"`
}

type Service struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	Replicas int32  `json:"replicas"`
}

func NewKustzFromConfig(cfg string) *Config {
	b, err := os.ReadFile(cfg)
	if err != nil {
		panic(err)
	}

	kz := &Config{}
	err = kubeutils.YAMLUnMarshal(b, kz)
	if err != nil {
		panic(err)
	}

	return kz
}
