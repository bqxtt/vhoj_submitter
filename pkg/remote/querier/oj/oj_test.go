package oj

import (
	"encoding/json"
	"fmt"
	"github.com/ecnuvj/vhoj_submitter/pkg/common"
	"testing"
)

func TestHDUQuerier_Query(t *testing.T) {
	info := &common.SubmissionInfo{RealRunId: "35259548"}
	status, err := HduQuerier.Query(info)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, _ := json.Marshal(status)
	fmt.Println(string(bytes))
}
