// Copyright 2018 Paul Greenberg (greenpau@outlook.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exporter

import "testing"

func TestNewExporter(t *testing.T) {
	pollTimeout := 2
	apiInventory := "../../assets/ansible/hosts"
	apiVault := "../../assets/ansible/vault.yml"
	apiVaultKey := "../../assets/ansible/vault.key"
	opts := Options{
		Timeout:       pollTimeout,
		InventoryFile: apiInventory,
		VaultFile:     apiVault,
		VaultKeyFile:  apiVaultKey,
	}
	if _, err := NewExporter(opts); err != nil {
		t.Errorf("expected no error, but got %q", err)
	}
}
