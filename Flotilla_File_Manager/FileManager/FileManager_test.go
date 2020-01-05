/*
* @Author: Ximidar
* @Date:   2018-10-21 17:54:57
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-04-02 16:23:04
 */

package FileManager

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/Ximidar/Flotilla/BuildResources/Test/CommonTestTools"
	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
)

var TestLocation = "/tmp/testing/FileManager"

func Test_Setup(t *testing.T) {
	fmt.Println("Testing File Manager Setup")

	_, err := NewFileManager(TestLocation)

	if err != nil {
		t.Fatal(err)
	}

}

func Test_Structure(t *testing.T) {
	fm, err := NewFileManager(TestLocation)

	if err != nil {
		t.Fatal(err)
	}

	structure, err := fm.GetJSONStructure()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(structure))
}

func Test_GetFileByPath(t *testing.T) {
	fmt.Println("Testing Test_GetFileByPath")
	// Remove file
	os.RemoveAll(TestLocation)

	// Create File Manager
	fm, err := NewFileManager(TestLocation)
	CommonTestTools.CheckErr(t, "Could not constuct File Manager", err)

	// get the bench path and create the destination path
	_, filename, _, _ := runtime.Caller(0)
	benchyOrgFile, _ := filepath.Abs(filepath.Clean(filepath.Dir(filename) + "/../Resources/3D_Benchy.gcode"))
	fmt.Println("Benchy Path is: ", benchyOrgFile)
	destPath := filepath.Clean(TestLocation + "/3D_Benchy.gcode")

	// copy source to destination
	err = simpleCopy(benchyOrgFile, destPath)
	CommonTestTools.CheckErr(t, "Could not copy benchy to Test Location", err)
	// Give a little bit of time to update the file structure
	fm.PrintStructure()

	_, err = fm.GetFileByPath(destPath)
	CommonTestTools.CheckErr(t, "Could not get file by path", err)

}

func Test_AddFile(t *testing.T) {
	tl, err := os.Open(TestLocation)
	CommonTestTools.CheckErr(t, "Could not open location", err)
	os.RemoveAll(TestLocation)
	tl.Sync()
	tl.Close()

	// make a filesystem
	fm, err := NewFileManager(TestLocation)
	CommonTestTools.CheckErr(t, "Could not create File manager. Test_AddFile", err)

	// Get file path
	_, filename, _, _ := runtime.Caller(0)
	benchyOrgFile, _ := filepath.Abs(filepath.Clean(filepath.Dir(filename) + "/../Resources/3D_Benchy.gcode"))
	fmt.Println("Benchy Path is: ", benchyOrgFile)
	destPath := filepath.Clean(fm.RootFolderPath + "/3D_Benchy.gcode")

	// Add a file by full path
	err = fm.AddFile(benchyOrgFile, destPath)
	CommonTestTools.CheckErr(t, "Add File Failed Test_AddFile", err)
	if ok := FileExistsInStructure(fm, "/3D_Benchy.gcode"); !ok {
		t.Fatal("Test_AddFile Failed. Could not add full path")
	}

	// Add a file by relative path
	destPath = "/3D_Relative_Benchy.gcode"
	err = fm.AddFile(benchyOrgFile, destPath)
	CommonTestTools.CheckErr(t, "Add File Failed on Relative add Test_AddFile", err)
	if ok := FileExistsInStructure(fm, "/3D_Relative_Benchy.gcode"); !ok {
		t.Fatal("Test_AddFile Failed. Could not add relative path")
	}

	// Add by Relative path with folders
	destPath = "/test/3D_Relative_Benchy.gcode"
	err = fm.AddFile(benchyOrgFile, destPath)
	CommonTestTools.CheckErr(t, "Add File Failed on Relative add Test_AddFile", err)
	if ok := FileExistsInStructure(fm, "/3D_Relative_Benchy.gcode"); !ok {
		t.Fatal("Test_AddFile Failed. Could not add relative path")
	}

	// Add by relative with multiple folders
	destPath = "/test/all/the/folders/3D_Relative_Benchy.gcode"
	err = fm.AddFile(benchyOrgFile, destPath)
	CommonTestTools.CheckErr(t, "Add File Failed on Relative add Test_AddFile", err)
	if ok := FileExistsInStructure(fm, "/3D_Relative_Benchy.gcode"); !ok {
		t.Fatal("Test_AddFile Failed. Could not add relative path")
	}

}

func Test_DeleteFile(t *testing.T) {
	// make a filesystem
	fm, err := NewFileManager(TestLocation)
	CommonTestTools.CheckErr(t, "Could not create File manager. Test_AddFile", err)

	// First verify that the file we want is not there
	DeleteAllFiles(fm)

	// add file
	_, filename, _, _ := runtime.Caller(0)
	benchyOrgFile, _ := filepath.Abs(filepath.Clean(filepath.Dir(filename) + "/../Resources/3D_Benchy.gcode"))
	fmt.Println("Benchy Path is: ", benchyOrgFile)
	destPath := filepath.Clean(fm.RootFolderPath + "/3D_Benchy.gcode")
	fm.AddFile(benchyOrgFile, destPath)

	// delete file by full path
	err = fm.DeleteFile(destPath)
	CommonTestTools.CheckErr(t, "Could not delete file by full path Test_DeleteFile", err)
	<-time.After(200 * time.Millisecond)
	if ok := FileExistsInStructure(fm, "/3D_Benchy.gcode"); ok {
		t.Fatal("Test_DeleteFile Failed. Did not actually delete file")
	}

	// add file
	fm.AddFile(benchyOrgFile, destPath)

	// delete file by relative path
	err = fm.DeleteFile("/3D_Benchy.gcode")
	CommonTestTools.CheckErr(t, "Could not delete file by relative path Test_DeleteFile", err)
	<-time.After(200 * time.Millisecond)
	if ok := FileExistsInStructure(fm, "/3D_Benchy.gcode"); ok {
		t.Fatal("Test_DeleteFile Failed. Did not actually delete file")
	}

	// try to delete non-existant file by relative and full path
	err = fm.DeleteFile(fm.RootFolderPath + "/non/existant/file.gcode")
	if err == nil {
		CommonTestTools.CheckErr(t, "Could not delete fake file by full path Test_DeleteFile", errors.New("There was not an error. this is not expected behaviour"))
	}

	err = fm.DeleteFile("/non/existant/file.gcode")
	if err == nil {
		CommonTestTools.CheckErr(t, "Could not delete fake file by relative path Test_DeleteFile", errors.New("There was not an error. this is not expected behaviour"))
	}
}

func Test_MoveFile(t *testing.T) {
	// make a filesystem
	fm, err := NewFileManager(TestLocation)
	CommonTestTools.CheckErr(t, "Could not create File manager. Test_AddFile", err)

	// First verify that the file we want is not there
	DeleteAllFiles(fm)

	// add file
	_, filename, _, _ := runtime.Caller(0)
	benchyOrgFile, _ := filepath.Abs(filepath.Clean(filepath.Dir(filename) + "/../Resources/3D_Benchy.gcode"))
	fmt.Println("Benchy Path is: ", benchyOrgFile)
	destPath := filepath.Clean(fm.RootFolderPath + "/3D_Benchy.gcode")
	fm.AddFile(benchyOrgFile, destPath)

	// add folder TODO update this part to use fm function instead of just making a directory
	os.Mkdir(fm.RootFolderPath+"/testfolder/", 0750)

	// move file to folder
	err = fm.MoveFile(destPath, fm.RootFolderPath+"/testfolder/3D_Benchy.gcode")
	CommonTestTools.CheckErr(t, "could not move file to folder Test_MoveFile", err)

	// try to move file outside controlled space
	fm.MoveFile(fm.RootFolderPath+"/testfolder/3D_Benchy.gcode", destPath)
	err = fm.MoveFile(destPath, "/not/a/real/folder")
	if err == nil {
		CommonTestTools.CheckErr(t, "Test_MoveFile failed to produce error", errors.New("There was not an error. this is not expected behaviour"))
	}

	// try to move non existant file
	err = fm.MoveFile("not/a/real/file.gcode", "/not/a/real/dest.gcode")
	if err == nil {
		CommonTestTools.CheckErr(t, "Test_MoveFile failed to produce error", errors.New("There was not an error. this is not expected behaviour"))
	}
}

//FileExistsInStructure will query the structure for existing files
func FileExistsInStructure(fm *FileManager, file string) bool {
	_, err := fm.GetFileByPath(file)
	if err != nil {
		fm.PrintStructure()
		fmt.Println(err)
		return false
	}
	return true

}

func DeleteAllFiles(fm *FileManager) {
	files := fm.Structure.GetContents()

	for _, value := range files {
		fmt.Println("Deleting", value.Path)
		if value.IsDir {
			os.RemoveAll(value.Path)
		} else {
			os.Remove(value.Path)
		}

	}
}

func GetBenchy(fm *FileManager) (*FS.File, error) {
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

func simpleCopy(src, dest string) error {

	fmt.Printf("Copying %v to %v\n", src, dest)
	buf := make([]byte, 5*1000)
	// #nosec
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()
	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	destination.Sync()

	// #nosec
	// Change Mode of destination file so it is executable to a normal user
	os.Chmod(dest, 0755)
	return nil
}
