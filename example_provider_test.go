// Copyright 2016 Arsham Shirvani <arshamshirvani@gmail.com>. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license
// License that can be found in the LICENSE file.

package expipe_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/arsham/expipe/internal"
	"github.com/arsham/expipe/reader"
	reader_testing "github.com/arsham/expipe/reader/testing"
	"github.com/arsham/expipe/recorder"
	recorder_testing "github.com/arsham/expipe/recorder/testing"
)

func getReader(log internal.FieldLogger) (map[string]reader.DataReader, func()) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		desire := `{"the key": "is the value!"}`
		_, err := io.WriteString(w, desire)
		if err != nil {
			panic(err)
		}
	}))

	red, err := reader_testing.New(
		reader.WithLogger(log),
		reader.WithEndpoint(ts.URL),
		reader.WithName("reader_example"),
		reader.WithTypeName("typeName"),
		reader.WithInterval(time.Millisecond*100),
		reader.WithTimeout(time.Second),
		reader.WithBackoff(5),
	)

	if err != nil {
		panic(err)
	}
	red.Pinged = true
	return map[string]reader.DataReader{red.Name(): red}, func() {
		ts.Close()
	}
}

func getRecorder(log internal.FieldLogger, url string) recorder.DataRecorder {
	rec, err := recorder_testing.New(
		recorder.WithLogger(log),
		recorder.WithEndpoint(url),
		recorder.WithName("recorder_example"),
		recorder.WithIndexName("indexName"),
		recorder.WithTimeout(time.Second),
		recorder.WithBackoff(5),
	)
	if err != nil {
		panic(err)
	}
	rec.Pinged = true
	return rec
}
