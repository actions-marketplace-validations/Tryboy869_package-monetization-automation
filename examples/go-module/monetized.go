package monetized

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "regexp"
)

type Client struct {
    LicenseKey  string
    Tier        string
    APIEndpoint string
}

func New(licenseKey, tier string) (*Client, error) {
    if licenseKey == "" && tier != "free" {
        return nil, fmt.Errorf("license key required for paid tiers")
    }
    
    return &Client{
        LicenseKey:  licenseKey,
        Tier:        tier,
        APIEndpoint: "https://api.yourpackage.com",
    }, nil
}

func (c *Client) ValidateLicense() bool {
    if c.Tier == "free" {
        return true
    }
    
    patterns := map[string]string{
        "basic":      "^bsc_[a-f0-9]{16}$",
        "pro":        "^pro_[a-f0-9]{16}$", 
        "enterprise": "^ent_[a-f0-9]{16}$",
    }
    
    pattern, exists := patterns[c.Tier]
    if !exists {
        return false
    }
    
    matched, _ := regexp.MatchString(pattern, c.LicenseKey)
    return matched
}

func (c *Client) APICall(data interface{}) (map[string]interface{}, error) {
    if !c.ValidateLicense() {
        return nil, fmt.Errorf("invalid license key")
    }
    
    payload := map[string]interface{}{
        "data":    data,
        "license": c.LicenseKey,
        "tier":    c.Tier,
    }
    
    jsonData, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }
    
    resp, err := http.Post(c.APIEndpoint+"/process", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }
    
    return result, nil
}