package handlers

import (
	"encoding/json"

	"github.com/nurimansyah89/terraform-elastis-provider/elastis/client"
)

type LocationHandler struct {
	c *client.ElastisClient
}

func NewLocation(c *client.ElastisClient) *LocationHandler {
	return &LocationHandler{
		c: c,
	}
}

func (h *LocationHandler) GetLocations() ([]client.LocationInfo, error) {
	var locations []client.LocationInfo
	url := h.c.EndpointURL + "/config/locations"

	data, err := h.c.Fetch(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return nil, err
	}

	return locations, nil
}
