package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/config"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/storage/conn"
	"github.com/xxxVitoxxx/JiaPen-backend/pkg/storage/mysqlDB"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "build tables of mysql",
	Long:  "build tables of mysql",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.LoadRun()
		if err != nil {
			panic(err)
		}

		db := conn.CheckConnect()

		if err := db.AutoMigrate(&mysqlDB.Message{}); err != nil {
			panic(err)
		}
	},
}
