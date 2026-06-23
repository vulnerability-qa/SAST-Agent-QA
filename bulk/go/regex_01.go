// CWE-1333: ReDoS — Go's RE2 is safe but incorrect use of backtracking lib
package main
import "regexp"
func validate(input string) bool {
	// RE2 doesn't backtrack, but developers sometimes switch to regexp/syntax which can
	r := regexp.MustCompile(`^(\w+\s?)+$`)
	return r.MatchString(input)
}
