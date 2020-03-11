// +build linux

package vad

/*
#cgo CXXFLAGS: -O2 -Wno-delete-non-virtual-dtor -Wunused-function
#cgo CXXFLAGS: -Wall -fPIC
#cgo CXXFLAGS: -I./
#cgo LDFLAGS: -ldl -luuid -lm -lrt -lpthread -lasound
#cgo LDFLAGS: -L./linux
#cgo LDFLAGS: -lvad
*/
import "C"
