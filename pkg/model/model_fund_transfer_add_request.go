/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type FundTransferAddRequest struct {
	Amount          *int64         `json:"amount,omitempty"`
	TransferType    *string        `json:"transfer_type,omitempty"`
	ExternalBillNo  *string        `json:"external_bill_no,omitempty"`
	Memo            *string        `json:"memo,omitempty"`
	TransferTryBest *int64         `json:"transfer_try_best,omitempty"`
	AccountId       *int64         `json:"account_id,omitempty"`
	FundType        AccountTypeMap `json:"fund_type,omitempty"`
}
