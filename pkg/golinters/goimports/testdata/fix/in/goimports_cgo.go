//golangcitest:config_path testdata/goimports.yml
//golangcitest:expected_exitcode 0
package p

/*
 #include <stdio.h>
 #include <stdlib.h>

 void myprint(char* s) {
 	printf("%d\n", s);
 }
*/
import "C"

import (
    "os"
    "fmt"
)

 func goimports(a, b int) int {
 	if a != b {
 		return 1 
	}
 	return 2
}
