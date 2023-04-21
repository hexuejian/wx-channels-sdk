package callbacks

import "github.com/tidwall/gjson"

// 优惠券领取通知
// 文档: https://developers.weixin.qq.com/doc/channels/API/coupon/ec_callback/channels_ec_coupon_receive.html

func init() {
	//添加可解析的回调事件
	supportCallback(ChannelsEcCouponReceive{})
}

type ChannelsEcCouponReceive struct {
	CreateTime   int64  `json:"CreateTime"`
	Event        string `json:"Event"`
	FromUserName string `json:"FromUserName"`
	MsgType      string `json:"MsgType"`
	ToUserName   string `json:"ToUserName"`
	ReceiveInfo  struct {
		CouponID     int64 `json:"coupon_id"`
		ReceiveTime  int64 `json:"receive_time"`
		UserCouponID int64 `json:"user_coupon_id"`
	} `json:"receive_info"`
}

func (ChannelsEcCouponReceive) GetMessageType() string {
	return "event"
}

func (ChannelsEcCouponReceive) GetEventType() string {
	return "channels_ec_coupon_receive"
}

func (m ChannelsEcCouponReceive) GetTypeKey() string {
	return m.GetMessageType() + ":" + m.GetEventType()
}

func (ChannelsEcCouponReceive) ParseFromJson(data []byte) (CallbackExtraInfoInterface, error) {
	var temp = ChannelsEcCouponReceive{
		CreateTime:   gjson.GetBytes(data, "CreateTime").Int(),
		Event:        gjson.GetBytes(data, "Event").String(),
		FromUserName: gjson.GetBytes(data, "FromUserName").String(),
		MsgType:      gjson.GetBytes(data, "MsgType").String(),
		ToUserName:   gjson.GetBytes(data, "ToUserName").String(),
		ReceiveInfo: struct {
			CouponID     int64 `json:"coupon_id"`
			ReceiveTime  int64 `json:"receive_time"`
			UserCouponID int64 `json:"user_coupon_id"`
		}{
			CouponID:     gjson.GetBytes(data, "receive_info.coupon_id").Int(),
			ReceiveTime:  gjson.GetBytes(data, "receive_info.receive_time").Int(),
			UserCouponID: gjson.GetBytes(data, "receive_info.user_coupon_id").Int(),
		},
	}
	return temp, nil
}
