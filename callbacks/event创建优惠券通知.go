package callbacks

import (
	"github.com/tidwall/gjson"
)

// 创建优惠券通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_coupon_create.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcCouponCreate{})
}

type ChannelsEcCouponCreate struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	CouponInfo   struct {
		CouponID   string `json:"coupon_id"`
		CreateTime int64  `json:"create_time"`
	} `json:"coupon_info"`
}

func (ChannelsEcCouponCreate) GetMessageType() string {
	return "event"
}

func (ChannelsEcCouponCreate) GetEventType() string {
	return "channels_ec_coupon_create"
}

func (m ChannelsEcCouponCreate) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcCouponCreate) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcCouponCreate{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		CouponInfo: struct {
			CouponID   string `json:"coupon_id"`
			CreateTime int64  `json:"create_time"`
		}{
			CouponID:   gjson.GetBytes(data, "coupon_info.coupon_id").String(),
			CreateTime: gjson.GetBytes(data, "coupon_info.create_time").Int(),
		},
	}
	return temp, nil
}
