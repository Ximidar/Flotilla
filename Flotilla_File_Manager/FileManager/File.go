/*
* @Author: Ximidar
* @Date:   2018-10-02 16:48:31
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-27 15:15:30
 */

package FileManager

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	ospath "path"

	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
)

// NewFile is a constructor for FS.File
func NewFile(path string, filetype string, previousPath string) *FS.File {
	file := new(FS.File)
	file.Path = path
	file.FileType = filetype
	file.PreviousPath = previousPath
	return file
}

// UpdateInfo will be called to update the meta data for the file
func UpdateInfo(CurrentFile *FS.File) {

	// sometimes file gets dereferenced before we can update it. Recover from calling a null pointer
	defer func() {
		if recover() != nil {
			fmt.Println("UpdateInfo Failed")
			return
		}
	}()

	stats, err := os.Stat(CurrentFile.Path)
	if err != nil {
		fmt.Println(err)
		return
	}

	populateFileInfoForFile(CurrentFile, stats)
}

// Indexfs will index all subdirectories if the file is a folder
func Indexfs(CurrentFile *FS.File) {

	//clear contents
	CurrentFile.Contents = make([]*FS.File, 0)

	// read dir
	files, err := ioutil.ReadDir(CurrentFile.Path)
	if err != nil {
		fmt.Printf("Error Accessing path: %v\nErr: %v", CurrentFile.Path, err)
		return
	}

	for _, fileinfo := range files {
		if fileinfo.IsDir() {
			filePath := ospath.Clean(CurrentFile.Path + "/" + fileinfo.Name())
			dir := NewFile(filePath, "folder", CurrentFile.Path)
			populateFileInfoForFile(dir, fileinfo)
			CurrentFile.Contents = append(CurrentFile.Contents, dir)
			go Indexfs(dir)
		} else {
			filePath := ospath.Clean(CurrentFile.Path + "/" + fileinfo.Name())
			pfile := NewFile(filePath, "file", CurrentFile.Path)
			populateFileInfoForFile(pfile, fileinfo)
			CurrentFile.Contents = append(CurrentFile.Contents, pfile)
		}

	}

}

func populateFileInfoForFile(CurrentFile *FS.File, info os.FileInfo) {
	CurrentFile.Name = info.Name()
	CurrentFile.Size = uint64(info.Size())
	CurrentFile.IsDir = info.IsDir()
	CurrentFile.UnixTime = info.ModTime().Unix()
}

// PullFile will check an array for a file by name
func PullFile(contents []*FS.File, name string) (*FS.File, error) {

	for _, file := range contents {
		if file.GetName() == name {
			return file, nil
		}
	}
	return nil, errors.New("file doesn't exist")
}

// DeleteFiles will take out files from an array
func DeleteFiles(contents []*FS.File, slicePos ...int) []*FS.File {
	for index, val := range slicePos {
		correctedPosition := val - index
		contents = append(contents[:correctedPosition], contents[correctedPosition+1:]...)
	}
	return contents
}
