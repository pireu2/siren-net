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
	Response string `json:"response"`
}

func AskLLM(c *gin.Context) {
	var request LLMRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		services.RespondError(c, http.StatusBadRequest, err)
		return
	}

	ollamaRequest := map[string]interface{}{
		"prompt": request.Prompt,
		"model":  "deepseek",
		"stream": false,
	}

	body, _ := json.Marshal(ollamaRequest)
	resp, err := http.Post("http://ollama:11434/api/generate", "application/json", bytes.NewBuffer(body))
	if err != nil {
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)

	c.JSON(http.StatusOK, string(respBody))
}
