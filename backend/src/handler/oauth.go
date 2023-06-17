package handler

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	traqoauth2 "github.com/ras0q/traq-oauth2"
	"github.com/traP-jp/h23s_01/backend/src/config"
)

var conf = config.GetTraqClientConf()

func authorizeHandler(c echo.Context) error {
	codeVerifier, err := traqoauth2.GenerateCodeVerifier()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to generate code verifier: %v", err))
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to get seeeion: %v", err))
	}

	sess.Values["code_verifier"] = codeVerifier
	sess.Options.SameSite = http.SameSiteNoneMode

	postmanRegexp := regexp.MustCompile(`^Postman`)
	if !postmanRegexp.MatchString(c.Request().UserAgent()) { //デバッグでPostmanを使う際にSecureがtrueだとちゃんとCookieを保存できない
		sess.Options.Secure = true
	}

	sess.Save(c.Request(), c.Response())

	codeChallengeMethod, ok := traqoauth2.CodeChallengeMethodFromStr(c.QueryParam("method"))
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid code challenge method: %v", c.QueryParam("method")))
	}

	codeChallenge, err := traqoauth2.GenerateCodeChallenge(codeVerifier, codeChallengeMethod)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to generate code challenge: %v", err))
	}

	authCodeURL := conf.AuthCodeURL(
		c.QueryParam("state"),
		traqoauth2.WithCodeChallenge(codeChallenge),
		traqoauth2.WithCodeChallengeMethod(codeChallengeMethod),
	)

	c.Response().Header().Set("Location", authCodeURL)

	return echo.NewHTTPError(http.StatusSeeOther, fmt.Sprintf("redirect to %s", authCodeURL))
}

func callbackHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to get session: %v", err))
	}

	codeVerifier, ok := sess.Values["code_verifier"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	code := c.QueryParam("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "no code")
	}

	ctx := c.Request().Context()
	token, err := conf.Exchange(ctx, code, traqoauth2.WithCodeVerifier(codeVerifier))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to exchange code into token: %v", err))
	}

	sess.Values["access_token"] = token.AccessToken
	sess.Values["expires_at"] = token.Expiry

	sess.Save(c.Request(), c.Response())

	return c.String(http.StatusOK, "ok")
}

type tokenStatus int

const (
	valid tokenStatus = iota
	expired
	noToken
)

func getToken(c echo.Context) (string, error) {
	sess, err := session.Get("session", c)
	if err != nil {
		return "", err
	}
	token, ok := sess.Values["access_token"].(string)
	if !ok {
		return "", errors.New("no access token")
	}
	return token, nil
}
