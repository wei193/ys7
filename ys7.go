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

//Ys7 萤石接口
type Ys7 struct {
	AppKey      string
	Secret      string
	AccessToken string
	ExpireTime  int64
	IsRAM       int
	AccountID   string
	Password    string
}

type respStatus struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//AccessToken 萤石密钥
type AccessToken struct {
	AccessToken string `json:"accessToken"`
	ExpireTime  int64  `json:"expireTime"`
}

//Device 萤石设备数据结构
type Device struct {
	DeviceSerial  string `json:"deviceSerial"`
	DeviceName    string `json:"deviceName"`
	DeviceType    string `json:"deviceType"`
	Status        int    `json:"status"`
	Defence       int    `json:"defence"`
	DeviceVersion string `json:"deviceVersion"`
}

//DeviceInfo 萤石设备数据结构
type DeviceInfo struct {
	DeviceSerial   string `json:"deviceSerial"`
	DeviceName     string `json:"deviceName"`
	Model          string `json:"model"`
	Status         int    `json:"status"`
	Defence        int    `json:"defence"`
	IsEncrypt      int    `json:"isEncrypt"`
	AlarmSoundMode int    `json:"alarmSoundMode"`
	OfflineNotify  int    `json:"offlineNotify"`
}

//Camera 萤石摄像头数据结构
type Camera struct {
	DeviceSerial string `json:"deviceSerial"`
	ChannelNo    int    `json:"channelNo"`
	ChannelName  string `json:"channelName"`
	PicURL       string `json:"picUrl"`
	IsShared     string `json:"isShared"`
	VideoLevel   int    `json:"videoLevel"`
	IsEncrypt    int    `json:"isEncrypt"`
	Status       int    `json:"status"`
}

//Account 萤石子账号ID
type Account struct {
	AccountID string `json:"accountId"`
}

//RAMAccount 萤石子账号
type RAMAccount struct {
	AccountID     string `json:"accountId"`
	AccountName   string `json:"accountName"`
	AppKey        string `json:"appKey"`
	AccountStatus int    `json:"accountStatus"`
	Policy        Policy `json:"policy"`
}

// Policy Policy
type Policy struct {
	Statement []Statement `json:"Statement"`
}

//Statement Statement
type Statement struct {
	Permission string   `json:"Permission"`
	Resource   []string `json:"Resource"`
}

//Page 分页数据
type Page struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Size  int `json:"size"`
}

//URL
const (
	MASTERACC = 0 //主账号
	RAMACC    = 1 //子账号

	//[用户]获取accessToken
	ACCESSTOKEN = "https://open.ys7.com/api/lapp/token/get" //获取accessToken

	//[用户]好友分享
	// --全部待实现

	//[直播管理]直播接口
	// URLVIDEOLIST  = "https://open.ys7.com/api/lapp/live/video/list"
	LIVEADDRESSLIMITED = "https://open.ys7.com/api/lapp/live/address/limited" //获取指定有效期的直播地址
	// URLVIDEOOPEN  = "https://open.ys7.com/api/lapp/live/video/open"
	// URLVIDEOCLOSE = "https://open.ys7.com/api/lapp/live/video/close"
	// URLLIVEGET    = "https://open.ys7.com/api/lapp/live/address/get"

	//[设备]管理
	DEVICEADD     = "https://open.ys7.com/api/lapp/device/add"         //添加设备到账号下
	DEVICEDELETE  = "https://open.ys7.com/api/lapp/device/delete"      //删除账号下指定设备
	DEVICEUPDATE  = "https://open.ys7.com/api/lapp/device/name/update" //修改设备名称
	DEVICECAPTURE = "https://open.ys7.com/api/lapp/device/capture"     //抓拍设备的当前画面
	URLIPCADD     = "https://open.ys7.com/api/lapp/device/ipc/add"     //NVR设备关联IPC
	URLIPCDELETE  = "hhttps://open.ys7.com/api/lapp/device/ipc/delete" //NVR设备删除IPC
	//修改设备视频加密密码
	//生成设备扫描配网二维码二进制数据
	CAMERANAMEUPDATE = "https://open.ys7.com/api/lapp/camera/name/update" // 修改通道名称

	//[设备]查询
	DEVICELIST = "https://open.ys7.com/api/lapp/device/list" //获取用户下的设备列表
	DEVICEINFO = "https://open.ys7.com/api/lapp/device/info" //获取指定设备的信息
	CAMERALIST = "https://open.ys7.com/api/lapp/camera/list" //获取用户下的摄像头列表
	//设备互联互通根据UUID查询抓拍的图片
	//根据序列号获取设备的状态信息
	//DEVICECAMERALIST = "https://open.ys7.com/api/lapp/device/camera/list"  //根据序列号获取设备的通道信息
	//根据设备型号以及设备版本号查询设备是否支持萤石协议
	//根据时间获取录像信息

	//[设备]配置
	// --全部待实现

	//[设备]升级
	// --全部待实现

	// [设备]云台
	URLPTZSTAR      = "https://open.ys7.com/api/lapp/device/ptz/start"
	URLPTZSTOP      = "https://open.ys7.com/api/lapp/device/ptz/stop"
	URLPTZMIRROR    = "https://open.ys7.com/api/lapp/device/ptz/mirror"
	URLPRESETADD    = "https://open.ys7.com/api/lapp/device/preset/add"
	URLPPRESETMOVE  = "https://open.ys7.com/api/lapp/device/preset/move"
	URLPPRESETCLEAR = "https://open.ys7.com/api/lapp/device/preset/clear"

	//[子账号]子账号接口
	ACCOUNTCREATE         = "https://open.ys7.com/api/lapp/ram/account/create"         //创建子账户
	ACCOUNTGET            = "https://open.ys7.com/api/lapp/ram/account/get"            //获取指定子账户信息
	ACCOUNTLIST           = "https://open.ys7.com/api/lapp/ram/account/list"           //分页获取应用下的子账户信息列表
	ACCOUNTUPDATEPASSWORD = "https://open.ys7.com/api/lapp/ram/account/updatePassword" //修改子账户密码
	POLICYSET             = "https://open.ys7.com/api/lapp/ram/policy/set"             //设置子账户的授权策略
	STATEMENTADD          = "https://open.ys7.com/api/lapp/ram/statement/add"          //增加子账户授权策略中的授权语句
	STATEMENTDELETE       = "https://open.ys7.com/api/lapp/ram/statement/delete"       //删除子账户授权策略中某个设备的所有授权语句
	RAMTOKENGET           = "https://open.ys7.com/api/lapp/ram/token/get"              //获取子账户AccessToken
	ACCOUNTDELETE         = "https://open.ys7.com/api/lapp/ram/account/delete"         //删除子账户

)

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
	ys.AccessToken = ac.AccessToken
	ys.ExpireTime = ac.ExpireTime
	return ac, nil
}

//GetLimited 获取直播流地址
func (ys *Ys7) GetLimited(deviceSerial string, channelNo, expireTime int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["expireTime"] = expireTime
	buf, err := ys.authorizeRequset("POST", LIVEADDRESSLIMITED, params, nil)
	if err != nil {
		return err
	}
	fmt.Println(4, string(buf))
	return
}

//AddDevice 添加设备
func (ys *Ys7) AddDevice(deviceSerial, validateCode string) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["validateCode"] = validateCode
	_, err = ys.authorizeRequset("POST", DEVICEADD, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//DeleteDevice 删除设备
func (ys *Ys7) DeleteDevice(deviceSerial string) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	_, err = ys.authorizeRequset("POST", DEVICEDELETE, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//UpdateDeviceName 修改通道名称
func (ys *Ys7) UpdateDeviceName(deviceSerial, deviceName string) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["deviceName"] = deviceName

	_, err = ys.authorizeRequset("POST", DEVICEUPDATE, params, nil)
	if err != nil {
		return
	}
	return nil
}

// GetDeviceCapture 设备抓拍图片
func (ys *Ys7) GetDeviceCapture(deviceSerial string, channelNo int) (picURL string, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	type st struct {
		PicURL string `json:"picUrl"`
	}
	var in st
	_, err = ys.authorizeRequset("POST", DEVICECAPTURE, params, &in)
	if err != nil {
		return
	}
	return in.PicURL, nil
}

//AddDeviceIpc NVR设备关联IPC
func (ys *Ys7) AddDeviceIpc(deviceSerial, ipcSerial string, channelNo int, validateCode string) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["ipcSerial"] = ipcSerial
	params["channelNo"] = channelNo
	params["validateCode"] = validateCode
	_, err = ys.authorizeRequset("POST", URLIPCADD, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//DeleteDeviceIpc NVR设备删除IPC
func (ys *Ys7) DeleteDeviceIpc(deviceSerial, ipcSerial string, channelNo int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["ipcSerial"] = ipcSerial
	params["channelNo"] = channelNo
	_, err = ys.authorizeRequset("POST", URLIPCDELETE, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//GetDeviceList 获取设备列表
func (ys *Ys7) GetDeviceList(pageStart, pageSize int) (devices []Device, page Page, err error) {
	params := make(map[string]interface{})
	params["pageStart"] = pageStart
	params["pageSize"] = pageSize
	buf, err := ys.authorizeRequset("POST", DEVICELIST, params, &devices) //获取用户下的设备列表
	if err != nil {
		return nil, page, err
	}
	return devices, getPage(buf), nil
}

//InfoDevice 获取单个设备信息
func (ys *Ys7) InfoDevice(deviceSerial string) (deviceinfo DeviceInfo, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	_, err = ys.authorizeRequset("POST", DEVICEINFO, params, &deviceinfo)
	if err != nil {
		return
	}
	return
}

//GetCameraList 获取摄像头列表
func (ys *Ys7) GetCameraList(pageStart, pageSize int) (cameras []Camera, page Page, err error) {
	params := make(map[string]interface{})
	params["pageStart"] = pageStart
	params["pageSize"] = pageSize
	buf, err := ys.authorizeRequset("POST", CAMERALIST, params, &cameras)
	if err != nil {
		return nil, page, err
	}
	return cameras, getPage(buf), nil
}

//UpdateCameraName 修改通道名称
func (ys *Ys7) UpdateCameraName(deviceSerial, name string, channelNo int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["name"] = name
	params["channelNo"] = channelNo

	_, err = ys.authorizeRequset("POST", CAMERANAMEUPDATE, params, nil)
	if err != nil {
		return
	}
	return nil
}

//StartPtz 开始云台控制
func (ys *Ys7) StartPtz(deviceSerial string, channelNo, direction, speed int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["direction"] = direction
	params["speed"] = speed

	_, err = ys.authorizeRequset("POST", URLPTZSTAR, params, nil)
	if err != nil {
		return
	}
	return nil
}

//StopPtz 停止云台转动
func (ys *Ys7) StopPtz(deviceSerial string, channelNo, direction int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["direction"] = direction

	_, err = ys.authorizeRequset("POST", URLPTZSTOP, params, nil)
	if err != nil {
		return
	}
	return nil
}

//MirrorPtz 镜像翻转
func (ys *Ys7) MirrorPtz(deviceSerial string, channelNo, command int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["command"] = command

	_, err = ys.authorizeRequset("POST", URLPTZMIRROR, params, nil)
	if err != nil {
		return
	}
	return nil
}

//AddPreset 添加预置点
func (ys *Ys7) AddPreset(deviceSerial string, channelNo int) (index int, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo

	type st struct {
		Index int `json:"index"`
	}
	var data st
	_, err = ys.authorizeRequset("POST", URLPRESETADD, params, &data)
	if err != nil {
		return
	}
	return data.Index, nil
}

//MovePreset 调用预置点
func (ys *Ys7) MovePreset(deviceSerial string, channelNo, index int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["index"] = index

	_, err = ys.authorizeRequset("POST", URLPPRESETMOVE, params, nil)
	if err != nil {
		return
	}
	return nil
}

//ClearPreset 清除预置点
func (ys *Ys7) ClearPreset(deviceSerial string, channelNo, index int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["index"] = index

	_, err = ys.authorizeRequset("POST", URLPPRESETCLEAR, params, nil)
	if err != nil {
		return
	}
	return nil
}

//CreateAccount 创建子账号
func (ys *Ys7) CreateAccount(accountName, password string) (acc *Account, err error) {
	params := make(map[string]interface{})
	params["accountName"] = accountName
	params["password"] = getPasswd(ys.AppKey, password)
	acc = &Account{}
	_, err = ys.authorizeRequset("POST", ACCOUNTCREATE, params, &acc)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

//CreateAccountAndToken 创建子账号并返回子账号对象
func (ys *Ys7) CreateAccountAndToken(accountName, password string) (y *Ys7, err error) {
	params := make(map[string]interface{})
	params["accountName"] = accountName
	params["password"] = getPasswd(ys.AppKey, password)
	acc := &Account{}
	_, err = ys.authorizeRequset("POST", ACCOUNTCREATE, params, &acc)
	if err != nil {
		return nil, err
	}
	ac, err := ys.RAMGetAccessToken(acc.AccountID)
	if err != nil {
		return y, err
	}
	y = &Ys7{
		AppKey:      ys.AppKey,
		AccessToken: ac.AccessToken,
		ExpireTime:  ac.ExpireTime,
		IsRAM:       RAMACC,
		Password:    password,
		AccountID:   acc.AccountID,
	}
	return y, nil
}

// RAMAccountGet  获取单个子账户信息
func (ys *Ys7) RAMAccountGet(accountID, accountName string) (acc *RAMAccount, err error) {
	params := make(map[string]interface{})
	if accountID != "" {
		params["accountId"] = accountID
	} else {
		params["accountName"] = accountName
	}
	acc = &RAMAccount{}
	_, err = ys.authorizeRequset("POST", ACCOUNTGET, params, &acc)
	if err != nil {
		return nil, err
	}
	return
}

// RAMAccountList 获取子账户信息列表
func (ys *Ys7) RAMAccountList(pageStart, pageSize int) (acc []RAMAccount, page Page, err error) {
	params := make(map[string]interface{})
	params["pageStart"] = pageStart
	params["pageSize"] = pageSize
	buf, err := ys.authorizeRequset("POST", ACCOUNTLIST, params, &acc)
	if err != nil {
		return nil, page, err
	}
	return acc, getPage(buf), nil
}

//RAMUpdatePassword 修改当前子账户密码
func (ys *Ys7) RAMUpdatePassword(accountID, old, new string) (err error) {
	params := make(map[string]interface{})
	params["accountId"] = accountID
	params["oldPassword"] = getPasswd(ys.AppKey, old)
	params["newPassword"] = getPasswd(ys.AppKey, new)
	_, err = ys.authorizeRequset("POST", ACCOUNTUPDATEPASSWORD, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//RAMSetSetPolicy 设置子账号权限
func (ys *Ys7) RAMSetSetPolicy(accountID string, policy Policy) (err error) {
	policyByte, err := json.Marshal(policy)
	if err != nil {
		return err
	}
	params := make(map[string]interface{})
	params["accountId"] = accountID
	params["policy"] = string(policyByte)
	_, err = ys.authorizeRequset("POST", POLICYSET, params, nil)
	if err != nil {
		return err
	}
	return nil
}

// RAMAddStatement 增加子账户权限
func (ys *Ys7) RAMAddStatement(accountID string, statement Statement) (err error) {
	StatementByte, err := json.Marshal(statement)
	if err != nil {
		return err
	}
	params := make(map[string]interface{})
	params["accountId"] = accountID
	params["statement"] = string(StatementByte)
	_, err = ys.authorizeRequset("POST", STATEMENTADD, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//RAMDeleteStatement 删除子账户权限
func (ys *Ys7) RAMDeleteStatement(accountID, deviceSerial string) (err error) {
	params := make(map[string]interface{})
	params["accountId"] = accountID
	params["deviceSerial"] = deviceSerial
	_, err = ys.authorizeRequset("POST", STATEMENTDELETE, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//RAMGetAccessToken 获取B模式子账户accessToken
func (ys *Ys7) RAMGetAccessToken(accountID string) (ac *AccessToken, err error) {
	params := make(map[string]interface{})
	params["accountId"] = accountID
	ac = &AccessToken{}
	_, err = ys.authorizeRequset("POST", RAMTOKENGET, params, &ac)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

// DeleteAccount  删除子账户
func (ys *Ys7) DeleteAccount(accountID string) (err error) {
	params := make(map[string]interface{})
	params["accountId"] = accountID
	_, err = ys.authorizeRequset("POST", ACCOUNTDELETE, params, nil)
	if err != nil {
		return err
	}
	return nil
}

func (ys *Ys7) requset(method, url string, params map[string]interface{}, data interface{}) (buf []byte, err error) {
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

	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return buf, err
	}
	if data != nil {

	}
	res := respStatus{
		Data: data,
	}
	err = json.Unmarshal(buf, &res)
	if err != nil {
		return buf, err
	}
	if res.Code != "200" {
		return buf, errors.New(res.Msg)
	}
	return buf, nil
}

func (ys *Ys7) authorizeRequset(method, url string, params map[string]interface{}, data interface{}) (buf []byte, err error) {
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
	var r http.Request
	r.ParseForm()
	r.Form.Add("accessToken", ys.AccessToken)
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

	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return buf, err
	}
	if data != nil {

	}
	res := respStatus{
		Data: data,
	}
	err = json.Unmarshal(buf, &res)
	if err != nil {
		return buf, err
	}
	if res.Code == "10002" {
		if _, err = ys.GetAccessToken(); err == nil {
			params["accessToken"] = ys.AccessToken
			return ys.requset(method, url, params, data)
		}
	}
	if res.Code != "200" {
		return buf, errors.New(res.Msg)
	}
	return buf, nil
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

