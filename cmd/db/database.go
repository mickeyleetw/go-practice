package db

import (
	"context"
	"os"

	"senao/pkg/core/database"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "resetdb",
		Short: "reset senao db",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
	VERSION string
)

func start() {
	log.Info().Msg("ResetDB Start")
	ctx := context.Background()

	// connect to db
	dsn := database.GetDSN()
	db, err := database.GetDB(dsn)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer db.Close()

	DB_SCHEMA := os.Getenv("DB_SCHEMA")

	resetDBHelper, err := database.NewResetDBHelper(DB_SCHEMA)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	if err := resetDBHelper.Reset(ctx, db.DB); err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg("ResetDB Done")
}
