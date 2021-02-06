package oj

import (
	"encoding/json"
	"fmt"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"testing"
)

func TestHDUQuerier_Query(t *testing.T) {
	info := &common.SubmissionInfo{RealRunId: "35301974"}
	status, err := HduQuerier.Query(info, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, _ := json.Marshal(status)
	fmt.Println(string(bytes))
}

func TestPOJQuerier_Query(t *testing.T) {
	info := &common.SubmissionInfo{RealRunId: "22378035"}
	account := &common.RemoteAccount{
		Username: "bqx",
		Password: "tcg19981108",
	}
	status, err := PojQuerier.Query(info, account)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, _ := json.Marshal(status)
	fmt.Println(string(bytes))
}
