// +build darwin

package vad

/*
#cgo CXXFLAGS: -O2 -Wno-delete-non-virtual-dtor -Wunused-function
#cgo CXXFLAGS: -Wall -fPIC
#cgo CXXFLAGS: -I./
#cgo LDFLAGS: -ldl -lm -lpthread
#cgo LDFLAGS: -L./mac
#cgo LDFLAGS: -lfvad
*/
import "C"
