// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetAttendanceUserDailyShift
//
// 支持查询多个用户的排班情况，查询的时间跨度不能超过 30 天。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/GetScheduledShifts
func (r *AttendanceService) GetAttendanceUserDailyShift(ctx context.Context, request *GetAttendanceUserDailyShiftReq, options ...MethodOptionFunc) (*GetAttendanceUserDailyShiftResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceUserDailyShift != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceUserDailyShift mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceUserDailyShift(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceUserDailyShift",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/user_daily_shifts/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceUserDailyShiftResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAttendanceGetAttendanceUserDailyShift(f func(ctx context.Context, request *GetAttendanceUserDailyShiftReq, options ...MethodOptionFunc) (*GetAttendanceUserDailyShiftResp, *Response, error)) {
	r.mockAttendanceGetAttendanceUserDailyShift = f
}

func (r *Mock) UnMockAttendanceGetAttendanceUserDailyShift() {
	r.mockAttendanceGetAttendanceUserDailyShift = nil
}

type GetAttendanceUserDailyShiftReq struct {
	EmployeeType  EmployeeType `query:"employee_type" json:"-"`   // 请求体中的 user_ids 的员工工号类型，可用值：【employee_id（员工的 employeeId），employee_no（员工工号）】，示例值：“employee_id”
	UserIDs       []string     `json:"user_ids,omitempty"`        // employee_no 或 employee_id 列表
	CheckDateFrom int64        `json:"check_date_from,omitempty"` // 查询的起始工作日
	CheckDateTo   int64        `json:"check_date_to,omitempty"`   // 查询的结束工作日
}

type getAttendanceUserDailyShiftResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *GetAttendanceUserDailyShiftResp `json:"data,omitempty"` // -
}

type GetAttendanceUserDailyShiftResp struct {
	UserDailyShifts []*GetAttendanceUserDailyShiftRespUserDailyShift `json:"user_daily_shifts,omitempty"` // 班表信息列表
}

type GetAttendanceUserDailyShiftRespUserDailyShift struct {
	GroupID    string `json:"group_id,omitempty"`    // 考勤组 ID
	ShiftID    string `json:"shift_id,omitempty"`    // 班次 ID，休息为 0
	Month      int64  `json:"month,omitempty"`       // 月份
	EmployeeNo string `json:"employee_no,omitempty"` // 用户
	DayNo      int64  `json:"day_no,omitempty"`      // 日期
}
