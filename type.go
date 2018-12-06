package ys7

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
	IpcSerial    string `json:"ipcSerial"`
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

type respStatus struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Page interface{} `json:"page"`
}

//Status 状态
type Status struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Buf  []byte `json:"-"`
}

//Live 直播信息
type Live struct {
	DeviceSerial string `json:"deviceSerial"`
	ChannelNo    int    `json:"channelNo"`
	DeviceName   string `json:"deviceName"`
	LiveAddress  string `json:"liveAddress,omitempty"`
	HdAddress    string `json:"hdAddress,omitempty"`
	Hls          string `json:"hls,omitempty"`
	HlsHd        string `json:"hlsHd,omitempty"`
	Rtmp         string `json:"rtmp,omitempty"`
	RtmpHd       string `json:"rtmpHd,omitempty"`
	Status       int    `json:"status"`
	Exception    int    `json:"exception"`
	BeginTime    int64  `json:"beginTime"`
	EndTime      int64  `json:"endTime"`
}

//LiveState 直播状态返回
type LiveState struct {
	DeviceSerial string `json:"deviceSerial"`
	ChannelNo    int    `json:"channelNo"`
	Ret          string `json:"ret"`
	Desc         string `json:"desc"`
}
