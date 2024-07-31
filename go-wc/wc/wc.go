package wc

import (
	"bytes"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type baseWc struct {
	log *logrus.Logger
}

func NewWc(log *logrus.Logger) *baseWc {
	return &baseWc{log: log}
}

func (b *baseWc) ReadFile(filepath string) ([]byte, error) {
	reader, err := os.Open(filepath)
	if err != nil {
		b.log.Error(err)
		return nil, err
	}
	filebyte, err := io.ReadAll(reader)
	if err != nil {
		b.log.Error(err)
		return nil, err
	}
	return filebyte, nil
}

func (b *baseWc) C(filepath string) (int, error) {
	filebyte, err := b.ReadFile(filepath)
	if err != nil {
		b.log.Error(err)
		return 0, err
	}
	return len(filebyte), nil
}

func (b *baseWc) L(filepath string) (int, error) {
	filebyte, err := b.ReadFile(filepath)
	if err != nil {
		b.log.Error(err)
		return 0, err
	}
	return bytes.Count(filebyte, []byte{'\n'}), nil
}

func (b *baseWc) W(filepath string) (int, error) {
	filebyte, err := b.ReadFile(filepath)
	if err != nil {
		b.log.Error(err)
		return 0, err
	}
	return len(bytes.Fields(filebyte)), nil
}

func (b *baseWc) M(filepath string) (int, error) {
	filebyte, err := b.ReadFile(filepath)
	if err != nil {
		b.log.Error(err)
		return 0, err
	}
	return len(bytes.Runes(filebyte)), nil
}
