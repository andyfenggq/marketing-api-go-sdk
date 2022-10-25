/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意组件信息
type CreativeComponentStruct struct {
	AccountId          *int64                       `json:"account_id,omitempty"`
	ComponentName      *string                      `json:"component_name,omitempty"`
	ComponentId        *int64                       `json:"component_id,omitempty"`
	ComponentType      CreativeComponentType        `json:"component_type,omitempty"`
	PromotedObjectType PromotedObjectType           `json:"promoted_object_type,omitempty"`
	PromotedObjectId   *string                      `json:"promoted_object_id,omitempty"`
	PromotedObjectName *string                      `json:"promoted_object_name,omitempty"`
	Status             AdStatus                     `json:"status,omitempty"`
	AuditStatus        CreativeComponentAuditStatus `json:"audit_status,omitempty"`
	AuditMsg           *string                      `json:"audit_msg,omitempty"`
	ComponentSpec      *CreativeComponentSpecStruct `json:"component_spec,omitempty"`
	PromotedObjectSpec *PromotedObjectSpec          `json:"promoted_object_spec,omitempty"`
	CreatedTime        *int64                       `json:"created_time,omitempty"`
	LastModifiedTime   *int64                       `json:"last_modified_time,omitempty"`
	IsPublishEnabled   *bool                        `json:"is_publish_enabled,omitempty"`
}
