package device

import (
	"github.com/azdbaaaaaa/util/log"
	"github.com/azdbaaaaaa/util/proto"
	"strconv"
	"strings"
)

const (
	ContextKeyDevice = "device"
)

type Device struct {
	IMEI         string           `json:"imei"`
	VersionName  string           `json:"version_name"`
	ScreenWidth  int              `json:"screen_width"`
	ScreenHeight int              `json:"screen_height"`
	Source       string           `json:"source"`
	SDK          string           `json:"sdk"`
	ClientType   proto.ClientType `json:"client_type"`
	PhoneModel   string           `json:"phone_model"`
	VersionCode  int              `json:"version_code"`
	Channel      string           `json:"channel"`
	ClientTime   int64            `json:"client_time"`
	IsEmulator   int              `json:"is_emulator"`
	Longitude    string           `json:"longitude"`
	Latitude     string           `json:"latitude"`
	CountryCode  string           `json:"country_code"`
	City         string           `json:"city"`
	MCC          string           `json:"mcc"`
	TimeZone     string           `json:"time_zone"`
	ZoneId       string           `json:"zone_id"`
	GAID         string           `json:"gaid"`
	Timestamp    int64            `json:"timestamp"`

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
			return err
		}
		d.ScreenWidth = sw
	case 3:
		sh, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		d.ScreenHeight = sh
	case 4:
		d.Source = v
	case 5:
		d.SDK = v
	case 6:
		ct, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		d.ClientType = proto.ClientType(ct)
	case 7:
		d.PhoneModel = v
	case 8:
		vc, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		d.VersionCode = vc
	case 9:
		d.Channel = v
	case 10:
		ct, err := strconv.ParseInt(v, 64, 10)
		if err != nil {
			return err
		}
		d.ClientTime = ct
	case 11:
		ie, err := strconv.Atoi(v)
		if err != nil {
			return err
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
			return err
		}
		d.Timestamp = ts
	}
	return
}

func (d *Device) GetAreaID() proto.AreaIdType {
	switch d.ClientType {
	case proto.ClientType_CLIENT_TYPE_IOS:
		return proto.AreaIdType_AREA_ID_IOS
	case proto.ClientType_CLIENT_TYPE_ANDROID:
		return proto.AreaIdType_AREA_ID_ANDROID
	default:
		return proto.AreaIdType_AREA_ID_ANDROID
	}
}
