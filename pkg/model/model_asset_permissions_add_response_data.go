/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AssetPermissionsAddResponseData struct {
	FailNum    *int64              `json:"fail_num,omitempty"`
	SuccessNum *int64              `json:"success_num,omitempty"`
	FailReason *[]FailReasonStruct `json:"fail_reason,omitempty"`
}