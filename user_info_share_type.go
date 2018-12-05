package alipay

// https://docs.open.alipay.com/api_2/alipay.user.info.share
type AliPayUserInfoShare struct {
	AppAuthToken string `json:"-"` // 可选
	AuthToken    string `json:"auth_token"`
}

func (this AliPayUserInfoShare) APIName() string {
	return "alipay.user.info.share"
}

func (this AliPayUserInfoShare) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayUserInfoShare) ExtJSONParamName() string {
	return ""
}

func (this AliPayUserInfoShare) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayUserInfoShareResponse struct {
	Body struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
	} `json:"alipay_system_oauth_token_response"`
	Sign string `json:"sign"`
}
