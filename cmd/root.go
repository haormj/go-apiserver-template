package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/haormj/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/haormj/go-apiserver-template/internal/option"
	"github.com/haormj/go-apiserver-template/internal/provider"
	"github.com/haormj/go-apiserver-template/internal/service"
	"github.com/haormj/go-apiserver-template/pkg/version"
)

var rootCmd = &cobra.Command{Version: version.FullVersion(),
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		if err := viper.ReadInConfig(); err != nil {
			log.Logger.Errorw("viper.ReadInConfig error", "err", err)
			return
		}

		var options option.Options
		if err := viper.Unmarshal(&options); err != nil {
			log.Logger.Errorw("viper.Unmarshal error", "err", err)
			return
		}

		svc, err := service.New()
		if err != nil {
			log.Logger.Errorw("service.New error", "err", err)
			return
		}

		p, err := provider.New(&options.Provider, svc)
		if err != nil {
			log.Logger.Errorw("provider.New error", "err", err)
			return
		}

		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		if err := p.Run(ctx); err != nil {
			log.Logger.Errorw("p.Run error", "err", err)
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
