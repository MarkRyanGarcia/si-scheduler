package drive

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
)

// GetClient retrieves the OAuth2 token from file or web flow.
// Returns the token and an error (if any).
func GetClient(config *oauth2.Config) (*oauth2.Token, error) {
	tok, err := getTokenFromFile()
	if err != nil {
		// If no token file, try web flow
		tok, err = getTokenFromWeb(config)
		if err != nil {
			return nil, fmt.Errorf("failed to get token from web: %w", err)
		}
		if err := saveToken(tok); err != nil {
			return nil, fmt.Errorf("failed to save token: %w", err)
		}
	}
	return tok, nil
}

func getTokenFromFile() (*oauth2.Token, error) {
	path := filepath.Join("token.json")
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tok := &oauth2.Token{}
	if err := json.NewDecoder(f).Decode(tok); err != nil {
		return nil, err
	}
	return tok, nil
}

func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser and enter the code:\n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, fmt.Errorf("unable to read authorization code: %w", err)
	}
	tok, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve token: %w", err)
	}
	return tok, nil
}

func saveToken(token *oauth2.Token) error {
	path := filepath.Join("token.json")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(token); err != nil {
		return err
	}
	fmt.Printf("Saved OAuth token to %s\n", path)
	return nil
}
