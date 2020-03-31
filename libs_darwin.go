// +build darwin

package vad

/*
#cgo CXXFLAGS: -O3 -Wno-delete-non-virtual-dtor -Wunused-function
#cgo CXXFLAGS: -Wall -fPIC
#cgo CXXFLAGS: -I./
#cgo LDFLAGS: -L./mac
#cgo LDFLAGS: -lfvad
*/
import "C"
