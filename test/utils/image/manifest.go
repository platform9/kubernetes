/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package image

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// RegistryList holds public and private image registries
type RegistryList struct {
	GcAuthenticatedRegistry string `yaml:"gcAuthenticatedRegistry"`
	DockerLibraryRegistry   string `yaml:"dockerLibraryRegistry"`
	DockerGluster           string `yaml:"dockerGluster"`
	E2eRegistry             string `yaml:"e2eRegistry"`
	E2eVolumeRegistry       string `yaml:"e2eVolumeRegistry"`
	PromoterE2eRegistry     string `yaml:"promoterE2eRegistry"`
	BuildImageRegistry      string `yaml:"buildImageRegistry"`
	InvalidRegistry         string `yaml:"invalidRegistry"`
	GcRegistry              string `yaml:"gcRegistry"`
	SigStorageRegistry      string `yaml:"sigStorageRegistry"`
	GcrReleaseRegistry      string `yaml:"gcrReleaseRegistry"`
	PrivateRegistry         string `yaml:"privateRegistry"`
	SampleRegistry          string `yaml:"sampleRegistry"`
}

// Config holds an images registry, name, and version
type Config struct {
	registry string
	name     string
	version  string
}

// SetRegistry sets an image registry in a Config struct
func (i *Config) SetRegistry(registry string) {
	i.registry = registry
}

// SetName sets an image name in a Config struct
func (i *Config) SetName(name string) {
	i.name = name
}

// SetVersion sets an image version in a Config struct
func (i *Config) SetVersion(version string) {
	i.version = version
}

func initReg() RegistryList {
	registry := RegistryList{
		GcAuthenticatedRegistry: "gcr.io/authenticated-image-pulling",
		DockerLibraryRegistry:   "docker.io/library",
		DockerGluster:           "docker.io/gluster",
		E2eRegistry:             "gcr.io/kubernetes-e2e-test-images",
		E2eVolumeRegistry:       "gcr.io/kubernetes-e2e-test-images/volume",
		PromoterE2eRegistry:     "k8s.gcr.io/e2e-test-images",
		BuildImageRegistry:      "k8s.gcr.io/build-image",
		InvalidRegistry:         "invalid.com/invalid",
		GcRegistry:              "k8s.gcr.io",
		SigStorageRegistry:      "k8s.gcr.io/sig-storage",
		GcrReleaseRegistry:      "gcr.io/gke-release",
		PrivateRegistry:         "gcr.io/k8s-authenticated-test",
		SampleRegistry:          "gcr.io/google-samples",
	}
	repoList := os.Getenv("KUBE_TEST_REPO_LIST")
	if repoList == "" {
		return registry
	}

	fileContent, err := ioutil.ReadFile(repoList)
	if err != nil {
		panic(fmt.Errorf("Error reading '%v' file contents: %v", repoList, err))
	}

	err = yaml.Unmarshal(fileContent, &registry)
	if err != nil {
		panic(fmt.Errorf("Error unmarshalling '%v' YAML file: %v", repoList, err))
	}
	return registry
}

var (
	registry                = initReg()
	dockerLibraryRegistry   = registry.DockerLibraryRegistry
	dockerGluster           = registry.DockerGluster
	e2eRegistry             = registry.E2eRegistry
	e2eVolumeRegistry       = registry.E2eVolumeRegistry
	promoterE2eRegistry     = registry.PromoterE2eRegistry
	buildImageRegistry      = registry.BuildImageRegistry
	gcAuthenticatedRegistry = registry.GcAuthenticatedRegistry
	gcRegistry              = registry.GcRegistry
	sigStorageRegistry      = registry.SigStorageRegistry
	gcrReleaseRegistry      = registry.GcrReleaseRegistry
	invalidRegistry         = registry.InvalidRegistry
	// PrivateRegistry is an image repository that requires authentication
	PrivateRegistry = registry.PrivateRegistry
	sampleRegistry  = registry.SampleRegistry

	// Preconfigured image configs
	imageConfigs = initImageConfigs()
)

const (
	// None is to be used for unset/default images
	None = iota
	// Agnhost image
	Agnhost
	// AgnhostPrivate image
	AgnhostPrivate
	// APIServer image
	APIServer
	// AppArmorLoader image
	AppArmorLoader
	// AuthenticatedAlpine image
	AuthenticatedAlpine
	// AuthenticatedWindowsNanoServer image
	AuthenticatedWindowsNanoServer
	// BusyBox image
	BusyBox
	// CheckMetadataConcealment image
	CheckMetadataConcealment
	// CudaVectorAdd image
	CudaVectorAdd
	// CudaVectorAdd2 image
	CudaVectorAdd2
	// DebianIptables Image
	// DebianIptables
	// EchoServer image
	EchoServer
	// DistrolessIptables Image
	DistrolessIptables
	// Etcd image
	Etcd
	// GlusterDynamicProvisioner image
	GlusterDynamicProvisioner
	// Httpd image
	Httpd
	// HttpdNew image
	HttpdNew
	// InvalidRegistryImage image
	InvalidRegistryImage
	// IpcUtils image
	IpcUtils
	// JessieDnsutils image
	JessieDnsutils
	// Kitten image
	Kitten
	// Nautilus image
	Nautilus
	// NFSProvisioner image
	NFSProvisioner
	// Nginx image
	Nginx
	// NginxNew image
	NginxNew
	// Nonewprivs image
	Nonewprivs
	// NonRoot runs with a default user of 1234
	NonRoot
	// Pause - when these values are updated, also update cmd/kubelet/app/options/container_runtime.go
	// Pause image
	Pause
	// Perl image
	Perl
	// PrometheusDummyExporter image
	PrometheusDummyExporter
	// PrometheusToSd image
	PrometheusToSd
	// Redis image
	Redis
	// RegressionIssue74839 image
	RegressionIssue74839
	// ResourceConsumer image
	ResourceConsumer
	// SdDummyExporter image
	SdDummyExporter
	// VolumeNFSServer image
	VolumeNFSServer
	// VolumeISCSIServer image
	VolumeISCSIServer
	// VolumeGlusterServer image
	VolumeGlusterServer
	// VolumeRBDServer image
	VolumeRBDServer
)

func initImageConfigs() map[int]Config {
	configs := map[int]Config{}
	configs[Agnhost] = Config{promoterE2eRegistry, "agnhost", "2.21"}
	configs[AgnhostPrivate] = Config{PrivateRegistry, "agnhost", "2.6"}
	configs[AuthenticatedAlpine] = Config{gcAuthenticatedRegistry, "alpine", "3.7"}
	configs[AuthenticatedWindowsNanoServer] = Config{gcAuthenticatedRegistry, "windows-nanoserver", "v1"}
	configs[APIServer] = Config{e2eRegistry, "sample-apiserver", "1.17"}
	configs[AppArmorLoader] = Config{e2eRegistry, "apparmor-loader", "1.0"}
	configs[BusyBox] = Config{dockerLibraryRegistry, "busybox", "1.29"}
	configs[CheckMetadataConcealment] = Config{promoterE2eRegistry, "metadata-concealment", "1.6"}
	configs[CudaVectorAdd] = Config{e2eRegistry, "cuda-vector-add", "1.0"}
	configs[CudaVectorAdd2] = Config{e2eRegistry, "cuda-vector-add", "2.0"}
        // We are switching from DebianIptables to DistrolessIptables for kube-proxy
	// configs[DebianIptables] = Config{buildImageRegistry, "debian-iptables", "buster-v1.6.7"}
	configs[DistrolessIptables] = Config{buildImageRegistry, "distroless-iptables", "v0.2.1"}
	configs[EchoServer] = Config{e2eRegistry, "echoserver", "2.2"}
	// configs[Etcd] = Config{gcRegistry, "etcd", "3.4.13-0"}
	configs[Etcd] = Config{gcRegistry, "etcd", "3.4.25"}
	configs[GlusterDynamicProvisioner] = Config{dockerGluster, "glusterdynamic-provisioner", "v1.0"}
	configs[Httpd] = Config{dockerLibraryRegistry, "httpd", "2.4.38-alpine"}
	configs[HttpdNew] = Config{dockerLibraryRegistry, "httpd", "2.4.39-alpine"}
	configs[InvalidRegistryImage] = Config{invalidRegistry, "alpine", "3.1"}
	configs[IpcUtils] = Config{e2eRegistry, "ipc-utils", "1.0"}
	configs[JessieDnsutils] = Config{e2eRegistry, "jessie-dnsutils", "1.0"}
	configs[Kitten] = Config{e2eRegistry, "kitten", "1.0"}
	configs[Nautilus] = Config{e2eRegistry, "nautilus", "1.0"}
	configs[NFSProvisioner] = Config{sigStorageRegistry, "nfs-provisioner", "v2.2.2"}
	configs[Nginx] = Config{dockerLibraryRegistry, "nginx", "1.14-alpine"}
	configs[NginxNew] = Config{dockerLibraryRegistry, "nginx", "1.15-alpine"}
	configs[Nonewprivs] = Config{e2eRegistry, "nonewprivs", "1.0"}
	configs[NonRoot] = Config{e2eRegistry, "nonroot", "1.0"}
/*=
======
func initImageConfigs(list RegistryList) (map[ImageID]Config, map[ImageID]Config) {
	configs := map[ImageID]Config{}
	configs[Agnhost] = Config{list.PromoterE2eRegistry, "agnhost", "2.39"}
	configs[AgnhostPrivate] = Config{list.PrivateRegistry, "agnhost", "2.6"}
	configs[AuthenticatedAlpine] = Config{list.GcAuthenticatedRegistry, "alpine", "3.7"}
	configs[AuthenticatedWindowsNanoServer] = Config{list.GcAuthenticatedRegistry, "windows-nanoserver", "v1"}
	configs[APIServer] = Config{list.PromoterE2eRegistry, "sample-apiserver", "1.17.5"}
	configs[AppArmorLoader] = Config{list.PromoterE2eRegistry, "apparmor-loader", "1.4"}
	configs[BusyBox] = Config{list.PromoterE2eRegistry, "busybox", "1.29-2"}
	configs[CudaVectorAdd] = Config{list.PromoterE2eRegistry, "cuda-vector-add", "1.0"}
	configs[CudaVectorAdd2] = Config{list.PromoterE2eRegistry, "cuda-vector-add", "2.2"}
	configs[DistrolessIptables] = Config{list.BuildImageRegistry, "distroless-iptables", "v0.2.1"}
	configs[Etcd] = Config{list.GcEtcdRegistry, "etcd", "3.5.7-0"}
	configs[GlusterDynamicProvisioner] = Config{list.PromoterE2eRegistry, "glusterdynamic-provisioner", "v1.3"}
	configs[Httpd] = Config{list.PromoterE2eRegistry, "httpd", "2.4.38-2"}
	configs[HttpdNew] = Config{list.PromoterE2eRegistry, "httpd", "2.4.39-2"}
	configs[InvalidRegistryImage] = Config{list.InvalidRegistry, "alpine", "3.1"}
	configs[IpcUtils] = Config{list.PromoterE2eRegistry, "ipc-utils", "1.3"}
	configs[JessieDnsutils] = Config{list.PromoterE2eRegistry, "jessie-dnsutils", "1.5"}
	configs[Kitten] = Config{list.PromoterE2eRegistry, "kitten", "1.5"}
	configs[Nautilus] = Config{list.PromoterE2eRegistry, "nautilus", "1.5"}
	configs[NFSProvisioner] = Config{list.SigStorageRegistry, "nfs-provisioner", "v3.0.1"}
	configs[Nginx] = Config{list.PromoterE2eRegistry, "nginx", "1.14-2"}
	configs[NginxNew] = Config{list.PromoterE2eRegistry, "nginx", "1.15-2"}
	configs[NodePerfNpbEp] = Config{list.PromoterE2eRegistry, "node-perf/npb-ep", "1.2"}
	configs[NodePerfNpbIs] = Config{list.PromoterE2eRegistry, "node-perf/npb-is", "1.2"}
	configs[NodePerfTfWideDeep] = Config{list.PromoterE2eRegistry, "node-perf/tf-wide-deep", "1.2"}
	configs[Nonewprivs] = Config{list.PromoterE2eRegistry, "nonewprivs", "1.3"}
	configs[NonRoot] = Config{list.PromoterE2eRegistry, "nonroot", "1.2"}
>>>>>>> 1ee1390... use new distroless kube-proxy image
*/
	// Pause - when these values are updated, also update cmd/kubelet/app/options/container_runtime.go
	configs[Pause] = Config{gcRegistry, "pause", "3.2"}
	configs[Perl] = Config{dockerLibraryRegistry, "perl", "5.26"}
	configs[PrometheusDummyExporter] = Config{gcRegistry, "prometheus-dummy-exporter", "v0.1.0"}
	configs[PrometheusToSd] = Config{gcRegistry, "prometheus-to-sd", "v0.5.0"}
	configs[Redis] = Config{dockerLibraryRegistry, "redis", "5.0.5-alpine"}
	configs[RegressionIssue74839] = Config{e2eRegistry, "regression-issue-74839-amd64", "1.0"}
	configs[ResourceConsumer] = Config{e2eRegistry, "resource-consumer", "1.5"}
	configs[SdDummyExporter] = Config{gcRegistry, "sd-dummy-exporter", "v0.2.0"}
	configs[VolumeNFSServer] = Config{e2eVolumeRegistry, "nfs", "1.0"}
	configs[VolumeISCSIServer] = Config{e2eVolumeRegistry, "iscsi", "2.0"}
	configs[VolumeGlusterServer] = Config{e2eVolumeRegistry, "gluster", "1.0"}
	configs[VolumeRBDServer] = Config{e2eVolumeRegistry, "rbd", "1.0.1"}
	return configs
}

// GetImageConfigs returns the map of imageConfigs
func GetImageConfigs() map[int]Config {
	return imageConfigs
}

// GetConfig returns the Config object for an image
func GetConfig(image int) Config {
	return imageConfigs[image]
}

// GetE2EImage returns the fully qualified URI to an image (including version)
func GetE2EImage(image int) string {
	return fmt.Sprintf("%s/%s:%s", imageConfigs[image].registry, imageConfigs[image].name, imageConfigs[image].version)
}

// GetE2EImage returns the fully qualified URI to an image (including version)
func (i *Config) GetE2EImage() string {
	return fmt.Sprintf("%s/%s:%s", i.registry, i.name, i.version)
}

// GetPauseImageName returns the pause image name with proper version
func GetPauseImageName() string {
	return GetE2EImage(Pause)
}

// ReplaceRegistryInImageURL replaces the registry in the image URL with a custom one
func ReplaceRegistryInImageURL(imageURL string) (string, error) {
	parts := strings.Split(imageURL, "/")
	countParts := len(parts)
	registryAndUser := strings.Join(parts[:countParts-1], "/")

	switch registryAndUser {
	case "gcr.io/kubernetes-e2e-test-images":
		registryAndUser = e2eRegistry
	case "gcr.io/kubernetes-e2e-test-images/volume":
		registryAndUser = e2eVolumeRegistry
	case "k8s.gcr.io":
		registryAndUser = gcRegistry
	case "k8s.gcr.io/sig-storage":
		registryAndUser = sigStorageRegistry
	case "gcr.io/k8s-authenticated-test":
		registryAndUser = PrivateRegistry
	case "gcr.io/google-samples":
		registryAndUser = sampleRegistry
	case "gcr.io/gke-release":
		registryAndUser = gcrReleaseRegistry
	case "docker.io/library":
		registryAndUser = dockerLibraryRegistry
	default:
		if countParts == 1 {
			// We assume we found an image from docker hub library
			// e.g. openjdk -> docker.io/library/openjdk
			registryAndUser = dockerLibraryRegistry
			break
		}

		return "", fmt.Errorf("Registry: %s is missing in test/utils/image/manifest.go, please add the registry, otherwise the test will fail on air-gapped clusters", registryAndUser)
	}

	return fmt.Sprintf("%s/%s", registryAndUser, parts[countParts-1]), nil
}
