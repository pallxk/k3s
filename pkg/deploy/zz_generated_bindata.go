// Code generated for package deploy by go-bindata DO NOT EDIT. (@generated)
// sources:
// manifests/ccm.yaml
// manifests/coredns.yaml
// manifests/local-storage.yaml
// manifests/metrics-server/aggregated-metrics-reader.yaml
// manifests/metrics-server/auth-delegator.yaml
// manifests/metrics-server/auth-reader.yaml
// manifests/metrics-server/metrics-apiservice.yaml
// manifests/metrics-server/metrics-server-deployment.yaml
// manifests/metrics-server/metrics-server-service.yaml
// manifests/metrics-server/resource-reader.yaml
// manifests/rolebindings.yaml
// manifests/traefik.yaml
//go:build !no_stage
// +build !no_stage

package deploy

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _ccmYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x94\x41\x8f\x13\x31\x0c\x85\xef\xf3\x2b\xa2\x1e\x91\xd2\x15\xe2\x82\xe6\x08\x07\xee\x2b\xc1\xdd\x4d\x1e\xdd\xd0\x4c\x1c\xd9\x4e\x61\xf9\xf5\x68\x3a\x5d\x51\x3a\xdb\xd2\x16\x10\x7b\x8b\xac\xf8\xf3\xf3\x73\x62\xaa\xe9\x13\x44\x13\x97\xde\xc9\x8a\xc2\x92\x9a\x3d\xb0\xa4\xef\x64\x89\xcb\x72\xf3\x56\x97\x89\xef\xb6\xaf\xbb\x4d\x2a\xb1\x77\xef\x73\x53\x83\xdc\x73\x46\x37\xc0\x28\x92\x51\xdf\x39\x57\x68\x40\xef\x36\x6f\xd4\x87\xcc\x2d\xfa\xc0\xc5\x84\x73\x86\xf8\x81\x0a\xad\x21\x9d\xb4\x0c\xed\x3b\xef\xa8\xa6\x0f\xc2\xad\xea\x98\xe8\x5d\x60\x96\x98\xca\x61\xbd\xce\x39\x81\x72\x93\x80\xfd\xa5\x0c\x52\x68\xe7\xdc\x16\xb2\xda\xc7\xd6\xb0\x09\x20\x20\xc3\xee\xd8\x6a\x1c\x8f\xb3\x1a\x8b\xc5\x1c\x89\x2d\x8a\x1d\x21\x0f\x50\x95\x2c\x3c\x5c\x0d\x2d\x1c\x8f\x65\x2e\x5e\x2d\xae\xc8\xbd\x53\x23\x6b\xba\x0b\x28\x64\x9b\xc2\x61\xec\x00\x3b\xe9\xbb\x08\xfc\xc4\x99\xf2\x38\x9e\xf0\x31\x27\x9d\x0e\x5f\x6f\x42\xcf\xb4\x5d\xeb\xdd\x9e\x45\x21\x70\x3b\x37\x99\x51\xef\x65\x86\xd2\x00\xad\x34\x93\xf7\x3b\x16\xd5\xaa\x73\x5a\x24\x0c\x5c\x14\xc7\xca\x9e\x9f\x6f\x4c\x1a\x78\x0b\x79\xdc\x3f\xe9\xe7\x1e\x60\x89\x95\x53\x31\xcd\x73\x07\x4f\xcd\xc4\xfb\xee\xf6\x1f\xfb\x2e\x95\x98\xca\xfa\xea\x8f\xcb\x19\xf7\xf8\x3c\xde\x7e\xea\xf2\x4c\xe5\xce\xb9\xf9\xaa\xb8\xa8\x8e\xb6\xd5\x17\x04\xdb\xed\x88\x09\xf1\x51\x21\x97\xe5\xba\x9f\xc3\xee\xdd\xa6\xad\xe0\xf5\x51\x0d\xc3\x7f\x71\xcc\x8f\x7c\x1f\x91\xb1\x26\xe3\xbf\x6a\xe0\xd4\x55\x7f\x54\xe0\xa5\x38\xf7\x87\x96\xa1\x58\x0a\x3b\xb2\x17\x50\x3c\x27\xee\x46\x4b\x7f\xf1\x12\xdf\x0c\x65\xec\xcd\x53\x4d\xe3\xf2\x39\x29\xe3\x9f\xf8\xfb\x23\x00\x00\xff\xff\xde\xc0\x02\x82\x7a\x07\x00\x00")

func ccmYamlBytes() ([]byte, error) {
	return bindataRead(
		_ccmYaml,
		"ccm.yaml",
	)
}

func ccmYaml() (*asset, error) {
	bytes, err := ccmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ccm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _corednsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x51\x6f\xdb\x38\x12\x7e\xf7\xaf\x20\x04\xe4\xe5\x70\x72\xe2\x0b\xda\xcb\xf1\x2d\x8d\xdd\x36\xb8\xc4\x35\x6c\xa7\x40\xb1\x58\x04\x34\x39\xb6\xb8\xa1\x38\x5c\x92\x72\xe2\xed\xe6\xbf\x2f\x28\xc9\xb2\x68\x2b\x69\x92\xed\xfa\xc5\x92\x86\xf3\x0d\xf9\x71\xf8\xcd\x90\x19\xf9\x15\xac\x93\xa8\x29\x59\x0f\x7a\x77\x52\x0b\x4a\x66\x60\xd7\x92\xc3\x39\xe7\x58\x68\xdf\xcb\xc1\x33\xc1\x3c\xa3\x3d\x42\x34\xcb\x81\x12\x8e\x16\x84\x76\xf5\xbb\x33\x8c\x03\x25\x77\xc5\x02\x52\xb7\x71\x1e\xf2\x5e\x9a\xa6\xbd\x36\xb4\x5d\x30\xde\x67\x85\xcf\xd0\xca\x3f\x98\x97\xa8\xfb\x77\x67\xae\x2f\xf1\xb8\x09\x7a\xa1\x0a\xe7\xc1\x4e\x51\x41\x14\x51\xb1\x05\x28\x17\x9e\x48\x19\xc2\x6a\xf0\x50\xba\x2e\x10\xbd\xf3\x96\x19\x23\xf5\xaa\x8a\x91\x0a\x58\xb2\x42\x79\xd7\x4c\xb5\x9a\x10\xdd\xce\xd8\x16\x0a\x1c\xed\xa5\x84\x19\xf9\xc9\x62\x61\x4a\xe4\x94\x24\x49\x8f\x10\x0b\x0e\x0b\xcb\xa1\xfe\x06\x5a\x18\x94\xba\x04\x4b\x89\xab\x48\xa9\x5e\x0c\x8a\xea\xa1\x59\x7f\x78\x5d\x83\x5d\xd4\xbe\x4a\x3a\x5f\x3e\xdc\x33\xcf\xb3\xc3\x78\x42\x3a\x8e\x6b\xb0\x9b\x9a\x87\x67\xa2\x2b\xf9\x43\xf4\xbf\xc5\xf6\x07\xa9\x85\xd4\xab\x88\x74\xa6\x35\xfa\xd2\xb3\x66\xbe\x0b\x32\xda\x0c\x56\x78\x2c\x8c\x60\x1e\x28\x49\xbc\x2d\x20\xf9\xf9\x7b\x87\x0a\xa6\xb0\x2c\xe7\x57\xb3\xf9\xcc\x5a\x7b\x84\x1c\x26\xd6\x13\xc8\xae\x58\xfc\x06\xdc\x97\x89\xd1\x79\x04\xde\x9c\xf8\x3b\xc2\x51\x2f\xe5\xea\x9a\x99\xb7\x1c\xa7\xed\xf0\x0b\xb4\xb0\x94\x0a\x28\xf9\xb3\xe4\xb4\x4f\xdf\x9d\x92\xef\xe5\x63\xf8\x81\xb5\x68\x5d\xf3\x9a\x01\x53\x3e\x6b\x5e\x2d\x30\xb1\x69\xde\x76\xdb\x41\x8e\xbe\x5f\x5c\xdd\xcc\xe6\xa3\xe9\xed\xf0\xcb\xf5\xf9\xe5\xf8\xf1\x88\x48\x9d\x32\x21\x6c\x9f\x59\xc3\x88\x34\xef\xab\x87\x5d\x24\x52\x9e\x00\x22\xb5\x03\x5e\x58\x68\x7d\x5f\x32\xa5\x7c\x66\xb1\x58\x65\xdd\x28\xcd\xd8\xc7\xdd\x44\xd1\x79\x47\x8e\xc1\xf3\xe3\x9a\x8a\xe3\x31\x0a\xf8\x5c\x7e\x6e\x07\xf5\x5e\x91\xf7\x27\xad\x0f\x16\x14\x32\x41\x06\xef\x5c\xf7\x14\x3a\x82\x19\x8b\x39\xf8\x0c\x0a\x47\xe8\xff\x06\xef\x4e\x1b\xc3\x12\xed\x3d\xb3\x82\xf4\xab\x99\x84\xe3\xa8\xd6\x7d\x8e\x7a\xd9\x0c\xe1\x8c\x67\x40\x4e\x77\x33\x50\x88\xa6\x17\x4f\xa6\x65\x63\x62\xc1\x14\xd3\x7c\xc7\x8f\xcc\x0d\x5a\x1f\x2f\x95\x17\xce\x63\x7e\xfc\xaf\x7e\xd0\x03\x2b\x45\x35\xba\x9a\xf0\xb3\xe3\x83\x22\x81\x3d\x48\x39\x66\x8c\xdb\x1d\xf4\x21\x18\x85\x9b\x1c\xde\xa6\xe3\x7b\x47\xf8\xcc\xa5\xcc\x98\x7a\x48\xe5\xb8\x7f\xb0\x2b\xe0\x24\x64\xea\x70\x3c\x4b\x7a\xce\x00\xa7\xa5\xba\xad\x65\x98\xdf\x67\xe9\x3c\xda\xcd\x95\xcc\xa5\xa7\x24\x30\x19\x64\xc0\xc3\x6a\x53\xc5\xf0\x1b\x03\x94\x4c\x51\x29\xa9\x57\x37\xa5\xa0\x54\x02\xd4\xfe\x42\x6b\x42\x73\xf6\x70\xa3\xd9\x9a\x49\xc5\x16\xe1\x54\x0c\x02\x1c\x28\xe0\x1e\x6d\x35\x26\x0f\x02\x79\xd5\x5a\x43\xf7\x2a\x3c\xe4\x46\x35\xc0\x6d\xa2\xca\x9d\x8c\xfc\x9f\xe2\x61\xbb\xd2\x2a\xc9\x24\x5a\xe9\x37\x17\x8a\x39\x37\xae\x28\xa9\x28\x4d\x79\x25\x47\x29\xb7\xd2\x4b\xce\x54\x52\xbb\xb8\x48\x71\xc6\x7b\xfb\x53\x52\x83\x0a\x6c\x5b\x94\xc3\x2f\x25\x77\xb0\x09\x84\xd7\x70\xe7\x42\xa0\x76\x5f\xb4\xda\x24\xad\x23\x81\x26\x78\xa2\xa5\x24\x19\x3d\x48\xe7\x5d\x72\x00\xa0\x51\x40\x1a\x24\x76\x4f\xd8\x39\x6a\x6f\x51\xa5\x46\x31\x0d\x2f\xc4\x24\x04\x96\x4b\xe0\x9e\x92\x64\x8c\x33\x9e\x81\x28\x14\xbc\x3c\x64\xce\x02\x43\x3f\x23\x56\x88\x30\x8b\x12\xe2\x30\x63\xd1\x51\xa2\xa4\x2e\x1e\x1a\x9a\x0d\x2a\x5c\x6d\x66\x26\x28\xe6\x05\xea\x90\xa0\xa1\x10\xb7\x49\xcf\xd9\xc3\xec\x0e\xee\xab\x94\xdb\xfe\xb6\x9e\xff\x0f\xab\x8b\x83\x04\x89\x0b\x47\xa3\x35\xfa\x3e\x03\x7d\xa3\x1d\xf3\xd2\x2d\x65\x95\xbf\x43\x1c\xa3\xdf\xae\xa1\x35\xb4\x4c\xc0\xc3\x75\x3c\x91\xe0\xcf\xa7\x29\x21\x61\x47\x99\xd4\x60\x1b\x8f\xf4\x40\x0f\xaa\x9f\xcc\xd9\x0a\x28\x39\xfa\x3e\xfb\x36\x9b\x8f\xae\x6f\x87\xa3\x8f\xe7\x37\x57\xf3\xdb\xe9\xe8\xd3\xe5\x6c\x3e\xfd\xf6\x78\x64\x99\xe6\x19\xd8\xe3\x5c\x86\xda\x03\x22\xad\x21\xb6\xff\x74\xd0\x1f\x9c\xf4\x07\x31\xe2\xa4\x50\x6a\x82\x4a\xf2\x0d\x25\x97\xcb\x31\xfa\x89\x05\x07\x65\x99\xad\x7e\x51\x2b\xd4\x90\x10\x24\x63\x6f\x91\x39\xe4\x68\x37\x94\x0c\xfe\x7b\x72\x2d\xa3\xba\xf0\x7b\x01\x6e\x7f\x34\x37\x05\x25\x83\x93\x93\xbc\x13\x23\x82\x60\x76\xe5\x28\xf9\x85\x24\x69\x28\x00\xc9\xbf\x49\x12\x69\xf0\xb6\x10\x27\xe4\xd7\xc6\x65\x8d\xaa\xc8\xe1\x3a\x9c\xde\x28\x55\xb6\xd4\x86\xfa\x9f\x56\x83\x5a\xf1\xf3\x30\x7e\xc2\x7c\x46\x23\x95\x8f\xd6\xc2\x44\x38\xcf\x94\x84\xb6\xea\x10\xb8\x2c\x07\xe9\x2b\xf1\xeb\x2a\xf2\xe3\x30\xa1\xfe\x44\xcb\x69\xb2\x67\x82\xd6\x53\xd2\x2a\xa0\xdb\xaa\x12\x4f\xdf\x58\xf4\xc8\x51\x51\x72\x33\x9c\xbc\x16\x27\xf5\xdc\x74\x62\xcd\x2f\x9e\xc1\x8a\xca\xfa\x16\x2d\x07\x6f\x25\xef\x9e\x59\x1b\xad\xec\x68\x82\x74\xa3\xf6\xf0\xe0\xdb\x19\xc4\x94\xc2\xfb\x89\x95\x6b\xa9\x60\x05\x23\xc7\x99\x2a\xe5\x98\x86\x96\xc3\xb5\x59\xe7\xcc\xb0\x85\x54\xd2\x4b\xd8\xcb\x41\x26\x44\xfc\x21\x25\xe3\xd1\xfc\xf6\xc3\xe5\x78\x78\x3b\x1b\x4d\xbf\x5e\x5e\x8c\x22\xb3\xb0\x68\xf6\x1d\x98\x52\x1d\x1b\x37\x45\xf4\x1f\xa5\x82\xba\xb7\x8d\xb7\x51\xc9\x35\x68\x70\x6e\x62\x71\x01\x6d\xbc\xcc\x7b\xf3\x09\x7c\x1c\xc2\x54\xf9\xb2\xd7\x40\x92\x3a\x1d\x28\x39\x3b\x39\x3b\x89\x3e\x3b\x9e\x41\x20\xf9\xf3\x7c\x3e\x69\x19\xa4\x96\x5e\x32\x35\x04\xc5\x36\x33\xe0\xa8\x85\xa3\x71\x03\x67\xc0\x4a\x14\x8d\x6d\xd0\xb6\x79\x99\x03\x16\x7e\x67\x6c\xd9\x5c\xc1\x39\x38\x37\xcf\x2c\xb8\x0c\x95\x88\xad\x4b\x26\x55\x61\xa1\x65\x3d\x8d\xda\x60\xf9\x6a\x2a\xe2\xe6\xb9\xc5\xc4\xe0\x6c\xf0\x66\x26\x9e\x21\xe2\x3f\xff\x30\x0f\x42\xbb\xad\x02\x0f\xab\x6b\x57\x6d\xa8\x04\xe4\x15\x02\xc6\xb7\x17\x9b\x98\xb7\xee\x82\x52\x52\xe1\x21\x77\xfb\x29\x5d\x36\x04\x5b\x55\xdd\xab\x63\xd5\x16\x74\x1a\x6b\xc7\xe6\xb6\xd0\xe9\x79\x68\x7d\xa1\x76\xbe\x64\x69\xe9\x81\x90\x86\x6e\x25\xa8\x02\x53\xf5\x19\x7c\xf2\x4e\x58\x5f\x32\x3b\x1a\xf3\x56\xc5\x7e\xb2\x33\x3f\xb8\xa3\xef\x6e\x36\xa1\xe3\xa8\xf2\x33\x09\x5a\x98\x74\x98\x1d\xb7\xcc\x3c\x79\x57\x7f\x41\xa3\xbf\xed\x63\xeb\xbe\xb5\x85\xf4\xd2\x2b\x41\xdc\xa9\x77\xc5\xac\x63\x5c\x4e\x68\xfb\x92\x3a\x9e\x3d\x1e\xf5\x5a\x95\x29\xdd\xab\x3b\xa6\x5d\x50\xf6\xcb\x4f\xda\x51\x5c\x9e\x70\xa8\xaa\x42\xda\x51\x3f\x4c\x5c\x66\x62\x97\xbf\x02\x00\x00\xff\xff\x2b\x8d\x4e\x50\x53\x13\x00\x00")

func corednsYamlBytes() ([]byte, error) {
	return bindataRead(
		_corednsYaml,
		"coredns.yaml",
	)
}

func corednsYaml() (*asset, error) {
	bytes, err := corednsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localStorageYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x5f\x6f\xdb\xb6\x16\x7f\xd7\xa7\x38\x57\xb7\x79\xb8\x17\xa5\x9d\x6c\x05\x32\xb0\xd8\x83\x9b\x38\x69\x80\xc4\x36\x6c\x77\x43\x51\x14\x06\x2d\x1d\xdb\x6c\x28\x92\x20\x29\xb7\x6a\x96\xef\x3e\x90\x94\x1d\xc9\x71\x13\x07\xdb\xde\xa6\x17\x81\x87\xe7\xef\xef\xfc\x23\xd3\xfc\x37\x34\x96\x2b\x49\x61\x7d\x92\xdc\x72\x99\x53\x98\xa0\x59\xf3\x0c\x7b\x59\xa6\x4a\xe9\x92\x02\x1d\xcb\x99\x63\x34\x01\x90\xac\x40\x0a\x42\x65\x4c\x10\xcd\xdc\x8a\x68\xa3\xd6\xdc\xcb\xa3\x21\x36\xca\x11\x56\x0b\x46\x76\xab\x59\x86\x14\x6e\xcb\x39\x12\x5b\x59\x87\x45\x42\x08\x49\x9a\x96\xcd\x9c\x65\x1d\x56\xba\x95\x32\xfc\x3b\x73\x5c\xc9\xce\xed\x2f\xb6\xc3\x55\x77\xeb\xd3\x99\x28\xad\x43\x33\x56\x02\x0f\x77\xc8\x78\x6e\x53\x0a\xb4\x34\x21\xc0\x34\xbf\x34\xaa\xd4\x96\xc2\xa7\x34\xfd\x9c\x00\x18\xb4\xaa\x34\x19\x06\x8a\x54\x39\xda\xf4\x35\xa4\xda\xbb\x65\x1d\x4a\xb7\x56\xa2\x2c\x30\x13\x8c\x17\xe1\x26\x53\x72\xc1\x97\x05\xd3\x36\x88\xaf\xd1\xcc\x83\xe8\x12\x9d\xbf\x16\xdc\x86\xff\x57\xe6\xb2\x55\xfa\xf9\x79\x93\x28\x73\xad\xb8\x74\x7b\xcd\x46\xa2\xca\x77\x6c\xfd\xff\x20\xc5\x6b\xf4\x5a\x5b\x82\x99\x41\xe6\x30\x28\xdd\xef\x9f\x75\xca\xb0\x25\xd6\xd0\x3f\x56\x5a\xdf\x67\x82\x59\x8b\x07\x22\xf0\x97\x12\xfd\x8e\xcb\x9c\xcb\xe5\xe1\xf9\x9e\x73\x99\x27\x3e\xe9\x63\x5c\x78\xe6\x4d\x78\x4f\x18\x4e\x00\x1e\x17\xd8\x21\x65\x65\xcb\xf9\x17\xcc\x5c\xa8\xac\xbd\x6d\xf3\x4f\x35\x0b\xd3\xda\x3e\xc0\x75\x8e\x5a\xa8\xaa\xc0\x17\xf4\xe9\x8f\x4d\x59\x8d\x19\x0d\x69\x8f\xbc\xef\xb9\xcf\x79\x75\xcd\x0b\xee\x28\x1c\x27\x00\xd6\x19\xe6\x70\x59\x79\x2e\x00\x57\x69\xa4\x30\x56\x42\x70\xb9\xfc\xa0\x73\xe6\x30\xd0\x4d\x93\x12\x59\x01\x0a\xf6\xed\x83\x64\x6b\xc6\x05\x9b\x0b\xa4\x70\xe2\xd5\xa1\xc0\xcc\x29\x13\x79\x0a\x5f\x35\xd7\x6c\x8e\xc2\x6e\x84\x98\xd6\x4f\x84\xe1\xb0\xd0\x62\x6b\xa2\x19\xbf\xff\x44\x4b\xd3\x73\xba\x00\x36\xd1\xfb\x4f\x1b\xae\x0c\x77\xd5\x99\x2f\xf6\x41\x00\x33\x8d\x20\x11\x3f\x27\x48\x66\xb8\xe3\x19\x13\x69\xcd\x6f\x5b\xb9\x1f\xbc\x2c\xf1\x01\x4a\x25\xd0\x84\xc2\x6c\x78\x0c\x40\xe0\x16\x2b\x0a\xe9\x59\x6d\xaf\x97\xe7\x4a\xda\xa1\x14\x55\xda\xe0\x02\x50\xda\x4b\x2b\x43\x21\xed\x7f\xe3\xd6\xd9\x74\x8f\x92\xe0\xb9\x2f\xde\x8e\x4f\xba\x91\xe8\x30\xf4\x5e\xa6\xa4\x33\x4a\x10\x2d\x98\xc4\x17\xe8\x05\xc0\xc5\x02\x33\x47\x21\x1d\xa8\x49\xb6\xc2\xbc\x14\xf8\x12\xc3\x05\xf3\x2d\xf7\x77\x59\xf4\x61\x30\x2e\xd1\x6c\x11\x24\xcf\xf5\x41\xfc\x78\xc1\x96\x48\xe1\xe8\x6e\xf2\x71\x32\xed\xdf\xcc\xce\xfb\x17\xbd\x0f\xd7\xd3\xd9\xb8\x7f\x79\x35\x99\x8e\x3f\xde\x1f\x19\x26\xb3\x15\x9a\xee\x7e\x45\x74\x7d\xdc\x39\xee\xfc\xf4\xa6\xad\x70\x54\x0a\x31\x52\x82\x67\x15\x85\xab\xc5\x40\xb9\x91\x41\x8b\xdb\x84\x7b\x7f\x8b\x82\xc9\xfc\x21\xdd\xe4\x39\x47\x09\x58\xc7\x8c\x6b\x9c\x09\x89\x3b\xa9\x41\xea\xa2\xcb\xba\x91\x5a\xff\x3a\x5f\xac\x92\x5b\x8e\xb8\x5d\x6e\x7c\xed\xd9\xa6\xed\x08\x55\x94\x20\x91\xa9\x81\x7c\xe1\xf9\x47\xcc\xad\x68\xcb\xc0\x96\x03\xe5\xfa\xb1\xb2\xd1\xf0\x7c\x36\xe8\xdd\xf4\x27\xa3\xde\x59\xbf\xa1\x6c\xcd\x44\x89\x17\x46\x15\xb4\x95\xdb\x05\x47\x91\xd7\xa3\xfb\x11\x3d\xda\xde\xf4\x78\x67\x3b\xc1\x92\x66\x54\x2f\x08\x28\xd2\x6f\x98\x6e\x5b\x7b\x54\x30\x35\xbe\xbb\x53\xb8\xbd\x2c\x1f\xe6\xf1\x24\xd2\xc3\xdc\x78\x72\x22\xfb\xf5\x24\xa5\x72\xcd\x9e\x6f\x6e\xd8\x9d\x56\xe1\x96\xe4\xb8\x60\xa5\x70\x24\x5c\x53\x48\x9d\x29\x31\x4d\x9a\x75\x08\x75\x9d\x7a\x81\x86\xa5\x18\x7b\xbd\x4d\x6f\x54\x8e\x14\x7e\x67\xdc\x5d\x28\x73\xc1\x8d\x75\x67\x4a\xda\xb2\x40\x93\x98\xf8\xd4\xd9\x14\xed\x39\x0a\x74\x18\x22\xaf\x57\xe4\x06\xb2\x64\xe7\xd9\xf8\xe4\xe6\xd9\x16\xe8\x0f\x96\xce\x46\xb0\x51\xab\x14\xfe\x20\x01\x90\xbb\x3a\x37\x61\x82\xf8\x0a\xb8\x61\x3a\xa5\x9f\x6a\xea\xdd\x36\x73\xe1\x3e\xa5\xe9\xa6\x73\x47\xbd\xe9\xfb\xd9\xc5\x70\x3c\x1b\x0c\x07\xb3\xeb\xab\xc9\xb4\x7f\x3e\x1b\x0c\xcf\xfb\x93\xf4\xf5\x83\x8c\xf7\xce\xa6\xf4\x53\x7a\x74\xb7\x91\xbb\x1e\x9e\xf5\xae\x67\x93\xe9\x70\xdc\xbb\xec\x07\x2d\xf7\x47\xe1\xa1\xe3\xbf\xfb\xfa\x1f\xcf\xf7\x61\x7d\x39\xff\xb8\xa8\x9d\xfd\xef\x7f\xba\x73\x2e\xbb\x76\x15\x4e\x5f\x57\x5c\x20\x2c\xd1\x29\xed\x2c\xa4\x05\xb5\x54\xd3\x14\x94\x8e\xed\x9b\xab\x87\x39\xc0\x2c\xc2\x2b\xa5\x1d\x70\xd9\xaa\x45\xfd\xbf\xd6\x91\xcd\xad\x12\xa5\x0b\x38\xfc\xfa\x6a\x38\x9a\xf6\xc6\x97\x2d\x86\xb7\x6f\x5b\x47\xdb\x16\xb7\xfc\x3b\x5e\xc9\x77\x95\x43\x7b\x88\x74\xd1\x96\x5e\x2b\xe1\x2b\xe7\x39\x49\xb4\x2c\xab\xe3\x93\xb1\xdb\x8a\xdb\x9c\x1b\x20\x05\x1c\x9f\x9e\x9e\x02\xd1\xf0\xea\xae\x19\x48\x04\x35\x5b\x15\x2a\x87\xd3\xe3\xe3\xdd\xdb\x6e\xa7\x13\xf6\x3c\x33\xb9\xfa\x2a\xff\x85\xfa\x49\xa8\x4d\x01\xc4\x2c\xf6\x00\xbc\x42\xa1\xd1\x8c\x54\xde\xa9\x58\x21\xb6\x28\xee\x74\xb1\x27\xc5\x46\x1f\xa9\x7c\xef\x8b\x2a\xf6\x76\xd4\x46\x74\xcd\xd4\x7c\x36\xfd\x78\x05\xef\x08\xc1\x8b\xd6\x6e\xc1\x8d\x51\x06\x73\x22\xf8\xdc\x30\x53\x91\x79\x69\xab\xb9\xfa\x46\x4f\x3a\x3f\xbf\xe9\x9c\x1c\xb8\x77\xff\x0c\x00\x00\xff\xff\x7c\x3e\x44\xe7\xec\x0e\x00\x00")

func localStorageYamlBytes() ([]byte, error) {
	return bindataRead(
		_localStorageYaml,
		"local-storage.yaml",
	)
}

func localStorageYaml() (*asset, error) {
	bytes, err := localStorageYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "local-storage.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAggregatedMetricsReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcf\x31\x6b\xf4\x30\x0c\xc6\xf1\xdd\x9f\x42\x78\x7e\x93\x97\x6e\xc5\x6b\x87\xee\x1d\xba\x94\x1b\x94\xf8\x21\x27\xce\xb1\x83\x24\xe7\x68\x3f\x7d\xb9\x70\xdc\x58\x68\x27\x0d\x7f\x7e\x0f\xe8\x22\x35\x27\x7a\x29\xdd\x1c\xfa\xd6\x0a\x02\x6f\xf2\x0e\x35\x69\x35\x91\x4e\x3c\x8f\xdc\xfd\xdc\x54\xbe\xd8\xa5\xd5\xf1\xf2\x6c\xa3\xb4\xff\xfb\x53\x58\xe1\x9c\xd9\x39\x05\xa2\xca\x2b\x12\xd9\xa7\x39\xd6\xc4\xcb\xa2\x58\xd8\x91\x87\x15\xae\x32\xdb\xa0\xe0\x0c\x0d\x44\x85\x27\x14\xbb\x11\xfa\x61\xfd\xb1\x30\x78\x1b\x76\xc1\x35\x51\x74\xed\x88\xbf\x71\xc8\xe2\x7f\x71\x9c\x57\xa9\x0f\xa8\xbd\xc0\x52\x18\x88\x37\x79\xd5\xd6\x37\x4b\xf4\x11\xef\x7f\xdd\x7d\x3c\x05\x22\x85\xb5\xae\x33\x8e\xbe\xb5\x6c\xf1\x1f\xc5\xda\x32\xec\xc8\x3b\x74\x3a\xd2\x02\xbf\x95\x22\x76\xdc\x2b\xfb\x7c\x8e\xa7\xf0\x1d\x00\x00\xff\xff\xe5\x1d\x7a\x17\x89\x01\x00\x00")

func metricsServerAggregatedMetricsReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAggregatedMetricsReaderYaml,
		"metrics-server/aggregated-metrics-reader.yaml",
	)
}

func metricsServerAggregatedMetricsReaderYaml() (*asset, error) {
	bytes, err := metricsServerAggregatedMetricsReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/aggregated-metrics-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAuthDelegatorYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\x31\x4e\xc4\x40\x0c\x45\xfb\x39\xc5\x5c\x60\x82\xe8\x90\x3b\xa0\xa0\x5f\x24\x7a\x67\xf2\x59\x4c\x92\x71\x64\x7b\x22\xc1\xe9\xd1\x8a\x88\x06\xd8\xf6\x4b\xff\xbd\x57\x4a\x49\xbc\xc9\x0b\xcc\x45\x1b\x65\x1b\xb9\x0e\xdc\xe3\x4d\x4d\x3e\x39\x44\xdb\x30\xdf\xf9\x20\x7a\xb3\xdf\xa6\x59\xda\x44\xf9\x71\xe9\x1e\xb0\x93\x2e\x78\x90\x36\x49\x3b\xa7\x15\xc1\x13\x07\x53\xca\xb9\xf1\x0a\xca\x2b\xc2\xa4\x7a\x71\xd8\x0e\x23\xff\xf0\xc0\x4a\x17\x70\x99\xb0\xe0\xcc\xa1\x96\x4c\x17\x9c\xf0\x7a\x79\xf1\x26\x4f\xa6\x7d\xbb\x52\x90\x72\xfe\x15\xf0\xe3\xfb\x5b\xe0\x7d\x7c\x47\x0d\xa7\x54\x8e\xef\x33\x6c\x97\x8a\xfb\x5a\xb5\xb7\xf8\x27\xf7\x98\x7d\xe3\x0a\xca\x73\x1f\x51\xbe\xf9\xe9\x2b\x00\x00\xff\xff\xa5\xb5\x26\x22\x2f\x01\x00\x00")

func metricsServerAuthDelegatorYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAuthDelegatorYaml,
		"metrics-server/auth-delegator.yaml",
	)
}

func metricsServerAuthDelegatorYaml() (*asset, error) {
	bytes, err := metricsServerAuthDelegatorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/auth-delegator.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerAuthReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\x3d\x4e\x04\x31\x0c\x46\xfb\x9c\x22\x17\xf0\x22\x3a\x94\x0e\x1a\xfa\x45\xa2\xf7\x64\x3e\xc0\xcc\x8e\x13\xd9\xce\x08\x38\x3d\x1a\xb4\xfc\x34\x4b\x6f\xbf\xef\x3d\x22\x4a\xdc\xe5\x11\xe6\xd2\xb4\x64\x9b\xb8\x1e\x78\xc4\x4b\x33\xf9\xe0\x90\xa6\x87\xe5\xc6\x0f\xd2\xae\xb6\xeb\xb4\x88\xce\x25\x1f\xdb\x09\x77\xa2\xb3\xe8\x73\x5a\x11\x3c\x73\x70\x49\x39\x2b\xaf\x28\x79\x45\x98\x54\x27\x87\x6d\x30\xda\x51\x64\xe0\x19\x76\x3e\xf1\xce\x15\x25\x2f\x63\x02\xf9\xbb\x07\xd6\x64\xed\x84\x23\x9e\x76\x08\x77\xb9\xb7\x36\xfa\x3f\x26\x29\xe7\x5f\x91\x9f\x5d\xbc\x05\x74\x6f\x20\xee\xf2\x67\x1c\x1a\x52\xbf\xde\xbf\x35\x7c\x4c\xaf\xa8\xe1\x25\xd1\x19\xf4\x00\xdb\xa4\xe2\xb6\xd6\x36\x34\x2e\xa4\x5c\xd6\xff\x0c\x00\x00\xff\xff\x2a\x39\xe6\xe4\x44\x01\x00\x00")

func metricsServerAuthReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerAuthReaderYaml,
		"metrics-server/auth-reader.yaml",
	)
}

func metricsServerAuthReaderYaml() (*asset, error) {
	bytes, err := metricsServerAuthReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/auth-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsApiserviceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8e\x4d\x6a\xc4\x30\x0c\x85\xf7\x3e\x85\x2e\xe0\x74\xb2\x2b\xde\x75\x59\x68\x61\x20\x65\xf6\x1a\x8f\x3a\x88\xe0\x1f\x24\xd9\x90\xdb\x97\x30\x71\xba\x13\x7a\xef\xfb\x24\xef\xbd\xc3\xca\x37\x12\xe5\x92\x03\x60\x65\xa1\x27\xab\x09\x1a\x97\x3c\xad\xef\x3a\x71\x79\xeb\xb3\x5b\x39\x3f\x02\x7c\x5c\x3f\x17\x92\xce\x91\x5c\x22\xc3\x07\x1a\x06\x07\x90\x31\x51\x80\x3e\xdf\xc9\x70\x9e\x12\x99\x70\xd4\x03\x76\x5a\x29\xee\x25\x7d\x81\xfb\x38\x88\xa3\xe9\xf7\x88\xe4\x0c\xb4\x62\xa4\x00\x6b\xbb\x93\xd7\x4d\x8d\x92\x03\x78\x4a\x69\xf5\x44\x86\x1c\xa0\x8f\xdf\x8f\xf3\x0e\x80\xb3\x52\x6c\x42\xcb\xca\xf5\xe7\x6b\xb9\x91\xf0\xef\x16\xc0\xa4\xd1\x10\x5d\x85\x8b\xb0\x6d\xdf\x9c\x39\xb5\x14\x60\xbe\x5c\xfe\x65\x23\x7d\xad\xff\x02\x00\x00\xff\xff\x14\x74\xa9\x1b\x25\x01\x00\x00")

func metricsServerMetricsApiserviceYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsApiserviceYaml,
		"metrics-server/metrics-apiservice.yaml",
	)
}

func metricsServerMetricsApiserviceYaml() (*asset, error) {
	bytes, err := metricsServerMetricsApiserviceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-apiservice.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsServerDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\xc1\x6e\xdb\x46\x10\xbd\xeb\x2b\x06\x2a\x7c\x2b\x6d\x49\xa9\xdd\x80\x80\x0f\x82\xc4\x44\x01\x6c\x47\x10\xe5\x16\x3e\x09\xeb\xe5\xc8\x5a\x78\xb9\xbb\x9d\x19\x2a\x61\x83\xfc\x7b\xb1\xa2\x43\x93\x8e\x1d\xb8\x68\xc3\x03\x0f\x33\x6f\xde\x3c\xbe\x9d\xe5\x24\x49\x32\x50\xc1\xfc\x81\xc4\xc6\xbb\x14\xf6\xe3\xc1\xbd\x71\x45\x0a\x39\xd2\xde\x68\x9c\x6a\xed\x2b\x27\x83\x12\x45\x15\x4a\x54\x3a\x00\x70\xaa\xc4\x14\x4a\x14\x32\x9a\x13\x46\xda\x23\x3d\x84\x39\x28\x8d\x29\xdc\x57\xb7\x98\x70\xcd\x82\xe5\xe0\x69\x07\x15\x02\x9f\xb4\x6d\xe6\x18\xac\xaf\x4b\xfc\x4f\x2d\x00\xac\xba\x45\xcb\xb1\x12\xe0\xfe\x2d\x27\x2a\x84\xef\xca\x39\xa0\x8e\x08\xc2\xbd\x89\x52\x16\x86\xc5\x53\x7d\x61\x4a\x23\x29\x8c\x06\x00\x2c\xa4\x04\xef\xea\x86\x47\xea\x80\x29\xac\xbc\xb5\xc6\xdd\x5d\x87\x42\x09\x1e\xe2\xd4\x8d\x34\x50\x80\x52\x7d\xbe\x76\x6a\xaf\x8c\x55\xb7\x16\x53\x18\x47\x3a\xb4\xa8\xc5\x53\x83\x29\x95\xe8\xdd\x45\x47\xe7\xcb\x4a\x01\x04\xcb\x60\x5b\xfa\xae\x33\xf1\x79\xc1\x9d\xf8\xd8\x5e\x83\x1f\xb5\x00\xf8\x66\x48\x7c\x02\x19\x4f\x46\xea\x99\x55\xcc\x57\x07\xfe\x61\xe3\x6e\xe2\x7c\x81\x89\x26\x23\x46\x2b\x3b\x7c\xc0\x73\x6f\x3c\xae\x5e\x16\x24\xde\x22\x29\x31\xde\x75\x54\x25\x70\x8f\x75\x0a\xc3\xd9\x03\xeb\xb4\x28\xbc\xe3\x8f\xce\xd6\xc3\x16\x03\xe0\x43\xac\xf4\x94\xc2\x30\xfb\x6c\x58\x78\xf8\x1d\xc1\x41\x1b\x79\x8b\xc7\x71\x1e\xc8\xa1\x20\x1f\x1b\x7f\xa2\xbd\x13\xf2\x36\x09\x56\x39\x7c\x25\x27\x00\x6e\xb7\xa8\x25\x85\xe1\x95\xcf\xf5\x0e\x8b\xca\xe2\xeb\x5b\x96\x8a\x05\xe9\xff\xe8\xb5\xf7\xb6\x2a\xb1\xb5\xeb\x17\x28\xa3\xc7\x60\x1c\x48\x19\x80\x3d\x7c\x42\xd0\xca\x01\xab\x2d\xda\x1a\x2a\x46\xd8\x92\x2f\x13\xd6\x14\x67\x0c\x4c\xa9\xee\x90\x41\xb9\xe2\xc4\x13\x10\xaa\x22\xf1\xce\xd6\x10\x4d\x51\xc6\x21\xf1\xe0\xdb\x27\x35\x93\x24\x65\x48\x0a\x43\xad\x3a\x2c\x83\xd4\x73\x43\x29\x7c\xf9\xfa\x10\x7c\xac\x4d\x9f\x14\x3f\x7b\xea\xd0\x88\x48\xe1\xe8\x4b\x7e\x93\xaf\xb3\xcb\xcd\x3c\x7b\x37\xbd\xbe\x58\x6f\x56\xd9\xfb\x0f\xf9\x7a\x75\xf3\xf5\x88\x94\xd3\x3b\xa4\x93\xd2\x10\x79\xc2\x22\xe9\x33\xa5\xfb\xd1\xf1\xd9\xf1\x9b\x96\x50\xd1\x5d\x6f\x82\x92\x44\x23\x49\xd4\x7d\x7e\x22\x65\xe8\x65\x18\x75\x45\x98\x04\x4f\x72\x3e\x1e\x4d\x4e\x47\xbd\x6c\x3c\x37\x8b\x92\x04\xc2\x2d\x52\xec\xac\x8a\x82\x90\x39\x89\x57\x9e\xcf\x8f\xbe\x2c\x57\xd9\xbb\x6c\xb5\xca\xe6\x9b\xe9\x7c\xbe\xca\xf2\x7c\xb3\xbe\x59\x66\xf9\xd7\xa3\x67\x79\x2a\xc6\xe6\x92\xb0\x28\xa9\xf8\xd0\xb6\x07\x6c\x3e\x2c\x21\x64\x6f\xab\x78\x15\xce\xc7\xa7\xdc\x43\x88\xe5\x44\x9b\xb0\x43\x4a\xb8\x32\x82\x7c\xbe\xbe\xc8\x37\xd9\x6c\xbe\xc8\xe2\x3b\x9f\x6e\xfe\xfc\xb0\x5e\x6c\xa6\x59\xbe\x99\x9c\x9e\x6d\xde\xcf\x2e\x37\xf9\x62\xfa\xe6\xed\x6f\xbf\x3e\xe2\x56\xaf\x42\x3d\x61\x1b\x4f\xde\x7e\xc3\x4d\x4e\xcf\x5e\x62\x7b\x11\xd5\x61\x9b\x2d\xa6\xb3\xc5\x74\x32\xda\x2c\x3f\x5e\xdc\x8c\xdf\x8c\x4e\x9f\x23\xfb\x0e\xd4\xba\x10\xcd\xa9\x48\x63\xe7\x8c\x63\xf0\xaf\x0a\x59\x7a\x31\x00\x1d\xaa\x14\xc6\xa3\x51\xd9\x8b\x96\x58\x7a\xaa\x53\xf8\x7d\x74\x69\xda\x44\x3c\x8a\xde\xd4\x34\x33\xbb\x13\x09\xdc\xa9\x6e\xa7\x7b\xe9\x49\x22\x77\x77\x64\xe2\xcf\xd1\x8b\xd7\xde\xa6\xb0\x9e\x2d\x3b\x8a\x55\x61\x1c\x32\x2f\xc9\xdf\x62\x57\x62\xa4\x7f\x8f\xd2\x57\x1d\x94\xec\x52\x38\x89\x55\xf5\xdf\xfd\xcc\xa1\xe9\x53\x4d\x00\xac\x77\x18\xd5\x2e\xd6\xeb\x65\xde\xc9\x18\x67\xc4\x28\x3b\x47\xab\xea\x1c\xb5\x77\x05\x37\xfb\xab\x25\x44\x32\xbe\x68\x53\x93\x4e\x4a\x4c\x89\xbe\x92\x36\x37\xee\xe4\xb8\xd2\x1a\x99\xd7\x3b\x42\xde\x79\x5b\xf4\xb3\x5b\x65\x6c\x45\xd8\xc9\x3e\xde\x4d\x6b\xf6\xf8\xaf\x9d\x88\x45\x3f\xc1\x88\xb3\x1f\x38\x31\x1e\xfd\x74\x2b\x0e\xbf\x9e\xb8\x48\xbd\x13\xfc\x2c\xfd\x69\x56\x45\xdc\x71\x2b\xef\xe5\x9d\xb1\xd8\xec\xd7\x14\x84\x2a\xec\xc2\x2a\x37\xe5\x2b\xef\x22\xec\xf9\xe4\x35\x23\x1d\x6e\x40\xf7\x73\x94\xb5\xfe\xd3\x92\xcc\xde\x58\xbc\xc3\x8c\xb5\xb2\x87\xb5\x9b\xc2\x56\x59\x7e\xe4\x68\xb6\xcb\x65\x5c\x29\xcf\xdc\x8c\xa7\xab\x00\x9a\xe5\xb3\x6c\x8e\x2c\xfe\x67\xff\x09\x00\x00\xff\xff\x70\xb0\x51\x48\x32\x0a\x00\x00")

func metricsServerMetricsServerDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsServerDeploymentYaml,
		"metrics-server/metrics-server-deployment.yaml",
	)
}

func metricsServerMetricsServerDeploymentYaml() (*asset, error) {
	bytes, err := metricsServerMetricsServerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-server-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerMetricsServerServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x3f\x4b\x04\x31\x10\xc5\xfb\x7c\x8a\x61\xfb\x28\xe2\x15\x92\xd6\x5a\x38\x50\xec\xe7\x72\x0f\x0d\x97\x4d\xc2\xcc\xec\x82\xdf\x5e\x76\xf6\x9a\x83\xed\x92\x37\xef\xcf\x2f\xc6\x18\x78\x94\x6f\x88\x96\xde\x12\xad\x2f\xe1\x56\xda\x35\xd1\x27\x64\x2d\x19\x61\x86\xf1\x95\x8d\x53\x20\x6a\x3c\x23\xd1\x0c\x93\x92\x35\x2a\x64\x85\xdc\x65\x1d\x9c\x91\xe8\xb6\x5c\x10\xf5\x4f\x0d\x73\x20\xaa\x7c\x41\xd5\x2d\x49\x7e\x91\x06\x83\x3e\x95\xfe\xbc\x37\x4d\x1f\x0f\x55\xd3\x81\x31\xd7\x45\x0d\xe2\x8e\xb2\x2d\x4c\x26\x0b\xa6\xa0\x03\x79\x2b\x56\x54\x64\xeb\x72\x1f\x79\xd3\xc8\x63\x1c\x30\x8e\x2e\xe6\x24\xd1\x9f\x89\x4e\xa7\x57\x8f\xec\x24\xbf\x66\x43\xfd\x3f\xa4\x5b\xcf\xbd\x26\xfa\x7a\x3f\xbb\x62\x2c\x3f\xb0\xb3\xa7\x76\xdf\x7f\x00\x00\x00\xff\xff\x7e\x3b\x1f\x83\x35\x01\x00\x00")

func metricsServerMetricsServerServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerMetricsServerServiceYaml,
		"metrics-server/metrics-server-service.yaml",
	)
}

func metricsServerMetricsServerServiceYaml() (*asset, error) {
	bytes, err := metricsServerMetricsServerServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/metrics-server-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricsServerResourceReaderYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x41\x4b\xc4\x30\x10\x85\xef\xf9\x15\xc3\xde\xd3\xc5\x9b\xe4\xa6\x1e\xbc\xaf\xe0\x3d\x4d\x9f\xbb\x63\xdb\xa4\xcc\x4c\x2a\xfa\xeb\xa5\xb6\x2a\xb8\xb0\x2c\x78\x4a\x98\xe4\x7d\x8f\xf9\xbc\xf7\x2e\x4e\xfc\x0c\x51\x2e\x39\x90\xb4\x31\x35\xb1\xda\xa9\x08\x7f\x44\xe3\x92\x9b\xfe\x56\x1b\x2e\xfb\xf9\xc6\xf5\x9c\xbb\x40\x0f\x43\x55\x83\x1c\xca\x00\x37\xc2\x62\x17\x2d\x06\x47\x94\xe3\x88\x40\xfa\xae\x86\x31\x8c\x30\xe1\xa4\x5e\x21\x33\xc4\x49\x1d\xa0\xc1\x79\x8a\x13\x3f\x4a\xa9\x93\x2e\x09\x4f\xbb\x9d\x23\x12\x68\xa9\x92\xb0\xcd\x72\xe9\xa0\xfb\x0d\xe0\x88\x66\x48\xbb\x3d\x1d\x61\xd7\x31\xa6\xd2\xe9\x2f\xec\x1c\xb2\x9c\x03\xeb\x7a\x79\x8b\x96\x4e\xee\x7f\x26\xee\x39\x77\x9c\x8f\xd7\x0b\x29\x03\x0e\x78\x59\xbe\x7d\xaf\x73\xa1\xd2\x11\x9d\xbb\xbf\x5c\xa0\xb5\x7d\x45\xb2\x2f\xe9\x6b\xf6\x09\x32\x73\xc2\x5d\x4a\xa5\x66\xfb\x89\xff\xc9\xad\x63\x9d\x62\x42\xa0\xbe\xb6\xf0\x2b\xdf\x7d\x06\x00\x00\xff\xff\xdb\x55\x9e\x61\x2a\x02\x00\x00")

func metricsServerResourceReaderYamlBytes() ([]byte, error) {
	return bindataRead(
		_metricsServerResourceReaderYaml,
		"metrics-server/resource-reader.yaml",
	)
}

func metricsServerResourceReaderYaml() (*asset, error) {
	bytes, err := metricsServerResourceReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metrics-server/resource-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rolebindingsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\x31\x6f\xe3\x30\x0c\x85\x77\xfd\x0a\x21\xbb\x72\x38\xdc\x72\xf0\xd8\x0e\xdd\x03\xb4\xbb\x2c\xb1\x09\x6b\x59\x14\x48\x39\x41\xfb\xeb\x0b\xc7\x4e\xd2\xc4\x76\xe0\xb4\xe9\x66\x0b\xe2\xfb\x48\xbe\x07\xd9\x84\x2f\xc0\x82\x14\x0b\xcd\xa5\x75\x4b\xdb\xe4\x0d\x31\x7e\xd8\x8c\x14\x97\xd5\x7f\x59\x22\xfd\xd9\xfe\x55\x15\x46\x5f\xe8\xc7\xd0\x48\x06\x5e\x51\x80\x07\x8c\x1e\xe3\x5a\xd5\x90\xad\xb7\xd9\x16\x4a\xeb\x68\x6b\x28\x74\xd5\x94\x60\x6c\x42\x01\xde\x02\x9b\xf6\x37\x40\x36\xd6\xd7\x18\x15\x53\x80\x15\xbc\xb6\xb7\x6d\xc2\x27\xa6\x26\x5d\x21\x2b\xad\x07\xe0\x23\x47\xde\x25\x43\x5d\x1c\xf5\x13\xf6\x0c\x69\xca\x37\x70\x59\x0a\x65\x6e\x82\x3c\x0b\xf0\xc4\x14\x4a\x19\x63\xd4\xf7\xb7\x35\xb2\xa6\x43\xfb\xff\xc4\x38\x8a\x99\x29\x04\x60\xc5\x4d\x80\xb3\xc6\xa5\xad\x30\x7a\xb1\x50\x5a\x33\x08\x35\xec\xa0\x3f\x8b\xe4\x41\x94\xd6\x5b\xe0\xb2\x3f\x5a\x43\x9e\x59\x6b\x6b\x90\x64\xdd\xa5\x40\x40\xc9\xfb\x8f\x9d\xcd\x6e\x33\xa2\x15\x21\xef\x88\x2b\x8c\xeb\x7e\xde\x31\xf1\xee\x4e\xa2\x80\x0e\xf7\x04\xa3\x5d\xb7\x0c\x87\x9e\x6f\x45\x8e\x10\x20\xfa\x44\x18\x73\xa7\x9d\xc8\x4f\x69\xb6\x0b\x39\x69\xff\xd0\xc5\xe9\xcc\x4f\x98\x79\xff\xb0\x9f\x03\x4e\x49\x6f\x67\x9c\xc7\xb8\x48\xfb\x75\xc0\xfd\x63\xff\x35\x07\xa6\x4d\xf0\x64\xe4\x07\x49\x1b\xc6\x60\x76\xa8\x7e\xcd\xf8\x91\x71\xee\x67\xfa\x50\xfc\xdc\xf0\xae\x72\x8f\x18\x3a\x79\x78\x1d\xe6\xb5\xf1\x19\x00\x00\xff\xff\x20\xa2\xda\xb0\x09\x06\x00\x00")

func rolebindingsYamlBytes() ([]byte, error) {
	return bindataRead(
		_rolebindingsYaml,
		"rolebindings.yaml",
	)
}

func rolebindingsYaml() (*asset, error) {
	bytes, err := rolebindingsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rolebindings.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _traefikYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x91\x5f\x6f\xdb\x3a\x0c\xc5\xdf\xfd\x29\x08\x03\x79\xba\x90\xdd\xe4\xa9\xd7\x6f\xb9\xa9\x7b\x57\x6c\xeb\x8a\x38\xdd\xd0\xa7\x80\x91\x99\x58\x88\x2c\x09\x14\x1d\x2c\xeb\xfa\xdd\x07\x25\xe9\x3f\xa0\xc0\x86\x61\x7b\x13\x44\xf2\x77\xc8\x73\x94\x52\x19\x06\xf3\x99\x38\x1a\xef\x2a\xe8\xc8\xf6\x85\x46\x11\x4b\x85\xf1\xe5\x6e\x9c\x6d\x8d\x6b\x2b\x78\x47\xb6\x9f\x75\xc8\x92\xf5\x24\xd8\xa2\x60\x95\x01\x38\xec\xa9\x02\x61\xa4\xb5\xd9\x2a\xcd\xed\xe9\x2f\x06\xd4\x54\xc1\x76\x58\x91\x8a\xfb\x28\xd4\x67\x31\x90\x4e\x23\x3a\x41\x2a\xe8\x44\x42\xac\xca\x72\x74\xff\xfe\xf6\xbf\x7a\x7e\x5d\x2f\xea\x66\x39\xbd\xb9\x7a\x18\x95\x51\x50\x8c\x2e\x0f\x8d\xb1\x7c\x01\x57\x93\x71\x31\x29\xc6\xff\x0c\xe1\xf0\x38\x2b\x64\xf3\x2d\xfb\x83\x07\xfc\xbd\xe5\xdf\x5a\x1c\x20\x92\x24\x28\xc0\xc6\xfa\x15\xda\xe2\x28\x76\x41\x6b\x1c\xac\xcc\x69\x63\xa2\xf0\xbe\x82\x7c\x74\xdf\xdc\x35\x8b\xfa\xe3\xf2\xa2\xbe\x9c\xde\x7e\x58\x2c\xe7\xf5\xff\x57\xcd\x62\x7e\xb7\x9c\x4f\xbf\x3c\x8c\xf2\x0c\x60\x87\x76\xa0\x38\xf3\x4e\xc8\x49\x05\xdf\xd5\x81\x1b\x7c\x3b\x75\xce\xa7\x95\xbc\x8b\x47\x2d\x80\xc0\xbe\x27\xe9\x68\x88\xc9\xa0\xe0\xd3\x45\xf9\xf9\xd9\xf9\x24\x7f\xb3\x21\x6a\xc6\x40\x15\xe4\xc2\x03\x1d\x5b\x02\xfb\x9d\x69\x89\x9f\x90\xc9\x2b\x76\x24\x14\xaf\xdc\x86\x29\x3e\x15\x00\xc2\xb0\xb2\x26\x76\xd4\x36\xc4\x3b\xa3\xe9\xb9\x02\x40\x0e\x57\x96\xda\x14\xc0\x40\x27\xb2\xf1\x6c\x64\x3f\xb3\x18\xe3\xf5\x21\x9c\xfc\x68\x8b\xd2\x76\x88\x42\xac\x34\x1b\x31\x1a\xed\x71\x15\xd3\xe3\xe6\x89\xc9\x14\x7c\x34\xe2\x0f\xae\x31\x3a\xdd\x11\x97\xbd\x61\xf6\x4c\xad\xb2\x66\xc5\xc8\x7b\x75\x0a\xe5\xf1\x5a\xc1\x4d\x05\xf9\xa4\xf8\xb7\x18\x9f\x1d\xff\xc4\x5b\xe2\x97\x9e\x29\xd8\x52\x42\xce\x4e\xd2\xd3\xb6\xf5\x2e\x7e\x72\x76\xff\x08\xf1\x21\x4d\x78\xae\x20\xaf\xbf\x9a\x28\x31\x7f\x35\xe8\x7c\x4b\x8a\xbd\xa5\xe2\xd9\xa9\xe4\xad\xf6\x4e\xd8\x5b\x15\x2c\x3a\xfa\x09\x0b\x80\xd6\x6b\xd2\x29\xac\x6b\xdf\xe8\x8e\xda\xc1\xd2\xaf\xc9\xf4\x98\x9c\xfb\x7d\x7e\x7c\x1d\x9d\x09\x97\xd8\x1b\xbb\xbf\xf1\xd6\xe8\xa4\x7b\xc3\xb4\x26\xbe\x18\xd0\x36\x82\x7a\x9b\x67\x3f\x02\x00\x00\xff\xff\x12\x80\xc2\x85\x56\x04\x00\x00")

func traefikYamlBytes() ([]byte, error) {
	return bindataRead(
		_traefikYaml,
		"traefik.yaml",
	)
}

func traefikYaml() (*asset, error) {
	bytes, err := traefikYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "traefik.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"ccm.yaml":           ccmYaml,
	"coredns.yaml":       corednsYaml,
	"local-storage.yaml": localStorageYaml,
	"metrics-server/aggregated-metrics-reader.yaml": metricsServerAggregatedMetricsReaderYaml,
	"metrics-server/auth-delegator.yaml":            metricsServerAuthDelegatorYaml,
	"metrics-server/auth-reader.yaml":               metricsServerAuthReaderYaml,
	"metrics-server/metrics-apiservice.yaml":        metricsServerMetricsApiserviceYaml,
	"metrics-server/metrics-server-deployment.yaml": metricsServerMetricsServerDeploymentYaml,
	"metrics-server/metrics-server-service.yaml":    metricsServerMetricsServerServiceYaml,
	"metrics-server/resource-reader.yaml":           metricsServerResourceReaderYaml,
	"rolebindings.yaml":                             rolebindingsYaml,
	"traefik.yaml":                                  traefikYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"ccm.yaml":           &bintree{ccmYaml, map[string]*bintree{}},
	"coredns.yaml":       &bintree{corednsYaml, map[string]*bintree{}},
	"local-storage.yaml": &bintree{localStorageYaml, map[string]*bintree{}},
	"metrics-server": &bintree{nil, map[string]*bintree{
		"aggregated-metrics-reader.yaml": &bintree{metricsServerAggregatedMetricsReaderYaml, map[string]*bintree{}},
		"auth-delegator.yaml":            &bintree{metricsServerAuthDelegatorYaml, map[string]*bintree{}},
		"auth-reader.yaml":               &bintree{metricsServerAuthReaderYaml, map[string]*bintree{}},
		"metrics-apiservice.yaml":        &bintree{metricsServerMetricsApiserviceYaml, map[string]*bintree{}},
		"metrics-server-deployment.yaml": &bintree{metricsServerMetricsServerDeploymentYaml, map[string]*bintree{}},
		"metrics-server-service.yaml":    &bintree{metricsServerMetricsServerServiceYaml, map[string]*bintree{}},
		"resource-reader.yaml":           &bintree{metricsServerResourceReaderYaml, map[string]*bintree{}},
	}},
	"rolebindings.yaml": &bintree{rolebindingsYaml, map[string]*bintree{}},
	"traefik.yaml":      &bintree{traefikYaml, map[string]*bintree{}},
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
