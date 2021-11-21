package main

import (
	"errors"
	"io"
	"math"
	"os"
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
	_, err = fromFile.Read(tempBuff)
	if err != nil {
		return err
	}
	_, err = fileTo.Write(tempBuff)
	if err != nil {
		return err
	}
	return nil
}
