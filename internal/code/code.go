package code

const (
	CodeOK                    = 0
	CodeInvalidParam          = 1
	CodeInternalServerError   = 2
	CodeNotSupportContentType = 3
	CodeRecordAlreadyExist    = 4
	CodeRecordNotFound        = 5
)

const (
	MsgOK                    = "ok"
	MsgInvalidParam          = "invalid param"
	MsgInternalServerError   = "internal server error"
	MsgNotSupportContentType = "not support content-type"
	MsgRecordAlreadyExist    = "record already exist"
	MsgRecordNotFound        = "record not found"
)
