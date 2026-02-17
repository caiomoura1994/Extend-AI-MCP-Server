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

// QueryValues returns url.Values for the pagination params
func (p *PaginationParams) QueryValues() url.Values {
	params := url.Values{}
	if p == nil {
		return params
	}

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

	return params
}

// QueryString builds URL query parameters from the pagination params
func (p *PaginationParams) QueryString() string {
	return BuildQueryString(p.QueryValues())
}

// BuildQueryString merges multiple url.Values into a single query string
func BuildQueryString(parts ...url.Values) string {
	merged := url.Values{}
	for _, part := range parts {
		for key, values := range part {
			for _, v := range values {
				merged.Set(key, v)
			}
		}
	}

	encoded := merged.Encode()
	if encoded == "" {
		return ""
	}
	return "?" + encoded
}

// PaginationInfo represents pagination metadata returned in list responses
type PaginationInfo struct {
	NextPageToken *string `json:"next_page_token,omitempty"`
}
