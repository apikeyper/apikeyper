package ratelimit

import (
	"apikeyper/internal/common"
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimitService struct {
	client *redis.Client
}

type RateLimitParams struct {
	Limit         int
	LimitPeriod   time.Duration // 1 hour for limitPeriodle
	CounterWindow time.Duration // 1 minute for example, 1/60 of the period
}

func New() *RateLimitService {

	config := common.GetRedisConfig()
	client := common.NewRedisClient(config, "ratelimit")

	return &RateLimitService{client}
}

func (r *RateLimitService) Increment(ctx context.Context, rateLimitParams RateLimitParams, key string, incr int) error {
	now := time.Now()
	timestamp := fmt.Sprint(now.Truncate(rateLimitParams.CounterWindow).Unix())

	val, err := r.client.HIncrBy(ctx, key, timestamp, int64(incr)).Result()
	if err != nil {
		return err
	}

	// check if current window has exceeded the limit
	if val >= int64(rateLimitParams.Limit) {
		// Otherwise, check if just this fixed window counter period is over
		slog.Info("Rate limit exceeded")
		return ErrRateLimitExceeded(
			0,
			rateLimitParams.Limit,
			rateLimitParams.LimitPeriod,
			now.Add(rateLimitParams.LimitPeriod),
		)
	}

	// create or move whole limit period window expiry
	r.client.Expire(ctx, key, rateLimitParams.LimitPeriod)

	// Get all the bucket values and sum them.
	vals, err := r.client.HGetAll(ctx, key).Result()
	if err != nil {
		return err
	}

	// The time to start summing from, any buckets before this are ignored.
	threshold := fmt.Sprint(now.Add(-rateLimitParams.LimitPeriod).Unix())

	// NOTE: This sums ALL the values in the hash, for more information, see the
	// "Practical Considerations" section of the associated Figma blog post.
	total := 0
	for k, v := range vals {
		if k > threshold {
			i, _ := strconv.Atoi(v)
			total += i
		} else {
			// Clear the old hash keys
			r.client.HDel(ctx, key, k)
		}
	}

	if total >= int(rateLimitParams.Limit) {
		return ErrRateLimitExceeded(
			0,
			rateLimitParams.Limit,
			rateLimitParams.LimitPeriod,
			now.Add(rateLimitParams.LimitPeriod),
		)
	}

	return nil
}
