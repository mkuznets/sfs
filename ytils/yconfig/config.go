package yconfig

import (
	"github.com/mitchellh/go-homedir"
	"mkuznets.com/go/sfs/ytils/yerr"
	"path"
	"strings"

	"github.com/rs/zerolog/log"

	"go.uber.org/config"
)

type Config interface {
	Validate() error
}

type Reader[T Config] struct {
	filename      string
	systemDirs    bool
	configDir     string
	defaultConfig string
}

func New[T Config](filename string) *Reader[T] {
	return &Reader[T]{
		filename: filename,
	}
}

func (r *Reader[T]) WithSystemDirs() *Reader[T] {
	r.systemDirs = true
	return r
}

func (r *Reader[T]) WithDefaults(defaults string) *Reader[T] {
	r.defaultConfig = defaults
	return r
}

func (r *Reader[T]) WithDir(dir string) *Reader[T] {
	r.configDir = dir
	return r
}

func (r *Reader[T]) Read() (T, error) {
	var cfg T

	copts, err := r.yamlOptions()
	if err != nil {
		return cfg, err
	}

	provider, err := config.NewYAML(copts...)
	if err != nil {
		return cfg, err
	}

	if err := provider.Get("").Populate(&cfg); err != nil {
		return cfg, err
	}

	if err := cfg.Validate(); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (r *Reader[T]) yamlOptions() ([]config.YAMLOption, error) {
	if r.systemDirs && r.configDir != "" {
		return nil, yerr.New("cannot use both WithSystemDirs and WithDir")
	}

	options := make([]config.YAMLOption, 0, 3)

	// Default config
	if r.defaultConfig != "" {
		options = append(options, config.Source(strings.NewReader(r.defaultConfig)))
	}

	switch {
	case r.systemDirs:
		if configPath, ok := searchConfig(r.filename); ok {
			log.Debug().Str("configDir", configPath).Msg("Config file")
			options = append(options, config.File(configPath))
		}
	case r.configDir != "":
		configPath, err := homedir.Expand(path.Join(r.configDir, r.filename))
		if err != nil {
			return nil, err
		}
		log.Debug().Str("configDir", configPath).Msg("Config file")
		options = append(options, config.File(configPath))
	default:
		return nil, yerr.New("no config directory specified")
	}

	return options, nil
}
