/*
* @Author: Ximidar
* @Date:   2018-12-18 11:44:58
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-12-18 11:46:18
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
