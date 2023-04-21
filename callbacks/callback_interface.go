package callbacks

var callbackParseExtraInfoMap = make(map[string]CallbackExtraInfoInterface)

type CallbackExtraInfoInterface interface {
	GetMessageType() string
	GetEventType() string
	GetTypeKey() string
	ParseFromJson(data []byte) (CallbackExtraInfoInterface, error)
}

func supportCallback(item CallbackExtraInfoInterface) {
	callbackParseExtraInfoMap[item.GetTypeKey()] = item
}

func SupportSdkCallback() {
	item1 := ChannelsEcCouponInvalid{}
	callbackParseExtraInfoMap[item1.GetTypeKey()] = item1
	item2 := ChannelsEcCouponDelete{}
	callbackParseExtraInfoMap[item2.GetTypeKey()] = item2
	item3 := ChannelsEcUserCouponUse{}
	callbackParseExtraInfoMap[item3.GetTypeKey()] = item3
	item4 := ChannelsEcCouponExpire{}
	callbackParseExtraInfoMap[item4.GetTypeKey()] = item4
	item5 := ChannelsEcUserCouponUnuse{}
	callbackParseExtraInfoMap[item5.GetTypeKey()] = item5
	item6 := ChannelsEcCouponReceive{}
	callbackParseExtraInfoMap[item6.GetTypeKey()] = item6
	item7 := ChannelsEcCouponCreate{}
	callbackParseExtraInfoMap[item7.GetTypeKey()] = item7
	item8 := ChannelsEcCouponInfoChange{}
	callbackParseExtraInfoMap[item8.GetTypeKey()] = item8
	item9 := ChannelsEcUserCouponExpire{}
	callbackParseExtraInfoMap[item9.GetTypeKey()] = item9
}
