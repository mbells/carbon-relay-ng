// Code generated by go-bindata.
// sources:
// admin_http_assets/app.css
// admin_http_assets/app.js
// admin_http_assets/index.html
// DO NOT EDIT!

package web

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

var _admin_http_assetsAppCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8a\xce\x4b\x8f\xb1\x4a\xce\xc9\x4f\xcc\x8e\xd5\x51\x00\x72\x74\xe1\xec\x94\xc4\x92\x44\x5d\x64\x81\x0a\x64\x9e\x1e\x8c\x0d\x64\x22\x24\x20\xe2\x19\x99\x29\xa9\x0a\xd5\x5c\x0a\x40\x90\x92\x59\x5c\x90\x93\x58\x69\xa5\x90\x97\x9f\x97\xaa\xa0\x98\x99\x5b\x90\x5f\x54\x92\x98\x57\x62\xcd\x55\x0b\x08\x00\x00\xff\xff\x9c\x9e\x21\xb0\x7a\x00\x00\x00")

func admin_http_assetsAppCssBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsAppCss,
		"admin_http_assets/app.css",
	)
}

func admin_http_assetsAppCss() (*asset, error) {
	bytes, err := admin_http_assetsAppCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/app.css", size: 122, mode: os.FileMode(436), modTime: time.Unix(1496310245, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _admin_http_assetsAppJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x58\x51\x6f\xdb\x36\x10\x7e\xcf\xaf\x20\x84\x00\xa2\x10\x47\xce\xf6\x30\x60\x0a\x82\x2e\xeb\x56\xac\x0f\x01\x86\xac\xdb\x1e\x32\x17\x60\x24\x46\x11\x42\x93\x06\x49\x39\x31\x5c\xfd\xf7\x1d\x29\x4a\xa2\x68\x35\x4d\x1d\x6f\x35\x10\xd8\xa2\xee\xbe\xfb\xee\xe3\xf1\x48\x66\x4d\x24\x22\xab\x15\xba\x40\x9c\x3e\x22\xc2\xcb\x9a\x11\x99\x2e\x45\x51\x33\x8a\xa3\x9c\xc8\x5b\xc1\x4f\x25\x65\x64\x73\xca\xcb\x68\x86\x6e\x22\x5e\x5e\x53\x25\x6a\x99\x53\x78\x8c\xea\x2a\xbd\x15\x42\x2b\x2d\xc9\x2a\x5a\x24\xe7\x47\x47\x80\x96\xe6\x82\x6b\x29\x18\xa3\x12\x47\x57\xa4\xe2\x6f\x35\xb3\xbe\xc7\x2a\x17\x2b\xeb\x77\x2c\x3d\x90\x63\x88\x47\x8c\xc5\x5d\xcd\x73\x5d\x09\x8e\x5b\xc3\x19\xea\xcd\xe0\xa7\x35\x4a\xb6\x47\x08\xb5\x6f\x53\x02\xf8\x5a\x01\xf5\x9b\xc5\x39\x8c\x9a\x54\xde\x0a\x7e\x57\x95\x30\xd4\x3b\xe2\x68\x9e\xdb\xc1\x79\x94\x74\x56\x1f\xc8\x2d\xa3\x81\x91\x36\x63\x9e\xcd\x35\x7d\x94\x95\xa6\x32\x30\x93\x6e\x58\xcd\xb3\x8a\x17\xf4\x69\x70\xf8\x99\x91\xfc\x81\x55\x4a\x07\x1e\xb7\xdd\xf8\xae\xcb\x65\x59\x4a\x5a\x12\x2d\xc2\x28\xa4\x7f\xb1\xeb\x74\x2d\x6a\x1d\x92\x97\x66\x0c\x4c\x1f\xe8\x06\x54\xdc\xc2\x57\x86\xe2\x9f\xe0\x2b\x6e\xe0\xb1\xe9\x7d\x7f\xa1\x4a\x57\x9c\x18\x89\x3f\x8f\x30\x2f\x06\x2b\x3f\xfc\xd1\xa0\xfc\x9a\xb0\xaa\xb8\x2c\x0a\x00\x30\xfa\xcf\x3f\xde\x7c\xcc\x16\x27\xff\x64\x37\x67\xa7\x3f\x2e\x4e\x70\x66\x1f\x93\x37\xc7\xf3\xf3\xc0\xc7\x92\xff\xb0\x59\x51\xeb\x85\x15\xe5\x05\xbe\x64\xec\xd3\xbb\x4a\x2a\x9d\x5c\x11\x9d\xdf\x27\x9f\x30\x4c\x98\x02\xc1\x28\xd7\xbf\x11\x75\x5f\xf1\x32\x99\x87\x38\xb4\xa4\x4f\x80\x81\xfb\x8a\x49\x90\xa9\x0c\xf3\x91\x54\xd7\x92\xf7\x8f\xe6\x03\xa9\xe9\x6c\xa8\x2e\x80\xa8\x69\x32\xb2\x30\x1f\xa3\x50\xa5\xfe\x32\xf8\x00\xad\x65\x4d\xcf\x03\x0b\x2d\x37\x3b\x5e\xc8\xae\x1c\x20\xf4\xeb\xd3\xca\x21\x87\x6e\x0d\xca\x4d\x66\x78\x22\x26\xf2\x22\xde\x11\xa6\x76\x42\x36\xc1\xb3\xcb\xce\x79\xf9\xd6\x9d\x65\x63\x06\x9b\x04\x9b\x29\x0b\x66\xac\x2c\xdf\x81\x06\x87\xd6\xad\xba\x43\xed\x2b\x74\x71\x81\x22\xb2\x2e\xa3\xa9\x3c\xfb\x00\x53\xc2\x86\x59\x8e\x21\x0b\xca\x34\x39\x34\x28\x23\x4a\x1f\x1a\x73\x49\x9e\x0e\x0e\x59\xf1\x43\x43\x2a\x5d\xd0\xf5\xc1\x41\xeb\xe5\x3e\x90\xd3\xd5\xbd\xb3\x12\x9e\xad\x6d\xd7\x74\xfb\x32\xad\x8a\xa7\xa4\x25\x62\x1b\x7d\x5a\x52\x3d\xd4\x7b\x41\x34\x49\x3a\x9a\x0e\x40\xbb\xfd\xc0\xbc\x6b\xa3\xb6\x2d\xb3\x09\xa3\x60\x3b\xdc\x6e\x32\x63\xd8\xfc\xae\xec\xb2\x77\xf6\x79\xb7\x15\xc1\xab\x96\xb5\x8f\x06\x3d\xc3\xdb\x61\xda\x0e\xd2\x3e\xb6\x31\xba\x2d\xae\x28\x3c\xb3\x9d\x45\x3b\xbd\x11\x4e\x05\x49\x8f\x15\x59\x53\x9c\xa4\xfa\x9e\xf2\x81\x36\x34\xf0\xd5\x30\x6b\x2f\xe4\xe6\x99\x0e\xa2\x40\x82\xb3\x1e\x96\x4a\xb9\x83\x3a\x90\xdc\x2e\x55\x99\x21\xb0\x49\x8d\xe0\x29\xfc\x10\xb2\x59\x04\xc2\xf7\x7e\x92\x2e\xc5\x9a\x4e\xa9\x30\xcc\xb3\x29\x45\xab\xb7\x5c\xe2\xf8\x52\x52\xb4\x11\x35\x52\xb5\xfb\xf1\x48\xb8\x46\x5a\x20\x68\x23\x14\xb6\xce\x6e\x0b\x47\xb0\xbf\x40\x3f\xe7\x22\x45\x31\x3a\x41\x06\x6d\x20\xdd\xeb\xd6\x3a\xe1\x6d\x6c\xf7\xc1\x38\x03\xb3\xe6\x59\x15\xc2\xba\x01\xfd\xa0\xef\x3a\x21\x87\x1d\x1f\x6f\xcd\x36\x98\x41\xbf\x2c\xcd\xb9\xea\x3d\x87\x60\xb0\x9a\xb2\x1f\xce\x66\xe8\x6f\x52\x41\xdb\xfd\xee\xfb\xb3\x26\xa8\x85\xd1\x81\x61\x8f\x6a\x00\xff\xaf\x2c\x04\xd8\xe4\x5f\x43\xfd\x5b\x94\xca\xb4\x48\x7b\x16\xcb\x70\x12\x7b\xb6\x5c\x86\x98\x07\x29\x98\xee\x8c\x67\xd7\x9e\xf9\x8d\xb7\xbf\x57\x70\x8a\x04\xd1\x6d\x6b\x9c\xa1\x3f\x56\x42\xb0\xcc\x36\xd5\x19\x0a\xde\xb9\xd9\x31\x87\x2b\x38\x5b\xd9\x43\x55\x14\x56\x52\x17\x61\x9f\x96\x62\x5c\xbf\xb6\x9f\xfc\x47\x09\xfd\x2f\xf5\x35\x08\x67\x72\xf6\x45\xb3\xa7\xe6\x2f\x29\xd7\xea\x65\xe5\xda\xc2\x71\xdc\xfa\xcc\x46\xc2\x8f\x33\x00\xf6\x2d\xdd\x20\x87\x17\xb3\x9f\x62\xde\x2e\x0d\xff\x8e\xf2\xda\x95\xd1\xdf\x6b\x9e\x5d\x18\x7d\xc4\xd7\xaf\x0b\xb7\x11\x84\x85\x0b\x37\x96\x2f\xd1\x7f\x13\xfb\x9d\xdd\x4e\x47\xcf\xc6\x5c\x92\xcc\xad\x67\x1f\x2e\xe3\xeb\x94\xcf\x68\x86\x5e\xa0\xea\x88\x96\x7f\x37\xdb\x65\x37\x43\x7b\xcb\x06\x7f\x3c\x14\x0d\xd9\x49\x77\xb1\xcd\xb5\xc7\xde\xad\xdf\x73\xa5\x09\xcf\xed\xe5\xd2\x0e\x58\x5f\xdc\x31\xd4\x74\xb9\x62\x44\xd3\x3f\x25\xac\xd3\xb8\x5e\x41\xcd\xb5\xb3\x71\x65\x6d\xef\xf5\x92\xc5\xae\x72\x11\x50\xbe\x15\x44\x16\xdd\x12\x76\xc3\xc3\xff\x05\x32\x8f\x4b\x7f\xd9\x1f\xb1\x70\x2b\xc5\x3f\x53\x76\xe2\xbb\x6c\xec\xf7\x79\xf8\x56\x3c\x8c\x12\x1d\x9f\x49\xc7\x21\xd2\x9c\x09\x45\xb1\x0f\xeb\xdd\xdd\x9a\x1d\xe8\xdc\x38\xb1\x97\xc3\x17\x95\x5a\x56\x4a\xe1\xb8\x75\x8c\xa7\xc0\xbb\xc5\x8e\x6c\x88\xcc\x85\x9a\xf5\x77\x32\x25\xd8\x1a\x86\x87\x28\x96\x66\xf6\x59\x0a\xee\x04\xdd\xfd\x27\x07\xc0\x36\xa3\x04\xd5\x0d\x4c\xfd\xc2\x67\x72\xe4\x7f\xbb\x53\x2a\x1a\x57\x04\xd4\xbb\xaa\x99\x1e\x37\x7b\x84\x83\x09\xf2\x9a\x24\xf6\xc5\x1c\xda\x51\x63\x02\xff\x1b\x00\x00\xff\xff\xe0\x3f\x4f\xcd\x73\x12\x00\x00")

func admin_http_assetsAppJsBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsAppJs,
		"admin_http_assets/app.js",
	)
}

func admin_http_assetsAppJs() (*asset, error) {
	bytes, err := admin_http_assetsAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/app.js", size: 4723, mode: os.FileMode(436), modTime: time.Unix(1497368977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _admin_http_assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x3b\x7b\x6f\xdc\x36\xf2\xff\xe7\x53\x30\x42\x81\x24\xf8\x59\x92\xed\xb4\x68\x6b\xec\x1a\x48\xd3\xe4\xd7\xde\x35\x4d\xe0\xb4\xf7\x40\xd1\x2b\xb8\x22\x77\x45\x9b\x12\x55\x92\x5a\xdb\x17\xf4\xbb\xdf\x0c\x25\xad\x1e\xab\x95\xb4\xf1\x2b\x69\x1d\x5b\xa4\x86\xc3\x79\xcf\x90\x14\x67\x8f\xbf\x7f\xfb\xf2\x97\x7f\xbf\x7b\x45\x62\x9b\xc8\xd3\x47\x33\xfc\x43\x24\x4d\x57\x73\x8f\xa7\x1e\x49\x57\x3e\xcd\xb2\xb9\x17\x51\xbd\x50\xa9\xaf\xb9\xa4\xd7\x7e\xba\xf2\x10\x92\x53\x76\xfa\x88\x90\x59\xc2\x2d\x25\x51\x4c\xb5\xe1\x76\xee\xe5\x76\xe9\x7f\xe3\xd5\x2f\x62\x6b\x33\x9f\xff\x99\x8b\xf5\xdc\xfb\x97\xff\xeb\x0b\xff\xa5\x4a\x32\x6a\xc5\x42\x72\x8f\x44\x2a\xb5\x3c\x85\x51\x3f\xbe\x9a\x73\xb6\xe2\x8d\x71\x29\x4d\xf8\xdc\x5b\x0b\x7e\x99\x29\x6d\x1b\xa0\x97\x82\xd9\x78\xce\xf8\x5a\x44\xdc\x77\x8d\x03\x22\x52\x61\x05\x95\xbe\x89\xa8\xe4\xf3\xa3\x2d\x34\x8c\x9b\x48\x8b\xcc\x0a\x95\x36\x30\x6d\x81\xd1\xdc\xc6\x4a\x6f\x41\x48\x91\x5e\x10\x60\x7d\xee\x89\x08\x11\xc4\x9a\x2f\xe7\x5e\x10\x84\xf0\xb3\xa4\x6b\xec\x0c\xe0\x17\x00\x23\xb4\x15\x56\xf2\xd3\x97\x85\xc0\xce\x9c\xc0\x7e\xfe\x7f\x42\x59\x22\xd2\x59\x58\xbc\x74\x70\x8f\x7d\x9f\xfc\x44\x2d\x37\x16\xe6\x4b\x32\x21\x39\x23\x34\x65\x04\xe0\xc4\x52\x40\xe3\xe5\xfb\xf7\xc4\xf7\x3b\x14\x18\x7b\x2d\xb9\x89\x39\xb7\x15\x1d\x61\x98\xd0\xab\x88\xa5\xc1\x42\x29\x6b\xac\xa6\x19\x36\x00\x65\xb8\xe9\x08\x9f\x07\xc7\xc1\x61\x18\x19\x53\xf7\x05\x30\x4f\x00\x3d\x5e\x4d\xcd\x5b\x27\x20\x2a\x89\x8d\x79\xc2\xef\x70\x6e\xdf\x4d\xd0\xa1\x60\x70\x1e\x30\xc2\x0e\xb1\x3f\xfc\xf2\xe6\xa7\xaf\x88\x89\x45\xe2\xa4\x76\xc6\x4d\xa6\x52\x16\x9c\x1b\xf2\xe3\xab\x6f\x88\xc9\x33\x34\x1b\xa2\x96\x25\x20\x97\x30\x63\x6a\x4d\x21\x62\xce\x04\x25\x7f\xe6\x5c\x0b\x6e\x2a\x3e\x01\xe9\x6f\x62\x49\xa4\x05\x04\xe4\xdb\xdf\x5d\x5f\x61\x35\xc4\xe8\x68\xee\xa1\x21\x9b\x93\x30\x54\xc6\x04\x25\xd7\xc8\x28\x3a\xcc\x57\x40\xc6\x1a\x18\xfd\x3a\x38\xae\xdb\x8e\xbd\x73\x20\x79\x16\x16\x68\xa6\x62\xd4\x05\x2b\xe1\x51\xf0\x25\xe0\x2b\x5b\xfd\xd8\x1e\xff\xc6\x53\x26\x96\xbf\x23\x0b\xb3\xb0\xf0\xc8\x47\xb3\x85\x62\xd7\x44\x2b\x89\x86\xaf\xa2\x1c\xf9\xf6\x48\x2d\xb9\xd7\xe2\x0a\xac\x2b\xa5\xeb\x05\xd5\x15\xf3\x4c\xac\x49\x24\xa9\x31\x73\xaf\x7c\x51\xfc\xf1\x45\xba\xe6\xe0\xd8\x5e\x89\x0f\x7a\xc5\x8a\x3a\x3f\xc2\x71\xed\x91\xe8\x36\x54\xa4\x5c\x97\xef\xfa\xf0\xfa\x48\x64\x03\x02\x60\x68\x07\x62\xa1\x41\x47\x1b\xcd\x7b\x5d\x57\x9a\x85\x74\x83\x3e\x04\xfc\x03\x73\x45\x4a\x4a\x9a\x19\x4e\xaa\x87\xe6\xb4\xb9\x6c\x40\x57\xec\xc2\x9f\x06\x8c\xb3\xca\x0a\x8a\x46\x56\xac\x79\xeb\xad\x23\x7e\x43\xe7\x0f\x2a\xe1\x0d\xe2\x0a\x02\xa5\xe8\xa0\xdb\x35\x7e\x41\xd9\x1b\x6e\xb5\x88\x4c\x78\xfc\x65\x0c\xaa\x46\x11\x7f\x47\xd1\x58\x5d\xef\x20\xe6\x59\x98\xcb\x7e\xa1\x80\xc2\xc3\x00\xb8\xaa\x65\x01\x1a\xaf\x61\xca\x87\x47\xbb\x14\x59\xaa\x3d\x81\x96\x4b\x06\xf8\x06\xba\x24\xd7\x73\xef\x0d\x74\xbe\xb4\xb2\x61\x08\x7d\xaa\xd0\xea\xb2\x18\x29\x15\xbd\x68\x6a\x1d\x70\x58\x7c\xa1\x79\xc6\x29\xc4\xda\xa2\x43\xa4\xc4\x3d\x80\xa5\x7f\xf8\xe0\x9e\x82\xc4\xac\xfe\xfa\x0b\xd8\xc7\x46\x03\x41\x8b\x5e\xe9\x27\xcc\x3f\x3a\x6e\xeb\x2e\x3e\x3e\xfd\x07\x95\x82\x39\x7b\x05\xf7\x38\xee\xc8\xde\x52\x48\x42\x15\x8e\xa2\xe1\x7e\x23\x97\x8c\xa7\x86\xb3\x8e\xb6\x71\x4c\x95\xf6\xba\xfd\x7a\xbb\xd3\x81\x97\xe6\x0b\xa1\x3f\xde\x05\x51\xea\x9d\x40\x9c\x1c\x02\x2b\x79\xe1\x44\x69\xf0\xa0\x7e\x48\xe8\xdd\x22\x04\x21\x7b\x88\x9e\x59\x8c\x13\x7b\xb0\xc2\x40\x23\x20\x99\xa5\x58\x05\xb5\x58\xff\x90\x7c\xcd\x25\xfc\x5e\xd1\xe8\x1a\xd5\x64\x7b\xa4\x33\x36\x3a\x39\x3e\xdc\x6f\x28\xff\xc3\xc9\x60\xd7\xa0\x1d\x52\xe8\xf2\x0b\x5d\xa8\xee\xa6\x1b\x35\x8c\x78\x8a\x89\x15\x46\xf6\x9d\xa4\xd1\x85\x14\xc6\x3e\x98\x8d\xbd\xa1\x36\x8a\x49\x06\x91\x44\x5c\x0d\x5a\x9a\x83\x33\xf9\x02\xd2\xb0\x48\x57\xe3\xa0\x9a\xaf\xf8\x20\xc6\x17\x11\xea\xd1\xdc\x92\x35\x36\xe3\xc1\x02\x63\x81\x93\x55\xb0\xa8\x04\xbc\x25\xab\x9d\x52\xa9\xec\x66\x11\x14\x52\xd9\x6d\x5f\x35\xe4\x46\x2e\x53\x80\x9d\x64\x46\x00\x37\x7a\xe7\x57\xd6\x8f\x20\x0d\x63\xe2\x83\x90\xef\xc2\xa1\x88\x2e\x20\x3a\xf2\x44\xad\xf9\xc6\x80\x9e\x7e\x21\xc0\x2a\xae\x9e\x01\xd4\x26\xef\xac\xe4\x75\x16\x63\x8d\x49\x36\x4f\x7e\x31\xcc\x8f\x84\x8e\xa0\x8c\x0e\x4f\x31\x39\xdc\xc8\x15\x7a\x9c\xc1\x75\x82\x3d\x9f\xf1\x4b\x2d\x80\x72\xd3\x63\xdd\x4b\xa5\x93\xb2\x76\xd6\x25\xd8\x6b\xe8\xf2\x2a\xda\xf1\xbd\x4f\x19\x23\xee\x61\xa5\x55\x9e\x91\x98\x1a\x7f\xc9\x39\x5b\x00\xcf\x55\x7e\x59\xba\x41\x20\x16\x50\x41\x22\x30\x19\x30\x56\xcd\xfb\xf4\x19\xbc\x51\xeb\xd2\xf1\x1f\xc4\xbd\xde\x4a\x36\xe4\x03\x3f\xf3\xcb\x61\x4f\x7a\x20\x0f\xd2\xb5\x07\x55\xda\x31\xfb\x7b\x90\x0e\x94\x64\xe3\x1e\xa1\x83\x94\x5f\x4e\x01\x83\x4a\xf7\x96\xdc\x66\x63\x21\x0f\xe7\x35\xb5\xe0\xf7\x93\x6b\xcb\x41\x7a\xfc\xa2\x7f\xa0\x1b\x2c\xd2\x2c\x77\x85\x53\xa2\x18\xae\x96\x40\xec\x95\x20\x02\x30\x54\xaf\xf4\x47\x85\x8f\xcd\x59\xca\xd2\xcd\x23\x19\x44\x1b\x1e\xc3\x7b\x2c\xe2\x1c\x98\xc6\xf5\xb9\xe6\x3b\x34\xe2\x66\xc5\x54\x88\xfe\x19\xab\xcb\xb6\xaf\xa3\x75\x04\xa0\x00\xe7\xa0\x03\x74\x3b\x2c\x26\xa3\xe9\x10\x1a\xae\xb5\xd2\x41\x46\x2d\xf4\x42\xf9\xfb\xea\x2a\xe3\x91\xc5\xc5\x0a\x68\x8f\x27\x99\xbd\x26\x55\xe2\x42\x4c\x03\xe4\xb6\x13\x79\xfb\xd5\x04\xcb\xbb\x65\xb5\x40\x80\xa8\xd4\x92\xe2\xe3\xb8\x5a\x10\xec\x21\xe8\xb7\xd7\x19\x12\x99\x27\x0b\x2c\xfe\xfb\xb9\x81\x78\x56\x71\x93\xe0\x63\x2f\x37\x37\x32\x29\x40\x7b\x1b\x26\xe5\xd0\xec\x32\x29\x08\x27\x8e\x4b\x62\x15\x81\x4c\xa1\xa0\xae\x06\x78\xa0\xdb\xe9\xa1\xd8\x2b\x78\x4a\x49\xa6\x8c\xc0\x75\x5f\x09\x7d\x00\xc5\x37\xf1\x8f\x30\x9f\x81\x51\x12\x29\x20\x5b\x3d\xbb\x2b\x6b\xdc\xf1\x62\x91\x5b\x0b\x71\xad\x14\xfb\xc2\xa6\x04\xfe\xf9\x26\x71\x7f\x32\x2d\x12\xaa\xaf\xdd\xf3\x42\x2a\x4c\xb1\x85\x4e\x8b\xcc\xea\x74\xca\x84\xc1\xa4\xc0\x3a\xe2\xaa\x25\xfe\x82\x41\xbe\x2b\xa6\xd9\x87\xec\xbe\xa8\xb9\x57\xb5\x11\xa2\x0d\x6d\x57\x20\x2f\x56\x2b\xa8\xb6\xa8\x55\x63\x35\x08\x5d\xad\x6e\xad\xfc\xa8\x27\xfd\x04\x0a\x90\xd7\x79\x1a\x15\x8b\xd8\xdd\x75\xc4\xd9\x58\xa9\xfe\x36\xb7\xe8\xe2\xc8\x2c\xb5\x43\x80\x3f\x62\xc6\x05\x7e\xc9\x53\xc3\xa3\x67\x43\x90\xff\xa4\xc2\x8e\x43\xdd\x5d\x85\x43\xeb\x0a\x87\xd6\x46\xb2\x7f\x8d\x43\x83\x65\x9e\x8e\x17\x2f\x74\x52\xd5\xef\x00\x41\xd8\xaf\x13\x3b\x05\xb2\x92\xf6\x14\x58\x94\xf7\x2d\x15\x4f\x0d\xfb\xbe\x71\xf9\x74\x93\x48\x50\x90\xfd\xd0\x15\x14\x48\x23\x00\x27\xab\xf2\xda\x12\x1f\xc7\xb3\x74\xe5\x96\x75\xba\x43\xac\x65\xae\x99\x7b\x2e\x64\x00\x66\x04\x1b\xa2\xa7\x95\x08\xcb\x18\x86\x06\xf9\x91\x39\xb0\x85\x61\x67\xfa\x2b\xd5\x0f\xd4\x93\x65\xc9\xc6\x09\xa1\xeb\xd5\x01\x64\x43\x69\xe9\x01\x01\xee\xed\x01\xe6\xc4\x03\x3c\x92\x38\x80\xca\x8b\xf1\x35\xa6\x3f\x93\x27\x9f\x47\xfd\x85\x4a\x75\x51\xd1\xdb\xac\x51\x5d\x63\x5c\xb1\x25\xe0\x96\x32\x0b\x6c\xfb\xaa\xd2\x61\xbb\xa1\x32\x4b\x1c\xbb\xd4\xe9\x30\xe3\x6e\x4d\x2e\xa9\x26\xfc\x2a\xd3\xdc\x18\x97\x30\x3e\x17\x45\x15\x01\x73\xb3\x7a\xc9\xed\x12\x5b\x13\x16\x30\xb9\x1d\xab\x36\x1f\xac\x64\x46\xbe\xaa\xf0\x5e\x71\x26\x36\xed\x5e\xde\x9a\x26\x17\xfe\xe7\xb7\x43\xff\xdb\xdf\xff\xef\x8b\x70\x6f\x93\xab\x66\xb9\xa1\xd5\xd5\x68\x46\xcb\xe8\xa7\x90\x89\xa1\x10\x80\xaa\xc7\x3c\x6b\xd4\xd4\x7f\xe6\x34\xb5\xe2\xbf\xb0\x66\x23\x15\xb2\x4f\xd6\x26\x47\x34\x89\xc9\xb7\xd2\xe2\xa5\x7b\xbe\x4b\x0d\xe2\x0c\x37\xd4\x5e\x81\xe2\x63\x35\x87\xab\x21\xc4\xf0\xf9\xae\x6e\x2a\x39\x7c\xb2\x0b\x9b\x33\x08\x5e\x6e\x37\x63\x70\x63\x15\x80\xf8\xed\xed\xaa\x22\xb6\x07\x5e\xd1\x14\xab\x02\xb7\x1c\x28\x4e\xa4\x2f\xf9\x13\x29\x09\xba\x20\x16\x9a\x86\xc4\x5c\x57\x9f\x1f\xf4\x8d\x74\x3c\x90\x0b\x7e\x3d\xb8\x1c\x72\x40\x68\x1d\xe3\xa7\x1c\x0f\x72\x70\xc2\x18\xe6\xe8\x21\x90\xf7\x99\x52\x72\x08\xe0\x1d\xd4\xf3\x72\x88\x3f\x3c\xf8\x46\xef\x9d\x1f\x93\x7b\xda\x67\x46\xa9\xef\xb1\x00\xc3\x38\x50\x44\xaf\x81\x85\x07\xe4\xfb\x6b\x8f\x50\x2d\xa8\x1f\x0b\x06\x36\x08\x76\xa9\x73\xee\xbe\x85\xc0\xd0\x34\x70\x72\x58\xa1\x15\xe9\x52\x79\x6e\xfb\x19\xcc\x66\xf0\xac\x71\x7b\x04\xda\xd0\x9e\x43\x12\xb4\x01\xae\x47\xcf\x9e\x06\x07\x4f\x38\x8e\x1a\x1c\x3f\xb2\x56\xed\x8e\xdd\xd8\x8a\xf7\xdc\x9b\x2c\xd2\x7a\xd0\x71\xff\x16\xbd\x0b\x37\x4e\xea\x77\xbc\x3d\xef\x4c\xac\x69\x95\x0c\xad\x52\x07\x8c\x1b\x88\xb1\xcd\x4f\x56\x7a\x78\x3a\xc5\x30\x34\x6a\x87\x20\xd5\xb5\x46\x42\xc5\x2a\xb6\x43\x06\x09\xa1\x6b\x38\xf5\xdd\xe0\x6d\x21\x62\x47\xe3\x87\x27\x8c\xa6\x2b\xae\x9f\x9c\x90\xc7\x2c\x50\xa9\x14\x29\x3f\x20\x4f\x50\x31\xd0\x55\xf5\xfc\x85\x56\xc1\x26\x9b\xe4\xad\x4c\x32\xf5\x24\xf5\xa6\xf3\x4c\x3a\x84\xfd\xc8\x39\x68\x11\xa0\x6f\x1f\x7b\xef\x0e\x4d\x3f\x7e\x67\x80\x9b\xba\x8e\x05\x06\x33\x82\x37\x64\xa3\x31\x63\xe0\x3a\x7d\x9e\x3b\xe0\xd1\x0d\x1e\x48\xc5\x04\xb9\x3b\x2e\x32\x97\xb6\x06\xd9\x50\x17\xbe\x11\xab\xf4\xbe\x58\x19\x0e\x62\xdf\xd7\x21\xa4\x08\x65\x07\x0f\x79\xe0\xf8\x11\xdf\xed\xdc\x47\x86\xdd\x6f\xd5\xd5\x7b\x5c\x86\xb9\x22\xf8\x3b\xbf\xae\x16\x59\xbf\x40\xe2\x1d\x39\x5e\xea\x6c\xdb\x54\x85\x61\xef\xac\xf7\xc8\x44\x41\xf9\x28\x17\x2d\xe2\x0d\x4f\xd9\x0b\x29\x5d\xf9\x38\xb4\xa1\xe8\x66\x70\x48\x77\x50\xd5\x3e\x57\xab\x56\x0f\x8e\xa6\xf1\x05\x65\xf7\x3c\xad\x33\x7c\xd7\x62\xb2\x49\xfc\x81\x6b\xbd\x16\xda\xd8\xb2\xad\x34\x7e\x52\x6d\x84\xc1\xaf\xaa\x7f\xa0\x26\x1e\x39\xc2\xdd\xb9\xa4\xbc\x47\x15\xbe\x73\x79\xb2\x52\x62\xd5\x1a\x57\x63\x09\xf9\xc0\xd4\xbf\xaf\x12\x70\xc5\x40\xa3\x63\x9c\x87\x1a\xf8\x81\xd9\x68\xed\xde\x9e\x4d\xdd\xbd\x3d\xdb\x7b\xf7\xf6\x1e\x59\x2a\xd7\x7d\x15\x53\x9b\xe6\x1e\x61\xae\x46\xb1\x75\xd8\x50\xbe\xb9\x4f\x1e\x8b\x0d\x18\x28\xc7\xa2\x8b\x85\xba\xea\x7e\x34\x50\x58\x63\x51\xb6\x94\x96\xd8\xaa\x61\x5a\xfc\x7e\x62\x64\xbf\x2b\xeb\x94\x32\x04\xb4\xab\x96\x89\x84\xef\xc8\xc8\x77\xbe\xf7\x55\x07\xee\xc9\xbb\x5f\xb7\xff\x01\x61\x77\xdf\xab\xfd\x69\x7a\xdd\x68\x7c\x85\xbe\xf9\x30\x1d\xf7\x85\xc2\xcd\x87\xe8\x6e\x37\xa8\xea\xfe\xae\xba\x47\x02\x59\x45\x73\xf2\x37\xba\xa6\xef\xdd\x9d\x08\x87\x6c\xbe\xf7\x7f\x8d\x1b\x20\xe4\x1d\x3a\x1a\x23\xd4\xe2\x25\x18\x02\x79\x0c\xaf\x90\xe0\x63\x75\x93\x82\x18\xe5\xda\x19\x5d\x71\x43\xa4\xa2\x8c\x2c\x29\x24\xb6\xcd\x55\x8a\xbe\x1b\x1e\xf4\x9c\x5e\x05\x2b\xa5\x56\x92\xd3\x4c\x18\x77\xcd\x03\xfb\x42\x29\x16\x26\x3c\xc7\x9b\x28\xd7\xe1\x51\x70\x04\x3f\x65\x6b\xfc\xf6\xc8\xf4\xbb\x37\xe7\xdd\x6b\x3f\x6d\xbc\xbb\x88\x8e\xc0\x1d\x02\xa8\xa0\xf1\x48\xe9\xdc\x04\x4a\xaf\x80\xc4\xe3\xe0\xe8\x30\x2c\x3b\xa7\xdf\x70\x19\x45\x05\xa5\xb2\x51\xb9\x8e\xf8\x14\xbe\x81\x4f\x40\x12\x49\x95\xb3\x25\x8c\xe5\x1d\x71\x56\x28\x73\xe1\xd7\x92\x38\x44\xe1\x1e\x86\xcd\x3e\xdf\x66\xd2\x8c\xc8\xc2\x5d\x3c\xda\x45\x4e\xe1\x7e\xb8\x00\x0a\xc1\xf7\x2c\x4f\x20\x4a\x5b\x88\x10\x02\xfc\x2f\xcf\x70\x6b\xd5\xc5\x91\x37\x8a\x51\x19\xe0\xbd\xa0\x9e\x3b\x33\x09\xbe\xec\x5e\x8a\x99\xc5\xcf\xdb\xef\xdd\xb5\x31\xef\xf4\x57\x87\x94\x38\xdf\x3e\x21\x1f\x3e\xb8\x87\x6a\x47\x2b\x7e\xde\x72\x26\xf7\xd8\xd8\x46\x5e\xde\x60\x07\x79\x6b\xab\x78\x9b\x03\x8c\x07\xcd\xdb\x35\x0d\x80\x7a\x8e\xce\xd5\x1a\xba\xe0\x12\x29\x98\x7b\x98\xc5\xbc\xd3\x77\x45\x2e\x9b\x85\xee\x4d\x0b\xb6\x9b\x4e\x0b\xce\x71\x40\x15\x9e\x1d\x8a\x29\xf5\x59\x59\xbb\x4e\x29\x0f\xda\xa5\x34\x22\x75\xa5\xef\xae\x2a\xba\x53\x3b\x37\xe0\x87\x8f\x6d\x87\x8e\x6b\xbb\x57\x12\x76\xdf\x50\x98\x20\x66\xdc\xd2\xf0\xea\x9d\xe7\xa9\x62\xc6\x01\x95\x98\x1d\x8a\x09\x62\x3e\x3a\xfe\x3a\x38\x84\xff\x8f\x4e\x8e\x0f\x0f\xbf\x1c\xfc\x40\xa2\xa7\x66\xe9\x11\x3c\x4e\xbc\x8f\xe0\x0b\xf8\x1d\x82\x3f\x21\x4f\x62\x65\xec\x09\xde\x11\x7c\xb2\x9f\xd0\x77\xde\x39\x2b\xfc\x60\x09\xd1\xa5\x7d\xbd\xad\x27\xa9\x63\x0a\x67\x7c\x49\x73\x69\xbd\xc6\x6e\x43\x44\xd3\x88\xcb\xa7\xcf\xbc\xd3\x97\x52\x19\xbe\x9d\xab\x77\x14\x08\x65\x65\xd0\xa9\x00\x96\xad\xe4\xdf\x98\x46\x5d\xe0\x14\x45\x2c\xe9\xce\xd1\x4a\xca\x55\xf2\xae\x63\x1f\x80\xbb\xbc\x3f\x0b\x8b\xeb\xc2\xff\x0b\x00\x00\xff\xff\x3e\x3f\xeb\xf0\x3f\x3c\x00\x00")

func admin_http_assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsIndexHtml,
		"admin_http_assets/index.html",
	)
}

func admin_http_assetsIndexHtml() (*asset, error) {
	bytes, err := admin_http_assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/index.html", size: 15423, mode: os.FileMode(436), modTime: time.Unix(1497369022, 0)}
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
	"admin_http_assets/app.css":    admin_http_assetsAppCss,
	"admin_http_assets/app.js":     admin_http_assetsAppJs,
	"admin_http_assets/index.html": admin_http_assetsIndexHtml,
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
	"admin_http_assets": {nil, map[string]*bintree{
		"app.css":    {admin_http_assetsAppCss, map[string]*bintree{}},
		"app.js":     {admin_http_assetsAppJs, map[string]*bintree{}},
		"index.html": {admin_http_assetsIndexHtml, map[string]*bintree{}},
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
