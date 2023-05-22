package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/nurimansyah89/terraform-elastis-provider/elastis/client"
)

type VMHandler struct {
	c *client.ElastisClient
}

func NewVM(c *client.ElastisClient) *VMHandler {
	return &VMHandler{
		c: c,
	}
}

func (h *VMHandler) GetVM(location string, uuid string) (*client.VMResponse, error) {
	var result client.VMResponse
	url := fmt.Sprintf("%s/%s/user-resource/vm?uuid=%s", h.c.EndpointURL, location, uuid)

	body, err := h.c.Fetch(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (h *VMHandler) CreateVM(location string, payload client.VMPayload) (*client.VMResponse, error) {
	var result client.VMResponse
	url := fmt.Sprintf("%s/%s/user-resource/vm", h.c.EndpointURL, location)

	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	body, err := h.c.Post(url, rb)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (h *VMHandler) UpdateVM(location string, payload client.VMPayloadUpdate, memory int, vcpu int) (*client.VMResponse, error) {
	var result client.VMResponse
	url := fmt.Sprintf("%s/%s/user-resource/vm", h.c.EndpointURL, location)

	// Check payload first
	if payload.Memory != memory && payload.VCPU != vcpu {
		// Stop VM first
		stopURL := fmt.Sprintf("%s/%s/user-resource/vm/stop", h.c.EndpointURL, location)
		p := map[string]string{
			"uuid": payload.UUID,
		}
		rb, err := json.Marshal(p)
		if err != nil {
			return nil, err
		}

		body, err := h.c.Post(stopURL, rb)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
	}

	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	body, err := h.c.Patch(url, rb)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	// Start only when memory or vcpu are changed
	if payload.Memory != memory && payload.VCPU != vcpu {
		stopURL := fmt.Sprintf("%s/%s/user-resource/vm/start", h.c.EndpointURL, location)
		p := map[string]string{
			"uuid": payload.UUID,
		}
		rb, err := json.Marshal(p)
		if err != nil {
			return nil, err
		}

		body, err := h.c.Post(stopURL, rb)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (h *VMHandler) DeleteVM(location string, payload client.VMPayloadDelete) error {
	url := fmt.Sprintf("%s/%s/user-resource/vm", h.c.EndpointURL, location)

	rb, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = h.c.Delete(url, rb)
	if err != nil {
		return err
	}

	// Delete floating IP
	if _, err := h.DeleteFloatingIP(location); err != nil {
		return err
	}

	return nil
}

func (h *VMHandler) DeleteFloatingIP(location string) (bool, error) {
	// First get floating IP based on VM uuid
	var floatingIPResult []client.FloatingIPInfo
	getUrl := fmt.Sprintf("%s/%s/network/ip_addresses", h.c.EndpointURL, location)

	body, err := h.c.Fetch(getUrl)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(body, &floatingIPResult)
	if err != nil {
		return false, err
	}

	// Filter un-assigned IP
	filteredIPS := []client.FloatingIPInfo{}

	for i := range floatingIPResult {
		if floatingIPResult[i].AssignedTo == nil {
			filteredIPS = append(filteredIPS, floatingIPResult[i])
		}
	}

	// If it has value
	if len(floatingIPResult) > 0 {
		for _, v := range filteredIPS {
			deleteUrl := fmt.Sprintf("%s/%s/network/ip_addresses/%s", h.c.EndpointURL, location, v.Address)
			_, err = h.c.Delete(deleteUrl, []byte(""))
			if err != nil {
				return false, err
			}
		}
	}

	return true, nil
}
