package conf

import (
	"flag"
	"net/url"
)

type Config struct {
	Port string
	Name string
	Url  string
}

func ReadConfig() (Config, error) {
	var port = flag.String("port", "80", "port number")
	var name = flag.String("name", "user", "username")
	var urr = flag.String("url", "http://localhost", "dest address")
	var err = error(nil)
	flag.Parse()

	u, err := url.ParseRequestURI(*urr)
	if err != nil {
		panic(err)
	}
	c := Config{
		Port: *port,
		Name: *name,
		Url:  u.String(),
	}

	return c, err
}
