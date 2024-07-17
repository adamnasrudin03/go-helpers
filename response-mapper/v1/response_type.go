package v1

// ResponseDefault represents the default structure for response handling.
type ResponseDefault struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Meta struct holds metadata information.
type Meta struct {
	Page         int `json:"page,omitempty"`
	Limit        int `json:"limit,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
	TotalPages   int `json:"total_pages,omitempty"`
}

// Pagination represents the structure for paginating data.
type Pagination struct {
	Meta Meta
	Data interface{}
}

// MultiLanguages represents a structure for multi-language support.
type MultiLanguages struct {
	ID string `json:"id"`
	EN string `json:"en"`
}

// Error returns the error message in the preferred language.
func (e *MultiLanguages) Error() string {
	if e.EN != "" {
		return e.EN
	} else if e.ID != "" {
		return e.ID
	}
	return "something went wrong"
}

// NewResponseMultiLang creates a new MultiLanguages instance.
func NewResponseMultiLang(languages MultiLanguages) *MultiLanguages {
	return &languages
}
