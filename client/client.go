package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	APIKey   string
	Endpoint string
}

func NewAppCheckClient(apiKey string, endpoint string) *Client {
	return &Client{APIKey: apiKey, Endpoint: endpoint}
}

func (client *Client) getRequest(path string) ([]byte, error) {
	apiURL := fmt.Sprintf("%s/api/v1/%s/%s", client.Endpoint, client.APIKey, path)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.APIKey))
	req.Header.Set("Content-Type", "application/json")

	clientHTTP := &http.Client{}
	resp, err := clientHTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get data: %s", resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

func (client *Client) postRequest(path string, body string) ([]byte, error) {
	apiURL := fmt.Sprintf("%s/api/v1/%s/%s", client.Endpoint, client.APIKey, path)
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.APIKey))
	req.Header.Set("Content-Type", "application/json")

	clientHTTP := &http.Client{}
	resp, err := clientHTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to post data: %s", resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

// Define methods for each API endpoint
func (client *Client) GetScanDetails(scanID string) ([]byte, error) {
	return client.getRequest(fmt.Sprintf("scan/%s", scanID))
}

func (client *Client) GetScanRunDetails(scanID, runID string) ([]byte, error) {
	return client.getRequest(fmt.Sprintf("scan/%s/run/%s", scanID, runID))
}

func (client *Client) GetScanRunVulnerabilities(scanID, runID string) ([]byte, error) {
	return client.getRequest(fmt.Sprintf("scan/%s/run/%s/vulnerabilities", scanID, runID))
}

func (client *Client) GetScanRuns(scanID string) ([]byte, error) {
	return client.getRequest(fmt.Sprintf("scan/%s/runs", scanID))
}

func (client *Client) GetScanStatus(scanID string) ([]byte, error) {
	return client.getRequest(fmt.Sprintf("scan/%s/status", scanID))
}

func (client *Client) GetScanVulnerabilities(scanID string) ([]byte, error) {
	return client.getRequest(fmt.Sprintf("scan/%s/vulnerabilities", scanID))
}

func (client *Client) GetScanProfiles() ([]byte, error) {
	return client.getRequest("scanprofiles")
}

func (client *Client) GetScans() ([]byte, error) {
	return client.getRequest("scans")
}

func (client *Client) GetVulnerabilities() ([]byte, error) {
	return client.getRequest("vulnerabilities")
}

func (client *Client) GetVulnerabilityDetails(vulnerabilityID string) ([]byte, error) {
	return client.getRequest(fmt.Sprintf("vulnerability/%s", vulnerabilityID))
}

func (client *Client) AbortScan(scanID string) ([]byte, error) {
	return client.postRequest(fmt.Sprintf("scan/%s/abort", scanID), "")
}

func (client *Client) DeleteScan(scanID string) ([]byte, error) {
	return client.postRequest(fmt.Sprintf("scan/%s/delete", scanID), "")
}

func (client *Client) PauseScan(scanID string) ([]byte, error) {
	return client.postRequest(fmt.Sprintf("scan/%s/pause", scanID), "")
}

func (client *Client) ResumeScan(scanID string) ([]byte, error) {
	return client.postRequest(fmt.Sprintf("scan/%s/resume", scanID), "")
}

func (client *Client) DeleteScanRun(scanID, runID string) ([]byte, error) {
	return client.postRequest(fmt.Sprintf("scan/%s/run/%s/delete", scanID, runID), "")
}

func (client *Client) StartScan(scanID string) ([]byte, error) {
	return client.postRequest(fmt.Sprintf("scan/%s/start", scanID), "")
}

func (client *Client) UpdateScan(scanID, body string) ([]byte, error) {
	return client.postRequest(fmt.Sprintf("scan/%s/update", scanID), body)
}

func (client *Client) CreateScan(body string) ([]byte, error) {
	return client.postRequest("scan/new", body)
}

func (client *Client) DeleteVulnerability(vulnID string) ([]byte, error) {
	return client.postRequest(fmt.Sprintf("vulnerability/%s/delete", vulnID), "")
}

func (client *Client) UpdateVulnerability(vulnerabilityID, body string) ([]byte, error) {
	return client.postRequest(fmt.Sprintf("vulnerability/%s/update", vulnerabilityID), body)
}
