/*
* @Author: Ximidar
* @Date:   2018-12-18 11:44:58
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-23 13:45:13
 */

package CommonTestTools

import "testing"

// CheckErr will take in an error and a message to display for any failed
// errors
func CheckErr(t *testing.T, mess string, err error) {
	if err != nil {
		t.Fatalf("Failed Check from %v, Error: %v", mess, err)
	}
}

// CheckEquals will check if two object equal eachother
func CheckEquals(t *testing.T, object1 interface{}, object2 interface{}) {
	if object1 != object2 {
		t.Fatal("Objects do not match OB1:", object1, "OB2:", object2)
	}
}
