package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	lineCnt := flag.Uint("n", 10, "output the last {n} lines")
	flag.Parse()

	if *lineCnt == 0 {
		return
	}

	if len(flag.Args()) != 1 {
		log.Fatalln("usage: go run main.go [-n] filename")
	}

	filename := flag.Arg(0)

	lines, err := createLines(filename, int(*lineCnt))
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < len(lines); i++ {
		println(lines[i])
	}
}

// ファイルを読み込み、所定行数分の文字列リストを生成する
func createLines(filename string, lineCnt int) ([]string, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	fileStat, err := fp.Stat()
	if err != nil {
		return nil, err
	}

	// 対象ファイルのファイルサイズを取得
	pos := fileStat.Size()
	line := ""
	lines := []string{}
	// EOFから遡って読み込み
	b := make([]byte, 1)
	for pos > 0 {
		_, err := fp.Seek(pos-1, io.SeekStart)
		if err != nil {
			return nil, err
		}

		_, err = fp.Read(b)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
		}

		if pos == fileStat.Size() {
			pos--
			continue
		}

		if b[0] == '\n' {
			lines = append([]string{line}, lines[:]...)
			// 必要行数分取得したら離脱
			if len(lines) == lineCnt {
				break
			}
			line = ""
		} else {
			line = string(b) + line
		}

		pos--

		if pos == 0 {
			lines = append([]string{line}, lines[0:]...)
		}
	}

	return lines, nil
}
