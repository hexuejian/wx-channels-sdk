package callbacks

import "github.com/tidwall/gjson"

// 优惠券作废通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_coupon_invalid.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcCouponInvalid{})
}

type ChannelsEcCouponInvalid struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	CouponInfo   struct {
		CouponID    int64 `json:"coupon_id"`
		InvalidTime int64 `json:"invalid_time"`
	} `json:"coupon_info"`
}

func (ChannelsEcCouponInvalid) GetMessageType() string {
	return "event"
}

func (ChannelsEcCouponInvalid) GetEventType() string {
	return "channels_ec_coupon_invalid"
}

func (m ChannelsEcCouponInvalid) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcCouponInvalid) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	//var temp ChannelsEcCouponInvalid
	//err := json.Unmarshal(data, &temp)
	var temp = ChannelsEcCouponInvalid{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		CouponInfo: struct {
			CouponID    int64 `json:"coupon_id"`
			InvalidTime int64 `json:"invalid_time"`
		}{
			CouponID:    gjson.GetBytes(data, "coupon_info.coupon_id").Int(),
			InvalidTime: gjson.GetBytes(data, "coupon_info.invalid_time").Int(),
		},
	}
	return temp, nil
}
