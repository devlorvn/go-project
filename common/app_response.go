package common

type successResponse struct {
	Data     interface{} `json:"data"`
	Metadata interface{} `json:"metadata,omitempty"`
}

func NewSuccessResponse(data, metadata interface{}) *successResponse {
	return &successResponse{Data: data, Metadata: metadata}
}

func SimpleSuccessResponse(data interface{}) *successResponse {
	return NewSuccessResponse(data, nil)
}
