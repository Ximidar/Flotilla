/*
* @Author: Ximidar
* @Date:   2018-10-10 02:38:49
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-27 15:53:28
 */

package FileManager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	ospath "path"
	"path/filepath"
	"strings"
	"sync"

	FS "github.com/Ximidar/Flotilla/DataStructures/FileStructures"
	"github.com/fsnotify/fsnotify"
)

// ErrFileDoesNotExist is an error for the file system
var ErrFileDoesNotExist = errors.New("file doesn't exist")

// ErrFileDoesNotExistInStructure is an error for the file system
var ErrFileDoesNotExistInStructure = errors.New("file doesn't exist in structure")

// ErrFileAlreadyExistsInStructure is an error for the file system
var ErrFileAlreadyExistsInStructure = errors.New("file already exists in structure")

// FileManager is an object for keeping track of a folder and
// it's contents.
type FileManager struct {
	RootFolderPath string
	Structure      *FS.File
	Watcher        *fsnotify.Watcher

	lastOp fsnotify.Op
	mux    sync.Mutex
}

// NewFileManager will contstruct a FileManager Object
func NewFileManager(rootPath string) (*FileManager, error) {
	fm := new(FileManager)

	fm.RootFolderPath = rootPath
	fm.mux.Lock()
	fm.Structure = NewFile(fm.RootFolderPath, "folder", "")
	fm.mux.Unlock()

	// Create notify watcher
	var err error
	fm.Watcher, err = fsnotify.NewWatcher()

	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil, err
	}

	fm.InitPaths()
	return fm, nil
}

// FSWatcher will watch for events and handle them as they come in
func (fm *FileManager) FSWatcher() {
	for {
		select {
		// watch for events
		case event := <-fm.Watcher.Events:
			fm.HandleEvent(event)

		// watch for errors
		case err := <-fm.Watcher.Errors:
			fmt.Println("ERROR", err)
			break
		}
	}
}

// PrintEventDetails will take an event and print the event details about it. If the last op is the same it will just print a dot
func (fm *FileManager) PrintEventDetails(event fsnotify.Event) {
	if event.Op != fm.lastOp {
		fmt.Printf("%v Event at %s\n", event.Op, event.Name)
	}

	fm.lastOp = event.Op

}

// HandleWrite will handle all write events
func (fm *FileManager) HandleWrite(event fsnotify.Event) {
	fmt.Println("Write Happened on", event.Name)
	file, err := fm.GetFileByPath(event.Name)
	if err != nil {
		// File is not in the structure yet
		return
	}
	if _, err := os.Stat(event.Name); !os.IsNotExist(err) {
		UpdateInfo(file)
	}
	fm.PrintStructure()
}

// HandleDelete will handle all Rename and Delete events
func (fm *FileManager) HandleDelete(event fsnotify.Event) {
	fmt.Println("Delete Happened on", event.Name)
	err := fm.RemoveFileFromStructure(event.Name)
	if err != nil {
		if err != ErrFileDoesNotExist {
			fm.PrintEventDetails(event)
			fmt.Println("Could not remove file from structure", event.Name)
			fmt.Println(err)
		}

		return
	}
	fm.PrintStructure()
}

// HandleCreate will handle all creation events
func (fm *FileManager) HandleCreate(event fsnotify.Event) {
	fmt.Println("Create Happened on", event.Name)
	err := fm.AddFileToStructure(event.Name)
	if err != nil {
		if err != ErrFileAlreadyExistsInStructure {
			fm.PrintEventDetails(event)
			message := fmt.Sprintf("Could not add File to structure: %v\nError: %v\n", event.Name, err)
			fmt.Println(message)
		}

		return
	}
	fm.PrintStructure()

}

// HandleEvent will handle a File system Event
func (fm *FileManager) HandleEvent(event fsnotify.Event) {
	op := event.Op

	switch op {

	case fsnotify.Create:
		go fm.HandleCreate(event)
	case fsnotify.Remove:
		go fm.HandleDelete(event)
	case fsnotify.Rename:
		go fm.HandleDelete(event)
	case fsnotify.Write:
		go fm.HandleWrite(event)

	}
}

// RemoveFileFromStructure Will remove a file object from the Structure Map
func (fm *FileManager) RemoveFileFromStructure(path string) error {
	file, err := fm.GetFileByPath(ospath.Dir(path))
	if err != nil {
		return err
	}
	basename := ospath.Base(path)
	var indexForDeletion = -1
	for index, pulledFile := range file.GetContents() {
		if pulledFile.GetName() == basename {
			indexForDeletion = index
			break
		}
	}

	if indexForDeletion == -1 {
		return errors.New("could not find file to delete")
	}

	file.Contents = DeleteFiles(file.GetContents(), indexForDeletion)

	return nil
}

// AddFileToStructure will add a file to the structure that was not previously there
func (fm *FileManager) AddFileToStructure(path string) error {

	if path == "" {
		return errors.New("path is empty. We will not add it")
	}
	_, err := fm.GetFileByPath(path)
	if err == nil {
		return ErrFileAlreadyExistsInStructure
	}

	file, err := fm.GetFileByPath(ospath.Dir(path))
	if err != nil {
		fmt.Printf("Could not get %v\n", ospath.Dir(path))
		return err
	}
	// tell the directory to re-index all it's files
	go Indexfs(file)

	// figure out if we need to add another watcher
	stats, fileerr := os.Stat(path)
	if fileerr != nil {
		fmt.Println("Could not stat path", path)
		return fileerr
	}
	if stats.IsDir() {
		fmt.Println("Adding path to fsnotify watcher", path)
		fm.Watcher.Add(path)
	}

	return nil

}

// GetFileByPath This function will be given a path and it will return the correct file
func (fm *FileManager) GetFileByPath(path string) (file *FS.File, err error) {

	defer func() {
		recoverErr := recover()
		if recoverErr != nil {
			file = nil
			errMess := fmt.Sprintf("recovered, File doesn't exist %v", recoverErr)
			err = errors.New(errMess)
		}
	}()

	// if it is not in the root path then it is a relative path
	if inPath := fm.isInRootPath(path); !inPath {
		path = fm.makeRelativePath(path)
	}

	// if the path is the root file, then get the root file
	if path == fm.RootFolderPath {
		fm.mux.Lock()
		defer fm.mux.Unlock()
		return fm.Structure, nil
	}
	pathCopy := path
	// cut off everything before the root file
	pathCopy = strings.Replace(pathCopy, fm.RootFolderPath+"/", "", 1)

	// split remaining names up by "/" char
	pathKeys := strings.Split(pathCopy, "/")

	// pull file out of structure

	pulledFile := new(FS.File)
	fm.mux.Lock()
	structure := fm.Structure.GetContents()
	fm.mux.Unlock()

	for _, key := range pathKeys {
		pulledFile, err = PullFile(structure, key)
		if err != nil {
			return nil, err
		}
		structure = pulledFile.GetContents()

	}

	if pulledFile.GetPath() == "" {
		return nil, ErrFileDoesNotExistInStructure
	}

	return pulledFile, nil
}

// InitPaths This function will check that our Root Folder Path Exists and is created, then it will add fsnotify
// watchers to every path
func (fm *FileManager) InitPaths() {

	// Create root folder path
	fm.createRootFolderPath()

	// Index FS
	fm.indexfs()

	go fm.FSWatcher()

}

// walk the root directory and index filesystem while also adding directories to the
// fs watcher
func (fm *FileManager) indexfs() {
	filepath.Walk(fm.RootFolderPath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			fm.Watcher.Add(path)
		}

		return nil

	})

	fm.mux.Lock()
	fm.Structure = NewFile(fm.RootFolderPath, "folder", "")
	Indexfs(fm.Structure)
	fm.mux.Unlock()

	fm.PrintStructure()
}

// GetStructure will return the structure safely
func (fm *FileManager) GetStructure() *FS.File {
	fm.mux.Lock()
	defer fm.mux.Unlock()
	copyStructure := fm.Structure
	return copyStructure
}

// PrintStructure will dump the Structure object to the console
func (fm *FileManager) PrintStructure() {
	marshed, err := fm.GetJSONStructure()
	if err != nil {
		fmt.Println("Couldn't get json structure:", err)
		return
	}
	fmt.Println(string(marshed))
}

// GetJSONStructure will create a JSON object of the Structure Object
func (fm *FileManager) GetJSONStructure() ([]byte, error) {
	fm.mux.Lock()
	marshed, err := json.MarshalIndent(fm.Structure, "", "||||")
	fm.mux.Unlock()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return marshed, nil
}

func (fm *FileManager) createRootFolderPath() {
	// Check and create basic data directory
	if _, err := os.Stat(fm.RootFolderPath); os.IsNotExist(err) {
		//#nosec
		err := os.MkdirAll(fm.RootFolderPath, 0755)
		if err != nil {
			panic(err)
		}
	}
}

////////// File Operations
/////////////////////////////////////////////////////////////////////////////
// reminder: You do not need to mess with Structure here at all.           //
// fsnotify functions will take care of maintaining the correct structure  //
/////////////////////////////////////////////////////////////////////////////

// AddFile will copy a file from src to dst
func (fm *FileManager) AddFile(src, dst string) error {

	if inPath := fm.isInRootPath(dst); !inPath {
		var err error
		dst, err = fm.solveRelativePath(dst)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Copying %v to %v\n", src, dst)
	err := fm.cp(src, dst)

	if err != nil {
		return err
	}

	return nil
}

// MoveFile will move a file from src to dst
func (fm *FileManager) MoveFile(src, dst string) (err error) {

	// detect relative paths for src and dst
	if inPath := fm.isInRootPath(src); !inPath {
		src = fm.makeRelativePath(src)
	}
	if inPath := fm.isInRootPath(dst); !inPath {
		dst = fm.makeRelativePath(dst)
	}

	// move the file
	err = fm.AddFile(src, dst)
	if err != nil {
		return err
	}
	err = fm.DeleteFile(src)
	return err
}

// DeleteFile will delete a file at path or relative path
func (fm *FileManager) DeleteFile(path string) error {

	// modify path if it is relative
	if inPath := fm.isInRootPath(path); !inPath {
		path = fm.makeRelativePath(path)
	}
	file, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !file.Mode().IsRegular() {
		return fmt.Errorf("non-regular file %s (%q)", file.Name(), file.Mode().String())
	}

	// delete the file
	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

func (fm *FileManager) isInRootPath(path string) (inPath bool) {
	defer func() {
		// recover from panic if one occurred.
		if recover() != nil {
			inPath = false
		}
	}()
	// Check one is to just compare paths
	inPath = false
	if path[:len(fm.RootFolderPath)] == fm.RootFolderPath {
		inPath = true
		return
	}

	// Check two is to grab the file info for the path and check it against the fileinfo for the root path
	f1, err := os.Stat(path[:len(fm.RootFolderPath)])
	// Don't check this error because we know RootFolderPath Exists and will stat just fine
	f2, _ := os.Stat(fm.RootFolderPath)

	if err != nil {
		inPath = false
		return
	}

	inPath = os.SameFile(f1, f2)
	return

}

// solveRelativePath will be given a path and it will rectify the path to put it in the root folder
// For example if someone puts the destination as "/com/boom/" it will grab the root path and then create the
// folders in order to satisfy the relative path
func (fm *FileManager) solveRelativePath(path string) (string, error) {
	rectifiedPath := filepath.Clean(fm.RootFolderPath + "/" + path)
	dirPath := filepath.Dir(rectifiedPath)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fileerr := os.MkdirAll(dirPath, 0750)
		if fileerr != nil {
			return "", fileerr
		}
	}
	return rectifiedPath, nil
}

// makeRelativePath will make a path relative without making folders
func (fm *FileManager) makeRelativePath(path string) string {
	rectifiedPath := filepath.Clean(fm.RootFolderPath + "/" + path)
	return rectifiedPath
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise copy the file contents from src to dst.
func (fm *FileManager) cp(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	err = fm.copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func (fm *FileManager) copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
