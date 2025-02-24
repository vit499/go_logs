package file

import (
	"context"
	"fmt"
	"go_logs/pkg/logger"
	"go_logs/pkg/utils"
	"os"
)

type Mfile struct {
	opened bool
	fname  string
	logger *logger.Logger
}

const (
	maxSize int = 1000000
)

func New(ctx context.Context, logger *logger.Logger) *Mfile {
	m := &Mfile{
		opened: false,
		fname:  "",
		logger: logger,
	}
	return m
}

func (m *Mfile) get_new_filename() string {
	t := utils.GetTime()
	fname := fmt.Sprintf("logs/%s.log", t)
	m.logger.Info().Msgf(" new file: %s", fname)
	return fname
}

func (m *Mfile) Write(buf []byte) {
	var f *os.File
	var err error
	if m.fname == "" {
		m.fname = m.get_new_filename()
		f, err = os.OpenFile(m.fname, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			m.logger.Info().Msgf("err open file, %v", err)
			return
		}
	} else {
		f, err = os.OpenFile(m.fname, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			m.logger.Info().Msgf("err open file, %v", err)
			return
		}
	}
	defer f.Close()
	f.Write(buf)

	fi, err := f.Stat()
	if err != nil {
		m.logger.Info().Msgf("err get file size, %v", err)
		return
	}
	sz := fi.Size()
	if sz > int64(maxSize) {
		m.logger.Info().Msgf("close file, sz = %d", sz)
		m.fname = "" // превышен размер, в следующий раз будет создан другой файл
	}

}
