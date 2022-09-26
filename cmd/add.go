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

var addCmd = &cobra.Command{
	Use:                   "add [item]",
	Short:                 "Add an item to your shopping list",
	Long:                  "Add an item to your shopping list.",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE:                  Run,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

type ErrorResponse struct {
	Message string
}

func Run(cmd *cobra.Command, args []string) error {
	err := AddItem(args[0])
	if err != nil {
		return err
	}
	fmt.Println("Successfully added!")
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
