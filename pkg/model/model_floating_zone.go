/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 浮层卡片创意内容
type FloatingZone struct {
	FloatingZoneSwitch        *bool            `json:"floating_zone_switch,omitempty"`
	FloatingZoneImageId       *string          `json:"floating_zone_image_id,omitempty"`
	FloatingZoneName          *string          `json:"floating_zone_name,omitempty"`
	FloatingZoneDesc          *string          `json:"floating_zone_desc,omitempty"`
	FloatingZoneButtonText    *string          `json:"floating_zone_button_text,omitempty"`
	FloatingZoneType          FloatingZoneType `json:"floating_zone_type,omitempty"`
	FloatingZoneSingleImageId *string          `json:"floating_zone_single_image_id,omitempty"`
}
