// Copyright 2016 Arsham Shirvani <arshamshirvani@gmail.com>. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license
// License that can be found in the LICENSE file.

package expvar

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/arsham/expvastic/lib"
	"github.com/arsham/expvastic/reader"
)

func TestExpvarReaderErrors(t *testing.T) {
	log := lib.DiscardLogger()
	ctx, cancel := context.WithCancel(context.Background())
	ctxReader := reader.NewMockCtxReader("nowhere")
	ctxReader.ContextReadFunc = func(ctx context.Context) (*http.Response, error) {
		return nil, fmt.Errorf("Error")
	}
	rdr, _ := NewExpvarReader(log, ctxReader, "my_reader", time.Second, time.Second)
	rdr.Start(ctx)
	defer cancel()

	rdr.JobChan() <- ctx
	select {
	case res := <-rdr.ResultChan():
		if res.Res != nil {
			t.Errorf("expecting no results, got(%v)", res.Res)
		}
		if res.Err == nil {
			t.Error("expecting error, got nothing")
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("expecting an error result back, got nothing")
	}
	cancel()
}

func TestExpvarReaderClosesStream(t *testing.T) {
	log := lib.DiscardLogger()
	ctxReader := reader.NewMockCtxReader("nowhere")
	ctx, cancel := context.WithCancel(context.Background())
	rdr, _ := NewExpvarReader(log, ctxReader, "my_reader", time.Second, time.Second)
	done := rdr.Start(ctx)
	rdr.JobChan() <- ctx

	select {
	case <-rdr.ResultChan():
	default:
		cancel()
	}
	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Error("The channel was not closed in time")
	}
}