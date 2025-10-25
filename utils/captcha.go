package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type capSolverTask struct {
	ClientKey string `json:"clientKey"`
	Task      struct {
		Module string `json:"module"`
		Type   string `json:"type"`
		Body   string `json:"body"`
	} `json:"task"`
}

type capSolverResponse struct {
	ErrorID  int `json:"errorId"`
	Solution struct {
		Text string `json:"text"`
	} `json:"solution"`
	ErrorDescription string `json:"errorDescription,omitempty"`
}

// ResolveCaptcha sends an image (base64) to CapSolver and returns the recognized text
func ResolveCaptcha(imageBase64 string, key string) (string, error) {
	task := capSolverTask{
		ClientKey: key,
	}
	task.Task.Module = "common"
	task.Task.Type = "ImageToTextTask"
	task.Task.Body = imageBase64

	payload, _ := json.Marshal(task)

	resp, err := http.Post("https://api.capsolver.com/createTask", "application/json", bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to contact CapSolver: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("CapSolver returned HTTP %d: %s", resp.StatusCode, string(body))
	}

	var result capSolverResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse CapSolver response: %v", err)
	}

	if result.ErrorID != 0 {
		return "", fmt.Errorf("CapSolver error: %s", result.ErrorDescription)
	}

	if result.Solution.Text == "" {
		return "", errors.New("no solution text returned from CapSolver")
	}

	return result.Solution.Text, nil
}
