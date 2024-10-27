// Copyright © 2023 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"portfolio/app/auth"

	"github.com/spf13/cobra"
)

// authCmd represents the api command
var authCmd = &cobra.Command{
	Use:   "auth API",
	Short: "Run portfolio auth app on ...",
	Long:  `Using api command to starts portfolio auth application`,
	Run: func(cmd *cobra.Command, args []string) {
		auth.Run()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
