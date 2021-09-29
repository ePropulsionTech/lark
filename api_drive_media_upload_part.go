// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
	"io"
)

// PartUploadDriveMedia 上传对应的文件块。
//
// 该接口不支持太高的并发，且调用频率上限为5QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_part
func (r *DriveService) PartUploadDriveMedia(ctx context.Context, request *PartUploadDriveMediaReq, options ...MethodOptionFunc) (*PartUploadDriveMediaResp, *Response, error) {
	if r.cli.mock.mockDrivePartUploadDriveMedia != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#PartUploadDriveMedia mock enable")
		return r.cli.mock.mockDrivePartUploadDriveMedia(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "PartUploadDriveMedia",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/medias/upload_part",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
		IsFile:                true,
	}
	resp := new(partUploadDriveMediaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDrivePartUploadDriveMedia(f func(ctx context.Context, request *PartUploadDriveMediaReq, options ...MethodOptionFunc) (*PartUploadDriveMediaResp, *Response, error)) {
	r.mockDrivePartUploadDriveMedia = f
}

func (r *Mock) UnMockDrivePartUploadDriveMedia() {
	r.mockDrivePartUploadDriveMedia = nil
}

type PartUploadDriveMediaReq struct {
	UploadID string    `json:"upload_id,omitempty"` // 分片上传事务ID, 示例值："123456"
	Seq      int64     `json:"seq,omitempty"`       // 块号，从0开始计数, 示例值：0
	Size     int64     `json:"size,omitempty"`      // 块大小, 示例值：4194304
	Checksum *string   `json:"checksum,omitempty"`  // 文件分块adler32校验和(可选), 示例值："12345678"
	File     io.Reader `json:"file,omitempty"`      // 文件分片内容, 示例值：file binary
}

type partUploadDriveMediaResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *PartUploadDriveMediaResp `json:"data,omitempty"`
}

type PartUploadDriveMediaResp struct{}
