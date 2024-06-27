/*
Copyright 2024 The Kubernetes Authors.

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

package version

import (
	"fmt"
	"sync/atomic"

	"k8s.io/apimachinery/pkg/util/version"
	baseversion "k8s.io/component-base/version"
)

type EffectiveVersion interface {
	BinaryVersion() *version.Version
	EmulationVersion() *version.Version
	MinCompatibilityVersion() *version.Version
	EqualTo(other EffectiveVersion) bool
	String() string
	Validate() []error
}

type MutableEffectiveVersion interface {
	EffectiveVersion
	Set(binaryVersion, emulationVersion, minCompatibilityVersion *version.Version)
	SetEmulationVersion(emulationVersion *version.Version)
	SetMinCompatibilityVersion(minCompatibilityVersion *version.Version)
}

type effectiveVersion struct {
	binaryVersion atomic.Pointer[version.Version]
	// If the emulationVersion is set by the users, it could only contain major and minor versions.
	// In tests, emulationVersion could be the same as the binary version, or set directly,
	// which can have "alpha" as pre-release to continue serving expired apis while we clean up the test.
	emulationVersion atomic.Pointer[version.Version]
	// minCompatibilityVersion could only contain major and minor versions.
	minCompatibilityVersion atomic.Pointer[version.Version]
}

func (m *effectiveVersion) BinaryVersion() *version.Version {
	return m.binaryVersion.Load()
}

func (m *effectiveVersion) EmulationVersion() *version.Version {
	ver := m.emulationVersion.Load()
	if ver != nil {
		// Emulation version can have "alpha" as pre-release to continue serving expired apis while we clean up the test.
		// The pre-release should not be accessible to the users.
		return ver.WithPreRelease(m.BinaryVersion().PreRelease())
	}
	return ver
}

func (m *effectiveVersion) MinCompatibilityVersion() *version.Version {
	return m.minCompatibilityVersion.Load()
}

func (m *effectiveVersion) EqualTo(other EffectiveVersion) bool {
	return m.BinaryVersion().EqualTo(other.BinaryVersion()) && m.EmulationVersion().EqualTo(other.EmulationVersion()) && m.MinCompatibilityVersion().EqualTo(other.MinCompatibilityVersion())
}

func (m *effectiveVersion) String() string {
	if m == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BinaryVersion: %s, EmulationVersion: %s, MinCompatibilityVersion: %s}",
		m.BinaryVersion().String(), m.EmulationVersion().String(), m.MinCompatibilityVersion().String())
}

func majorMinor(ver *version.Version) *version.Version {
	if ver == nil {
		return ver
	}
	return version.MajorMinor(ver.Major(), ver.Minor())
}

func (m *effectiveVersion) Set(binaryVersion, emulationVersion, minCompatibilityVersion *version.Version) {
	m.binaryVersion.Store(binaryVersion)
	m.emulationVersion.Store(majorMinor(emulationVersion))
	m.minCompatibilityVersion.Store(majorMinor(minCompatibilityVersion))
}

func (m *effectiveVersion) SetEmulationVersion(emulationVersion *version.Version) {
	m.emulationVersion.Store(majorMinor(emulationVersion))
}

func (m *effectiveVersion) SetMinCompatibilityVersion(minCompatibilityVersion *version.Version) {
	m.minCompatibilityVersion.Store(majorMinor(minCompatibilityVersion))
}

func (m *effectiveVersion) Validate() []error {
	var errs []error
	// Validate only checks the major and minor versions.
	binaryVersion := m.binaryVersion.Load().WithPatch(0)
	emulationVersion := m.emulationVersion.Load()
	minCompatibilityVersion := m.minCompatibilityVersion.Load()

	// emulationVersion can only be 1.{binaryMinor-1}...1.{binaryMinor}.
	maxEmuVer := binaryVersion
	minEmuVer := binaryVersion.SubtractMinor(1)
	if emulationVersion.GreaterThan(maxEmuVer) || emulationVersion.LessThan(minEmuVer) {
		errs = append(errs, fmt.Errorf("emulation version %s is not between [%s, %s]", emulationVersion.String(), minEmuVer.String(), maxEmuVer.String()))
	}
	// minCompatibilityVersion can only be 1.{binaryMinor-1} for alpha.
	maxCompVer := binaryVersion.SubtractMinor(1)
	minCompVer := binaryVersion.SubtractMinor(1)
	if minCompatibilityVersion.GreaterThan(maxCompVer) || minCompatibilityVersion.LessThan(minCompVer) {
		errs = append(errs, fmt.Errorf("minCompatibilityVersion version %s is not between [%s, %s]", minCompatibilityVersion.String(), minCompVer.String(), maxCompVer.String()))
	}
	return errs
}

func newEffectiveVersion(binaryVersion *version.Version) MutableEffectiveVersion {
	effective := &effectiveVersion{}
	compatVersion := binaryVersion.SubtractMinor(1)
	effective.Set(binaryVersion, binaryVersion, compatVersion)
	return effective
}

func NewEffectiveVersion(binaryVer string) MutableEffectiveVersion {
	if binaryVer == "" {
		return &effectiveVersion{}
	}
	binaryVersion := version.MustParse(binaryVer)
	return newEffectiveVersion(binaryVersion)
}

// DefaultBuildEffectiveVersion returns the MutableEffectiveVersion based on the
// current build information.
func DefaultBuildEffectiveVersion() MutableEffectiveVersion {
	verInfo := baseversion.Get()
	binaryVersion := version.MustParse(verInfo.String()).WithInfo(verInfo)
	if binaryVersion.Major() == 0 && binaryVersion.Minor() == 0 {
		return DefaultKubeEffectiveVersion()
	}
	return newEffectiveVersion(binaryVersion)
}

// DefaultKubeEffectiveVersion returns the MutableEffectiveVersion based on the
// latest K8s release.
func DefaultKubeEffectiveVersion() MutableEffectiveVersion {
	binaryVersion := version.MustParse(baseversion.DefaultKubeBinaryVersion).WithInfo(baseversion.Get())
	return newEffectiveVersion(binaryVersion)
}
