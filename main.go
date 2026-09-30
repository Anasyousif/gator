package main

import (
	"fmt"
	"log"
	"os"

	"gator/internal/config"
)

// state holds application state, including a pointer to the config.
type state struct {
	cfg *config.Config
}

// command represents a CLI command and its arguments.
type command struct {
	Name string
	Args []string
}

// commands holds registered CLI handlers mapped by command name.
type commands struct {
	handlers map[string]func(*state, command) error
}

// register adds a new command handler to the map.
func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

// run executes a command if it is registered in the map.
func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.handlers[cmd.Name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}

// handlerLogin sets the current user in the config file.
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("the login handler expects a single argument: <username>")
	}

	username := cmd.Args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}

	fmt.Printf("User has been set to: %s\n", username)
	return nil
}

func main() {
	// 1. Read config and store in state
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	appState := &state{
		cfg: &cfg,
	}

	// 2. Initialize commands struct and register handlers
	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	// 3. Parse command-line arguments
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error: not enough arguments provided")
		os.Exit(1)
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	cmd := command{
		Name: cmdName,
		Args: cmdArgs,
	}

	// 4. Run command
	err = cmds.run(appState, cmd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}