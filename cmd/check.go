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
	"bufio"
	"bytes"
	"crypto/sha1"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/alegrey91/hibp/utils"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check \"yourpassword\"",
	Short: "Check if your password has been found in some data set.",
	Long: `hibp uses the haveibeenpwned.com APIs to check your password 
validity in order to protect the value of the source password 
using the k-anonymity principle.`,

	Run: func(cmd *cobra.Command, args []string) {

        var typedPassword []byte
		var definedURL = "https://api.pwnedpasswords.com/range/"

		if len(args) != 1 {
		    // Check if arguments are passed from Stdin pipe.
		    if utils.IsStdinPresent() {
		        reader := bufio.NewReader(os.Stdin)
		        input, err := reader.ReadString('\n')
		        input = strings.TrimSuffix(input, "\n")
		        if err != nil {
		            log.Fatal("Unable to read from Stdin.")
		        }
		        typedPassword = []byte(input)
		    } else {
			    log.Fatal("Not enough arguments.")
			}
		} else {
		    typedPassword = []byte(args[0])
		}

		hashedPassword := sha1.Sum(typedPassword)

		// Configuuring TLS for network communication.
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		// Convert hex format to string.
		hashStr := hex.EncodeToString(hashedPassword[:])
		// Retrieve first 5 characters to query the password.
		hashSubStr := hashStr[:5]
		// Compose API url.
		definedURL += hashSubStr

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
		bodyReader := bytes.NewReader(bodyBytes)
		scanner := bufio.NewScanner(bodyReader)

		for scanner.Scan() {
			// Extracting the provided hashes line by line.
			resultSubStr := strings.Split(scanner.Text(), ":")
			hashResult := strings.ToUpper(hashSubStr + resultSubStr[0])
			hashString := strings.ToUpper(hashStr)

			if hashResult == hashString {
				fmt.Printf("Your password appears %s times in data set.\n", resultSubStr[1])
				os.Exit(1)
			}
		}

		fmt.Printf("Your password is good.\n")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
