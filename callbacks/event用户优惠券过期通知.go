package callbacks

import (
	"github.com/tidwall/gjson"
)

// 用户优惠券过期通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_user_coupon_expire.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcUserCouponExpire{})
}

type ChannelsEcUserCouponExpire struct {
	CreateTime     int64  `json:"CreateTime"`
	Event          string `json:"Event"`
	FromUserName   string `json:"FromUserName"`
	MsgType        string `json:"MsgType"`
	ToUserName     string `json:"ToUserName"`
	UserCouponInfo struct {
		CouponID     string `json:"coupon_id"`
		ExpireTime   int64  `json:"expire_time"`
		UserCouponID string `json:"user_coupon_id"`
	} `json:"user_coupon_info"`
}

func (ChannelsEcUserCouponExpire) GetMessageType() string {
	return "event"
}

func (ChannelsEcUserCouponExpire) GetEventType() string {
	return "channels_ec_user_coupon_expire"
}

func (m ChannelsEcUserCouponExpire) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcUserCouponExpire) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcUserCouponExpire{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		UserCouponInfo: struct {
			CouponID     string `json:"coupon_id"`
			ExpireTime   int64  `json:"expire_time"`
			UserCouponID string `json:"user_coupon_id"`
		}{
			CouponID:     gjson.GetBytes(data, "user_coupon_info.coupon_id").String(),
			ExpireTime:   gjson.GetBytes(data, "user_coupon_info.expire_time").Int(),
			UserCouponID: gjson.GetBytes(data, "user_coupon_info.user_coupon_id").String(),
		},
	}
	return temp, nil
}
