package cmd

import (
	"fmt"
	"os"

	"github.com/lbajolet/qdpep8/cpu"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qdpep8cli",
	Short: "A quick-and-dirty implementation of a PEP/8 emulator",
	Args:  cobra.ExactArgs(1),
	RunE:  runCmd,
}

var inputFile *string
var outputFile *string
var simMode *bool
var traceMode *bool

func runCmd(cmd *cobra.Command, args []string) error {
	cpu := cpu.NewPep8Cpu()
	err := cpu.LoadFromFile(args[0])
	if err != nil {
		return fmt.Errorf("load error: %s", err)
	}

	if *inputFile != "" {
		in, err := os.Open(*inputFile)
		if err != nil {
			return fmt.Errorf("input file error: %s", err)
		}
		cpu.In = in
	}

	if *outputFile != "" {
		out, err := os.Create(*outputFile)
		if err != nil {
			return fmt.Errorf("output file error: %s", err)
		}
		cpu.Out = out
	}

	if *simMode {
		cpu.NoEOFChariStop = true
	}

	if *traceMode {
		cpu.Trace = true
	}

	return cpu.Run()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	inputFile = rootCmd.Flags().StringP("input", "i", "", "path to the input file for stdin")
	outputFile = rootCmd.Flags().StringP("output", "o", "", "path to the output file for stdout")
	simMode = rootCmd.Flags().BoolP("eof", "e", false, "run the tests as in simulator mode, i.e. on EOF return some \\x00 rather than immediately stopping")
	traceMode = rootCmd.Flags().BoolP("trace", "t", false, "print the state of the CPU after each cycle")
}
