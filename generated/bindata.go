// Code generated by go-bindata.
// sources:
// files/cf-mgmt.sh
// files/cf-mgmt.yml
// files/pipeline.yml
// files/security-group.json
// files/vars.yml
// DO NOT EDIT!

package generated

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

var _filesCfMgmtSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xd4\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\x4f\xcd\x2b\x53\x48\x4a\x2c\xce\xe0\x2a\x4e\x2d\x51\xd0\x4d\xe5\x4a\x4e\x51\x48\xce\xcf\x4b\xcb\x4c\xd7\x2d\x4a\x2d\xc8\xe7\x4a\x4e\xd3\xcd\x4d\xcf\x2d\x51\x50\x71\x76\x8b\xf7\x75\xf7\x0d\x89\x77\xf6\xf7\xf5\x75\xf4\x73\xe1\x02\x04\x00\x00\xff\xff\x4b\xdb\x02\xbb\x43\x00\x00\x00")

func filesCfMgmtShBytes() ([]byte, error) {
	return bindataRead(
		_filesCfMgmtSh,
		"files/cf-mgmt.sh",
	)
}

func filesCfMgmtSh() (*asset, error) {
	bytes, err := filesCfMgmtShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/cf-mgmt.sh", size: 67, mode: os.FileMode(493), modTime: time.Unix(1495821632, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesCfMgmtYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\xcd\x41\x8a\xc3\x30\x0c\x85\xe1\xbd\x4f\x21\xb2\x1e\x4f\xf6\xbe\xcc\x60\x3c\x8a\x23\x12\x5b\x46\x92\x43\x43\xe9\xdd\xeb\x94\x52\xe8\xf6\xfd\xf0\x3d\xef\xbd\x6b\x7b\xb4\x85\xa5\x04\xd8\xa9\xf6\x9b\x73\x54\x62\xc6\x3f\x41\xe5\x2e\x09\x83\x03\xb0\xb3\x61\x80\x7f\x4e\x1b\x8a\x7f\xe5\x31\xbe\x33\xdc\x05\x1b\x2b\x19\xcb\x19\xa0\xd1\xc1\x16\x77\x45\x39\x28\xa1\xce\x69\xf1\x25\x17\xfb\x01\x8b\x39\xc0\x34\xae\x50\x6d\x7a\x8c\x93\xda\xba\xe9\x85\x7b\xa8\xb1\x0c\x27\x71\x5d\x28\xfb\x4b\x73\x4e\x7a\xbd\x5a\x8b\xb6\x7e\x95\x39\xd1\x6c\x51\xb7\x8f\xfc\xab\xab\x7b\x06\x00\x00\xff\xff\x0a\x5a\x24\x71\xc6\x00\x00\x00")

func filesCfMgmtYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesCfMgmtYml,
		"files/cf-mgmt.yml",
	)
}

func filesCfMgmtYml() (*asset, error) {
	bytes, err := filesCfMgmtYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/cf-mgmt.yml", size: 198, mode: os.FileMode(420), modTime: time.Unix(1495821632, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesPipelineYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x96\x51\x6f\x9b\x30\x10\xc7\xdf\xf3\x29\xee\x0b\xd0\x68\x0f\x7b\xe1\x2d\x0a\xb4\x42\x2a\x49\x15\x32\x4d\xd3\x34\x59\x2e\x5c\x98\x37\x8c\x3d\xdb\x6c\x8a\xa2\x7c\xf7\x19\x48\x43\x42\x49\xda\x25\x6b\x1f\xe0\x8d\xf6\x8e\xfb\xdf\xfd\xfe\x9c\x63\x85\x5a\x14\x2a\x46\xed\x8e\x1c\xc8\x29\x47\x17\x62\x91\xaf\x58\xea\x28\x94\x62\x04\x60\xd6\xd2\xfe\x2f\x65\xc6\x3e\xd7\xa9\xae\x7d\x02\x28\x14\x73\x61\xb3\xb1\x01\x52\x66\x12\xfb\xf7\x76\x5b\x45\x1e\x15\xcd\xe3\xef\x2e\x70\xaa\x0d\xaa\x7d\xd9\x0f\x1f\xf9\xbe\x9c\x61\x1c\x9b\x7a\xb0\x61\xb9\xcd\xfc\x4d\xb3\x2a\x6b\x3b\x1a\xfd\x10\x8f\x87\x0d\x29\xa4\x06\x1d\xa1\x52\x6d\x5f\x92\x19\xcd\xcb\x16\x1c\x48\xd1\xb4\xbb\xb5\x02\x8a\xa5\x29\x2a\xab\xa1\x0a\xac\xd2\x0c\xd5\x3f\xdb\x45\x00\x56\x2c\x3b\x9e\x75\x1c\xb3\x71\x99\xaa\xc7\xf1\xca\xe1\x29\x37\x37\x6b\x9e\x55\xa9\x92\x2a\xca\x75\x3d\x36\x40\xf4\x25\x5a\xfa\x21\xf1\xe6\xe1\x24\x98\x95\x08\xf4\xda\xce\xc9\x49\x22\x38\x65\xf9\x8e\x01\xc0\xa7\xc8\x5f\x90\xc0\x2b\x13\x0a\x8d\x8a\xb0\x64\x1f\x7a\x98\x44\xd1\xe7\xf9\xa2\x8a\x49\xaa\xf5\x1f\xa1\x9a\xe0\x74\x3e\xbb\x0d\xee\x88\x17\x2c\x5c\xb8\x19\xd7\xfd\x3d\x85\xee\x03\x7f\xb6\x24\x91\x3f\x5d\xf8\xcb\xf2\xe5\x38\x63\x98\x1b\xa2\xd1\x0e\x67\x9a\x0a\xb7\x24\xbc\x0b\x97\x64\x3a\x0f\xc3\xc9\xcc\x3b\x1e\xfd\x89\x69\x82\x19\x5e\xc1\xb4\x4a\xab\x1d\x3d\x8d\xfc\x58\x63\x50\xc8\x0f\x47\x6f\x7d\xc6\x5a\x52\xbb\x6e\x2f\x43\x2f\xdb\xc4\xc4\x85\xaf\x07\xf6\x7d\x7b\xcd\x27\xbe\x17\xe8\x33\xf1\x7b\x6f\xf2\x40\x0e\x35\xb2\x84\x4a\xf2\x5c\xe8\xc4\x2e\xec\x18\xb5\xb6\xe1\x2d\xac\xf9\x97\x4d\x19\x84\x73\xa7\x76\xa5\x65\x49\x21\x93\x6b\xb6\xa5\x7e\xef\xfc\xbe\xb4\x25\x06\x45\xfd\x78\xf8\x2e\xea\x4e\xd9\xe2\x7f\x47\xff\xaa\x7d\xe8\x6c\xa3\xcf\xf6\x5c\x7a\x9c\x75\x80\xea\xb4\xf2\x57\x21\x2c\xa7\xb7\x5f\xa3\x46\xa8\xcf\x6e\x9d\xb7\x61\x87\xa0\xd3\x07\x5b\xc8\x5e\x94\xcd\xda\x49\x95\x28\xe4\x7b\x18\xf2\x5c\x71\xb8\xce\xb4\x59\xb4\x2c\xb2\x3f\xe3\x97\x9e\x79\x57\xdf\x00\x3a\x5a\xe8\xb3\x4f\x57\x9e\x77\x0d\xa6\x0e\x0b\x2f\x3d\xeb\x5e\xbe\x60\x77\x89\xf4\xd9\xa5\x73\xf8\x77\x00\xfe\x06\x00\x00\xff\xff\xdd\xfe\x38\xd9\x40\x10\x00\x00")

func filesPipelineYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesPipelineYml,
		"files/pipeline.yml",
	)
}

func filesPipelineYml() (*asset, error) {
	bytes, err := filesPipelineYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/pipeline.yml", size: 4160, mode: os.FileMode(420), modTime: time.Unix(1497466783, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesSecurityGroupJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8a\x8e\xe5\x02\x04\x00\x00\xff\xff\x44\xd2\x68\x70\x03\x00\x00\x00")

func filesSecurityGroupJsonBytes() ([]byte, error) {
	return bindataRead(
		_filesSecurityGroupJson,
		"files/security-group.json",
	)
}

func filesSecurityGroupJson() (*asset, error) {
	bytes, err := filesSecurityGroupJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/security-group.json", size: 3, mode: os.FileMode(420), modTime: time.Unix(1495821632, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesVarsYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x8e\x41\x6e\x43\x21\x0c\x44\xf7\x9c\x62\x6e\xd0\x7d\x55\x71\x15\xe4\xf2\xfd\x53\x4b\xf9\x18\xd9\x46\x51\x6e\x5f\x28\xca\xae\xab\xac\x98\x79\x0c\xc3\xdc\x24\x8a\x71\xd7\x32\x4c\x3e\xf1\xf5\xd4\x61\x98\x0c\x8b\x61\xb2\x9c\xfc\xe9\xc1\x57\x39\xf4\x22\x69\xaf\x48\x3d\xb1\x39\x36\xcf\x69\x38\x5b\x91\x63\x06\x96\x02\xd5\xaa\xa3\x05\x1e\x12\x3f\xe8\x6c\x97\xb8\x8b\x36\x84\xa2\x1a\x53\x30\xd4\x6e\xfe\xe1\x9d\x2a\x7b\x4e\x9d\xdc\x1f\x6a\xeb\xf9\x4b\x42\x4f\xbc\x53\x55\xef\xc2\x2d\x8a\xf3\xbc\x8c\xd9\xb7\x3d\xb6\xc7\xa9\x86\x41\x54\xb7\xd8\x9b\x73\xba\x1f\xd4\xcb\x7f\x1b\xe6\x1f\xdf\xd2\xfe\xce\x95\xc9\xe9\x37\x00\x00\xff\xff\x27\x4c\x31\xa2\x30\x01\x00\x00")

func filesVarsYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesVarsYml,
		"files/vars.yml",
	)
}

func filesVarsYml() (*asset, error) {
	bytes, err := filesVarsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/vars.yml", size: 304, mode: os.FileMode(420), modTime: time.Unix(1495821632, 0)}
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
	"files/cf-mgmt.sh": filesCfMgmtSh,
	"files/cf-mgmt.yml": filesCfMgmtYml,
	"files/pipeline.yml": filesPipelineYml,
	"files/security-group.json": filesSecurityGroupJson,
	"files/vars.yml": filesVarsYml,
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
	"files": &bintree{nil, map[string]*bintree{
		"cf-mgmt.sh": &bintree{filesCfMgmtSh, map[string]*bintree{}},
		"cf-mgmt.yml": &bintree{filesCfMgmtYml, map[string]*bintree{}},
		"pipeline.yml": &bintree{filesPipelineYml, map[string]*bintree{}},
		"security-group.json": &bintree{filesSecurityGroupJson, map[string]*bintree{}},
		"vars.yml": &bintree{filesVarsYml, map[string]*bintree{}},
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

