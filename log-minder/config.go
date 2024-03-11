package logminder

import (
	"fmt"
	"strings"
)

type config struct {
	addr     string
	username string
	password string
	debug    bool
}

func NewConfig(addr, username, password string, debug bool) (*config, error) {

	if addr == "" {
		return nil, ErrInvalidAddress
	}

	addr = strings.ReplaceAll(addr, "http://", "")
	addr = strings.ReplaceAll(addr, "https://", "")

	return &config{
		addr:     fmt.Sprintf("http://%s", addr),
		username: username,
		password: password,
		debug:    debug,
	}, nil
}
