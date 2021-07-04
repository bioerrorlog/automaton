/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// gen_2d()
		gen_1d()
	},
}

func gen_1d() {
	grid := initialize_1d(5, 5)
	fmt.Println(grid)

	count := 10
	for i := 0; i < count; i++ {
		grid = calc_1d(grid)
		fmt.Println(grid)
	}
}

func calc_1d(grid [][]int) [][]int {
	return grid
}

func initialize_1d(x int, y int) [][]int {
	grid := [][]int{}
	row := []int{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			row = append(row, rand.Intn(2))
		}
		grid = append(grid, row)
		row = nil
	}

	return grid
}

func gen_2d() {
	grid := initialize_2d(5, 5)
	fmt.Println(grid)

	count := 10
	for i := 0; i < count; i++ {
		grid = calc_2d(grid)
		fmt.Println(grid)
	}
}

func calc_2d(grid [][]int) [][]int {
	return grid
}

func initialize_2d(x int, y int) [][]int {
	grid := [][]int{}
	row := []int{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			row = append(row, rand.Intn(2))
		}
		grid = append(grid, row)
		row = nil
	}

	return grid
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
