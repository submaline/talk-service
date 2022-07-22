package talk_service

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	talkv1 "github.com/submaline/talk-service/gen/talk/v1"
)

type TalkService struct{}

func (ts *TalkService) SendMessage(
	_ context.Context, request *connect.Request[talkv1.SendMessageRequest]) (
	*connect.Response[talkv1.SendMessageResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}

func (ts *TalkService) SendReadReceipt(
	_ context.Context, request *connect.Request[talkv1.SendReadReceiptRequest]) (
	*connect.Response[talkv1.SendReadReceiptResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}
