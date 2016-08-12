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

var _templatesAppTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x3d\x6b\x73\xdb\x46\x92\xdf\xf3\x2b\xa6\x50\xb9\x92\xbc\xa6\xa8\x47\x72\x7b\x7b\xcc\x79\xab\x64\x8a\xb6\x95\xe8\xc1\x23\x65\x6f\x5d\x6c\x55\x0a\x02\x87\x22\x22\x12\x60\x00\x50\x8f\x68\xf9\xdf\xaf\x7b\x1e\xc0\x3c\x01\x50\x8f\x5d\xed\x5a\x22\x31\x3d\x3d\x3d\x3d\x3d\xfd\x9a\xc6\xe4\xf1\x91\x4c\xe8\x34\x4e\x28\x09\xc2\xe5\x32\x20\xeb\xf5\x77\x84\xc0\xc3\xef\xe1\x1b\xe9\xbd\x23\xdd\x43\xf8\x5b\x3e\x5c\x84\x49\x3c\xa5\x79\xc1\x5a\x4e\xe5\x17\xde\x0c\xff\x08\x09\x0e\xff\x31\xbe\xa0\x8b\xe5\x3c\x2c\xe8\x87\x34\x5b\x84\xc5\x17\x9a\xe5\x71\x9a\x04\xa4\x47\x82\x83\xbd\xfd\xbd\x9d\xbd\xff\x86\xff\x07\x1d\x0e\xde\x4f\x93\x49\x5c\x40\x7b\x1e\xf4\x04\x0a\x36\x52\x21\x70\x90\xe0\x2a\x9c\x87\x49\x44\xb3\x9d\xa8\x02\x35\xc7\xb6\x3a\x2d\xb3\x34\xa2\x79\xbe\x51\x9f\x8c\x5e\xc7\x79\x91\x3d\x34\x75\x0a\x8e\x93\x82\x66\x49\x38\x47\x8a\x49\xf0\x21\xe9\xf5\x06\x7f\xac\xc2\x39\xce\xe0\x2b\x3e\x19\xd1\x29\x7c\xac\xc0\xc8\xba\x43\x82\xff\xa3\x80\xed\x12\x3e\x4a\x2c\xc3\x2c\xbe\x85\x71\x1b\x90\x48\x28\x37\x8e\xf7\xc0\x9a\x9b\x31\x8d\x56\x59\x5c\x3c\x7c\xcc\xd2\xd5\x12\xd9\xfc\xa8\xa2\x83\xef\x5f\x1f\x19\x36\x5c\x00\x1d\x16\x71\x06\x97\x7c\x5e\x02\x69\x30\x0c\xb3\x70\x41\x81\x72\xd6\xb5\x7e\x45\x96\x08\xbb\xc1\x6a\x38\xe1\xe5\x5c\xfa\xf3\x55\x0e\xc3\x2a\x62\x00\x0f\x2f\x1e\x96\x94\x13\x5e\x64\x71\x72\x1d\x74\xaa\xa6\x23\x3a\x0d\x57\xf3\x82\xb5\xea\xcf\xf3\x28\x8b\x97\x85\x94\xb9\x40\x34\x55\x5c\x3b\xa2\xcb\x79\xfa\xb0\xa0\x49\x71\x1a\xde\xc7\x8b\xd5\xc2\x31\x26\x74\x3c\x5b\x2d\xae\x80\x1e\xc7\x90\x4c\x92\xf7\x7c\x83\x42\xab\xc0\x4b\x96\x34\x8b\x60\x98\xf0\x9a\x92\x74\x4a\x04\x1b\x68\x4e\x8a\x94\xdc\x50\xba\x24\xd9\x2a\x49\x60\x5a\xe4\x6e\x16\xcf\x29\xec\x43\xa4\x0b\xa7\x59\x47\x72\x9c\x3c\x91\xe4\xfd\x7a\x92\x39\xde\x97\x23\x79\x90\xdc\xc6\x59\x9a\x20\xcd\x6e\x62\xfd\x4b\x5a\xb3\xa2\xce\x05\x55\x37\x64\xbb\x71\x34\x84\xe7\xc9\xfc\x81\x84\xf3\x79\x7a\x47\xc2\x08\xa7\x8b\x93\x2d\x66\x71\x4e\x50\x07\x4e\xb3\x74\x41\xe2\x24\x8f\x27\x14\x1e\x52\xf2\x65\xd8\xf7\xd0\x7c\x96\xaa\x0d\x87\x88\x90\x4e\xbe\x84\xf3\x15\xe5\xbb\x9a\xed\xdf\x0e\x83\x23\x97\xd6\x24\x7e\xa1\x0f\xaf\xcd\x27\x45\xe5\x3c\x81\x4d\x9f\x73\x4a\xc6\xab\xab\x84\x16\xb9\x40\x84\x7c\xca\x97\x34\x8a\xa7\x0f\xc8\x96\x1d\xc6\xa3\x79\x1a\x4e\x88\x54\x11\x84\x26\x93\x65\x1a\x27\x45\xfe\x2a\x3c\x1b\xd1\x39\x0d\x73\xd7\x84\x5e\x5a\x67\x8c\xe8\x32\xcd\xe3\x22\xcd\x5c\x8b\xf4\xbc\xc1\xc6\xe9\x0a\xb6\x1c\x89\x52\x60\x5e\x56\x0d\x63\x91\xa0\xeb\xee\x97\xa6\xe2\x02\x44\xfb\x44\x5b\xba\x5c\x8c\x47\xae\x71\x40\x32\x4d\xb3\x72\x53\x38\x88\xe3\x82\xe1\x21\xeb\x04\x0c\xeb\xff\x80\x67\x00\x56\xa9\x7f\xd0\xeb\x71\xe0\x5e\xef\x78\xf2\xf7\xa7\x90\x0a\xa2\x46\x72\x3e\x5e\x3b\xaa\xfc\x72\xff\x3a\xc4\x2d\xc5\xf6\x68\x47\xa4\x74\x90\x34\xea\x8c\xbd\xb7\x3d\x1a\xfc\xef\xe7\xe3\xd1\xe0\xe8\x0d\x39\x09\x17\x57\x93\x90\xf4\xc1\x5a\xa6\x8b\x8b\x74\x19\x47\xe4\x53\x98\x4c\xe6\xb0\x62\x62\x3b\x10\x89\x51\x21\x13\xd4\xfb\x09\x4d\xae\x8b\x19\x23\x72\x5f\x6d\x32\x14\x80\x4d\x1f\x28\x3c\x37\xe7\x2a\xa6\x01\x0c\x72\xec\xa9\x0c\x6b\x60\xd0\xb0\xdf\x3f\x3e\x1a\xbd\xb8\xc8\xe3\xc8\x88\xd8\x3d\xbc\xe6\x15\x9d\x42\x0b\x8c\xa2\xca\x77\x30\x4c\xb3\x62\x98\xa5\x45\x1a\xa5\x86\xe5\x99\x15\xc5\x92\xfb\x75\x28\x5b\x34\xa1\x99\x02\x17\x7c\xba\xb8\x18\xa2\x4a\x3b\x4e\xf2\x02\x77\x9a\xab\x8d\xed\x75\xea\x83\x18\x07\x15\x77\xc4\x70\x79\xfd\x78\xe3\x67\x0f\xa8\x8d\x58\x44\x35\xf3\xbb\xe8\x7b\xa7\x27\x9a\xfc\x83\x8d\xc7\x27\xe6\x50\xf3\x9a\xa9\x21\xf8\xf3\x86\x22\x6b\xe7\x7a\x8f\x68\xce\xb4\xb2\xb6\xe0\xca\x96\x1b\xa5\x73\x8f\x19\x65\x7b\xe2\xf8\xf0\xb4\xd7\x63\x30\xca\x4c\x60\x70\x70\xae\x8a\x98\xea\x5a\x12\xcd\x5e\x9e\xaf\x16\x14\xe1\x87\xe9\x3c\x8e\x1e\x8e\xd2\x68\x65\xf9\x4d\x7c\x2b\x94\xba\x02\x63\xa9\x83\x1d\x08\xa7\xf6\xff\x4b\x19\x84\x01\x8d\x0b\x50\x3e\xa2\xff\x57\xad\x89\x18\xf8\x18\xf8\x60\x3a\xa5\x11\x33\xc6\xcc\xfc\x1a\xd8\x04\xe9\x71\x12\xc5\x4b\x19\xf2\x8c\x69\x76\x1b\x47\x94\x1b\xe8\x39\xd3\x47\xdd\x70\x11\xfe\x99\x26\xe1\x5d\xde\x8d\xd2\x85\x16\xa5\xa8\x13\x8d\x84\x42\x83\x7e\x79\x91\xf7\xaa\x89\x57\xd6\x5d\xfe\xac\xb5\xef\x6a\xab\x86\x19\x02\x16\x50\x6a\x40\xfc\x6e\xa0\x3f\x46\x4e\x72\x5e\xeb\x3c\x30\x39\xc0\x21\x1f\xce\x20\xea\x61\x3c\x98\x2c\xc0\x13\x86\x38\x30\x04\x2b\x6c\xf1\x22\x68\x58\x20\x06\xd3\x66\x91\x18\xa0\xb6\x50\xc8\x58\x6b\x29\x14\x96\x05\x7f\xc1\xaf\x52\x30\xf9\x03\xb2\x6e\x60\x9b\xfa\xad\x82\x5c\x5b\x2a\x56\x11\xed\x1a\xb1\xe6\xa6\xa7\xd7\xfb\xb0\x4a\x38\x55\xad\xa4\xbb\x0f\x8e\x8d\x2d\xc9\xe3\x1f\xde\xaf\xa2\x1b\x5a\x54\xf1\xef\xcf\xe0\x27\x72\xd1\xd8\x81\x99\xc2\x1f\x88\xc3\x6f\xd3\x7b\xf8\x5c\x85\xc3\x8c\x8c\x11\x04\xea\xa8\xc2\x61\xf2\xb6\x9c\x01\x62\xe1\x49\x9b\x58\x39\xd2\x8c\xdb\xc8\x5d\x0d\x6d\x99\xa3\xc0\x88\x78\x97\x4b\xf4\xee\x94\xa5\x2f\xe0\x71\xf7\xcf\x78\x19\xf0\xb1\xbc\x52\x28\x4c\x30\x22\x8b\x93\x09\xbd\xef\xd2\x7b\x11\x93\x68\x60\xa7\x74\x01\xbe\xdd\x38\xfe\x93\x31\x75\xff\xe0\x6f\x7a\xb3\x54\x2b\x9c\xf4\x8f\xb4\x38\x2c\xb8\x6c\x58\xba\x07\x25\x23\x4b\xac\x7d\x16\x8c\x56\x49\x11\x73\x49\x4e\x80\xef\xbf\xe7\xfa\x00\x17\xd0\x96\xae\x98\x84\xfd\xb0\x17\xf8\x05\xc2\x1d\xef\x67\xa5\x56\x24\x5d\x4f\xa8\x1f\x41\xc8\xf7\x7b\x7a\xd5\x06\x54\x66\x05\x54\xd0\x96\x89\x84\x9c\x2b\xa0\x1a\xe4\x65\x32\x67\x03\xec\xe0\xe6\x16\x8e\x59\xba\x40\xa5\x6f\x1c\x78\x50\xe5\x05\x4f\xda\xe8\x56\xe5\x7c\x55\x2c\x57\x45\x73\xa6\x2b\x15\x70\xa4\x5b\xcf\x86\x0a\xae\x6d\x6a\xab\x7d\x0f\xce\x8b\x0a\x5e\x4d\xd5\x8c\x69\x51\x18\xde\x10\xaa\x3d\x8c\xda\xb8\xf4\x8a\x6d\x55\xc2\x99\x56\xf6\x3b\xfc\x07\xe3\x41\x74\xc8\xf0\x2a\x49\x48\x57\xe6\x4e\xa6\x1f\xb3\x30\xb9\xa6\xe4\xfb\x1b\x96\x7d\x1c\x24\x30\x21\xd4\xda\xb9\x9c\x02\x4f\x87\x01\xdc\x6a\x09\xaa\x08\xe1\xd6\xeb\xca\x54\xd9\x59\x36\xd4\x08\x81\xb2\x85\x82\x41\x12\x5e\xcd\xe9\x44\xc7\x50\x75\x3d\x4b\xd9\x66\x74\xa6\xeb\xf0\xc9\x18\x54\x4b\xc4\xf7\xeb\x9e\xaa\x5d\x74\x7c\x1f\xa4\x5a\xe1\x0a\x0c\x35\xce\xce\x3e\xa3\x42\x10\x52\xf1\xa5\x9e\x43\x32\x9b\x66\x70\x87\xfa\xb8\x53\x91\x41\x35\x32\x14\x0f\x47\x6a\xfb\x7e\xba\x58\x84\x47\x74\x1e\x2f\xe2\x82\x4e\xd0\xf3\x0a\x94\x54\x54\x95\x51\xea\xec\x75\x0e\xfe\xf3\xaf\x6a\x9b\x16\xb5\xf0\x74\x94\x95\x47\xca\x56\x49\x87\xf4\x87\x9f\xc9\x2a\x89\x0b\xfe\x84\xe2\x8e\xa6\x1d\x02\x5a\x94\x9c\xbe\xc7\x1e\xa3\xc3\x53\xa5\x25\xa8\xf6\x51\x5b\xf6\x94\xa2\xcb\xe6\x1f\x9c\xa4\xd7\x7a\xe0\xec\x90\xd7\x12\x86\x4b\x68\xa7\x61\x04\x45\x51\xf8\xc6\xd0\xcd\x67\x7a\x9d\xb3\xdf\x1c\xa8\xcd\x10\x95\xa2\x6b\x95\x82\xf7\xa4\xed\xe3\x69\xd5\xad\xfb\x29\x84\x78\x58\xae\x86\x90\x0d\x43\x7a\x2a\x60\xb1\x7d\x72\x25\xfb\xad\x88\x51\x17\x05\x0c\x9a\x06\xfd\xf1\x45\x98\xdf\x1c\x21\xf1\x71\xe1\x88\x65\x97\x30\xc5\xfc\x9c\xd9\x61\xcd\xd5\xe8\x94\xbe\x24\x33\x6a\x97\x8e\xa8\x94\x83\x63\x98\x69\x8e\xa1\x00\x2b\x1e\xd7\x7e\x77\xaf\x9d\x5b\x22\x06\xbe\x48\x6f\x68\xd2\x68\x73\xbd\xf6\x56\xb8\x8d\x1e\x17\xc6\x70\x5c\xc0\xdf\x8b\x6e\x58\x0f\xb6\xed\x71\xb9\x4a\x1e\x06\xb6\x33\xa3\xa6\xb7\x4a\x44\xf2\x99\x01\x6a\x64\x5b\x4b\x70\xf5\xb9\xd1\xa5\x74\x93\x04\x28\x7e\x37\x40\x90\xe3\x5c\xc5\xe9\x2e\x34\x9f\xed\xf1\x94\xcf\xd5\x50\xbb\xe5\x94\xa4\xee\x35\x3d\x5f\x87\xcb\x2c\xbd\x6f\x9d\x25\x0e\x97\xf9\x78\x11\x5e\x2b\x90\xec\xab\x1b\xf4\xf1\x11\xc5\x9e\x76\x99\x2e\x4b\x26\xdd\xc3\x2c\x0b\x1f\xd6\x6b\x0b\x8e\x08\x75\x97\x4c\x1c\xe1\x92\x44\x25\xb7\x07\xf3\xed\x3a\x80\x76\xce\xfc\x76\xb6\x59\xda\x0c\xa1\x92\xc4\x70\xac\xd7\x9d\xc7\x47\x3a\xcf\xe9\x7a\x0d\x7f\x93\x49\x4d\x2f\x98\xaa\x1c\x0f\x26\xea\x21\xd0\x87\xe2\xd2\xc5\x18\x1c\x17\x95\x42\x42\x55\xea\x79\x22\x05\x2c\x62\x13\x93\x80\x1d\xb7\xa8\x3b\x1d\x9d\xd7\x8e\xa8\xcf\x47\x5c\xd0\x5f\xae\xaa\xad\xa3\x18\xcf\x7d\xb7\xf1\x2c\xe5\xc2\xb2\xa0\x36\x6a\xee\x67\x3b\xb1\x1f\x3c\x1f\xbb\xef\x68\x43\x01\x39\x1c\x0e\xa5\x94\xa2\x2a\xae\x11\x69\xdc\xe7\x87\xfd\x5f\x04\x34\x4d\x6e\xc5\x77\x2f\x34\xa8\x92\xdf\x46\x83\x8f\xc7\xe7\x67\x6a\x1f\xe5\xa9\xaf\xa7\xe2\x41\xd1\x07\x10\x62\xbe\x88\x5c\x84\x95\x29\x11\xe7\xfa\x33\xd9\x45\x81\xe1\xbd\x82\xc0\x0d\x46\xb8\x85\xc0\x11\x84\xef\x54\x0a\x0b\xff\xe3\x92\x90\x3a\x01\xae\xec\xe3\x06\x13\xea\x9e\xc4\xc9\xcd\x97\x30\xcb\x7d\x44\x5a\x34\x36\x50\xe7\xa7\x21\x38\x39\xff\xf8\xdb\xc7\xd1\xf9\xe7\xa1\xcf\x95\x70\x2e\xe2\x70\x74\xde\x1f\x8c\xc7\xb6\xce\xb3\x80\x5d\x02\xf8\x25\x9d\xaf\x16\x8e\xdc\x86\xc9\x16\xda\x3d\x4d\x21\x20\x44\xaf\x56\x74\xf1\x31\x84\x7b\x09\xf4\x0f\xd2\xfd\x94\x82\x3b\x11\xec\xde\x86\xd9\x2e\x38\x6a\xbb\x93\x14\xa2\xf4\xac\x9b\xc3\x1f\xff\x92\xe3\x24\x58\xc7\xf5\xba\x07\x9f\xfa\x29\x8c\x09\x7e\x4c\xe6\x11\x44\xce\x51\x54\x42\x5e\x84\x9e\xe8\x7d\xf7\x96\x4f\x63\xd7\xce\x0a\x18\xc6\x75\x17\xb5\x26\xe3\x2a\x6a\x58\x0f\x71\xae\x04\x82\x4a\x62\x8d\xe0\xf9\xdb\x48\x79\xda\xce\xe8\x3a\x4b\xb9\xa3\x49\x6c\x60\x87\x72\x0e\x06\xf7\x45\x16\x22\xb5\xcd\xab\xeb\xd8\xc5\x65\xe7\xd3\x70\xe9\x5d\x6a\xdf\x1a\x62\x47\xd5\x10\x8b\xfd\xe1\x66\x0f\xda\xe2\xe5\xe1\x64\x02\x4e\x70\x2e\x3b\xc8\x3d\xe4\x36\x52\x1b\x6f\xad\x67\x72\x52\xfa\xae\x3e\x3e\x3e\x0f\x3b\x66\xff\x95\x53\x81\xda\x95\xea\x22\xb0\x7f\xeb\x99\xa2\xde\x43\x59\xf7\xef\x0b\xbf\xe9\xc2\x61\xe0\x71\xf7\xbd\x3c\xc2\x5b\xaf\x71\x4d\x3d\x3a\x88\x4d\x03\xc1\xcb\x1d\xe1\x5d\x3a\xef\x36\x79\xc5\xe5\xc3\x43\xbb\x78\x4e\xaf\xe9\xa4\x52\x92\xd5\x33\x07\xa9\x16\x81\x7e\x02\x8c\xac\x69\xdb\x2c\xa9\x3b\xf4\x51\x52\x0c\x55\xdc\x51\xd6\x35\x71\x27\xdc\xc8\x2b\xb8\x5c\x5c\x3d\x4c\xfa\x4e\xe3\x31\xf7\x64\x21\x72\x53\x56\x56\x9b\x41\xa0\x46\x86\xe5\xfa\xcb\x64\x31\x1b\xcc\xe3\x57\xbb\x56\x50\xf1\x77\x41\xc1\xb0\x54\x90\x1e\x16\x8e\xf0\x91\x25\xd0\x01\x7b\x9c\xc9\x83\x1a\x24\x89\xf5\xed\x96\x47\x37\x20\x9f\xd0\x69\xb4\x9a\xd3\xaa\x51\x3c\x6c\x24\x4a\x0f\xc7\x1c\x91\x1c\x0b\x21\xab\xe7\x0d\xc1\xa4\x1d\x9c\xfc\x1b\xa3\x48\xa5\xf6\xa9\xdc\xdf\xf2\x99\x01\x7a\x44\xf3\x38\xa3\x93\x3e\xda\x75\xa7\x7f\xeb\x49\x3d\xb5\xf2\x6f\x5f\x37\x9c\xb5\x73\x04\x75\x74\xda\x01\xbf\x19\x1d\x97\xc9\xf2\x32\xdd\x58\xc9\x81\xc5\x35\xa3\x7e\x6a\xc8\xcb\x9c\x2a\xd5\x62\x41\x98\x32\x69\x97\x8d\xd5\xe1\x10\x05\x60\x26\x0e\xac\xb0\x90\xbb\xd3\x71\x34\xa5\x38\x6e\x1e\x93\x61\xad\xcc\x3f\xff\xe9\xb4\x16\xaa\xd2\x60\xdb\xb2\x56\x47\xf7\x1c\x7b\x57\xb5\x22\x17\x61\x76\x4d\x0b\xcb\x56\xb8\x6c\x42\xad\x67\xa7\x2d\x78\x7b\x7d\xe5\x1b\xc7\x6d\x79\x9a\x93\x08\x4d\x46\xcf\x61\xee\xdc\xc3\x19\x1b\xc2\x69\xe1\xbc\x06\x45\x45\xd9\x94\xd4\x74\x16\xc1\xea\x89\xdf\x92\x8b\x6a\x56\xef\x7b\x91\x48\x64\xcc\x00\x15\xce\x37\x59\x77\xa8\x3c\x55\x80\xe5\x28\xc3\x0c\x86\xbd\x47\xf8\x25\x44\xf5\xc5\x94\x04\x12\xf7\x7f\xc0\xb0\x1a\x4e\x33\x81\xd8\x55\xfd\x22\x25\x6b\x28\x73\x44\xe6\x18\x4e\xa7\xa5\x8f\xba\x76\x1a\x47\x56\xcd\x8e\xb7\x4c\xd6\x9c\x6a\x23\x5a\xa6\xb1\xac\x92\xb2\x27\x2d\x89\x3b\x0f\xef\x5e\x8e\xb2\xb8\x0a\x83\xd4\xd6\xcc\xab\xb4\xa3\xec\x6f\xac\xe0\x26\x3c\x7c\x95\xf2\xb8\xa7\x50\xc8\x5c\xd3\xa7\x90\x86\xba\x0d\x73\x68\xca\x60\xa3\x30\x99\xa4\x8b\x9c\x6c\xc7\x45\x1a\x56\xa3\xbc\xb1\x5c\xac\xda\x89\x3c\x69\xf9\xf5\x73\x06\x5f\x0a\x5e\x2c\xf0\xa9\xa9\xf2\x9a\xa5\xa3\xdc\x7b\x25\x8f\x0d\xd6\x1a\x7c\xac\x77\x3d\x8d\xbe\xd5\xd1\x8d\x72\x1a\x62\x7a\x30\xb8\x6e\x9a\x6a\xc6\x7e\xc0\xcc\xb3\x31\x37\xfe\x97\x7a\xf1\xcc\xab\x88\xb3\xfc\xb8\x89\x97\xed\xc1\xae\xb9\x7c\x62\xd6\x81\x31\xdc\xcb\x48\xb8\xfc\x28\xbd\xaa\x57\x20\x5c\x15\x9b\xae\xea\x5d\xb0\x3e\x45\x86\x89\xb7\xb0\xca\x3f\xbe\x80\xbc\x9b\xa7\x5e\xaf\x20\xf1\x0e\x81\xf3\x15\xbf\x3e\x93\x93\xfa\x01\x1d\xaf\xff\xd4\x46\x52\x6a\xa7\x9d\x21\x40\xc0\xc0\xf4\x73\x4f\xcb\x47\x23\x2d\xbc\xe7\x1d\x49\xaa\xe5\x39\xeb\x85\xbf\xc7\xc9\xb5\x48\xb9\x7c\x75\x45\x6c\x9e\x3d\x27\xa0\x4c\x67\x92\x39\x89\x98\x35\x62\xa5\x2a\xb6\x33\x15\xf4\xe3\x49\x76\x8c\xfc\x0e\xf6\xba\xec\x7f\xbb\x7b\xf6\x71\x8e\xcf\xe3\xab\x7a\x2b\x95\x36\xa2\x96\xd3\x11\xad\x7b\x1c\xba\xe0\x78\xa9\x96\xed\x61\xe9\xa1\x55\x95\xf5\x21\x4b\x17\x38\x71\xd7\x4e\xb6\x80\x2f\x52\x1f\xa8\x9e\x14\x68\x8c\x92\x1b\x9d\x3f\x35\x04\xfd\xb2\x8c\x8e\x27\x26\x2b\xac\xc2\x88\x8e\x77\x03\xb8\x8e\xe9\xb9\xd0\xce\xc3\xbc\x88\xa3\x6a\xef\xc3\xca\xe3\x29\x73\xa5\x0a\x2a\x21\x7e\x9a\x69\xd0\x12\x13\x2d\x76\x67\x35\x6f\xdf\xae\xa9\xb2\xcd\xe3\x68\x46\x01\x45\x10\x57\x2f\x4b\x69\x41\x3f\x6f\x0f\x7a\x0a\x84\x1e\xfb\x57\x75\xe7\xfa\xa1\xa4\xac\xf9\xee\x78\x1c\x76\xa3\x34\xdc\xf2\xf7\x4d\x40\xc3\xa9\x57\xe1\x9d\x1b\xa0\xa2\xdc\x20\xac\x7c\x59\xa5\xa3\xce\xc9\x2f\x4d\x56\x18\xed\x9d\xf2\xb1\x0b\x9b\x3d\x4f\xe7\xdc\xec\x19\xe9\xe2\x8e\xa2\x93\x50\x56\xf5\x77\x94\x41\xdc\x04\x42\xc6\x4b\x21\x39\x19\x42\x96\xe0\x0b\x9a\x9c\x8e\x5a\x68\xf6\xd7\x3d\x4d\x99\x55\x78\xd4\x2a\x22\x20\x7e\x32\xa7\x4a\x75\x1a\x0a\x99\xf2\xc8\x8c\x02\x83\x7e\x96\xe6\xf9\xaf\x69\x42\xe5\x90\x55\xd3\x27\x1a\xce\x8b\x59\x7f\x46\xa3\x1b\x33\x5b\xc3\x9b\x1e\x2e\x66\xa0\x42\x67\xe9\x9c\x65\x16\x0f\x74\x81\x62\x4c\xbc\x65\x95\xae\x8c\x08\xde\x45\x3e\xb5\x02\x7a\x1e\x19\xbb\x94\xbe\x15\x96\x2b\xe8\xa4\x42\x03\x74\x3d\xaf\x84\xfa\x36\xa6\x74\x34\x04\x2a\x9e\x3e\x73\xe6\x79\xd5\x11\xc3\x62\x66\xa8\x38\x47\x82\x46\xe7\x3f\xef\xa9\xac\x80\x06\xfc\x39\x99\x39\xb9\x59\x85\xbb\xca\x9a\xc8\xbc\xc2\x4b\xda\x2d\xcd\xb8\x73\x76\x76\x9d\xf9\x00\xd5\x7c\xe8\xfe\x92\x51\x52\xce\xfa\xb7\xb7\x6f\x3a\x6a\x57\x51\x84\xe5\xba\x3f\x31\x80\xeb\x54\xf5\xec\x58\xb5\xee\x48\xc6\x7b\xad\xa7\x6a\x08\x5a\x9b\x48\x57\xbd\xbc\xc6\x39\x13\xc0\xcd\xb9\x0a\x0f\x1f\xd8\x95\xef\xdb\x30\x5e\xb4\x67\x8e\x2c\x51\x78\x25\x8d\xec\xeb\xad\x85\x9d\x8a\xf2\xa9\xee\x3a\xd0\xe7\x92\x61\x9d\xa2\x5c\x9a\xa5\xdf\x2f\xeb\xbf\x78\xea\xe9\x5b\x6e\x60\x7b\xc3\xde\x3f\xd4\xed\x5a\x47\x89\x90\x5e\xa6\xcf\x0d\x8e\x86\xc7\xf9\xe2\x02\xeb\x24\xfd\x25\x0d\x5c\x69\x72\xbd\xa9\x50\x14\x59\x7c\x85\xa7\x1c\x38\x61\xf7\x91\x69\x59\xb5\xd4\x44\x06\x03\x2e\x43\x4d\x34\x57\x8e\x63\x7e\xd7\x19\x98\xba\x7f\x38\x21\x2f\xb1\x83\xac\x57\x06\x9a\x53\x9d\xcf\x97\x9f\x93\xf7\xfd\x34\xbd\x89\xe9\x18\x7c\xd5\x1b\x88\x28\xf3\xbc\xf4\x1f\x70\x56\xfa\xea\x86\x53\x96\xf4\xc7\x7a\x34\x0d\x87\x11\xdf\xf2\xf8\xb7\x45\xd8\xeb\x0b\xa6\xc4\xeb\xf1\xa5\xb6\x20\xba\x70\xbb\xde\xaf\xd7\x76\x55\x79\xbe\xd8\xe8\x13\xaf\xdd\xfd\x0c\xa0\x8a\x73\xe5\x22\x29\x81\x41\x53\x88\x6e\xd6\x6c\x2b\xe5\x9f\x55\x18\xae\x1f\xdd\x89\xb4\xbf\x1a\x5c\xd4\x24\x82\x04\xb0\x37\xf7\xd3\x92\x40\x47\x06\x61\x33\x12\x1b\xe3\x9f\x2f\x07\xbe\x08\xc8\x17\xc6\x57\x07\x5f\xde\x18\x5d\x91\xa1\xcd\x5c\x7b\xf1\x99\x16\x3b\xd3\x10\xa9\x0b\xea\xe2\x7b\x33\xb3\xcd\x67\x3f\xae\xde\x37\xd0\x76\xd5\x46\xe1\x4f\xcb\xe0\xa7\x36\xf4\xb9\xf4\xbd\x30\xa7\x93\xe9\x5d\xac\x97\xc9\xb0\x94\x8a\x51\x7b\xad\x7c\xa3\x3c\x89\x21\x6d\xa5\x2f\x68\xdb\x2e\x67\x06\xc2\xce\x39\x94\xe7\xd8\x66\x86\x41\x69\x50\x13\x22\x8d\x91\x9b\x92\x2a\x51\xd3\x2f\xb8\x04\xff\xbe\xe4\x84\x91\x32\x6c\xcb\x4c\xc7\xf1\xa2\x64\xcb\xc6\xfb\x5a\xe0\x68\x16\x1c\x71\x66\xc0\xdf\xa0\x2b\xf7\x95\x1c\x64\x9a\x66\x77\x61\x36\x61\xeb\xc5\x82\x38\x26\x36\x78\x3c\xdf\xab\x3d\x16\x15\x74\xf3\x3e\xe5\xb1\xc4\xda\xd8\x98\xaa\xf6\x71\xe3\xd4\x55\x8a\x53\x60\x94\x66\xf3\x95\xdd\x9a\x4d\xd8\x48\xec\xa6\x1c\x57\xd8\xd3\xcc\x74\x39\x91\xfd\xbd\x83\x1f\x6b\x67\xf0\x72\xe2\xa7\x57\xa9\x34\xd8\x91\xc6\xea\x94\xf2\xac\xdb\xa8\x4f\xf9\x57\x70\xcd\xe5\xde\x71\x2a\x44\x6d\x83\xee\xdb\xb9\x27\x60\xc5\x49\x2f\xc4\xfb\xe7\x16\xf8\x3c\x75\xa3\x23\xb6\x66\x0e\xbe\xec\x2e\xdf\x4c\x34\xac\xdd\xaf\xdd\x6a\xc5\xdf\x0f\x8b\x29\x4f\x90\x2c\xc3\x62\xb6\x03\xbf\x50\xe1\x23\x55\xea\x5d\x27\x15\x66\x59\xb2\xc1\xb5\xfd\xa5\x23\x93\xf2\xc4\x29\x98\x82\x11\xa7\xc2\x5e\xf3\x53\xd5\x08\x36\x93\x7a\xfc\xd5\xde\xb7\x73\xbc\xf8\xa9\x3a\x77\x58\x03\xdf\x07\x90\x9f\xd3\xab\xdc\x7e\xf9\x10\xd3\x79\x49\xf9\x56\xb1\x4f\x50\xbc\xaf\x1f\x7b\x85\x42\x42\x2a\x55\xad\x5a\x81\xfd\x0e\xea\x48\xa5\x3c\x7f\x07\x67\xa1\xee\x0e\xeb\xb5\xde\x99\x78\xa0\xc0\x3c\xf3\xa5\xdd\xda\x57\x76\x95\x94\xdc\xfe\x9e\x26\x60\xd6\x3b\xd5\xc1\xaf\xf1\xf2\x43\xcc\x48\x31\x93\x91\xc1\xb7\xc4\xce\x46\x6e\xad\x72\x4a\x72\x88\x72\xa3\x62\xeb\x27\xf3\x0a\x01\xf3\xfb\x6d\x98\x11\x76\xcd\x04\x79\x47\x32\xfa\xc7\x2a\xce\xe8\xf6\x16\x7b\xb0\xf5\xc6\xea\x8c\xc0\xe1\x9d\x06\x0a\x5f\x77\xf2\xc9\x8d\x07\x38\x9a\xa7\xab\x49\xf9\xba\x35\xf4\x4b\xe8\x1d\x62\xe8\xf6\xb1\xa1\x2c\x58\xdb\x76\xf7\xfe\x63\x45\xb3\x87\x9c\xbf\x7f\xa2\x0c\xa9\x3c\x76\x0c\xeb\x42\xc4\x0b\x39\x00\xc7\xa3\xd9\x8a\xbf\xca\x58\xa0\x47\xb6\x4c\x39\xda\x32\x3b\xac\x1b\x07\x14\xd5\x2a\x5d\xc0\xd3\x3d\x3b\x3f\x1a\xfc\x76\x71\x82\x2f\x70\xfc\x3c\xe8\x5f\xfc\xf6\xf9\xec\xf0\xf3\xc5\xa7\xf3\xd1\xf1\xaf\x83\x23\x20\x67\x6b\xaf\x79\x81\xe8\xfd\x12\x73\x04\x52\x3e\xa1\xd7\x54\xc8\xfe\x36\xbd\xa5\x49\xd1\x21\x51\x0a\x4e\xe6\x7d\xf1\xc6\x3d\x3b\x68\xcd\x41\x44\xbb\xf3\xf4\x7a\x7b\x0b\x2f\x9b\x19\x8c\x2f\xc8\x68\xd0\x1f\x1c\x7f\x19\x1c\xc1\x8c\xc9\x5b\xf2\xf3\xf8\xfc\xac\xcb\x19\x1a\x4f\x1f\x38\xda\x37\xcd\x9c\x65\xd8\xb5\x05\xee\x4e\x98\x33\x7f\x45\x19\x4b\xf3\x6d\xce\xf7\x8e\x42\x71\x96\x75\xc8\x24\x2c\x42\x0f\xb1\xf8\x03\x4a\x05\xe1\x6a\x20\xcc\x69\x21\xb4\x45\xad\xfc\xb0\xe6\x39\xde\x5a\x64\x28\x23\x93\x30\x9f\x5d\xa5\x60\x51\xbc\x98\x24\xe4\x32\xcc\xf3\xbb\xb4\x05\xa0\x88\xd7\x61\xc9\x70\xca\x5d\xce\x94\xaf\x7b\x97\x5d\xf1\xba\x78\x8b\x91\xe4\x95\x7d\x36\x8e\xea\x3a\xbf\x56\x0b\x25\x7f\x60\xad\xb6\x11\x75\xfc\x6e\xef\x27\x49\x60\x77\xce\x2e\x10\xfa\x7b\xfc\x13\x89\xdf\xbe\x6d\x60\x3c\xfe\xe0\x12\x89\xbe\x5f\x63\x39\x9d\x5f\xe8\x03\x79\x07\x32\x7d\x24\x19\xb9\xd5\x02\x13\xfe\x94\x9c\x87\x49\x5a\x58\x99\x15\xad\x65\x14\xfe\x98\x19\xea\xa0\xb6\xb1\x89\x3f\x8c\x37\x15\xef\x9f\xc4\x9e\xaa\x3b\xce\xa5\x5c\x2c\xc9\xa4\xa1\x90\xa1\xb6\x3c\x92\x32\x07\x2c\xf2\x60\xfe\x57\x33\x8a\x09\x38\x0b\xdb\xbd\x8a\x55\xfd\x99\x81\xef\xda\xab\x96\xba\xd3\x04\x8f\x4a\xaf\x47\x7e\xfc\xf1\x87\x66\x48\xf0\xbb\x40\x8f\xed\x86\xcb\x65\xbe\x8b\xea\x8c\xe9\xaf\x2e\x56\xae\xbc\x85\xc7\xe5\x0b\xea\x4a\x9b\x78\xc6\xda\xb3\x55\xb2\xd5\x38\x06\x70\x78\x96\x4e\x60\x94\xe1\xf9\xf8\xa2\x19\x7c\x46\xc3\x09\xac\x4f\xaf\xdd\xda\x6e\x1d\x46\x11\x5d\x16\x5b\x80\x1e\x88\x9e\xe3\x29\x00\x70\x75\xf7\xf7\x3c\x6d\x41\x19\x43\x80\x25\xaf\x30\xaf\x1d\x74\xaa\x4c\x34\xf7\x3b\x77\x77\x77\x3b\xa8\xa2\x77\x56\x19\x48\x32\xde\x46\x37\x69\x89\xf7\x73\x4e\xb3\x9d\xc3\x6b\x40\x8d\x58\x23\xe8\xbf\x6b\x19\x44\xb3\xd3\xba\x11\x75\xb8\x62\x0b\xc6\x2f\x81\xe9\xe1\xaa\x48\xe9\xae\x15\xce\x8d\xd4\x1c\x4a\x27\x38\x0c\x20\x99\xcc\x97\xe9\xa2\xf3\x00\x6e\xe9\xb6\x90\x58\xc5\x24\x81\x5b\xdb\x62\x0f\x22\xbe\xab\x74\xf2\x80\x46\xdb\xb6\xd9\x26\xb4\x66\x73\xf1\x2e\xa0\x55\xde\xdb\xea\xe0\x95\x07\xdd\x9c\x7d\x43\x2f\xcf\x6f\xb2\x9c\x68\x3e\x09\x99\x02\x3c\x86\xbd\x46\xb4\x42\xe2\x1c\x56\xdb\x44\xca\x88\xa0\xc5\x00\x05\x01\x10\x6c\x6f\xad\x8a\xe9\xdf\x1c\x7e\x94\xab\x1f\xb0\x6b\x0b\xcd\xd0\x96\xc2\xbf\x68\xb6\x4a\x6e\x80\x83\x9c\x3d\x6f\xdf\x11\xf6\x80\xac\xdb\x63\x84\x98\x43\x45\xd8\x52\x23\xaa\xcc\xc1\xa1\x9b\xc7\x13\xbd\xd0\x57\xea\x4e\xd2\x84\x3a\xbc\x4e\x13\xbc\x61\x16\x8e\xe6\x3a\x70\x10\x42\x3e\xe1\x2c\x4b\xb3\x2d\xdd\x2d\x4a\x9b\x1c\x1e\x73\xce\x5b\x03\xec\x84\xe2\xc0\x7b\xb7\x91\x26\x36\xf5\x69\x18\xcf\xb7\x5b\xf4\xd9\x70\x72\xcc\x5f\x01\x25\x7f\x04\xf2\x01\xbb\x44\x71\xd2\x15\x51\x6d\x98\x21\xe8\x04\xf6\xd2\xf9\x16\xe9\x09\x4d\x2d\x1e\xbc\xf0\x22\xdc\x41\x68\x4c\xb7\x25\xb5\xf5\x6c\x40\x78\x90\xd0\x1a\x61\x71\xda\x54\x07\x55\xa0\xc1\x94\x07\xfa\xad\x53\xad\x13\xb2\x0d\xe1\x36\x86\xbd\xe3\x19\xd8\x4e\x11\xf7\xb2\x14\x8b\x37\x3f\x83\x2c\xce\x7b\xbd\x76\x79\x18\x25\xd4\xee\x9e\xa4\xc9\xb5\x8c\xac\xf3\x68\x46\x27\x2b\xfd\x46\xbc\xb1\x78\x36\xb8\x5f\x62\xae\x5e\xe4\xf6\x19\x71\xa2\xc5\x48\x83\xf2\x04\x8c\x75\x24\x1b\x94\xc9\x10\x2b\x02\x57\xf3\x0a\xbe\xb7\xaf\x8e\x27\x0e\x82\x45\x61\x8f\x0e\x98\x2c\x45\x75\xcc\x37\xbc\x2c\xfd\x1b\x7c\xfc\x16\x88\x14\x82\xe8\xf6\x0d\x46\xf9\x26\x83\xbb\x0a\x40\x1c\xcd\x96\x00\x42\x60\x2b\x00\x71\x8f\x02\x03\x50\xaa\x74\x5c\xc7\x80\x8e\xf5\xe3\x99\x91\x21\xcd\x16\x71\x9e\xbb\x52\x28\xc4\xcc\xa1\x28\xb0\xae\x25\x25\x8e\xdc\x1a\x43\xc2\xef\x46\xeb\x1d\x83\x51\xbe\xa1\xae\x7b\xe0\xb4\x94\x0b\x79\xe2\xa2\x28\x17\x0e\xe2\xa0\x6c\x93\xe7\xc6\x15\x83\xaa\x1c\xb1\x8c\x13\x43\xe3\x2d\x60\xb7\xc4\x5d\x19\x78\xf3\x6c\x97\xf3\xae\x78\x7e\x09\x10\xbf\x9d\xee\x53\x98\x0f\xfa\xea\x95\xa5\x8c\xa8\xf3\x4c\x3b\x8f\x72\xde\x24\xe5\xbc\xe9\x0e\x88\x5d\xe5\x3b\x34\xcc\x0b\x76\x57\x94\x5a\xaa\xb7\x21\x8e\x3b\xd0\x03\x3b\x07\xcf\xc0\x41\x57\x1c\x07\xa3\x43\xa0\x40\x29\xf5\xde\x65\xe4\xba\x8a\xad\x62\x15\xb6\x38\x2e\x57\x76\xbc\xb2\x39\x52\xc0\xc4\x3e\xd0\xea\x59\x35\xc6\x0b\x00\xe7\x2b\x9d\x6e\xdd\xf5\xfc\x57\x39\x95\x4b\xa2\x9b\x8f\x93\xbf\x93\x12\xd7\xcc\x37\xe3\x26\x2b\xc9\x35\x9e\xb6\x6f\xcb\x08\xeb\xb2\x2b\x46\x11\xc4\x14\xf8\x22\x29\xa0\x92\xb4\x34\x2c\xcb\x26\x03\x59\xe5\x03\x36\xde\x8e\xc5\x34\xb1\x1f\xbd\x5c\x31\x6e\xde\xd3\x5e\xfa\x74\x5f\xbb\x65\x5e\xc6\xea\x59\xfe\x56\x97\xb0\x7a\xef\x56\x35\x6e\xb7\xac\x2e\x3a\xd5\x9e\x13\xeb\xda\x53\xad\xd9\x28\x01\x6a\xb8\x99\x55\xbf\x95\xd5\x1c\x47\xb9\xa3\xd5\x68\xc2\x6c\x61\x64\x6a\x53\x03\xc6\xff\x7e\xbc\xe7\x00\xd9\x79\xe5\xa9\xd4\xab\x95\x85\xb1\xee\x6a\x75\xd7\x95\x69\xc6\x5d\x2f\x0d\x52\xd7\xbb\x63\x83\xf9\xaf\xce\x7d\xf9\x5b\x71\xbd\x8b\xcc\x5a\x69\x74\xd0\x3b\x84\x18\x36\xcd\xe2\x3f\xa9\xb3\x34\xc1\x46\x29\xba\x1d\x89\xfc\xe8\x5f\x3c\x20\xfc\xc0\x0c\x2f\xda\xbf\x2a\x2b\xfe\x8f\x28\xd7\x17\x34\x2b\x6b\x23\xb0\x42\xc1\x59\x05\xb3\x09\x36\xe9\x72\x6d\xd2\x77\x73\xea\x47\x26\xed\xff\x88\x8b\xd9\x13\x69\x1f\x19\x94\x5b\x1d\x5d\x95\x76\xca\xbd\xba\x5f\xd9\xcd\xba\x56\x91\x9c\xbe\x53\x9d\x3b\x42\x35\x87\x0d\xca\x5d\xbd\x20\xd4\xbe\x4f\x53\x57\x60\xe3\x1f\x7a\x3d\x71\x59\xae\xd0\x60\x47\x74\x4e\x51\xf6\xca\x12\x3a\xa0\x1f\xdf\x66\x6e\xd0\x70\xec\x3f\x6e\x81\x59\xa0\x8c\x9f\x05\x9b\x45\x42\xe0\x5f\x1b\x97\x99\x3c\xca\x6b\xe2\x82\x1c\xc2\x34\xba\x28\x4f\x2c\xf1\x91\xb8\x9f\x97\xe8\x9e\x84\x80\xc7\xab\xd4\x3b\x5e\xbb\xa3\x5a\x42\x17\xdb\x14\xae\xfd\x7f\x00\x00\x00\xff\xff\x85\xf5\xd4\x38\xa1\x68\x00\x00")

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

