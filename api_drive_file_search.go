// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// SearchDriveFile 该接口用于根据搜索条件进行文档搜索。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ugDM4UjL4ADO14COwgTN
func (r *DriveService) SearchDriveFile(ctx context.Context, request *SearchDriveFileReq, options ...MethodOptionFunc) (*SearchDriveFileResp, *Response, error) {
	if r.cli.mock.mockDriveSearchDriveFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#SearchDriveFile mock enable")
		return r.cli.mock.mockDriveSearchDriveFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Drive",
		API:                 "SearchDriveFile",
		Method:              "POST",
		URL:                 "https://open.feishu.cn/open-apis/suite/docs-api/search/object",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(searchDriveFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveSearchDriveFile(f func(ctx context.Context, request *SearchDriveFileReq, options ...MethodOptionFunc) (*SearchDriveFileResp, *Response, error)) {
	r.mockDriveSearchDriveFile = f
}

func (r *Mock) UnMockDriveSearchDriveFile() {
	r.mockDriveSearchDriveFile = nil
}

type SearchDriveFileReq struct {
	SearchKey *string  `json:"search_key,omitempty"` // 搜索关键字
	Count     *int64   `json:"count,omitempty"`      // 搜索返回数量，0 <= count <= 50
	Offset    *int64   `json:"offset,omitempty"`     // 搜索偏移位，offset >= 0，offset + count < 10000
	OwnerIDs  []string `json:"owner_ids,omitempty"`  // 文档所有者
	ChatIDs   []string `json:"chat_ids,omitempty"`   // 文档所在群
	DocsTypes []string `json:"docs_types,omitempty"` // 文档类型，支持："doc", "sheet", "slide", "bitable", "mindnote", "file"
}

type searchDriveFileResp struct {
	Code int64                `json:"code,omitempty"`
	Msg  string               `json:"msg,omitempty"`
	Data *SearchDriveFileResp `json:"data,omitempty"`
}

type SearchDriveFileResp struct {
	DocsEntities *SearchDriveFileRespDocsEntities `json:"docs_entities,omitempty"` // 搜索匹配文档列表
	HasMore      bool                             `json:"has_more,omitempty"`      // 搜索偏移位结果列表后是否还有数据
	Total        int64                            `json:"total,omitempty"`         // 搜索匹配文档总数量
}

type SearchDriveFileRespDocsEntities struct {
	DocsToken string `json:"docs_token,omitempty"` // 文档token
	DocsType  string `json:"docs_type,omitempty"`  // 文档类型
	Title     string `json:"title,omitempty"`      // 标题
	DocsOwner string `json:"docs_owner,omitempty"` // 文件所有者
}
