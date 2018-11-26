package daga

import (
	"fmt"
	"github.com/dexidp/dex/connector"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type dagaConnector struct {}


// NewDAGAConnector returns a dex callback connector. that use DAGA to authenticate users
func NewDAGAConnector() connector.Connector {

	// TODO
	return connector.CallbackConnector(nil)
}


// LoginURL returns the URL to redirect the user to login with.
func (dc *dagaConnector) LoginURL(s connector.Scopes, callbackURL, state string) (string, error) {
	u, err := url.Parse(callbackURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse callbackURL %q: %v", callbackURL, err)
	}
	v := u.Query()
	v.Set("state", state)
	u.RawQuery = v.Encode()
	return u.String(), nil
}

// HandleCallback parses the request and returns the user's identity
func (dc *dagaConnector) HandleCallback(s connector.Scopes, r *http.Request) (connector.Identity, error) {

	// TODO
	panic("implement me")
	return connector.Identity{}, nil
}

// Config holds configuration options for daga logins.
type Config struct{}

func (c *Config) Open(id string, logger logrus.FieldLogger) (connector.Connector, error) {
	panic("implement me")
	return connector.Connector(nil), nil
}