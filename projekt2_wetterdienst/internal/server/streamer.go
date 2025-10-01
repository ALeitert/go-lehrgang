package server

import (
	"context"
)

var (
	postChan = make(chan postMsg, 256)
	listChan = make(chan listenerMsg, 256)

	listeners = map[string][]listener{}
)

type postMsg struct {
	TempMessage
	city string
}

type listenerMsg struct {
	listener
	city string
}

type listener struct {
	ctx     context.Context
	msgChan chan TempMessage
}

type Streamer struct{}

func (str *Streamer) Name() string                   { return "Streamer" }
func (str *Streamer) Init(ctx context.Context) error { return nil }
func (str *Streamer) Stop() error                    { return nil }

func (str *Streamer) Run(ctx context.Context) error {
	for done := false; !done; {
		select {
		case <-ctx.Done():
			done = true
			continue

		case msg := <-postChan:
			listList := listeners[msg.city]
			for idx := 0; idx < len(listList); idx++ {
				listener := &listList[idx]

				select {
				case <-listener.ctx.Done():
					close(listener.msgChan)

					// Remove by swapping with last.
					lastIdx := len(listList) - 1
					listList[idx] = listList[lastIdx]
					listList = listList[:lastIdx]

					idx--
					continue

				default:
				}

				// Second select to give priority to ctx.Done().
				select {
				case listener.msgChan <- msg.TempMessage:
				default:
				}
			}
			listeners[msg.city] = listList

		case reg := <-listChan:
			listeners[reg.city] = append(listeners[reg.city], reg.listener)
		}
	}

	for _, listenerList := range listeners {
		for _, listener := range listenerList {
			close(listener.msgChan)
		}
	}

	return nil
}

func Post(city string, msg TempMessage) {
	postChan <- postMsg{msg, city}
}

func Listen(ctx context.Context, city string) <-chan TempMessage {
	msgChan := make(chan TempMessage, 256)
	listChan <- listenerMsg{listener{ctx, msgChan}, city}

	return msgChan
}
