/*
Copyright 2016 Skippbox, Ltd.

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

package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/bitnami-labs/kubewatch/config"
	"github.com/spf13/cobra"
)

// toastConfigCmd represents the toast subcommand
var toastConfigCmd = &cobra.Command{
	Use:   "toast",
	Short: "specific toast configuration",
	Long:  `specific toast configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.New()
		if err != nil {
			logrus.Fatal(err)
		}

		mode, err := cmd.Flags().GetString("mode")
		if err == nil {
			if len(mode) > 0 {
				conf.Handler.Toast.Mode = mode
			}
		} else {
			logrus.Fatal(err)
		}

		if err = conf.Write(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	toastConfigCmd.Flags().StringP("mode", "m", "", "Specify toast mode")
}
