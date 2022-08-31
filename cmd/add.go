/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: Run,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type ErrorResponse struct {
	Message string
}

func Run(cmd *cobra.Command, args []string) error {
	err := AddItem(args[0])
	if err != nil {
		return err
	}
	fmt.Println("Added!")
	return nil
}

func AddItem(itemName string) error {
	err := CallAddItemApi(itemName)
	if err != nil {
		return err
	}
	return nil
}

func CallAddItemApi(name string) error {
	param := `{"name":"` + name + `"}`

	url, err := url.JoinPath(viper.GetString("url"), "/api/items")
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(param)),
	)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+viper.GetString("api_key"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 201 {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var error = ErrorResponse{}
	json.Unmarshal(body, &error)

	return fmt.Errorf(error.Message)
}
