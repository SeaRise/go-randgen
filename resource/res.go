package resource

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _resource_default_zz_lua = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\xdf\x8b\xdb\x30\x0c\xc7\xdf\xf3\x57\x88\xbc\x78\x03\xb7\x38\xdd\x95\x5d\x19\x7d\xd8\xb1\x63\x1c\xdc\x8f\x97\x0e\xb6\xa7\xc3\x69\x94\xd6\xd4\x71\x8a\xe3\x6c\x17\x46\xff\xf7\x21\xcb\x69\x69\x3b\xd8\xcb\x0a\x95\xbf\x52\x3e\x92\x62\x5b\x09\xba\xb4\xd8\xc1\x12\x7e\x67\x00\x00\xbe\xfd\x15\x1d\x28\x24\x14\x4a\xc2\x8c\xfe\x73\x09\x73\x25\xe1\xe3\x9c\x62\x0a\x64\xa1\x94\x3a\xc8\xc8\xaf\xb7\xda\x77\x18\x38\x47\xf4\xa1\xbe\x6d\xca\x1b\x58\xb7\xd6\xea\x80\xcb\xe4\xbf\x6e\xd0\xa1\xd7\xf6\x75\x6d\x84\xbc\xa2\x60\x09\x23\x57\x1a\x47\x80\xd5\xc1\xb8\x82\x54\x69\x9c\xf6\x83\x80\xd4\x6d\xaf\x7d\x30\xc1\xb4\x6e\xec\xe7\x2a\xac\xe3\xd3\xc3\xa7\x2c\xab\x0d\xda\xea\xb4\x95\xc9\x04\x56\x5b\x84\xba\xed\x3d\xdc\x3d\xbe\xdc\x41\x18\xf6\xd8\x81\xf6\x08\xab\x87\xe7\x1f\x14\x92\xc0\xf6\xe9\xfe\xcb\xc3\xb7\x27\xd6\xda\x55\xf0\xf8\xf2\xfc\x95\xbc\xe9\x55\xa1\xd5\xfd\xf7\xd5\x45\x21\x0a\x49\x60\xcb\x85\x58\x8f\x85\xc8\xe3\x42\x9c\x17\xdf\xdc\xb8\x40\xfb\xab\x74\xc0\x60\x1a\x1c\x35\xad\xe4\x77\x41\x37\x7b\x72\x6a\xdb\x6a\x26\xdb\xbe\xb4\x28\xf8\x1c\x4e\x3f\x41\x17\xf0\x6e\xa6\xde\x13\xf3\x53\xfb\x93\x7b\x04\x27\x93\x8b\x14\x74\x7d\x43\x78\x87\xe1\x1c\x0b\x6d\xd5\xc2\x0e\x07\x28\x6d\x5b\xc6\xf7\x0f\xf8\x16\x60\xef\xb1\x36\x6f\x57\xf5\x44\x30\x6e\x20\x52\x48\x91\x96\x06\x2b\xd3\x37\xc9\xb1\xad\xdb\xb0\xfc\x7b\x2a\xd5\x16\x52\xa4\x85\x53\x93\x43\xa9\x2c\x63\x6a\xba\xfd\xce\x6c\x1c\x9f\x1e\x29\xac\xe2\x2c\xb9\xa4\x47\x68\x87\x43\x3a\xe2\x1d\x0e\x4c\xf0\x8c\x64\x87\x2c\xab\x74\xd0\xc7\xf9\x00\x00\x3a\x09\x86\xf3\x21\x27\x38\x2f\x79\xf9\x5c\xae\x59\xb8\xde\xda\x5c\x80\x14\x79\x91\x0b\x29\xf2\x59\xb4\x2a\x17\xa9\x9d\xeb\x9b\x12\x7d\xea\x48\xb0\x90\xa2\x98\x51\xaa\x22\x33\x89\x53\xcc\x26\x06\x6e\xa2\x9c\x4d\x17\x0b\x0e\x4e\x95\x5a\x44\x70\x31\x2d\x6e\x3f\x1c\x77\xd1\x05\x6f\xdc\xe6\xac\x2c\xdd\xdb\xc6\x9a\x6e\xfb\x1f\xe4\xd8\x66\x1c\xbf\xf3\x3e\xc7\xa1\x1c\xb1\xe3\x44\xfe\x83\xab\xf8\x6b\xbe\x40\xf8\xfb\xfc\x13\x00\x00\xff\xff\xf1\x1b\xad\x88\x68\x04\x00\x00")

func resource_default_zz_lua() ([]byte, error) {
	return bindata_read(
		_resource_default_zz_lua,
		"resource/default.zz.lua",
	)
}

var _resource_english_txt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xc9\x3d\x0a\x02\x31\x14\x00\xe1\x7e\x4e\xf1\x60\x6b\x21\x01\xb1\x77\x11\x05\x7f\xaa\x54\x2b\xaf\x31\x71\x0d\x68\xdc\x26\xb1\x88\x37\xf0\x06\x76\x5e\xd1\x23\x88\xa4\x99\xe2\x9b\xb8\x3a\x66\x36\xe7\xa7\x63\xb8\xdf\x1c\xf5\xb0\xcb\x38\xbb\x30\xe4\x7f\xae\x29\x5f\xd8\xee\xdd\x9a\xe9\x91\x12\x95\x01\x4f\xcf\xc9\x07\x96\x3e\x10\xa3\xe8\x44\x08\xa2\x86\x71\x14\x2d\x68\x91\x86\x5a\xa4\xb9\x16\x69\xcb\x4b\x27\x81\x5e\x3a\xa9\x20\x3d\x16\xc3\xcc\x62\x0d\x73\xbe\x9f\xf7\xeb\x17\x00\x00\xff\xff\x08\xfe\x8a\x6a\x8a\x00\x00\x00")

func resource_english_txt() ([]byte, error) {
	return bindata_read(
		_resource_english_txt,
		"resource/english.txt",
	)
}

var _resource_expression_zz_lua = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xc1\x6a\xdc\x30\x10\xbd\xef\x57\x0c\xbe\x4c\x0a\xda\x45\x4a\x7a\x88\x29\x39\x84\x10\xc8\xad\x90\x04\x7a\x96\xe3\xd9\x8d\x88\x2c\x2f\x92\xdc\xd6\x04\xff\x7b\x19\x49\xb6\xd7\xf4\x90\xdb\x93\xde\x1b\xcf\x7b\xcf\x8a\xba\xb1\x14\xe0\x0e\x3e\x77\x00\x00\xbe\xff\x93\x0e\x4a\x80\x92\x02\xae\xa5\x80\x1b\x29\xa0\x96\x93\x48\xfc\x7e\x0f\x2f\x4f\x3f\x7f\xc1\xc3\xd3\xfd\xf3\xfd\xc3\xeb\xe3\x33\xbc\x3c\xbe\xfe\x48\xd4\xdb\xbb\xf6\x81\x62\x1a\xc7\x21\x1e\x6f\x51\x20\xae\x63\x67\xed\xa3\x89\xa6\x77\xe0\x86\xae\x21\x9f\xee\x97\xcb\x32\xe5\x5a\x3a\xf2\xcc\xb4\xdb\x1d\x0d\xd9\x76\x35\x16\xc7\x73\xb6\x89\xc6\x45\x14\x80\xd1\xb8\xb1\xc0\xd0\x69\x6b\x0b\x6e\xcc\xa9\xa0\x96\xde\x4c\xa7\xed\xd5\xf7\x94\xe3\x1b\x66\x27\x80\x47\xdb\xeb\x2c\xe8\x87\xc6\x12\x23\xb6\x7e\x95\x34\x80\xbf\xb5\xbf\x3c\x92\x1b\xba\xb4\x83\xf2\x8c\x8e\x14\x4d\x47\xf3\xd7\x00\x9b\xbe\xb7\x28\x80\x0d\x75\x14\xa2\xee\xce\xac\x1b\x49\xfb\x59\xcf\x81\x92\x3a\x98\x93\x4b\x11\x18\x50\xcb\xfc\xe0\x0a\x2e\x45\x7d\xd0\x18\x58\xf1\x41\x23\x4e\x5c\x43\xab\xa3\x5e\x4a\xc8\xcd\xe5\x1a\xdc\x60\xed\xa6\x87\x62\x08\x00\xd5\xf5\xa1\xae\x15\x93\xea\x20\x65\xcd\x60\x5f\x1f\xd4\xed\x0d\x0a\x94\x28\x70\xaf\x50\xa0\xc2\x34\x50\xf6\x72\xcc\xbb\x4f\xac\xc6\x0a\x05\x56\x4d\x95\x04\x02\xab\xbf\xe9\xcc\xbb\xaa\xd9\x22\x70\xe2\xf2\x46\xa4\x00\xe6\xa6\x65\x39\xe7\xde\xd8\x4b\x45\xac\xfc\xdc\xdf\x46\xb3\x94\xba\xea\x96\x32\xbf\x12\xf2\xd5\x7f\x9a\x0b\x3e\x44\x6f\xdc\x69\x5b\x99\xa5\x18\xc9\xe7\xbf\x7b\xb2\x26\xbc\x97\xb4\x32\x3f\xbd\x7f\x01\x00\x00\xff\xff\x89\x82\x00\xfe\x13\x03\x00\x00")

func resource_expression_zz_lua() ([]byte, error) {
	return bindata_read(
		_resource_expression_zz_lua,
		"resource/expression.zz.lua",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"resource/default.zz.lua": resource_default_zz_lua,
	"resource/english.txt": resource_english_txt,
	"resource/expression.zz.lua": resource_expression_zz_lua,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"resource": &_bintree_t{nil, map[string]*_bintree_t{
		"default.zz.lua": &_bintree_t{resource_default_zz_lua, map[string]*_bintree_t{
		}},
		"english.txt": &_bintree_t{resource_english_txt, map[string]*_bintree_t{
		}},
		"expression.zz.lua": &_bintree_t{resource_expression_zz_lua, map[string]*_bintree_t{
		}},
	}},
}}
