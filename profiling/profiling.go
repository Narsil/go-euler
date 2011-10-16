package profiling

import "time"

func Time(function func() string) int64 {
	timeBefore := time.Nanoseconds()
	function()
	return time.Nanoseconds() - timeBefore
}

func TimeArg(function func(int) string, arg int) (int64, string) {
	timeBefore := time.Nanoseconds()
	result := function(arg)
	return time.Nanoseconds() - timeBefore, result

}
