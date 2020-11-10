package es

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	clientSingleNode *ClientSingleNode
	clientMultiNode  *ClientMultiNode
)

type ClientHandler interface {
	write(index string, docId string, body []byte) (map[string]interface{}, error)
}

type ClientSingleNode struct {
	Username, Password  string
	Address             string //   format: https://host:port
	Host, Port          string
	MaxIdleConnsPerHost int
	client              *elasticsearch.Client
}

type ClientMultiNode struct {
	Username, Password  string
	Addresses           []string
	MaxIdleConnsPerHost int
	client              *elasticsearch.Client
}

/*
 * create client as ES's single node
 */
func (conf ClientSingleNode) Init() error {
	var err error
	getDefaultConf(&conf.Host, &conf.Port, &conf.MaxIdleConnsPerHost)
	if conf.Address == "" {
		conf.Address = fmt.Sprintf("http://%s:%s", conf.Host, conf.Port)
	}
	addresses := make([]string, 0)
	addresses = append(addresses, conf.Address)
	conf.client, err = initClient(conf.Username, conf.Password, addresses, conf.MaxIdleConnsPerHost)
	clientSingleNode = &conf
	return err
}

/*
 * create client as ES's multi node
 */
func (conf ClientMultiNode) Init() error {
	var err error
	conf.client, err = initClient(conf.Username, conf.Password, conf.Addresses, conf.MaxIdleConnsPerHost)
	clientMultiNode = &conf
	return err
}

/*
 * create default localhost and default port config
 */
func getDefaultConf(host, port *string, maxIdleConnsPerHost *int) {
	if host == nil {
		hostDefault := "localhost"
		host = &hostDefault
	}
	if port == nil {
		portDefault := "9200"
		port = &portDefault
	}
	if maxIdleConnsPerHost == nil {
		maxIdleConnsPerHostDefault := 10
		maxIdleConnsPerHost = &maxIdleConnsPerHostDefault
	}
}
/*
 * init client
 */

func initClient(username, password string, addresses []string, maxIdleConnsPerHost int) (*elasticsearch.Client, error) {
	esConf := elasticsearch.Config{
		Addresses: addresses,
		Username:  username,
		Password:  password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   maxIdleConnsPerHost,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
			},
		},
	}
	return elasticsearch.NewClient(esConf)
}

func (conf ClientSingleNode) write(index string, docId string, body []byte) (map[string]interface{}, error) {
	return send(index, docId, body, conf.client)
}

func (conf ClientMultiNode) write(index string, docId string, body []byte) (map[string]interface{}, error) {
	return send(index, docId, body, conf.client)
}

/*
 * send ....
 */
func send(index string, docId string, body []byte, trans esapi.Transport) (map[string]interface{}, error) {
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: docId,
		Body:       strings.NewReader(string(body)),
		Refresh:    "true",
	}
	// Perform the request with the client.
	res, err := req.Do(context.Background(), trans)
	if err != nil {
		fmt.Printf("Error getting response: %s \n", err)
	}
	defer func() {
		err = res.Body.Close()

	}()
	var r map[string]interface{}
	if res.IsError() {
		fmt.Printf("[%s] Error indexing document ID=%v\n", res.Status(), docId)
	} else {
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			fmt.Printf("Error parsing the response body: %s\n", err)
		} else {
			fmt.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	return r, err
}

/*
 * get SingleNode has been created
 */
func GetSingleNodeClient() *ClientSingleNode {
	return clientSingleNode
}

/*
 * get MiltiNode has been created
 */
func GetMultiNodeClient() *ClientMultiNode {
	return clientMultiNode
}
