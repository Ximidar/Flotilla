/*
* @Author: Ximidar
* @Date:   2018-10-21 17:54:57
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-10-22 18:38:24
 */

package FileManagerTest

import (
	"fmt"
	"testing"

	"github.com/ximidar/Flotilla/Flotilla_File_Manager/FileManager"
)

func Test_Setup(t *testing.T) {
	fmt.Println("Testing File Manager Setup")

	_, err := FileManager.NewFileManager()

	if err != nil {
		t.Fatal(err)
	}

}

func Test_Structure(t *testing.T) {
	fm, err := FileManager.NewFileManager()

	if err != nil {
		t.Fatal(err)
	}

	structure, err := fm.GetJSONStructure()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(structure))
}
