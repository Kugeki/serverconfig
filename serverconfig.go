package serverconfig

import (
	"errors"
	"io"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Addr                string `yaml:"Addr"`
	ReadTimeoutMs       uint64 `yaml:"ReadTimeoutMs,omitempty"`
	ReadHeaderTimeoutMs uint64 `yaml:"ReadHeaderTimeoutMs,omitempty"`
	WriteTimeoutMs      uint64 `yaml:"WriteTimeoutMs,omitempty"`
	IdleTimeoutMs       uint64 `yaml:"IdleTimeoutMs,omitempty"`
	MaxHeaderBytes      int    `yaml:"MaxHeaderBytes,omitempty"`
}

var (
	ErrAddrIsEmpty = errors.New("yaml: addr is empty")
)

func FromReader(r io.Reader) (*Config, error) {
	conf := Config{}

	err := yaml.NewDecoder(r).Decode(&conf)
	if err != nil {
		return nil, err
	}

	if conf.Addr == "" {
		return nil, ErrAddrIsEmpty
	}

	return &conf, nil
}
