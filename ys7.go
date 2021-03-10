package ys7

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//URL
const (
	MASTERACC = 0 //主账号
	RAMACC    = 1 //子账号

	//[用户]获取accessToken
	ACCESSTOKEN = "https://open.ys7.com/api/lapp/token/get" //获取accessToken

	//[用户]好友分享
	// --全部待实现

)

//NewYs7 创建Ys7对象
func NewYs7(AppKey, Secret string) (ys *Ys7, err error) {
	ys = &Ys7{
		AppKey: AppKey,
		Secret: Secret,
		IsRAM:  MASTERACC,
	}
	_, err = ys.GetAccessToken()
	return
}

//NewRAMYs7 创建子账号对象
func NewRAMYs7(AppKey, Secret, AccountID string) (ys *Ys7, err error) {
	ys = &Ys7{
		AppKey:    AppKey,
		Secret:    Secret,
		IsRAM:     RAMACC,
		AccountID: AccountID,
	}
	_, err = ys.GetAccessToken()
	return
}

//GetAccessToken 获取token
func (ys *Ys7) GetAccessToken() (ac *AccessToken, err error) {
	params := make(map[string]interface{})
	params["appKey"] = ys.AppKey
	params["appSecret"] = ys.Secret
	ac = &AccessToken{}
	_, err = ys.requset("POST", ACCESSTOKEN, params, &ac)
	if err != nil {
		return nil, err
	}
	if ys.IsRAM == MASTERACC {
		ys.AccessToken = ac.AccessToken
		ys.ExpireTime = ac.ExpireTime
	} else {

		ys.AccessToken = ac.AccessToken
		ac, err = ys.RAMGetAccessToken(ys.AccountID)
		if err != nil {
			ys.AccessToken = ""
			return
		}
		ys.AccessToken = ac.AccessToken
		ys.ExpireTime = ac.ExpireTime
	}
	return ac, nil
}

func (ys *Ys7) requset(method, url string, params map[string]interface{}, data interface{}, page ...interface{}) (status *Status, err error) {
	defer func() {
		if Rerr := recover(); Rerr != nil {
			err = errors.New("recover error")
			return
		}
	}()
	var r http.Request
	r.ParseForm()
	for k, v := range params {
		r.Form.Add(k, fmt.Sprint(v))
	}
	req, err := http.NewRequest(method, url, strings.NewReader(r.Form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res respStatus
	if len(page) == 1 {
		res = respStatus{
			Data: data,
			Page: page[0],
		}
	} else {
		res = respStatus{
			Data: data,
		}
	}

	err = json.Unmarshal(buf, &res)
	if err != nil {
		return nil, err
	}
	status = &Status{
		Code: res.Code,
		Msg:  res.Msg,
		Buf:  buf,
	}
	if res.Code != "200" {
		return status, errors.New(res.Msg)
	}
	return status, nil
}

func (ys *Ys7) authorizeRequset(method, url string, params map[string]interface{}, data interface{}, page ...interface{}) (status *Status, err error) {
	exTime := time.Unix(ys.ExpireTime/1000, 0)
	if exTime.Unix() < time.Now().Unix() {
		ys.GetAccessToken()
	}
	defer func() {
		if Rerr := recover(); Rerr != nil {
			err = errors.New("recover error")
			return
		}
	}()
	params["accessToken"] = ys.AccessToken
	status, err = ys.requset(method, url, params, data, page...)
	return
}

func getPage(buf []byte) Page {
	type st struct {
		Page Page
	}
	var gPage st
	json.Unmarshal(buf, &gPage)
	return gPage.Page
}

func getPasswd(appkey, password string) string {
	hash := md5.New()
	hash.Write([]byte(appkey + "#" + password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
