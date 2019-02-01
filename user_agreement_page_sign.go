package alipay

// ConfirmUserAgreementPageSign https://docs.alipay.com/pre-open/api_pre/alipay.user.agreement.page.sign
func (this *AliPay) ConfirmUserAgreementPageSign(param AliPayUserAgreementPageSign) (results *AliPayUserAgreementPageSignResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}
