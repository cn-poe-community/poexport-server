package rest

type restResp struct {
	Code    string
	Message string
	Data    map[string]any
}

var EmptyReturnData = map[string]any{}

func newRestResp(code, message string, data map[string]any) *restResp {
	return &restResp{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (r *restResp) toMap() map[string]any {
	m := map[string]any{}
	m["code"] = r.Code
	m["msg"] = r.Message
	m["data"] = r.Data
	return m
}

func Resp(code, message string, data map[string]any) map[string]any {
	resp := &restResp{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return resp.toMap()
}

func EmptyResp(code, message string) map[string]any {
	resp := &restResp{
		Code:    code,
		Message: message,
		Data:    map[string]any{},
	}
	return resp.toMap()
}
