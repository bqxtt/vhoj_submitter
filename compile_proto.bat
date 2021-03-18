cd pkg/sdk
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative submitterpb/service.proto
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative submitterpb/submission.proto
cd ../..

xcopy pkg\sdk\submitterpb\*.go ..\vhoj_rpc\model\submitterpb\ /s /y