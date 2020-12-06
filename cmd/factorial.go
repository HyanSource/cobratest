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

// factorialCmd represents the factorial command
var factorialCmd = &cobra.Command{
	Use:   "factorial",
	Short: "階層公式",
	Long: `階層公式 
	輸入1個數字並計算它的階層
	4 = 4*3*2*1 = 24
	5 = 5*4*3*2*1 = 120`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			fmt.Println("參數不對")
			return
		}

		t, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("未輸入正確數字")
			return
		}

		fmt.Println("計算結果為:", pc.F(t))
	},
}

func init() {
	rootCmd.AddCommand(factorialCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// factorialCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// factorialCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
