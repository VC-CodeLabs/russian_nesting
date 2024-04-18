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

func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

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

	var wtt, rtt int64 = 0, 0

	for r := 0; r < TEST_RUNS; r++ {
		fmt.Println("Writing test [][w#,h#] array...")
		ts := start()
		envs := make(EnvelopesA, 0)
		for i := 0; i < ENV_MAX; i++ {
			envs = append(envs, EnvArray{i + 1, i + 1})

		}
		td := finish(ts)
		fmt.Printf("Array with %d envelopes [w#,h#] made in %15s\n", len(envs), sanitizeDuration(td))
		wtt += td.Nanoseconds()

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
		rtt += rtd.Nanoseconds()
	}
	wta := time.Duration(float64(wtt/TEST_RUNS) * float64(time.Nanosecond))
	rta := time.Duration(float64(rtt/TEST_RUNS) * float64(time.Nanosecond))
	fmt.Printf("*** Average time to make [][w#,h#] %15s read %15s in %d runs\n", sanitizeDuration(wta), sanitizeDuration(rta), TEST_RUNS)
}

func testEnvelopesStructOps() {

	var wtt, rtt int64 = 0, 0
	for r := 0; r < TEST_RUNS; r++ {
		fmt.Println("Writing test []{w#,h#} array...")
		ts := start()
		envs := make(EnvelopesS, 0)
		for i := 0; i < ENV_MAX; i++ {
			envs = append(envs, EnvStruct{i + 1, i + 1})

		}
		td := finish(ts)
		fmt.Printf("Array with %d envelopes {w#,h#} made in %15s\n", len(envs), sanitizeDuration(td))
		wtt += td.Nanoseconds()

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
		rtt += rtd.Nanoseconds()

	}
	wta := time.Duration(float64(wtt/TEST_RUNS) * float64(time.Nanosecond))
	rta := time.Duration(float64(rtt/TEST_RUNS) * float64(time.Nanosecond))
	fmt.Printf("*** Average time to make []{w#,h#} %15s read %15s in %d runs\n", sanitizeDuration(wta), sanitizeDuration(rta), TEST_RUNS)

	var envA = EnvelopesA{{1, 2}, {3, 4}}
	var envS = EnvelopesS{{1, 2}, {3, 4}}

	assert(len(envA) == len(envS), "mismatched lengths")
}
