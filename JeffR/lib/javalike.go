package lib

// provides an approximation of good ol' Java ternary operator:
//
//	<condBool> ? <valueIfTrue> : <valueIfFalse>
//
// the equivalent of
//
//	if <condBool> {
//		return <valueIfTrue>
//	} else {
//		return <valueIfFalse>
//	}
//
// (in fact, this is the impl in golang)
// *BUT* is nestable inside another expression e.g.
//
//	fmt.Printf("%s", ternary(flag,"foo","bar"))
//
// vs
//
//	     if flag {
//				fmt.Print("foo")
//			} else {
//				fmt.Print("bar")
//			}
//
// with one *MAJOR* difference: the <valueIfTrue> and <valueIfFalse>
// are both eval'd regardless of the <condBool> outcome
// SOOO don't use anything that has side-effects for 2nd or 3rd args!!!
//
// # Java:
//
//	int y = 1;
//	int z = 1;
//	int a = y < 10 ? y++ : z++;
//	System.out.println( "y=" + y + " z=" + z );
//	// console output: y=2 z=1
//
// # Golang:
//
//	y := 1
//	z := 1
//	a := ternary( y < 10, y++, z++ )
//	fmt.Printf( "y=%d z=%d\n", y, z)
//	// console output: y=2 z=2
func Ternary[B bool, V int | int64 | float32 | float64 | string](cond bool, ifTrue V, ifFalse V) V {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}

// the rough equivalent of java assert,
// minus the nifty auto-stringification of check expression (1st arg)
func Assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}
