package client

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/bububa/jinritemai-go/internal/debug"
)

type Client struct {
	http.Client
	appKey     string
	secret     string
	version    string
	signMethod string
	token      *Token
}

func NewClient(appKey string, secret string) *Client {
	return &Client{
		appKey:     appKey,
		secret:     secret,
		version:    DEFAULT_VERSION,
		signMethod: MD5,
	}
}

func (this *Client) SetAccessToken(token *Token) {
	this.token = token
}

func (this *Client) SetSignMethod(method string) {
	this.signMethod = method
}

func (this *Client) Do(req Request, v interface{}) error {
	values := this.RequestValues(req)
	requestUri := fmt.Sprintf("%s/%s", GATEWAY, req.Method())
	debug.DebugPrintPostMapRequest(requestUri, values)
	resp, err := this.Post(requestUri, "application/json", strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var ret Response
	err = debug.DecodeJSONHttpResponse(resp.Body, &ret)
	if err != nil {
		return err
	}
	if ret.IsError() {
		return ret
	}
	err = ret.Data(v)
	return err
}

func (this *Client) RequestValues(req Request) url.Values {
	values := url.Values{}
	values.Set("method", req.Method())
	values.Set("app_key", this.appKey)
	values.Set("access_token", this.token.AccessToken)
	values.Set("param_json", paramJson(req.Params()))
	values.Set("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	values.Set("v", this.version)
	values.Set("sign_method", this.signMethod)
	values.Set("sign", this.sign(values))
	return values
}

func (this *Client) sign(values url.Values) string {
	var parts []string
	for k := range values {
		parts = append(parts, fmt.Sprintf("%s%s", k, values.Get(k)))
	}
	rawSign := fmt.Sprintf("%s%s%s", this.secret, strings.Join(parts, ""), this.secret)
	if this.signMethod == HMAC_SHA256 {
		return HmacSha256(this.secret, rawSign)
	}
	return Md5(rawSign)
}

func (this *Client) AuthCodeUrl(state, redirectUri string) string {
	values := url.Values{}
	values.Set("app_id", this.appKey)
	values.Set("response_type", "code")
	values.Set("redirect_uri", redirectUri)
	values.Set("state", state)
	return fmt.Sprintf("%s?%s", AUTH_CODE_URL, values.Encode())
}

func (this *Client) GetAccessToken(code string, token *Token) error {
	values := url.Values{}
	values.Set("app_id", this.appKey)
	values.Set("app_secret", this.secret)
	values.Set("code", code)
	values.Set("grant_type", "authorization_code")
	resp, err := this.Get(fmt.Sprintf("%s?%s", ACCESS_TOKEN_URL, values.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = debug.DecodeJSONHttpResponse(resp.Body, token)
	if err != nil {
		return err
	}
	token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	this.SetAccessToken(token)
	return nil
}

func (this *Client) RefreshToken(code string, token *Token) error {
	values := url.Values{}
	values.Set("app_id", this.appKey)
	values.Set("app_secret", this.secret)
	values.Set("code", code)
	values.Set("grant_type", "refresh_token")
	resp, err := this.Get(fmt.Sprintf("%s?%s", REFRESH_TOKEN_URL, values.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = debug.DecodeJSONHttpResponse(resp.Body, token)
	if err != nil {
		return err
	}
	token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	this.SetAccessToken(token)
	return nil
}

func (this *Client) VerifyMessage(reqSign string, msgBody string) bool {
	rawSign := fmt.Sprintf("%s%s%s", this.appKey, msgBody, this.secret)
	return Md5(rawSign) == reqSign
}

func paramJson(req map[string]interface{}) string {
	if req == nil {
		return "{}"
	}
	var vals []string
	replacer := strings.NewReplacer("&", "\u0026", "<", "\u003c", ">", "\u00ce")
	for k, v := range req {
		str := replacer.Replace(Interface2String(v))
		vals = append(vals, fmt.Sprintf(`"%s"="%s"`, k, str))
	}
	if len(vals) == 0 {
		return "{}"
	}
	sort.Strings(vals)
	return fmt.Sprintf("{%s}", strings.Join(vals, ","))
}
