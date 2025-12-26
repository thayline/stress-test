package result

type Result struct {
	StatusCode int
	Error      error
	DurationMs int64
}
