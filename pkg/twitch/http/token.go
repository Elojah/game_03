package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/twitch"
	"golang.org/x/oauth2"
)

const (
	twitchIDURL = "https://id.twitch.tv"
)

func (c Client) GetToken(ctx context.Context, code string, cfg oauth2.Config) (twitch.Token, error) {
	b, err := url.Parse(twitchIDURL)
	if err != nil {
		return twitch.Token{}, err
	}

	r, err := url.Parse("/oauth2/token")
	if err != nil {
		return twitch.Token{}, err
	}

	q := r.Query()
	q.Set("client_id", cfg.ClientID)
	q.Set("client_secret", cfg.ClientSecret)
	q.Set("code", code)
	q.Set("grant_type", "authorization_code")
	q.Set("redirect_uri", cfg.RedirectURL)
	r.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, b.ResolveReference(r).String(), nil)
	if err != nil {
		return twitch.Token{}, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return twitch.Token{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		raw, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("err:", string(raw))
		return twitch.Token{}, errors.ErrInvalidStatus{Status: resp.Status}
	}

	var result twitch.Token

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return twitch.Token{}, err
	}

	return result, nil
}
