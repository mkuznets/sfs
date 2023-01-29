package yfs

import (
	"github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog/log"
	"mkuznets.com/go/sfs/ytils/y"
	"mkuznets.com/go/sfs/ytils/yerr"
	"os"
	"path/filepath"
)

func EnsureDir(parts ...string) (string, error) {
	path := filepath.Join(parts...)

	expandedPath, err := homedir.Expand(path)
	if err != nil {
		return "", yerr.New("failed to expand path %s", path).Err(err)
	}

	absPath, err := filepath.Abs(expandedPath)
	if err != nil {
		return "", yerr.New("failed to get absolute path for %s", absPath).Err(err)
	}

	if err := os.MkdirAll(absPath, 0755); err != nil {
		return "", yerr.New("failed to create target directory: %s", absPath).Err(err)
	}

	if err := IsWritableDir(absPath); err != nil {
		return "", err
	}

	return absPath, nil
}

func IsWritableDir(path string) (err error) {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return yerr.New("path does not exist: %s", path)
		}
		return yerr.New("could not open dir: %s", path).Err(err)
	}

	if !fi.IsDir() {
		return yerr.New("path is not a directory: %s", path)
	}

	f, err := os.CreateTemp(path, ".tmp*")
	if err != nil {
		return yerr.New("path is not writable: %s", path).Err(err)
	}
	defer func() {
		if err := os.Remove(f.Name()); err != nil {
			log.Debug().Stack().Err(yerr.New("Remove error").Err(err)).Send()
		}
		y.Close(f)
	}()

	return nil
}
