/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var kubeCluster string
var userToken string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cluster",
	Short: "CLI to manage clusters",
	Long:  "CLI to manage clusters",
	Run:   func(cmd *cobra.Command, args []string) { fmt.Println("Hello CLI") },
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
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configFlags := genericclioptions.NewConfigFlags(true)
	kubeconfig, err := configFlags.ToRawKubeConfigLoader().RawConfig()

	kubeCluster = os.Getenv("CLUSTER_TEMPLATES_API")
	authInfo := kubeconfig.Contexts[kubeconfig.CurrentContext].AuthInfo
	userToken = kubeconfig.AuthInfos[authInfo].Token

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
