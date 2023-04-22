package callbacks

import (
	"github.com/tidwall/gjson"
)

// 优惠券返还通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_user_coupon_unuse.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcUserCouponUnuse{})
}

type ChannelsEcUserCouponUnuse struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	UseInfo      struct {
		CouponID     string `json:"coupon_id"`
		OrderID      string `json:"order_id"`
		UnuseTime    int64  `json:"unuse_time"`
		UserCouponID string `json:"user_coupon_id"`
	} `json:"use_info"`
}

func (ChannelsEcUserCouponUnuse) GetMessageType() string {
	return "event"
}

func (ChannelsEcUserCouponUnuse) GetEventType() string {
	return "channels_ec_user_coupon_unuse"
}

func (m ChannelsEcUserCouponUnuse) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcUserCouponUnuse) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcUserCouponUnuse{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		UseInfo: struct {
			CouponID     string `json:"coupon_id"`
			OrderID      string `json:"order_id"`
			UnuseTime    int64  `json:"unuse_time"`
			UserCouponID string `json:"user_coupon_id"`
		}{
			CouponID:     gjson.GetBytes(data, "use_info.coupon_id").String(),
			OrderID:      gjson.GetBytes(data, "use_info.order_id").String(),
			UserCouponID: gjson.GetBytes(data, "use_info.user_coupon_id").String(),
			UnuseTime:    gjson.GetBytes(data, "use_info.unuse_time").Int(),
		},
	}
	return temp, nil
}
