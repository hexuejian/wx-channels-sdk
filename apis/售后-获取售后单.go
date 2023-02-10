package apis

import (
	"encoding/json"
)

// 获取售后单
// 文档：https://developers.weixin.qq.com/doc/channels/API/aftersale/getaftersaleorder.html

type ReqAftersaleGetaftersaleorder struct {
	AfterSaleOrderID string `json:"after_sale_order_id"`
}

var _ bodyer = ReqAftersaleGetaftersaleorder{}

func (x ReqAftersaleGetaftersaleorder) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespAftersaleGetaftersaleorder struct {
	AfterSaleOrder struct {
		AfterSaleOrderID string `json:"after_sale_order_id"`
		CreateTime       int    `json:"create_time"`
		Details          struct {
			CancelTime     int           `json:"cancel_time"`
			Desc           string        `json:"desc"`
			ProveImgs      []interface{} `json:"prove_imgs"`
			ReceiveProduct int           `json:"receive_product"`
			TelNumber      string        `json:"tel_number"`
		} `json:"details"`
		MerchantUploadInfo struct {
			RefundCertificates []interface{} `json:"refund_certificates"`
			RejectReason       string        `json:"reject_reason"`
		} `json:"merchant_upload_info"`
		Openid      string `json:"openid"`
		OrderID     string `json:"order_id"`
		ProductInfo struct {
			Count     int    `json:"count"`
			ProductID string `json:"product_id"`
			SkuID     string `json:"sku_id"`
		} `json:"product_info"`
		Reason     string `json:"reason"`
		RefundInfo struct {
			Amount int `json:"amount"`
		} `json:"refund_info"`
		RefundResp struct {
			Code    string `json:"code"`
			Message string `json:"message"`
			Ret     int    `json:"ret"`
		} `json:"refund_resp"`
		ReturnInfo struct {
			DeliveryID   string `json:"delivery_id"`
			DeliveryName string `json:"delivery_name"`
			WaybillID    string `json:"waybill_id"`
		} `json:"return_info"`
		Status     string `json:"status"`
		Type       string `json:"type"`
		UpdateTime int    `json:"update_time"`
	} `json:"after_sale_order"`
	CommonResp
}

var _ bodyer = RespAftersaleGetaftersaleorder{}

func (x RespAftersaleGetaftersaleorder) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecAftersaleGetaftersaleorder(req ReqAftersaleGetaftersaleorder) (RespAftersaleGetaftersaleorder, error) {
	var resp RespAftersaleGetaftersaleorder
	err := c.executeWXApiPost("/channels/ec/aftersale/getaftersaleorder", req, &resp, true)
	if err != nil {
		return RespAftersaleGetaftersaleorder{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAftersaleGetaftersaleorder{}, bizErr
	}
	return resp, nil
}
