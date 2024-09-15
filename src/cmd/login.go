/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stored_api_key string
var stored_api_secret string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use: "login",
	Run: func(cmd *cobra.Command, args []string) {
		api_key, _ := cmd.Flags().GetString("client-id")
		api_secret, _ := cmd.Flags().GetString("client-secret")
		url := "https://login.aruba.it/auth/realms/cmp-new-apikey/protocol/openid-connect/token"
		method := "POST"

		payload := strings.NewReader("grant_type=client_credentials&client_id=" + api_key + "&client_secret=" + api_secret)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Cookie", "COMMON_BACKEND=gob-com5")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// Define a struct to capture the necessary fields from the response
		var data struct {
			AccessToken      string `json:"access_token"`
			ExpiresIn        int    `json:"expires_in"`
			RefreshExpiresIn int    `json:"refresh_expires_in"`
			TokenType        string `json:"token_type"`
			NotBeforePolicy  int    `json:"not_before_policy"`
			Scope            string `json:"scope"`
		}

		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Println("Error unmarshalling response:", err)
			return
		}

		stored_api_key = api_key
		fmt.Println("Access Token:", data.AccessToken)
		viper.AutomaticEnv()
		viper.Set("APIKEY", "asdi")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().String("client-id", "", "specifies the api-key")
	loginCmd.Flags().String("client-secret", "", "specifies the secret key")
	viper.AddConfigPath("./cmd")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
