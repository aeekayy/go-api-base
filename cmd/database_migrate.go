package cmd

/*
Copyright © 2021 Farye Nwede <farye@aeekay.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/aeekayy/go-api-base/pkg/database"
)

// databaseMigrateCmd represents the migrate command
var databaseMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Starting database migration")

		dbConfig := database.NewConfig()

		// setup the database pool
		importDbConfig := viper.GetStringMap("db")
		mapstructure.Decode(&importDbConfig, &dbConfig)
		err := database.MigrateDatabase(nil, dbConfig)

		if err != nil {
			log.Errorf("Migration error: %s", err)
		}
	},
}

func init() {
	databaseCmd.AddCommand(databaseMigrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// databaseMigrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// databaseMigrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
