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

var _uiExplorerGlade = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5b\xdb\x6e\xe3\x36\x13\xbe\xcf\x53\xf0\xe7\xed\x0f\xc5\x49\x8a\x1e\x2e\x6c\x2d\x90\x6e\x13\x2c\x50\x2c\x16\xd8\x6c\x7b\x49\x50\xe4\x58\xe6\x86\xe6\xa8\x14\xe5\x43\x9f\xbe\x90\xbc\x39\x59\xa4\x6c\xc9\xb2\x13\x27\xb9\x53\x12\x8e\xe6\xf4\xcd\xc7\xe1\x50\x19\x7e\x58\x4c\x35\x99\x81\xcd\x15\x9a\x11\x3d\x3f\x3d\xa3\x04\x8c\x40\xa9\x4c\x3a\xa2\xdf\x6e\xae\xa2\xdf\xe8\x87\xf8\x64\xf8\xbf\x28\x22\xd7\x60\xc0\x72\x07\x92\xcc\x95\x9b\x90\x54\x73\x09\xe4\xa7\xd3\x8b\xb3\xd3\x33\x12\x45\xf1\xc9\x50\x19\x07\x76\xcc\x05\xc4\x27\x84\x0c\x2d\xfc\x53\x28\x0b\x39\xd1\x2a\x19\xd1\xd4\xdd\xfe\x9f\x3e\x28\x2a\xc5\xe8\xa0\x5a\x87\xc9\x77\x10\x8e\x08\xcd\xf3\x7c\x44\xaf\xdd\xed\x25\x2e\x28\x51\x72\x44\x2d\xa2\xa3\xe5\x1a\x42\x86\x99\xc5\x0c\xac\x5b\x12\xc3\xa7\x30\xa2\x33\x95\xab\x44\x03\x8d\x6f\x6c\x01\xc3\xc1\xdd\x5f\xfd\x8b\x05\x37\x6c\x8c\xa2\xc8\x69\x7c\xc5\x75\xbe\x71\x3d\x5a\x05\xc6\x71\xa7\xd0\xd0\x78\x06\xd6\x29\xc1\x75\x4d\x48\x4c\x94\x96\xab\x67\x9f\x17\x37\x88\x3a\xe1\x96\xde\xad\x68\xe9\x43\x17\x3f\xea\x66\x85\x4d\xbb\x2c\x9c\x43\xb3\x8a\xb3\x4b\x98\x28\x1c\x7d\x2c\xd4\x4d\xb9\x4f\xce\x20\xcb\x27\x38\x67\x5c\xeb\xb0\xa7\x3e\x41\xcd\x13\xd0\x94\x38\xcb\x4d\xae\xb9\xe3\x89\x86\x11\x5d\x42\x4e\xe3\xdf\x0b\xb7\xed\x4b\x8a\x1c\x58\x61\x24\x58\xad\x4c\x43\xa4\x7d\xa2\x4a\xa0\x61\xe5\x23\x8d\x41\x2a\x17\x89\x90\xd6\xe1\x60\x15\xe1\x27\xbf\xcb\xb8\xb8\x55\x26\x6d\x56\x01\x8b\x8c\x1b\xd9\x32\xa0\x13\x9c\x62\x0a\x06\xb0\x4c\x45\xd8\xa1\xe1\xa0\x66\xc2\x70\xb0\x06\x8e\xae\x60\xc1\x6c\x79\x44\x68\xc1\x6c\xf9\x1c\x70\x09\xaa\x7d\x73\x78\xc9\x78\xee\xe0\x78\x00\xf3\xa5\x34\xf7\x19\x10\x93\x85\xf5\xbe\x39\xc8\x48\xd0\x70\x4c\x98\xf9\x58\xd9\xfb\x0c\xa0\x91\x0d\x8a\x5f\x37\x6a\xbe\x42\xc6\x2d\x77\x68\x4b\xf8\x7c\x72\x30\xbd\x07\x4f\x7e\xfe\xbc\xc0\x79\xdd\x71\xf7\x54\xab\x81\x39\x93\xca\x82\x70\x68\x8f\xa8\x33\xf8\x0c\x73\xf2\xf1\xce\xec\xc3\xd6\xee\x18\xb5\x04\x1b\x19\x98\xbf\x43\xe8\x1e\x42\x63\xa5\x8f\x88\xf2\x4b\xf4\x5c\x29\x7d\x60\xd2\x97\x28\x8a\x29\x18\xf7\x56\xa1\xd3\xc0\xfa\x17\xef\xac\x7f\xd0\x92\x2d\xb2\x0d\x01\xdf\x38\xe0\xe8\x33\x4b\xe1\x42\xfd\x96\x1d\xb6\x44\x53\x8c\x8a\x80\xca\xd7\x8f\x91\xaa\x24\x9f\x13\x16\x35\x3b\xfd\xb6\xfe\x61\x9c\x5d\xae\xa0\x0c\xe5\x23\x13\xdc\x41\x5a\x6f\x5f\x3a\x3b\xb0\xc1\x89\x66\x51\x1f\x4e\x3c\x49\xe9\x03\x50\x6d\x92\xf0\x04\x4f\x0d\x69\xd8\xe7\xc1\x10\xe7\xe6\x38\x68\xe7\x1a\x0f\x4e\x3b\x65\x70\x5e\x17\xf1\xac\xdb\x5d\xb7\xb9\xb5\xbd\xeb\x02\x63\xd5\xb4\xaf\xaf\xaf\xce\x30\x57\xab\x8b\x81\xb3\xfa\xf2\x35\x7f\x9e\xf8\xb2\xe9\xbe\xe0\x0b\x37\x20\xf7\x74\x5b\xd0\xde\xb9\xf3\x9f\x7f\x6d\x23\xc0\x72\x70\xdb\xab\x99\x2b\x09\x6c\xc2\x8d\x6c\x74\x67\x9b\x66\x50\x58\xd4\x1a\xe4\xdf\xca\x48\x9c\xef\x99\x16\xda\x88\xe1\x0c\xac\xe6\x4b\x96\x57\x06\x2a\x93\xf6\xb4\x6b\xdd\x58\x80\xbf\x14\xcc\x7f\xd0\xe1\xba\xcb\x9d\xdd\xde\xc1\x75\x9f\x28\x98\x92\x07\x99\xb3\x00\xac\xe4\xb4\x0e\xaf\xe0\xc2\xa9\x19\x77\xc0\x4a\x70\x29\x93\x6a\x60\x42\x2b\x71\xbb\xc5\x9b\xaa\x58\x92\xea\x32\xd2\x70\x1d\x55\x3f\x8e\x68\x0e\x1a\xc4\x0a\xdc\x35\x91\x50\xa8\xbf\xde\xcb\x0c\x3c\x7a\xea\xbb\x31\x39\xcc\xd6\x6d\x21\x57\xff\x42\x4b\x4a\xce\x27\x56\x99\xc6\xf0\xf5\xb5\x6b\x7f\x46\x07\x09\xe2\xed\x0b\xaa\xc7\xea\xfc\x96\xa0\x95\x60\x7b\xaa\xc4\x46\xea\xe9\xec\xee\x0e\x2e\xfb\x44\xdb\xd2\x50\x38\x00\xfe\x20\x7c\x12\x68\x1e\xe8\x48\xf9\xc2\xb0\x53\x28\x76\x0c\x87\x4f\x7c\xca\x6d\xaa\x0c\x8d\x7f\x69\x2b\x78\xcf\x1f\x6c\x8a\x12\x68\x3c\x2d\xb4\x53\x59\x78\xfc\xf3\xe3\x2d\x5e\x36\x20\x3d\xd0\xc7\x3d\xd1\xb9\x65\x06\x23\xea\x78\x52\x8b\xbe\x27\x61\x53\x9e\xae\x8f\xd8\x7c\xce\xf6\x01\xd6\xcd\x48\x0b\xf7\xb2\xaa\xb4\x33\x5a\x44\x65\x0f\x69\x95\x68\x7b\x50\x0a\x11\xa9\x4f\xab\xe3\x09\x5b\xb5\x80\x1b\x0c\xf6\xd0\xe3\x86\xd4\xbc\xb3\xc8\xd6\x3b\xed\xe3\xa6\x66\x7d\x8e\xb7\x73\x2c\x76\x8c\x07\xe9\xda\x54\x84\xdd\x6d\x6c\x2c\x48\x98\x1d\x48\x7f\x94\xd2\xa6\x46\x1e\x9d\x0b\xfa\xac\x8f\x57\x49\x5d\x63\xb4\x53\xee\xa2\xef\x45\xee\xd4\x78\x19\x95\xd4\xb2\x4f\xfa\xda\x2a\x35\x07\x61\xbd\xdd\xfb\xd9\x76\x1d\xdd\xee\xed\x6c\xf7\xe1\xc2\xb6\x27\xdd\xae\xb3\x05\x4f\x2e\x77\x99\x2d\x5c\xe2\xe2\x25\x7e\x87\x78\x65\xab\x8a\x79\x61\xc3\x3c\xb6\xe0\x5a\xa5\xfe\x01\x4f\x08\x8b\x5c\xe2\x9c\x95\x6c\x46\x63\x15\x98\xc0\x6d\xdf\x14\xfc\xb9\x1a\x29\x96\x9b\xa1\x4e\x34\xab\x4e\xc2\xc0\x66\xe5\x0e\xf9\x42\xe9\x2f\x3c\x04\x15\x85\xad\x2c\xef\x61\xd4\xbd\xb6\x6b\xac\x72\xa5\xea\xd7\x0d\xa5\x79\x9a\x0b\x98\x54\x57\xec\x83\x3d\xf2\x56\xa7\xd1\x68\x33\x25\xf8\x24\x32\x2e\x65\xd5\x75\x05\x19\xbe\xd5\x90\x92\xf4\x78\xd4\x7f\xaf\xe0\x6d\x2a\x58\x14\xd6\x82\x71\x87\xbe\x62\xea\xb3\x8c\x79\xf0\x8b\xeb\xf7\x0a\xde\x6f\x05\x07\x44\x8e\xf1\xea\xa4\x31\x0a\xe1\x08\x5c\xb4\x69\x86\x1e\xbc\x1e\x0e\x1e\xfd\x3f\xca\x7f\x01\x00\x00\xff\xff\x20\x51\x71\xcb\xe8\x32\x00\x00")

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

	info := bindataFileInfo{name: "ui/explorer.glade", size: 13032, mode: os.FileMode(420), modTime: time.Unix(1516479064, 0)}
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

