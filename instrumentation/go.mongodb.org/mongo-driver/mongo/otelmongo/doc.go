// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package otelmongo instruments go.mongodb.org/mongo-driver/mongo.
//
// This package is compatible with v0.2.0 of
// go.mongodb.org/mongo-driver/mongo.
//
// NewMonitor will return an event.CommandMonitor which is used to trace
// requests.
//
// This code was originally based on the following:
//   - https://github.com/DataDog/dd-trace-go/tree/02f0449efa3cb382d499fadc873957385dcb2192/contrib/go.mongodb.org/mongo-driver/mongo
//   - https://github.com/DataDog/dd-trace-go/tree/v1.23.3/ddtrace/ext
//
// The "OTEL_SEMCONV_STABILITY_OPT_IN" environment variable can be used to opt
// into the latest semantic conventions:
//   - "database": emit the latest semantic conventions
//   - "": emit v1.21.0 (default) semantic conventions
//   - "database/dup": emit v1.21.0 (default) and the latest semantic
//     conventions
//
// By default, otelmongo only emits v1.21.0.
package otelmongo // import "go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
