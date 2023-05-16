package services

import (
	config "ImageService/configs"
	"ImageService/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	attempts = 3
	timeout  = 5
	protocol = "http"
	baseURL  = "/api/v1/"
)

func GWConnection() {
	logger := log.Default()

	var errorCnt int
	var er error
	for errorCnt < attempts {
		time.Sleep(time.Second * timeout)
		logger.Printf("Attempt to connect to APIGateway")
		if err := pingAPIGateway(); err != nil {
			errorCnt++
			er = err
		} else {
			break
		}
	}
	if errorCnt >= attempts {
		logger.Fatal(fmt.Errorf("failed to send info to the APIGateway %v", er))
	}
	err := helloAPIGateway()
	if err != nil {
		logger.Fatal(err)
	}
}

func helloAPIGateway() error {
	gwHost := config.EnvGWIP()
	gwPort := config.EnvGWPort()
	gatewayURL, err := url.Parse(
		protocol + "://" + gwHost + ":" + gwPort + baseURL + "hello")
	if err != nil {
		return fmt.Errorf("can't parse url for endpoint 'hello'")
	}

	log.Printf("connection gateway url %s", gatewayURL.String())

	info := &models.Service{
		Name:      "image",
		Port:      config.EnvPort(),
		IP:        config.EnvIP(),
		Endpoints: nil,
	}
	jsonStr, err := json.Marshal(info)
	if err != nil {
		return fmt.Errorf("failed to marshal data")
	}

	go knock(gatewayURL.String(), jsonStr)

	return nil
}

func knock(url string, payload []byte) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if resp == nil {
		log.Println("can't say Hello to Gateway", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if err != nil {
		log.Println("knock() Post Error", err)
	}
	if resp.StatusCode == http.StatusOK {
		log.Println("Successfully greet ApiGateway")
	}
}

func pingAPIGateway() error {
	gwURL, err := gatewayURL()
	if err != nil {
		return err
	}
	resp, err := http.Get(gwURL.String())

	if resp == nil {
		return fmt.Errorf("can't ping APIGateway")
	}
	if err != nil {
		return fmt.Errorf("error pingAPIGateway http.Get(gwURL)")
	}
	if err := resp.Body.Close(); err != nil {
		log.Println("can't close response body")
	}

	return nil
}

func gatewayURL() (*url.URL, error) {
	gwHost := config.EnvGWIP()
	gwPort := config.EnvGWPort()
	isPort := config.EnvPort()
	gwURL, err := url.Parse(protocol + "://" + gwHost + ":" + gwPort + baseURL + "hello?service=image&port=" + isPort)
	if err != nil {
		return nil, fmt.Errorf("failed to parse gatewayURL")
	}
	log.Printf("Connection gateway url %s", gwURL.String())

	return gwURL, nil
}
