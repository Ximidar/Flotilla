/*
* @Author: Ximidar
* @Date:   2018-10-14 10:56:57
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-27 17:19:06
 */

package file_structure_test

import (
	"fmt"
	"testing"

	FS "github.com/ximidar/Flotilla/DataStructures/FileStructures"
)

// TODO Take this out
func TestSetup(t *testing.T) {
	fmt.Println("Testing Setup!")
	test_path := "somepath"
	fa, err := FS.NewFileAction(FS.FileAction_SelectFile, test_path)

	if err != nil {
		t.Fatal(err)
	}

	if fa.Action != FS.FileAction_SelectFile {
		t.Fatal("Action was not set correctly")
	}

	if fa.Path != test_path {
		t.Fatal("Path was not set correctly")
	}
}
