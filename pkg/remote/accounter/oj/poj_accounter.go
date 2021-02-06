package oj

import (
	ytask "github.com/ecnuvj/vhoj_submitter/pkg/bootstrap/ytask/client"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/accounter"
)

var PojAccounter accounter.IAccounter = &POJAccounter{}

type POJAccounter struct {
	accounter.DefaultAccounterImpl
}

func (P *POJAccounter) FindAccount() (*common.RemoteAccount, error) {
	//todo
	return &common.RemoteAccount{
		Username: "bqx",
		Password: "tcg19981108",
	}, nil
}

func (P *POJAccounter) ReleaseAccount(account *common.RemoteAccount) error {
	//todo
	return nil
}

func (P *POJAccounter) HandleSubmit(info *common.SubmissionInfo) error {
	remoteAccount, err := P.FindAccount()
	if err != nil {
		return err
	}
	if remoteAccount != nil {
		_, err = ytask.Client.Send("code", "submit_code", info, remoteAccount)
		if err != nil {
			return err
		}
	} else {
		// todo join waiting queue
	}
	return nil
}

func (P *POJAccounter) HandleQuery(info *common.SubmissionInfo, account *common.RemoteAccount) error {
	err := P.ReleaseAccount(account)
	if err != nil {
		return err
	}
	_, err = ytask.Client.Send("code", "query_result", info, account)
	if err != nil {
		return err
	}
	return nil
}
