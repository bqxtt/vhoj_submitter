package loginer

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"net/http"
)

type ILoginer interface {
	GetOJInfo() *remote_oj.RemoteOJInfo
	Login(account *common.RemoteAccount) ([]*http.Cookie, error)
	mustEmbedDefaultLoginer()
}

type DefaultLoginerImpl struct{}

func (DefaultLoginerImpl) GetOJInfo() *remote_oj.RemoteOJInfo {
	panic("implement me")
}

func (DefaultLoginerImpl) Login(*common.RemoteAccount) ([]*http.Cookie, error) {
	panic("implement me")
}

func (DefaultLoginerImpl) mustEmbedDefaultLoginer() {
}
