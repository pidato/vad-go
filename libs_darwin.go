//go:build darwin
// +build darwin

package vad

/*
#cgo CXXFLAGS: -O3 -Wno-delete-non-virtual-dtor -Wunused-function
#cgo CXXFLAGS: -Wall -fPIC
#cgo CXXFLAGS: -I./
#cgo darwin,amd64 LDFLAGS: -L./mac/amd64
#cgo darwin,arm64 LDFLAGS: -L./mac/arm64
#cgo LDFLAGS: -lfvad
*/
import "C"
