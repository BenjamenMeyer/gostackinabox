package common_test

import (
    "net/http"
    "testing"

    "github.com/TestInABox/gostackinabox/common"
)

func Test_Common_HttpReply(t *testing.T) {
    hr := &common.HttpReply{
        Status: common.HttpStatus_RouteNotHandled,
        Headers: http.Header{},
    }
    if hr == nil {
        t.Error("Unexpected nil")
    }
}
