// +build linux

package vad

/*
#cgo CXXFLAGS: -O3 -Wno-delete-non-virtual-dtor -Wunused-function
#cgo CXXFLAGS: -Wall -fPIC
#cgo CXXFLAGS: -I./
#cgo LDFLAGS: -L./linux
#cgo LDFLAGS: -lfvad
*/
import "C"
