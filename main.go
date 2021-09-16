package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/TylerBrock/colorjson"
	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	flags := struct {
		consumerKey    string
		consumerSecret string
		geocode        string
		count          int
	}{}

	flag.StringVar(&flags.consumerKey, "consumer-key", "", "Twitter Consumer Key")
	flag.StringVar(&flags.consumerSecret, "consumer-secret", "", "Twitter Consumer Secret")
	flag.StringVar(&flags.geocode, "geocode", "", "lat,lon,accuracy")
	flag.IntVar(&flags.count, "count", 1000, "max num of results")
	flag.Parse()
	flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")

	if flags.consumerKey == "" || flags.consumerSecret == "" {
		log.Fatal("Application Access Token required")
	}

	// parse the geocode argument to ensure it's in the correct format
	// expecting "lat,lon,radius+km|mi" e.g. "50.1234,-5.1234,10km"
	g := strings.Split(flags.geocode, ",")
	if len(g) < 3 {
		log.Fatal("Incorrect geocode format, expecting like '50.1234,-5.1234,10km'")
	}

	decimal_regex := regexp.MustCompile(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`)
	unit_regex := regexp.MustCompile(`^\d+(km|mi)$`)

	if !(decimal_regex.MatchString(g[0]) && decimal_regex.MatchString(g[1]) && unit_regex.MatchString(g[2])) {
		log.Fatal("Incorrect geocode format, expecting like '50.1234,-5.1234,10km'")
	}

	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     flags.consumerKey,
		ClientSecret: flags.consumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// search tweets
	searchTweetParams := &twitter.SearchTweetParams{
		Geocode:   flags.geocode,
		Count:     flags.count,
		TweetMode: "extended",
	}

	search, _, _ := client.Search.Tweets(searchTweetParams)

	// pretty print the json output
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(search)

	var obj map[string]interface{}
	json.Unmarshal([]byte(reqBodyBytes.Bytes()), &obj)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)
	fmt.Println(string(s))

}
