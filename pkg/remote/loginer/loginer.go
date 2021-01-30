package loginer

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"net/http"
)

type ILoginer interface {
	GetOJInfo() *constants.RemoteOJInfo
	Login(account *common.RemoteAccount) ([]*http.Cookie, error)
	mustEmbedDefaultLoginer()
}

type DefaultLoginerImpl struct{}

func (DefaultLoginerImpl) GetOJInfo() *constants.RemoteOJInfo {
	panic("implement me")
}

func (DefaultLoginerImpl) Login(*common.RemoteAccount) ([]*http.Cookie, error) {
	panic("implement me")
}

func (DefaultLoginerImpl) mustEmbedDefaultLoginer() {
}
