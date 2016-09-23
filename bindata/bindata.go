// Code generated by go-bindata.
// sources:
// queries/trip/getAllTrips.sql
// queries/trip/getTrip.sql
// queries/trip/insertTrip.sql
// DO NOT EDIT!

package bindata

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

var _queriesTripGetalltripsSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x0a\x76\xf5\x71\x75\x0e\x51\xd0\x52\x70\x0b\xf2\xf7\x55\x28\x29\xca\x2c\x00\x04\x00\x00\xff\xff\x6a\x40\x12\x2a\x12\x00\x00\x00")

func queriesTripGetalltripsSqlBytes() ([]byte, error) {
	return bindataRead(
		_queriesTripGetalltripsSql,
		"queries/trip/getAllTrips.sql",
	)
}

func queriesTripGetalltripsSql() (*asset, error) {
	bytes, err := queriesTripGetalltripsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "queries/trip/getAllTrips.sql", size: 18, mode: os.FileMode(438), modTime: time.Unix(1472712624, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queriesTripGettripSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x0a\x76\xf5\x71\x75\x0e\x51\xd0\x52\x70\x0b\xf2\xf7\x55\x28\x29\xca\x2c\x50\x08\xf7\x70\x0d\x72\x55\xc8\x4c\x51\xb0\x55\x50\x31\x04\x04\x00\x00\xff\xff\xa9\xba\x03\x98\x20\x00\x00\x00")

func queriesTripGettripSqlBytes() ([]byte, error) {
	return bindataRead(
		_queriesTripGettripSql,
		"queries/trip/getTrip.sql",
	)
}

func queriesTripGettripSql() (*asset, error) {
	bytes, err := queriesTripGettripSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "queries/trip/getTrip.sql", size: 32, mode: os.FileMode(438), modTime: time.Unix(1472712617, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queriesTripInserttripSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xf2\xf4\x0b\x76\x0d\x0a\x51\xf0\xf4\x0b\xf1\x57\x28\x29\xca\x2c\xd0\x48\xce\x4f\x49\xd5\x51\x48\x4c\xca\x2f\x2d\xd1\x51\xc8\x4b\xcc\x4d\xd5\x54\x08\x73\xf4\x09\x75\x0d\x56\xd0\x50\x31\xd4\x51\x50\x31\x02\x62\x63\x4d\x6b\x40\x00\x00\x00\xff\xff\x97\xa8\x5c\xaf\x38\x00\x00\x00")

func queriesTripInserttripSqlBytes() ([]byte, error) {
	return bindataRead(
		_queriesTripInserttripSql,
		"queries/trip/insertTrip.sql",
	)
}

func queriesTripInserttripSql() (*asset, error) {
	bytes, err := queriesTripInserttripSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "queries/trip/insertTrip.sql", size: 56, mode: os.FileMode(438), modTime: time.Unix(1472713355, 0)}
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
	"queries/trip/getAllTrips.sql": queriesTripGetalltripsSql,
	"queries/trip/getTrip.sql": queriesTripGettripSql,
	"queries/trip/insertTrip.sql": queriesTripInserttripSql,
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
	"queries": &bintree{nil, map[string]*bintree{
		"trip": &bintree{nil, map[string]*bintree{
			"getAllTrips.sql": &bintree{queriesTripGetalltripsSql, map[string]*bintree{}},
			"getTrip.sql": &bintree{queriesTripGettripSql, map[string]*bintree{}},
			"insertTrip.sql": &bintree{queriesTripInserttripSql, map[string]*bintree{}},
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
