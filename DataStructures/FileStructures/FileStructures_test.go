/*
* @Author: Ximidar
* @Date:   2018-10-14 10:56:57
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-04-02 16:19:16
 */

package FileStructures

import (
	"fmt"
	"testing"
)

// TODO Take this out
func TestSetup(t *testing.T) {
	fmt.Println("Testing Setup!")
	test_path := "somepath"
	fa, err := NewFileAction(FileAction_SelectFile, test_path)

	if err != nil {
		t.Fatal(err)
	}

	if fa.Action != FileAction_SelectFile {
		t.Fatal("Action was not set correctly")
	}

	if fa.Path != test_path {
		t.Fatal("Path was not set correctly")
	}
}
