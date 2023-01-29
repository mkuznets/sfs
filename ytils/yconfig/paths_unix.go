package yconfig

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

func initBaseDirs(home string) {
	DirData = xdgPath("XDG_DATA_HOME", filepath.Join(home, ".local", "share"))
	DirConfig = xdgPath("XDG_CONFIG_HOME", filepath.Join(home, ".config"))
	DirCache = xdgPath("XDG_CACHE_HOME", filepath.Join(home, ".cache"))
}

func xdgPath(name, defaultPath string) string {
	dir, err := homedir.Expand(os.Getenv(name))
	if err == nil && dir != "" && filepath.IsAbs(dir) {
		return dir
	}
	return defaultPath
}
