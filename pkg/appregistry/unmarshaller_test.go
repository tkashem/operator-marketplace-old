package appregistry

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshall(t *testing.T) {
	// Do not use tabs for indentation
	data := `
publisher: redhat
data:
  customResourceDefinitions: "my crds"
  clusterServiceVersions: "my csvs"
  packages: "my packages"
`

	u := blobUnmarshallerImpl{}
	manifest, err := u.Unmarshal([]byte(data))

	require.NoError(t, err)

	assert.Equal(t, "redhat", manifest.Publisher)
	assert.Equal(t, "my crds", manifest.Data.CRDs)
	assert.Equal(t, "my csvs", manifest.Data.CSVs)
	assert.Equal(t, "my packages", manifest.Data.Packages)
}
