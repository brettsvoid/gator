package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/brettsvoid/gator/internal/commands"
	"github.com/brettsvoid/gator/internal/config"
	"github.com/brettsvoid/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	state := &commands.State{
		DB:     dbQueries,
		Config: &cfg,
	}
	cmds := commands.Commands{
		Handlers: map[string]func(*commands.State, commands.Command) error{},
	}
	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegister)
	cmds.Register("reset", commands.HandlerReset)
	cmds.Register("users", commands.HandlerListUsers)
	cmds.Register("agg", commands.HandlerAgg)
	cmds.Register("addfeed", commands.HandlerAddFeed)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	err = cmds.Run(state, commands.Command{
		Name: cmdName,
		Args: cmdArgs,
	})
	if err != nil {
		log.Fatal(err)
	}
}
