package ys7

import (
	"testing"
)

func TestDeviceAdd(t *testing.T) {
	err := ys.AddDevice(DeviceSerial, ValidateCode)
	if err != nil {
		t.Error(err)
		return
	}
	tlog("AddDevice success")
}

func TestGetDeviceList(t *testing.T) {
	devices, page, err := ys.GetDeviceList(0, 10)
	if err != nil {
		t.Error(err)
		return
	}
	tlog("GetDeviceList", devices, page)
}

func TestUpdateDeviceName(t *testing.T) {
	err := ys.UpdateDeviceName(DeviceSerial, "ys7-sdk")
	if err != nil {
		t.Error(err)
		return
	}
	tlog("UpdateDeviceName success")
}

func TestInfoDevice(t *testing.T) {
	info, err := ys.InfoDevice(DeviceSerial)
	if err != nil {
		t.Error(err)
		return
	}
	tlog("InfoDevice ", info)
}

func TestGetDeviceCapture(t *testing.T) {
	pic, err := ys.GetDeviceCapture(DeviceSerial, ChannelNo)
	if err != nil {
		t.Error(err)
		return
	}
	tlog("GetDeviceCapture", pic)
}

func TestGetCameraList(t *testing.T) {
	list, page, err := ys.GetCameraList(0, 10)
	if err != nil {
		t.Error(err)
		return
	}
	tlog("GetCameraList", list, page)
}

func UpdateCameraName(t *testing.T) {
	err := ys.UpdateCameraName(DeviceSerial, "ys7-camera", ChannelNo)
	if err != nil {
		t.Error(err)
		return
	}
	tlog("UpdateCameraName success")
}

func TestDeleteDevice(t *testing.T) {
	err := ys.DeleteDevice(DeviceSerial)
	if err != nil {
		t.Error(err)
		return
	}
	tlog("DeleteDevice success")
}
