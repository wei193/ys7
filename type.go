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

//Protocol 萤石协议
type Protocol struct {
	Model     string `json:"model"`
	Version   string `json:"version"`
	IsSupport string `json:"isSupport"`
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
	SupportDefence            string `json:"support_defence"`              //是否支持布撤防,活动检测开关
	SupportTalk               string `json:"support_talk"`                 //是否支持对讲: 0-不支持, 1-全双工, 3-半双工
	SupportDefenceplan        string `json:"support_defenceplan"`          //是否支持布撤防计划 0-不支持， 1-支持,2-支持新的设备计划协议
	SupportDisk               string `json:"support_disk"`                 //是否支持存储格式化 0-不支持, 1-支持
	SupportPrivacy            string `json:"support_privacy"`              //是否支持隐私保护 0-不支持, 1-支持
	SupportMessage            string `json:"support_message"`              //是否支持留言 0-不支持, 1-支持
	SupportAlarmVoice         string `json:"support_alarm_voice"`          //是否支持告警声音配置 0-不支持, 1-支持
	SupportAutoOffline        string `json:"support_auto_offline"`         //是否支持设备自动上下线 0-不支持, 1-支持
	SupprotEncrypt            string `json:"supprot_encrypt"`              //是否支持视频图像加密 0-不支持, 1-支持
	SupportUpgrade            string `json:"support_upgrade"`              //是否支持设备升级 0-不支持, 1-支持
	SupportCloud              string `json:"support_cloud"`                //该设备型号是否支持云存储 0-不支持, 1-支持
	SupportCloudVersion       string `json:"support_cloud_version"`        //该设备版本是否支持云存储 0-不支持, 1-支持
	SupportWifi               string `json:"support_wifi"`                 //是否支持WI-FI:
	SupportCapture            string `json:"support_capture"`              //是否支持封面抓图: 0-不支持, 1-支持
	SupportModifyPwd          string `json:"support_modify_pwd"`           //是否支持修改设备加密密码: 0-不支持, 1-支持
	SupportResolution         string `json:"support_resolution"`           //视频播放比例 16-9表示16:9分辨率,默认16-9
	SupportMultiScreen        string `json:"support_multi_screen"`         //是否支持多画面播放 0-不支持, 1-支持(客户端使用,与设备无关)
	SupportUploadCloudFile    string `json:"support_upload_cloud_file"`    //是否支持手机拍照上传到云存储 `0-不支持，1-支持
	SupportAddDelDetector     string `json:"support_add_del_detector"`     //是否支持app远程添加删除外设(探测器): 0-不支持, 1-支持
	SupportIpcLink            string `json:"support_ipc_link"`             //是否支持IPC与A1联动关系设置: 0-不支持, 1-支持
	SupportWeixin             string `json:"support_weixin"`               //是否支持微信互联:0-不支持, 1-支持
	SupportSsl                string `json:"support_ssl"`                  //是否支持声源定位:0-不支持, 1-支持
	SupportRemoteAuthRandcode string `json:"support_remote_auth_randcode"` //是否支持设备远程授权获取密码, 0-不支持, 1-支持
	PtzTopBottom              string `json:"ptz_top_bottom"`               //是否支持云台上下转动 0-不支持, 1-支持
	PtzLeftRight              string `json:"ptz_left_right"`               //是否支持云台左右转动 0-不支持, 1-支持
	Ptz45                     string `json:"ptz_45"`                       //是否支持云台45度方向转动 0-不支持, 1-支持
	PtzZoom                   string `json:"ptz_zoom"`                     //是否支持云台缩放控制 0-不支持, 1-支持
	SupportPtz                string `json:"support_ptz"`                  //是否支持云台控制 0-不支持, 1-支持, 注:新设备的该能力集拆分为30-33这四个能力
	PtzPreset                 string `json:"ptz_preset"`                   //是否支持云台预置点 0-不支持, 1-支持
	PtzCommonCruise           string `json:"ptz_common_cruise"`            //是否支持普通巡航 0-不支持, 1-支持
	PtzFigureCruise           string `json:"ptz_figure_cruise"`            //是否支持花样巡航0-不支持, 1-支持
	PtzCenterMirror           string `json:"ptz_center_mirror"`            //是否支持中心镜像0-不支持, 1-支持
	PtzLeftRightMirror        string `json:"ptz_left_right_mirror"`        //是否支持左右镜像 0-不支持, 1-支持
	PtzTopBottomMirror        string `json:"ptz_top_bottom_mirror"`        //是否支持上下镜像 0-不支持, 1-支持
	PtzCloseScene             string `json:"ptz_close_scene"`              //是否支持关闭镜头 0-不支持, 1-支持
	SupportIntelligentTrack   string `json:"support_intelligent_track"`    //是否支持智能跟踪 0-不支持, 1-支持(C6B等云台摄像机支持)
	SupportP2pMode            string `json:"support_p2p_mode"`             //默认0，表示老的p2p协议；配置为1，表示该版本支持新的p2p协议
	SupportPresetAlarm        string `json:"support_preset_alarm"`         //是否支持预置点告警联动 0-不支持, 1-支持(C6B等云台摄像机支持)
	SupportRelatedDevice      string `json:"support_related_device"`       //是否支持关联设备 0-无关联设备, 1-关联监控点或N1, 2-关联探测器或A1, 3-关联监控点探测器或R1, 4关联多通道设备
	SupportFullscreenPtz      string `json:"support_fullscreen_ptz"`       //是否支持全景云台功能 0-不支持, 1-支持(C6B等云台摄像机支持).如存在能力集support_fullscreen_ptz_12(序号82),则优先参考能力集support_fullscreen_ptz_12
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

//Picture 图片
type Picture struct {
	PicURL string `json:"picUrl"`
}

//DeviceStatusInfo 设备状态信息
type DeviceStatusInfo struct {
	PrivacyStatus  int    `json:"privacyStatus"`
	PirStatus      int    `json:"pirStatus"`
	AlarmSoundMode int    `json:"alarmSoundMode"`
	BattryStatus   int    `json:"battryStatus"`
	LockSignal     int    `json:"lockSignal"`
	DiskNum        int    `json:"diskNum"`
	DiskState      string `json:"diskState"`
	CloudType      int    `json:"cloudType"`
	CloudStatus    int    `json:"cloudStatus"`
	NvrDiskNum     int    `json:"nvrDiskNum"`
	NvrDiskState   string `json:"nvrDiskState"`
}

// SwitchStatus 开关状态
type SwitchStatus struct {
	DeviceSerial string `json:"deviceSerial"`
	ChannelNo    int    `json:"channelNo"`
	Enable       int    `json:"enable"`
}

//TrafficTotal 账号下流量消耗汇总
type TrafficTotal struct {
	TotalFlow      int64 `json:"totalFlow"`
	UsedFlow       int64 `json:"usedFlow"`
	AverageConsume int   `json:"averageConsume"`
}

//TrafficUserDetail 账户下流量消耗详情
type TrafficUserDetail struct {
	FlowDate     int64 `json:"flowDate"`
	DeviceCount  int   `json:"deviceCount"`
	ChannelCount int   `json:"channelCount"`
	HlsFlow      int64 `json:"hlsFlow"`
	AppFlow      int64 `json:"appFlow"`
	RtmpFlow     int64 `json:"rtmpFlow"`
	FlowCount    int64 `json:"flowCount"`
}
