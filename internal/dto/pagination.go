package dto

import (
	"fmt"
	"net/url"
)

// PaginationParams represents common query parameters for paginated list endpoints
type PaginationParams struct {
	MaxPageSize   *int    `json:"max_page_size,omitempty"`
	NextPageToken *string `json:"next_page_token,omitempty"`
	SortBy        *string `json:"sort_by,omitempty"`  // "updatedAt" or "createdAt"
	SortDir       *string `json:"sort_dir,omitempty"` // "asc" or "desc"
}

// QueryString builds URL query parameters from the pagination params
func (p *PaginationParams) QueryString() string {
	if p == nil {
		return ""
	}

	params := url.Values{}

	if p.MaxPageSize != nil {
		params.Set("maxPageSize", fmt.Sprintf("%d", *p.MaxPageSize))
	}
	if p.NextPageToken != nil && *p.NextPageToken != "" {
		params.Set("nextPageToken", *p.NextPageToken)
	}
	if p.SortBy != nil && *p.SortBy != "" {
		params.Set("sortBy", *p.SortBy)
	}
	if p.SortDir != nil && *p.SortDir != "" {
		params.Set("sortDir", *p.SortDir)
	}

	encoded := params.Encode()
	if encoded == "" {
		return ""
	}
	return "?" + encoded
}

// PaginationInfo represents pagination metadata returned in list responses
type PaginationInfo struct {
	NextPageToken *string `json:"next_page_token,omitempty"`
}
