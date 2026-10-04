package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, errors.New("input must be a positive integer");
    }
    
	result := n;
    steps := 0
    
    for result != 1 {
		if result % 2 == 0 {
			result /= 2
		} else {
			result = result * 3 + 1
		}
		steps++
	}
    return steps, nil;
}
