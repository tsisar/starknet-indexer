package abiutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tsisar/extended-log-go/log"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type rpcRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type rpcResponse struct {
	Result struct {
		Abi json.RawMessage `json:"abi"`
	} `json:"result"`
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func FetchAndSaveABI(contractName, contractAddress, rpcURL, outputDir string) error {
	payload := rpcRequest{
		Jsonrpc: "2.0",
		Method:  "starknet_getClassAt",
		Params:  []interface{}{"latest", contractAddress},
		ID:      1,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	resp, err := http.Post(rpcURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("http post: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respText, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("RPC error: %s", string(respText))
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}

	var parsed rpcResponse
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return fmt.Errorf("unmarshal response: %w", err)
	}

	if parsed.Error != nil {
		return fmt.Errorf("RPC response error: %s", parsed.Error.Message)
	}

	if len(parsed.Result.Abi) == 0 {
		return fmt.Errorf("no ABI returned for %s", contractAddress)
	}

	prettyAbi := &bytes.Buffer{}
	if err := json.Indent(prettyAbi, parsed.Result.Abi, "", "  "); err != nil {
		return fmt.Errorf("format ABI: %w", err)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("create dir: %w", err)
	}

	outputFile := filepath.Join(outputDir, strings.ToLower(contractName)+".json")
	if err := os.WriteFile(outputFile, prettyAbi.Bytes(), 0644); err != nil {
		return fmt.Errorf("write ABI file: %w", err)
	}

	log.Debugf("ABI saved for %s → %s", contractName, outputFile)
	return nil
}
