// Copyright 2016 Arsham Shirvani <arshamshirvani@gmail.com>. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license
// License that can be found in the LICENSE file.

package config

import "fmt"

var (
    // EmptyConfigErr is an error when the config file is empty
    EmptyConfigErr = &StructureErr{"", "empty configuration file", nil}
)

// NotSpecifiedErr represents an error when a section is not specified
type NotSpecifiedErr interface {
    NotSpecified()
}

// RoutersErr represents an error when routes are not configured correctly
// The section on this error is the subsection of the route
type RoutersErr interface {
    Routers()
}

// NotSupportedErr is for when something is still not supported
type NotSupportedErr interface {
    NotSupported()
}

// StructureErr represents an error reading the configuration file
type StructureErr struct {
    Section string // The section that error happened
    Reason  string // The reason behind the error
    Err     error  // Err is the error that occurred during the operation.
}

func (e *StructureErr) Error() string {
    if e == nil {
        return "<nil>"
    }

    s := e.Section
    if e.Reason != "" {
        s += " " + e.Reason
    }

    if e.Err != nil {
        s += ": " + e.Err.Error()
    }
    return s
}

type notSpecifiedErr struct{ StructureErr }

func newNotSpecifiedErr(section, reason string, err error) *notSpecifiedErr {
    return &notSpecifiedErr{StructureErr{section, reason, err}}
}

func (e *notSpecifiedErr) NotSpecified() {}
func (e *notSpecifiedErr) Error() string {
    if e == nil {
        return "<nil>"
    }

    s := e.Section
    if e.Reason != "" {
        s += " " + e.Reason
    }
    s += " not specified"

    if e.Err != nil {
        s += ": " + e.Err.Error()
    }
    return s
}

type routersErr struct{ StructureErr }

func newRoutersErr(section, reason string, err error) *routersErr {
    return &routersErr{StructureErr{section, reason, err}}
}

func (routersErr) Routers() {}
func (e *routersErr) Error() string {
    if e == nil {
        return "<nil>"
    }

    s := "not specified: " + e.Section
    if e.Reason != "" {
        s += " " + e.Reason
    }

    if e.Err != nil {
        s += ": " + e.Err.Error()
    }
    return s
}

type notSupportedErr string

func (notSupportedErr) NotSupported() {}
func (n notSupportedErr) Error() string {
    return fmt.Sprintf("%s is not supported", string(n))
}
