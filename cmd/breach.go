/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"os"
	"log"
	"strings"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/alegrey91/hibp/private"
)

// breachCmd represents the breach command
var breachCmd = &cobra.Command{
	Use:   "breach \"site\"",
	Short: "Check if a specific site has been compromised.",
	Long: `hibp uses the haveibeenpwned.com APIs to check if the specified site
has been compromised. This show you informations about the compromise.`,
	Run: func(cmd *cobra.Command, args []string) {

        var typedSite []byte
		var definedURL = "https://haveibeenpwned.com/api/v3/breach/"
		type Site struct {
            Name string                     `json:"Name"`
            Title string                    `json:"Title"`
            Domain string                   `json:"Domain"`
            BreachDate string               `json:"BreachDate"`
            AddedDate string                `json:"AddedDate"`
            ModifedDate string              `json:"ModifiedDate"`
            PwnCount int                    `json:"PwnCount"`
            Description string              `json:"Description"`
            LogoPath string                 `json:"LogoPath"`
            DataClasses []string            `json:"DataClasses"`
            IsVerified bool                 `json:"IsVerified"`
            IsFabricated bool               `json:"IsFabricated"`
            IsSensitive bool                `json:"IsSensitive"`
            IsRetired bool                  `json:"IsRetired"`
            IsSpamList bool                 `json:"IsSpamList"`
        }

		if len(args) != 1 {
		    // Check if arguments are passed from Stdin pipe.
		    if utils.IsStdinPresent() {
		        reader := bufio.NewReader(os.Stdin)
		        input, err := reader.ReadString('\n')
		        input = strings.TrimSuffix(input, "\n")
		        if err != nil {
		            log.Fatal("Unable to read from Stdin.")
		        }
		        typedSite = []byte(input)
		    } else {
			    log.Fatal("Not enough arguments.")
			}
		} else {
		    typedSite = []byte(args[0])
		}

		// Configuuring TLS for network communication.
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

        definedURL += string(typedSite)

		// Make request to server API.
		req, err := http.NewRequest("GET", definedURL, nil)
		if err != nil {
			log.Fatal(err)
		}

		// Retrieving response.
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			log.Fatal("Unexpected Status Code.")
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

        site := Site{}
        jsonResponse := json.Unmarshal(bodyBytes, &site)
        if jsonResponse != nil {
            log.Fatal(jsonResponse)
        }
		
        fmt.Printf("Name: %s\n", site.Name)
        fmt.Printf("Title: %s\n", site.Title)
        fmt.Printf("Domain: %s\n", site.Domain)
        fmt.Printf("BreachDate: %s\n", site.BreachDate)
        fmt.Printf("PwnCount: %d\n", site.PwnCount)
        fmt.Printf("Description: %s\n", site.Description)
	},
}

func init() {
	rootCmd.AddCommand(breachCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// breachCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// breachCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
