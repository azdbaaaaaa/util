package ximg

import (
	"context"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGetImageTypeFromURL(t *testing.T) {
	tp, err := GetImageTypeFromURL(context.Background(), "https://pro-docusign-agreement.s3.eu-west-1.amazonaws.com/receiver/id_card/0307be884c366ed03bb9f8d0a112fa03.png")
	assert.Equal(t, err, nil)
	t.Log(tp)
}
