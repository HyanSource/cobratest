/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"strconv"

	"github.com/HyanSource/cobratest/pc"
	"github.com/spf13/cobra"
)

// permutationsCmd represents the permutations command
var permutationsCmd = &cobra.Command{
	Use:   "permutations",
	Short: "排列公式",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("參數不對")
			return
		}

		n, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("未輸入正確數字")
			return
		}
		k, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("未輸入正確數字")
			return
		}

		fmt.Println("計算結果為:", pc.P(n, k))
	},
}

func init() {
	rootCmd.AddCommand(permutationsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// permutationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// permutationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
