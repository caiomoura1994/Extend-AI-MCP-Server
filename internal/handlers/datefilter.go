package handlers

import (
	"context"
	"fmt"
	"time"
)

const (
	maxAutoPages    = 50
	autoPageSize    = 100
)

// DateFilterParams holds optional date range boundaries for client-side filtering.
type DateFilterParams struct {
	CreatedAfter  *time.Time
	CreatedBefore *time.Time
}

// HasDateFilter returns true if at least one date boundary is set.
func (d *DateFilterParams) HasDateFilter() bool {
	return d.CreatedAfter != nil || d.CreatedBefore != nil
}

// PageFetcher fetches a single page of results given a pagination token (nil for first page).
// Returns the items, the next page token (nil if no more pages), and any error.
type PageFetcher[T any] func(ctx context.Context, token *string) (items []T, nextToken *string, err error)

// TimeExtractor returns the createdAt time from an item.
type TimeExtractor[T any] func(item T) time.Time

// FetchWithDateFilter auto-paginates through an API endpoint, collecting items whose
// createdAt falls within the date range. It assumes results are sorted by createdAt desc
// (newest first) and stops early when it finds items older than CreatedAfter.
func FetchWithDateFilter[T any](
	ctx context.Context,
	params DateFilterParams,
	fetch PageFetcher[T],
	getTime TimeExtractor[T],
) ([]T, error) {
	var collected []T
	var nextToken *string

	for page := 0; page < maxAutoPages; page++ {
		items, token, err := fetch(ctx, nextToken)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch page %d: %w", page+1, err)
		}

		allOlderThanRange := len(items) > 0
		for _, item := range items {
			t := getTime(item)

			if params.CreatedBefore != nil && t.After(*params.CreatedBefore) {
				// Item is newer than the range upper bound -- skip but keep going
				allOlderThanRange = false
				continue
			}
			if params.CreatedAfter != nil && t.Before(*params.CreatedAfter) {
				// Item is older than the range lower bound -- skip
				// Since results are desc-sorted, all remaining items are also older
				continue
			}

			// Item is within range
			allOlderThanRange = false
			collected = append(collected, item)
		}

		// If every item on this page was older than created_after, we can stop
		if allOlderThanRange {
			break
		}

		// Check if the oldest item on this page is already past our range
		if len(items) > 0 && params.CreatedAfter != nil {
			oldest := getTime(items[len(items)-1])
			if oldest.Before(*params.CreatedAfter) {
				break
			}
		}

		nextToken = token
		if nextToken == nil {
			break
		}
	}

	return collected, nil
}

// ParseOptionalTime parses an RFC 3339 string pointer into a time.Time pointer.
// Returns nil if the input is nil or empty.
func ParseOptionalTime(s *string) (*time.Time, error) {
	if s == nil || *s == "" {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, fmt.Errorf("invalid date format (expected RFC 3339, e.g. 2026-02-17T00:00:00Z): %w", err)
	}
	return &t, nil
}
