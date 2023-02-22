module github.com/hashgraph/hedera-protobufs-go/examples

go 1.15

require (
	github.com/hashgraph/hedera-protobufs-go v0.2.0
	google.golang.org/genproto v0.0.0-20230221151758-ace64dc21148 // indirect
	google.golang.org/grpc v1.53.0
)

replace github.com/hashgraph/hedera-protobufs-go => ../
