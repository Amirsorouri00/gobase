package cmd

import (
	"fmt"
	"os"

	"portfolio/services/infrastructure/config/auth"
	"portfolio/services/infrastructure/log"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:               "us",
	Short:             "Run server",
	DisableAutoGenTag: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eways.yaml)")
}

func initConfig() {
	cfg := auth.LoadConfig()
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         cfg.Sentry.DSN,
		Environment: cfg.Environment,
	}); err != nil {
		log.Errorf("Sentry initialization failed: %v\n", err)
	}
}
