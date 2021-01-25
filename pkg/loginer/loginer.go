package loginer

import "github.com/bqxtt/vhoj_common/pkg/common/constants"

type ILoginer interface {
	GetOJInfo() *constants.RemoteOJInfo
	Login() error
	mustEmbedDefaultLoginer()
}

type DefaultLoginerImpl struct{}

func (DefaultLoginerImpl) GetOJInfo() *constants.RemoteOJInfo {
	panic("implement me")
}

func (DefaultLoginerImpl) Login() error {
	panic("implement me")
}

func (DefaultLoginerImpl) mustEmbedDefaultLoginer() {
}
