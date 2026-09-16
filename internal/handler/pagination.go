package handler

import (
	"net/http"
	"strconv"
)

// PaginationParams represents the query parameters for pagination in API requests. It includes a cursor for the starting point of the pagination and a limit for the number of items to be returned in the response.
type PaginationParams struct {
	Cursor string
	Limit  int
}

const (
	defaultLimit = 20
	maxLimit     = 100
)

func ParsePagination(request *http.Request) PaginationParams {
	cursor := request.URL.Query().Get("cursor")
	limit := defaultLimit

	if limitQuery := request.URL.Query().Get("limit"); limitQuery != "" {
		if parsed, err := strconv.Atoi(limitQuery); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	return PaginationParams{
		Cursor: cursor,
		Limit:  limit,
	}
}

type PaginatedResponse[T any] struct {
	Data       []T    `json:"data"`
	NextCursor string `json:"next_cursor,omitempty"`
	HasMore    bool   `json:"has_more"`
}
