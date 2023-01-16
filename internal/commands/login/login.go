package login

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
)

var Command = cli.Command{
	Name:    "login",
	Aliases: []string{},
	Usage:   "",
	Flags:   []cli.Flag{},
	Action: func(context *cli.Context) error {
		payload := strings.NewReader(fmt.Sprintf("client_id=%s&scope=scope", os.Getenv("AUTH0_CLIENT_ID")))
		req, err := http.Post(fmt.Sprintf("https://%s/oauth/device/code", os.Getenv("AUTH0_DOMAIN")), "application/x-www-form-urlencoded", payload)

		if err != nil {
			pterm.Error.Println(fmt.Errorf("auth %w", err))
		}

		defer req.Body.Close()
		rawBody, _ := io.ReadAll(req.Body)
		var body models.AuthTokens
		_ = json.Unmarshal(rawBody, &body)
		openBrowser(body.VerificationURIComplete)

		return nil
	},
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		pterm.Error.Println(fmt.Errorf("%w", err))
	}
}
