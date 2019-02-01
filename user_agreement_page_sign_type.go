package alipay

// https://docs.alipay.com/pre-open/api_pre/alipay.user.agreement.page.sign
type AliPayUserAgreementPageSignIdentityParams struct {
	Username     string `json:"user_name"`
	CertNo       string `json:"cert_no"`
	IdentityHash string `json:"identity_hash"`
	SignUserId   string `json:"sign_user_id"`
}
type AliPayUserAgreementPageSign struct {
	AppAuthToken        string                                    `json:"-"` // 可选
	ReturnURL           string                                    `json:"-"`
	NotifyURL           string                                    `json:"notify_url"`
	PersonalProductCode string                                    `json:"personal_product_code"`
	ExternalLogonId     string                                    `json:"external_logon_id"` //用户在商户网站的登录账号，用于在签约页面展示，如果为空，则不展示
	IdentityParams      AliPayUserAgreementPageSignIdentityParams `json:"identity_params"`
}

func (this AliPayUserAgreementPageSign) APIName() string {
	return "alipay.user.agreement.page.sign"
}

func (this AliPayUserAgreementPageSign) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	m["personal_product_code"] = this.PersonalProductCode
	return m
}

func (this AliPayUserAgreementPageSign) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayUserAgreementPageSign) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayUserAgreementPageSignResponse struct {
	Body struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`

		//success
		ExternalAgreementNo string `json:"external_agreement_no"`
		PersonalProductCode string `json:"personal_product_code"`
		ValidTime           string `json:"valid_time"`
		SignScene           string `json:"sign_scene"`
		AgreementNo         string `json:"agreement_no"`
		ZmOpenId            string `json:"zm_open_id"`
		InvalidTime         string `json:"invalid_time"`
		SignTime            string `json:"sign_time"`
		AlipayUserId        string `json:"alipay_user_id"`
		Status              string `json:"status"`
		ForexEligible       string `json:"forex_eligible"`
		ExternalLogonId     string `json:"external_logon_id"`
		AlipayLogonId       string `json:"alipay_logon_id"`
	} `json:"alipay_user_agreement_page_sign_response"`
	ErrorResp struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	Sign string `json:"sign"`
}

func (this *AliPayUserAgreementPageSignResponse) IsSuccess() (bool, string) {
	if this.Body.Code == K_SUCCESS_CODE {
		return true, ""
	}
	return false, marshal(this.ErrorResp)
}
