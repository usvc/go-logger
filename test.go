package logger

import (
	"bytes"
	"io"
	"os"
	"sync"
)

// captureStdoutFrom is used in tests only to confirm output
// from the fmt package
func captureStdoutFrom(thisFunction func()) string {
	stdoutReader, stdoutWriter, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	originalStdout := os.Stdout
	defer func() {
		os.Stdout = originalStdout
		if err = stdoutReader.Close(); err != nil {
			panic(err)
		}
	}()
	os.Stdout = stdoutWriter
	capturedOutput := make(chan string)
	waiter := new(sync.WaitGroup)
	waiter.Add(1)
	go func() {
		var output bytes.Buffer
		waiter.Done()
		io.Copy(&output, stdoutReader)
		capturedOutput <- output.String()
	}()
	thisFunction()
	waiter.Wait()
	if err = stdoutWriter.Close(); err != nil {
		panic(err)
	}
	return <-capturedOutput
}

// captureStderrFrom is used in tests only to confirm output
// from the fmt package
func captureStderrFrom(thisFunction func()) string {
	stderrReader, stderrWriter, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	originalStderr := os.Stderr
	defer func() {
		os.Stderr = originalStderr
		if err = stderrReader.Close(); err != nil {
			panic(err)
		}
	}()
	os.Stderr = stderrWriter
	capturedOutput := make(chan string)
	waiter := new(sync.WaitGroup)
	waiter.Add(1)
	go func() {
		var output bytes.Buffer
		waiter.Done()
		io.Copy(&output, stderrReader)
		capturedOutput <- output.String()
	}()
	waiter.Wait()
	thisFunction()
	if err = stderrWriter.Close(); err != nil {
		panic(err)
	}
	return <-capturedOutput
}
