package client

import (
	"time"

	"github.com/Rican7/retry/strategy"
)

// Config holds various configuration parameters for a dqlite client.
type Config struct {
	AttemptTimeout  time.Duration       // Timeout for each individual connection attempt.
	RetryStrategies []strategy.Strategy // Strategies used for retrying to connect to a leader.
}
