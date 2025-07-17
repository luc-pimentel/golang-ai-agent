package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

// TerminalTool provides functionality to execute terminal commands
type TerminalTool struct{}

// NewTerminalTool creates a new terminal tool instance
func NewTerminalTool() *TerminalTool {
	return &TerminalTool{}
}

// Execute runs a terminal command and returns the output
func (t *TerminalTool) Execute(command string) (string, error) {
	if command == "" {
		return "", fmt.Errorf("command cannot be empty")
	}

	// Split command into parts (simple approach)
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return "", fmt.Errorf("invalid command")
	}

	// Execute the command
	cmd := exec.Command(parts[0], parts[1:]...)
	
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	
	// Combine stdout and stderr for complete output
	output := stdout.String()
	if stderr.Len() > 0 {
		if output != "" {
			output += "\n"
		}
		output += "STDERR: " + stderr.String()
	}

	if err != nil {
		return output, fmt.Errorf("command failed: %v", err)
	}

	return output, nil
}

// WebTool provides functionality to access web URLs
type WebTool struct {
	client *http.Client
}

// NewWebTool creates a new web tool instance
func NewWebTool() *WebTool {
	return &WebTool{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Get performs an HTTP GET request to the specified URL
func (w *WebTool) Get(url string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("URL cannot be empty")
	}

	// Add http:// if no protocol specified
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	resp, err := w.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return fmt.Sprintf("Status: %d\nContent:\n%s", resp.StatusCode, string(body)), nil
}

// Post performs an HTTP POST request to the specified URL with data
func (w *WebTool) Post(url, contentType, data string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("URL cannot be empty")
	}

	// Add http:// if no protocol specified
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	if contentType == "" {
		contentType = "application/json"
	}

	resp, err := w.client.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("failed to post to URL: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return fmt.Sprintf("Status: %d\nContent:\n%s", resp.StatusCode, string(body)), nil
}

// GetHeaders gets just the headers from a URL (HEAD request)
func (w *WebTool) GetHeaders(url string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("URL cannot be empty")
	}

	// Add http:// if no protocol specified
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	resp, err := w.client.Head(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch headers: %v", err)
	}
	defer resp.Body.Close()

	var headers strings.Builder
	headers.WriteString(fmt.Sprintf("Status: %d\n", resp.StatusCode))
	headers.WriteString("Headers:\n")
	
	for name, values := range resp.Header {
		for _, value := range values {
			headers.WriteString(fmt.Sprintf("%s: %s\n", name, value))
		}
	}

	return headers.String(), nil
}

// ToolManager manages all available tools
type ToolManager struct {
	Terminal *TerminalTool
	Web      *WebTool
}

// NewToolManager creates a new tool manager with all tools initialized
func NewToolManager() *ToolManager {
	return &ToolManager{
		Terminal: NewTerminalTool(),
		Web:      NewWebTool(),
	}
}