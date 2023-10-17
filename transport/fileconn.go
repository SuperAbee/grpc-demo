package transport

import (
	"context"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"
)

func CustomeDialer(ctx context.Context, addr string) (net.Conn, error) {
	return NewFileConn("../ch/"+ClientName(addr), "../ch/"+addr), nil
}

type FileAddr struct {
	path string
}

func (f *FileAddr) Network() string {
	return "file"
}

func (f *FileAddr) String() string {
	return "file://" + f.path
}

func NewFileConn(local, remote string) *FileConn {
	read, err := os.OpenFile(local, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	write, err := os.OpenFile(remote, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	return &FileConn{
		readFile:  read,
		writeFile: write,
	}
}

type FileConn struct {
	readFile  *os.File
	writeFile *os.File
}

func (f *FileConn) Read(b []byte) (n int, err error) {
	for {
		n, err = f.readFile.Read(b)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			time.Sleep(50 * time.Millisecond)
			continue
		}

		log.Printf("read %d bytes", n)
		// err = f.readFile.Truncate(0)
		// if err != nil {
		// 	panic(err)
		// }
		// _, err = f.readFile.Seek(0, 0)
		// if err != nil {
		// 	panic(err)
		// }
		break
	}
	return
}

func (f *FileConn) Write(b []byte) (n int, err error) {
	log.Printf("write %d bytes", len(b))
	return f.writeFile.Write(b)
}

func (f *FileConn) Close() error {
	return nil
}

func (f *FileConn) LocalAddr() net.Addr {
	path, err := filepath.Abs(f.readFile.Name())
	if err != nil {
		panic(err)
	}
	return &FileAddr{
		path: path,
	}
}

func (f *FileConn) RemoteAddr() net.Addr {
	path, err := filepath.Abs(f.writeFile.Name())
	if err != nil {
		panic(err)
	}
	return &FileAddr{
		path: path,
	}
}

func (f *FileConn) SetDeadline(t time.Time) error {
	return nil
}

func (f *FileConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (f *FileConn) SetWriteDeadline(t time.Time) error {
	return nil
}
