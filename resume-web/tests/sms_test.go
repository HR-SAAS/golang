package main

import (
	"fmt"
	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"testing"
)

func Test_sms(t *testing.T) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-qingdao", "LTAI5tAAjLhEXmxpigb7BFYz", "XwpYel3yJ98sx8D0zcoRl9SUiOK2cM")
	/* use STS Token
	client, err := dysmsapi.NewClientWithStsToken("cn-qingdao", "<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = "15070055362"   //接收短信的手机号码
	request.SignName = "yblog"             //短信签名名称
	request.TemplateCode = "SMS_148865003" //短信模板ID
	request.TemplateParam = "{\"code\":\"1111\"}"
	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
