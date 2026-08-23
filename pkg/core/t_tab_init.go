package core

import "context"

func (t *TTab) Init(ctx context.Context) error {
	if err := t.setUpStdoutDescriptor(); err != nil {
		return WrapErrors(ErrTInitError, err)
	}
	if err := t.setUpSize(); err != nil {
		return WrapErrors(ErrTInitError, err)
	}
	go t.listenResize(ctx)

	return nil
}
