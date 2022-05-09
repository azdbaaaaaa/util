package ximg

import (
	"context"
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

	//buff := make([]byte, 512)
	filetype := http.DetectContentType(b[:512])

	for i := 0; i < len(ExtList); i++ {
		if strings.Contains(ExtList[i], filetype[6:len(filetype)]) {
			return ExtList[i], nil
		}
	}
	return "", errors.New("Invalid image type")

}
