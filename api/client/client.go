package client

import (
	"fmt"
	"github.com/meow-pad/joytalk-helpers/api"
	"github.com/meow-pad/joytalk-helpers/api/familyapi"
	"github.com/meow-pad/joytalk-helpers/api/payapi"
	"github.com/meow-pad/joytalk-helpers/api/userapi"
	"github.com/meow-pad/joytalk-helpers/utils/jwt"
	"github.com/meow-pad/persian/utils/json"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

var ErrLessRequestUri = fmt.Errorf("request uri is empty")

func NewClient(appId, secret string, opts ...Option) (*Client, error) {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}
	if err := options.complete(); err != nil {
		return nil, err
	}
	return &Client{
		inner:      &fasthttp.Client{},
		appId:      appId,
		secret:     secret,
		jwtHeader:  jwt.BuildBase64JoytalkJWTHeader(jwt.NewJoytalkJWTHeader()),
		requestUri: options.RequestUri,
		payReqUri:  options.PayReqUri,
	}, nil
}

type Client struct {
	inner *fasthttp.Client

	appId     string
	secret    string
	jwtHeader string

	requestUri string
	payReqUri  string
}

func request[RespData any](client *Client, requestUri string, reqMsg any,
	handler func(err error, data *RespData), timeout time.Duration) {
	expDuration := int64(600)
	nowSec := time.Now().Unix()
	expSec := nowSec + expDuration
	if rReq, _ := reqMsg.(api.Request); rReq != nil {
		rReq.SetIat(nowSec)
		rReq.SetExp(expSec)
	}
	// 构建签名
	payload := jwt.BuildBase64JoytalkJWTPayload(&jwt.JoyTalkJWTPayload{
		Iat:   nowSec,
		Exp:   expSec,
		AppId: client.appId,
	})
	token := jwt.BuildJoytalkToken(client.jwtHeader, payload, []byte(client.secret))
	// 构建请求
	req := &fasthttp.Request{}
	req.SetRequestURI(requestUri)
	req.Header.SetMethod("POST")
	req.Header.Set("AppID", client.appId)
	req.Header.Set("Token", token)
	req.Header.Set("Content-Type", api.JsonContentType)
	req.SetBody([]byte(json.ToString(reqMsg)))
	// 发起请求
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := client.inner.DoTimeout(req, resp, timeout)
	if err != nil {
		handler(err, nil)
		return
	}

	if resp.StatusCode() != http.StatusOK {
		handler(fmt.Errorf("request failed, status code: %d", resp.StatusCode()), nil)
		return
	} else {
		var respMsg api.Response[RespData]
		if err = json.Unmarshal(resp.Body(), &respMsg); err != nil {
			handler(err, nil)
			return
		}
		if respMsg.ErrCode != api.ErrCodeSuccess {
			handler(fmt.Errorf("request failed, bizcode: %d", respMsg.ErrCode), nil)
			return
		}
		handler(nil, &respMsg.Data)
	}
}

func (client *Client) BatchGetUsers(userIds []string,
	handler func(err error, data *userapi.BatchGetUserData), timeout time.Duration) {
	if len(client.requestUri) <= 0 {
		handler(ErrLessRequestUri, nil)
		return
	}
	if len(userIds) > api.MaxRequestElemNum {
		userIds = userIds[:api.MaxRequestElemNum]
	}
	requestUri := client.requestUri + userapi.BatchGetUserPath
	request[userapi.BatchGetUserData](client, requestUri,
		&userapi.BatchGetUserRequest{UserIds: userIds}, handler, timeout)
}

func (client *Client) BatchGetFamilies(familyIds []string,
	handler func(err error, data *familyapi.BatchGetFamilyData), timeout time.Duration) {
	if len(client.requestUri) <= 0 {
		handler(ErrLessRequestUri, nil)
		return
	}
	if len(familyIds) > api.MaxRequestElemNum {
		familyIds = familyIds[:api.MaxRequestElemNum]
	}
	requestUri := client.requestUri + familyapi.BatchGetClanPath
	request[familyapi.BatchGetFamilyData](client, requestUri,
		&familyapi.BatchGetFamilyRequest{FamilyIds: familyIds}, handler, timeout)
}

func (client *Client) BatchGetUserFamilyInfo(userIds []string,
	handler func(err error, data *userapi.BatchGetUserFamilyInfoData), timeout time.Duration) {
	if len(client.requestUri) <= 0 {
		handler(ErrLessRequestUri, nil)
		return
	}
	if len(userIds) > api.MaxRequestElemNum {
		userIds = userIds[:api.MaxRequestElemNum]
	}
	requestUri := client.requestUri + userapi.BatchGetUserFamilyInfoPath
	request[userapi.BatchGetUserFamilyInfoData](client,
		requestUri, &userapi.BatchGetUserFamilyInfoRequest{UserIds: userIds}, handler, timeout)
}

func (client *Client) OrderConsume(consumeReq *payapi.OrderConsumeRequest,
	handler func(err error, _ *any), timeout time.Duration) {
	if len(client.payReqUri) <= 0 {
		handler(ErrLessRequestUri, nil)
		return
	}
	requestUri := client.payReqUri + payapi.OrderConsumerPath
	request[any](client, requestUri, consumeReq, handler, timeout)
}

func (client *Client) OrderDetail(consumeReq *payapi.OrderDetailRequest,
	handler func(err error, data *payapi.OrderDetailData), timeout time.Duration) {
	if len(client.payReqUri) <= 0 {
		handler(ErrLessRequestUri, nil)
		return
	}
	requestUri := client.payReqUri + payapi.OrderDetailPath
	request[payapi.OrderDetailData](client, requestUri, consumeReq, handler, timeout)
}

func (client *Client) UserBalance(balanceReq *payapi.BalanceRequest,
	handler func(err error, data *payapi.BalanceData), timeout time.Duration) {
	if len(client.payReqUri) <= 0 {
		handler(ErrLessRequestUri, nil)
		return
	}
	requestUri := client.payReqUri + payapi.BalancePath
	request[payapi.BalanceData](client, requestUri, balanceReq, handler, timeout)
}
