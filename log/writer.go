package log

import (
	"bufio"
	"context"
	"io"
	"runtime"
	"strings"
)

// SafeWriter
// @Description: 获取一个io.Writer，方便集成
// @return *io.PipeWriter
func SafeWriter(ctx context.Context) *io.PipeWriter {
	reader, writer := io.Pipe()
	go scan(ctx, reader)
	runtime.SetFinalizer(writer, writerFinalizer)
	return writer
}

func scan(ctx context.Context, reader *io.PipeReader) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(scanLinesOrGiveLong)
	for scanner.Scan() {
		if strings.Replace(scanner.Text(), " ", "", -1) == "" {
			continue
		}

		Infof(ctx, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		Errorf(ctx, "Error while reading from Writer: %s", err)
	}
	_ = reader.Close()
}

const maxTokenLength = bufio.MaxScanTokenSize / 2

func scanLinesOrGiveLong(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanLines(data, atEOF)
	if advance > 0 || token != nil || err != nil {
		return
	}
	if len(data) < maxTokenLength {
		return
	}
	return maxTokenLength, data[0:maxTokenLength], nil
}

func writerFinalizer(writer *io.PipeWriter) {
	_ = writer.Close()
}
