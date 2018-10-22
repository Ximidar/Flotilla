/*
* @Author: Ximidar
* @Date:   2018-10-21 17:54:57
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-10-21 17:57:21
 */

package FileManagerTest

import (
	"fmt"
	"testing"

	"github.com/ximidar/Flotilla/Flotilla_File_Manager/FileManager"
)

func Test_Setup(t *testing.T) {
	fmt.Println("Testing File Manager Setup")

	FileManager.NewFileManager()

}
