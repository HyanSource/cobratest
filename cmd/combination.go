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

// combinationCmd represents the combination command
var combinationCmd = &cobra.Command{
	Use:   "combination",
	Short: "組合公式和重複組合公式",
	Long: `組合公式和重複組合公式
	輸入兩個數字 第一個為n 第二個為k 在n個物件中取出k個的組合數量
	當輸入-r為 n種物件均超過k個 而取出k個的組合數量
	`,
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

		if r {
			fmt.Println("計算結果為:", pc.H(n, k))
		} else {
			fmt.Println("計算結果為:", pc.C(n, k))
		}

	},
}

var (
	r bool
)

func init() {
	rootCmd.AddCommand(combinationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// combinationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// combinationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	combinationCmd.Flags().BoolVarP(&r, "repetition", "r", false, "重複組合")
}
