package ocr

type GetOcrRequest struct {
	Images []string `json:"images"`
}

type GetOcrResponse struct {
	Msg     string `json:"msg"`
	Results [][]struct {
		Confidence float64 `json:"confidence"`
		Text       string  `json:"text"`
		TextRegion [][]int `json:"text_region"`
	} `json:"results"`
	Status string `json:"status"`
}
