package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type File struct {
	Groups []struct {
		Path string
		Aliases []string
	} `toml:"group"`
	Tasks []struct {
		Name string
		Aliases []string
		Groups []string
		Targets []string
		Commands []string
	} `toml:"task"`
}

func load(path string) (string, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(body), err
}

func parse(data string) (*File, error) {
	var file File
	_, err := toml.Decode(data, &file)
	return &file, err
}

func convert(file *File) (*Config, error) {
	return nil, nil
}

func GetConfigFromPath(path string) (*Config, error) {
	data, err := load(path)
	if err != nil {
		return nil, err
	}

	file, err := parse(data)
	if err != nil {
		return nil, err
	}

	return convert(file)
}