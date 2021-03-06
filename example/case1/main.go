// Copyright 2022 helloshaohua <wu.shaohua@foxmail.com>;
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

package main

import (
	"embed"
	"fmt"
	"github.com/coolstina/embedsfs"
	"log"
)

//go:embed embeds
var embeds embed.FS

func main() {
	// Initialize embedsfs instance.
	ins := embedsfs.NewEmbedsFS(embeds, embedsfs.WithPath("embeds"))

	// Get golang internal structure embed.FS
	embeds := ins.Embeds()
	fmt.Println(embeds)

	// Get golang embeds filesystem filename
	name := ins.Filename("readme.txt")
	fmt.Println(name)

	// Get golang embeds filesystem filename
	data, err := ins.Content("readme.txt")
	if err != nil {
		log.Fatalf("get embeds file content to failure: %v", err)
	}

	fmt.Println(string(data))
}
