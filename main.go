package main

import (
	"database/sql"
	"gator/internal/config"
	"gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name string
	args []string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no command given.")
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)

	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	s := &state{
		db:     dbQueries,
		config: &cfg,
	}

	cmds := &commands{
		commands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)

	if err := cmds.run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
