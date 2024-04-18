package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const DIM_MIN = 1

var DIM_MAX = int(math.Pow(10, 5))

const ENV_MIN = 1

var ENV_MAX = DIM_MAX

type EnvStruct struct {
	width  int
	height int
}

type EnvelopesS []EnvStruct

type EnvArray [2]int

type EnvelopesA [][2]int

func main() {
	testEnvelopeOps()
}

func start() time.Time {
	return time.Now()
}

func finish(started time.Time) time.Duration {
	return time.Since(started)
}

// the rough equivalent of java assert,
// minus the nifty auto-stringification of check expression (1st arg)
func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

// sanitizes stringified duration so the microsecond indicator
// can pass thru console pipes like find cmd;
// otherwise, it prints as "#┬╡s" instead of "#µs"
func sanitizeDuration(d time.Duration) string {
	s := fmt.Sprintf("%s", d)
	s = strings.Replace(s, "µs", "us", 1)
	return s
}

//////////////////////////////////////////////////////////////////
// tests
//

const TEST_RUNS = 100

func testEnvelopeOps() {

	testEnvelopesArrayOps()

	testEnvelopesStructOps()

}

func testEnvelopesArrayOps() {

	var wtt, wtc, rtt, rtc int64 = 0, 0, 0, 0

	for r := 0; r < TEST_RUNS; r++ {
		fmt.Println("Writing test [][w#,h#] array...")
		ts := start()
		envs := make(EnvelopesA, 0)
		for i := 0; i < ENV_MAX; i++ {
			envs = append(envs, EnvArray{i + 1, i + 1})

		}
		td := finish(ts)
		fmt.Printf("Array with %d envelopes [w#,h#] made in %15s\n", len(envs), sanitizeDuration(td))
		wns := td.Nanoseconds()
		wtt += wns
		if wns > 0 {
			wtc++
		}

		fmt.Println("Reading test [][w#,h#] array...")
		rts := start()
		for i := 0; i < len(envs); i++ {
			env := envs[i]
			envW := env[0]
			envH := env[1]
			// be sure the read doesn't get optimized out!!!
			if envW > DIM_MAX || envH > DIM_MAX {
				fmt.Println("Never!!!")
			}
		}
		rtd := finish(rts)
		fmt.Printf("Array with %d envelopes [w#,h#] read in %15s\n", len(envs), sanitizeDuration(rtd))
		rns := rtd.Nanoseconds()
		rtt += rns
		if rns > 0 {
			rtc++
		}
	}
	wtd := time.Duration(float64(wtt) * float64(time.Nanosecond))
	wta := time.Duration(float64(wtt/TEST_RUNS) * float64(time.Nanosecond))
	wtz := time.Duration(float64(wtt/max(wtc, 1)) * float64(time.Nanosecond))
	fmt.Printf("*** Average time to make [%d][w#,h#] %15s / %d = %15s (!0 %15s x %d)\n",
		ENV_MAX, sanitizeDuration(wtd), TEST_RUNS, sanitizeDuration(wta), sanitizeDuration(wtz), wtc)

	rtd := time.Duration(float64(rtt) * float64(time.Nanosecond))
	rta := time.Duration(float64(rtt/TEST_RUNS) * float64(time.Nanosecond))
	rtz := time.Duration(float64(rtt/max(rtc, 1)) * float64(time.Nanosecond))
	fmt.Printf("*** Average time to read [%d][w#,h#] %15s / %d = %15s (!0 %15s x %d)\n",
		ENV_MAX, sanitizeDuration(rtd), TEST_RUNS, sanitizeDuration(rta), sanitizeDuration(rtz), rtc)
}

func testEnvelopesStructOps() {

	var wtt, wtc, rtt, rtc int64 = 0, 0, 0, 0
	for r := 0; r < TEST_RUNS; r++ {
		fmt.Println("Writing test []{w#,h#} array...")
		ts := start()
		envs := make(EnvelopesS, 0)
		for i := 0; i < ENV_MAX; i++ {
			envs = append(envs, EnvStruct{i + 1, i + 1})

		}
		td := finish(ts)
		fmt.Printf("Array with %d envelopes {w#,h#} made in %15s\n", len(envs), sanitizeDuration(td))
		wns := td.Nanoseconds()
		wtt += wns
		if wns > 0 {
			wtc++
		}

		fmt.Println("Reading test []{w#,h#} array...")
		rts := start()
		for i := 0; i < len(envs); i++ {
			env := envs[i]
			envW := env.width
			envH := env.height
			// be sure the read doesn't get optimized out!!!
			if envW > DIM_MAX || envH > DIM_MAX {
				fmt.Println("Never!!!")
			}
		}
		rtd := finish(rts)
		fmt.Printf("Array with %d envelopes {w#,h#} read in %15s\n", len(envs), sanitizeDuration(rtd))
		rns := rtd.Nanoseconds()
		rtt += rns
		if rns > 0 {
			rtc++
		}

	}

	wtd := time.Duration(float64(wtt) * float64(time.Nanosecond))
	wta := time.Duration(float64(wtt/TEST_RUNS) * float64(time.Nanosecond))
	wtz := time.Duration(float64(wtt/max(wtc, 1)) * float64(time.Nanosecond))
	fmt.Printf("*** Average time to make [%d]{w#,h#} %15s / %d = %15s (!0 %15s x %d)\n",
		ENV_MAX, sanitizeDuration(wtd), TEST_RUNS, sanitizeDuration(wta), sanitizeDuration(wtz), wtc)

	rtd := time.Duration(float64(rtt) * float64(time.Nanosecond))
	rta := time.Duration(float64(rtt/TEST_RUNS) * float64(time.Nanosecond))
	rtz := time.Duration(float64(rtt/max(rtc, 1)) * float64(time.Nanosecond))
	fmt.Printf("*** Average time to read [%d]{w#,h#} %15s / %d = %15s (!0 %15s x %d)\n",
		ENV_MAX, sanitizeDuration(rtd), TEST_RUNS, sanitizeDuration(rta), sanitizeDuration(rtz), rtc)

	var envA = EnvelopesA{{1, 2}, {3, 4}}
	var envS = EnvelopesS{{1, 2}, {3, 4}}

	assert(len(envA) == len(envS), "mismatched lengths")
}
