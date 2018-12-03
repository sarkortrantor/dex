package daga

import (
	"errors"
	"fmt"
	"github.com/dexidp/dex/connector"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

// dagaConnector is a connector that uses a daga cothority to authenticate users
type dagaConnector struct {
	Logger   logrus.FieldLogger
}

// NewDAGAConnector returns a dex "callback connector" (not really now..). that use DAGA to authenticate users
func NewDAGAConnector(logger logrus.FieldLogger) connector.Connector {
	return &dagaConnector{
		Logger: logger,
	}
}

// LoginURL returns the URL to redirect the user to login with.
func (dc *dagaConnector) LoginURL(s connector.Scopes, _, state string) (string, error) {
	// TODO enhancement, later speak OAuth with daga_auth located elsewhere (I think this shouldn't be too difficult)
	//  => need to setup callbackurls in connector config and at remote and follow same flow/design/conventions as existing OAuth connectors
	// TODO get rid of the magic string and load that information from connector config(.yaml) => becomes kind of generic OAuth connector => later refactor existing connectors to use that as my little useless contribution (if makes sense)
	u, err := url.Parse("/daga_auth")
	if err != nil {
		return "", fmt.Errorf("failed to parse URL %q: %v", "/daga_auth", err)
	}
	v := u.Query()
	// pass state in url because remote daga login endpoint will need to give it to us again for the handlecallback/finalize stage
	v.Set("state", state)
	u.RawQuery = v.Encode()
	return u.String(), nil
}

// HandleCallback parses the request and returns the user's identity
func (dc *dagaConnector) HandleCallback(s connector.Scopes, r *http.Request) (connector.Identity, error) {
	// extract tag from dex request state (don't use user agent for that, on login postback, store tag in request state, then retrieve it here)
	// TODO too complex, can already log user in at previous stage, but this is too use the "dex callback connector flow"
	// FIXME if I don't use OAuth with remote daga idp and keep using handbacked dex solution => create new type of connector and finalize login on daga_auth postback
	// FIXME or keep callback connector and short circuit the handlecallback stage, redirect to finalize on daga_auth postback
 	var tag string = "implement me"
	// build identity
	identity := connector.Identity{
		UserID: tag,
	}
	return identity, errors.New("unimplemented, not needed for now")
}

// Config holds configuration options for daga logins.
type Config struct{}

// Open returns the daga authentication strategy.
func (c *Config) Open(id string, logger logrus.FieldLogger) (connector.Connector, error) {
	return NewDAGAConnector(logger), nil
}