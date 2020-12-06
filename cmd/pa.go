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

	"github.com/spf13/cobra"
)

// paCmd represents the pa command
var paCmd = &cobra.Command{
	Use:   "pa",
	Short: "把所有輸入的數字加總起來",
	Long:  `把所有輸入的數字加總起來 比如123 456 789 就會是1368`,
	Run: func(cmd *cobra.Command, args []string) {

		c := 0.0

		for _, v := range args {
			t, err := strconv.Atoi(v)
			if err != nil {
				panic("number error")
			}
			c += float64(t)
		}
		c -= m
		fmt.Println("pa called ", c)
	},
}

var (
	m float64
)

func init() {
	rootCmd.AddCommand(paCmd)

	paCmd.Flags().Float64VarP(&m, "minus", "m", 0, "需要減去的數字")
}
