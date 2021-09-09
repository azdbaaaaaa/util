package metadata

import (
	"context"
	"github.com/azdbaaaaaa/util/log"
	"github.com/azdbaaaaaa/util/proto/common"
	"github.com/azdbaaaaaa/util/xutil/xerror"
	"strconv"
	"strings"
)

const (
	ContextKeyDevice = "device"
)

type Device struct {
	IMEI         string            `json:"imei"`          // deviceId
	VersionName  string            `json:"version_name"`  // 版本号 1.6.2.5
	ScreenWidth  int32             `json:"screen_width"`  // 屏幕宽度
	ScreenHeight int32             `json:"screen_height"` // 屏幕高度
	Source       string            `json:"source"`        // 第三方使用的source
	SDK          string            `json:"sdk"`           // 字符串
	ClientType   common.ClientType `json:"client_type"`   // 客户端类型  1：Android 5：iOS
	PhoneModel   string            `json:"phone_model"`   // 手机型号
	VersionCode  int32             `json:"version_code"`  // 版本号的数字版本，1.6.0以后都是4位，老版本是3位
	Channel      string            `json:"channel"`       // sdk为sdk_phoenix
	ClientTime   int64             `json:"client_time"`   // 客户端时间戳
	IsEmulator   int32             `json:"is_emulator"`   // 是否模拟器
	Longitude    string            `json:"longitude"`     // 经度
	Latitude     string            `json:"latitude"`      // 纬度
	CountryCode  string            `json:"country_code"`  // 国家
	City         string            `json:"city"`          // 城市
	MCC          string            `json:"mcc"`           // 运营商编码
	TimeZone     string            `json:"time_zone"`     // 时区 GMT+08:00
	ZoneId       string            `json:"zone_id"`       // 地区 Asia/Shanghai
	GAID         string            `json:"gaid"`          // gaid
	Timestamp    int64             `json:"timestamp"`     // 同client_time 客户端时间戳

	UserAgent string `json:"user_agent"`
}

func New(text string) (d *Device) {
	d = &Device{}
	parts := strings.Split(text, "|")
	for i, v := range parts {
		v = strings.TrimSpace(v)
		err := d.ValueFromIdx(i, v)
		if err != nil {
			log.Errorw("ValueFromIdx", "i", i, "v", v)
			continue
		}
	}
	return
}

func DeviceFromContext(ctx context.Context) (d Device, err error) {
	if v := ctx.Value(ContextKeyDevice); v != nil {
		if d, ok := v.(Device); ok {
			return d, nil
		}
	}
	return d, xerror.ErrNoDeviceError
}

func (d *Device) ValueFromIdx(i int, v string) (err error) {
	switch i {
	case 0:
		d.IMEI = v
	case 1:
		d.VersionName = v
	case 2:
		sw, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return err
		}
		d.ScreenWidth = int32(sw)
	case 3:
		sh, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return err
		}
		d.ScreenHeight = int32(sh)
	case 4:
		d.Source = v
	case 5:
		d.SDK = v
	case 6:
		ct, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		d.ClientType = common.ClientType(ct)
	case 7:
		d.PhoneModel = v
	case 8:
		vc, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return err
		}
		d.VersionCode = int32(vc)
	case 9:
		d.Channel = v
	case 10:
		ct, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		d.ClientTime = ct
	case 11:
		ie, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return err
		}
		d.IsEmulator = int32(ie)
	case 12:
		d.Longitude = v
	case 13:
		d.Latitude = v
	case 14:
		d.CountryCode = v
	case 15:
		d.City = v
	case 16:
		d.MCC = v
	case 17:
		d.TimeZone = v
	case 18:
		d.ZoneId = v
	case 19:
		if strings.ToLower(v) == "null" {
			d.GAID = ""
		} else {
			d.GAID = v
		}
	case 20:
		ts, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		d.Timestamp = ts
	}
	return
}
