// Code generated by go-bindata.
// sources:
// migrations/01_initial.sql
// DO NOT EDIT!

package db

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _migrations01_initialSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x54\x41\x6e\xc2\x30\x10\x3c\xe3\x57\xec\x11\xd4\xf2\x02\x4e\x29\x2c\x28\x6a\x49\x90\x31\x12\x9c\x22\x37\x31\xc1\x12\xc4\xd1\xda\x88\xf2\xfb\xba\xa8\x84\xd0\x12\xaa\x2a\xaa\xd4\x53\x56\x33\xb1\x67\x3c\xeb\x75\xbf\x0f\x0f\x3b\x9d\x93\x74\x0a\x16\x25\x1b\x72\x0c\x04\x82\x08\x9e\x5e\x10\xc2\x31\x44\xb1\x00\x5c\x86\x73\x31\x87\x83\xa1\x6d\x66\xa1\xcb\x3a\x3a\x83\x30\x12\x38\x41\x0e\x33\x1e\x4e\x03\xbe\x82\x67\x5c\x41\xb0\x10\x71\x18\xf9\x1d\xa6\x18\x89\x47\xd6\x29\xe4\x4e\x81\xc0\xa5\xaf\x01\x1c\xc9\xc2\x6e\xbd\x4a\x96\x54\x38\xeb\x0d\xd8\x3d\x45\x52\xb9\x36\xc5\x1d\x49\x2f\x72\x72\x95\x5c\xe8\x6b\xdd\x8e\x3b\x96\x55\x3d\x8e\x39\x86\x93\xe8\x63\x61\xf7\xbc\xac\x07\x1c\xc7\xc8\x31\x1a\xe2\xf9\x84\x5d\x8f\xfe\xe4\x6c\x5f\x64\x8a\x72\x32\xfe\x9b\xfc\x5f\x97\x56\x3b\xd5\xca\xd7\x5b\x92\x1a\x43\x75\xf6\xf8\x15\xf1\x9d\x6d\x61\xfe\x9e\x7b\x49\x4e\xaf\x65\xea\x5a\x9d\xc0\x27\xb0\xfb\x83\x64\x37\xda\x3a\x43\x3a\x95\xdb\x64\xad\xf3\x3d\xb5\x8b\x99\x64\x5a\xd5\xa9\xb4\xae\x1a\x1a\x59\x96\x4a\x92\xba\x8a\xfb\x55\x93\xdb\x24\x47\x8f\xd7\xd1\x4c\xc9\x1b\xa8\xb4\xd6\xa4\xfa\x34\x74\xb5\x2e\xfd\x3e\x89\xfa\x23\x31\x32\x87\x82\xb1\x11\x8f\x67\x97\x64\xae\x5e\x88\xc1\x6d\xf2\x73\x4c\x1a\xd8\x1b\x03\xd5\xf0\xe7\xe9\x52\x37\x70\xd5\x95\x69\xe0\xbf\xb7\x6d\xc0\xde\x03\x00\x00\xff\xff\x75\x89\x10\x30\x01\x05\x00\x00")

func migrations01_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations01_initialSql,
		"migrations/01_initial.sql",
	)
}

func migrations01_initialSql() (*asset, error) {
	bytes, err := migrations01_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/01_initial.sql", size: 1281, mode: os.FileMode(436), modTime: time.Unix(1443591016, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"migrations/01_initial.sql": migrations01_initialSql,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"migrations": &bintree{nil, map[string]*bintree{
		"01_initial.sql": &bintree{migrations01_initialSql, map[string]*bintree{
		}},
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

