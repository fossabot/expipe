// Copyright 2016 Arsham Shirvani <arshamshirvani@gmail.com>. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license
// License that can be found in the LICENSE file.

package config

import (
    "context"
    "time"

    "github.com/Sirupsen/logrus"
    "github.com/arsham/expvastic/reader"
    "github.com/arsham/expvastic/recorder"
)

// Conf will return necessary information for setting up an endpoint, for readers or recorders.
type Conf interface {
    Endpoint() string
    Interval() time.Duration
    Timeout() time.Duration
    Backoff() int
    Logger() logrus.FieldLogger
    //TODO: Add TLS stuff
}

// ReaderConf defines a behaviour of a reader.
type ReaderConf interface {
    Conf

    // NewInstance should return an intialised Reader instance.
    NewInstance(context.Context) (reader.TargetReader, error)
}

// RecorderConf defines a behaviour that requies the recorders to have the concept
// of Index and Type. If any of these are not applicable, just return an empty string.
type RecorderConf interface {
    Conf

    // NewInstance should return an intialised Recorder instance.
    NewInstance(context.Context) (recorder.DataRecorder, error)

    // IndexName comes from the configuration, but the engine takes over.
    // Recorders should not intercept the engine for its decision, unless they have a
    // valid reason.
    IndexName() string

    // TypeName is usually the application name.
    // Recorders should not intercept the engine for its decision, unless they have a
    // valid reason.
    TypeName() string
}