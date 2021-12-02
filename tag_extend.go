package json

var avoidOmitempty bool

func AvoidOmit() {
	avoidOmitempty = true
}

func tagFilter(om bool) bool {
	if avoidOmitempty {
		return false
	}
	return om
}