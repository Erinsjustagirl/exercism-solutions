// twofer does the thing it asks for.
package twofer

import "fmt"

// ShareWith is the program.
func ShareWith(name string) string {
    result := "One for you, one for me.";
	if name != "" {
	    result = fmt.Sprintf("One for %s, one for me.", name);
	}
	return result;
}
