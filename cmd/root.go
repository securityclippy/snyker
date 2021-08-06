// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/securityclippy/snyker/pkg/outwriter"
	"github.com/securityclippy/snyker/pkg/snykclient"
	"github.com/sirupsen/logrus"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"encoding/json"
)

var cfgFile string
var SnykClient *snykclient.SnykClient
var log *logrus.Logger
var output string
var Out *outwriter.OutWriter
var Org string
var SnykConfig SnykCfg


type SnykCfg struct {
	Orgs []SnykOrg `json:"orgs"`
}

type SnykOrg struct {
	Name string `json:"name"`
	APIKey string `json:"api_key"`
}
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "snyker",
	Short: "A Golang CLI for the snyk api",
	Long: `An easy way to interact with the snyk api via the command line.  THis is a WIP. 

NOTE: Currently pulls its access token from the SNYK_TOKEN env var`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	log = logrus.New()
	SnykConfig = SnykCfg{}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.snyker.yaml)")
	rootCmd.PersistentFlags().StringVar(&Org, "org", "", "snyk org to use")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&output, "output", "", "json", "output type")
	//rootCmd.PersistentFlags().StringVarP()

	Orgs := SnykConfig.Orgs


	for _, o := range Orgs {
		fmt.Println(o)
	}

	var err error
	switch {
	case len(Orgs) > 1:
		for _, o := range Orgs {
			fmt.Println(o)
			if o.Name == Org {
				SnykClient, err = snykclient.NewSnykClient(o.APIKey)
				if err != nil {
					log.Fatal(err)
				}
			}
			fmt.Printf("Using Org: %s", o.Name)
		}
	case len(Orgs) == 1:
		SnykClient, err = snykclient.NewSnykClient(Orgs[0].APIKey)
		if err != nil {
			log.Fatal(err)
		}
	default:
		SnykClient, err = snykclient.NewSnykClient(os.Getenv("SNYK_TOKEN"))
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Using Org: %s\n", SnykClient.Org)
	log = logrus.New()
	Out = &outwriter.OutWriter{}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".snyker" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".snyker")
	}

	viper.AutomaticEnv() // read in environment variables that match

	js, err := ioutil.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(js, &SnykConfig)

	if err != nil {
		log.Fatal(err)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

}
