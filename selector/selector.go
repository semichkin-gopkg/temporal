package selector

import wf "go.temporal.io/sdk/workflow"

type ImplSelector[T any] struct {
	sdkSelector wf.Selector
	response    *response[T]
}

type response[T any] struct {
	Ok   bool
	Data T
}

func NewSelector[T any](
	ctx wf.Context,
	channels ...wf.ReceiveChannel,
) *ImplSelector[T] {
	r := &response[T]{}

	sdkSelector := wf.NewSelector(ctx)

	for _, channel := range channels {
		sdkSelector.AddReceive(channel, func(channel wf.ReceiveChannel, more bool) {
			var data T

			if more {
				channel.Receive(ctx, &data)
			}

			r.Ok = more
			r.Data = data
		})
	}

	sdkSelector.AddReceive(ctx.Done(), func(channel wf.ReceiveChannel, more bool) {
		channel.Receive(ctx, nil)
		r.Ok = false
	})

	return &ImplSelector[T]{
		sdkSelector: sdkSelector,
		response:    r,
	}
}

func (s *ImplSelector[T]) Select(ctx wf.Context) (T, bool) {
	s.sdkSelector.Select(ctx)
	return s.response.Data, s.response.Ok
}

func (s *ImplSelector[T]) HasPending() bool {
	return s.sdkSelector.HasPending()
}
