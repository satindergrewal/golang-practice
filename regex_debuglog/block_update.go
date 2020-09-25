package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"time"
)

var debugLog string = `2020-09-25 06:52:26 UpdateTip: new best=0a9d5c0d9a75aa8e689fccb9f100a63cf35c8bac8720faae10487ce44b3af1bb  height=2066410  log2_work=51.711356  log2_stake=-inf  tx=11557057  date=2020-09-25 06:52:25 progress=1.000000  cache=3.3MiB(7611tx)
2020-09-25 06:55:39 ConnectBlock: wrote 5 block notarisations in block: 0d7651e5a52e7428d94d23c42f341ffca1aa986dcc939fe60d2d14c030d1d2b0
2020-09-25 06:55:39 UpdateTip: new best=0d7651e5a52e7428d94d23c42f341ffca1aa986dcc939fe60d2d14c030d1d2b0  height=2066411  log2_work=51.711357  log2_stake=-inf  tx=11557067  date=2020-09-25 06:55:37 progress=1.000000  cache=3.3MiB(7619tx)
2020-09-25 07:04:17 ConnectBlock: wrote 6 block notarisations in block: 03318bb6514aaf1d965d4526ed9503739bdfca74fbd1fa171b174905f9637b15
2020-09-25 07:04:17 UpdateTip: new best=03318bb6514aaf1d965d4526ed9503739bdfca74fbd1fa171b174905f9637b15  height=2066412  log2_work=51.711359  log2_stake=-inf  tx=11557082  date=2020-09-25 07:04:17 progress=1.000000  cache=3.3MiB(7630tx)
2020-09-25 07:05:20 ConnectBlock: wrote 2 block notarisations in block: 00000000ec98b0606715c46c3405d5baad35e9f548ef7519ae335616205c264c
2020-09-25 07:05:20 UpdateTip: new best=00000000ec98b0606715c46c3405d5baad35e9f548ef7519ae335616205c264c  height=2066413  log2_work=51.71136  log2_stake=-inf  tx=11557116  date=2020-09-25 07:05:12 progress=1.000000  cache=3.3MiB(7637tx)
`

func tail(filename string, out io.Writer) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	info, err := f.Stat()
	if err != nil {
		panic(err)
	}
	oldSize := info.Size()
	for {
		for line, prefix, err := r.ReadLine(); err != io.EOF; line, prefix, err = r.ReadLine() {
			if prefix {
				// fmt.Fprint(out, string(line))
				var expNewBlk = regexp.MustCompile(`(?-s).*UpdateTip.*(?s)`)
				foundBlk := expNewBlk.FindAllString(string(line), -1)
				// fmt.Println(len(foundBlk))
				fmt.Println(foundBlk)
			} else {
				// fmt.Fprintln(out, string(line))
				var expNewBlk = regexp.MustCompile(`(?-s).*UpdateTip.*(?s)`)
				foundBlk := expNewBlk.FindAllString(string(line), -1)
				// fmt.Println(len(foundBlk))
				fmt.Println(foundBlk)
			}
		}
		pos, err := f.Seek(0, io.SeekCurrent)
		if err != nil {
			panic(err)
		}
		for {
			time.Sleep(time.Second)
			newinfo, err := f.Stat()
			if err != nil {
				panic(err)
			}
			newSize := newinfo.Size()
			if newSize != oldSize {
				if newSize < oldSize {
					f.Seek(0, 0)
				} else {
					f.Seek(pos, io.SeekStart)
				}
				r = bufio.NewReader(f)
				oldSize = newSize
				break
			}
		}
	}
}

func main() {
	tail("/Users/satinder/Library/Application Support/Komodo/debug.log", os.Stdout)
	// var expNewBlk = regexp.MustCompile(`(?-s).*UpdateTip.*(?s)`)
	// newBlk := expNewBlk.FindString(debugLog)
	// fmt.Println(newBlk)
	// foundBlk := expNewBlk.FindAllString(logFile, -1)
	// fmt.Println(len(foundBlk))
	// for i, match := range expNewBlk.FindAllString(debugLog, -1) {
	// 	fmt.Println(match, "found at index", i)
	// }
}
