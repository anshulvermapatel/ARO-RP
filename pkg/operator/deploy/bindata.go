// Code generated for package deploy by go-bindata DO NOT EDIT. (@generated)
// sources:
// deploy/staticresources/aro.openshift.io_clusters.yaml
// deploy/staticresources/master/deployment.yaml
// deploy/staticresources/master/rolebinding.yaml
// deploy/staticresources/master/service.yaml
// deploy/staticresources/master/serviceaccount.yaml
// deploy/staticresources/namespace.yaml
// deploy/staticresources/worker/deployment.yaml
// deploy/staticresources/worker/role.yaml
// deploy/staticresources/worker/rolebinding.yaml
// deploy/staticresources/worker/serviceaccount.yaml
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

var _aroOpenshiftIo_clustersYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\xcd\x72\xe3\xc8\x0d\xbe\xeb\x29\x50\xce\xc1\x87\x58\xb4\xb5\x93\x99\xd9\xd5\xcd\xe5\xd9\x4d\xa9\xb2\xc9\xba\xc6\xae\xb9\xac\xf7\x00\x36\x21\x0a\x71\xb3\x9b\xe9\x06\xed\xd1\xa4\xf2\xee\x29\x34\x49\x89\x92\x49\xcb\x93\xaa\xf0\xa2\x22\xba\x1b\x3f\x1f\x80\xaf\x21\xce\xe6\xf3\xf9\x0c\x6b\xfe\x42\x21\xb2\x77\x4b\xc0\x9a\xe9\xab\x90\xd3\xb7\x98\x3d\xfe\x18\x33\xf6\x97\x4f\x8b\x9c\x04\x17\xb3\x47\x76\xc5\x12\x6e\x9a\x28\xbe\xfa\x4c\xd1\x37\xc1\xd0\x27\x5a\xb3\x63\x61\xef\x66\x15\x09\x16\x28\xb8\x9c\x01\xa0\x73\x5e\x50\xc5\x51\x5f\x01\x8c\x77\x12\xbc\xb5\x14\xe6\x25\xb9\xec\xb1\xc9\x29\x6f\xd8\x16\x14\x92\x85\xde\xfe\xd3\x55\xf6\x2e\x5b\xcc\xaf\xb2\x1f\xae\x7e\xb8\xba\xfa\xb0\xf8\xf8\xc3\x62\xf1\xe1\xea\xfd\xfc\xc3\xfb\xc5\x4f\x57\xef\xfe\xf2\xf1\xe3\xe2\xc7\xf7\x33\x00\x13\x28\x29\xbf\xe7\x8a\xa2\x60\x55\x2f\xc1\x35\xd6\xce\x00\x1c\x56\xb4\x04\x63\x9b\x28\x14\x62\x86\xc1\x67\xbe\x26\x17\x37\xbc\x96\x8c\xfd\x2c\xd6\x64\xd4\xa3\x32\xf8\xa6\x5e\xc2\x8b\xf5\x56\x43\xe7\x74\x17\x70\xab\x2c\x49\x2c\x47\xf9\xdb\x50\xfa\x2b\x47\x49\x2b\xb5\x6d\x02\xda\xbd\xe9\x24\x8c\xec\xca\xc6\x62\xd8\x89\x67\x00\xd1\xf8\x9a\x86\x5a\x63\x93\x87\x0e\xcd\xce\x6e\x14\x94\x26\x2e\xe1\xdf\xff\x99\x01\x3c\xa1\xe5\x22\x45\xdb\x2e\xaa\xbb\xd7\xb7\xab\x2f\xef\xee\xcc\x86\x2a\x6c\x85\x00\x05\x45\x13\xb8\x4e\xfb\x7a\xe5\xc0\x11\x64\x43\xd0\xee\x84\xb5\x0f\xe9\xb5\x77\x11\xae\x6f\x57\xdd\xe9\x3a\xf8\x9a\x82\x70\xef\x81\x3e\x83\xba\xd8\xc9\x8e\xec\x9c\xab\x23\xed\x1e\x28\xb4\x12\xa8\x35\xd8\xe5\x93\x0a\x88\xad\x69\xbf\x06\xd9\x70\x84\x40\x75\xa0\x48\xae\xad\x8d\x81\x5a\xd0\x2d\xe8\xc0\xe7\xff\x24\x23\x19\xdc\x51\x50\x25\x10\x37\xbe\xb1\x85\x96\xcf\x13\x05\x81\x40\xc6\x97\x8e\xbf\xed\x34\x47\x10\x9f\x4c\x5a\x14\xea\x52\xd1\x3f\xec\x84\x82\x43\xab\x10\x36\x74\x01\xe8\x0a\xa8\x70\x0b\x81\xd4\x06\x34\x6e\xa0\x2d\x6d\x89\x19\xfc\xdd\x07\x02\x76\x6b\xbf\x84\x8d\x48\x1d\x97\x97\x97\x25\x4b\xdf\x09\xc6\x57\x55\xe3\x58\xb6\x97\xa9\x9e\x39\x6f\xc4\x87\x78\x59\xd0\x13\xd9\xcb\xc8\xe5\x1c\x83\xd9\xb0\x90\x91\x26\xd0\x25\xd6\x3c\x4f\x8e\xbb\xd4\x08\x59\x55\xfc\x69\x97\xe8\xf3\x81\xa7\xb2\xd5\x82\x88\x12\xd8\x95\x3b\x71\xaa\xbd\x49\xdc\xb5\x06\x35\xbb\xd8\x1d\x6b\xfd\xdf\xc3\xab\x22\x45\xe5\xf3\xcf\x77\xf7\xd0\x1b\x4d\x29\x38\xc4\x3c\xa1\xbd\x3f\x16\xf7\xc0\x2b\x50\xec\xd6\x14\xda\xc4\xad\x83\xaf\x92\x46\x72\x45\xed\xd9\x49\x57\x49\x4c\xee\x10\xf4\xd8\xe4\x15\x8b\x66\xfa\x5f\x0d\x45\xd1\xfc\x64\x70\x93\xf8\x00\x72\x82\xa6\x2e\x50\xa8\xc8\x60\xe5\xe0\x06\x2b\xb2\x37\x18\xe9\xff\x0e\xbb\x22\x1c\xe7\x0a\xe9\x69\xe0\x87\x34\x76\xb8\xb1\x45\x6b\x27\xee\xa9\x64\x34\x43\x5d\x07\xde\xd5\x64\x0e\x3a\xa3\xa0\xc8\x41\xab\x57\x50\x48\x6b\x7e\xc8\x2e\xd3\xbd\x98\xfa\xd1\x84\x4f\xbe\x42\x76\x87\xe2\x89\x30\xa0\xed\xe0\x95\x93\xd5\xed\xdb\x0f\x7c\xfb\xd9\x3d\x71\xf0\xae\x22\x27\x6f\x3e\x55\x7c\x9f\x57\x6b\x42\xcd\x53\x3c\x3e\x70\x80\xdf\x2f\xdd\xa6\x03\x00\xaf\x3f\xff\xa6\xf4\x17\x50\x7c\xe8\xd5\x40\xa9\x8d\x7f\xa4\x6a\x0a\xc2\xb4\xa6\xdc\x14\x85\x9c\xdc\x06\x5f\x91\x6c\xa8\x19\xd9\xd5\xfb\x9f\x7b\x6f\x09\xdd\x68\x6c\x47\xf5\xa0\x4f\x49\x8e\x9e\xf0\x57\x5f\x96\xec\xca\x63\xad\xaf\x79\x65\xbc\x5b\x73\x39\xc2\xb7\xbb\xc3\x28\xca\x66\x4b\x38\xff\xfd\x6a\xfe\xd3\x1f\x7f\xce\xda\x9f\xf3\x49\xcf\x47\x90\xd7\xa7\xf2\x8e\xc5\xeb\xd2\x5f\x6f\xee\x5e\xc9\xb6\x3e\xe4\x9a\x6a\x4c\x3e\x87\x4f\x8c\xa5\xf3\x51\xd8\xc4\xdb\xe0\x8b\xd1\x3d\xf7\xc7\x7c\x7c\xd2\xbb\x49\x58\xd9\x95\x81\x62\xfc\x8e\x4a\x6e\xb9\x9f\xe4\x66\x43\xe6\x91\xc2\xf7\xa4\xa2\x09\x76\xb4\x20\x58\xa8\x1a\x5d\x38\x81\x78\xbf\x8c\x21\xe0\xf6\xad\x11\x5b\x6f\x06\x77\xfe\x1b\x2c\xf5\x24\xbf\x2a\x5e\xed\xab\x7e\x6e\x5b\x7d\xea\x87\x83\xeb\x6f\xda\x45\xfb\xe3\xed\x5d\x4d\x83\x89\xe5\x4d\xf6\x9f\x1c\xc9\x4b\xdb\x13\xdb\xc7\xf9\xb4\x9d\x7a\x4e\x31\x6a\xda\x75\xc0\xa9\x3e\x8f\x7a\x69\xfd\x4f\xa4\x6a\xbc\x2b\x78\x30\xa7\x4e\x19\xdf\x6d\xeb\x6e\x5d\x92\x64\xa7\x17\x03\xbb\x28\xe8\x0c\xc5\xec\x48\xcd\x44\xd5\x1c\x68\x3f\xdb\xeb\xd9\x5f\xc5\xed\x34\xa4\x91\xa5\x42\x38\x98\x8f\xce\x63\x1b\xeb\xb1\x31\x7d\x06\xae\x62\x20\x3d\xb3\x9b\xe8\xa1\x22\xb3\x41\xc7\xb1\x4a\xfd\xe1\x0a\x2a\x74\x7c\xd2\x6b\x39\x52\x01\xcf\x1b\x72\x0a\xe8\x88\xd2\x82\x04\xd9\xc6\x9d\x13\x7b\xb7\xd4\x86\xde\xed\x08\x75\x60\x1f\x18\x1e\x9d\x7f\x76\xe0\x03\x3c\xa7\xc9\x2d\xad\xd5\xb5\x3d\xae\xfc\x54\x06\x1e\xd0\xda\x3d\x76\x49\x3d\x94\xfc\x44\x0e\x74\xc2\xc9\xe0\xc1\x0d\xe3\xe9\x86\xc1\x9c\x00\x8b\x82\xc6\x58\x47\x3c\xd0\xd7\xda\xb2\x61\xb1\xdb\x76\x6a\xdc\x0e\x72\x0f\xb2\x41\xd1\x60\x43\x4c\xd3\xa0\xf1\x55\xed\x5d\x42\xdb\x24\xb0\x72\xdf\x8c\x71\x56\x40\xd9\xa4\x49\x08\x5d\x1a\x6c\x38\xb4\x03\x96\x8f\x74\xa0\x3d\x61\x99\xa6\x26\xbd\xe3\xd3\xcc\xe4\xf5\xe4\x88\xca\x01\x86\x31\x83\xdf\x9c\xa1\xae\xa6\x8b\x8b\x54\xd4\x15\xa1\x53\x23\x09\x92\x7d\x7d\x18\x74\xd0\x8e\x52\x23\x3a\x35\xb9\x25\x15\x80\x21\x67\x09\x18\xd8\x6e\x61\x0e\xac\xbb\x8d\xaf\x28\x42\x8d\x41\xfa\xfe\xbe\xbe\x5d\xb5\x23\xf1\x06\xdb\x36\x8a\x58\x8d\x29\xcd\xd1\x3c\x3e\x63\x28\xe2\x3c\xed\x5e\xfb\xd0\xbe\x29\x76\x28\x9c\xb3\x65\x49\x50\x1b\x0a\xae\xab\x90\x6d\x1b\x76\xb2\x37\x16\xfb\xce\x83\xec\xec\xe5\xfd\xfc\x0a\x35\x03\x58\x8c\x72\x1f\xd0\x45\xee\xff\xff\x8d\x33\xf2\xda\x87\x0a\x65\x09\x3a\x6d\xce\x85\x47\x23\x3b\xc9\xdb\x15\xc5\x88\xe5\x84\x85\x13\x67\x03\x61\x1c\xbf\xc9\xa7\xa8\xe5\x73\x3a\xa1\xfc\x72\xd4\x9c\x08\xde\xd1\xfc\xd9\x87\xe2\x62\x3f\x33\x8f\x2a\x86\xa3\x3f\x58\x3b\x2e\x47\xa1\xd2\x87\xad\xbe\x1b\x6c\x22\xed\x16\x9a\x10\xc8\x49\xc7\xbd\x63\x74\xa2\xcf\x4a\x46\xbc\x4a\x94\xc1\x2e\x65\x9e\x55\x63\x23\x75\x23\x17\x10\x1b\xb3\x01\x8c\xc9\x67\xcb\x6e\xca\xd1\xc7\x26\x27\x23\x16\x4a\x65\xd2\xee\xa8\xd6\x17\x3b\x88\x4d\x55\x61\xe0\x6f\xa9\xfc\x4d\xeb\x62\xc7\x0e\xc9\xf9\x09\x3f\x4f\x24\xe4\xe5\xf5\xf2\xe6\xa3\x69\xf9\x74\x26\xf7\x34\x7e\xbf\xad\xa9\xbf\x5f\xf5\xf0\x0e\xee\x5d\x1f\xa7\x50\x8f\xc7\xd6\x81\x3d\x36\x68\xed\x56\x5b\xbf\x4f\x78\x01\x5a\x01\x4a\xac\x71\xe3\x83\x40\xbd\x09\xe9\xaf\xd3\x90\x22\x93\xb1\x29\xad\x1d\x7b\xb2\x2b\x58\xeb\xa1\xbb\x2d\x39\x51\x3e\x3c\x9c\x61\xee\xb4\x67\xec\x5c\x42\x43\x0f\x67\x50\x7b\x8b\x81\x65\x9b\xc1\x2f\x7e\x8c\xc0\xf4\xa1\xaf\x58\xd5\x96\x2e\x80\x8f\xe3\xeb\xad\xc4\xf6\x56\x41\x55\xc7\x66\xdb\xd6\x51\xfa\xa4\x71\x31\x15\x7c\xf2\x86\x63\xfb\xe1\xe3\xe1\x0c\x0c\xc6\x04\x66\x1d\x7c\x8e\xb9\xdd\xa6\x1d\xea\xeb\x05\x44\x7f\x68\xf6\xf5\xc8\x73\x6d\x04\x6b\xa9\x80\x87\xb3\x95\xeb\xd4\x8f\x30\x10\x9c\xaa\x88\xf6\x0a\xa0\x17\xd3\x8e\x4e\xbd\x6d\x99\x8d\x2c\xa8\xc6\x17\xe2\xc9\xf9\x6f\x7a\x68\xec\xff\x00\x4d\xfc\x57\x78\xfb\xc4\x75\x24\xda\x7f\x83\x5b\xa0\xad\x37\xb8\xd8\xcb\x52\xd7\xcc\xbb\x6f\x6b\x83\x65\x80\x76\xf0\x5a\x82\xd6\x4c\xf7\xe9\xca\x07\xa5\xcc\x56\xb2\xef\x39\x34\x86\x6a\xa1\xe2\x1f\xc7\x5f\xd7\xce\xce\x0e\x3e\x9f\xa5\xd7\xc1\x50\x06\xbf\xff\x31\x6b\xb5\x52\xf1\xa5\xf7\x46\x85\xff\x0d\x00\x00\xff\xff\x72\x8a\x49\x5c\xbb\x14\x00\x00")

func aroOpenshiftIo_clustersYamlBytes() ([]byte, error) {
	return bindataRead(
		_aroOpenshiftIo_clustersYaml,
		"aro.openshift.io_clusters.yaml",
	)
}

func aroOpenshiftIo_clustersYaml() (*asset, error) {
	bytes, err := aroOpenshiftIo_clustersYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aro.openshift.io_clusters.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xc1\x6e\xdb\x30\x0c\xbd\xfb\x2b\x88\xde\xdd\xa4\xb7\x42\xb7\x62\x0d\x7a\x19\x82\x62\x59\x77\x67\x64\x26\x16\x22\x8b\x02\x49\x07\x75\xbf\x7e\x10\x92\x28\xce\x0a\x64\x3a\x19\x7c\x8f\xef\x3d\xd2\xc4\x1c\xfe\x90\x68\xe0\xe4\x00\x73\xd6\xc5\xf1\xa9\x39\x84\xd4\x39\x78\xa5\x1c\x79\x1a\x28\x59\x33\x90\x61\x87\x86\xae\x01\x88\xb8\xa5\xa8\xe5\x0b\x4a\x83\x03\x14\x6e\x39\x93\xa0\xb1\xb4\x03\xaa\x91\x34\x00\x09\x07\xba\x87\x69\x46\x4f\x0e\x38\x53\xd2\x3e\xec\xac\xc5\xaf\x51\xa8\x92\x1b\xcd\xe4\x8b\x89\x50\x8e\xc1\xa3\x3a\x78\x6a\x00\x94\x22\x79\x63\x39\xd9\x0f\x68\xbe\xff\x39\xcb\x73\x37\x91\x9a\xa0\xd1\x7e\x3a\x51\x85\x63\x0c\x69\xff\x91\x3b\x34\xba\x74\x0f\xf8\xb9\x19\x65\x4f\x27\xb3\x73\xe5\x23\xe1\x11\x43\xc4\x6d\x24\x07\xcb\x06\xc0\x68\xc8\xb1\x76\xcd\x77\x53\x5e\xbc\xc9\x73\x37\x11\xc0\x65\xca\xf2\x3c\x27\xc3\x90\x48\x6a\x73\x0b\x9e\x87\x01\x53\x77\x55\x6b\x8b\xd4\x55\x5b\xf6\x3a\xc7\xea\xf6\xae\xa5\x99\x59\x79\x61\xc0\x32\xde\xdb\x6a\xbd\xfa\xf5\xf2\x7b\xf5\x5a\x81\xef\xff\xab\x42\x99\xc5\x6e\x6c\x6a\xd2\x77\x16\x73\xf0\xbc\x7c\x5e\x56\xf4\xa2\xd4\x9b\xe5\x5a\x8c\xe1\x48\x89\x54\xdf\x85\xb7\xe4\x66\xdc\xc2\x7a\x23\x9b\x97\x00\x32\x5a\xef\x60\xd1\x13\x46\xeb\xbf\x16\x42\xd8\x4d\xb7\x84\x7f\x6d\x13\x77\xb4\xb9\x39\x8d\x4b\xb5\x15\x8e\xf4\x78\x18\xb7\x24\x89\x8c\xf4\x31\xf0\xe2\xb4\x12\x07\x0f\x0f\x67\xaa\x92\x1c\x83\xa7\x17\xef\x79\x4c\xb6\xbe\x73\xb9\xdf\xd9\xf7\x98\x59\x02\x4b\xb0\xe9\x47\x44\xd5\x93\xac\x4e\x6a\x34\xb4\x3e\x8e\x85\xd7\x7a\x09\x16\x3c\xc6\x73\x83\x71\x2c\x3a\x81\xd3\xec\x06\x0e\x34\xb9\xff\xcc\x52\x47\xbe\xe4\x70\xb0\xfa\x0c\x6a\x5a\x01\xda\xed\xc8\x9b\x83\x35\x6f\x7c\x4f\xdd\x18\xa9\xf9\x1b\x00\x00\xff\xff\x57\x5c\x5d\xa2\xfa\x03\x00\x00")

func masterDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterDeploymentYaml,
		"master/deployment.yaml",
	)
}

func masterDeploymentYaml() (*asset, error) {
	bytes, err := masterDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xb1\x4e\x03\x31\x0c\x40\x77\x7f\x85\x7f\x20\x87\xd8\x50\x36\x60\x60\x2f\x12\xbb\x9b\xb8\xd4\xf4\x62\x47\x8e\xd3\xa1\x5f\x8f\xaa\xa2\x5b\x90\x6e\xb5\xdf\xf3\x33\x75\xf9\x62\x1f\x62\x9a\xd1\x8f\x54\x16\x9a\x71\x36\x97\x1b\x85\x98\x2e\x97\x97\xb1\x88\x3d\x5d\x9f\xe1\x22\x5a\x33\xbe\xaf\x73\x04\xfb\xc1\x56\x7e\x13\xad\xa2\xdf\xd0\x38\xa8\x52\x50\x06\x44\xa5\xc6\x19\xc9\x2d\x59\x67\xa7\x30\x4f\x8d\xee\x02\xb8\xad\x7c\xe0\xd3\x1d\xa2\x2e\x1f\x6e\xb3\xef\x04\x01\xf1\x5f\x6f\x3b\x5f\x1e\xb3\x44\xb5\x89\xc2\x98\xc7\x1f\x2e\x31\x32\xa4\x3f\xe7\x93\xfd\x2a\x85\x5f\x4b\xb1\xa9\xb1\xfb\xd5\x63\x37\x3a\x15\xce\x68\x9d\x75\x9c\xe5\x14\x89\x6e\xd3\x79\x83\xe1\x37\x00\x00\xff\xff\x4f\x98\xa4\x7c\x24\x01\x00\x00")

func masterRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterRolebindingYaml,
		"master/rolebinding.yaml",
	)
}

func masterRolebindingYaml() (*asset, error) {
	bytes, err := masterRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8d\x41\xca\xc2\x40\x0c\x46\xf7\x73\x8a\x5c\x60\xa0\xff\xae\xcc\x29\x7e\x10\xdc\x87\xe9\xa7\x1d\xb4\x93\x90\xc4\x2e\x3c\xbd\xd4\x16\x5d\xb9\x0b\xef\x7b\xbc\xb0\xb6\x33\xcc\x9b\xf4\x42\xeb\x5f\xba\xb5\x3e\x15\x3a\xc1\xd6\x56\x91\x16\x04\x4f\x1c\x5c\x12\x51\xe7\x05\x85\xd8\x24\x8b\xc2\x38\xc4\xf2\xc2\x1e\xb0\x63\x73\xe5\x8a\x42\xa2\xe8\x3e\xb7\x4b\x64\x7e\x3e\x0c\x1f\x39\xb9\xa2\x6e\x1d\xc7\x1d\x35\xc4\xb6\x9b\x88\x55\x7f\x45\x55\x2c\x7c\xb7\xf2\xf1\x7d\x8e\xd0\x37\xd8\xd7\x42\xe3\x30\x0e\x07\x08\xb6\x2b\xe2\xff\x8b\x5f\x01\x00\x00\xff\xff\x10\x70\xf6\x36\xda\x00\x00\x00")

func masterServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterServiceYaml,
		"master/service.yaml",
	)
}

func masterServiceYaml() (*asset, error) {
	bytes, err := masterServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\x31\x8e\x02\x31\x0c\x05\xd0\x3e\xa7\xf0\x05\x52\x6c\xeb\x6e\xcf\x80\x44\xff\x95\xf9\x08\x0b\xc5\x8e\x1c\xcf\x14\x9c\x9e\x06\x51\xbf\x87\x65\x77\xe6\xb6\x70\x95\xeb\xaf\xbd\xcc\x0f\x95\x1b\xf3\xb2\xc1\xff\x31\xe2\xf4\x6a\x93\x85\x03\x05\x6d\x22\x8e\x49\x15\x64\xf4\x58\x4c\x54\x64\x9f\xd8\xc5\xfc\xda\x5e\x18\x54\x89\x45\xdf\x4f\x7b\x54\xc7\xfb\x4c\xfe\x72\xfb\x04\x00\x00\xff\xff\xe4\xf5\x04\x25\x70\x00\x00\x00")

func masterServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterServiceaccountYaml,
		"master/serviceaccount.yaml",
	)
}

func masterServiceaccountYaml() (*asset, error) {
	bytes, err := masterServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _namespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xb1\x0d\x02\x31\x0c\x05\xd0\x3e\x53\x58\xd7\x07\x44\x9b\x21\x28\xe9\xbf\x2e\x1f\x61\x41\xec\x28\x36\x14\x4c\x8f\xa8\xae\x7f\x98\x7a\xe3\x0a\x75\x6b\xf2\xb9\x94\xa7\x5a\x6f\x72\xc5\x60\x4c\xec\x2c\x83\x89\x8e\x44\x2b\x22\x86\xc1\x26\x3e\x69\xf1\xd0\x7b\x56\x7c\xdf\x8b\xd5\x27\x17\xd2\x57\x11\x81\x99\x27\x52\xdd\xe2\xef\xe5\xb0\x27\xf5\xb3\x79\x67\x0d\xbe\xb8\xa7\xaf\x26\xdb\x56\x7e\x01\x00\x00\xff\xff\xc1\xaf\xa6\x4c\x7c\x00\x00\x00")

func namespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_namespaceYaml,
		"namespace.yaml",
	)
}

func namespaceYaml() (*asset, error) {
	bytes, err := namespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\xcb\x6e\xdb\x40\x0c\xbc\xeb\x2b\x88\xdc\x15\x27\xb7\x60\x6f\x41\x63\xe4\x52\x04\x45\xd3\xf4\x4e\xaf\xa6\xd6\xc2\xfb\x02\x49\xbb\x55\xbe\xbe\x10\x64\xcb\x32\x02\x88\x27\x61\x38\x9c\x19\x2e\xc5\x35\xfc\x86\x68\x28\xd9\x11\xd7\xaa\x9b\xd3\x63\x73\x08\xb9\x73\xf4\x82\x1a\xcb\x90\x90\xad\x49\x30\xee\xd8\xd8\x35\x44\x91\x77\x88\x3a\x7e\xd1\x38\xe0\x88\xa5\xb4\xa5\x42\xd8\x8a\xb4\x7f\x8b\x1c\x20\x0d\x51\xe6\x84\xb5\x9e\x56\xf6\x70\x54\x2a\xb2\xf6\xe1\x8f\xb5\xfc\x79\x14\xcc\xe4\x46\x2b\xfc\x68\x22\xa8\x31\x78\x56\x47\x8f\x0d\x91\x22\xc2\x5b\x91\xc9\x3e\xb1\xf9\xfe\xfb\x22\xcf\x6a\x22\x35\x61\xc3\x7e\x98\xa8\x52\x62\x0c\x79\xff\x51\x3b\x36\x5c\xa6\x13\xff\x7b\x3f\xca\x1e\x93\xd9\x19\xf9\xc8\x7c\xe2\x10\x79\x17\xe1\xe8\xa1\x21\x32\xa4\x1a\xe7\xa9\xe5\xdb\x8c\x15\x6f\xf2\xac\x26\x22\xba\x6c\x39\x96\x2f\xd9\x38\x64\xc8\x3c\xdc\x92\x2f\x29\x71\xee\xae\x6a\xed\x28\x75\xd5\x96\xbd\x2e\x7b\xf3\xeb\x5d\xa1\x85\xd9\x58\x21\xf1\xb8\xde\xeb\xf6\x6d\xfb\xf3\xf9\xd7\xf6\x65\x6e\x7c\xbd\xd7\xdc\x8a\xe1\x84\x0c\xd5\x1f\x52\x76\xb8\xda\x11\xf5\x66\xf5\x15\xb6\x84\x88\x2a\x5b\xef\x68\xd3\x83\xa3\xf5\x9f\x1b\x01\x77\xc3\x2d\xa1\x88\x39\x7a\x7a\x78\x7a\x38\xc3\xb9\x74\x78\xbf\x39\xec\x05\x6d\xa5\x44\xdc\x1f\x8e\x3b\x48\x86\x41\xef\x43\xd9\x4c\x0b\x39\xba\xbb\x3b\x53\x15\x72\x0a\x1e\xcf\xde\x97\x63\xb6\xb7\x95\xff\xee\x2b\x7b\x8d\x59\x25\x14\x09\x36\x7c\x8b\xac\x3a\xc9\xea\xa0\x86\xd4\xfa\x78\x54\x83\xb4\x5e\x82\x05\xcf\xb1\xf9\x1f\x00\x00\xff\xff\x4f\x57\x4a\x02\x45\x03\x00\x00")

func workerDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerDeploymentYaml,
		"worker/deployment.yaml",
	)
}

func workerDeploymentYaml() (*asset, error) {
	bytes, err := workerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8e\xb1\x6e\x2c\x31\x08\x45\x7b\xbe\x82\x1f\xb0\x57\xaf\x7b\x72\x9b\x22\x7d\x14\xa5\x67\x3d\x24\x83\xc6\x63\x2c\xc0\xbb\x52\xbe\x3e\x9a\xd9\x6d\x53\xa5\xe2\x0a\x1d\x0e\x17\x52\x4a\x40\x43\x3e\xd8\x5c\xb4\x17\xb4\x2b\xd5\x4c\x33\x56\x35\xf9\xa6\x10\xed\x79\xfb\xef\x59\xf4\x72\xfb\x07\x9b\xf4\xa5\xe0\x4b\x9b\x1e\x6c\x6f\xda\x18\x76\x0e\x5a\x28\xa8\x00\x62\x35\x3e\x0f\xde\x65\x67\x0f\xda\x47\xc1\x3e\x5b\x03\xc4\x4e\x3b\x17\x24\xd3\xa4\x83\x8d\x42\x2d\xdd\xd5\x36\x36\xb0\xd9\xd8\x0b\x24\xa4\x21\xaf\xa6\x73\xf8\x61\x4a\x07\x9b\x75\x70\xf7\x55\x3e\x23\x8b\x02\xa2\xb1\xeb\xb4\xca\x4f\xa2\x3e\x5a\x38\x20\xde\xd8\xae\xcf\xed\x17\xc7\x39\x9b\xf8\x23\xdc\x29\xea\xfa\x17\xff\xc5\x83\x62\xfe\xf2\x66\x9c\xf6\x23\xcd\xb1\x50\x30\xfc\x04\x00\x00\xff\xff\x30\x78\x19\x41\x50\x01\x00\x00")

func workerRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerRoleYaml,
		"worker/role.yaml",
	)
}

func workerRoleYaml() (*asset, error) {
	bytes, err := workerRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8d\x31\x6e\xc3\x30\x0c\x45\x77\x9d\x82\x17\x90\x8b\x6e\x85\xb6\xb6\x43\x77\x17\xe8\x4e\xcb\x74\xcd\xda\x26\x05\x8a\x72\x01\x9f\x3e\x08\x12\x64\x09\xe0\xf9\xbf\xf7\x1f\x16\xfe\x21\xab\xac\x92\xc0\x06\xcc\x1d\x36\x9f\xd5\xf8\x40\x67\x95\x6e\x79\xab\x1d\xeb\xcb\xfe\x1a\x16\x96\x31\xc1\xe7\xda\xaa\x93\xf5\xba\xd2\x07\xcb\xc8\xf2\x1b\x36\x72\x1c\xd1\x31\x05\x00\xc1\x8d\x12\xa0\x69\xd4\x42\x86\xae\x16\xff\xd5\x16\xb2\x60\xba\x52\x4f\xd3\x15\xc2\xc2\x5f\xa6\xad\x9c\x04\x03\xc0\x53\xef\xf4\xbe\xb6\xe1\x8f\xb2\xd7\x14\xe2\xdd\xfc\x26\xdb\x39\xd3\x7b\xce\xda\xc4\x4f\xe5\xdb\x56\x0b\x66\x4a\xa0\x85\xa4\xce\x3c\x79\xc4\xa3\x19\x3d\xe0\x70\x09\x00\x00\xff\xff\x73\xce\x57\x9b\x2a\x01\x00\x00")

func workerRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerRolebindingYaml,
		"worker/rolebinding.yaml",
	)
}

func workerRolebindingYaml() (*asset, error) {
	bytes, err := workerRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\x31\x8a\xc3\x40\x0c\x05\xd0\x7e\x4e\xa1\x0b\x4c\xb1\xad\xba\x3d\x43\x20\xfd\x67\xfc\x43\x84\xb1\x34\x68\x64\x07\x72\xfa\x34\x21\xf5\x7b\x98\x76\x67\x2e\x0b\x57\xb9\xfe\xda\x6e\xbe\xa9\xdc\x98\x97\x0d\xfe\x8f\x11\xa7\x57\x3b\x58\xd8\x50\xd0\x26\xe2\x38\xa8\x82\x8c\x1e\x93\x89\x8a\xec\xaf\xc8\x9d\xf9\xb5\x35\x31\xa8\x12\x93\xbe\x9e\xf6\xa8\x8e\xf7\x99\xfc\xe5\xf6\x09\x00\x00\xff\xff\xe3\x3c\x43\x66\x70\x00\x00\x00")

func workerServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerServiceaccountYaml,
		"worker/serviceaccount.yaml",
	)
}

func workerServiceaccountYaml() (*asset, error) {
	bytes, err := workerServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"aro.openshift.io_clusters.yaml": aroOpenshiftIo_clustersYaml,
	"master/deployment.yaml":         masterDeploymentYaml,
	"master/rolebinding.yaml":        masterRolebindingYaml,
	"master/service.yaml":            masterServiceYaml,
	"master/serviceaccount.yaml":     masterServiceaccountYaml,
	"namespace.yaml":                 namespaceYaml,
	"worker/deployment.yaml":         workerDeploymentYaml,
	"worker/role.yaml":               workerRoleYaml,
	"worker/rolebinding.yaml":        workerRolebindingYaml,
	"worker/serviceaccount.yaml":     workerServiceaccountYaml,
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
	"aro.openshift.io_clusters.yaml": {aroOpenshiftIo_clustersYaml, map[string]*bintree{}},
	"master": {nil, map[string]*bintree{
		"deployment.yaml":     {masterDeploymentYaml, map[string]*bintree{}},
		"rolebinding.yaml":    {masterRolebindingYaml, map[string]*bintree{}},
		"service.yaml":        {masterServiceYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {masterServiceaccountYaml, map[string]*bintree{}},
	}},
	"namespace.yaml": {namespaceYaml, map[string]*bintree{}},
	"worker": {nil, map[string]*bintree{
		"deployment.yaml":     {workerDeploymentYaml, map[string]*bintree{}},
		"role.yaml":           {workerRoleYaml, map[string]*bintree{}},
		"rolebinding.yaml":    {workerRolebindingYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {workerServiceaccountYaml, map[string]*bintree{}},
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
