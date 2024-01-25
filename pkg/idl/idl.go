package idl

type Input struct {
	Action   string `json:"action" validate:"required" mapstructure:"action"`
	PageNum  int64  `json:"page_num"`
	PageSize int64  `json:"page_size"`
}

type Output struct {
	Code  int64  `json:"code"`
	Msg   string `json:"msg"`
	Total int64  `json:"total"`
}
