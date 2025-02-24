package file

import (
	"context"
	"fmt"
	"go_logs/pkg/bench"
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
	maxSize int = 1000000 // 1048576
)

func New(ctx context.Context, logger *logger.Logger) *Mfile {
	m := &Mfile{
		ended:  true,
		fname:  "logs/rx.log",
		logger: logger,
	}
	return m
}

func (m *Mfile) get_new_filename() string {
	t := utils.GetTime()
	fname := fmt.Sprintf("logs/%s.log", t)
	// m.logger.Info().Msgf(" new file: %s", fname)
	return fname
}

//===================

func (m *Mfile) Copy1() {
	defer bench.Duration(bench.Track("copied "))
	fsrc, err := os.OpenFile(m.fname, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		m.logger.Info().Msgf("err open file, %v", err)
		return
	}
	// m.logger.Info().Msgf("copy file")
	fname1 := m.get_new_filename()
	fdst, err := os.OpenFile(fname1, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		m.logger.Info().Msgf("err create file, %v", err)
		return
	}

	defer func() {
		fsrc.Close()
		fdst.Close()
	}()

	sz1 := maxSize + 100000
	b := make([]byte, sz1)
	sz := 0
	for {
		len, err := fsrc.Read(b)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			// m.logger.Info().Msgf("read, len=%d", len)
			break
		}
		fdst.Write(b[:len])
		sz = sz + len
	}
	// m.logger.Info().Msgf("copied, sz=%d", sz)
}

//====================

func (m *Mfile) Write(buf []byte) {
	var f *os.File
	var err error

	if m.ended {
		// m.logger.Info().Msgf("clear file")
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

	fi, err := f.Stat() // проверяем размер файла
	if err != nil {
		m.logger.Info().Msgf("err get file size, %v", err)
		return
	}
	sizeFile := fi.Size()

	if sizeFile < int64(maxSize) {
		return
	}
	f.Close()

	m.ended = true
	m.Copy1() // переносим в другой файл при переполнении

}
