/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 素材组信息
type MaterialGroupUpdateStruct struct {
	MaterialGroupId *int64                  `json:"material_group_id,omitempty"`
	Materials       *[]MaterialUpdateStruct `json:"materials,omitempty"`
	Previews        *[]PreviewUpdateStruct  `json:"previews,omitempty"`
}