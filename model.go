package netgear_nighthawk_m5_api

import (
	"encoding/json"
	"fmt"
	"github.com/silvercory/netgear_nighthawk_m5_api/models"
	"net/http"
)

// ModelInfo is a general information request
func (n *NetgearNighthawkM5Api) ModelInfo() (models.ModelInfo, error) {
	req, err := http.NewRequest(http.MethodGet, n.routerBaseURL+"/api/model.json?internalapi=1&x="+randomString(5), nil)
	if err != nil {
		return models.ModelInfo{}, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := n.httpClient.Do(req)
	if err != nil {
		return models.ModelInfo{}, fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()
	var response models.ModelInfo
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return models.ModelInfo{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return response, nil
}
