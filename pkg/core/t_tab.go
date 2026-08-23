package core

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/term"
)

type TTabSize struct {
	Height int
	Width  int
}

type TStdoutDescriptor int

func (d TStdoutDescriptor) ToInt() int {
	return int(d)
}

type drawer interface {
	Draw(context.Context, TTabSize) error
}

type defaultDrawer struct{}

func (dd *defaultDrawer) Draw(context.Context, TTabSize) error { return nil }

type TTab struct {
	size             TTabSize
	stdoutDescriptor TStdoutDescriptor

	drawTimeout time.Duration
	drawer      drawer

	mtx sync.RWMutex
}

func NewTTab() *TTab {
	return &TTab{
		drawTimeout: time.Second,
		drawer:      &defaultDrawer{},
	}
}

func (t *TTab) setUpStdoutDescriptor() error {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	fd := int(os.Stdout.Fd())

	// Проверяем, запущен ли код в реальном терминале
	if !term.IsTerminal(fd) {
		return ErrTInvalidTerminalDescriptor
	}

	t.stdoutDescriptor = TStdoutDescriptor(fd)

	return nil
}

func (t *TTab) setUpSize() error {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	width, height, err := term.GetSize(t.stdoutDescriptor.ToInt())
	if err != nil {
		return WrapErrors(ErrTGetSize, err)
	}

	t.size.Height = height
	t.size.Width = width

	return nil
}

func (t *TTab) GetSize() TTabSize {
	t.mtx.RLock()
	defer t.mtx.RUnlock()

	size := t.size

	return size
}

func (t *TTab) listenResize(ctx context.Context) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGWINCH)

	for {
		select {
		case <-ctx.Done():
			return
		case <-sigChan:
			_ = t.setUpSize()
			ctx, _ := context.WithTimeout(ctx, time.Second)
			_ = t.drawer.Draw(ctx, t.GetSize())
		}
	}
}
