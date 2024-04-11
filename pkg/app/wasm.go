//go:build wasm
// +build wasm

package app

import (
	"runtime"

	"github.com/carmel/go-pwa/pkg/errors"
)

func GenerateStaticWebsite(dir string, h *Handler, pages ...string) error {
	panic(errors.New("unsupported instruction").
		WithTag("architecture", runtime.GOARCH))
}

func wasmExecJS() string {
	return ""
}
