// Copyright 2016 Arsham Shirvani <arshamshirvani@gmail.com>. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license
// License that can be found in the LICENSE file.

package config

import (
	"bytes"
	"testing"

	"github.com/arsham/expvastic/lib"
	"github.com/spf13/viper"
)

func TestLoadConfiguration(t *testing.T) {
	t.Parallel()
	v := viper.New()
	log := lib.DiscardLogger()
	v.SetConfigType("yaml")

	input := bytes.NewBuffer([]byte(`
    readers:
        reader_1: # populating to get to the passing tests
            interval: 1s
            timeout: 1s
            endpoint: localhost:8200
            backoff: 9
            type_name: erwer
    recorders:
        recorder_1:
            interval: 1s
            timeout: 1s
            endpoint: localhost:8200
            backoff: 9
            index_name: erwer
    routes: blah
    `))
	v.ReadConfig(input)

	readers := map[string]string{"reader_1": "not_exists"}
	recorders := map[string]string{"recorder_1": "elasticsearch"}
	routeMap := map[string]route{"routes": {
		readers:   []string{"reader_1"},
		recorders: []string{"recorder_1"},
	}}
	_, err := loadConfiguration(v, log, routeMap, readers, recorders)
	if _, ok := err.(interface {
		NotSupported()
	}); !ok {
		t.Errorf("want InvalidEndpoint, got (%v)", err)
	}

	readers = map[string]string{"reader_1": "expvar"}
	recorders = map[string]string{"recorder_1": "not_exists"}
	_, err = loadConfiguration(v, log, routeMap, readers, recorders)
	if _, ok := err.(interface {
		NotSupported()
	}); !ok {
		t.Errorf("want InvalidEndpoint, got (%v)", err)
	}

	readers = map[string]string{"reader_1": "expvar"}
	recorders = map[string]string{"recorder_1": "elasticsearch"}
	_, err = loadConfiguration(v, log, routeMap, readers, recorders)
	if err != nil {
		t.Errorf("want (nil), got (%v)", err)
	}

}