package callbacks

import (
	"errors"
	"github.com/zsmhub/wx-channels-sdk/internal/envelope"
	"github.com/zsmhub/wx-channels-sdk/internal/signature"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type callbackUrlVars struct {
	Signature string
	Timestamp int64
	Nonce     string
	EchoStr   string
}

type CallbackHandler struct {
	token               string // 回调 token
	ep                  *envelope.Processor
	callbackParseImpMap map[string]CallbackExtraInfoInterface
	//rwLock              sync.RWMutex
}

func NewCallbackHandler(token string, encodingAESKey string) (*CallbackHandler, error) {
	ep, err := envelope.NewProcessor(token, encodingAESKey)
	if err != nil {
		return nil, err
	}

	handler := &CallbackHandler{token: token, ep: ep, callbackParseImpMap: make(map[string]CallbackExtraInfoInterface)}
	handler.RegisterCallBackImp()
	return handler, nil
}

// 解析并获取回调数据
func (cb *CallbackHandler) GetCallbackMsg(r *http.Request) (CallbackMessage, error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return CallbackMessage{}, err
	}

	// 验签
	ev, err := cb.ep.HandleIncomingMsg(r.URL, body)
	if err != nil {
		return CallbackMessage{}, err
	}

	// 解析json
	message, err := CallbackMessage{}.ParseMessageFromJson(ev.Msg, cb.callbackParseImpMap)
	if err != nil {
		return message, err
	}

	return message, nil
}

// 后台回调配置URL，申请校验
func (cb *CallbackHandler) EchoTestHandler(rw http.ResponseWriter, r *http.Request) {
	if !signature.VerifyHTTPRequestSignature(cb.token, r.URL) {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	args, err := cb.parseUrlVars(r.URL.Query())
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write([]byte(args.EchoStr))
}

func (cb *CallbackHandler) parseUrlVars(urlVars url.Values) (callbackUrlVars, error) {
	var errMalformedArgs = errors.New("malformed arguments for echo test API")

	var msgSignature string
	{
		l := urlVars["signature"]
		if len(l) != 1 {
			return callbackUrlVars{}, errMalformedArgs
		}
		msgSignature = l[0]
	}

	var timestamp int64
	{
		l := urlVars["timestamp"]
		if len(l) != 1 {
			return callbackUrlVars{}, errMalformedArgs
		}
		timestampStr := l[0]

		timestampInt, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			return callbackUrlVars{}, errMalformedArgs
		}

		timestamp = timestampInt
	}

	var nonce string
	{
		l := urlVars["nonce"]
		if len(l) != 1 {
			return callbackUrlVars{}, errMalformedArgs
		}
		nonce = l[0]
	}

	var echoStr string
	{
		l := urlVars["echostr"]
		if len(l) != 1 {
			return callbackUrlVars{}, errMalformedArgs
		}
		echoStr = l[0]
	}

	return callbackUrlVars{
		Signature: msgSignature,
		Timestamp: timestamp,
		Nonce:     nonce,
		EchoStr:   echoStr,
	}, nil
}

// RegisterCallBackImp 注册回调函数实现
func (cb *CallbackHandler) RegisterCallBackImp() {
	item1 := ChannelsEcCouponInvalid{}
	cb.callbackParseImpMap[item1.GetTypeKey()] = item1
	item2 := ChannelsEcCouponDelete{}
	cb.callbackParseImpMap[item2.GetTypeKey()] = item2
	item3 := ChannelsEcUserCouponUse{}
	cb.callbackParseImpMap[item3.GetTypeKey()] = item3
	item4 := ChannelsEcCouponExpire{}
	cb.callbackParseImpMap[item4.GetTypeKey()] = item4
	item5 := ChannelsEcUserCouponUnuse{}
	cb.callbackParseImpMap[item5.GetTypeKey()] = item5
	item6 := ChannelsEcCouponReceive{}
	cb.callbackParseImpMap[item6.GetTypeKey()] = item6
	item7 := ChannelsEcCouponCreate{}
	cb.callbackParseImpMap[item7.GetTypeKey()] = item7
	item8 := ChannelsEcCouponInfoChange{}
	cb.callbackParseImpMap[item8.GetTypeKey()] = item8
	item9 := ChannelsEcUserCouponExpire{}
	cb.callbackParseImpMap[item9.GetTypeKey()] = item9
}
