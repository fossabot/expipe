// Copyright 2016 Arsham Shirvani <arshamshirvani@gmail.com>. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license
// License that can be found in the LICENSE file.

package testing

import (
	"time"

	"github.com/arsham/expipe/internal"
	"github.com/arsham/expipe/reader"
)

// GetReader provides a SimpleReader for using in the example.
func GetReader(url string) *Reader {
	log := internal.DiscardLogger()
	red, err := New(
		reader.WithLogger(log),
		reader.WithEndpoint(url),
		reader.WithName("reader_example"),
		reader.WithTypeName("reader_example"),
		reader.WithInterval(10*time.Millisecond),
		reader.WithTimeout(time.Second),
		reader.WithBackoff(10),
	)
	if err != nil {
		panic(err)
	}
	return red
}
