package config

import (
	"AlekseyPromet/authorization/internal/app"
	"AlekseyPromet/authorization/internal/logger"
	"AlekseyPromet/authorization/internal/models"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "file", "", "config file (default is config/sso-api.yaml)")
	rootCmd.PersistentFlags().StringP("port", "p", "33088", "servise start on port")
	viper.RegisterAlias("f", "file")
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.SetDefault("author", "Aleksey Perminov <promet.alex@gmail.com>")
	viper.SetDefault("license", "GPL-3.0")
}

var (
	// Used for flags.
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "sso [options]",
		Short: "Simple SSO server over grpc",
		Long:  "Complete documentation is available at https://github.com/AlekseyPromet/simple_sso/blob/main/README.md",
		RunE: func(cmd *cobra.Command, args []string) error {

			config := &models.Config{
				Env:       models.TypeEnv(viper.GetString("env")),
				StorgeDsn: viper.GetString("storge_dsn"),
				TokenTTL:  viper.GetDuration("token_ttl"),
				Grpc: models.Grpc{
					Port:    viper.GetString("port"),
					Timeout: viper.GetDuration("timeout"),
				},
			}

			service := app.NewApp(config)
			log := logger.NewLogger(config.Env)

			fx.New(
				fx.Provide(service.Run),
				fx.Invoke(func(*http.Server) {}),
				fx.WithLogger(log),
			).Run()

			return nil
		},
	}
)

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in config directory with name "sso-api.yaml"
		viper.AddConfigPath("./config")
		viper.SetConfigType("yaml")
		viper.SetConfigName("sso-api.yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
