package common

// JsonResponse Json 返回体结构
type JsonResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Count   int64       `json:"count"`
	Data    interface{} `json:"data,omitempty"`
}

func (res *JsonResponse) Clear() *JsonResponse {
	res.Message = "success"
	res.Code = 20000
	//res.Count = 0
	res.Data = nil
	return res
}
