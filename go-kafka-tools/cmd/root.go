/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var conn *kafka.Conn // todo get rid of global connectiopn object ??

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-kafka-tools",
	Short: "some sample command line tools for kafka-go libraries",
	Long: ` some sample command line tools for exercising 
	   the kafka-go libraries `,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// create a connection before any command
		dialer := getDialer()
		c, err := dialer.DialContext(context.Background(), "tcp", viper.GetString("bootstrap.servers"))
		if err != nil {
			panic(err.Error())
		}
		conn = c
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		conn.Close()
	},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-kafka-tools.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name "." (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("properties")
		viper.SetConfigName(".go-kafka-tools")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		fmt.Println(" bootstrap servers = ", viper.Get("bootstrap.servers"))
	}
}

func getDialer() kafka.Dialer {
	dialer := kafka.Dialer{
		Timeout: 10 * time.Second,
	}
	if viper.GetString("sasl.mechanisms") == "" { // return palin dialer if no sasl mechanism in config
		return dialer
	}
	// add in sasl parameters from config
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	dialer.SASLMechanism =
		plain.Mechanism{
			Username: viper.GetString("sasl.username"),
			Password: viper.GetString("sasl.password"),
		}
	dialer.DualStack = true
	dialer.TLS = &tls.Config{
		InsecureSkipVerify: true,
		RootCAs:            rootCAs,
	}
	return dialer
}
