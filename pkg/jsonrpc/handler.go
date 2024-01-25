package jsonrpc

import (
	"context"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/haormj/util"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"

	"github.com/haormj/go-apiserver-template/internal/code"
	"github.com/haormj/go-apiserver-template/pkg/idl"
	"github.com/haormj/go-apiserver-template/pkg/jsonrpc/invoker"
	"github.com/haormj/go-apiserver-template/pkg/jsonrpc/invoker/receiver"
	"github.com/haormj/go-apiserver-template/pkg/log"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewHandleFunc(ctx context.Context, service any) (http.HandlerFunc, error) {
	inv := receiver.NewInvoker(service)
	if err := inv.Init(); err != nil {
		return nil, err
	}
	validate := validator.New()

	fn := func(w http.ResponseWriter, req *http.Request) {
		// 仅支持 json
		contentType := req.Header.Get("Content-Type")
		if !strings.Contains(contentType, "application/json") {
			errOutput := idl.Output{
				Code: code.CodeNotSupportContentType,
				Msg:  code.MsgNotSupportContentType,
			}

			JSONResponse(w, http.StatusOK, errOutput)
			return
		}

		errOutput := idl.Output{
			Code: code.CodeInternalServerError,
			Msg:  code.MsgInternalServerError,
		}

		defer req.Body.Close()
		b, err := io.ReadAll(req.Body)
		if err != nil {
			JSONResponse(w, http.StatusOK, errOutput)
			return
		}
		action := json.Get(b, "action").ToString()
		// 通过action寻找对应的函数
		fn, err := inv.Function(action)
		if err != nil {
			JSONResponse(w, http.StatusOK, errOutput)
			return
		}
		// 获取函数的输入参数，并初始化
		in := fn.In()
		inputVal := util.InitPointer(in[1])
		outputVal := util.InitPointer(in[2])
		util.SetStructField(outputVal.Interface(), code.CodeOK, "Output", "Code")
		util.SetStructField(outputVal.Interface(), code.MsgOK, "Output", "Msg")
		inputMap := make(map[string]any)
		if err := json.Unmarshal(b, &inputMap); err != nil {
			errOutput.Code = code.CodeInvalidParam
			errOutput.Msg = code.MsgInvalidParam
			JSONResponse(w, http.StatusOK, errOutput)
			return
		}
		// 主要是客户端在传参的时候，可能会存在传递字符串数字这类情况，服务端需要兼容这种情况
		config := &mapstructure.DecoderConfig{
			Metadata:         nil,
			Result:           inputVal.Addr().Interface(),
			TagName:          "json",
			WeaklyTypedInput: true,
			Squash:           true,
			DecodeHook:       mapstructure.ComposeDecodeHookFunc(StringTrimSpaceHookFunc()),
		}
		decoder, err := mapstructure.NewDecoder(config)
		if err != nil {
			errOutput.Code = code.CodeInvalidParam
			errOutput.Msg = code.MsgInvalidParam
			JSONResponse(w, http.StatusOK, errOutput)
			return
		}
		if err := decoder.Decode(inputMap); err != nil {
			errOutput.Code = code.CodeInvalidParam
			errOutput.Msg = code.MsgInvalidParam
			JSONResponse(w, http.StatusOK, errOutput)
			return
		}
		// 检验输入参数是否合法
		if err := validate.Struct(inputVal.Interface()); err != nil {
			errOutput.Code = code.CodeInvalidParam
			errOutput.Msg = code.MsgInvalidParam
			JSONResponse(w, http.StatusOK, errOutput)
			return
		}
		// pageSize 配置默认值为10
		pageSize, ok := util.GetStructField(inputVal.Interface(), "Input", "PageSize").Int()
		if ok && pageSize == 0 {
			util.SetStructField(inputVal.Interface(), 10, "Input", "PageSize")
		}
		mi := invoker.NewMessage()
		mi.SetFuncName(action)
		mi.SetParameters([]any{ctx, inputVal.Interface(), outputVal.Interface()})
		mo, err := inv.Invoke(ctx, mi)
		// 调用报错，返回错误信息
		if err != nil {
			JSONResponse(w, http.StatusOK, errOutput)
			return
		}

		// 调用未报错，返回处理后数据
		if mo.Parameters()[0] == nil {
			JSONResponse(w, http.StatusOK, outputVal.Interface())
			return
		}
		// 一定是 error 类型
		err, ok = mo.Parameters()[0].(error)
		if !ok {
			JSONResponse(w, http.StatusOK, errOutput)
			return
		}

		if err != nil {
			JSONResponse(w, http.StatusOK, errOutput)
			return
		}
	}

	return fn, nil
}

// JSONResponse response body with json format
func JSONResponse(w http.ResponseWriter, statusCode int, body any) {
	data, err := json.Marshal(body)
	if err != nil {
		log.Errorf("json.Marshal error err: %v", err)
		return
	}

	w.Write(data)
	w.WriteHeader(statusCode)
}

// StringTrimSpaceHookFunc returns a DecodeHookFunc that trim string space
func StringTrimSpaceHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Kind,
		t reflect.Kind,
		data any) (any, error) {
		if f != reflect.String {
			return data, nil
		}

		raw := data.(string)

		return strings.TrimSpace(raw), nil
	}
}
