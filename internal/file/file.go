package file

import (
	"context"
	"fmt"
	"go_logs/pkg/logger"
	"go_logs/pkg/utils"
	"io"
	"os"
)

type Mfile struct {
	ended  bool
	fname  string
	logger *logger.Logger
}

const (
	maxSize int = 1000
)

func New(ctx context.Context, logger *logger.Logger) *Mfile {
	m := &Mfile{
		ended:  true,
		fname:  "logs/log1.log",
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

func (m *Mfile) Write1(buf []byte) {
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

//====================

func (m *Mfile) Write(buf []byte) {
	var f *os.File
	var err error

	if m.ended {
		m.logger.Info().Msgf("clear file")
		f, err = os.OpenFile(m.fname, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			m.logger.Info().Msgf("err open file, %v", err)
			return
		}
		m.ended = false
	} else {
		f, err = os.OpenFile(m.fname, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
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
		m.logger.Info().Msgf("copy file, sz = %d", sz)
		fname1 := m.get_new_filename()
		f1, err := os.OpenFile(fname1, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			m.logger.Info().Msgf("err create file, %v", err)
			return
		}
		// f.Close()
		// err := os.Rename(m.fname, fname1)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		defer f1.Close()
		m.ended = true
		_, err = io.Copy(f1, f)
		if err != nil {
			m.logger.Info().Msgf("err copy file, %v", err)
		}

		// _, err = f.Seek(0, 0)

		// err = f.Truncate(0)

	}

}
