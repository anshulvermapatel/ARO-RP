// Code generated for package deploy by go-bindata DO NOT EDIT. (@generated)
// sources:
// deploy/staticresources/aro.openshift.io_clusters.yaml
// deploy/staticresources/master/deployment.yaml
// deploy/staticresources/master/rolebinding.yaml
// deploy/staticresources/master/service.yaml
// deploy/staticresources/master/serviceaccount.yaml
// deploy/staticresources/namespace.yaml
// deploy/staticresources/preview.aro.openshift.io_previewfeatures.yaml
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

var _aroOpenshiftIo_clustersYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x18\xcb\x6e\xdb\x46\xf0\xae\xaf\x18\xa4\x07\x1f\x5a\x51\x49\x0b\x14\xad\x6e\xae\x9c\xa4\x42\xd2\x58\xb0\x8d\x5c\x82\x1c\x46\xe4\x90\x9a\x98\xdc\x65\x77\x86\x4a\x9c\xa2\xff\x5e\xec\x92\x94\x28\x99\x7a\xd0\x41\x2e\xe5\x45\xe2\xec\xbc\xdf\xcb\xd1\x78\x3c\x1e\x61\xc9\xef\xc9\x09\x5b\x33\x05\x2c\x99\xbe\x28\x19\xff\x26\xd1\xfd\x6f\x12\xb1\x9d\xac\x5f\x8c\xee\xd9\x24\x53\x98\x55\xa2\xb6\xb8\x21\xb1\x95\x8b\xe9\x8a\x52\x36\xac\x6c\xcd\xa8\x20\xc5\x04\x15\xa7\x23\x00\x34\xc6\x2a\x7a\xb0\xf8\x57\x80\xd8\x1a\x75\x36\xcf\xc9\x8d\x33\x32\xd1\x7d\xb5\xa4\x65\xc5\x79\x42\x2e\x30\x6f\x45\xaf\x9f\x47\xbf\x46\x3f\x8f\x00\x62\x47\x81\xfc\x8e\x0b\x12\xc5\xa2\x9c\x82\xa9\xf2\x7c\x04\x60\xb0\xa0\x29\xc4\x79\x25\x4a\x4e\x22\x74\x36\xb2\x25\x19\x59\x71\xaa\x11\xdb\x91\x94\x14\x7b\x99\x99\xb3\x55\x39\x85\x47\xe7\x35\x87\x46\xad\xc6\xa4\x9a\x59\x80\xe4\x2c\xfa\xa6\x0b\x7d\xcb\xa2\xe1\xa4\xcc\x2b\x87\xf9\x56\x74\x00\x0a\x9b\xac\xca\xd1\x6d\xc0\x23\x00\x89\x6d\x49\x5d\xae\x8d\x79\x41\xe6\xb8\x31\x60\xfd\x02\xf3\x72\x85\x2f\x6a\x2e\xf1\x8a\x0a\xac\x55\x02\xf0\xea\x5e\x2e\xe6\xef\x7f\xb9\xdd\x01\x03\x24\x24\xb1\xe3\x52\x83\xab\x1a\xf6\xc0\x02\xba\x22\xa8\x71\x21\xb5\x2e\xbc\xb6\x4a\xc2\xe5\x62\xbe\xa1\x2f\x9d\x2d\xc9\x29\xb7\xd6\xd7\x4f\x27\xf4\x1d\xe8\x9e\xb4\x0b\xaf\x50\x8d\x05\x89\x8f\x39\xd5\x62\x1b\xd3\x28\x69\x6c\x00\x9b\x82\xae\x58\xc0\x51\xe9\x48\xc8\xd4\x59\xb0\xc3\x18\x3c\x12\x1a\xb0\xcb\x4f\x14\x6b\x04\xb7\xe4\x3c\x1b\x90\x95\xad\xf2\xc4\xa7\xca\x9a\x9c\x82\xa3\xd8\x66\x86\xbf\x6e\x78\x0b\xa8\x0d\x42\x73\x54\x6a\x82\xb2\x7d\xd8\x28\x39\x83\x39\xac\x31\xaf\xe8\x27\x40\x93\x40\x81\x0f\xe0\xc8\x4b\x81\xca\x74\xf8\x05\x14\x89\xe0\x2f\xeb\x08\xd8\xa4\x76\x0a\x2b\xd5\x52\xa6\x93\x49\xc6\xda\xa6\x7c\x6c\x8b\xa2\x32\xac\x0f\x93\x90\xbd\xbc\xac\xd4\x3a\x99\x24\xb4\xa6\x7c\x22\x9c\x8d\xd1\xc5\x2b\x56\x8a\xb5\x72\x34\xc1\x92\xc7\x41\x75\x13\xd2\x3e\x2a\x92\x1f\x5c\x53\x24\x72\xb1\xa3\xab\x3e\xf8\xf4\x10\x75\x6c\xb2\xce\x41\xc8\xc5\x23\x11\xf0\x59\xe9\xa3\x8d\x0d\x69\x6d\xc5\xd6\xd1\x1e\xe4\xbd\x73\xf3\xf2\xf6\x0e\x5a\xd1\x21\x18\xfb\xde\x0f\x7e\xdf\x12\xca\x36\x04\xde\x61\x6c\x52\x72\x75\x10\x53\x67\x8b\xc0\x93\x4c\x52\x5a\x36\xda\xe4\x16\x93\xd9\x77\xbf\x54\xcb\x82\xd5\xc7\xfd\xef\x8a\x44\x7d\xac\x22\x98\x85\x3e\x00\x4b\x82\xaa\x4c\x50\x29\x89\x60\x6e\x60\x86\x05\xe5\x33\x14\xfa\xee\x01\xf0\x9e\x96\xb1\x77\xec\x79\x21\xe8\xb6\xb0\x7d\xe4\xda\x6b\x9d\x83\xb6\xd1\x1c\x88\x57\x53\x9f\xb7\x25\xc5\x3b\x15\x93\x90\xb0\xf3\x39\xad\xa8\xe4\x2b\xa1\xdb\x7d\xda\xa7\xbf\x52\xfd\x83\xb1\xbb\xb2\x05\xb2\xd9\x3f\x38\x68\x14\xd4\x35\x3e\x37\x3a\x5f\x0c\x23\xea\x78\xb7\xb7\x43\x6c\xe9\x7d\xf1\x65\x7b\x36\x00\xe0\xd7\x97\x66\xcd\xce\x9a\x82\x8c\x0e\x12\xbd\x44\x63\xc8\x3d\x26\xd9\xf1\xf0\x1f\x01\x69\xe3\x5c\x4e\x01\x5b\x58\xd3\x4a\x96\xe4\xff\x7d\x36\x6d\xe3\x88\xc3\xec\x7a\xa4\xe7\x31\x7f\x43\x33\xbc\x7a\x2d\x38\x61\xc5\xc1\xd4\x09\x4c\xeb\xb0\xb7\x83\xf4\xb5\x1f\x57\xf3\x64\x90\x97\x92\xe1\x89\x90\xa1\xd2\x67\x7c\xa8\x53\xa8\xc7\x58\x56\x2a\x7a\x7d\x70\x86\x99\xe8\x1c\x3e\xf4\xcb\x5b\x38\x5e\xa3\xd2\xcb\xa6\x8d\x0c\x4c\xc4\x8c\x0c\xad\xf1\xad\xcd\x32\x36\xd9\x63\xca\x93\xc1\x4b\x39\x3b\x98\xbf\x81\x01\xaa\x9f\x1d\x53\xb8\xf8\xf0\x7c\xfc\xfb\xc7\x1f\xa3\xfa\xe7\x62\x78\xbc\x01\x0a\x6b\x58\xad\x3f\x7c\x3d\xbb\xbd\x8c\x63\x5b\x1d\x4a\x1c\x32\x55\xd1\x7f\x32\x86\xcb\x9b\xeb\x76\xfd\xb0\x99\xcc\xdf\xdd\x9d\x85\xb7\xb8\xb9\xbe\x3a\x0b\xf1\x9b\x0d\x3b\x5a\xd7\xa7\x8c\xbb\x62\xcc\x8c\x15\xe5\x58\x16\xce\x26\x07\xb0\xee\x1e\x8f\xf8\xf6\x68\x86\xaf\x90\x5d\x8a\x5f\xbe\xd9\x8e\x77\x7e\x15\x2c\x31\xa6\xff\x41\x88\x8e\xf4\x1a\x36\xa9\xc3\x81\xcd\x85\x4d\xe6\x48\x64\x60\xa9\xd6\x5b\x18\xe9\x6c\x45\xf1\x7d\x5f\x07\x3f\x5e\xac\x95\xcb\x7b\xe1\x47\x1a\xd3\x09\x85\xba\x08\x7d\x0d\xea\xa8\xdf\x72\x1b\x87\xf5\x75\x90\x0b\xbc\x7d\xa8\xd6\xa5\x39\x66\x3d\x1a\x63\x92\x84\xfb\x12\xe6\x8b\xa3\xae\x38\x6a\xd3\xce\x1c\xbc\x6e\x04\xbe\xf2\x02\x37\xe3\x30\x25\xf4\x83\x3b\x74\x60\xd9\xdc\x0b\x2e\x6f\xae\x37\xf8\x43\x3c\xd1\xee\x93\x7d\x49\xb4\xa3\x4c\x3b\xcf\xe6\x57\xed\xcd\xe4\xf2\xab\x57\x63\xcb\xa0\xbe\x22\x50\xe7\xc2\x74\xb6\xe1\xa2\xd6\x61\x46\xb7\x55\x9a\xf2\x97\x41\x41\x59\x1b\xd2\x41\x05\x70\x68\xe9\x53\xd4\x4a\xce\x58\xfb\x02\xde\xce\xe2\x67\x97\xe2\xb7\xec\x27\x6f\x7e\xb1\x35\x75\xe6\x0c\x19\xdb\xbd\x89\x32\x6b\x39\xf9\x10\x7d\xaa\xa4\x5e\xeb\x45\xd1\x24\xe8\x92\xad\x20\x48\x99\xf2\x44\xa2\x1e\xbe\xc7\xcb\x18\x20\x47\xd1\x3b\x87\x46\xb8\xbd\xc5\x1f\xaa\xdd\xd4\xba\x02\x75\x0a\xfe\x7e\x30\x56\x2e\xe8\xa9\x35\x5e\x90\x08\x66\x07\xe5\x9c\xa4\x77\x84\x72\x68\x43\x38\x83\xbc\x2f\x33\x06\x90\x07\x84\xa7\x11\x1f\x29\xdb\x63\x9d\xaf\xed\x53\x27\x56\xfb\x5e\xb1\x8e\x92\x3f\x51\xdf\xd0\x83\x2c\xea\x3b\xe4\xf7\xde\x24\x7b\x6d\x7c\x04\xac\x0b\x6c\x0a\xea\xaa\x3a\x8d\x9a\x86\xd1\x85\x54\xcb\xcd\xb5\xbc\xd5\xae\x09\x1d\xfc\xf3\xef\x68\x1b\x45\x8c\x63\x2a\x95\x92\x77\xfb\x5f\x8b\x9e\x3d\xdb\xf9\x1c\x14\x5e\x3b\xb5\x09\x1f\x3e\x8e\x6a\xc1\x94\xbc\x6f\x3f\xfc\x78\xe0\x7f\x01\x00\x00\xff\xff\x00\xf5\x06\x7a\x68\x13\x00\x00")

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

var _previewAroOpenshiftIo_previewfeaturesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4d\x73\xdb\x38\x12\xbd\xeb\x57\x74\x65\x0f\xbe\x84\xb4\x14\x67\xb5\x1b\xdd\xbc\xce\x47\xb9\x36\x89\x53\x96\xcb\x53\x53\x53\x73\x68\x11\x4d\x0a\x31\x08\x30\x40\x43\x1a\xcd\xd4\xfc\xf7\x29\x80\xa4\x24\x52\xa4\xed\x49\x65\x70\x22\xf1\xf1\xba\xf1\xfa\x35\x1a\x98\x24\x49\x32\xc1\x4a\xde\x93\x75\xd2\xe8\x05\x60\x25\xe9\x37\x26\x1d\xfe\x5c\xfa\xf0\x5f\x97\x4a\x73\xbe\x99\x4d\x1e\xa4\x16\x0b\xb8\xf2\x8e\x4d\x79\x4b\xce\x78\x9b\xd1\x5b\xca\xa5\x96\x2c\x8d\x9e\x94\xc4\x28\x90\x71\x31\x01\x40\xad\x0d\x63\xe8\x76\xe1\x17\x20\x33\x9a\xad\x51\x8a\x6c\x52\x90\x4e\x1f\xfc\x8a\x56\x5e\x2a\x41\x36\x82\xb7\xa6\x37\xd3\x74\x9e\x5e\x24\xd3\xf4\xd5\xf4\xd5\x6c\xfa\x66\x36\x9f\x5d\x4c\xff\xf3\x7a\x9e\xbc\x79\xfd\x7a\x3a\x9b\xff\x7b\x86\xf3\xec\x62\x02\x90\x59\x8a\xe0\x77\xb2\x24\xc7\x58\x56\x0b\xd0\x5e\xa9\x09\x80\xc6\x92\x16\x50\x59\xda\x48\xda\xe6\x84\xec\x2d\xb9\xb4\xf9\x4f\xd1\x9a\xd4\x54\xa4\xdd\x5a\xe6\x9c\x4a\x33\x71\x15\x65\xc1\xc1\xc2\x1a\x5f\xed\xd7\x9d\xce\xab\x81\x9b\xbd\xd4\x3c\x7c\xa9\xe7\xbe\xaf\x6d\xc4\x01\x25\x1d\xff\x7f\x60\xf0\xa3\x74\x1c\x27\x54\xca\x5b\x54\x27\xfe\xc5\x31\x27\x75\xe1\x15\xda\xfe\xe8\x04\xc0\x65\xa6\xa2\x05\x5c\x29\xef\x98\xec\x04\xa0\xe1\x2b\xfa\x93\x34\x7b\xde\xcc\x50\x55\x6b\x9c\xd5\x60\xd9\x9a\x4a\xac\xdd\x05\x08\x5b\xb9\xfc\x72\x7d\x7f\xb1\xec\x74\x03\x08\x72\x99\x95\x15\x47\xee\xbb\x3e\x83\x74\xc0\x6b\x82\x7a\x09\xe4\xc6\xc6\xdf\xc6\x37\x68\x9c\x83\xcb\x2f\xd7\x7b\xb4\xca\x9a\x8a\x2c\xcb\x96\xa7\xba\x1d\x29\xeb\xa8\xb7\x67\xfb\x2c\xb8\x57\xcf\x02\x11\x24\x45\xb5\xf5\x66\xa3\x24\x9a\x1d\x81\xc9\x81\xd7\xd2\x81\xa5\xca\x92\x23\x5d\x8b\xac\x03\x0c\x61\x12\x6a\x30\xab\xaf\x94\x71\x0a\x4b\xb2\x01\x06\xdc\xda\x78\x25\x82\x12\x37\x64\x19\x2c\x65\xa6\xd0\xf2\xf7\x3d\xb6\x03\x36\xd1\xa8\x42\xa6\x26\x60\x87\x26\x35\x93\xd5\xa8\x60\x83\xca\xd3\x4b\x40\x2d\xa0\xc4\x1d\x58\x0a\x56\xc0\xeb\x23\xbc\x38\xc5\xa5\xf0\xc9\x04\x1e\x75\x6e\x16\xb0\x66\xae\xdc\xe2\xfc\xbc\x90\xdc\x66\x54\x66\xca\xd2\x6b\xc9\xbb\xf3\x98\x1c\x72\xe5\xd9\x58\x77\x2e\x68\x43\xea\xdc\xc9\x22\x41\x9b\xad\x25\x53\x16\x88\x3e\xc7\x4a\x26\xd1\x75\x1d\xb3\x2a\x2d\xc5\xbf\x6c\x93\x83\xee\xac\xe3\x2b\xef\x82\x58\x1c\x5b\xa9\x8b\xa3\x81\xa8\xda\x47\x22\x10\x84\x1b\x82\x8e\xcd\xd2\x7a\x17\x07\xa2\x43\x57\x60\xe7\xf6\xdd\xf2\x0e\x5a\xd3\x31\x18\x7d\xf6\x23\xef\x87\x85\xee\x10\x82\x40\x98\xd4\x39\xd9\x3a\x88\xb9\x35\x65\xc4\x24\x2d\x2a\x23\x35\xc7\x9f\x4c\x49\xd2\x7d\xfa\x9d\x5f\x95\x92\x43\xdc\xbf\x79\x72\x1c\x62\x95\xc2\x55\x3c\x66\x60\x45\xe0\x2b\x81\x4c\x22\x85\x6b\x0d\x57\x58\x92\xba\x42\x47\xff\x78\x00\x02\xd3\x2e\x09\xc4\x3e\x2f\x04\xc7\x27\x64\x7f\x72\xcd\xda\xd1\x40\x7b\x34\x8d\xc4\xab\x9b\xad\xcb\x8a\xb2\x4e\xe2\xf4\xf3\x34\xe4\xef\xe5\xed\x4d\x07\x6f\x38\x5f\x43\xd3\xae\x78\xaf\xcc\xf6\xa3\x29\x4e\x86\x7a\x6e\x7c\x5e\x7e\x68\x67\xc6\x33\x1e\xa5\x8e\x1f\xb9\x2c\xbc\x8d\xb9\x19\x4d\x7f\x5e\x7e\x80\x5c\x99\x2d\x28\x53\xb8\xf4\x04\x12\xe0\xa6\x94\x4d\xf4\x3b\x6b\x65\x0e\x3b\xe3\x41\x18\x7d\xc6\xb0\x45\xbd\x9f\xd3\x14\x93\x90\xb2\x21\xf1\x74\x26\x15\x0d\xc0\x3e\x61\x77\x9c\x81\xd0\x48\xe3\x4a\x91\x18\x1a\xea\x27\xcf\xbb\x7a\x6a\x27\x04\x2b\x5a\xe3\x46\x1a\x6f\xeb\x43\x8b\x0e\x8e\xda\x61\x48\x80\xed\x9a\x34\xb0\xf5\xd4\xdf\xe6\x56\x2a\x05\x6c\x77\xf5\x7e\x31\x5b\xc7\x09\x82\x9c\xb4\x24\xba\x9c\xbd\x7c\x0c\x3b\x47\xe5\x08\x24\x77\x00\x85\x74\xc1\xfb\x08\x79\x20\xeb\x6c\x10\xa7\x16\xeb\xca\x18\x45\xd8\x3f\x76\x43\xb3\xc4\x75\x82\x8c\xb1\x96\xa3\x57\xbc\x80\x57\xb3\xf9\x74\xfd\x34\xb1\xb7\x2d\x1c\x54\x64\xa5\x11\x51\x4c\x21\x85\x86\x34\x34\x9a\x77\x6d\x73\x6c\x2c\x16\x74\x99\x65\xc6\x6b\x6e\xaf\x30\xd7\xcf\x89\xf0\x72\x78\xe9\xdb\xb6\xaa\xac\xa8\x89\xbd\x08\x9a\xfd\xdf\xcf\x37\x49\xb3\x62\x24\x1a\x58\x03\x85\x03\xd7\xbb\x70\x70\x7d\xf2\x2e\x9e\x64\x52\xc7\x38\x38\x2c\x83\x60\x8a\xb0\x75\x93\xef\xc3\xf2\x5d\xdb\x66\x8b\x79\x2e\xb3\x4b\x8d\x6a\xc7\x32\x73\xd7\xa1\x94\x6d\x50\x3d\x11\xa2\xf9\xb4\x7c\x86\xf2\x5b\x30\x40\x86\xed\x5a\x06\x65\x9a\x20\x48\xe1\x33\xae\xdd\xc6\xd6\x6e\x0a\xf7\xb1\x30\x8e\xa9\x7f\x3e\x2d\x5f\xc2\x6c\x5a\xa6\xf0\xf6\xc8\x87\x11\x21\x92\xf6\xe5\x30\x50\x32\xea\x79\x12\xd0\x7f\x04\x83\x1f\x4d\xb1\xff\xfe\xc9\xd8\x07\x57\xe1\x33\x75\x74\x4b\xdf\x7c\xcc\xd9\x20\xe4\xbb\x1e\xec\xa0\x0a\x46\xc8\xfa\x41\xda\xd8\x0c\x5d\xcd\x0e\x9e\xb7\xd9\xfa\xf4\xc6\xfa\xb7\xb7\x06\x38\x38\x78\x7c\x08\x0f\x3b\xf9\x58\x30\x67\x23\xfd\xc3\x3e\xd5\xdb\x0d\x97\xb5\x22\xde\x94\xbb\xcd\x36\xec\x9f\xda\x4a\xda\xd3\xfe\x64\x64\xa4\x3c\x8f\xd7\x6d\x46\xf6\xee\xf9\x95\x3b\x4e\xef\x14\x0e\xb3\x72\xe1\xbe\x24\x22\x14\x05\x06\x07\xde\x1a\x6d\x1b\x2f\x61\x21\x09\xe5\xd1\x13\xec\xb8\x49\xa6\x72\x30\x0f\x3b\x9e\xde\x54\x64\x91\x8d\xbd\x6a\x91\xc2\x71\xf5\x35\x48\x34\x8a\x93\x51\x0b\xb4\xe2\x60\x08\x72\x49\x4a\x0c\xd6\xf7\xc7\x2b\x2d\x80\x42\xc7\x77\x16\xb5\x93\xed\xbb\x6e\xec\x94\xc8\x8d\x2d\x91\x17\xa1\x0a\x50\xc2\x72\x34\x41\x9e\x10\x7e\xb8\x8d\x39\x87\xc5\xa8\x9d\x27\xd7\x5b\x42\x37\x96\x39\xcf\x58\x3e\xa4\x93\xbf\xb1\x3c\x4e\xf8\xbe\xc5\xa3\x8a\x3e\x0c\xa2\xb5\xb8\xeb\xdf\xeb\x1b\x35\x0c\xbe\xe6\x1e\x35\x3b\x68\xf0\xa4\xb3\x16\xfd\x22\xde\x7f\xea\x8e\xba\x82\x1e\xf7\xf8\xd5\xfe\xd1\xd3\xda\x6f\x78\x84\x3f\xfe\x9c\x1c\x28\xc5\x2c\xa3\x8a\x49\x7c\xee\xbf\xda\x5f\xbc\xe8\x3c\xc4\xe3\xef\x51\xa2\xc0\x2f\xbf\x4e\x6a\xc3\x24\xee\xdb\x47\x76\xe8\xfc\x2b\x00\x00\xff\xff\x68\x5e\x36\x49\x25\x11\x00\x00")

func previewAroOpenshiftIo_previewfeaturesYamlBytes() ([]byte, error) {
	return bindataRead(
		_previewAroOpenshiftIo_previewfeaturesYaml,
		"preview.aro.openshift.io_previewfeatures.yaml",
	)
}

func previewAroOpenshiftIo_previewfeaturesYaml() (*asset, error) {
	bytes, err := previewAroOpenshiftIo_previewfeaturesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "preview.aro.openshift.io_previewfeatures.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"aro.openshift.io_clusters.yaml":                aroOpenshiftIo_clustersYaml,
	"master/deployment.yaml":                        masterDeploymentYaml,
	"master/rolebinding.yaml":                       masterRolebindingYaml,
	"master/service.yaml":                           masterServiceYaml,
	"master/serviceaccount.yaml":                    masterServiceaccountYaml,
	"namespace.yaml":                                namespaceYaml,
	"preview.aro.openshift.io_previewfeatures.yaml": previewAroOpenshiftIo_previewfeaturesYaml,
	"worker/deployment.yaml":                        workerDeploymentYaml,
	"worker/role.yaml":                              workerRoleYaml,
	"worker/rolebinding.yaml":                       workerRolebindingYaml,
	"worker/serviceaccount.yaml":                    workerServiceaccountYaml,
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
	"preview.aro.openshift.io_previewfeatures.yaml": {previewAroOpenshiftIo_previewfeaturesYaml, map[string]*bintree{}},
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
