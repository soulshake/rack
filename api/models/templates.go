// Code generated by go-bindata.
// sources:
// models/templates/app.tmpl
// DO NOT EDIT!

package models

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

var _templatesAppTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x3d\x69\x73\xdb\x46\xb2\xdf\xf3\x2b\xa6\x50\x79\x25\x79\x4d\x51\x47\xb2\xfb\xf6\x31\xcf\x5b\x25\x53\xb4\xad\x44\x07\x1f\x29\x7b\xeb\xc5\x76\xa5\x20\x62\x28\xc2\x22\x01\x06\x00\x75\x44\xcb\xff\xbe\xdd\x73\x00\x73\x02\xa0\x8e\xac\x12\x5b\x24\xd0\xd3\xd3\xd3\xd3\xd3\xd7\xf4\x8c\x1f\x1e\x48\x44\xa7\x71\x42\x49\x10\x2e\x97\x01\x59\xaf\xbf\x23\x04\x1e\x7e\x0f\xdf\x48\xef\x0d\xe9\x1e\xc2\xef\xf2\xe1\x22\x4c\xe2\x29\xcd\x0b\xf6\xe6\x54\x7e\xe1\xaf\xe1\x0f\x21\xc1\xe1\x3f\xc7\x17\x74\xb1\x9c\x87\x05\x7d\x97\x66\x8b\xb0\xf8\x44\xb3\x3c\x4e\x93\x80\xf4\x48\x70\xb0\xb7\xbf\xb7\xb3\xf7\x3f\xf0\x7f\xd0\xe1\xe0\xfd\x34\x89\xe2\x02\xde\xe7\x41\x4f\xa0\x60\x3d\x15\x02\x07\x09\x2e\xc3\x79\x98\x4c\x68\xb6\x33\xa9\x40\xcd\xbe\xad\x46\xcb\x2c\x9d\xd0\x3c\xdf\xa8\x4d\x46\xaf\xe2\xbc\xc8\xee\x9b\x1a\x05\xc7\x49\x41\xb3\x24\x9c\x23\xc5\x24\x78\x97\xf4\x7a\x83\xdf\x57\xe1\x1c\x47\xf0\x19\x9f\x8c\xe8\x14\x3e\x56\x60\x64\xdd\x21\xc1\xff\x53\xc0\xf6\x15\x3e\x4a\x2c\xc3\x2c\xbe\x81\x7e\x1b\x90\x48\x28\x37\x8e\xb7\xc0\x9a\xeb\x31\x9d\xac\xb2\xb8\xb8\x7f\x9f\xa5\xab\x25\xb2\xf9\x41\x45\x07\xdf\x3f\x3f\x30\x6c\x38\x01\x3a\x2c\xe2\x0c\xbe\xf2\x71\x09\xa4\xc1\x30\xcc\xc2\x05\x05\xca\x59\xd3\xfa\x19\x59\x22\xec\x06\xb3\xe1\x84\x97\x63\xe9\xcf\x57\x39\x74\xab\x88\x01\x3c\xbc\xb8\x5f\x52\x4e\x78\x91\xc5\xc9\x55\xd0\xa9\x5e\x1d\xd1\x69\xb8\x9a\x17\xec\xad\xfe\x3c\x9f\x64\xf1\xb2\x90\x32\x17\x88\x57\x15\xd7\x8e\xe8\x72\x9e\xde\x2f\x68\x52\x9c\x86\x77\xf1\x62\xb5\x70\xf4\x09\x0d\xcf\x56\x8b\x4b\xa0\xc7\xd1\x25\x93\xe4\x3d\x5f\xa7\xf0\x56\xe0\x25\x4b\x9a\x4d\xa0\x9b\xf0\x8a\x92\x74\x4a\x04\x1b\x68\x4e\x8a\x94\x5c\x53\xba\x24\xd9\x2a\x49\x60\x58\xe4\x76\x16\xcf\x29\xac\x43\xa4\x0b\x87\x59\x47\x72\x9c\x3c\x92\xe4\xfd\x7a\x92\x39\xde\xe7\x23\x79\x90\xdc\xc4\x59\x9a\x20\xcd\x6e\x62\xfd\x53\x5a\x33\xa3\xce\x09\x55\x17\x64\xbb\x7e\x34\x84\xe7\xc9\xfc\x9e\x84\xf3\x79\x7a\x4b\xc2\x09\x0e\x17\x07\x5b\xcc\xe2\x9c\xa0\x0e\x9c\x66\xe9\x82\xc4\x49\x1e\x47\x14\x1e\x52\xf2\x69\xd8\xf7\xd0\x7c\x96\xaa\x2f\x0e\x11\x21\x8d\x3e\x85\xf3\x15\xe5\xab\x9a\xad\xdf\x0e\x83\x23\x5f\xad\x41\xfc\x42\xef\x5f\x9a\x4f\x8a\xca\x79\x04\x9b\x3e\xe6\x94\x8c\x57\x97\x09\x2d\x72\x81\x08\xf9\x94\x2f\xe9\x24\x9e\xde\x23\x5b\x76\x18\x8f\xe6\x69\x18\x11\xa9\x22\x08\x4d\xa2\x65\x1a\x27\x45\xfe\x22\x3c\x1b\xd1\x39\x0d\x73\xd7\x80\x9e\x5b\x67\x8c\xe8\x32\xcd\xe3\x22\xcd\x5c\x93\xf4\xb4\xce\xc6\xe9\x0a\x96\x1c\x99\xa4\xc0\xbc\xac\xea\xc6\x22\x41\xd7\xdd\xcf\x4d\xc5\x05\x88\xf6\x89\x36\x75\xb9\xe8\x8f\x5c\x61\x87\x64\x9a\x66\xe5\xa2\x70\x10\xc7\x05\xc3\x43\xd6\x09\x18\xd6\xff\x05\xcf\x00\xac\x52\xff\xa0\xd7\xe3\xc0\xbd\xde\x71\xf4\x8f\xc7\x90\x0a\xa2\x46\x72\xde\x5f\x3b\xaa\xfc\x72\xff\x32\xc4\x2d\xc5\xf2\x68\x47\xa4\x74\x90\x34\xea\x8c\xb5\xb7\x3d\x1a\xfc\xdf\xc7\xe3\xd1\xe0\xe8\x15\x39\x09\x17\x97\x51\x48\xfa\x60\x2d\xd3\xc5\x45\xba\x8c\x27\xe4\x43\x98\x44\x73\x98\x31\xb1\x1c\x88\xc4\xa8\x90\x09\xea\xfd\x84\x26\x57\xc5\x8c\x11\xb9\xaf\xbe\x32\x14\x80\x4d\x1f\x28\x3c\x37\xe7\x2a\xa6\x01\x0c\x72\xec\xb1\x0c\x6b\x60\xd0\xb0\xdf\x3f\x3e\x1a\x3d\xbb\xc8\x63\xcf\x88\xd8\xdd\xbd\xe6\x15\x9d\xc2\x1b\xe8\x45\x95\xef\x60\x98\x66\xc5\x30\x4b\x8b\x74\x92\x1a\x96\x67\x56\x14\x4b\xee\xd7\xa1\x6c\xd1\x84\x66\x0a\x5c\xf0\xe1\xe2\x62\x88\x2a\xed\x38\xc9\x0b\x5c\x69\xae\x77\x6c\xad\x53\x1f\xc4\x38\xa8\xb8\x23\xba\xcb\xeb\xfb\x1b\x3f\xb9\x43\xad\xc7\x62\x52\x33\xbe\x8b\xbe\x77\x78\xe2\x95\xbf\xb3\xf1\xf8\xc4\xec\x6a\x5e\x33\x34\x04\x7f\x5a\x57\x64\xed\x9c\xef\x11\xcd\x99\x56\xd6\x26\x5c\x59\x72\xa3\x74\xee\x31\xa3\x6c\x4d\x1c\x1f\x9e\xf6\x7a\x0c\x46\x19\x09\x74\x0e\xce\x55\x11\x53\x5d\x4b\xa2\xd9\xcb\xf3\xd5\x82\x22\xfc\x30\x9d\xc7\x93\xfb\xa3\x74\xb2\xb2\xfc\x26\xbe\x14\x4a\x5d\x81\xb1\xd4\xc1\x0e\x84\x53\xfb\xff\xad\x74\xc2\x80\xc6\x05\x28\x1f\xd1\xfe\xb3\xf6\x8a\x18\xf8\x18\xf8\x60\x3a\xa5\x13\x66\x8c\x99\xf9\x35\xb0\x09\xd2\xe3\x64\x12\x2f\x65\xc8\x33\xa6\xd9\x4d\x3c\xa1\xdc\x40\xcf\x99\x3e\xea\x86\x8b\xf0\x8f\x34\x09\x6f\xf3\xee\x24\x5d\x68\x51\x8a\x3a\xd0\x89\x50\x68\xd0\x2e\x2f\xf2\x5e\x35\xf0\xca\xba\xcb\x9f\xb5\xf6\x5d\x7d\xab\x61\x86\x80\x05\x94\x1a\x10\xbf\x1b\xe8\x8f\x91\x93\x9c\xd7\x3a\x0f\x4c\x0e\x70\xc8\xfb\x33\x88\x7a\x18\x0f\xa2\x05\x78\xc2\x10\x07\x86\x60\x85\x2d\x5e\x04\x0d\x13\xc4\x60\xda\x4c\x12\x03\xd4\x26\x0a\x19\x6b\x4d\x85\xc2\xb2\xe0\x2f\xf8\x55\x0a\x26\x7f\x40\xd6\x0d\x6c\x53\xbf\x55\x90\x6b\x4b\xc5\x2a\xa2\x5d\x23\xd6\xdc\xf4\xf4\x7a\xef\x56\x09\xa7\xaa\x95\x74\xf7\xc1\xb1\xb1\x25\x79\xfc\xc3\xdb\xd5\xe4\x9a\x16\x55\xfc\xfb\x33\xf8\x89\x5c\x34\x76\x60\xa4\xf0\x0b\xe2\xf0\x9b\xf4\x0e\x3e\x57\xe1\x30\x23\x63\x04\x81\x3a\xaa\x70\x18\xbc\x2d\x67\x80\x58\x78\xd2\x26\x56\x8e\x34\xe3\x36\x72\x57\x43\x5b\xe6\x28\x30\x22\xde\xe5\x12\xbd\x3b\x65\xe9\x0b\x78\xdc\xfd\x23\x5e\x06\xbc\x2f\xaf\x14\x0a\x13\x8c\xc8\xe2\x24\xa2\x77\x5d\x7a\x27\x62\x12\x0d\xec\x94\x2e\xc0\xb7\x1b\xc7\x7f\x30\xa6\xee\x1f\xfc\x5d\x7f\x2d\xd5\x0a\x27\xfd\x3d\x2d\x0e\x0b\x2e\x1b\x96\xee\x41\xc9\xc8\x12\x6b\x9d\x05\xa3\x55\x52\xc4\x5c\x92\x13\xe0\xfb\xb7\x5c\xef\xe0\x02\xde\xa5\x2b\x26\x61\x3f\xec\x05\x7e\x81\x70\xc7\xfb\x59\xa9\x15\x49\xd7\x13\xea\x4f\x20\xe4\xfb\x96\x5e\xb6\x01\x95\x59\x01\x15\xb4\x65\x22\x21\xe7\x0a\xa8\x06\x79\x99\xcc\xd9\x00\x3b\xb8\xb9\x85\x63\x94\x2e\x50\xe9\x1b\x07\x1e\x54\x79\xc1\x93\x36\xba\x55\x39\x5f\x15\xcb\x55\xd1\x9c\xe9\x4a\x05\x1c\xe9\xd6\xb3\xa1\x82\x6b\x9b\xda\x6a\xdf\x82\xf3\xa2\x82\x57\x53\x35\x63\x5a\x14\x86\x37\x84\x6a\x0f\xa3\x36\x2e\xbd\x62\x59\x95\x70\xa6\x95\xfd\x0e\xff\x40\x7f\x10\x1d\x32\xbc\x4a\x12\xd2\x95\xb9\x93\xe9\xc7\x2c\x4c\xae\x28\xf9\xfe\x9a\x65\x1f\x07\x09\x0c\x08\xb5\x76\x2e\x87\xc0\xd3\x61\x00\xb7\x5a\x82\x2a\x42\xb8\xf5\xba\x32\x55\x76\x96\x0d\x35\x42\xa0\x2c\xa1\x60\x90\x84\x97\x73\x1a\xe9\x18\xaa\xa6\x67\x29\x5b\x8c\xce\x74\x1d\x3e\x19\x83\x6a\x99\xf0\xf5\xba\xa7\x6a\x17\x1d\xdf\x3b\xa9\x56\xb8\x02\x43\x8d\xb3\xb3\xcf\xa8\x10\x84\x54\x7c\xa9\xe7\x90\xcc\xa6\x19\xdc\xa1\x3e\xee\x54\x64\x50\x8d\x0c\xc5\xc3\x91\xda\xbe\x9f\x2e\x16\xe1\x11\x9d\xc7\x8b\xb8\xa0\x11\x7a\x5e\x81\x92\x8a\xaa\x32\x4a\x9d\xbd\xce\xc1\x5f\xff\xa6\xbe\xd3\xa2\x16\x9e\x8e\xb2\xf2\x48\xd9\x2a\xe9\x90\xfe\xf0\x23\x59\x25\x71\xc1\x9f\x50\x5c\xd1\xb4\x43\x40\x8b\x92\xd3\xb7\xd8\x62\x74\x78\xaa\xbc\x09\xaa\x75\xd4\x96\x3d\xa5\xe8\xb2\xf1\x07\x27\xe9\x95\x1e\x38\x3b\xe4\xb5\x84\xe1\x12\xda\x69\xe8\x41\x51\x14\xbe\x3e\x74\xf3\x99\x5e\xe5\xec\x6f\x0e\xd4\xa6\x8b\x4a\xd1\xb5\x4a\xc1\x7b\xd2\xf6\xf1\xb4\x6a\xd6\xfd\x10\x42\x3c\x2c\x67\x43\xc8\x86\x21\x3d\x15\xb0\x58\x3e\xb9\x92\xfd\x56\xc4\xa8\x8b\x02\x06\xaf\x06\xfd\xf1\x45\x98\x5f\x1f\x21\xf1\x71\xe1\x88\x65\x97\x30\xc4\xfc\x9c\xd9\x61\xcd\xd5\xe8\x94\xbe\x24\x33\x6a\x5f\x1d\x51\x29\x07\xc7\x30\xd3\xec\x43\x01\x56\x3c\xae\xfd\xee\x5e\x3b\xb7\x44\x74\x7c\x91\x5e\xd3\xa4\xd1\xe6\x7a\xed\xad\x70\x1b\x3d\x2e\x8c\xe1\xb8\x80\xbf\x37\xb9\x66\x2d\xd8\xb2\xc7\xe9\x2a\x79\x18\xd8\xce\x8c\x9a\xde\x2a\x11\xc9\x67\x06\xa8\x91\x6d\x2d\xc1\xd5\xe7\x46\x93\xd2\x4d\x12\xa0\xf8\xdd\x00\x41\x8e\x73\x15\xa7\xbb\xd0\x7c\xb4\xc7\x53\x3e\x56\x43\xed\x96\x43\x92\xba\xd7\xf4\x7c\x1d\x2e\xb3\xf4\xbe\x75\x96\x38\x5c\xe6\xe3\x45\x78\xa5\x40\xb2\xaf\x6e\xd0\x87\x07\x14\x7b\xda\x65\xba\x2c\x89\xba\x87\x59\x16\xde\xaf\xd7\x16\x1c\x11\xea\x2e\x89\x1c\xe1\x92\x44\x25\x97\x07\xf3\xed\x3a\x80\x76\xce\xfc\x76\xb6\x58\xda\x74\xa1\x92\xc4\x70\xac\xd7\x9d\x87\x07\x3a\xcf\xe9\x7a\x0d\xbf\x93\xa8\xa6\x15\x0c\x55\xf6\x07\x03\xf5\x10\xe8\x43\xf1\xd5\xc5\x18\xec\x17\x95\x42\x42\x55\xea\x79\x22\x05\x2c\x62\x13\x93\x80\x1d\x37\xa8\x3b\x1d\x8d\xd7\x8e\xa8\xcf\x47\x5c\xd0\x5f\xae\xaa\xa5\xa3\x18\xcf\x7d\xb7\xf1\x2c\xe5\xc2\xb2\xa0\x36\x6a\xee\x67\x3b\xb1\x1f\x3c\x1d\xbb\x6f\x6b\x43\x01\x39\x1c\x0e\xa5\x94\xa2\x2a\xae\x11\x69\x5c\xe7\x87\xfd\x5f\x04\x34\x4d\x6e\xc4\x77\x2f\x34\xa8\x92\xdf\x46\x83\xf7\xc7\xe7\x67\x6a\x1b\xe5\xa9\xaf\xa5\xe2\x41\xd1\x7b\x10\x62\x3e\x89\x5c\x84\x95\x21\x11\xe7\xfc\x33\xd9\x45\x81\xe1\xad\x82\xc0\x0d\x46\xb8\x85\xc0\x1e\x84\xef\x54\x0a\x0b\xff\xe5\x92\x90\x3a\x01\xae\xec\xe3\x06\x03\xea\x9e\xc4\xc9\xf5\xa7\x30\xcb\x7d\x44\x5a\x34\x36\x50\xe7\xa7\x21\x38\x39\x7f\xff\xdb\xfb\xd1\xf9\xc7\xa1\xcf\x95\x70\x4e\xe2\x70\x74\xde\x1f\x8c\xc7\xb6\xce\xb3\x80\x5d\x02\xf8\x29\x9d\xaf\x16\x8e\xdc\x86\xc9\x16\xda\x3d\x4d\x21\x20\x44\xaf\x56\x34\xf1\x31\x84\x7b\x09\xf4\x77\xd2\xfd\x90\x82\x3b\x11\xec\xde\x84\xd9\x2e\x38\x6a\xbb\x51\x0a\x51\x7a\xd6\xcd\xe1\x97\x7f\xca\x71\x10\xac\xe1\x7a\xdd\x83\x4f\xfd\x14\xfa\x04\x3f\x26\xf3\x08\x22\xe7\x28\x2a\x21\x2f\x42\x4f\xf4\xbe\x7b\xc3\x87\xb1\x6b\x67\x05\x0c\xe3\xba\x8b\x5a\x93\x71\x15\x35\xac\x87\x38\x57\x02\x41\x25\xb1\x46\xf0\xfc\xef\x48\xb9\xdb\xce\xe8\x3a\x4b\xb9\xa3\x49\x6c\x60\x87\x72\x0e\x06\x77\x45\x16\x22\xb5\xcd\xb3\xeb\x58\xc5\x65\xe3\xd3\x70\xe9\x9d\x6a\xdf\x1c\x62\x43\xd5\x10\x8b\xf5\xe1\x66\x0f\xda\xe2\xe5\x61\x14\x81\x13\x9c\xcb\x06\x72\x0d\xb9\x8d\xd4\xc6\x4b\xeb\x89\x9c\x94\xbe\xab\x8f\x8f\x4f\xc3\x8e\xd9\x7f\x65\x57\xa0\x76\xa6\xba\x08\xec\x5f\x7a\xa6\xa8\xf7\x50\xd6\xfd\xeb\xc2\x6f\xba\xb0\x1b\x78\xdc\x7d\x2b\xb7\xf0\xd6\x6b\x9c\x53\x8f\x0e\x62\xc3\x40\xf0\x72\x45\x78\xa7\xce\xbb\x4c\x5e\x70\xfa\x70\xd3\x2e\x9e\xd3\x2b\x1a\x55\x4a\xb2\x7a\xe6\x20\xd5\x22\xd0\x4f\x80\x91\x35\x6d\x9b\x25\x75\x87\x3e\x4a\x8a\xa1\x8a\x3b\xca\xba\x26\xee\x84\x1b\x79\x05\x97\x8b\xab\x87\x49\xdf\x69\x3c\xe6\x9e\x2c\x44\x6e\xca\xcc\x6a\x23\x08\xd4\xc8\xb0\x9c\x7f\x99\x2c\x66\x9d\x79\xfc\x6a\xd7\x0c\x2a\xfe\x2e\x28\x18\x96\x0a\xd2\xc3\xc2\x11\x3e\xb2\x04\xba\x6a\xf6\x0d\x9a\x2d\xc3\x62\xc6\x5a\xb1\xf6\x5d\x4c\xd3\x3b\x96\x40\xc0\x50\x65\x72\x73\x07\x87\xc1\xe1\xcb\xed\x1e\x90\x69\x68\x36\x5a\xcd\x69\xc5\x3f\x0e\x22\x18\x2f\x46\x55\x35\x15\x4d\xb0\x47\x7c\xf8\xcd\x56\x62\x9e\x41\x5b\xcf\xf4\x10\xd0\x11\x3d\xb2\xb0\xb5\x7a\xde\x10\xc0\xda\x01\xd1\x7f\x30\x72\x55\xea\xad\x4a\x9d\x22\x9f\x19\xa0\x47\x34\x8f\x33\x1a\xf5\xd1\x97\x70\xfa\xd4\x9e\x74\x57\x2b\x9f\xfa\x65\x43\x68\x3b\x2f\x51\x47\xa7\x9d\x64\x30\x23\xf2\x32\x41\x5f\xa6\x38\x2b\x39\xb0\xb8\x66\xd4\x6c\x0d\x79\x69\x55\xa5\xce\x2c\x08\x53\x52\xed\x52\xb5\x3a\x1c\xa2\xe8\xcc\xc4\x81\x55\x1d\x52\x23\x38\xb6\xc3\x14\x67\xd1\x63\xa6\xac\x99\xf9\xd7\xbf\x9c\x16\x4a\x55\x54\x6c\x59\xd7\xda\x85\x9e\x63\xed\xab\x96\xeb\x22\xcc\xae\x68\xe1\x98\x24\xcb\x64\xb9\x4c\x53\xad\x83\xa9\xc9\x40\x7b\xb5\xe9\xeb\xc7\x6d\x00\x9b\x73\x19\x4d\xb6\xd7\x61\x75\xdd\xdd\x19\x6b\xc4\x69\x68\xbd\x76\x4d\x45\xd9\x94\x5b\x75\xd6\xe2\xea\xf9\xe7\x92\x8b\x6a\x72\xf1\x7b\x91\xcf\x64\xcc\x00\x9b\xc0\xa7\xb4\x3b\x54\x9e\x2a\xc0\xb2\x97\x61\x06\xdd\xde\x21\xfc\x32\x8b\x93\x62\x4a\x02\x89\xfb\xbf\xa0\x5b\x0d\xa7\x99\xc7\xec\xaa\xee\x99\x92\xbc\x94\xa9\x2a\xb3\x0f\xa7\xef\xd4\x47\xf5\x3b\x8d\x27\x56\xe9\x90\xb7\x5a\xd7\x1c\x6a\x23\x5a\xa6\xc4\xac\xca\xb6\x47\x4d\x89\x7b\x3b\xc0\x3d\x1d\x65\x8d\x17\xc6\xca\xad\x99\x57\xad\x45\xd9\xde\x98\xc1\x4d\x78\xf8\x22\x55\x7a\x8f\xa1\x90\x79\xc8\x8f\x21\x0d\xd5\x1d\xa6\xf2\x94\xce\x46\x61\x12\xa5\x8b\x9c\x6c\xc7\x45\x1a\x56\xbd\xbc\xb2\x3c\xbd\xda\x81\x3c\x6a\xfa\xf5\xed\x0e\xdf\x4e\x80\x98\xe0\x53\x53\xe5\x35\x4b\x47\xb9\xf6\x4a\x1e\x1b\xac\x35\xf8\x58\xef\x01\x1b\x6d\xab\x1d\x24\x65\x53\xc6\x74\x6a\x70\xde\x34\xd5\x8c\xed\x80\x99\x67\x63\xee\x0f\x7c\xd5\x6b\x78\x5e\x44\x9c\xe5\xc7\x4d\x9c\x7d\x0f\x76\xcd\x0b\x14\xa3\x0e\x8c\xee\x9e\x47\xc2\xe5\x47\xe9\x68\xbd\x00\xe1\xaa\xd8\x74\x55\x87\x83\xb5\x29\x32\xcc\xff\x85\x55\x1a\xf4\x19\xe4\xdd\xdc\x7c\x7b\x01\x89\x77\x08\x9c\xaf\x06\xf7\x89\x9c\xd4\xf7\x09\x79\x19\xaa\xd6\x93\x52\xc2\xed\x8c\x0a\x02\x06\xa6\x6f\xbf\x5a\x6e\x1b\x69\xe1\x50\xef\x48\x52\x2d\x67\x5a\xaf\x3f\x3e\x4e\xae\x44\xe6\xe7\xb3\x2b\x70\xf4\xac\x39\x01\x65\xfa\x97\xcc\x6f\xc4\xe4\x15\xab\x98\x71\x04\x89\xfd\x38\xca\x8e\x91\xdf\xc1\x5e\x97\xfd\xb7\xbb\x67\xef\x2a\xf9\x3c\xbe\xaa\xb5\x52\xf0\x23\x4a\x4a\x1d\x49\x03\x8f\x43\x17\x1c\x2f\xd5\xea\x41\xac\x80\xb4\x8a\xc3\xde\x65\xe9\x02\x07\xee\x5a\xc9\x16\xf0\x45\xea\x03\xd5\x73\x13\x8d\xc1\x7a\xa3\xf3\xa7\x46\xa5\x9f\x96\x93\xe3\xc8\x64\x85\x55\x9f\xd1\xf1\x2e\x00\x57\xb5\x00\x17\xda\x79\x98\x17\xf1\xa4\x5a\xfb\x30\xf3\xb8\xd9\x5d\xa9\x82\x4a\x88\x1f\x67\x1a\xb4\xfc\x48\x8b\xd5\x59\x8d\xdb\xb7\x6a\xaa\xa4\xf7\x78\x32\xa3\x80\x22\x88\xab\x33\x5b\x5a\x1e\x80\xbf\x0f\x7a\x0a\x84\x9e\x0e\xa8\xca\xdf\xf5\xbd\x51\x59\x7a\xde\xf1\x38\xec\x46\x85\xba\xe5\xef\x9b\x80\x86\x53\xaf\xc2\x3b\x17\x40\x45\xb9\x41\x58\x79\x66\xa6\xa3\x8e\xc9\x2f\x4d\x56\x64\xed\x1d\xf2\xb1\x0b\x9b\x3d\x4e\xe7\xd8\xec\x11\xe9\xe2\x8e\xa2\x93\x50\x56\x7c\x78\x94\x41\xdc\x04\x42\xc6\x2b\x32\x39\x19\x42\x96\xe0\x0b\x9a\x9c\x8e\x5a\xef\xf6\xb7\x3d\x4d\x99\x55\x78\xd4\x62\x26\x20\x3e\x9a\x53\xa5\x48\x0e\x85\x4c\x79\x64\x46\x81\x41\x3f\x4b\xf3\xfc\xd7\x34\xa1\xb2\xcb\xea\xd5\x07\x1a\xce\x8b\x59\x7f\x46\x27\xd7\x66\x02\x87\xbf\xba\xbf\x98\x81\x0a\x9d\xa5\x73\x96\xe0\x3c\xd0\x05\x8a\x31\xf1\x86\x15\xdc\x32\x22\x78\x13\xf9\xd4\x8a\xf1\x79\xb0\xec\x52\xfa\x56\xa4\xae\xa0\x93\x0a\x0d\xd0\xf5\xbc\x12\xea\x5b\x98\xd2\xd1\x10\xa8\x78\xce\xcd\x99\x6e\x56\x7b\xc4\xa4\xa0\xae\xe2\x1c\x39\x1b\x9d\xff\xbc\xa5\x32\x03\x1a\xf0\xc7\x64\xe6\xe4\x66\x15\xee\x2a\x73\x22\x53\x0d\xcf\x69\xb7\x34\xe3\xce\xd9\xd9\x75\xe6\x03\x54\xf3\xa1\xfb\x4b\x46\x65\x3b\x6b\xdf\xde\xbe\xe9\xa8\x5d\xb5\x19\x96\xeb\xfe\xc8\x00\xae\x53\x95\xd5\x63\xf1\xbc\x63\x4f\xc0\x6b\x3d\x55\x43\xd0\xda\x44\xba\xca\xf6\x35\xce\x99\x00\x6e\xce\x55\x78\x78\xc7\xae\x14\xe0\x86\xf1\xa2\x3d\x72\x64\x89\xc2\x2b\x69\x64\x5f\x6e\x2e\xec\x54\x94\x4f\x75\xd7\x81\x3e\x95\x0c\x6b\x33\xe7\xab\x59\x81\xfe\xbc\xfe\x8b\xa7\xac\xbf\xe5\x02\xb6\x17\xec\xdd\x7d\xdd\xaa\x75\x54\x2a\xe9\xa7\x05\xb8\xc1\xd1\xf0\x38\xcf\x4f\xb0\x46\xd2\x5f\xd2\xc0\x95\x57\xae\x03\x13\x45\x91\xc5\x97\xb8\xd9\x82\x03\x76\xef\xdc\x96\xc5\x53\x4d\x64\x30\xe0\x32\xd4\x44\x73\xe5\xa8\x36\x70\x6d\xc5\xa9\xeb\x87\x13\xf2\x1c\x2b\xc8\x3a\xb9\xf0\xb8\x3d\x9a\xcd\xe4\xe7\xe4\x6d\x3f\x4d\xaf\x63\x3a\x06\x5f\xf5\x1a\x22\xca\x3c\x2f\xfd\x07\x1c\x95\x3e\xbb\xe1\x94\xed\x03\x60\x59\x9c\x86\xc3\x88\x6f\x79\xfc\xdb\x22\xec\xf5\x05\x53\xe2\x94\x7e\xa9\x2d\x88\x2e\xdc\xae\x63\xfe\xda\xaa\x2a\xb7\x39\x1b\x7d\xe2\xb5\xbb\x9d\x01\x54\x71\xae\x9c\x24\x25\x30\x68\x0a\xd1\xcd\xd2\x71\xa5\x0a\xb5\x0a\xc3\xf5\x1d\x44\xb1\x13\xa0\x06\x17\x35\x89\x20\x01\xec\xcd\xfd\xb4\x24\xd0\x91\x41\xd8\x8c\xc4\xc6\xf8\xe7\xd3\x81\x2f\x02\xf2\x85\xf1\xd5\x5e\x98\x37\x46\x57\x64\x68\x33\xd7\x5e\x7c\xa6\xc5\xce\x34\x44\xea\x82\xba\xf8\xde\xcc\x6c\xf3\xd1\x8f\xab\x63\x0f\xda\xaa\xda\x28\xfc\x69\x19\xfc\xd4\x86\x3e\x5f\x7d\xe7\xf6\x74\x32\xbd\x93\xf5\x3c\x19\x96\x52\x31\x6a\xa7\xdb\x37\xca\x93\x18\xd2\x56\xfa\x82\xb6\xed\x72\x66\x20\xec\x9c\x83\x70\xc4\x99\xaf\x96\x7a\x5e\xa8\x09\x91\xc6\xc8\x4d\x49\x95\xa8\xe9\x17\x9c\x82\xff\x5c\x72\xc2\x48\x19\xb6\x65\xa6\x63\xc7\x51\xb2\x65\xe3\x75\x2d\x70\x34\x0b\x8e\xd8\x33\xe0\x07\xf9\xca\x75\x25\x3b\x99\xa6\xd9\x6d\x98\x45\x6c\xbe\x58\x10\xc7\xc4\x06\x77\xec\x7b\xb5\x3b\xa5\x82\x6e\xde\xa6\xdc\x96\x58\x1b\x0b\x53\xd5\x3e\x6e\x9c\xba\x4a\x71\x0a\x8c\xf2\xda\x3c\x39\x5c\xb3\x08\x1b\x89\xdd\x94\xe3\x0a\x7b\x9a\x99\x2e\x07\xb2\xbf\x77\xf0\x63\xed\x08\x9e\x43\xfc\xbe\xcf\xf0\xdc\x4a\xef\x0d\xca\x61\x94\x2e\xc8\x5f\xf7\x0c\xc9\xd4\xeb\x68\x1a\x4c\x4c\x63\x2d\x8c\xb9\x33\xde\xaa\x1a\xe6\x4f\x61\xb7\xcb\x2f\x74\x51\x27\x9c\x42\xf7\xf0\xac\x00\xeb\x99\x26\x6d\xb3\x02\xa5\x3f\xab\x30\xc9\x7f\xda\xb6\x49\xf3\x60\xcf\x2d\xcf\x97\x3f\xab\xea\x79\x4e\xa1\x34\x15\x16\x31\x6e\x04\xe3\x67\xeb\x62\xca\xb3\x3a\x38\x63\x3b\xf0\x17\x5a\x29\xa4\x5a\xbd\x27\x06\xd5\x0c\xe1\x35\x40\x26\x46\x49\xe9\x23\x47\x66\xd6\xf5\x80\xbb\x92\x0a\xef\x02\x7b\x5d\x8a\xaf\x42\x0b\x7c\x1f\x8b\x89\xad\x12\x50\x52\x14\x37\xdf\xed\x72\x9c\xa6\x55\x5d\x55\x3c\x58\xd0\x07\x90\x9f\xd3\xcb\xdc\x3e\xd1\x89\xc9\xc9\xa4\x3c\xaa\xed\x5b\xfe\xde\x33\xdd\xde\xb5\x2e\x21\x95\x52\x61\xed\xd4\xc2\x0e\x4e\x85\x72\xe6\x61\x07\x47\xa1\x0a\xaa\x75\x56\x7a\x26\x1e\x28\x30\x4f\x3c\x09\x5d\x7b\x0e\x5a\x49\x30\xee\xef\x69\x39\x61\xeb\xa0\x7a\xf0\x6b\xbc\x7c\x17\x33\x52\xcc\xd4\x6a\xf0\x25\xb1\x73\xab\x5b\xab\x9c\x92\x1c\x62\xf6\x49\xb1\xf5\x93\x79\x2f\x83\xf9\xfd\x26\xcc\x08\xbb\xbb\x83\x80\xe5\xa0\xbf\xaf\xe2\x8c\x6e\x6f\xb1\x07\x5b\xaf\xac\xc6\x08\x1c\xde\x6a\xa0\xf0\x75\x27\x8f\xae\x3d\xc0\x93\x79\xba\x8a\xca\x33\xec\xd0\x2e\xa1\xb7\x88\xa1\xdb\xc7\x17\x65\x45\xde\xb6\xbb\xf5\xef\x2b\x9a\xdd\xe7\xfc\x50\x8f\xd2\xa5\xf2\xd8\xd1\xad\x0b\x11\x2f\x4b\x01\x1c\x0f\xe6\x5b\xfc\xab\x8c\x6c\x7a\x64\xcb\x94\xa3\x2d\xb3\xc1\xba\xb1\x43\x51\x7b\xd3\x05\x3c\xdd\xb3\xf3\xa3\xc1\x6f\x17\x27\x78\x2a\xe6\xe7\x41\xff\xe2\xb7\x8f\x67\x87\x1f\x2f\x3e\x9c\x8f\x8e\x7f\x1d\x1c\x01\x39\x5b\x7b\xcd\x13\x44\xef\x96\x98\xf1\x90\xf2\x09\xad\xa6\x42\xf6\xb7\xe9\x0d\x4d\x8a\x0e\x99\xa4\xe0\x32\xdf\x15\xaf\xdc\xa3\x83\xb7\x39\x88\x68\x77\x9e\x5e\x6d\x6f\xe1\x0d\x3e\x83\xf1\x05\x19\x0d\xfa\x83\xe3\x4f\x83\x23\x18\x31\x79\x4d\x7e\x1e\x9f\x9f\x75\x39\x43\xe3\xe9\x3d\x47\xfb\xaa\x99\xb3\x0c\xbb\x36\xc1\xdd\x88\x85\x26\x97\x94\xb1\x34\xdf\xe6\x7c\xef\x28\x14\x67\x59\x87\x44\x61\x11\x7a\x88\xc5\x1f\x50\x2a\x08\x57\x03\x61\x0e\x0b\xa1\x2d\x6a\xe5\x87\x35\xcf\x58\xd7\x22\x43\x19\x89\xc2\x7c\x76\x99\x82\x29\xf2\x62\x92\x90\xcb\x30\xcf\x6f\xd3\x16\x80\x22\xfb\x00\x53\x86\x43\xee\x72\xa6\x7c\xde\xfb\xda\x15\x67\xf0\x5b\xf4\x24\xef\x41\xb4\x71\x54\x77\x24\xb6\x9a\x28\xf9\x03\x73\xb5\x8d\xa8\xe3\x37\x7b\x3f\x49\x02\xbb\x73\x76\x2b\xd3\x3f\xe2\x9f\x48\xfc\xfa\x75\x03\xe3\xf1\x07\xa7\x48\xb4\xfd\x1c\xcb\xe1\xfc\x42\xef\xc9\x1b\x90\xe9\x23\xc9\xc8\xad\x16\x98\xf0\xa7\xe4\x3c\x0c\xd2\xc2\xca\xcc\x6b\x2d\xa3\xf0\xc7\xcc\xb7\x07\xb5\x2f\x9b\xf8\xc3\x78\x53\xf1\xfe\x51\xec\xa9\x9a\xe3\x58\xca\xc9\x92\x4c\x1a\x0a\x19\x6a\xcb\x23\x29\x73\xc0\x22\x0f\xe6\x3f\x9b\x51\x4c\xc0\x59\x12\xc2\xab\x58\xd5\x9f\x19\x38\xd4\xbd\x6a\xaa\x3b\x4d\xf0\xa8\xf4\x7a\xe4\xc7\x1f\x7f\x68\x86\x04\x87\x0c\xf4\xd8\x6e\xb8\x5c\xe6\xbb\xa8\xce\x98\xfe\xea\x62\x1d\xce\x6b\x78\x5c\x9e\xfa\x57\xde\x89\x67\xec\x7d\xb6\x4a\xb6\x1a\xfb\x00\x0e\xcf\xd2\x08\x7a\x19\x9e\x8f\x2f\x9a\xc1\x67\x34\x8c\x60\x7e\x7a\xed\xe6\x76\xeb\x70\x32\xa1\xcb\x62\x0b\xd0\x03\xd1\x73\xdc\xd3\x00\xae\xee\x7e\xcb\xd3\x16\x94\x31\x04\x58\xc0\x0b\xe3\xda\x41\xa7\xca\x44\x73\xb7\x73\x7b\x7b\xbb\x83\x2a\x7a\x67\x95\x81\x24\xe3\x15\x7f\x51\x4b\xbc\x1f\x73\x9a\xed\x1c\x5e\x01\x6a\xc4\x3a\x81\xf6\xbb\x96\x41\x34\x1b\xad\x1b\x51\x87\x2b\x36\x61\xfc\x66\x9d\x1e\xce\x8a\x94\xee\x5a\xe1\xdc\x48\xcd\xa1\x74\x82\xc3\x00\x92\xc9\x7c\x99\x2e\x3a\x0f\xe0\x96\x6e\x0b\x89\x55\x4c\x12\xb8\xb5\x2d\xd6\x20\xe2\xbb\x4c\xa3\x7b\x34\xda\xb6\xcd\x36\xa1\x35\x9b\x8b\x17\x2c\xad\xf2\xde\x56\x07\xef\x91\xe8\xe6\xec\x1b\x7a\x79\x7e\x93\xe5\x44\xf3\x41\xc8\x14\xe0\x31\xec\x35\xa2\x15\x12\xe7\xb0\xda\x26\x52\x46\x04\x2d\x06\x28\x08\x80\x60\x7b\x6b\x55\x4c\xff\xee\xf0\xa3\x5c\xed\x80\x5d\x5b\x68\x86\xb6\x14\xfe\x4d\x66\xab\xe4\x1a\x38\xc8\xd9\xf3\xfa\x0d\x61\x0f\xc8\xba\x3d\x46\x88\x39\x54\x84\x2d\x35\xa2\xca\x1c\xec\xba\xb9\x3f\xd1\x0a\x7d\xa5\x6e\x94\x26\xd4\xe1\x75\x9a\xe0\x0d\xa3\x70\xbc\xae\x03\x07\x21\xe4\x03\xce\xb2\x34\xdb\xd2\xdd\xa2\xb4\xc9\xe1\x31\xc7\xbc\x35\xc0\x46\x28\x0e\xbc\x75\x1b\x69\x62\x43\x9f\x86\xf1\x7c\xbb\x45\x9b\x0d\x07\xc7\xfc\x15\x50\xf2\x47\x20\x1f\xb0\x4a\x14\x27\x5d\x11\xd5\x86\x11\x82\x4e\x60\x27\xf9\xb7\x48\x4f\x68\x6a\xf1\xe0\x99\x27\xe1\x16\x62\x65\xba\x2d\xa9\xad\x67\x03\xc2\x83\x84\xd6\x08\x8b\xd3\xa6\x3a\xa8\x02\x0d\xa6\x3c\xd0\xaf\xf2\x6a\x9d\x5e\x6e\x08\xb7\x31\xec\x1d\xcf\xc0\x76\x8a\xb8\x97\xe5\x67\xbc\x59\x37\x64\x71\xde\xeb\x69\x49\x9c\xa6\x6d\xa1\x80\x6f\x33\x26\x57\x32\xb2\xce\x27\x33\x1a\x19\x69\xa0\xb1\x78\x36\xb8\x5b\xe2\xce\x83\xd8\xa9\x60\xc4\x89\x37\x46\x52\x97\x67\x6e\xac\x0d\xe6\xa0\x4c\x97\x58\x11\xb8\x9a\x57\xf0\x1d\x2f\x3b\x8e\x1c\x04\x8b\x32\x25\x1d\x30\x59\x8a\x5a\x9f\x2f\x78\x03\xfd\x17\xf8\xf8\x25\x10\x29\x04\xd1\xec\x0b\xf4\xf2\x45\x06\x77\x15\x80\xd8\x68\x2e\x01\x84\xc0\x56\x00\xe2\x72\x0a\x06\xa0\x26\x64\x1c\x9b\x9a\x8e\xf9\xe3\x99\x91\x21\xcd\x16\x71\x9e\xbb\x52\x28\xc4\xcc\xa1\x28\xb0\xae\x29\x25\xda\x9c\x8a\x8b\xfd\x10\x09\xbf\x70\xae\x77\x0c\x46\xf9\x9a\xba\x2e\xd7\xd3\x52\x2e\xe4\x91\x93\xa2\xdc\xe2\x88\x9d\xb2\x45\x9e\x1b\xf7\x36\xaa\x72\xc4\x32\x4e\x0c\x8d\xb7\x1c\xdf\x12\x77\xa5\xe3\xba\x55\xe5\xd9\x97\x75\x5d\xc0\xcf\x6f\x56\xe2\x57\xfe\x7d\x08\xf3\x41\x5f\xbd\x07\x96\x11\x75\x9e\x69\xbb\x6b\xce\xeb\xb9\x9c\xd7\x07\x02\xb1\xab\x7c\x87\x86\x79\xc1\x2e\xe0\x52\x0b\x0f\x37\xc4\x71\x0b\x7a\x60\xe7\xe0\x09\x38\xe8\x8a\xe3\x60\x74\x08\x14\x28\xa5\xde\x0b\xa2\x5c\xf7\xdb\x55\xac\xc2\x37\x8e\x1b\xab\x1d\x67\x52\x47\x0a\x98\x58\x07\x5a\x75\xae\xc6\x78\x01\xe0\x3c\xb3\xea\xd6\x5d\x4f\x3f\xab\xaa\xdc\xbc\xdd\xbc\x39\xfe\x9d\x94\xb8\x66\xbe\x19\xd7\x83\x49\xae\xf1\xbd\x84\xb6\x8c\xb0\x6e\x10\x63\x14\x41\x4c\x81\x27\x65\x01\x95\xa4\xa5\x61\x5a\x36\xe9\xc8\x2a\x86\xb0\xf1\x76\x2c\xa6\x89\xf5\xe8\xe5\x8a\x71\x9d\xa1\x76\xaa\xd5\x7d\x97\x99\x79\xc3\xad\x67\xfa\x5b\xdd\x6c\xeb\xbd\xb0\xd6\xb8\x32\xb4\xba\x3d\x56\x7b\x4e\xac\xbb\x64\xb5\xd7\x46\x41\x53\xc3\x75\xb7\xfa\x55\xb7\x66\x3f\xca\xc5\xb7\xc6\x2b\xcc\x16\x4e\x4c\x6d\x6a\xc0\xf8\x2f\x1d\xf0\x6c\x87\x3b\xef\x91\x95\x7a\xb5\xb2\x30\xd6\x05\xb8\xee\x2a\x39\xcd\xb8\xeb\x85\x4e\xea\x7c\x77\x6c\x30\xff\x7d\xc4\xcf\x7f\xd5\xb0\x77\x92\xd9\x5b\x3a\x39\xe8\x1d\x42\x0c\x9b\x66\xf1\x1f\xd4\x59\x68\x61\xa3\x14\xcd\x8e\x44\x7e\xf4\x2f\x1e\x10\xbe\xdb\x86\xff\x7a\xc1\x65\x79\x7e\xe1\x88\x72\x7d\x41\xb3\xb2\xd2\x03\xeb\x2d\x9c\x35\x3d\x9b\x60\x93\x2e\xd7\x26\x6d\x37\xa7\x7e\x64\xd2\xfe\xcf\xb8\x98\x3d\x92\xf6\x91\x41\xb9\xd5\xd0\x55\x37\xa8\x5c\x56\xfc\x99\x5d\x57\x6c\x95\xfc\xe9\x2b\xd5\xb9\x22\x54\x73\xd8\xa0\xdc\xd5\x5b\x57\xed\x4b\x4a\x75\x05\x36\xfe\xa1\xd7\x13\x37\x10\x0b\x0d\x76\x44\xe7\x14\x65\xaf\x2c\x08\x04\xfa\xf1\x6c\x76\x83\x86\x63\xff\x62\x08\x66\x81\x32\xbe\x41\x6d\x96\x3c\x81\x7f\x6d\xdc\x10\xf3\x20\xef\xde\x0b\x72\x08\xd3\xe8\xa2\xdc\xca\xc4\x47\xe2\xd2\x63\xa2\x7b\x12\x02\x1e\xef\xa7\xef\x78\xed\x8e\x6a\x09\x5d\x6c\x53\xb8\xf6\xef\x00\x00\x00\xff\xff\x67\x19\xc7\x7c\xf6\x69\x00\x00")

func templatesAppTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAppTmpl,
		"templates/app.tmpl",
	)
}

func templatesAppTmpl() (*asset, error) {
	bytes, err := templatesAppTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/app.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/app.tmpl": templatesAppTmpl,
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
		"app.tmpl": &bintree{templatesAppTmpl, map[string]*bintree{}},
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

