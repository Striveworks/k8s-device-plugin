/**
# Copyright (c) NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package nvcdi

import (
	"github.com/container-orchestrated-devices/container-device-interface/pkg/cdi"
	"github.com/container-orchestrated-devices/container-device-interface/specs-go"
	"gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvlib/device"
)

const (
	// ModeAuto configures the CDI spec generator to automatically detect the system configuration
	ModeAuto = "auto"
	// ModeNvml configures the CDI spec generator to use the NVML library.
	ModeNvml = "nvml"
	// ModeWsl configures the CDI spec generator to generate a WSL spec.
	ModeWsl = "wsl"
)

// Interface defines the API for the nvcdi package
type Interface interface {
	GetCommonEdits() (*cdi.ContainerEdits, error)
	GetAllDeviceSpecs() ([]specs.Device, error)
	GetGPUDeviceEdits(device.Device) (*cdi.ContainerEdits, error)
	GetGPUDeviceSpecs(int, device.Device) (*specs.Device, error)
	GetMIGDeviceEdits(device.Device, device.MigDevice) (*cdi.ContainerEdits, error)
	GetMIGDeviceSpecs(int, device.Device, int, device.MigDevice) (*specs.Device, error)
}