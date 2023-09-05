package main

import (
	"encoding/json"
	"flag"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"level0/internal/model"
	DataToSend "level0/prep"
	"log"
)

func main() {
	var (
		clusterID string
		clientID  string
		URL       string
		async     bool
		userCreds string
	)

	flag.StringVar(&URL, "s", stan.DefaultNatsURL, "The nats server URLs (separated by comma)")
	flag.StringVar(&URL, "server", stan.DefaultNatsURL, "The nats server URLs (separated by comma)")
	flag.StringVar(&clusterID, "c", "test-cluster", "The NATS Streaming cluster ID")
	flag.StringVar(&clusterID, "cluster", "test-cluster", "The NATS Streaming cluster ID")
	flag.StringVar(&clientID, "id", "stan-pub", "The NATS Streaming client ID to connect with")
	flag.StringVar(&clientID, "clientid", "stan-pub", "The NATS Streaming client ID to connect with")
	flag.BoolVar(&async, "a", false, "Publish asynchronously")
	flag.BoolVar(&async, "async", false, "Publish asynchronously")
	flag.StringVar(&userCreds, "cr", "", "Credentials File")
	flag.StringVar(&userCreds, "creds", "", "Credentials File")

	log.SetFlags(0)

	flag.Parse()

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Streaming Example Publisher")}
	// Use UserCredentials
	if userCreds != "" {
		opts = append(opts, nats.UserCredentials(userCreds))
	}

	//Connecting and publishing to the nats.
	newConn, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, URL)
	}
	defer newConn.Close()
	output := &model.Orders{}
	json.Unmarshal([]byte(DataToSend.RawData), output)
	err = newConn.Publish("str", []byte(DataToSend.RawData))
	if err != nil {
		log.Fatal("unanable to send msg")
	}

}
