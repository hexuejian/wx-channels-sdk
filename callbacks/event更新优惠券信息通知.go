package callbacks

import (
	"github.com/tidwall/gjson"
)

// 更新优惠券信息通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_coupon_info_change.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcCouponInfoChange{})
}

type ChannelsEcCouponInfoChange struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	CouponInfo   struct {
		ChangeTime int64  `json:"change_time"`
		CouponID   string `json:"coupon_id"`
	} `json:"coupon_info"`
}

func (ChannelsEcCouponInfoChange) GetMessageType() string {
	return "event"
}

func (ChannelsEcCouponInfoChange) GetEventType() string {
	return "channels_ec_coupon_info_change"
}

func (m ChannelsEcCouponInfoChange) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcCouponInfoChange) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcCouponInfoChange{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		CouponInfo: struct {
			ChangeTime int64  `json:"change_time"`
			CouponID   string `json:"coupon_id"`
		}{
			CouponID:   gjson.GetBytes(data, "coupon_info.coupon_id").String(),
			ChangeTime: gjson.GetBytes(data, "coupon_info.change_time").Int(),
		},
	}
	return temp, nil
}
