package collection

import (
	"reflect"

	"github.com/xh3b4sd/anna/spec"
)

// DefaultArg represents an identifier that, when used, signals to use the
// default argument, if any given. DefaultArg can be used as follows.
//
//     args := []interface{}{"foo", DefaultArg{}, "baz"}
//     def := ArgToString(args, 1, "bar")
//
// Here the second argument expresses to use the given default value. Thus
// DefaultArg acts as placeholder to be able to define further arguments. See
// documentation for spec.CLGIndex for more information.
//
type DefaultArg struct{}

// Arg

// ArgToArg converts the argument under index to interface{}, if possible.
// Note that this is only not possible if the argument under index is nil, or
// if there is no argument under the given index at all.
func ArgToArg(args []interface{}, index int) (interface{}, error) {
	if len(args) < index+1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if a, ok := args[index].(interface{}); ok {
		return a, nil
	}

	return nil, maskAnyf(wrongArgumentTypeError, "expected interface{} got %T", args[index])
}

// ArgToArgs converts the argument under index to []interface{}, if possible.
func ArgToArgs(args []interface{}, index int) ([]interface{}, error) {
	if len(args) < index+1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if as, ok := args[index].([]interface{}); ok {
		return as, nil
	}

	return nil, maskAnyf(wrongArgumentTypeError, "expected []interface{} got %T", args[index])
}

// ArgToArgsList converts the argument under index to [][]interface{}, if
// possible.
func ArgToArgsList(args []interface{}, index int) ([][]interface{}, error) {
	if len(args) < index+1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if asl, ok := args[index].([][]interface{}); ok {
		return asl, nil
	}

	return nil, maskAnyf(wrongArgumentTypeError, "expected [][]interface{} got %T", args[index])
}

// ArgToBool converts the argument under index to bool, if possible.
func ArgToBool(args []interface{}, index int) (bool, error) {
	if len(args) < index+1 {
		return false, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if b, ok := args[index].(bool); ok {
		return b, nil
	}

	return false, maskAnyf(wrongArgumentTypeError, "expected bool got %T", args[index])
}

// ArgToDistribution converts the argument under index to spec.Distribution, if
// possible.
func ArgToDistribution(args []interface{}, index int) (spec.Distribution, error) {
	if len(args) < index+1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if d, ok := args[index].(spec.Distribution); ok {
		return d, nil
	}

	return nil, maskAnyf(wrongArgumentTypeError, "expected spec.Distribution got %T", args[index])
}

// ArgToFeature converts the argument under index to spec.Feature, if
// possible.
func ArgToFeature(args []interface{}, index int) (spec.Feature, error) {
	if len(args) < index+1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if f, ok := args[index].(spec.Feature); ok {
		return f, nil
	}

	return nil, maskAnyf(wrongArgumentTypeError, "expected spec.Feature got %T", args[index])
}

// ArgToFeatures converts the argument under index to []spec.Feature, if
// possible.
func ArgToFeatures(args []interface{}, index int) ([]spec.Feature, error) {
	if len(args) < index+1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if fs, ok := args[index].([]spec.Feature); ok {
		return fs, nil
	}

	return nil, maskAnyf(wrongArgumentTypeError, "expected []spec.Feature got %T", args[index])
}

// ArgToFeatureSet converts the argument under index to spec.FeatureSet, if
// possible.
func ArgToFeatureSet(args []interface{}, index int) (spec.FeatureSet, error) {
	if len(args) < index+1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if fs, ok := args[index].(spec.FeatureSet); ok {
		return fs, nil
	}

	return nil, maskAnyf(wrongArgumentTypeError, "expected spec.FeatureSet got %T", args[index])
}

// ArgToFloat64 converts the argument under index to float64, if possible.
func ArgToFloat64(args []interface{}, index int) (float64, error) {
	if len(args) < index+1 {
		return 0, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if f, ok := args[index].(float64); ok {
		return f, nil
	}

	return 0, maskAnyf(wrongArgumentTypeError, "expected float64 got %T", args[index])
}

// ArgToFloat64Slice converts the argument under index to []float64, if
// possible.
func ArgToFloat64Slice(args []interface{}, index int, def ...[]float64) ([]float64, error) {
	// In any case we should make sure the defaults are validated. This causes a
	// more strict usage and understanding of the argument helper APIs.
	if len(def) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 default got %d", len(def))
	}

	if len(args) < index+1 {
		if len(def) == 1 {
			// There is no argument given, thus we use the default.
			return def[0], nil
		}
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if _, ok := args[index].(DefaultArg); ok {
		if len(def) < 1 {
			return nil, maskAnyf(notEnoughArgumentsError, "expected 1 default got 0")
		}
		// There is no argument given, thus we use the default.
		return def[0], nil
	}

	if fs, ok := args[index].([]float64); ok {
		return fs, nil
	}
	return nil, maskAnyf(wrongArgumentTypeError, "expected []float64 got %T", args[index])
}

// ArgToFloat64SliceSlice converts the argument under index to [][]float64, if
// possible.
func ArgToFloat64SliceSlice(args []interface{}, index int, def ...[][]float64) ([][]float64, error) {
	// In any case we should make sure the defaults are validated. This causes a
	// more strict usage and understanding of the argument helper APIs.
	if len(def) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 default got %d", len(def))
	}

	if len(args) < index+1 {
		if len(def) == 1 {
			// There is no argument given, thus we use the default.
			return def[0], nil
		}
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if _, ok := args[index].(DefaultArg); ok {
		if len(def) < 1 {
			return nil, maskAnyf(notEnoughArgumentsError, "expected 1 default got 0")
		}
		// There is no argument given, thus we use the default.
		return def[0], nil
	}

	if fss, ok := args[index].([][]float64); ok {
		return fss, nil
	}
	return nil, maskAnyf(wrongArgumentTypeError, "expected [][]float64 got %T", args[index])
}

// ArgToInt converts the argument under index to int, if possible. Optionally
// it takes one default value that is returned in case there is no argument
// available for the given index.
func ArgToInt(args []interface{}, index int, def ...int) (int, error) {
	// In any case we should make sure the defaults are validated. This causes a
	// more strict usage and understanding of the argument helper APIs.
	if len(def) > 1 {
		return 0, maskAnyf(tooManyArgumentsError, "expected 1 default got %d", len(def))
	}

	if len(args) < index+1 {
		if len(def) == 1 {
			// There is no argument given, thus we use the default.
			return def[0], nil
		}
		return 0, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if _, ok := args[index].(DefaultArg); ok {
		if len(def) < 1 {
			return 0, maskAnyf(notEnoughArgumentsError, "expected 1 default got 0")
		}
		// There is no argument given, thus we use the default.
		return def[0], nil
	}

	if i, ok := args[index].(int); ok {
		return i, nil
	}
	return 0, maskAnyf(wrongArgumentTypeError, "expected int got %T", args[index])
}

// ArgToIntSlice converts the argument under index to []int, if possible.
func ArgToIntSlice(args []interface{}, index int) ([]int, error) {
	if len(args) < index+1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if i, ok := args[index].([]int); ok {
		return i, nil
	}

	return nil, maskAnyf(wrongArgumentTypeError, "expected []int got %T", args[index])
}

// ArgToString converts the argument under index to string, if possible.
// Optionally it takes one default value that is returned in case there is no
// argument available for the given index.
func ArgToString(args []interface{}, index int, def ...string) (string, error) {
	// In any case we should make sure the defaults are validated. This causes a
	// more strict usage and understanding of the argument helper APIs.
	if len(def) > 1 {
		return "", maskAnyf(tooManyArgumentsError, "expected 1 default got %d", len(def))
	}

	if len(args) < index+1 {
		if len(def) == 1 {
			// There is no argument given, thus we use the default.
			return def[0], nil
		}
		return "", maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if _, ok := args[index].(DefaultArg); ok {
		if len(def) < 1 {
			return "", maskAnyf(notEnoughArgumentsError, "expected 1 default got 0")
		}
		// There is no argument given, thus we use the default.
		return def[0], nil
	}

	if s, ok := args[index].(string); ok {
		return s, nil
	}
	return "", maskAnyf(wrongArgumentTypeError, "expected string got %T", args[index])
}

// ArgToStringSlice converts the argument under index to []string, if possible.
// Optionally it takes one default value that is returned in case there is no
// argument available for the given index.
func ArgToStringSlice(args []interface{}, index int, def ...[]string) ([]string, error) {
	// In any case we should make sure the defaults are validated. This causes a
	// more strict usage and understanding of the argument helper APIs.
	if len(def) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 default got %d", len(def))
	}

	if len(args) < index+1 {
		if len(def) == 1 {
			// There is no argument given, thus we use the default.
			return def[0], nil
		}
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d args(s) got %d", index+1, len(args))
	}

	if _, ok := args[index].(DefaultArg); ok {
		if len(def) < 1 {
			return nil, maskAnyf(notEnoughArgumentsError, "expected 1 default got 0")
		}
		// There is no argument given, thus we use the default.
		return def[0], nil
	}

	if ss, ok := args[index].([]string); ok {
		return ss, nil
	}
	return nil, maskAnyf(wrongArgumentTypeError, "expected []string got %T", args[index])
}

// Args

// ArgsToValues converts the given arguments to reflect values.
func ArgsToValues(args []interface{}) []reflect.Value {
	values := make([]reflect.Value, len(args))

	for i := range args {
		values[i] = reflect.ValueOf(args[i])
	}

	return values
}

// ValuesToArgs converts the given []reflect.Value to []interface{}.
func ValuesToArgs(values []reflect.Value) ([]interface{}, error) {
	if len(values) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(values))
	}
	if len(values) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected 2 got %d", len(values))
	}

	if !values[1].IsValid() || values[1].IsNil() {
		return values[0].Interface().([]interface{}), nil
	}

	return nil, maskAny(values[1].Interface().(error))
}
