package cmd

import (
	"log"

	"github.com/lbryio/chainquery/apiactions"
	"github.com/lbryio/chainquery/config"
	"github.com/lbryio/chainquery/db"
	"github.com/lbryio/chainquery/lbrycrd"
	"github.com/lbryio/chainquery/swagger/apiserver"
	"github.com/lbryio/chainquery/twilio"
	"github.com/spf13/cobra"
)

func init() {
	serveCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitSlack()
		twilio.InitTwilio()
		apiactions.AutoUpdateCommand = config.GetAutoUpdateCommand()
		//Main Chainquery DB connection
		dbInstance, err := db.Init(config.GetMySQLDSN(), config.GetDebugQueryMode())
		if err != nil {
			log.Panic(err)
		}
		defer db.CloseDB(dbInstance)

		lbrycrdClient := lbrycrd.Init()
		defer lbrycrdClient.Shutdown()

		swagger.InitApiServer(config.GetAPIHostAndPort())
	},
}
