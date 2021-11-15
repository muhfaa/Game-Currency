package response

type ConversionResponseData struct {
	Result int `json:"result"`
}

func NewResponse(data int) ConversionResponseData {
	return ConversionResponseData{
		Result: data,
	}
}
