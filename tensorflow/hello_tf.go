// tensorflow example file

package main

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
	"fmt"
)

func main() {
	// Construct a graph with an operation that produces a string constant.
	s := op.NewScope()
	c := op.Const(s, "Hello from Tensorflow version" + tf.Version())
	graph, err := s.Finalize()
	if err != nil {
		panic(err)
	}

	// Execute the graph in a session
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		panic(err)
	}

	output, err := session.Run(nil, []tf.Output{c}, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(output[0].Value())
}