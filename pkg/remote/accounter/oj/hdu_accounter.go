package oj

import (
	ytask "github.com/bqxtt/vhoj_submitter/pkg/bootstrap/ytask/client"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"github.com/bqxtt/vhoj_submitter/pkg/remote/accounter"
)

var HduAccounter accounter.IAccounter = &HDUAccounter{}

type HDUAccounter struct {
	accounter.DefaultAccounterImpl
	remoteAccounts []*common.RemoteAccount
}

func (H *HDUAccounter) FindAccount() (*common.RemoteAccount, error) {
	//todo
	return &common.RemoteAccount{
		Username: "bqxtt233",
		Password: "tcg19981108",
	}, nil
}

func (H *HDUAccounter) ReleaseAccount(account *common.RemoteAccount) error {
	//todo
	return nil
}

func (H *HDUAccounter) HandleSubmit(submission *common.SubmissionInfo) error {
	remoteAccount, err := H.FindAccount()
	if err != nil {
		return err
	}
	if remoteAccount != nil {
		_, err = ytask.Client.Send("code", "submit_code", submission, remoteAccount)
		if err != nil {
			return err
		}
	} else {
		// todo join waiting queue
	}
	return nil
}

func (H *HDUAccounter) HandleQuery(info *common.SubmissionInfo, account *common.RemoteAccount) error {
	err := H.ReleaseAccount(account)
	if err != nil {
		return err
	}
	_, err = ytask.Client.Send("code", "query_result", info)
	if err != nil {
		return err
	}
	return nil
}
