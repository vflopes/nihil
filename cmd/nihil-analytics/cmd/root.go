package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/spf13/cobra"
	"github.com/vflopes/nihil/pkg/analysis"
	"github.com/vflopes/nihil/pkg/worker"
	"golang.org/x/net/context"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "nihil-analytics [input-file] [output-file]",
	Args:  cobra.ExactArgs(2),
	Short: "Executes nihil analytics",
	Long: `Executes nihil analytics.
	- The input file must be a JSON encoded nihil.analytics.Pipeline message
	- The output file will be a JSON encoded nihil.analytics.Axis message
`,
	Run: func(cmd *cobra.Command, args []string) {

		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC3339Nano,
			},
		)

		input, err := ioutil.ReadFile(args[0])
		if err != nil {
			log.Fatal().Err(err).Msg("unable to read input file")
		}

		pipeline := new(analysis.Pipeline)

		err = jsonpb.Unmarshal(bytes.NewReader(input), pipeline)
		if err != nil {
			log.Fatal().Err(err).Msg("unable to parse input file JSON")
		}

		log.Info().Msg("starting nihil-analytics")

		worker := worker.NewWorker()

		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("ANALYSIS_TIMEOUT"))
		defer cancel()

		axis := worker.Execute(ctx, pipeline)
		marshaler := jsonpb.Marshaler{}

		fdout, err := os.Create(args[1])
		if err != nil {
			log.Fatal().Err(err).Msg("unable to create output file")
		}

		err = marshaler.Marshal(fdout, axis)
		if err != nil {
			log.Fatal().Err(err).Msg("unable to encode output file JSON")
		}

		log.Info().Str("output_file", args[1]).Msg("wrote output file")

		err = fdout.Close()
		if err != nil {
			log.Fatal().Err(err).Msg("unable to close output file")
		}

		log.Info().Msg("end of nihil-analytics")

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nihil.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.SetDefault("ANALYSIS_TIMEOUT", 5*time.Second)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".nihil")
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()
}
