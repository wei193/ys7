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

//Ys7 萤石协议
type Ys7Protocol struct {
	Model		string `json:"model"`
	Version 	string `json:"version"`
	IsSupport	string `json:"isSupport"`
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

//DeviceCapacity 萤石设备能力集
type DeviceCapacity struct {
	Support_cloud					string `json:"support_cloud"`
	Support_intelligent_track		string `json:"support_intelligent_track"`
	Support_p2p_mode				string `json:"support_p_2_p_mode"`
	Support_resolution				string `json:"support_resolution"`
	Support_talk					string `json:"support_talk"`
	Support_wifi_userId				string `json:"support_wifi_user_id"`
	Support_remote_auth_randcode	string `json:"support_remote_auth_randcode"`
	Support_upgrade					string `json:"support_upgrade"`
	Support_smart_wifi				string `json:"support_smart_wifi"`
	Support_ssl						string `json:"support_ssl"`
	Support_weixin					string `json:"support_weixin"`
	Ptz_close_scene					string `json:"ptz_close_scene"`
	Support_preset_alarm			string `json:"support_preset_alarm"`
	Support_related_device			string `json:"support_related_device"`
	Support_message					string `json:"support_message"`
	Ptz_preset						string `json:"ptz_preset"`
	Support_wifi					string `json:"support_wifi"`
	Support_cloud_version			string `json:"support_cloud_version"`
	Ptz_center_mirror				string `json:"ptz_center_mirror"`
	Support_defence					string `json:"support_defence"`
	Ptz_top_bottom					string `json:"ptz_top_bottom"`
	Support_fullscreen_ptz			string `json:"support_fullscreen_ptz"`
	Support_defenceplan				string `json:"support_defenceplan"`
	Support_disk					string `json:"support_disk"`
	Support_alarm_voice				string `json:"support_alarm_voice"`
	Ptz_left_right					string `json:"ptz_left_right"`
	Support_modify_pwd				string `json:"support_modify_pwd"`
	Support_capture					string `json:"support_capture"`
	Support_privacy					string `json:"support_privacy"`
	Support_encrypt					string `json:"support_encrypt"`
	Support_auto_offline			string `json:"support_auto_offline"`
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

// 图片
type Picture struct {
	PicUrl		string `json:"picUrl"`
}

// 设备状态信息
type DeviceStatusInfo struct {
	PrivacyStatus	int `json:"privacyStatus"`
	PirStatus		int `json:"pirStatus"`
	AlarmSoundMode  int `json:"alarmSoundMode"`
	BattryStatus	int `json:"battryStatus"`
	LockSignal		int `json:"lockSignal"`
	DiskNum			int `json:"diskNum"`
	DiskState		string `json:"diskState"`
	CloudType		int `json:"cloudType"`
	CloudStatus		int `json:"cloudStatus"`
	NvrDiskNum		int `json:"nvrDiskNum"`
	NvrDiskState	string `json:"nvrDiskState"`
}

// SwitchStatus 开关状态
type SwitchStatus struct {
	DeviceSerial string `json:"deviceSerial"`
	ChannelNo    int    `json:"channelNo"`
	Enable		 int	`json:"enable"`
}