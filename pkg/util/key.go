package util

import (
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
)

func GetCookieKey(remoteOJ remote_oj.RemoteOJ, account *common.RemoteAccount) string {
	return fmt.Sprintf("oj_%v_username_%v", remoteOJ, account.Username)
}
