# Proto Publisher

Proto publisher is a tool to publish messages to RabbitMQ by using [Protobuf (Google Protocol Buffers)](https://developers.google.com/protocol-buffers).

### Install tool
    
    go install publisher.go

### Install go protobuf 

	go get github.com/golang/protobuf
	go get github.com/golang/protobuf/protoc-gen-go

### Complie protobuf
    
    protoc resources/proto/*.proto --go_out=pkg/
    
### Publish message

    publisher
