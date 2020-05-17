/*
   Network Next. Copyright © 2017 - 2020 Network Next, Inc. All rights reserved.
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runCommand(command string, args []string) bool {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Printf("runCommand error: %v\n", err)
		return false
	}
	return true
}

func bash(format string, args ...interface{}) bool {
	command := fmt.Sprintf(format, args...)
	return runCommand("bash", []string{"-c", command})
}

func main() {

	version := "3.4.5"

	basedir := fmt.Sprintf("next-%s", version)

	fmt.Printf("\nMaking release %s in \"%s\"\n", version, basedir)

	// create the release directory clean
	bash("rm -rf %s", basedir)
	bash("mkdir -p %s", basedir)

	// copy across the premake file
	bash("cp premake5.lua %s", basedir)

	// copy across the README and LICENCE text files
	bash("cp README %s", basedir)
	bash("cp LICENCE %s", basedir)

	// copy across the unit test and soak test
	bash("cp test.cpp %s", basedir)
	bash("cp soak.cpp %s", basedir)

	// copy across the examples
	bash("mkdir -p %s/%s", basedir, "examples")
	bash("cp examples/*.cpp %s/examples", basedir)

	// copy across the include files
	bash("mkdir -p %s/include", basedir)
	bash("cp include/*.h %s/include", basedir)

	// copy across the source code
	bash("mkdir -p %s/%s", basedir, "source")
	bash("cp source/*.h %s/source", basedir)
	bash("cp source/*.cpp %s/source", basedir)

	// copy across the debug win64 libraries
	libdir := "lib/Win64/Debug"
	bash("mkdir -p %s/%s", basedir, libdir)
	bash("cp %s/next-win64-%s.lib %s/%s/next-win64-%s.lib", libdir, version, basedir, libdir, version)
	bash("cp %s/next-win64-%s.pdb %s/%s/next-win64-%s.pdb", libdir, version, basedir, libdir, version)
	bash("cp %s/next-win64-%s.dll %s/%s/next-win64-%s.dll", libdir, version, basedir, libdir, version)
	bash("cp %s/next-win64-%s.exp %s/%s/next-win64-%s.exp", libdir, version, basedir, libdir, version)

	// copy across the release win64 libraries
	libdir = "lib/Win64/Release"
	bash("mkdir -p %s/%s", basedir, libdir)
	bash("cp %s/next-win64-%s.lib %s/%s/next-win64-%s.lib", libdir, version, basedir, libdir, version)
	bash("cp %s/next-win64-%s.pdb %s/%s/next-win64-%s.pdb", libdir, version, basedir, libdir, version)
	bash("cp %s/next-win64-%s.dll %s/%s/next-win64-%s.dll", libdir, version, basedir, libdir, version)
	bash("cp %s/next-win64-%s.exp %s/%s/next-win64-%s.exp", libdir, version, basedir, libdir, version)

	// copy across the debug win32 libraries
	libdir = "lib/Win32/Debug"
	bash("mkdir -p %s/%s", basedir, libdir)
	bash("cp %s/next-win32-%s.lib %s/%s/next-win32-%s.lib", libdir, version, basedir, libdir, version)
	bash("cp %s/next-win32-%s.pdb %s/%s/next-win32-%s.pdb", libdir, version, basedir, libdir, version)
	bash("cp %s/next-win32-%s.dll %s/%s/next-win32-%s.dll", libdir, version, basedir, libdir, version)
	bash("cp %s/next-win32-%s.exp %s/%s/next-win32-%s.exp", libdir, version, basedir, libdir, version)

	// copy across the release win64 libraries
	libdir = "lib/Win32/Release"
	bash("mkdir -p %s/%s", basedir, libdir)
	bash("cp %s/next-win32-%s.lib %s/%s/next-win32-%s.lib", libdir, version, basedir, libdir, version)
	bash("cp %s/next-win32-%s.pdb %s/%s/next-win32-%s.pdb", libdir, version, basedir, libdir, version)
	bash("cp %s/next-win32-%s.dll %s/%s/next-win32-%s.dll", libdir, version, basedir, libdir, version)
	bash("cp %s/next-win32-%s.exp %s/%s/next-win32-%s.exp", libdir, version, basedir, libdir, version)

	// copy across the debug x1 libraries
	libdir = "lib/XBoxOne/Debug"
	bash("mkdir -p %s/%s", basedir, libdir)
	bash("cp %s/next-xboxone-%s.lib %s/%s/next-xboxone-%s.lib", libdir, version, basedir, libdir, version)
	bash("cp %s/next-xboxone-%s.pdb %s/%s/next-xboxone-%s.pdb", libdir, version, basedir, libdir, version)
	bash("cp %s/next-xboxone-%s.dll %s/%s/next-xboxone-%s.dll", libdir, version, basedir, libdir, version)
	bash("cp %s/next-xboxone-%s.exp %s/%s/next-xboxone-%s.exp", libdir, version, basedir, libdir, version)

	// copy across the release x1 libraries
	libdir = "lib/XBoxOne/Release"
	bash("mkdir -p %s/%s", basedir, libdir)
	bash("cp %s/next-xboxone-%s.lib %s/%s/next-xboxone-%s.lib", libdir, version, basedir, libdir, version)
	bash("cp %s/next-xboxone-%s.pdb %s/%s/next-xboxone-%s.pdb", libdir, version, basedir, libdir, version)
	bash("cp %s/next-xboxone-%s.dll %s/%s/next-xboxone-%s.dll", libdir, version, basedir, libdir, version)
	bash("cp %s/next-xboxone-%s.exp %s/%s/next-xboxone-%s.exp", libdir, version, basedir, libdir, version)

	// copy across the debug ps4 libraries
	libdir = "lib/Playstation4/Debug"
	bash("mkdir -p %s/%s", basedir, libdir)
	bash("cp %s/next-ps4-%s_stub_weak.a %s/%s/next-ps4-%s_stub_weak.a", libdir, version, basedir, libdir, version)
	bash("cp %s/next-ps4-%s_stub.a %s/%s/next-ps4-%s_stub.a", libdir, version, basedir, libdir, version)
	bash("cp %s/next-ps4-%s.prx %s/%s/next-ps4-%s.prx", libdir, version, basedir, libdir, version)

	// copy across the release ps4 libraries
	libdir = "lib/Playstation4/Release"
	bash("mkdir -p %s/%s", basedir, libdir)
	bash("cp %s/next-ps4-%s_stub_weak.a %s/%s/next-ps4-%s_stub_weak.a", libdir, version, basedir, libdir, version)
	bash("cp %s/next-ps4-%s_stub.a %s/%s/next-ps4-%s_stub.a", libdir, version, basedir, libdir, version)
	bash("cp %s/next-ps4-%s.prx %s/%s/next-ps4-%s.prx", libdir, version, basedir, libdir, version)

	// copy across the debug nx64 libraries
	libdir = "lib/NintendoSwitch-NX64/Debug"
	bash("mkdir -p %s/%s", basedir, libdir)
	bash("cp %s/next-nx64-%s.nro %s/%s/next-nx64-%s.nro", libdir, version, basedir, libdir, version)
	bash("cp %s/next-nx64-%s.nrr %s/%s/next-nx64-%s.nrr", libdir, version, basedir, libdir, version)
	bash("cp %s/next-nx64-%s.nrs %s/%s/next-nx64-%s.nrs", libdir, version, basedir, libdir, version)

	// copy across the release nx64 libraries
	libdir = "lib/NintendoSwitch-NX64/Release"
	bash("mkdir -p %s/%s", basedir, libdir)
	bash("cp %s/next-nx64-%s.nro %s/%s/next-nx64-%s.nro", libdir, version, basedir, libdir, version)
	bash("cp %s/next-nx64-%s.nrr %s/%s/next-nx64-%s.nrr", libdir, version, basedir, libdir, version)
	bash("cp %s/next-nx64-%s.nrs %s/%s/next-nx64-%s.nrs", libdir, version, basedir, libdir, version)

	// copy across the debug nx32 libraries
	libdir = "lib/NintendoSwitch-NX32/Debug"
	bash("mkdir -p %s/%s", basedir, libdir)
	bash("cp %s/next-nx32-%s.nro %s/%s/next-nx32-%s.nro", libdir, version, basedir, libdir, version)
	bash("cp %s/next-nx32-%s.nrr %s/%s/next-nx32-%s.nrr", libdir, version, basedir, libdir, version)
	bash("cp %s/next-nx32-%s.nrs %s/%s/next-nx32-%s.nrs", libdir, version, basedir, libdir, version)

	// copy across the release nx32 libraries
	libdir = "lib/NintendoSwitch-NX32/Release"
	bash("mkdir -p %s/%s", basedir, libdir)
	bash("cp %s/next-nx32-%s.nro %s/%s/next-nx32-%s.nro", libdir, version, basedir, libdir, version)
	bash("cp %s/next-nx32-%s.nrr %s/%s/next-nx32-%s.nrr", libdir, version, basedir, libdir, version)
	bash("cp %s/next-nx32-%s.nrs %s/%s/next-nx32-%s.nrs", libdir, version, basedir, libdir, version)

	// build manifest

	manifest := make([]string, 0)

	manifest = append(manifest, fmt.Sprintf("%s/premake5.lua", basedir))

	manifest = append(manifest, fmt.Sprintf("%s/README", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/LICENCE", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/test.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/soak.cpp", basedir))

	manifest = append(manifest, fmt.Sprintf("%s/examples/simple_client.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/examples/simple_server.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/examples/upgraded_client.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/examples/upgraded_server.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/examples/complex_client.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/examples/complex_server.cpp", basedir))

	manifest = append(manifest, fmt.Sprintf("%s/include/next.h", basedir))

	manifest = append(manifest, fmt.Sprintf("%s/source/next.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_mac.h", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_mac.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_linux.h", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_linux.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_windows.h", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_windows.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_windows.h", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_windows.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_ios.h", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_ios.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_ps4.h", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_ps4.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_xboxone.h", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_xboxone.cpp", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_switch.h", basedir))
	manifest = append(manifest, fmt.Sprintf("%s/source/next_switch.cpp", basedir))

	manifest = append(manifest, fmt.Sprintf("%s/lib/Win64/Debug/next-win64-%s.lib", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Win64/Debug/next-win64-%s.pdb", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Win64/Debug/next-win64-%s.dll", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Win64/Debug/next-win64-%s.exp", basedir, version))

	manifest = append(manifest, fmt.Sprintf("%s/lib/Win64/Release/next-win64-%s.lib", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Win64/Release/next-win64-%s.pdb", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Win64/Release/next-win64-%s.dll", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Win64/Release/next-win64-%s.exp", basedir, version))

	manifest = append(manifest, fmt.Sprintf("%s/lib/Win32/Debug/next-win32-%s.lib", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Win32/Debug/next-win32-%s.pdb", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Win32/Debug/next-win32-%s.dll", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Win32/Debug/next-win32-%s.exp", basedir, version))

	manifest = append(manifest, fmt.Sprintf("%s/lib/Win32/Release/next-win32-%s.lib", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Win32/Release/next-win32-%s.pdb", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Win32/Release/next-win32-%s.dll", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Win32/Release/next-win32-%s.exp", basedir, version))

	manifest = append(manifest, fmt.Sprintf("%s/lib/XBoxOne/Debug/next-xboxone-%s.lib", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/XBoxOne/Debug/next-xboxone-%s.pdb", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/XBoxOne/Debug/next-xboxone-%s.dll", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/XBoxOne/Debug/next-xboxone-%s.exp", basedir, version))

	manifest = append(manifest, fmt.Sprintf("%s/lib/XBoxOne/Release/next-xboxone-%s.lib", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/XBoxOne/Release/next-xboxone-%s.pdb", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/XBoxOne/Release/next-xboxone-%s.dll", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/XBoxOne/Release/next-xboxone-%s.exp", basedir, version))

	manifest = append(manifest, fmt.Sprintf("%s/lib/Playstation4/Debug/next-ps4-%s_stub_weak.a", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Playstation4/Debug/next-ps4-%s_stub.a", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Playstation4/Debug/next-ps4-%s.prx", basedir, version))

	manifest = append(manifest, fmt.Sprintf("%s/lib/Playstation4/Release/next-ps4-%s_stub_weak.a", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Playstation4/Release/next-ps4-%s_stub.a", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/Playstation4/Release/next-ps4-%s.prx", basedir, version))

	manifest = append(manifest, fmt.Sprintf("%s/lib/NintendoSwitch-NX64/Debug/next-nx64-%s.nro", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/NintendoSwitch-NX64/Debug/next-nx64-%s.nrr", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/NintendoSwitch-NX64/Debug/next-nx64-%s.nrs", basedir, version))

	manifest = append(manifest, fmt.Sprintf("%s/lib/NintendoSwitch-NX64/Release/next-nx64-%s.nro", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/NintendoSwitch-NX64/Release/next-nx64-%s.nrr", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/NintendoSwitch-NX64/Release/next-nx64-%s.nrs", basedir, version))

	manifest = append(manifest, fmt.Sprintf("%s/lib/NintendoSwitch-NX32/Debug/next-nx32-%s.nro", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/NintendoSwitch-NX32/Debug/next-nx32-%s.nrr", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/NintendoSwitch-NX32/Debug/next-nx32-%s.nrs", basedir, version))

	manifest = append(manifest, fmt.Sprintf("%s/lib/NintendoSwitch-NX32/Release/next-nx32-%s.nro", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/NintendoSwitch-NX32/Release/next-nx32-%s.nrr", basedir, version))
	manifest = append(manifest, fmt.Sprintf("%s/lib/NintendoSwitch-NX32/Release/next-nx32-%s.nrs", basedir, version))

	fmt.Printf("\nManifest:\n\n" )
	for _, file := range manifest {
		fmt.Printf("    %s\n", file)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Printf("\n    ^---- error: missing file!\n\n")
			os.Exit(1)
		}
	}

	fmt.Printf("\nSuccessfully made release %s\n\n", version)

}
