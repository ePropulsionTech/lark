// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// RecognizeSpeechStream 语音流式接口，将整个音频文件分片进行传入模型。能够实时返回数据。建议每个音频分片的大小为 100-200ms
//
// 单租户限流：20 路（一个 stream_id 称为一路会话），同租户下的应用没有限流，共享本租户的 20路限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/speech_to_text-v1/speech/stream_recognize
func (r *AIService) RecognizeSpeechStream(ctx context.Context, request *RecognizeSpeechStreamReq, options ...MethodOptionFunc) (*RecognizeSpeechStreamResp, *Response, error) {
	if r.cli.mock.mockAIRecognizeSpeechStream != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] AI#RecognizeSpeechStream mock enable")
		return r.cli.mock.mockAIRecognizeSpeechStream(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "RecognizeSpeechStream",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/speech_to_text/v1/speech/stream_recognize",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(recognizeSpeechStreamResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAIRecognizeSpeechStream(f func(ctx context.Context, request *RecognizeSpeechStreamReq, options ...MethodOptionFunc) (*RecognizeSpeechStreamResp, *Response, error)) {
	r.mockAIRecognizeSpeechStream = f
}

func (r *Mock) UnMockAIRecognizeSpeechStream() {
	r.mockAIRecognizeSpeechStream = nil
}

type RecognizeSpeechStreamReq struct {
	Speech *RecognizeSpeechStreamReqSpeech `json:"speech,omitempty"` // 语音资源
	Config *RecognizeSpeechStreamReqConfig `json:"config,omitempty"` // 配置属性
}

type RecognizeSpeechStreamReqSpeech struct {
	Speech *string `json:"speech,omitempty"` // base64 后的音频文件进行, 示例值："base64 后的音频内容"
}

type RecognizeSpeechStreamReqConfig struct {
	StreamID   string `json:"stream_id,omitempty"`   // 仅包含字母数字和下划线的 16 位字符串作为同一数据流的标识，用户生成, 示例值："asd1234567890ddd"
	SequenceID int64  `json:"sequence_id,omitempty"` // 数据流分片的序号，序号从 0 开始，每次请求递增 1, 示例值：1
	Action     int64  `json:"action,omitempty"`      // 数据流标记：1 首包，2 正常结束，等待结果返回，3 中断数据流不返回最终结果, 示例值：1
	Format     string `json:"format,omitempty"`      // 语音格式，目前仅支持：pcm, 示例值："pcm"
	EngineType string `json:"engine_type,omitempty"` // 引擎类型，目前仅支持：16k_auto 中英混合, 示例值："16k_auto"
}

type recognizeSpeechStreamResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *RecognizeSpeechStreamResp `json:"data,omitempty"`
}

type RecognizeSpeechStreamResp struct {
	StreamID        string `json:"stream_id,omitempty"`        // 16 位 String 随机串作为同一数据流的标识
	SequenceID      int64  `json:"sequence_id,omitempty"`      // 数据流分片的序号，序号从 0 开始，每次请求递增 1
	RecognitionText string `json:"recognition_text,omitempty"` // 语音流识别后的文本信息
}
