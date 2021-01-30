package util

import (
	"fmt"
	"github.com/bqxtt/vhoj_common/pkg/common/constants"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
)

func GetCookieKey(remoteOJ constants.RemoteOJ, account *common.RemoteAccount) string {
	return fmt.Sprintf("oj_%v_username_%v", remoteOJ, account.Username)
}
