package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

type bres struct {
	tInit  time.Duration
	tStart time.Duration
	tBcast time.Duration
	tDone  time.Duration
}

func doWork(val int) {
	// r := rand.NewSource(int64(val))
	// buflen := 1024 * 32
	// b := make([]byte, buflen)
	// for i := 0; i < buflen; i += 8 {
	// 	v := r.Int63()
	// 	b[i+0] = byte((v >> (0 * 8)) % 0xff)
	// 	b[i+1] = byte((v >> (1 * 8)) % 0xff)
	// 	b[i+2] = byte((v >> (2 * 8)) % 0xff)
	// 	b[i+3] = byte((v >> (3 * 8)) % 0xff)
	// 	b[i+4] = byte((v >> (4 * 8)) % 0xff)
	// 	b[i+5] = byte((v >> (5 * 8)) % 0xff)
	// 	b[i+6] = byte((v >> (6 * 8)) % 0xff)
	// 	b[i+7] = byte((v >> (7 * 8)) % 0xff)
	// }
	// sha512.Sum512(b)
}

func main() {
	buf := bytes.NewBufferString("")
	for i := 0; i < 512; i++ {
		ra := methodA(16, 1024*8)
		rb := methodB(16, 1024*8)
		// log.Printf("init % 15s\tstart % 15s\tbcast % 15s\tdone % 15s", r.tInit, r.tStart, r.tBcast, r.tDone)
		buf.WriteString(fmt.Sprint(i, "\t", ra.tBcast.Microseconds(), "\t", rb.tBcast.Microseconds(), "\n"))
	}
	log.Println(os.WriteFile("results.dat", buf.Bytes(), 0644))
}
