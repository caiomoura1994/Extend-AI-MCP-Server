package dto

// ParserRunMetrics contains metrics about the parsing process
type ParserRunMetrics struct {
	ProcessingTimeMs float64 `json:"processingTimeMs"`
	PageCount        float64 `json:"pageCount"`
}

// ParserRunChunkMetadataPageRange represents the page range of a chunk
type ParserRunChunkMetadataPageRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// ParserRunChunkMetadata contains metadata about a chunk
type ParserRunChunkMetadata struct {
	PageRange ParserRunChunkMetadataPageRange `json:"pageRange"`
}

// ParserRunChunk represents a parsed chunk from a document
type ParserRunChunk struct {
	Object   string                   `json:"object"`
	Type     string                   `json:"type"` // page, document, section
	Content  string                   `json:"content"`
	Metadata ParserRunChunkMetadata   `json:"metadata"`
	Blocks   []map[string]interface{} `json:"blocks,omitempty"`
}

// ParserRun represents a parser run result (GET /parser_runs/{id})
type ParserRun struct {
	Object        string                 `json:"object"` // parser_run or parser_run_status
	ID            string                 `json:"id"`
	FileID        string                 `json:"fileId,omitempty"`
	Status        string                 `json:"status"` // PROCESSING, PROCESSED, FAILED
	FailureReason string                 `json:"failureReason,omitempty"`
	Chunks        []ParserRunChunk       `json:"chunks,omitempty"`
	Metrics       *ParserRunMetrics      `json:"metrics,omitempty"`
	Config        map[string]interface{} `json:"config,omitempty"`
	Usage         map[string]interface{} `json:"usage,omitempty"`
}

// GetParserRunResponse wraps the get parser run API response
type GetParserRunResponse struct {
	Success   bool      `json:"success"`
	Warning   string    `json:"warning,omitempty"`
	ParserRun ParserRun `json:"parserRun"`
}
