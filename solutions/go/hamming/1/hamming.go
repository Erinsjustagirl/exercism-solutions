package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
	    return 0, errors.New("darn");
	} else {
        distance := 0;
        for i := 0; i < len(a); i++ { // HEAVILY referenced some of this code from marcalperapoch's solution. - Erin
            if a[i] != b[i] {
                distance++
            }
        }
	    return distance, nil;
	}
}
