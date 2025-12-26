package executor

import (
	"net/http"
	"time"

	"github.com/thayline/stress-test/internal/result"
)

func Run(url string, rotines int, concurrency int) []result.Result {
	jobs := make(chan int, rotines)
	results := make(chan result.Result, rotines)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for w := 0; w < concurrency; w++ {
		go func() {
			for range jobs {
				start := time.Now()
				resp, err := client.Get(url)

				res := result.Result{}
				res.DurationMs = time.Since(start).Milliseconds()

				if err != nil {
					res.Error = err
				} else {
					res.StatusCode = resp.StatusCode
					resp.Body.Close()
				}

				results <- res
			}
		}()
	}

	for i := 0; i < rotines; i++ {
		jobs <- i
	}
	close(jobs)

	var final []result.Result
	for i := 0; i < rotines; i++ {
		final = append(final, <-results)
	}

	return final
}
