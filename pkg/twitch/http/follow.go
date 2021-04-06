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

	if f.After != nil {
		req.Header.Set("after", *f.After)
	}

	if f.First != nil {
		req.Header.Set("first", *f.First)
	}
}

func (c Client) GetFollows(
	ctx context.Context,
	a twitch.Auth,
	f twitch.FollowFilter,
	callback func(twitch.Follow) error,
) (twitch.Cursor, error) {
	b, err := url.Parse(twitchURL)
	if err != nil {
		return twitch.Cursor{}, err
	}

	r, err := url.Parse("/helix/users/follows")
	if err != nil {
		return twitch.Cursor{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, b.ResolveReference(r).String(), nil)
	if err != nil {
		return twitch.Cursor{}, err
	}

	auth(a).set(req)

	followFilter(f).set(req)

	resp, err := c.Do(req)
	if err != nil {
		return twitch.Cursor{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return twitch.Cursor{}, errors.ErrInvalidStatus{Status: resp.Status}
	}

	var result struct {
		Total      uint
		Data       []twitch.Follow
		Pagination struct {
			Cursor string
		}
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return twitch.Cursor{}, err
	}

	for _, fo := range result.Data {
		if err := callback(fo); err != nil {
			return twitch.Cursor{}, err
		}
	}

	return twitch.Cursor{
		Total:  result.Total,
		Cursor: result.Pagination.Cursor,
	}, nil
}
