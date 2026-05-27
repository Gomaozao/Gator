package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"Gator/internal/config"
)

type state struct {
	config_ptr *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandName map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdName, ok := c.commandName[cmd.name]
	if ok {
		return cmdName(s, cmd)
	}
	return errors.New("Command not found.")
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandName[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("No Username provided.")
	}
	if err := s.config_ptr.SetUser(cmd.args[0]); err != nil {
		return errors.New("Something went wrong with the name.")
	}
	fmt.Println("Name has been set.")
	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	foo := state{&cfg}
	bar := commands{commandName: make(map[string]func(*state, command) error)}
	bar.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("error")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	args := os.Args[2:]
	usrcmd := command{cmdName, args}
	if err := bar.run(&foo, usrcmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
