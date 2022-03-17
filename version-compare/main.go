package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	v1 := "v0.9.1-2"
	v2 := "v0.9.1-2"

	upgrade := false

	noV1Str := strings.TrimLeft(v1, "v")
	fmt.Println("noV1Str:", noV1Str)
	v1Slice := strings.Split(noV1Str, ".")
	fmt.Println("v1Slice:", v1Slice)
	v1NoDash := strings.Split(v1Slice[2], "-")
	fmt.Printf("v1NoDash: %v\nv1NoDash Length: %v\n", v1NoDash, len(v1NoDash))

	v1Major, _ := strconv.Atoi(v1Slice[0])
	v1Minor, _ := strconv.Atoi(v1Slice[1])
	v1Patch, _ := strconv.Atoi(v1NoDash[0])
	var v1Build int
	if len(v1NoDash) != 1 {
		v1Build, _ = strconv.Atoi(v1NoDash[1])
	}

	fmt.Printf("v1Major: %v\nv1Minor: %v\nv1Patch: %v\nv1Build: %v\n", v1Major, v1Minor, v1Patch, v1Build)

	fmt.Println()

	noV2Str := strings.TrimLeft(v2, "v")
	fmt.Println("noV2Str:", noV2Str)
	v2Slice := strings.Split(noV2Str, ".")
	fmt.Println("v2Slice:", v2Slice)
	v2NoDash := strings.Split(v2Slice[2], "-")
	fmt.Printf("v2NoDash: %v\nv2NoDash Length: %v\n", v2NoDash, len(v2NoDash))

	v2Major, _ := strconv.Atoi(v2Slice[0])
	v2Minor, _ := strconv.Atoi(v2Slice[1])
	v2Patch, _ := strconv.Atoi(v2NoDash[0])
	var v2Build int
	if len(v2NoDash) != 1 {
		v2Build, _ = strconv.Atoi(v2NoDash[1])
	}

	fmt.Printf("v2Major: %v\nv2Minor: %v\nv2Patch: %v\nv2Build: %v\n\n", v2Major, v2Minor, v2Patch, v2Build)

	if v1Build == 0 {
		fmt.Println("v1Build is empty")
	}
	if v2Build == 0 {
		fmt.Println("v2Build is empty")
	}

	fmt.Println()
	fmt.Printf("%v < %v: %v\n", v1Minor, v2Minor, v1Minor < v2Minor)
	fmt.Printf("%v == %v: %v\n", v1Major, v2Major, v1Major == v2Major)
	if v1Major < v2Major {
		fmt.Println("Major: v2 is new version")
	}
	if v1Major == v2Major {
		fmt.Printf("v2 major version (%v) is same as v1 version (%v). Checking minor versions...\n", v2Major, v1Major)
		if v1Minor < v2Minor {
			fmt.Println("Major: v2 is new version")
			upgrade = true
		}
		if v1Minor == v2Minor {
			fmt.Printf("v2 minor version (%v) is same as v1 version (%v). Checking patch versions...\n", v2Minor, v1Minor)
			if v1Patch < v2Patch {
				fmt.Println("Patch: v2 is new version")
				upgrade = true
			}
			if v1Patch == v2Patch {
				fmt.Printf("v2 patch version (%v) is same as v1 version (%v). Checking build versions...\n", v2Patch, v1Patch)
				if v1Build < v2Build {
					fmt.Println("Build: v2 is new version")
					upgrade = true
				}
			}
		}
	}
	fmt.Printf("Upgrade: %v\n", upgrade)
}
