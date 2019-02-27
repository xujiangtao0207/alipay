package alipay

// UserAgreementPageSign https://docs.alipay.com/pre-open/api_pre/alipay.user.agreement.page.sign
func (this *AliPay) UserAgreementPageSign(param AliPayUserAgreementPageSign) (results *AliPayUserAgreementPageSignResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// UserAgreementQuery https://docs.alipay.com/pre-open/api_pre/alipay.user.agreement.query
func (this *AliPay) UserAgreementQuery(param AliPayUserAgreementQuery) (results *AliPayUserAgreementQueryResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}
