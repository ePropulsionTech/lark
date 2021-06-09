// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// EventV2ContactUserUpdatedV3
//
// 通过该事件订阅员工变更。old_object中只展示更新的字段的原始值。事件只有在应用有数据权限的字段改动才会发送,具体的数据权限与字段的关系请参考[应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=contact&version=v3&resource=user&event=updated)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/events/updated
func (r *EventCallbackService) HandlerEventV2ContactUserUpdatedV3(f eventV2ContactUserUpdatedV3Handler) {
	r.cli.eventHandler.eventV2ContactUserUpdatedV3Handler = f
}

type eventV2ContactUserUpdatedV3Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2ContactUserUpdatedV3) (string, error)

type EventV2ContactUserUpdatedV3 struct {
	Object    *EventV2ContactUserUpdatedV3Object    `json:"object,omitempty"`     // 变更后信息
	OldObject *EventV2ContactUserUpdatedV3OldObject `json:"old_object,omitempty"` // 变更前信息
}

type EventV2ContactUserUpdatedV3Object struct {
	OpenID        string                                         `json:"open_id,omitempty"`        // 用户的open_id
	UserID        string                                         `json:"user_id,omitempty"`        // 租户内用户的唯一标识, 字段权限要求:  获取用户 userid
	Name          string                                         `json:"name,omitempty"`           // 用户名, 最小长度：`1` 字符
	EnName        string                                         `json:"en_name,omitempty"`        // 英文名
	Email         string                                         `json:"email,omitempty"`          // 邮箱, 字段权限要求:  获取用户邮箱
	Mobile        string                                         `json:"mobile,omitempty"`         // 手机号, 字段权限要求:  获取用户手机号
	Gender        int64                                          `json:"gender,omitempty"`         // 性别, 可选值有: `0`：未知, `1`：男, `2`：女
	Avatar        *EventV2ContactUserUpdatedV3ObjectAvatar       `json:"avatar,omitempty"`         // 用户头像信息
	Status        *EventV2ContactUserUpdatedV3ObjectStatus       `json:"status,omitempty"`         // 用户状态
	DepartmentIDs []string                                       `json:"department_ids,omitempty"` // 用户所属部门的ID列表
	LeaderUserID  string                                         `json:"leader_user_id,omitempty"` // 用户的直接主管的用户ID
	City          string                                         `json:"city,omitempty"`           // 城市
	Country       string                                         `json:"country,omitempty"`        // 国家
	WorkStation   string                                         `json:"work_station,omitempty"`   // 工位
	JoinTime      int64                                          `json:"join_time,omitempty"`      // 入职时间, 取值范围：`1` ～ `2147483647`
	EmployeeNo    string                                         `json:"employee_no,omitempty"`    // 工号
	EmployeeType  int64                                          `json:"employee_type,omitempty"`  // 员工类型, 可选值有: `1`：正式员工, `2`：实习生, `3`：外包, `4`：劳务, `5`：顾问
	Orders        []*EventV2ContactUserUpdatedV3ObjectOrder      `json:"orders,omitempty"`         // 用户排序信息
	CustomAttrs   []*EventV2ContactUserUpdatedV3ObjectCustomAttr `json:"custom_attrs,omitempty"`   // 自定义属性
}

type EventV2ContactUserUpdatedV3ObjectAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`     // 72*72像素头像链接
	Avatar240    string `json:"avatar_240,omitempty"`    // 240*240像素头像链接
	Avatar640    string `json:"avatar_640,omitempty"`    // 640*640像素头像链接
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 原始头像链接
}

type EventV2ContactUserUpdatedV3ObjectStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`    // 是否暂停
	IsResigned  bool `json:"is_resigned,omitempty"`  // 是否离职
	IsActivated bool `json:"is_activated,omitempty"` // 是否激活
}

type EventV2ContactUserUpdatedV3ObjectOrder struct {
	DepartmentID    string `json:"department_id,omitempty"`    // 排序信息对应的部门ID
	UserOrder       int64  `json:"user_order,omitempty"`       // 用户在其直属部门内的排序，数值越大，排序越靠前
	DepartmentOrder int64  `json:"department_order,omitempty"` // 用户所属的多个部门间的排序，数值越大，排序越靠前
}

type EventV2ContactUserUpdatedV3ObjectCustomAttr struct {
	Type  string                                            `json:"type,omitempty"`  // 自定义属性类型
	ID    string                                            `json:"id,omitempty"`    // 自定义属性ID
	Value *EventV2ContactUserUpdatedV3ObjectCustomAttrValue `json:"value,omitempty"` // 自定义属性取值
}

type EventV2ContactUserUpdatedV3ObjectCustomAttrValue struct {
	Text  string `json:"text,omitempty"`   // 属性文本
	URL   string `json:"url,omitempty"`    // URL
	PcURL string `json:"pc_url,omitempty"` // PC上的URL
}

type EventV2ContactUserUpdatedV3OldObject struct {
	OpenID        string                                            `json:"open_id,omitempty"`        // 用户的open_id
	UserID        string                                            `json:"user_id,omitempty"`        // 租户内用户的唯一标识, 字段权限要求:  获取用户 userid
	Name          string                                            `json:"name,omitempty"`           // 用户名, 最小长度：`1` 字符
	EnName        string                                            `json:"en_name,omitempty"`        // 英文名
	Email         string                                            `json:"email,omitempty"`          // 邮箱, 字段权限要求:  获取用户邮箱
	Mobile        string                                            `json:"mobile,omitempty"`         // 手机号, 字段权限要求:  获取用户手机号
	Gender        int64                                             `json:"gender,omitempty"`         // 性别, 可选值有: `0`：未知, `1`：男, `2`：女
	Avatar        *EventV2ContactUserUpdatedV3OldObjectAvatar       `json:"avatar,omitempty"`         // 用户头像信息
	Status        *EventV2ContactUserUpdatedV3OldObjectStatus       `json:"status,omitempty"`         // 用户状态
	DepartmentIDs []string                                          `json:"department_ids,omitempty"` // 用户所属部门的ID列表
	LeaderUserID  string                                            `json:"leader_user_id,omitempty"` // 用户的直接主管的用户ID
	City          string                                            `json:"city,omitempty"`           // 城市
	Country       string                                            `json:"country,omitempty"`        // 国家
	WorkStation   string                                            `json:"work_station,omitempty"`   // 工位
	JoinTime      int64                                             `json:"join_time,omitempty"`      // 入职时间, 取值范围：`1` ～ `2147483647`
	EmployeeNo    string                                            `json:"employee_no,omitempty"`    // 工号
	EmployeeType  int64                                             `json:"employee_type,omitempty"`  // 员工类型, 可选值有: `1`：正式员工, `2`：实习生, `3`：外包, `4`：劳务, `5`：顾问
	Orders        []*EventV2ContactUserUpdatedV3OldObjectOrder      `json:"orders,omitempty"`         // 用户排序信息
	CustomAttrs   []*EventV2ContactUserUpdatedV3OldObjectCustomAttr `json:"custom_attrs,omitempty"`   // 自定义属性
}

type EventV2ContactUserUpdatedV3OldObjectAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`     // 72*72像素头像链接
	Avatar240    string `json:"avatar_240,omitempty"`    // 240*240像素头像链接
	Avatar640    string `json:"avatar_640,omitempty"`    // 640*640像素头像链接
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 原始头像链接
}

type EventV2ContactUserUpdatedV3OldObjectStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`    // 是否暂停
	IsResigned  bool `json:"is_resigned,omitempty"`  // 是否离职
	IsActivated bool `json:"is_activated,omitempty"` // 是否激活
}

type EventV2ContactUserUpdatedV3OldObjectOrder struct {
	DepartmentID    string `json:"department_id,omitempty"`    // 排序信息对应的部门ID
	UserOrder       int64  `json:"user_order,omitempty"`       // 用户在其直属部门内的排序，数值越大，排序越靠前
	DepartmentOrder int64  `json:"department_order,omitempty"` // 用户所属的多个部门间的排序，数值越大，排序越靠前
}

type EventV2ContactUserUpdatedV3OldObjectCustomAttr struct {
	Type  string                                               `json:"type,omitempty"`  // 自定义属性类型
	ID    string                                               `json:"id,omitempty"`    // 自定义属性ID
	Value *EventV2ContactUserUpdatedV3OldObjectCustomAttrValue `json:"value,omitempty"` // 自定义属性取值
}

type EventV2ContactUserUpdatedV3OldObjectCustomAttrValue struct {
	Text  string `json:"text,omitempty"`   // 属性文本
	URL   string `json:"url,omitempty"`    // URL
	PcURL string `json:"pc_url,omitempty"` // PC上的URL
}
