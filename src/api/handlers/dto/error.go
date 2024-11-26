package dto

type ErrorMessage struct {
	StatusCode    int            `json:"status_code,omitempty"`
	Message       string         `json:"message"`
	InvalidFields []InvalidField `json:"invalid_fields,omitempty"`
	isInternal    bool
}

type InvalidField struct {
	FieldName   string `json:"field_name"`
	Description string `json:"description"`
}

func (instance *ErrorMessage) IsInternal() bool {
	return instance.isInternal
}
