//  Method to print report from stress-test with total duration, number of requests
//  and status codes that were received.

package report

import (
	"fmt"
	"time"

	"github.com/thayline/stress-test/internal/result"
)

func Print(results []result.Result, duration time.Duration) {
	statusCount := make(map[int]int)
	success := 0
	failures := 0

	for _, r := range results {
		if r.Error != nil {
			fmt.Printf("[Failures %v] Result %v has an error: %v\n", failures, r, r.Error)
			failures++
			continue
		}

		statusCount[r.StatusCode]++
		if r.StatusCode == 200 {
			success++
		}
	}

	fmt.Println(":===== Stress Test Report =====:")
	fmt.Printf("Total time: %v\n", duration)
	fmt.Printf("Total requests: %d\n", len(results))
	fmt.Printf("Status 200: %d\n", success)
	fmt.Println("Other status codes:")
	for code, count := range statusCount {
		if code != 200 {
			fmt.Printf("  %d -> %d\n", code, count)
		}
	}
}
