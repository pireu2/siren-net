package handlers

import (
	"backend/internal/services"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type LLMRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

type LLMResponse struct {
	Model         string `json:"model"`
	Response      string `json:"response"`
	TotalDuration int64  `json:"total_duration,omitempty"`
}

func AskLLM(c *gin.Context) {
	var request LLMRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		services.RespondError(c, http.StatusBadRequest, err)
		return
	}

	ollamaRequest := map[string]interface{}{
		"prompt":     request.Prompt,
		"model":      "deepseek",
		"stream":     false,
		"keep_alive": -1,
	}

	body, _ := json.Marshal(ollamaRequest)
	resp, err := http.Post("http://ollama:11434/api/generate", "application/json", bytes.NewBuffer(body))
	if err != nil {
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	var ollamaResponse map[string]interface{}
	if err := json.Unmarshal(respBody, &ollamaResponse); err != nil {
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	response := LLMResponse{
		Model:    ollamaResponse["model"].(string),
		Response: ollamaResponse["response"].(string),
	}

	if duration, ok := ollamaResponse["total_duration"]; ok {
		response.TotalDuration = int64(duration.(float64))
	}

	c.JSON(http.StatusOK, response)
}
