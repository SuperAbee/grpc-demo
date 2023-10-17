package transport

import (
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func NewFileListener(listenPath, interestFilename string) *FileListener {
	return &FileListener{
		listenPath:       listenPath,
		interestFilename: interestFilename,
		alreadyAccept:    make(map[string]struct{}),
	}
}

type FileListener struct {
	listenPath       string
	interestFilename string
	alreadyAccept    map[string]struct{}
}

func (f *FileListener) Accept() (net.Conn, error) {
	var conn *FileConn
	for {
		err := filepath.Walk(f.listenPath, func(path string, fi os.FileInfo, err error) error {
			if fi == nil {
				return err
			}
			if fi.IsDir() {
				return nil
			}
			if strings.HasPrefix(fi.Name(), f.interestFilename) {
				if _, ok := f.alreadyAccept[fi.Name()]; ok {
					return nil
				}
				log.Printf("accept new conn %s<--->%s", fi.Name(), ClientName(fi.Name()))
				f.alreadyAccept[fi.Name()] = struct{}{}
				conn = NewFileConn(f.listenPath+"/"+fi.Name(), f.listenPath+"/"+ClientName(fi.Name()))
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
		if conn != nil {
			return conn, nil
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func (f *FileListener) Close() error {
	return nil
}

func (f *FileListener) Addr() net.Addr {
	path, err := filepath.Abs(f.listenPath)
	if err != nil {
		panic(err)
	}
	return &FileAddr{
		path: path,
	}
}
