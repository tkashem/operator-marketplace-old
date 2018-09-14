package appregistry

import (
	"errors"
	"fmt"
	"strings"
)

// OperatorMetadata encapsulates operator metadata and manifest assocated with a package
type OperatorMetadata struct {
	Namespace  string
	Repository string
	Release    string
	Digest     string
	Manifest   *Manifest
}

func (om *OperatorMetadata) ID() string {
	return fmt.Sprintf("%s/%s", om.Namespace, om.Repository)
}

type client struct {
	adapter      apprApiAdapter
	decoder      blobDecoder
	unmarshaller blobUnmarshaller
}

func (c *client) RetrieveAll() ([]*OperatorMetadata, error) {
	packages, err := c.adapter.ListPackages()
	if err != nil {
		return nil, err
	}

	list := make([]*OperatorMetadata, len(packages))
	for i, pkg := range packages {
		manifest, err := c.RetrieveOne(pkg.Name, pkg.Default)
		if err != nil {
			return nil, err
		}

		list[i] = manifest
	}

	return list, nil
}

func (c *client) RetrieveOne(name, release string) (*OperatorMetadata, error) {
	namespace, repository, err := split(name)
	if err != nil {
		return nil, err
	}

	metadata, err := c.adapter.GetPackageMetadata(namespace, repository, release)
	if err != nil {
		return nil, err
	}

	digest := metadata.Content.Digest
	blob, err := c.adapter.DownloadOperatorManifest(namespace, repository, digest)
	if err != nil {
		return nil, err
	}

	decoded, err := c.decoder.Decode(blob)
	if err != nil {
		return nil, err
	}

	manifest, err := c.unmarshaller.Unmarshal(decoded)
	if err != nil {
		return nil, err
	}

	om := &OperatorMetadata{
		Namespace:  namespace,
		Repository: repository,
		Release:    release,
		Manifest:   manifest,
		Digest:     digest,
	}

	return om, nil
}

func split(name string) (namespace string, repository string, err error) {
	// we expect package name to comply to this format - {namespace}/{repository}
	split := strings.Split(name, "/")
	if len(split) != 2 {
		return "", "", errors.New(fmt.Sprintf("package name should be specified in this format {namespace}/{repository}"))
	}

	namespace = split[0]
	repository = split[1]

	return namespace, repository, nil
}
