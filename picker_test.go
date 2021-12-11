package picker

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {

	newPath, err := Generate(5, "testdata/my-test.log")
	require.NoError(t, err)

	t.Logf("Generate() newPath: %v", newPath)

	require.Equal(t, filepath.Base(newPath), "my-test-3.log")
}

func TestGenerateFail(t *testing.T) {
	_, err := Generate(2, "testdata/my-test.log")
	require.Error(t, err)

	_, err = Generate(0, "testdata/my-test.log")
	require.Error(t, err)
}
