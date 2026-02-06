package tunnel

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type Tunnel struct {
	src int
	dst int

	listener net.Listener
}

func New(src, dst int) *Tunnel {
	return &Tunnel{
		src: src,
		dst: dst,
	}
}

func (t *Tunnel) Init() error {
	log.Printf("creating listener on tcp 0.0.0.0:%d", t.src)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", t.src))
	if err != nil {
		return err
	}

	t.listener = ln

	return nil
}

func (t *Tunnel) Start() {
	for {
		srcConn, err := t.listener.Accept()
		if err != nil {
			log.Printf("listener.Accept returned error: %s", err)
			continue
		}

		log.Printf("got connection from %v", srcConn.RemoteAddr())

		dstConn, err := net.Dial("tcp", fmt.Sprintf(":%d", t.dst))
		if err != nil {
			srcConn.Close()
			log.Printf("failure for net.Dial on port %d: %s", t.dst, err)
			continue
		}

		log.Printf("created connections with %s and %s", srcConn.RemoteAddr(), dstConn.RemoteAddr())

		go t.bidirectionalCopy(srcConn, dstConn)
	}
}

func (t *Tunnel) bidirectionalCopy(srcConn, dstConn net.Conn) {
	var once sync.Once
	closeBoth := func() {
		log.Printf("closing connections with %s & %s", srcConn.RemoteAddr(), dstConn.RemoteAddr())
		_ = srcConn.Close()
		_ = dstConn.Close()
	}

	defer once.Do(closeBoth)

	var wg sync.WaitGroup
	wg.Add(2)

	copy := func(src, dst net.Conn) {
		defer wg.Done()

		n, err := io.Copy(src, dst)
		if err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, net.ErrClosed) {
				log.Printf("%s -> %s: connections are closed", src.RemoteAddr(), dst.RemoteAddr())
			} else {
				log.Printf("%s -> %s: io.Copy error: %s", src.RemoteAddr(), dst.RemoteAddr(), err)
			}
		}

		log.Printf("copied %s -> %s total: %s", src.RemoteAddr(), dst.RemoteAddr(), formatBytes(n))
		once.Do(closeBoth)
	}

	// Copy data from src to dst
	go copy(srcConn, dstConn)

	// Copy data from dst to src
	go copy(dstConn, srcConn)

	wg.Wait() // Wait for both copy operations to finish

	log.Printf("connections with %s & %s are destroyed. returning from bidirectionalCopy", srcConn.RemoteAddr(), dstConn.RemoteAddr())
}

func formatBytes(bytes int64) string {
	kb := float64(bytes) / 1024
	if kb < 1 {
		return fmt.Sprintf("%d B", bytes)
	}

	mb := kb / 1024
	if mb < 1 {
		return fmt.Sprintf("%f KB", kb)
	}

	gb := mb / 1024
	if gb < 1 {
		return fmt.Sprintf("%f MB", mb)
	}

	return fmt.Sprintf("%f GB", gb)
}
