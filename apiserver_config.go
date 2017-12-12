package test

import "github.com/asaskevich/govalidator"

// APIServerConfig is a struct holding data to configure the API Server process
type APIServerConfig struct {
	APIServerURL string `valid:"url"`
}

// Validate checks that the config contains only valid URLs
func (c *APIServerConfig) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	return err
}
