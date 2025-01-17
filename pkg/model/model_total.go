/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 总计返回结构
type Total struct {
	Cpc                       *string `json:"cpc,omitempty"`
	Ctr                       *string `json:"ctr,omitempty"`
	Cost                      *string `json:"cost,omitempty"`
	ViewCount                 *string `json:"view_count,omitempty"`
	ValidClickCount           *string `json:"valid_click_count,omitempty"`
	ConversionsCount          *string `json:"conversions_count,omitempty"`
	ConversionsByClickCount   *string `json:"conversions_by_click_count,omitempty"`
	ConversionsByDisplayCount *string `json:"conversions_by_display_count,omitempty"`
	ConversionsRate           *string `json:"conversions_rate,omitempty"`
	ConversionsByDisplayRate  *string `json:"conversions_by_display_rate,omitempty"`
	ConversionsByClickRate    *string `json:"conversions_by_click_rate,omitempty"`
	ConversionsCost           *string `json:"conversions_cost,omitempty"`
	ConversionsByDisplayCost  *string `json:"conversions_by_display_cost,omitempty"`
	ConversionsByClickCost    *string `json:"conversions_by_click_cost,omitempty"`
	DeepConversionsCount      *string `json:"deep_conversions_count,omitempty"`
	DeepConversionsRate       *string `json:"deep_conversions_rate,omitempty"`
	DeepConversionsCost       *string `json:"deep_conversions_cost,omitempty"`
	ThousandDisplayPrice      *string `json:"thousand_display_price,omitempty"`
}
