package main

import (
	"bytes"
	"errors"
	"io"
	"math"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	fileTo, err := os.OpenFile(toPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		return err
	}

	defer func() {
		fromFile.Close()
		fileTo.Close()
	}()

	fromInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	if !fromInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > fromInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	if offset == 0 && limit == 0 {
		_, err = io.Copy(fileTo, fromFile)
		if err != nil {
			return err
		}
	}
	var tempBuff []byte
	if offset == 0 && limit != 0 {
		if limit >= fromInfo.Size() {
			tempBuff = make([]byte, fromInfo.Size())
		} else {
			tempBuff = make([]byte, limit)
		}
	}
	if offset > 0 && limit > 0 {
		_, err := fromFile.Seek(offset, 0)
		if err != nil {
			return err
		}
		if (offset + limit) > fromInfo.Size() {
			tempBuff = make([]byte, int64(math.Abs(float64(offset-fromInfo.Size()))))
		} else {
			tempBuff = make([]byte, limit)
		}
	}

	err = progressBar(fromFile, fileTo, tempBuff, limit)
	if err != nil {
		return err
	}
	return nil
}

func progressBar(fromFile, fileTo *os.File, tempBuff []byte, limit int64) error {
	_, err := fromFile.Read(tempBuff)
	if err != nil {
		return err
	}

	newReader := bytes.NewReader(tempBuff)
	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(newReader)
	bar.Set(pb.Bytes, true)
	_, err = io.CopyN(fileTo, barReader, limit)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return err
		}
	}
	bar.Finish()
	return nil
}
