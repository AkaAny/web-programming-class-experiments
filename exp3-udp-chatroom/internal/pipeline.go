package internal

import (
	"encoding/binary"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/frame"
	"math"
)

func WithProtocol(pipeline netty.Pipeline, fn func(pipeline netty.Pipeline)) {
	pipeline.AddLast(frame.LengthFieldCodec(binary.BigEndian, math.MaxInt, 0, 8, 0, 8))
	//pipeline.AddLast(frame.LengthFieldPrepender(binary.BigEndian, 8, 0, false))
	// decode to map[string]interface{}
	// session recorder.
	fn(pipeline)
	// chat handler.

}
