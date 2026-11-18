package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	uptimescout "github.com/jsonscraper/uptimescout-api-go"
	"net/url"
	"os"
	"time"
)

func run() int {
	limit := flag.String("limit", "10", "limit parameter")
	offset := flag.String("offset", "0", "offset parameter")
	region := flag.String("region", "GB", "region parameter")
	lang := flag.String("lang", "en", "lang parameter")
	flag.Parse()
	key := os.Getenv("UPTIMESCOUT_LICENSE_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "Set UPTIMESCOUT_LICENSE_KEY in your environment")
		return 1
	}
	client := uptimescout.New(key)
	if base := os.Getenv("UPTIMESCOUT_BASE_URL"); base != "" {
		client.BaseURL = base
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, err := client.Request(ctx, "incidents", url.Values{
		"limit":  {*limit},
		"offset": {*offset},
		"region": {*region},
		"lang":   {*lang},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "UptimeScout request failed; check key, plan, parameters and network")
		return 1
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(response); err != nil {
		return 1
	}
	return 0
}

func main() { os.Exit(run()) }
