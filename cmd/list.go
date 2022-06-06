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
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

type ClusterReponse struct {
	Name              string `json:"Name"`
	Type              string `json:"Type"`
	KubeadminPassword string `json:"KubeadminPassword"`
	URL               string `json:"URL"`
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		values := map[string]string{"token": userToken}
		json_data, err := json.Marshal(values)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		endpoint := kubeCluster + "clusters"
		resp, err := client.Post(endpoint, "application/json", bytes.NewBuffer(json_data))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var res []ClusterReponse

		json.NewDecoder(resp.Body).Decode(&res)

		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', tabwriter.AlignRight)

		fmt.Fprintf(writer, "NAME\tTYPE\tKUBEADMIN_PASS\tURL\n")
		for _, cluster := range res {
			fmt.Fprintf(writer, "%v\t%v\t%v\t%v\n", cluster.Name, cluster.Type, cluster.KubeadminPassword, cluster.URL)
		}
		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
