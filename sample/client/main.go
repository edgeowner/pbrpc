package main

import (
	"context"
	"fmt"

	"github.com/let-z-go/pbrpc"
	"github.com/let-z-go/pbrpc/sample"
)

type ClientServiceHandler struct {
	sample.ClientServiceHandlerBase
}

func (ClientServiceHandler) GetNickname(context_ context.Context, channel pbrpc.Channel) (*sample.GetNicknameResponse, error) {
	response := &sample.GetNicknameResponse{
		Nickname: "007",
	}

	return response, nil
}

func main() {
	channelPolicy := (&pbrpc.ChannelPolicy{}).RegisterServiceHandler(ClientServiceHandler{})
	channel := (&pbrpc.ClientChannel{}).Initialize(channelPolicy, []string{"127.0.0.1:8888"}, nil)

	go func() {
		client := sample.ServerServiceClient{channel, nil}

		request := &sample.SayHelloRequest{
			ReplyFormat: "Hello, %v!",
		}

		response, _ := client.SayHello(request, true)
		fmt.Println(response.Reply)
		channel.Stop()
	}()

	channel.Run()
}
