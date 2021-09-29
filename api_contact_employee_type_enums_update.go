// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateEmployeeTypeEnumPatch 更新自定义人员类型
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/update
func (r *ContactService) UpdateEmployeeTypeEnumPatch(ctx context.Context, request *UpdateEmployeeTypeEnumPatchReq, options ...MethodOptionFunc) (*UpdateEmployeeTypeEnumPatchResp, *Response, error) {
	if r.cli.mock.mockContactUpdateEmployeeTypeEnumPatch != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#UpdateEmployeeTypeEnumPatch mock enable")
		return r.cli.mock.mockContactUpdateEmployeeTypeEnumPatch(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "UpdateEmployeeTypeEnumPatch",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/employee_type_enums/:enum_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateEmployeeTypeEnumPatchResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockContactUpdateEmployeeTypeEnumPatch(f func(ctx context.Context, request *UpdateEmployeeTypeEnumPatchReq, options ...MethodOptionFunc) (*UpdateEmployeeTypeEnumPatchResp, *Response, error)) {
	r.mockContactUpdateEmployeeTypeEnumPatch = f
}

func (r *Mock) UnMockContactUpdateEmployeeTypeEnumPatch() {
	r.mockContactUpdateEmployeeTypeEnumPatch = nil
}

type UpdateEmployeeTypeEnumPatchReq struct {
	EnumID      string                                     `path:"enum_id" json:"-"`       // 枚举值id, 示例值："exGeIjow7zIqWMy+ONkFxA=="
	Content     string                                     `json:"content,omitempty"`      // 枚举内容, 示例值："专家", 长度范围：`1` ～ `100` 字符
	EnumType    int64                                      `json:"enum_type,omitempty"`    // 类型, 示例值：2, 可选值有: `1`：内置类型, `2`：自定义
	EnumStatus  int64                                      `json:"enum_status,omitempty"`  // 使用状态, 示例值：1, 可选值有: `1`：激活, `2`：未激活
	I18nContent *UpdateEmployeeTypeEnumPatchReqI18nContent `json:"i18n_content,omitempty"` // i18n定义
}

type UpdateEmployeeTypeEnumPatchReqI18nContent struct {
	Locale *string `json:"locale,omitempty"` // 语言版本, 示例值："zh_cn"
	Value  *string `json:"value,omitempty"`  // 字段名, 示例值："专家"
}

type updateEmployeeTypeEnumPatchResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *UpdateEmployeeTypeEnumPatchResp `json:"data,omitempty"`
}

type UpdateEmployeeTypeEnumPatchResp struct {
	EmployeeTypeEnum *UpdateEmployeeTypeEnumPatchRespEmployeeTypeEnum `json:"employee_type_enum,omitempty"` // 更新后的人员类型字段
}

type UpdateEmployeeTypeEnumPatchRespEmployeeTypeEnum struct {
	EnumID      string                                                      `json:"enum_id,omitempty"`      // 枚举值id
	EnumValue   string                                                      `json:"enum_value,omitempty"`   // 枚举值
	Content     string                                                      `json:"content,omitempty"`      // 枚举内容, 长度范围：`1` ～ `100` 字符
	EnumType    int64                                                       `json:"enum_type,omitempty"`    // 类型, 可选值有: `1`：内置类型, `2`：自定义
	EnumStatus  int64                                                       `json:"enum_status,omitempty"`  // 使用状态, 可选值有: `1`：激活, `2`：未激活
	I18nContent *UpdateEmployeeTypeEnumPatchRespEmployeeTypeEnumI18nContent `json:"i18n_content,omitempty"` // i18n定义
}

type UpdateEmployeeTypeEnumPatchRespEmployeeTypeEnumI18nContent struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}
