package vad

/*
#include "fvad.h"
*/
import "C"
import (
	"errors"
	"reflect"
	"sync"
	"unsafe"
)

type Result C.int

const (
	Active    = Result(1)
	NonActive = Result(0)
	Invalid   = Result(1)
)

type Mode C.int

const (
	Quality        = Mode(0)
	LowBitrate     = Mode(1)
	Aggressive     = Mode(2)
	VeryAggressive = Mode(3)
)

type VAD struct {
	ptr    *C.Fvad
	mu     sync.Mutex
	closed bool
}

func New() *VAD {
	inst := C.fvad_new()
	if inst == nil {
		return nil
	}
	return &VAD{
		ptr:    inst,
		mu:     sync.Mutex{},
		closed: false,
	}
}

//
// Re-initializes a VAD instance, clearing all state and resetting mode and
// sample rate to defaults.
//
func (v *VAD) Reset() {
	v.mu.Lock()
	defer v.mu.Unlock()
	C.fvad_reset(v.ptr)
}

//
// Changes the VAD operating ("aggressiveness") mode of a VAD instance.
//
// A more aggressive (higher mode) VAD is more restrictive in reporting speech.
// Put in other words the probability of being speech when the VAD returns 1 is
// increased with increasing mode. As a consequence also the missed detection
// rate goes up.
//
// Valid modes are 0 ("quality"), 1 ("low bitrate"), 2 ("aggressive"), and 3
// ("very aggressive"). The default mode is 0.
//
// Returns 0 on success, or -1 if the specified mode is invalid.
//
func (v *VAD) SetMode(mode Mode) bool {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.closed {
		return false
	}
	return C.fvad_set_mode(v.ptr, (C.int)(mode)) == 0
}

//
// Sets the input sample rate in Hz for a VAD instance.
//
// Valid values are 8000, 16000, 32000 and 48000. The default is 8000. Note
// that internally all processing will be done 8000 Hz; input data in higher
// sample rates will just be downsampled first.
//
// Returns 0 on success, or -1 if the passed value is invalid.
//
func (v *VAD) SetSampleRate(sampleRate int32) bool {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.closed {
		return false
	}
	return C.fvad_set_sample_rate(v.ptr, (C.int)(sampleRate)) == 0
}

//
// Calculates a VAD decision for an audio frame.
//
// `frame` is an array of `length` signed 16-bit samples. Only frames with a
// length of 10, 20 or 30 ms are supported, so for example at 8 kHz, `length`
// must be either 80, 160 or 240.
//
// Returns              : 1 - (active voice),
//                        0 - (non-active Voice),
//                       -1 - (invalid frame length).
//
func (v *VAD) Process(data []int16) Result {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.closed {
		return Invalid
	}
	return Result(
		C.fvad_process(
			v.ptr,
			(*C.int16_t)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data)).Data)),
			(C.size_t)(uint(len(data)))))
}

// Closes and frees up native resources.
func (v *VAD) Close() error {
	v.mu.Lock()
	if v.closed {
		v.mu.Unlock()
		return nil
	}
	v.closed = true
	v.mu.Unlock()

	C.fvad_free(v.ptr)
	return nil
}

func DownSampleBy2(in []int16, out []int16, state *[8]int32) error {
	if len(out) != len(in)/2 {
		return errors.New("out must be have the size of in")
	}
	C.fvad_downsample_by_2(
		(*C.int16_t)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&in)).Data)),
		(*C.int16_t)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&out)).Data)),
		(*C.int32_t)(unsafe.Pointer(state)),
		C.size_t(len(in)),
	)
	return nil
}

func UpSampleBy2(in []int16, out []int16, state *[8]int32) error {
	if len(out) != len(in)*2 {
		return errors.New("out must be have the size of in")
	}
	C.fvad_upsample_by_2(
		(*C.int16_t)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&in)).Data)),
		C.size_t(len(in)),
		(*C.int16_t)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&out)).Data)),
		(*C.int32_t)(unsafe.Pointer(state)),
	)
	return nil
}
