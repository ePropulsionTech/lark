// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetDriveDocContent
//
// 在使用此接口前，请仔细阅读[概述](https://open.feishu.cn/document/ukTMukTMukTM/ukjM5YjL5ITO24SOykjN)和[准备接入文档 API](//ssl:ttdoc/ukTMukTMukTM/ugzNzUjL4czM14CO3MTN/guide/getting-start)了解文档调用的规则和约束，确保你的文档数据不会丢失或出错。
// 文档数据结构定义可参考：[文档数据结构概述](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)
// 该接口用于获取结构化的文档内容。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDM2YjL1AjN24SNwYjN
func (r *DriveService) GetDriveDocContent(ctx context.Context, request *GetDriveDocContentReq, options ...MethodOptionFunc) (*GetDriveDocContentResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveDocContent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDriveDocContent mock enable")
		return r.cli.mock.mockDriveGetDriveDocContent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveDocContent",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/doc/v2/:docToken/content",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveDocContentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveGetDriveDocContent(f func(ctx context.Context, request *GetDriveDocContentReq, options ...MethodOptionFunc) (*GetDriveDocContentResp, *Response, error)) {
	r.mockDriveGetDriveDocContent = f
}

func (r *Mock) UnMockDriveGetDriveDocContent() {
	r.mockDriveGetDriveDocContent = nil
}

type GetDriveDocContentReq struct {
	DocToken string `path:"docToken" json:"-"` // 获取方式详见 [准备接入文档 API](https://open.feishu.cn/document/ukTMukTMukTM/ugzNzUjL4czM14CO3MTN/guide/getting-start)
}

type getDriveDocContentResp struct {
	Code int64                   `json:"code,omitempty"`
	Msg  string                  `json:"msg,omitempty"`
	Data *GetDriveDocContentResp `json:"data,omitempty"`
}

type GetDriveDocContentResp struct {
	Content  string `json:"content,omitempty"`  // 详情参考[文档数据结构](https://open.feishu.cn/document/ukTMukTMukTM/ukDM2YjL5AjN24SOwYjN)
	Revision int64  `json:"revision,omitempty"` // 文档当前版本号
}
