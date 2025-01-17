/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LeadCluesGetRequest struct {
	AccountId             *int64             `json:"account_id,omitempty"`
	TimeType              TimeType           `json:"time_type,omitempty"`
	TimeRange             *TimeRange         `json:"time_range,omitempty"`
	Filtering             *[]FilteringStruct `json:"filtering,omitempty"`
	Page                  *int64             `json:"page,omitempty"`
	PageSize              *int64             `json:"page_size,omitempty"`
	LastSearchAfterValues *[]string          `json:"last_search_after_values,omitempty"`
}
