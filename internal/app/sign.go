package app

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"git.sr.ht/~jamesponddotco/bunnysign/internal/token"
	"github.com/urfave/cli/v2"
)

// SignAction is the main action for the application.
func SignAction(ctx *cli.Context) error {
	var (
		uri      = ctx.String("url")
		key      = ctx.String("key")
		duration = ctx.Duration("expiration")
	)

	u, err := url.Parse(uri)
	if err != nil {
		return fmt.Errorf("failed to parse URL to sign; please check the URL and try again: %w", err)
	}

	var (
		expiration     = time.Now().Unix() + int64(duration.Seconds())
		baseString     = key + u.Path + strconv.FormatInt(expiration, 10)
		generatedToken = token.Generate(baseString)
		finalURI       = fmt.Sprintf("%s://%s%s?token=%s&expires=%d", u.Scheme, u.Host, u.Path, generatedToken, expiration)
	)

	fmt.Fprintf(ctx.App.Writer, "%s\n", finalURI)

	return nil
}
