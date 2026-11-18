package transport

import (
	"bufio"
	"context"
	"io"
	"net/http"
	"strings"

	norbixerr "github.com/norbix-code/sdk-go/v2/norbix/errors"
)

// SSEEvent is one parsed server-sent events frame: the "event:" field
// ("message" when the frame has none), the "data:" lines joined with "\n",
// and the last "id:" seen.
type SSEEvent struct {
	Event string
	Data  string
	ID    string
}

// Stream opens a server-sent events stream with GET and calls onEvent for
// every frame, in order, until the server closes the stream (nil), ctx ends
// (ctx.Err()) or onEvent returns false (nil). req.Path is the stream path
// (the gateway's is "/event-stream", without {version}); req.Body becomes
// the query string. Auth and the project / env / region headers are those of
// Send. Comment lines (":" heartbeats) are skipped.
//
// The client timeout does not apply: a stream stays open as long as ctx
// does. There is no retry and no reconnect — a refused subscription (401 /
// 403 / 404, e.g. AiChatChannelRefused) is returned as the gateway's error
// and would only be refused again.
func (t *Transport) Stream(ctx context.Context, req Request, onEvent func(SSEEvent) bool) error {
	if ctx == nil {
		ctx = context.Background()
	}
	base := t.apiBase()
	version := t.Cfg.APIVersion
	if req.Target == TargetHub {
		base = t.Cfg.BaseURLHub
		version = t.Cfg.HubVersion
	}
	urlStr, _, err := buildURLAndBody(base, req.Path, version, http.MethodGet, req.PathParams, req.Body)
	if err != nil {
		return err
	}
	headers, err := t.requestHeaders(req, "text/event-stream")
	if err != nil {
		return err
	}
	headers["Cache-Control"] = "no-cache"

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		return norbixerr.New(err.Error(), norbixerr.CodeNetwork).WithWrapped(err)
	}
	for k, v := range headers {
		httpReq.Header.Set(k, v)
	}
	resp, err := t.client.Do(httpReq)
	if err != nil {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		return norbixerr.New(err.Error(), norbixerr.CodeNetwork).WithWrapped(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		raw, _ := io.ReadAll(resp.Body)
		return norbixerr.FromResponse(raw, resp.StatusCode, resp.Header)
	}

	err = ReadSSE(resp.Body, onEvent)
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return err
}

// ReadSSE parses SSE framing from r — "event:", "data:" (several lines are
// joined with "\n"), "id:", ":" comments and the blank line that ends a frame
// — and calls onEvent per frame until r ends or onEvent returns false.
func ReadSSE(r io.Reader, onEvent func(SSEEvent) bool) error {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 64*1024), 4*1024*1024)
	event := "message"
	var data []string
	lastID := ""
	for sc.Scan() {
		line := strings.TrimSuffix(sc.Text(), "\r")
		if line == "" {
			if len(data) > 0 || event != "message" {
				if !onEvent(SSEEvent{Event: event, Data: strings.Join(data, "\n"), ID: lastID}) {
					return nil
				}
			}
			event, data = "message", nil
			continue
		}
		if strings.HasPrefix(line, ":") {
			continue
		}
		field, value, _ := strings.Cut(line, ":")
		value = strings.TrimPrefix(value, " ")
		switch field {
		case "event":
			event = value
		case "data":
			data = append(data, value)
		case "id":
			lastID = value
		}
	}
	return sc.Err()
}
