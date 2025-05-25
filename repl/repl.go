package repl

import (
	"fmt"
	"bufio"
	"os"

	"github.com/morganb27/pokedexcli/commands"
	"github.com/morganb27/pokedexcli/utils"
)

func StartREPL(cfg *commands.Config) {
	reader := bufio.NewScanner(os.Stdin)
	commands := commands.GetCommands()
	for {
		fmt.Print("Pokedex > ")
		reader.Scan() 
		input := utils.CleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		cmd, ok := commands[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.Callback(cfg); err != nil {
			fmt.Printf("Error with command %s: %v", commandName, err)
		}
	}
}