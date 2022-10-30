package cmd

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	CONFIG_DIR string
	LOG_DIR    string

	LogFlag bool
)

var rootCmd = &cobra.Command{
	Use:   "csview",
	Short: "Open and edit csv files vim-style",
	Run:   func(cmd *cobra.Command, args []string) {},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		InitializeLogger()
		log.Info().Interface("args", os.Args).Msg("Execute started")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		DestroyLogger()
	},
}

var logfile *os.File

func InitializeLogger() {
	logfile, err := os.OpenFile(path.Join(CONFIG_DIR, "log"),
		os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in creating log file: %v\n", err)
	}

	var writer io.Writer = logfile

	if LogFlag {
		writer = io.MultiWriter(writer, os.Stderr)
	}

	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        writer,
		TimeFormat: time.RFC3339,
	}).
		With().Timestamp().Caller().Logger()

	log.Logger = logger
}

func DestroyLogger() {
	logfile.Close()
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// global vars
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in initializing: %v\n", err)
	}

	CONFIG_DIR = path.Join(homedir, ".config/csview")

	err = os.MkdirAll(CONFIG_DIR, 0775)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in creating config dir: %v\n", err)
	}

	// flags
	rootCmd.PersistentFlags().BoolVar(&LogFlag, "log", false, "turn on logging to stderr")
}
