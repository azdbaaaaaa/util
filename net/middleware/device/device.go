package device

import (
	"github.com/azdbaaaaaa/util/log"
	"strconv"
	"strings"
)

const (
	ContextKeyDevice = "device"
)

type (
	ClientType int
	AreaID     int
)

const (
	ClientType_ANDROID ClientType = 1
	ClientType_IOS     ClientType = 5
)

const (
	AreaID_ANDROID AreaID = 30
	AreaID_IOS     AreaID = 40
	AreaID_WEB     AreaID = 99
	AreaID_H5      AreaID = 98
)

type Device struct {
	IMEI         string     `json:"imei"`
	VersionName  string     `json:"version_name"`
	ScreenWidth  int        `json:"screen_width"`
	ScreenHeight int        `json:"screen_height"`
	Source       string     `json:"source"`
	SDK          string     `json:"sdk"`
	ClientType   ClientType `json:"client_type"`
	PhoneModel   string     `json:"phone_model"`
	VersionCode  int        `json:"version_code"`
	Channel      string     `json:"channel"`
	ClientTime   int64      `json:"client_time"`
	IsEmulator   int        `json:"is_emulator"`
	Longitude    string     `json:"longitude"`
	Latitude     string     `json:"latitude"`
	CountryCode  string     `json:"country_code"`
	City         string     `json:"city"`
	MCC          string     `json:"mcc"`
	TimeZone     string     `json:"time_zone"`
	ZoneId       string     `json:"zone_id"`
	GAID         string     `json:"gaid"`
	Timestamp    int64      `json:"timestamp"`

	UserAgent string `json:"user_agent"`
}

func New(text string, userAgent string) (d *Device) {
	d = &Device{
		UserAgent: userAgent,
	}
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

func (d *Device) ValueFromIdx(i int, v string) (err error) {
	switch i {
	case 0:
		d.IMEI = v
	case 1:
		d.VersionName = v
	case 2:
		sw, err := strconv.Atoi(v)
		if err != nil {
			return
		}
		d.ScreenWidth = sw
	case 3:
		sh, err := strconv.Atoi(v)
		if err != nil {
			return
		}
		d.ScreenHeight = sh
	case 4:
		d.Source = v
	case 5:
		d.SDK = v
	case 6:
		ct, err := strconv.Atoi(v)
		if err != nil {
			return
		}
		d.ClientType = ClientType(ct)
	case 7:
		d.PhoneModel = v
	case 8:
		vc, err := strconv.Atoi(v)
		if err != nil {
			return
		}
		d.VersionCode = vc
	case 9:
		d.Channel = v
	case 10:
		ct, err := strconv.ParseInt(v, 64, 10)
		if err != nil {
			return
		}
		d.ClientTime = ct
	case 11:
		ie, err := strconv.Atoi(v)
		if err != nil {
			return
		}
		d.IsEmulator = ie
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
		ts, err := strconv.ParseInt(v, 64, 10)
		if err != nil {
			return
		}
		d.Timestamp = ts
	}
	return
}

func (d *Device) GetAreaID() AreaID {
	switch d.ClientType {
	case ClientType_ANDROID:
		return AreaID_ANDROID
	case ClientType_IOS:
		return AreaID_IOS
	default:
		return AreaID_ANDROID
	}
}
