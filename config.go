package main

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// Config represents a config file
type Config struct {
	Controller string
	Reader     string

	Twitch twitch
	IRC    irc
}

type twitch struct {
	Username string
	Password string
	Channel  string
}

type irc struct {
	Server  string
	Port    int
	Channel string
}

// ReadConfig reads the given config file
func ReadConfig(filename string) (c *Config, err error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c = &Config{}
	toml.Decode(string(contents), c)

	return c, nil
}
