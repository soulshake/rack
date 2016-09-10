// Code generated by go-bindata.
// sources:
// templates/git-restore-mtime
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesGitRestoreMtime = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x54\x4d\x8f\xe3\x36\x0c\xbd\xeb\x57\xb0\x08\x8a\xd8\x40\x24\xef\x4c\x07\x6d\x37\x40\x50\x6c\xdb\x39\xf4\x50\x60\x51\xa0\xa7\xe9\x1c\x14\x9b\xb6\xb5\x91\x45\xaf\x48\x27\x93\x7f\x5f\xc8\x1f\x99\xa4\x41\x57\x17\x29\xe2\x23\x1f\xdf\xa3\xe2\xd5\x77\xc5\xc0\xb1\xd8\xbb\x50\x60\x38\x42\x7f\x96\x96\x82\x52\x2b\x68\x45\xfa\x6d\x51\xb0\xd8\xf2\x40\x47\x8c\xb5\xa7\x93\x29\xa9\x2b\xbe\x0e\xc8\xe2\x28\x70\xf1\xf0\xf1\xc7\xa7\xa7\x9f\x3e\x14\xa7\xd6\x0a\x6b\x69\x51\xe3\xd7\xc1\x1d\xad\xc7\x20\x9a\x6a\x3d\x30\xea\x92\xba\xce\x89\x16\xd7\x21\xeb\x9a\xa2\x6e\x9c\x14\x0f\x3f\x3c\xfe\xfc\xf4\xf8\xf8\x71\xb5\x1c\x12\xe1\xaf\x36\xa2\xde\x53\x40\x86\x23\x46\x76\x14\x0c\xfc\x36\xc4\x88\x41\xa0\x72\x11\xba\x81\x05\xf6\x08\x42\xbd\xf6\x78\x44\x0f\x54\xc3\x89\xe2\x01\x24\x22\x1a\xb5\x82\xbf\xd9\x36\xb8\x85\xc6\x89\x8e\xc8\x42\x11\x75\x97\x88\xf5\xde\x46\x84\x97\xde\x4a\xcb\x3d\x96\x6c\x8c\x79\x4d\x84\x67\xa8\xb0\xb6\x83\x17\x18\xfa\xca\x0a\x82\xf5\x1e\x6a\xe7\x91\xd5\x0a\x9e\xdf\x6c\xd7\x7b\xdc\x82\x10\x50\xf0\xe7\x05\x33\x9e\xa5\x45\xf8\xeb\xf9\xd3\xef\x7f\x3e\x83\x0d\xd5\x94\x03\x2e\x80\x29\x2a\x2a\xb7\x6a\xf5\x7f\x3d\xcc\x39\x15\x95\x4a\xb9\xae\xa7\x28\xc0\xc3\xbe\x8f\x54\x22\xf3\x06\xb8\xf5\xf8\x76\x09\x9c\x79\x03\xc4\x26\xb5\xad\x54\xa2\xf0\x8e\x05\x76\xc0\x28\x59\xae\x6a\x8a\x90\x42\x89\x36\xe3\x33\x1b\x1b\x9b\xe3\xcb\xc3\xf6\x15\x28\xc2\xcb\x9c\x67\xca\x21\x56\x2e\xbe\xe6\x5b\x05\x00\xe0\xea\xa5\xa0\x71\x9c\x2a\x66\xe9\x9c\xa7\x8c\xf7\x7b\xef\xc2\x61\xba\x9f\x92\xd2\x5a\xd8\x8d\xad\xaa\x6c\x81\x46\xf4\x69\x9f\xb0\xf9\x88\x45\x7f\x43\x51\xb9\x78\x57\x89\x22\x44\x22\xd9\x24\xe1\x95\x8b\xbc\x79\x77\x8f\xd8\x9c\xac\xbf\x23\x9f\x3b\x5f\x9b\xc6\xc9\x3a\xe1\xe6\xcc\x5b\x44\x5a\x73\xc0\x44\xec\xe8\x88\xd9\x94\x91\xdf\xc0\x12\x7f\x22\x4c\x75\x46\xe2\xfb\x2a\xdf\x14\xbb\xfc\xfe\x42\x2e\x64\x93\x90\x84\xcf\xf3\x5c\xa9\x71\xd0\xb0\x83\x0f\xaa\x71\x42\xfb\x2f\x69\x56\x97\xe9\x9a\xcf\xd4\x63\xc8\xc6\x11\x1b\xee\xbd\x93\x6c\xdd\x38\x81\xf4\xe7\x29\x5b\x1b\x1a\xac\x40\xeb\x3e\xa2\xc8\x79\xf7\xbd\x95\x75\xbe\xb9\xeb\xec\x4a\xa9\x54\x34\xc8\xee\xba\xfc\x1f\x9f\x9f\xa7\x57\xe1\x5d\x18\xe5\x4d\x4d\x98\x09\x3a\xc9\x1c\x43\xbb\x71\x33\x2c\xd1\xf5\x59\xbe\x3c\x8c\x40\x32\xde\x6f\xa1\xa4\x20\x2e\x0c\xa8\x96\xd0\x0c\xb7\x51\xf8\xe4\xa4\xcd\xd6\xdb\xf5\x7f\xde\xc6\xa5\xe6\xa4\xeb\x1f\x59\xe7\x2f\xfa\xe1\x55\x5d\x8d\xef\xda\xf4\xe4\xee\xad\xef\x17\xcf\xe7\xd1\x8d\x9e\xde\x20\x56\x7d\x74\x41\x60\xf4\x78\xf2\xfc\x26\x4c\x6c\x86\x14\x1a\x33\x37\x90\xcd\xb8\x71\xbb\x3c\x4e\xc6\x77\xd6\x65\x58\x9e\x42\x93\xa5\xe6\xf3\x49\xef\x0a\x3e\x2d\x5f\x01\xa8\x28\xe0\x2f\xd7\x06\xdd\x37\xbf\x8f\x68\x0f\xea\xdf\x00\x00\x00\xff\xff\xd5\xa9\x9e\x5b\x46\x05\x00\x00")

func templatesGitRestoreMtimeBytes() ([]byte, error) {
	return bindataRead(
		_templatesGitRestoreMtime,
		"templates/git-restore-mtime",
	)
}

func templatesGitRestoreMtime() (*asset, error) {
	bytes, err := templatesGitRestoreMtimeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/git-restore-mtime", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/git-restore-mtime": templatesGitRestoreMtime,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"git-restore-mtime": &bintree{templatesGitRestoreMtime, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
