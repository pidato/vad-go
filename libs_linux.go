//go:build linux
// +build linux

package vad

/*
#cgo CXXFLAGS: -O3 -Wno-delete-non-virtual-dtor -Wunused-function
#cgo CXXFLAGS: -Wall -fPIC
#cgo CXXFLAGS: -I./
#cgo linux,amd64 LDFLAGS: -L./linux/amd64
#cgo LDFLAGS: -lfvad
*/
import "C"
