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

type TemplatesResponse struct {
	Type      string `json:"Type"`
	Available int32  `json:"Available"`
}

// templatesCmd represents the templates command
var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "Gets available templates",
	Long:  "Gets available templates",
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

		endpoint := kubeCluster + "templates"
		resp, err := client.Post(endpoint, "application/json", bytes.NewBuffer(json_data))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var res []TemplatesResponse

		json.NewDecoder(resp.Body).Decode(&res)

		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', tabwriter.AlignRight)

		fmt.Fprintf(writer, "TYPE\tAVAILABLE\n")
		for _, template := range res {
			fmt.Fprintf(writer, "%v\t%v\n", template.Type, template.Available)
		}
		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(templatesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// templatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// templatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
