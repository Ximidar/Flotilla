/*
* @Author: Ximidar
* @Date:   2018-10-21 17:54:57
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-11-10 14:30:48
 */

package FileManagerTest

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/ximidar/Flotilla/Flotilla_File_Manager/FileManager"
	"github.com/ximidar/Flotilla/Flotilla_File_Manager/Files"
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

func Test_AddFile(t *testing.T) {
	// make a filesystem
	fm, err := FileManager.NewFileManager()
	check_err(t, "Could not create File manager. Test_AddFile", err)

	// First verify that the file we want is not there
	DeleteAllFiles(fm)

	// Get file path
	_, filename, _, _ := runtime.Caller(0)
	benchyOrgFile, _ := filepath.Abs(filepath.Clean(filepath.Dir(filename) + "/../Resources/3D_Benchy.gcode"))
	fmt.Println("Benchy Path is: ", benchyOrgFile)
	destPath := filepath.Clean(fm.RootFolderPath + "/3D_Benchy.gcode")

	// Add a file by full path
	err = fm.AddFile(benchyOrgFile, destPath)
	check_err(t, "Add File Failed Test_AddFile", err)
	if ok := FileExistsInStructure(fm, "/3D_Benchy.gcode"); !ok {
		t.Fatal("Test_AddFile Failed. Could not add full path")
	}

	// Add a file by relative path
	destPath = "/3D_Relative_Benchy.gcode"
	err = fm.AddFile(benchyOrgFile, destPath)
	check_err(t, "Add File Failed on Relative add Test_AddFile", err)
	if ok := FileExistsInStructure(fm, "/3D_Relative_Benchy.gcode"); !ok {
		t.Fatal("Test_AddFile Failed. Could not add relative path")
	}

	// Add by Relative path with folders
	destPath = "/test/3D_Relative_Benchy.gcode"
	err = fm.AddFile(benchyOrgFile, destPath)
	check_err(t, "Add File Failed on Relative add Test_AddFile", err)
	if ok := FileExistsInStructure(fm, "/3D_Relative_Benchy.gcode"); !ok {
		t.Fatal("Test_AddFile Failed. Could not add relative path")
	}

	// Add by relative with multiple folders
	destPath = "/test/all/the/folders/3D_Relative_Benchy.gcode"
	err = fm.AddFile(benchyOrgFile, destPath)
	check_err(t, "Add File Failed on Relative add Test_AddFile", err)
	if ok := FileExistsInStructure(fm, "/3D_Relative_Benchy.gcode"); !ok {
		t.Fatal("Test_AddFile Failed. Could not add relative path")
	}

}

//FileExistsInStructure will query the structure for existing files
func FileExistsInStructure(fm *FileManager.FileManager, file string) bool {

	_, err := fm.GetFileByPath(file)
	if err != nil {
		fm.PrintStructure()
		fmt.Println(err)
		return false
	}
	return true

}

func DeleteAllFiles(fm *FileManager.FileManager) {
	files := fm.Structure["root"].Contents

	for _, value := range files {
		fmt.Println("Deleting", value.Path)
		os.Remove(value.Path)
	}
}

func GetBenchy(fm *FileManager.FileManager) (*Files.File, error) {
	_, filename, _, _ := runtime.Caller(0)
	benchyOrgFile, _ := filepath.Abs(filepath.Clean(filepath.Dir(filename) + "/../Resources/3D_Benchy.gcode"))
	fmt.Println("Benchy Path is: ", benchyOrgFile)
	destPath := filepath.Clean(fm.RootFolderPath + "/3D_Benchy.gcode")

	if _, err := os.Stat(benchyOrgFile); !os.IsNotExist(err) {
		err := fm.AddFile(benchyOrgFile, destPath)
		if err != nil {
			return nil, err
		}
		file, err := fm.GetFileByPath("3D_Benchy.gcode")
		if err != nil {
			structure, _ := fm.GetJSONStructure()
			fmt.Println(string(structure))
			return nil, err
		}
		return file, nil
	}

	return nil, fmt.Errorf("Could not get file %v", ".3D_Benchy.gcode")

}

func check_err(t *testing.T, mess string, err error) {
	if err != nil {
		t.Fatalf("Failed Check from %v, Error: %v", mess, err)
	}
}
