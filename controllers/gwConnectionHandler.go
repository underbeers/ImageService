package controllers

import (
	config "ImageService/configs"
	"ImageService/models"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

const (
	contentType = "Content-Type"
	appJSON     = "application/json"
)

func HandleInfo(c *fiber.Ctx) error {
	c.Append(contentType, appJSON)
	serviceInfo := GetServiceInfo(config.EnvIP(), config.EnvPort(), c)
	payload, err := json.Marshal(serviceInfo)
	if err != nil {
		logrus.Error(err.Error())
	}

	err = c.Send(payload)
	if err != nil {
		logrus.Error(err.Error())
	}

	return nil
}

func GetServiceInfo(ip, port string, c *fiber.Ctx) *models.Service {
	handles, err := getHandles(c)
	if err != nil {
		logrus.Fatalf("failed to getHandles, %v", err)
	}

	instance := models.Service{
		Name:      "image",
		Label:     "pl_image_service",
		IP:        ip,
		Port:      port,
		Endpoints: nil,
	}
	unprotected, err := getUnprotected()
	if err != nil {
		logrus.Fatalf("failed to getHandles, %v", err)
	}

	for k, v := range handles {
		// skip endpoint-info
		if k == "endpoint-info/" {
			continue
		}
		endpoint := models.Endpoint{
			URL:       k,
			Protected: true,
			Methods:   v,
		}
		if unprotected[k] {
			endpoint.Protected = false
		}
		instance.Endpoints = append(instance.Endpoints, endpoint)
	}

	return &instance
}

func getHandles(c *fiber.Ctx) (map[string][]string, error) {
	data := make(map[string][]string)
	handlers := c.App().GetRoutes()
	logrus.Info(handlers)

	for _, v := range handlers {
		path := v.Path
		method := []string{v.Method}
		path = strings.Split(path, "/api/v1/")[1]
		data[path] = append(data[path], method...)
	}

	return data, nil

	/*err := srv.router.Walk(
		func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			path, _ := route.GetPathTemplate()
			n, _ := route.GetMethods()
			path = strings.Split(path, "/api/v1/")[1]
			d, ok := data[path]
			if ok {
				n = append(n, d...)
				data[path] = n

				return nil
			}
			data[path] = n

			return nil
		})
	if err != nil {
		return nil, err
	}

	return data, nil*/
}

func getUnprotected() (map[string]bool, error) {
	// Read's list of unprotected endpoints
	lst, err := os.OpenFile("service.json", os.O_RDONLY, 0o600) //nolint:gomnd
	if err != nil {
		return nil, errors.New("can't open file")
	}
	reader, err := io.ReadAll(lst)
	if err != nil {
		return nil, errors.New("can't read file")
	}
	data := struct {
		URLS []string `json:"urls"`
	}{}
	err = json.Unmarshal(reader, &data)
	if err != nil {
		return nil, errors.New("can't unmarshal data")
	}
	result := make(map[string]bool)
	for _, k := range data.URLS {
		result[k] = true
	}

	return result, nil
}
