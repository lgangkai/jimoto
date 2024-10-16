# execute to generate .pb.go and .pb.micro.go files
cd ..
cd proto
protoc --micro_out=./ --go_out=./ commodity/commodity.proto
protoc --micro_out=./ --go_out=./ account/account.proto
protoc --micro_out=./ --go_out=./ order/order.proto
cd .. && cd ..
