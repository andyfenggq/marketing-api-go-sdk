/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 页面元素配置
type PageSpecsListStruct struct {
	BgColor              *string                       `json:"bg_color,omitempty"`
	PageElementsSpecList *[]PageElementsSpecListStruct `json:"page_elements_spec_list,omitempty"`
}