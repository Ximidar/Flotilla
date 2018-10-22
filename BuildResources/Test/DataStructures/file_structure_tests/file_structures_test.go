/*
* @Author: Ximidar
* @Date:   2018-10-14 10:56:57
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-10-17 14:13:22
 */

package file_structure_test

import (
	"fmt"
	"testing"

	FS "github.com/ximidar/Flotilla/DataStructures/FileStructures"
)

func TestSetup(t *testing.T) {
	fmt.Println("Testing Setup!")
	test_path := "somepath"
	fa, err := FS.NewFileAction(FS.SelectFile, test_path)

	if err != nil {
		t.Fatal(err)
	}

	if fa.Action != FS.SelectFile {
		t.Fatal("Action was not set correctly")
	}

	if fa.Path != test_path {
		t.Fatal("Path was not set correctly")
	}
}
