// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetHelpdeskAgentSkill 该接口用于获取客服技能
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill/get
func (r *HelpdeskService) GetHelpdeskAgentSkill(ctx context.Context, request *GetHelpdeskAgentSkillReq, options ...MethodOptionFunc) (*GetHelpdeskAgentSkillResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetHelpdeskAgentSkill != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetHelpdeskAgentSkill mock enable")
		return r.cli.mock.mockHelpdeskGetHelpdeskAgentSkill(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetHelpdeskAgentSkill",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/agent_skills/:agent_skill_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getHelpdeskAgentSkillResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHelpdeskGetHelpdeskAgentSkill(f func(ctx context.Context, request *GetHelpdeskAgentSkillReq, options ...MethodOptionFunc) (*GetHelpdeskAgentSkillResp, *Response, error)) {
	r.mockHelpdeskGetHelpdeskAgentSkill = f
}

func (r *Mock) UnMockHelpdeskGetHelpdeskAgentSkill() {
	r.mockHelpdeskGetHelpdeskAgentSkill = nil
}

type GetHelpdeskAgentSkillReq struct {
	AgentSkillID string `path:"agent_skill_id" json:"-"` // agent skill id, 示例值："6941215891786825756"
}

type getHelpdeskAgentSkillResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *GetHelpdeskAgentSkillResp `json:"data,omitempty"`
}

type GetHelpdeskAgentSkillResp struct {
	AgentSkill *GetHelpdeskAgentSkillRespAgentSkill `json:"agent_skill,omitempty"` // 技能
}

type GetHelpdeskAgentSkillRespAgentSkill struct {
	ID        string                                      `json:"id,omitempty"`         // 技能id
	Name      string                                      `json:"name,omitempty"`       // 技能名
	Rules     []*GetHelpdeskAgentSkillRespAgentSkillRule  `json:"rules,omitempty"`      // 技能rules
	AgentIDs  []string                                    `json:"agent_ids,omitempty"`  // 具有此技能的客服ids
	IsDefault bool                                        `json:"is_default,omitempty"` // 默认技能
	Agents    []*GetHelpdeskAgentSkillRespAgentSkillAgent `json:"agents,omitempty"`     // 客服 info
}

type GetHelpdeskAgentSkillRespAgentSkillRule struct {
	ID               string  `json:"id,omitempty"`                // rule id, 看[获取客服技能rules](https://open.feishu.cn/document/ukTMukTMukTM/ucDOyYjL3gjM24yN4IjN/list-agent-skill-rules) 用于获取rules options
	SelectedOperator int64   `json:"selected_operator,omitempty"` // 运算符compare, 看[客服技能运算符options](https://open.feishu.cn/document/ukTMukTMukTM/ucDOyYjL3gjM24yN4IjN/operator-options)
	OperatorOptions  []int64 `json:"operator_options,omitempty"`  // rule操作数value，[客服技能及运算符](https://open.feishu.cn/document/ukTMukTMukTM/ucDOyYjL3gjM24yN4IjN/operator-options)
	Operand          string  `json:"operand,omitempty"`           // rule操作数value
	Category         int64   `json:"category,omitempty"`          // rule type，1-知识库，2-工单信息，3-用户飞书信息
	DisplayName      string  `json:"display_name,omitempty"`      // rule 名
}

type GetHelpdeskAgentSkillRespAgentSkillAgent struct {
	ID        string `json:"id,omitempty"`         // user id
	AvatarURL string `json:"avatar_url,omitempty"` // user avatar url
	Name      string `json:"name,omitempty"`       // user name
}
