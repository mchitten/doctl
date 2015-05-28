package main

import (
	"encoding/json"

	"github.com/codegangsta/cli"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

type tokenSource struct {
	AccessToken string
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: t.AccessToken,
	}, nil
}

func main() {
	app := cli.NewApp()
	app.Name = "docli"
	app.Usage = "DigitalOcean API CLI"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		tokenFlag(),
	}

	app.Commands = []cli.Command{
		accountCommands(),
		actionCommands(),
		domainCommands(),
		dropletCommands(),
		dropletActionCommands(),
		imageActionCommands(),
		imageCommands(),
		regionCommands(),
		sizeCommands(),
		sshKeyCommands(),
	}

	app.RunAndExitOnError()
}

func tokenFlag() cli.Flag {
	return cli.StringFlag{
		Name:   "token",
		Usage:  "DigitalOcean API V2 Token",
		EnvVar: "DO_TOKEN",
	}
}

func toJSON(item interface{}) (string, error) {
	b, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func newClient(c *cli.Context) *godo.Client {
	pat := c.GlobalString("token")
	tokenSource := &tokenSource{
		AccessToken: pat,
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	return godo.NewClient(oauthClient)
}