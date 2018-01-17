// Code generated by go-bindata.
// sources:
// ui/explorer.glade
// DO NOT EDIT!

package ui

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

var _uiExplorerGlade = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5b\x4b\x6f\xe3\x36\x10\xbe\xe7\x57\xb0\xbc\x16\x8a\x93\x14\x7d\x1c\x6c\x2d\x90\x6e\x13\x2c\x50\x2c\x16\xd8\x6c\x7b\x24\x28\x72\x2c\x73\x43\x73\x54\x8a\xf2\xa3\xbf\xbe\x90\xdc\xbc\x2c\x52\xb6\x64\xd9\x89\x93\xdc\x9c\x84\xa3\x99\xf9\xf8\xcd\xc7\xe1\xc8\x19\x7e\x58\x4c\x35\x99\x81\xcd\x15\x9a\x11\x3d\x3f\x3d\xa3\x04\x8c\x40\xa9\x4c\x3a\xa2\xdf\x6e\xae\xa2\xdf\xe8\x87\xf8\x64\xf8\x43\x14\x91\x6b\x30\x60\xb9\x03\x49\xe6\xca\x4d\x48\xaa\xb9\x04\xf2\xd3\xe9\xc5\xd9\xe9\x19\x89\xa2\xf8\x64\xa8\x8c\x03\x3b\xe6\x02\xe2\x13\x42\x86\x16\xfe\x29\x94\x85\x9c\x68\x95\x8c\x68\xea\x6e\x7f\xa4\x0f\x8e\x4a\x33\x3a\xa8\xd6\x61\xf2\x1d\x84\x23\x42\xf3\x3c\x1f\xd1\x6b\x77\x7b\x89\x0b\x4a\x94\x1c\x51\x8b\xe8\x68\xb9\x86\x90\x61\x66\x31\x03\xeb\x96\xc4\xf0\x29\x8c\xe8\x4c\xe5\x2a\xd1\x40\xe3\x1b\x5b\xc0\x70\x70\xf7\x57\xff\x62\xc1\x0d\x1b\xa3\x28\x72\x1a\x5f\x71\x9d\x6f\x5c\x8f\x56\x81\x71\xdc\x29\x34\x34\x9e\x81\x75\x4a\x70\x5d\x33\x12\x13\xa5\xe5\xea\xb3\x2f\x8b\x1b\x44\x9d\x70\x4b\xef\x56\xb4\xcc\xa1\x4b\x1e\xf5\xb0\xc2\xa1\x5d\x16\xce\xa1\x59\xe1\xec\x12\x26\x0a\x47\x1f\x1b\x75\x73\xee\xb3\x33\xc8\xf2\x09\xce\x19\xd7\x3a\x9c\xa9\xcf\x50\xf3\x04\x34\x25\xce\x72\x93\x6b\xee\x78\xa2\x61\x44\x97\x90\xd3\xf8\xf7\xc2\x6d\xfb\x90\x22\x07\x56\x18\x09\x56\x2b\xd3\x80\xb4\xcf\x54\x09\x34\xac\xfc\x48\x63\x90\xca\x45\x22\xe4\x75\x38\x58\x21\xfc\xe4\x77\x19\x17\xb7\xca\xa4\xcd\x2e\x60\x91\x71\x23\x5b\x02\x3a\xc1\x29\xa6\x60\x00\xcb\xad\x08\x27\x34\x1c\xd4\x42\x18\x0e\xd6\xc8\xd1\x95\x2c\x98\x2d\x8f\x88\x2d\x98\x2d\x9f\x83\x2e\x41\xb7\x6f\x8e\x2f\x19\xcf\x1d\x1c\x0f\x61\xbe\x94\xe1\x3e\x03\x63\xb2\xb0\xdf\x37\x47\x19\x09\x1a\x8e\x89\x33\x1f\xab\x78\x9f\x81\x34\xb2\xc1\xf1\xeb\x66\xcd\x57\xc8\xb8\xe5\x0e\x6d\x49\x9f\x4f\x0e\xa6\xf7\xe4\xc9\xcf\x9f\x97\x38\xaf\x1b\x77\x4f\xb5\x1a\x98\x33\xa9\x2c\x08\x87\xf6\x88\x3a\x83\xcf\x30\x27\x1f\xef\xc2\x3e\x6c\xed\x8e\x51\x4b\xb0\x91\x81\xf9\x5b\xa4\x50\x43\xe9\x5e\xbc\x97\xee\x41\x4b\xb7\xc8\x36\x00\xbe\xf1\x96\xda\xe7\x2e\x85\x6b\xf5\x5b\x76\xd8\x02\x4d\x31\x2a\x02\x2e\x5f\x3f\x47\xaa\x92\x7c\x4e\x5a\xd4\xe2\xf4\xc7\xfa\x87\x71\x76\xb9\xa2\x32\x94\x1f\x99\xe0\x0e\xd2\xfa\x19\xd4\x39\x81\x0d\x49\x34\x9b\xfa\x78\xe2\xd9\x94\x3e\x08\xd5\x66\x13\x9e\xf0\xa9\x61\x1b\xf6\xd9\xdd\xe3\xdc\x1c\x87\xec\x5c\xe3\xc1\x65\xa7\x04\xe7\x75\x09\xcf\x7a\xdc\xf5\x98\x5b\xc7\xbb\x6e\x30\x56\x4d\xe7\xfa\xfa\xea\x0c\x73\xb5\x9a\xee\x9e\xd5\x97\xaf\xe5\xf3\x24\x97\x4d\x43\xdf\x2f\xdc\x80\xdc\xd3\xc8\xb7\x7d\x72\xe7\x3f\xff\xda\xc6\x80\xe5\xe0\xb6\x77\x33\x57\x12\xd8\x84\x1b\xd9\x98\xce\x36\xcd\xa0\xb0\xa8\x35\xc8\xbf\x95\x91\x38\xdf\xb3\x2c\xb4\x31\xc3\x19\x58\xcd\x97\x2c\xaf\x02\x54\x26\xed\xe9\xd4\xba\xb1\x00\x7f\x29\x98\xff\x2f\x87\xeb\x29\x77\x4e\x7b\x87\xd4\x7d\xa6\x60\x4a\x1d\x64\xce\x02\xb0\x52\xd3\x3a\x3c\x82\x0b\xa7\x66\xdc\x01\x2b\xc9\xa5\x4c\xaa\x81\x09\xad\xc4\xed\x16\x4f\xaa\xb0\x24\xd5\x1b\x25\xc3\x75\x54\xfd\x38\xa2\x39\x68\x10\x2b\x72\xd7\x4c\x42\x50\x7f\xbd\xb7\x19\x78\xfc\xd4\x4f\x63\x72\x98\xa3\xdb\x42\xae\xfe\x85\x96\x92\x9c\x4f\xac\x32\x8d\xf0\xf5\x75\x6a\x7f\x46\x07\x09\xe2\xed\x0b\xaa\xc7\xea\xfe\x96\xa0\x95\x60\x7b\xaa\xc4\x46\xe9\xe9\x9c\xee\x0e\x29\xfb\x4c\xdb\xca\x50\x18\x00\x3f\x08\x9f\x04\x9a\x07\x39\x52\x3e\x18\x76\x82\x62\x47\x38\x7c\xe6\x53\x6e\x53\x65\x68\xfc\x4b\x5b\xc3\x7b\xfd\x60\x53\x94\x40\xe3\x69\xa1\x9d\xca\xf4\x86\x00\xfc\x6a\x40\x7a\x90\x8f\x7b\xa1\x73\xcb\x0c\x46\xd4\xf1\xa4\x86\xbe\x67\xc3\xa6\x3c\x5d\x1f\x8d\xfb\x92\xed\x83\xac\x9b\x99\x16\xee\x65\x55\x19\x67\xb4\x88\xca\x1e\xd2\x2a\xd1\xf6\xa2\x14\x12\x52\x9f\x57\xc7\x13\xb6\x6a\x01\x37\x04\xec\x91\xc7\x0d\x5b\xf3\xae\x22\x5b\x9f\xb4\x8f\x9b\x9a\xf5\x39\xde\xce\x58\xec\x88\x07\xe9\xda\x54\x84\xd3\x6d\x6c\x2c\x48\x58\x1d\x48\x7f\x92\xd2\xa6\x46\x1e\xdd\x0b\xfa\xac\x8f\x57\x29\x5d\x63\xb4\x53\xee\xa2\xef\x45\xee\xd4\x78\x19\x95\xd2\xb2\x4f\xf9\xda\x6a\x6b\x0e\xa2\x7a\xbb\xf7\xb3\xed\x3a\xba\xdd\xdb\xd9\xee\xc3\x85\x6d\x6f\xba\x5d\x67\x0b\x9e\xbd\xdc\x65\xb6\x70\x89\x8b\x97\xf8\x65\xb2\x2b\x5b\x55\xcc\x0b\x1b\xe6\xb1\x05\xd7\x2a\xf5\x0f\x78\x42\x5c\xe4\x12\xe7\xac\x54\x33\x1a\xab\xc0\x04\x6e\xfb\xa6\xe0\xcf\xd5\x48\xb1\x3c\x0c\x75\xa2\x59\x75\x13\x06\x36\x2b\x4f\xc8\x17\x2a\x7f\xe1\x21\xa8\x28\x6c\x15\x79\x0f\xa3\xee\xb5\x53\x63\xb5\x57\xaa\xfe\xba\xa1\x0c\x4f\x73\x01\x93\xea\x3d\xe9\x60\x8f\xba\xd5\x69\x34\xda\x2c\x09\x3e\x8b\x8c\x4b\x59\x75\x5d\x41\x85\x6f\x35\xa4\x24\x3d\x5e\xf5\xdf\x2b\x78\x9b\x0a\x16\x85\xb5\x60\xdc\xa1\x5f\x31\xf5\x59\xc6\x3c\xf8\xb5\xd9\xf7\x0a\xde\x6f\x05\x07\x4c\x8e\xf1\xd5\x49\x23\x0a\x61\x04\x2e\xda\x34\x43\x0f\x59\x0f\x07\x8f\xfe\xa9\xe0\xbf\x00\x00\x00\xff\xff\xea\x1c\x49\xb8\xad\x30\x00\x00")

func uiExplorerGladeBytes() ([]byte, error) {
	return bindataRead(
		_uiExplorerGlade,
		"ui/explorer.glade",
	)
}

func uiExplorerGlade() (*asset, error) {
	bytes, err := uiExplorerGladeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/explorer.glade", size: 12461, mode: os.FileMode(420), modTime: time.Unix(1515872193, 0)}
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
	"ui/explorer.glade": uiExplorerGlade,
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
	"ui": &bintree{nil, map[string]*bintree{
		"explorer.glade": &bintree{uiExplorerGlade, map[string]*bintree{}},
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
