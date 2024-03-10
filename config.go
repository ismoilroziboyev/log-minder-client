package main

type config struct {
	addr     string
	username string
	password string
	debug    bool
}

func NewConfig(addr, username, password string, debug bool) *config {
	return &config{
		addr:     addr,
		username: username,
		password: password,
		debug:    debug,
	}
}
