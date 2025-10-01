package station

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"

	"weather-service/internal/config"
	"weather-service/internal/server"
)

type City string

func (c City) Name() string               { return string(c) }
func (City) Init(_ context.Context) error { return nil }
func (City) Stop() error                  { return nil }

func (c City) Run(ctx context.Context) error {
	time.Sleep(time.Duration(rand.IntN(1000)) * time.Millisecond)
	ticker := time.NewTicker(time.Second)

	for {
		var ts time.Time
		select {
		case <-ctx.Done():
			return nil
		case ts = <-ticker.C:
		}

		temp := rand.IntN(35)
		fmt.Printf("[%s] %s: %d\n", ts.UTC().Format(time.DateTime), c, temp)

		msg, err := json.Marshal(server.TempMessage{
			Temp: temp,
			Time: ts,
		})
		if err != nil {
			return fmt.Errorf("failed to marshall into json: %w", err)
		}

		req, err := http.NewRequestWithContext(
			ctx, http.MethodPost,
			fmt.Sprintf("http://localhost:%d/cities/%s", config.C.APIPort, c),
			bytes.NewReader(msg),
		)
		if err != nil {
			return fmt.Errorf("failed to build request: %w", err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("failed to submit request: %w", err)
		} else if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to submit request: %v", resp)
		}

		err = resp.Body.Close()
		if err != nil {
			return fmt.Errorf("failed to close body: %w", err)
		}
	}
}
