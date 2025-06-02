package handlers

import (
	"backend/internal/config"
	"backend/internal/services"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type SDRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

func callSDAPI(cfg *config.Config, endpoint string, payload interface{}) (map[string]interface{}, error) {

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", cfg.SDUrl+"/"+endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var sdResponse map[string]interface{}
	if err := json.Unmarshal(respBody, &sdResponse); err != nil {
		return nil, err
	}

	return sdResponse, nil
}

func TextToImage(c *gin.Context, cfg *config.Config) {
	var request SDRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		services.RespondError(c, http.StatusBadRequest, err)
		return
	}

	requestPayload := map[string]interface{}{
		"prompt":          request.Prompt,
		"negative_prompt": "sdxl_cyberrealistic_simpleneg-neg",
		"width":           512,
		"height":          512,
		"steps":           30,
		"cfg_scale":       3,
		"sampler_name":    "DPM++ 2S a Karras",
		"sampler_index":   "DPM++ 2S a Karras",
		"batch_size":      1,
		"n_iter":          1,
	}

	sdResponse, err := callSDAPI(cfg, "sdapi/v1/txt2img", requestPayload)
	if err != nil {
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, sdResponse)
}
