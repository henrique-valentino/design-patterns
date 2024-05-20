package main

import "fmt"

type Server struct {
	config Config
}

func NewServer(config Config) (*Server, error) {
	return &Server{config: config}, nil
}

type Config struct {
	listenAdd string
	id        string
	name      string
}

func (c Config) WithListenAddr(addr string) Config {
	c.listenAdd = addr
	return c
}

func (c Config) WithID(id string) Config {
	c.id = id
	return c
}

func (c Config) WithName(name string) Config {
	c.name = name
	return c
}

func NewConfig() Config {
	return Config{
		listenAdd: ":3000",
		id:        "randon_id",
		name:      "rando_name",
	}
}

func main() {
	config := NewConfig().
		WithID("foor").
		WithListenAddr(":555").
		WithName("bar")

	Server, _ := NewServer(config)
	fmt.Println(Server)
}
