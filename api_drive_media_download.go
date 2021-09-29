// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
	"io"
)

// DownloadDriveMedia 使用该接口可以下载素材。素材表示在各种创作容器里的文件，如Doc文档内的图片，文件均属于素材。支持range下载。
//
// 该接口不支持太高的并发，且调用频率上限为5QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/download
func (r *DriveService) DownloadDriveMedia(ctx context.Context, request *DownloadDriveMediaReq, options ...MethodOptionFunc) (*DownloadDriveMediaResp, *Response, error) {
	if r.cli.mock.mockDriveDownloadDriveMedia != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DownloadDriveMedia mock enable")
		return r.cli.mock.mockDriveDownloadDriveMedia(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DownloadDriveMedia",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/medias/:file_token/download",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(downloadDriveMediaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveDownloadDriveMedia(f func(ctx context.Context, request *DownloadDriveMediaReq, options ...MethodOptionFunc) (*DownloadDriveMediaResp, *Response, error)) {
	r.mockDriveDownloadDriveMedia = f
}

func (r *Mock) UnMockDriveDownloadDriveMedia() {
	r.mockDriveDownloadDriveMedia = nil
}

type DownloadDriveMediaReq struct {
	Extra     *string  `query:"extra" json:"-"`     // 扩展信息, 示例值："[请参考-上传点类型及对应Extra说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/introduction)"
	FileToken string   `path:"file_token" json:"-"` // 文件的 token，获取方式见 [概述](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/files/guide/introduction), 示例值："boxcnabCdefg12345"
	Range     [2]int64 `header:"range" json:"-"`    // 指定文件下载部分，示例:bytes=0-1024
}

type downloadDriveMediaResp struct {
	IsFile bool
	Code   int64
	Msg    string
	Data   *DownloadDriveMediaResp
}

func (r *downloadDriveMediaResp) SetReader(file io.Reader) {
	if r.Data == nil {
		r.Data = &DownloadDriveMediaResp{}
	}
	r.Data.File = file
}

func (r *downloadDriveMediaResp) SetFilename(filename string) {
	if r.Data == nil {
		r.Data = &DownloadDriveMediaResp{}
	}
	r.Data.Filename = filename
}

type DownloadDriveMediaResp struct {
	File     io.Reader
	Filename string // 文件名
}
