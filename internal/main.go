package internal

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env/v7"
	"github.com/cloudflare/cloudflare-go"
)

type Input struct {
	Debug          bool     `env:"PLUGIN_DEBUG" envDefault:false`
	ApiToken       string   `env:"PLUGIN_API_TOKEN,required"`
	ZoneIdentifier string   `env:"PLUGIN_ZONE_IDENTIFIER,required"`
	Action         string   `env:"PLUGIN_ACTION,required"`
	List           []string `env:"PLUGIN_LIST" envSeparator:","`
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func handleCloudflareError(err error) {
	if err != nil {
		fmt.Printf("err: %#v\n", err)
		os.Exit(3)
	}
}

func Run() {
	// Define valid actions
	validActions := []string{
		"purge_everything",
		"purge_hosts",
		"purge_files",
		"purge_tags",
	}
	// Parse environmental variables
	input := Input{}
	if err := env.Parse(&input); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	// Normalize input
	input.Action = strings.ToLower(input.Action)
	// Semantic checks
	if !contains(validActions, input.Action) {
		fmt.Printf(
			"err: PLUGIN_ACTION must equal one of [%s]\n",
			strings.Join(validActions, ","),
		)
		os.Exit(2)
	} else if input.Action != "purge_everything" && len(input.List) < 1 {
		fmt.Printf(
			"err: PLUGIN_LIST must be set with '%s' action\n",
			input.Action,
		)
		os.Exit(2)
	}
	// Create client for cloudflare api communication
	api, err := cloudflare.NewWithAPIToken(input.ApiToken)
	handleCloudflareError(err)
	// Print input if debug mode is on
	if input.Debug {
		fmt.Printf("Input: %+v\n", input)
	}
	// Execute action
	request := cloudflare.PurgeCacheRequest{}
	switch input.Action {
	case "purge_everything":
		request.Everything = true
		break
	case "purge_hosts":
		request.Hosts = input.List
		break
	case "purge_files":
		request.Files = input.List
		break
	case "purge_tags":
		request.Tags = input.List
		break
	}
	res, err := api.PurgeCache(context.Background(), input.ZoneIdentifier, request)
	handleCloudflareError(err)
	if input.Debug {
		fmt.Printf("PurgeCache: %+v\n", res)
	}
}
