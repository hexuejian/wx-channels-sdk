package callbacks

import (
	"github.com/tidwall/gjson"
)

// 优惠券过期通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_coupon_expire.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcCouponExpire{})
}

type ChannelsEcCouponExpire struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	CouponInfo   struct {
		CouponID   string `json:"coupon_id"`
		ExpireTime int64  `json:"expire_time"`
	} `json:"coupon_info"`
}

func (ChannelsEcCouponExpire) GetMessageType() string {
	return "event"
}

func (ChannelsEcCouponExpire) GetEventType() string {
	return "channels_ec_coupon_expire"
}

func (m ChannelsEcCouponExpire) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcCouponExpire) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcCouponExpire{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		CouponInfo: struct {
			CouponID   string `json:"coupon_id"`
			ExpireTime int64  `json:"expire_time"`
		}{
			CouponID:   gjson.GetBytes(data, "coupon_info.coupon_id").String(),
			ExpireTime: gjson.GetBytes(data, "coupon_info.expire_time").Int(),
		},
	}
	return temp, nil
}
