// Copyright 2016 Arsham Shirvani <arshamshirvani@gmail.com>. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license
// License that can be found in the LICENSE file.

package testing_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/arsham/expvastic/lib"
	"github.com/arsham/expvastic/reader"
	reader_test "github.com/arsham/expvastic/reader/testing"
)

var (
	log        logrus.FieldLogger
	testServer *httptest.Server
)

func TestMain(m *testing.M) {
	log = lib.DiscardLogger()
	testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	exitCode := m.Run()
	testServer.Close()
	os.Exit(exitCode)
}

type Construct struct {
	name     string
	typeName string
	endpoint string
	interval time.Duration
	timeout  time.Duration
	backoff  int
}

func (c *Construct) SetName(name string)                { c.name = name }
func (c *Construct) SetTypename(typeName string)        { c.typeName = typeName }
func (c *Construct) SetEndpoint(endpoint string)        { c.endpoint = endpoint }
func (c *Construct) SetInterval(interval time.Duration) { c.interval = interval }
func (c *Construct) SetTimeout(timeout time.Duration)   { c.timeout = timeout }
func (c *Construct) SetBackoff(backoff int)             { c.backoff = backoff }
func (c *Construct) TestServer() *httptest.Server       { return testServer }
func (c *Construct) Object() (reader.DataReader, error) {
	return reader_test.New(log, c.endpoint, c.name, c.typeName, c.interval, c.timeout, c.backoff)
}

func TestSimpleReader(t *testing.T) {
	reader_test.TestReader(t, &Construct{})
}