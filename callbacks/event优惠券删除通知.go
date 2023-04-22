package callbacks

import (
	"github.com/tidwall/gjson"
)

// 优惠券删除通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_coupon_delete.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcCouponDelete{})
}

type ChannelsEcCouponDelete struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	CouponInfo   struct {
		CouponID   string `json:"coupon_id"`
		DeleteTime int64  `json:"delete_time"`
	} `json:"coupon_info"`
}

func (ChannelsEcCouponDelete) GetMessageType() string {
	return "event"
}

func (ChannelsEcCouponDelete) GetEventType() string {
	return "channels_ec_coupon_delete"
}

func (m ChannelsEcCouponDelete) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcCouponDelete) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcCouponDelete{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		CouponInfo: struct {
			CouponID   string `json:"coupon_id"`
			DeleteTime int64  `json:"delete_time"`
		}{
			CouponID:   gjson.GetBytes(data, "coupon_info.coupon_id").String(),
			DeleteTime: gjson.GetBytes(data, "coupon_info.delete_time").Int(),
		},
	}
	return temp, nil
}
