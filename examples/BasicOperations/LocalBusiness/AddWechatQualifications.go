/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type WechatQualificationsAddExample struct {
	TAds                             *ads.SDKClient
	AccessToken                      string
	AccountId                        int64
	WechatQualificationType          string
	WechatQualificationFile          *os.File
	WechatQualificationFileSignature string
}

func (e *WechatQualificationsAddExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = 789
	e.WechatQualificationType = "wechatQualificationType_example"
	file, err := os.Open("YOUR FILE PATH")
	if err != nil {
		e.WechatQualificationFile = file
	}
	e.WechatQualificationFileSignature = "wechatQualificationFileSignature_example"
}

func (e *WechatQualificationsAddExample) RunExample() (model.WechatQualificationsAddResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.WechatQualifications().Add(ctx, e.AccountId, e.WechatQualificationType, e.WechatQualificationFile, e.WechatQualificationFileSignature)
}

func main() {
	e := &WechatQualificationsAddExample{}
	e.Init()
	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Headers:", headers)
}