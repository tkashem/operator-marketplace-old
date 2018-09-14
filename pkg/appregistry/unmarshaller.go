package appregistry

import (
	yaml "gopkg.in/yaml.v2"
)

type Manifest struct {
	Publisher string `yaml:"publisher"`
	Data      Data   `yaml:"data"`
}

type Data struct {
	CRDs     string `yaml:"customResourceDefinitions"`
	CSVs     string `yaml:"clusterServiceVersions"`
	Packages string `yaml:"packages"`
}

type blobUnmarshaller interface {
	// Unmarshall unmarshals package blob into structured representations
	Unmarshal(in []byte) (*Manifest, error)
}

type blobUnmarshallerImpl struct{}

func (*blobUnmarshallerImpl) Unmarshal(in []byte) (*Manifest, error) {
	m := &Manifest{}
	if err := yaml.Unmarshal(in, m); err != nil {
		return nil, err
	}

	return m, nil
}
