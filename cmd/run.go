/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mount986/advent2023/advent"
	"github.com/spf13/cobra"
)

var days []int
var part int

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run a solution for the provided day/part",
	Long: `Flags:
	--day 	Comma Separated List of days to run
	--part	Which part to run for the provided day(s)

SubCommands:
	all		Run all parts of all days

Examples: 
	advent2023 run --day 1 --part 2
	advent2023 run --day 12
	advent2023 run all`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			advent advent.Advent
		)

		advent.Days = days
		advent.Part = part
		advent.Run()

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")
	runCmd.PersistentFlags().IntSliceVarP(&days, "day", "d", []int{}, "Comma Separated List of days to run")
	runCmd.PersistentFlags().IntVarP(&part, "part", "p", 0, "Which part to run")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
