package cli

import (
	"context"
	"fmt"
)

type EventCMD struct {
	Send   *SendEventCMD   `arg:"subcommand:send"`
	Listen *ListenEventCMD `arg:"subcommand:listen"`
}

func (e *EventCMD) handle(ctx context.Context) error {

	switch {
	case e.Send != nil:
		e.Send.handle(ctx)
	case e.Listen != nil:
		e.Listen.handle(ctx)
	}
	return nil
}

type SendEventCMD struct {
	Channel string `arg:"positional,required"`
	Message string `arg:"positional,required"`
}

func (e *SendEventCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

	fmt.Println(e.Channel, e.Message)
	client.SendEvent(e.Channel, e.Message)
	return nil
}

type ListenEventCMD struct {
}

func (e *ListenEventCMD) handle(ctx context.Context) error {
	return nil
}
