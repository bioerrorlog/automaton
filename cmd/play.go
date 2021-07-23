package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	Mark string
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch Type {
		case "1d":
			play_1d(Mark)
		case "2d":
			play_2d(Mark)
			// default
			// TODO: error handling
		}
	},
}

func play_1d(mark string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(replace_1d(scanner.Text(), mark))
		// fmt.Printf("\r%s", replace_1d(scanner.Text(), mark))
		time.Sleep(time.Second / 30)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func replace_1d(str string, mark string) string {
	str = strings.Replace(str, "[", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "]", "", -1)
	str = strings.Replace(str, "0", " ", -1)
	str = strings.Replace(str, "1", mark, -1)
	return str
}

func play_2d(mark string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// fmt.Print("\033[?25l")     // Hide the cursor
		fmt.Print("\033[H\033[2J") // Clear terminal
		fmt.Print(replace_2d(scanner.Text(), mark))
		// fmt.Print("\033[?25h") // Show the cursor
		time.Sleep(time.Second / 30)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func replace_2d(str string, mark string) string {
	str = strings.Replace(str, "[", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "]", "\n", -1)
	str = strings.Replace(str, "0", " ", -1)
	str = strings.Replace(str, "1", mark, -1)
	return str
}

func init() {
	rootCmd.AddCommand(playCmd)

	rootCmd.PersistentFlags().StringVarP(&Mark, "mark", "m", "@", "String displayed as cells.")
}
