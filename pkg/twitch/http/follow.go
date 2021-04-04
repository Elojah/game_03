package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/twitch"
)

type followFilter twitch.FollowFilter

func (f followFilter) set(req *http.Request) {
	if f.FromID != nil {
		req.Header.Set("from_id", *f.FromID)
	}

	if f.ToID != nil {
		req.Header.Set("to_id", *f.ToID)
	}
}

func (c Client) GetFollows(
	ctx context.Context,
	a twitch.Auth,
	f twitch.FollowFilter,
	callback func(twitch.Follow) error,
) error {
	b, err := url.Parse(twitchURL)
	if err != nil {
		return err
	}

	r, err := url.Parse("/helix/users/follows")
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, b.ResolveReference(r).String(), nil)
	if err != nil {
		return err
	}

	auth(a).set(req)

	followFilter(f).set(req)

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return errors.ErrInvalidStatus{Status: resp.Status}
	}

	var result struct {
		Total      int
		Data       []twitch.Follow
		Pagination struct {
			Cursor string
		}
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	for _, fo := range result.Data {
		if err := callback(fo); err != nil {
			return err
		}
	}

	return nil
}
