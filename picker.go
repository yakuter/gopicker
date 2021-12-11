package picker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// setZipFilePath generates and returns unique zip file name with full path.
func Generate(tryCount uint, path string) (string, error) {

	absPath := tryAbsPath(path)

	if _, err := os.Lstat(absPath); os.IsNotExist(err) {
		return absPath, nil
	}

	if tryCount == 0 {
		return "", fmt.Errorf("file '%s' already exist ", path)
	}

	dir := filepath.Dir(absPath)
	filename := filepath.Base(absPath)
	ext := filepath.Ext(filename)
	base := trimExt(filename)

	for i := 1; i < int(tryCount); i++ {
		newPath := filepath.Join(dir, fmt.Sprintf("%s-%d%s", base, i, ext))
		if _, err := os.Lstat(newPath); os.IsNotExist(err) {
			return newPath, nil
		}
	}

	return "", fmt.Errorf("unable to pick a unique file name for %s in %d tries", path, tryCount)
}

func trimExt(filename string) string {
	newFilename := strings.TrimSuffix(filename, filepath.Ext(filename))

	// This is necessary for files like ".bashrc"
	if newFilename == "" {
		newFilename = filename
	}

	return newFilename
}

func tryAbsPath(path string) string {
	if !filepath.IsAbs(path) {
		p, err := filepath.Abs(path)
		if err == nil {
			return p
		}
	}
	return path
}
