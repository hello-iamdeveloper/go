// Code generated by go-bindata.
// sources:
// latest.sql
// migrations/01_init.sql
// migrations/02_table_names.sql
// DO NOT EDIT!

package db

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

var _latestSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x49\x2c\x49\x2d\x2e\xd1\x2b\x2e\xcc\xe1\x02\x04\x00\x00\xff\xff\xff\x32\xce\x0c\x0b\x00\x00\x00")

func latestSqlBytes() ([]byte, error) {
	return bindataRead(
		_latestSql,
		"latest.sql",
	)
}

func latestSql() (*asset, error) {
	bytes, err := latestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "latest.sql", size: 11, mode: os.FileMode(420), modTime: time.Unix(1524741625, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations01_initSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x41\x6f\x82\x40\x10\x85\xef\xfb\x2b\xe6\x28\xa9\x5e\x9a\xea\x85\x13\xad\x34\x21\xb5\x68\x09\x24\xf5\xb4\x19\xdd\x45\x27\x65\xc1\x2c\x4b\xd5\xfe\xfa\x86\x5a\x85\xad\xa2\xe9\x75\xdf\xdb\x99\xf7\x3e\xd8\xc1\x00\xee\x14\xad\x34\x1a\x09\xc9\x86\x3d\x45\xbe\x17\xfb\x10\x7b\x8f\x13\x1f\xbc\xca\xac\x0b\x4d\x5f\x52\xc4\x1a\xf3\x12\x97\x86\x8a\x1c\x7a\x0c\x80\x04\x2c\x68\x55\x4a\x4d\x98\xf5\x19\x80\x69\x74\x4e\x02\x3e\x51\x2f\xd7\xa8\x7b\xa3\x07\x07\xc2\x69\x0c\x61\x32\x99\xd4\x36\x25\x55\xd1\x29\xb6\x67\xec\x84\x06\x23\x77\xc6\x32\xe0\x29\x0e\x47\x03\x86\x94\x2c\x0d\xaa\x8d\xe5\x11\x68\xf0\xfc\x26\x03\x98\x45\xc1\xab\x17\xcd\xe1\xc5\x9f\x43\x8f\x84\xc3\x1c\x97\xfd\x69\x9b\x65\xc5\x56\x8a\xe7\xe0\x62\xc3\x1c\x95\x3c\x45\xbf\x1f\x0e\xed\xec\xa2\x50\x48\x79\xb7\xbe\xa9\x16\x19\x2d\xf9\x87\xdc\xc3\x8f\x61\x38\xb2\x75\x3c\xec\xee\xee\x75\x16\x9f\x39\xd0\x14\x48\xc2\xe0\x2d\xf1\x21\x08\xc7\xfe\x3b\x60\x4a\x7c\xb1\xe7\xbf\x91\xa6\x61\xbb\xd8\xe1\xd0\x71\xaf\x5d\x6c\x65\xb5\x2f\x37\x42\x17\xbb\xa4\x94\xfa\x22\xbd\x94\xf8\x75\x80\x29\xf1\x5b\x0c\x53\xe2\xb7\x30\x56\xa5\xd4\xed\xff\xef\x6c\xc6\xff\x39\x3b\x5d\x94\xab\x9a\x95\x95\x89\x1f\xd7\x37\xd8\x0e\x40\x2c\x57\xff\x98\xb2\x9e\xcc\xda\xcf\x6f\x5c\x6c\x73\x36\x8e\xa6\xb3\x6b\xcf\xcf\xb5\x1c\xc7\x8f\x73\xe9\xb4\xde\xed\xb2\xef\x00\x00\x00\xff\xff\x02\xc5\x23\x8a\xe0\x03\x00\x00")

func migrations01_initSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations01_initSql,
		"migrations/01_init.sql",
	)
}

func migrations01_initSql() (*asset, error) {
	bytes, err := migrations01_initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/01_init.sql", size: 992, mode: os.FileMode(420), modTime: time.Unix(1524490129, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations02_table_namesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x70\x2c\x2d\xc9\xc8\x2f\xca\xac\x4a\x4d\x09\x29\x4a\xcc\x2b\x4e\x4c\x2e\xc9\xcc\xcf\x53\x08\x72\xf5\x73\xf4\x75\x55\x08\xf1\x57\x48\x84\xcb\xc7\x97\x20\x14\x58\xa3\x9a\x91\x93\x93\x5f\x9e\x9a\xe2\xe6\x89\xac\x0f\x22\x16\x9f\x96\x89\x55\x6d\x68\x71\x6a\x11\x16\xd5\xa5\xc5\xa9\x45\xd6\x5c\x5c\xc8\xee\x75\xc9\x2f\xcf\x43\x31\x01\xbb\x8b\x90\x0c\xc3\xea\x25\x54\x57\x20\x5c\x87\xac\x0f\xe6\x0d\xec\x6a\x4b\x51\x9d\x8c\xe4\x11\x6b\x2e\x40\x00\x00\x00\xff\xff\x43\x1a\x22\x86\x61\x01\x00\x00")

func migrations02_table_namesSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations02_table_namesSql,
		"migrations/02_table_names.sql",
	)
}

func migrations02_table_namesSql() (*asset, error) {
	bytes, err := migrations02_table_namesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/02_table_names.sql", size: 353, mode: os.FileMode(420), modTime: time.Unix(1524490129, 0)}
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
	"latest.sql": latestSql,
	"migrations/01_init.sql": migrations01_initSql,
	"migrations/02_table_names.sql": migrations02_table_namesSql,
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
	"latest.sql": &bintree{latestSql, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"01_init.sql": &bintree{migrations01_initSql, map[string]*bintree{}},
		"02_table_names.sql": &bintree{migrations02_table_namesSql, map[string]*bintree{}},
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

