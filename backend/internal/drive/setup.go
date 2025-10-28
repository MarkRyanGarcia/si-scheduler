package drive

import (
	"context"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func Setup() (*sheets.Service, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}
	config, err := google.ConfigFromJSON(b, sheets.DriveReadonlyScope)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	token, err := GetClient(config)
	if err != nil {
		return nil, err
	}
	client := config.Client(ctx, token)
	return sheets.NewService(ctx, option.WithHTTPClient(client))
}
