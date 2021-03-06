package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var (
	Width int
	// Height int
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
			Gen2d()
			// default
			// TODO: error handling
		}
	},
}

func Gen1d() {
	line := Init1d(Width)
	fmt.Println(line)

	count := 100000
	for i := 0; i < count; i++ {
		line = Calc1d(line)
		fmt.Println(line)
	}
}

func Calc1d(line []int) []int {
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

func Gen2d() {
	grid := Init2d(Width, 50)
	fmt.Println(grid)

	count := 100000
	for i := 0; i < count; i++ {
		grid = Calc2d(grid)
		fmt.Println(grid)
	}
}

func Calc2d(grid [][]int) [][]int {
	rGrid := [][]int{}
	for y := 0; y < len(grid); y++ {
		row := grid[y]
		rRow := []int{}
		for x := 0; x < len(row); x++ {
			// c[0] c[1] c[2]
			// c[3]   m  c[4]
			// c[5] c[6] c[7]
			abo := y - 1
			bel := y + 1
			lef := x - 1
			rig := x + 1
			if abo < 0 {
				abo = len(grid) - 1
			}
			if bel >= len(grid) {
				bel = 0
			}
			if lef < 0 {
				lef = len(row) - 1
			}
			if rig >= len(row) {
				rig = 0
			}

			c := []int{}
			c = append(c, grid[abo][lef])
			c = append(c, grid[abo][x])
			c = append(c, grid[abo][rig])
			c = append(c, grid[y][lef])
			c = append(c, grid[y][rig])
			c = append(c, grid[bel][lef])
			c = append(c, grid[bel][x])
			c = append(c, grid[bel][rig])

			m := grid[y][x]

			s := 0
			for i := 0; i < len(c); i++ {
				s += c[i]
			}

			r := m
			if m == 1 && s > 3 {
				r = 0
			} else if m == 1 && (s == 2 || s == 3) {
				r = 1
			} else if m == 1 && s < 2 {
				r = 0
			} else if m == 0 && s == 3 {
				r = 1
			}
			rRow = append(rRow, r)
		}
		rGrid = append(rGrid, rRow)
	}
	return rGrid
}

func Init2d(x int, y int) [][]int {
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

	genCmd.PersistentFlags().IntVarP(&Width, "width", "w", 200, "Cell automaton width")
	// rootCmd.PersistentFlags().IntVarP(&Height, "height", "h", 50, "Cell automaton height. If -t 1d, this value ignored.")
}
