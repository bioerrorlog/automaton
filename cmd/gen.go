/*
Copyright_index © 2021 NAME HERE <EMAIL ADDRESS>

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
		switch Type {
		case "1d":
			Gen1d()
		case "2d":
			gen_2d()
			// default
			// TODO: error handling
		}
	},
}

func Gen1d() {
	line := Init1d(100)
	fmt.Println(line)

	count := 100000
	for i := 0; i < count; i++ {
		line = calc_1d(line)
		fmt.Println(line)
	}
}

func calc_1d(line []int) []int {
	result := []int{}
	var left_index, center_index, right_index, pattern int

	for i := range line {
		if i == 0 {
			left_index = len(line) - 1
		} else {
			left_index = i - 1
		}
		center_index = i
		if i == len(line)-1 {
			right_index = 0
		} else {
			right_index = i + 1
		}
		pattern = line[left_index]*100 + line[center_index]*10 + line[right_index]

		// Rule 30
		switch pattern {
		case 0:
			result = append(result, 0)
		case 1:
			result = append(result, 1)
		case 10:
			result = append(result, 1)
		case 11:
			result = append(result, 1)
		case 100:
			result = append(result, 1)
		case 101:
			result = append(result, 0)
		case 110:
			result = append(result, 0)
		case 111:
			result = append(result, 0)
			// default
			// TODO: error handling
		}
	}

	return result
}

func Init1d(x int) []int {
	line := []int{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < x; i++ {
		line = append(line, rand.Intn(2))
	}

	return line
}

func gen_2d() {
	grid := initialize_2d(50, 20)
	fmt.Println(grid)

	count := 100000
	for i := 0; i < count; i++ {
		grid = calc_2d(grid)
		fmt.Println(grid)
	}
}

func calc_2d(grid [][]int) [][]int {
	for y := 0; y < len(grid); y++ {
		row := grid[y]
		for x := 0; x < len(row); x++ {
			ty := y - 1
			if ty < 0 {
				ty = len(grid) - 1
			}
			tx := x - 1
			if tx < 0 {
				tx = len(row) - 1
			}
			if grid[ty][x] == 1 && grid[y][tx] == 0 {
				grid[y][x] = 1
			} else {
				grid[y][x] = 0
			}
		}
	}
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
}
