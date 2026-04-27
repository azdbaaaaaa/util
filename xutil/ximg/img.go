package ximg

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/azdbaaaaaa/util/log"
	"github.com/kirinlabs/HttpRequest"
	"net/http"
	"strings"
)

func GetImageTypeFromURL(ctx context.Context, url string) (tp string, err error) {
	resp, err := HttpRequest.NewRequest().Get(url)
	if err != nil {
		log.WithContext(ctx).Errorw("error to get url", "err", err, "url", url)
		return
	}
	defer resp.Close()
	if resp.StatusCode() > http.StatusIMUsed {
		log.WithContext(ctx).Errorw("error to get url", "err", err, "url", url, "status_code", resp.StatusCode())
		err = errors.New("error to get url")
		return
	}

	b, err := resp.Body()
	if err != nil {
		log.WithContext(ctx).Errorw("error to read body", "err", err)
		return
	}

	contentType := http.DetectContentType(b[:512])

	for i := 0; i < len(ExtList); i++ {
		if strings.Contains(ExtList[i], contentType[6:len(contentType)]) {
			return ExtList[i], nil
		}
	}
	return "", errors.New("Invalid image type")

}

func GetContentType(base64String string) (string, error) {
	if strings.HasPrefix(base64String, "data:image/") {
		base64String = base64String[strings.Index(base64String, ",")+1:]
	}
	res, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return "", err
	}
	return http.DetectContentType(res[:512]), nil
}

func GetImageType(base64String string) (string, error) {
	contentType, err := GetContentType(base64String)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(ExtList); i++ {
		if strings.Contains(ExtList[i], contentType[6:len(contentType)]) {
			return ExtList[i], nil
		}
	}
	return "", errors.New("invalid image type")
}
