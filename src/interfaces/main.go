package main

import "fmt"
import "io"
import "bytes"

type Writer interface {
  Write ([]byte) (int, error)
}

type Closer interface {
  Close () error
}

// composed interface: combination of two or more interfaces
type WriterCloser interface {
  Writer
  Closer
}

type ConsoleWriter struct {}

type BufferedWriterCloser struct {
  buffer *bytes.Buffer
}

func NewBufferedWriterCloser () *BufferedWriterCloser {
  return &BufferedWriterCloser {
    buffer: bytes.NewBuffer([]byte{}),
  }
}

// implementation

// value-type implementation
func (cw ConsoleWriter) Write(data []byte) (int, error) {
  n, err := fmt.Println(string(data))
  return n, err
}

// pointer-type implementation
func (bwc *BufferedWriterCloser) Write (data []byte) (int, error) {
  n, err := bwc.buffer.Write(data)
  if err != nil {
    return 0, err
  }

  v := make([]byte, 8)
  for bwc.buffer.Len() > 8 {
    _, err = bwc.buffer.Read(v)
    if err != nil {
      return 0, err
    }
    _, err = fmt.Println(string(v))
    if err != nil {
      return 0, err
    }
  }
  return n, nil
}

func (bwc *BufferedWriterCloser) Close () error {
  for bwc.buffer.Len() > 0 {
    data := bwc.buffer.Next(8)
    _, err := fmt.Println(string(data))
    if err != nil {
      return err
    }
  }
  return nil
}

// what is an interface?
// an interface is a contract that defined the behaviour of an 'object'

// use cases
// for defining the behaviour of an 'object' in go

// empty interfaces
// these are interfaces defined on the fly that have no methods on it

// best practice
// if the interface has one method (which isnamed after it),
// the interface name should end with 'er'
// e.g: Write -> Writer

func main () {
  var writer Writer = ConsoleWriter {}
  writer.Write([]byte("Hello world!"))

  var writerCloser WriterCloser = NewBufferedWriterCloser()
  writerCloser.Write([]byte("Hello world"))
  writerCloser.Close()

  // converting an interface to a another interface (could panic if the operation is not possible)
  var bwc *BufferedWriterCloser = writerCloser.(*BufferedWriterCloser)
  fmt.Println(bwc)

  // converting an interface to another interface (without panicing)
  r, ok := writerCloser.(io.Reader)
  if ok {
    fmt.Println(r)
  } else {
    fmt.Println("Conversion failed!")
  }
}
