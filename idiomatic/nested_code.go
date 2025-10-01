// Package idiomatic provides examples of idiomatic Go code patterns.
//
// This package demonstrates how to write Go code following best practices
// and common idioms, with specific examples of refactoring non-idiomatic
// code into more readable and maintainable forms.
package idiomatic

import "errors"

// Join concatenates two strings and returns the result truncated to max length if necessary.
//
// If either s1 or s2 is empty, an error is returned.
// If the concatenated string exceeds max length, it is truncated to max characters.
// Otherwise, the full concatenated string is returned.
/*func Join(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)

			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}*/

// Join concatenates two strings and returns the result truncated to max length if necessary.
//
// If either s1 or s2 is empty, an error is returned.
// If the concatenated string exceeds max length, it is truncated to max characters.
// Otherwise, the full concatenated string is returned.
//
// This implementation avoids unnecessary nesting for better readability
func Join(s1, s2 string, max int) (string, error) {

	if s1 == "" {
		return "", errors.New("s1 is empty")
	}

	if s2 == "" {
		return "", errors.New("s2 is empty")
	}

	concat, err := concatenate(s1, s2)

	if err != nil {
		return "", err
	}

	if len(concat) > max {
		return concat[:max], nil
	}

	return concat, nil

}

// concatenate combines two strings and returns the result.
func concatenate(s1, s2 string) (string, error) {
	return s1 + s2, nil
}
