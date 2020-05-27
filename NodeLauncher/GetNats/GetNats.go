package GetNats

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/Ximidar/Flotilla/NodeLauncher/RootFolder"
)

type GithubData struct {
	URL     string  `json:"url"`
	HTMLURL string  `json:"html_url"`
	Name    string  `json:"name"`
	Assets  []Asset `json:"assets"`
}

type Asset struct {
	URL                string      `json:"url"`
	BrowserDownloadURL string      `json:"browser_download_url"`
	ID                 int         `json:"id"`
	NodeID             string      `json:"node_id"`
	Name               string      `json:"name"`
	Label              string      `json:"label"`
	State              string      `json:"state"`
	ContentType        string      `json:"content_type"`
	Size               int         `json:"size"`
	DownloadCount      int         `json:"download_count"`
	CreatedAt          string      `json:"created_at"`
	UpdatedAt          string      `json:"updated_at"`
	Uploader           interface{} `json:"uploader"`
}

// GetNats will go get the latest version of NATS
func GetNatsLinks() []Asset {
	url := "https://api.github.com/repos/nats-io/nats-server/releases/latest"
	response, _ := http.Get(url)
	fmt.Println(response.Status)
	data, _ := ioutil.ReadAll(response.Body)
	gd := new(GithubData)
	err := json.Unmarshal(data, gd)
	if err != nil {
		panic(err)
	}

	download := make([]Asset, 0)
	lookFor := []string{
		"linux-amd64.zip",
		"linux-arm64.zip",
		"linux-arm7.zip",
	}

	for _, asset := range gd.Assets {
		for _, end := range lookFor {
			if strings.HasSuffix(asset.Name, end) {
				fmt.Println("Asset Found ", asset.Name)
				download = append(download, asset)
			}
		}
	}
	return download

}

// DownloadNats will download the nats server to a package folder
func DownloadNats(rf *RootFolder.RootFolder) error {

	downloads := GetNatsLinks()

	// Download files
	lookFor := make(map[string]string)
	lookFor = map[string]string{
		"amd64": "linux-amd64.zip",
		"arm64": "linux-arm64.zip",
		"arm":   "linux-arm7.zip",
	}

	for key, val := range lookFor {
		fmt.Println(key, " ", val)
		downloadLoc, ok := rf.ArchPath[key]
		if !ok {
			fmt.Println("Cannot find key: ", key)
			fmt.Println(rf.ArchPath)
			panic("Cannot find key")
		}

		for _, asset := range downloads {
			if strings.HasSuffix(asset.Name, val) {
				fmt.Println("Downloading Asset: ", asset.Name)
				fmt.Println("Asset Link: ", asset.BrowserDownloadURL)
				fmt.Println("Download LOC: ", downloadLoc)

				// download
				response, err := http.Get(asset.BrowserDownloadURL)
				if err != nil {
					panic(err)
				}

				defer response.Body.Close()

				// create file
				filename := path.Base(asset.BrowserDownloadURL)
				fullPath := path.Clean(downloadLoc + "/" + filename)
				out, err := os.Create(fullPath)
				if err != nil {
					panic(err)
				}

				defer out.Close()

				_, err = io.Copy(out, response.Body)
				if err != nil {
					panic(err)
				}

			}
		}
	}

	return nil
}

func UnzipAndClean(rf *RootFolder.RootFolder) error {

	for key, val := range rf.ArchPath {
		fmt.Println("Unzipping file for arch: ", key)

		// find file ending in .zip
		files, err := ioutil.ReadDir(val)
		if err != nil {
			fmt.Println("UnzipAndClean Failed on ", err)
			return err
		}

		var zipfile os.FileInfo
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".zip") {
				zipfile = file
			}

		}

		if zipfile == nil {
			return errors.New("Zipfile not found in DIR: " + val)
		}

		// open ZipFile
		zipReader, err := zip.OpenReader(val + "/" + zipfile.Name())
		if err != nil {
			fmt.Println("Could not open Zipfile")
			return err
		}
		defer zipReader.Close()

		// find nats-server
		for _, file := range zipReader.File {
			fmt.Println("Found File: ", file.Name)
			if strings.HasSuffix(file.Name, "nats-server") {
				fmt.Println("Unzipping File!")
				zipped, err := file.Open()
				if err != nil {
					fmt.Println("Could not open file")
					return err
				}

				unzipName := val + "/" + path.Base(file.Name)
				unzipf, err := os.OpenFile(unzipName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
				if err != nil {
					fmt.Println("Could not create unzip file", unzipName)
					return err
				}

				defer unzipf.Close()

				_, err = io.Copy(unzipf, zipped)
				if err != nil {
					return err
				}
			}

		}

		// now delete after closing
		defer os.Remove(val + "/" + zipfile.Name())

	}

	return nil
}
