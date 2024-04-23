package librn

//
// define constants for the upper & lower limits of the problem space
//

const DIM_MIN = 1 // the minimum value for both envelope width and height (must be corporeal :))

const DIM_MAX = int(1e5) // the maximum value for both envelope width and height
//// interesting:
//// other approaches are available for DIM_MAX definition,
//// but they require a var, not a const; this would also
//// force ENV_MAX to be a var or repeat the definition
// var DIM_MAX = int(math.Pow(10, 5))
// var DIM_MAX = int(math.Pow10(5))

const ENV_MIN = 1 // the minimum # of envelopes in input

const ENV_MAX = DIM_MAX // the maximum value of envelopes in input, matching DIM_MAX
