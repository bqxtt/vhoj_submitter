package oj

import (
	"encoding/json"
	"fmt"
	"github.com/bqxtt/vhoj_submitter/pkg/common"
	"testing"
)

func TestHDUQuerier_Query(t *testing.T) {
	info := &common.SubmissionInfo{RealRunId: "35246418"}
	status, err := HduQuerier.Query(info)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, _ := json.Marshal(status)
	fmt.Println(string(bytes))
}
