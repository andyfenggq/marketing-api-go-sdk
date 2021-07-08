/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// PageTypeRead : 落地页类型
type PageTypeRead string

// List of PageTypeRead
const (
	PageTypeRead_DEFAULT                         PageTypeRead = "PAGE_TYPE_DEFAULT"
	PageTypeRead_TSA_APP                         PageTypeRead = "PAGE_TYPE_TSA_APP"
	PageTypeRead_TSA_WEB_NONE_ECOMMERCE          PageTypeRead = "PAGE_TYPE_TSA_WEB_NONE_ECOMMERCE"
	PageTypeRead_FENGYE_ECOMMERCE                PageTypeRead = "PAGE_TYPE_FENGYE_ECOMMERCE"
	PageTypeRead_CANVAS                          PageTypeRead = "PAGE_TYPE_CANVAS"
	PageTypeRead_MINI_PROGRAM                    PageTypeRead = "PAGE_TYPE_MINI_PROGRAM"
	PageTypeRead_CANVAS_WECHAT                   PageTypeRead = "PAGE_TYPE_CANVAS_WECHAT"
	PageTypeRead_MINI_PROGRAM_WECHAT             PageTypeRead = "PAGE_TYPE_MINI_PROGRAM_WECHAT"
	PageTypeRead_UNSUPPORTED                     PageTypeRead = "PAGE_TYPE_UNSUPPORTED"
	PageTypeRead_MINI_GAME_WECHAT                PageTypeRead = "PAGE_TYPE_MINI_GAME_WECHAT"
	PageTypeRead_FENGYE_EC_WECHAT_MINIPROGRAM    PageTypeRead = "PAGE_TYPE_FENGYE_EC_WECHAT_MINIPROGRAM"
	PageTypeRead_MINI_PROGRAM_QQ                 PageTypeRead = "PAGE_TYPE_MINI_PROGRAM_QQ"
	PageTypeRead_MINI_GAME_QQ                    PageTypeRead = "PAGE_TYPE_MINI_GAME_QQ"
	PageTypeRead_MINI_PROGRAM_CANVAS_WECHAT      PageTypeRead = "PAGE_TYPE_MINI_PROGRAM_CANVAS_WECHAT"
	PageTypeRead_MOMENTS_SIMPLE_NATIVE_WECHAT    PageTypeRead = "PAGE_TYPE_MOMENTS_SIMPLE_NATIVE_WECHAT"
	PageTypeRead_FULL_SCREEN_WECHAT              PageTypeRead = "PAGE_TYPE_FULL_SCREEN_WECHAT"
	PageTypeRead_YUEBAO_QUICKAPP                 PageTypeRead = "PAGE_TYPE_YUEBAO_QUICKAPP"
	PageTypeRead_YUEBAO_OFFICIAL_ACCOUNT_ARTICLE PageTypeRead = "PAGE_TYPE_YUEBAO_OFFICIAL_ACCOUNT_ARTICLE"
	PageTypeRead_XIJING_QUICK                    PageTypeRead = "PAGE_TYPE_XIJING_QUICK"
	PageTypeRead_WECHAT_CHANNELS                 PageTypeRead = "PAGE_TYPE_WECHAT_CHANNELS"
	PageTypeRead_CHANNELS_WATCH_LIVE             PageTypeRead = "PAGE_TYPE_CHANNELS_WATCH_LIVE"
	PageTypeRead_CHANNELS_RESERVE_LIVE           PageTypeRead = "PAGE_TYPE_CHANNELS_RESERVE_LIVE"
	PageTypeRead_YOUZAN_SINGLE                   PageTypeRead = "PAGE_TYPE_YOUZAN_SINGLE"
	PageTypeRead_YOUZAN_TOGETHER                 PageTypeRead = "PAGE_TYPE_YOUZAN_TOGETHER"
	PageTypeRead_YOUZAN_WECHAT_MINIPROGRAM       PageTypeRead = "PAGE_TYPE_YOUZAN_WECHAT_MINIPROGRAM"
	PageTypeRead_YIYE_FORM                       PageTypeRead = "PAGE_TYPE_YIYE_FORM"
	PageTypeRead_WEIMOB_PRODUCTSET               PageTypeRead = "PAGE_TYPE_WEIMOB_PRODUCTSET"
	PageTypeRead_WEIMOB_PROMOTION                PageTypeRead = "PAGE_TYPE_WEIMOB_PROMOTION"
	PageTypeRead_WEIMOB_PRODUCT                  PageTypeRead = "PAGE_TYPE_WEIMOB_PRODUCT"
	PageTypeRead_WEIMOB_H5                       PageTypeRead = "PAGE_TYPE_WEIMOB_H5"
	PageTypeRead_WECHAT_OFFICIAL_ACCOUNT_DETAIL  PageTypeRead = "PAGE_TYPE_WECHAT_OFFICIAL_ACCOUNT_DETAIL"
)
