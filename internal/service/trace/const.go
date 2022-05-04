package trace

type SendCMSRequest struct {
	Phone      string `json:"p090983d"` // 手机号
	TimeString string `json:"s324234m"` // 时间字符串 "2022-05-05 00:01:58"
	TimeMD5    string `json:"s234234s"` // 时间字符串MD5
	UUID       string `json:"q435434f"` // uuid
}

//{"status":"2","code":"01","errorDesc":"非法访问","result":null,"queryId":null}
//{"status":"1","code":"00","errorDesc":"请求成功","result":"短信发送中,请注意查收","queryId":"0435c23b172cb1471f18e27185ae1611"}
type SendCMSResponse struct {
	Status    string `json:"status"`
	Code      string `json:"code"`
	ErrorDesc string `json:"errorDesc"`
	Result    string `json:"result"`
	QueryId   string `json:"queryId"` // 就是请求验证码时候自己生成的uuid
}

// var _0x18d3 = ['shift', 'createElement', 'text/css', 'onload', 'name', 'head', '__esModule', '56d7', 'powered-by-title', 'caict', 'e260', 'div', 'form-list', 'loginForm', 'phonebtn', '100%', 'veryCode', 'composing', 'form-item', '0.02rem', 'ckValue', 'target', 'v-show', 'primary', 'mt-button', 'show', 'none', '#007aff', 'tel:10010', 'detail2', '1rem', '0.2rem', '26%', 'goUrl', '使用\x20“通信行程卡”\x20小程序或手机APP，不用每次输验证码', '$createElement', '25cd', 'flex', 'code', 'location', 'href', 'checkPhoneValue', 'phone', 'uuid', 'loginkey', 'OcpqZSOIZOxr0', 'getItem', '请输入正确的手机号码！', '获取验证码', 'dateFormat', '$xhrlogin', 'getTime', 'getSeconds', 'timeStap', 'rpage', 'body', 'headers', 'data', 'default', 'prototype', 'Swipe', 'component', 'console'];
type GetTraceInfoRequest struct {
	Phone   string `json:"p234324"` // 手机号
	QueryId string `json:"q363209"` // 请求验证码返回的queryId，就是请求验证码时候自己生成的uuid
	Token   string `json:"y892342"` // 验证码6位
	Code    string `json:"c098465"` // url的code参数，一般为空
	Crypto  string `json:"s456hr8"` // MD5(_0x18d3[46]+phone)， _0x18d3[46]=OcpqZSOIZOxr0
}
type GetTraceInfoResponse struct {
	Status    string `json:"status"`
	Code      string `json:"code"`
	ErrorDesc string `json:"errorDesc"`
	Result    struct {
		Color   string `json:"color"`
		Phone   string `json:"phone"`
		Time    string `json:"time"`
		Message string `json:"message"`
	} `json:"result"`
	QueryId string `json:"queryId"`
}
