// Code generated by "stringer -type=Capability"; DO NOT EDIT.

package device

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VideoCapture-0]
	_ = x[AudioCapture-1]
	_ = x[ApplicationMonitoring-2]
	_ = x[Display-3]
	_ = x[AudioOutput-4]
}

const _Capability_name = "VideoCaptureAudioCaptureApplicationMonitoringDisplayAudioOutput"

var _Capability_index = [...]uint8{0, 12, 24, 45, 52, 63}

func (i Capability) String() string {
	if i < 0 || i >= Capability(len(_Capability_index)-1) {
		return "Capability(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Capability_name[_Capability_index[i]:_Capability_index[i+1]]
}