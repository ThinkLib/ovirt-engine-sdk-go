//
// Copyright (c) 2017 Joey <majunjiev@gmail.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"time"

	"github.com/imjoey/go-ovirt"
)

func main() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"

	conn, err := ovirtsdk4.NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("qwer1234").
		Insecure(true).
		Compress(true).
		Timeout(time.Second * 10).
		Build()
	if err != nil {
		fmt.Printf("Make connection failed, reason: %v\n", err)
		return
	}
	defer conn.Close()

	// Get the reference to the "StorageDomains" service
	sdsService := conn.SystemService().StorageDomainsService()

	// Find the storage domain with unregistered VM
	sd := sdsService.List().Search("name=mysd").MustSend().MustStorageDomains().Slice()[0]

	// Locate the service that manages the storage domain, as that is where the action methods are defined
	sdService := sdsService.StorageDomainService(sd.MustId())

	// Locate the service that manages the VMs in storage domain
	sdVmsService := sdService.VmsService()

	// Find the unregistered VM we want to register
	unregVMSlice := sdVmsService.List().Unre

}
