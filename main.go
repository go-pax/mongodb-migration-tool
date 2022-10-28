package main

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"main/commands"
	"os"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer logger.Sync()
	log := logger.Sugar()

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		if !errors.Is(err, commands.ErrHelp) {
			log.Infow("startup", "ERROR", err.Error())
		}
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	return processCommands(args, log)
}

// processCommands handles the execution of the commands specified on
// the command line.
func processCommands(args Args, log *zap.SugaredLogger) error {
	if args == nil || len(args) < 3 {
		fmt.Println("migrate sets up mongodb")
		fmt.Println("provide a command, connection string, and migration folder.")
		fmt.Println("\tcommand is, up or down")
		fmt.Println("\tconnection string, is the working connection string to mongodb")
		fmt.Println("\tmigration folder, is a folder with your JSON migrations begins with file://")
		fmt.Println("migrate up mongodb://localhost:27017/my_db file://./my_migrations")
		return commands.ErrHelp
	}

	command := args.Num(0)
	connectionString := args.Num(1)
	migrationDir := args.Num(2)
	if err := commands.Migrate(log, command, connectionString, migrationDir); err != nil {
		return fmt.Errorf("migrating mongodb: %w", err)
	}
	return nil
}
