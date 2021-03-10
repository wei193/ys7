package ys7

//接口地址
const (
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

	UUIDPICTURE      = "https://open.ys7.com/api/lapp/device/uuid/picture"  //设备互联互通根据UUID查询抓拍的图片
	DEVICESTATUSINFO = "https://open.ys7.com/api/lapp/device/status/get"    //根据序列号获取设备的状态信息
	DEVICECAMERALIST = "https://open.ys7.com/api/lapp/device/camera/list"   //根据序列号获取设备的通道信息
	DEVICESUPPORT    = "https://open.ys7.com/api/lapp/device/support/ezviz" //根据设备型号以及设备版本号查询设备是否支持萤石协议
	DEVICECAPACITY   = "https://open.ys7.com/api/lapp/device/capacity"      //根据设备序列号查询设备能力集

	//根据时间获取录像信息

	//[设备]配置
	DEFENCESET        = "https://open.ys7.com/api/lapp/device/defence/set"         //设置设备活动检测开关状态
	OFFENCRYPT        = "https://open.ys7.com/api/lapp/device/encrypt/off"         //关闭设备视频加密开关
	ONENCRYPT         = "https://open.ys7.com/api/lapp/device/encrypt/on"          //打开设备视频加密开关
	SWITICHSTATUS     = "https://open.ys7.com/api/lapp/device/sound/switch/status" //获取wifi配置或者设备重启提示音开关状态
	SETSOUND          = "https://open.ys7.com/api/lapp/device/sound/switch/set"    //设置wifi配置或者设备重启提示音开关状态
	SCRNESWITCHSTATUS = "https://open.ys7.com/api/lapp/device/scene/switch/status" //获取镜头遮蔽开关状态
	SETSCRENSWITCH    = "https://open.ys7.com/api/lapp/device/scene/switch/set"    //设置镜头遮蔽开关
	//获取声源定位开关状态
	//设置声源定位开关
	//获取设备布撤防（活动检测）时间计划
	//设置布撤防（活动检测）时间计划
	//获取摄像机指示灯开关状态
	//设置摄像机指示灯开关
	//获取全天录像开关状态
	//设置全天录像开关
	//获取智能算法配置信息
	//设置智能算法模式
	//设置设备告警声音模式
	//开启或关闭设备下线通知
	//获取设备麦克风即声音开关状态
	//设置设备麦克风即声音开关
	//设置设备移动跟踪开关
	//获取设备移动跟踪开关状态
	//设置设备预览时osd名称
	//获取设备智能检测开关状态
	//设置智能检测开关

	//[设备]升级
	// --全部待实现
)

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

//GetAllDeviceList 获取所有设备列表
func (ys *Ys7) GetAllDeviceList() (devices []Device, err error) {
	var page Page
	devices, page, err = ys.GetDeviceList(0, 50)
	if err != nil {
		return
	}
	for pageNum := 1; pageNum <= page.Total/50; pageNum++ {
		list, _, err := ys.GetDeviceList(pageNum, 50)
		if err != nil {
			return nil, err
		}
		devices = append(devices, list...)
	}
	return
}

//GetDeviceList 获取设备列表
func (ys *Ys7) GetDeviceList(pageStart, pageSize int) (devices []Device, page Page, err error) {
	params := make(map[string]interface{})
	params["pageStart"] = pageStart
	params["pageSize"] = pageSize
	_, err = ys.authorizeRequset("POST", DEVICELIST, params, &devices, &page) //获取用户下的设备列表
	if err != nil {
		return nil, page, err
	}
	return devices, page, nil
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

//GetAllCameraList 获取所有摄像头列表
func (ys *Ys7) GetAllCameraList() (cameras []Camera, err error) {
	var page Page
	cameras, page, err = ys.GetCameraList(0, 50)
	if err != nil {
		return
	}
	for pageNum := 1; pageNum <= page.Total/50; pageNum++ {
		list, _, err := ys.GetCameraList(pageNum, 50)
		if err != nil {
			return nil, err
		}
		cameras = append(cameras, list...)
	}
	return
}

//GetCameraList 获取摄像头列表
func (ys *Ys7) GetCameraList(pageStart, pageSize int) (cameras []Camera, page Page, err error) {
	params := make(map[string]interface{})
	params["pageStart"] = pageStart
	params["pageSize"] = pageSize
	_, err = ys.authorizeRequset("POST", CAMERALIST, params, &cameras, &page)
	if err != nil {
		return nil, page, err
	}
	return cameras, page, nil
}

//GetDeviceCameraList 获取指定设备的通道信息
func (ys *Ys7) GetDeviceCameraList(deviceSerial string) (cameras []Camera, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	_, err = ys.authorizeRequset("POST", DEVICECAMERALIST, params, &cameras)
	if err != nil {
		return nil, err
	}
	return cameras, nil
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

//GetPictureByUUID 设备互联互通根据UUID查询抓拍的图片
func (ys *Ys7) GetPictureByUUID(uuid string, size int) (pic *Picture, err error) {
	params := make(map[string]interface{})
	params["uuid"] = uuid
	params["size"] = size
	pic = &Picture{}
	_, err = ys.authorizeRequset("POST", UUIDPICTURE, params, &pic)
	if err != nil {
		return
	}
	return

}

//GetDeviceStatusInfo 根据序列号获取设备的状态信息
func (ys *Ys7) GetDeviceStatusInfo(deviceSerial string, channel int) (devInfo *DeviceStatusInfo, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channel"] = channel
	devInfo = &DeviceStatusInfo{}
	_, err = ys.authorizeRequset("POST", DEVICESTATUSINFO, params, &devInfo)
	if err != nil {
		return
	}
	return
}

//IsSupportEzviz 根据设备型号以及设备版本号查询设备是否支持萤石协议
func (ys *Ys7) IsSupportEzviz(model, version string) (protocol *Protocol, err error) {
	params := make(map[string]interface{})
	params["appKey"] = ys.AppKey
	params["model"] = model
	params["version"] = version
	protocol = &Protocol{}
	_, err = ys.authorizeRequset("POST", DEVICESUPPORT, params, &protocol)
	if err != nil {
		return
	}
	return
}

//GetDeviceCap 根据设备序列号查询设备能力集
func (ys *Ys7) GetDeviceCap(deviceSerial string) (deviceCap *DeviceCapacity, err error) {
	params := make(map[string]interface{})
	params["accessToken"] = ys.AccessToken
	params["deviceSerial"] = deviceSerial
	deviceCap = &DeviceCapacity{}
	_, err = ys.authorizeRequset("POST", DEVICECAPACITY, params, &deviceCap)
	if err != nil {
		return
	}
	return

}

// SetDefence 设备布撤防
func (ys *Ys7) SetDefence(deviceSerial string, isDefence int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["isDefence"] = isDefence
	_, err = ys.authorizeRequset("POST", DEFENCESET, params, nil)
	if err != nil {
		return
	}
	return nil
}

//OffEncrypt 关闭设备视频加密开关
func (ys *Ys7) OffEncrypt(deviceSerial, validateCode string) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["validateCode"] = validateCode
	_, err = ys.authorizeRequset("POST", OFFENCRYPT, params, nil)
	if err != nil {
		return
	}
	return nil
}

//OnEncrypt 开启设备视频加密开关
func (ys *Ys7) OnEncrypt(deviceSerial string) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	_, err = ys.authorizeRequset("POST", ONENCRYPT, params, nil)
	if err != nil {
		return
	}
	return nil
}

//GetSoundSwitchStatus 获取wifi配置提示音开关状态
func (ys *Ys7) GetSoundSwitchStatus(deviceSerial string) (soundStatus *SwitchStatus, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	soundStatus = &SwitchStatus{}
	_, err = ys.authorizeRequset("POST", SWITICHSTATUS, params, &soundStatus)
	if err != nil {
		return
	}
	return
}

//SetSoundSwitch 设置wifi配置或设备启动提示音开关
func (ys *Ys7) SetSoundSwitch(deviceSerial string, enable, channelNo int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["enable"] = enable
	params["channelNo"] = channelNo
	_, err = ys.authorizeRequset("POST", SETSOUND, params, nil)
	if err != nil {
		return
	}
	return
}

//GetSceneSwitchStatus 获取镜头遮蔽开关状态
func (ys *Ys7) GetSceneSwitchStatus(deviceSerial string) (switchStatus *SwitchStatus, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	switchStatus = &SwitchStatus{}
	_, err = ys.authorizeRequset("POST", SCRNESWITCHSTATUS, params, &switchStatus)
	if err != nil {
		return
	}
	return
}

//SetSceneSwitch 设置设备镜头遮蔽开关状态
func (ys *Ys7) SetSceneSwitch(deviceSerial string, enable, channelNo int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["enable"] = enable
	params["channelNo"] = channelNo
	_, err = ys.authorizeRequset(deviceSerial, SETSCRENSWITCH, params, nil)
	if err != nil {
		return
	}
	return
}
