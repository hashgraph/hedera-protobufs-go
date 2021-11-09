package main

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	// which module are we building
	moduleName := os.Args[1]

	// get project directory

	// nolint:dogsled
	_, buildFilename, _, _ := runtime.Caller(0)
	projectDir := path.Join(buildFilename, "../../../..")

	// remove all existing files

	removeAllWithExt(projectDir, moduleName, ".pb.go")

	// invoke the build command for this module

	switch moduleName {
	case "services":
		buildServices(projectDir)

	case "mirror":
		buildMirror(projectDir)

	case "streams":
		buildStreams(projectDir)
	}
}

func buildServices(dir string) {
	// collect all files from proto/services/_
	var servicesProtoFiles []string

	var servicesModuleDecls []string

	err := filepath.Walk(path.Join(dir, "proto/services"), func(filename string, info fs.FileInfo, err error) error {
		if !strings.HasSuffix(filename, ".proto") {
			return nil
		}

		pathFromRoot := strings.TrimPrefix(filename, dir+"/")
		pathBase := path.Base(filename)
		servicesProtoFiles = append(servicesProtoFiles, pathFromRoot)

		// the -M argument generation is what allows us to avoid
		// requiring "option go_package"

		servicesModuleDecls = append(servicesModuleDecls,
			fmt.Sprintf("--go_opt=M%v=github.com/hashgraph/hedera-protobufs-go/services", pathBase),
			fmt.Sprintf("--go-grpc_opt=M%v=github.com/hashgraph/hedera-protobufs-go/services", pathBase),
		)

		return nil
	})
	if err != nil {
		panic(err)
	}

	// generate proto files for services code

	cmdArguments := []string{
		"--go_out=services/",
		"--go_opt=paths=source_relative",
		"--go-grpc_out=services/",
		"--go-grpc_opt=paths=source_relative",
		"-Iproto/services",
	}

	cmdArguments = append(cmdArguments, servicesModuleDecls...)
	cmdArguments = append(cmdArguments, servicesProtoFiles...)

	cmd := exec.Command("protoc", cmdArguments...)
	cmd.Dir = dir

	mustRunCommand(cmd)
	renamePackageDeclGrpcFiles(dir, "proto", "services")
}

func buildMirror(dir string) {
	cmd := exec.Command("protoc",
		"--go_out=mirror/",
		"--go_opt=Mbasic_types.proto=github.com/hashgraph/hedera-protobufs-go/services",
		"--go_opt=Mtimestamp.proto=github.com/hashgraph/hedera-protobufs-go/services",
		"--go_opt=Mconsensus_submit_message.proto=github.com/hashgraph/hedera-protobufs-go/services",
		"--go_opt=Mconsensus_service.proto=github.com/hashgraph/hedera-protobufs-go/mirror",
		"--go_opt=paths=source_relative",
		"--go-grpc_out=mirror/",
		"--go-grpc_opt=Mbasic_types.proto=github.com/hashgraph/hedera-protobufs-go/services",
		"--go-grpc_opt=Mtimestamp.proto=github.com/hashgraph/hedera-protobufs-go/services",
		"--go-grpc_opt=Mconsensus_submit_message.proto=github.com/hashgraph/hedera-protobufs-go/services",
		"--go-grpc_opt=Mconsensus_service.proto=github.com/hashgraph/hedera-protobufs-go/mirror",
		"--go-grpc_opt=paths=source_relative",
		"-Iproto/mirror",
		"-Iproto/services",
		"proto/mirror/consensus_service.proto",
	)

	cmd.Dir = dir

	mustRunCommand(cmd)
	renamePackageDeclGrpcFiles(dir, "com_hedera_mirror_api_proto", "mirror")
}

func buildStreams(dir string) {
	cmd := exec.Command("protoc",
		"--go_out=streams/",
		"--go_opt=Mbasic_types.proto=github.com/hashgraph/hedera-protobufs-go/services",
		"--go_opt=Mtimestamp.proto=github.com/hashgraph/hedera-protobufs-go/services",
		"--go_opt=Maccount_balance_file.proto=github.com/hashgraph/hedera-protobufs-go/streams",
		"--go_opt=paths=source_relative",
		"--go-grpc_out=streams/",
		"--go-grpc_opt=Mbasic_types.proto=github.com/hashgraph/hedera-protobufs-go/services",
		"--go-grpc_opt=Mtimestamp.proto=github.com/hashgraph/hedera-protobufs-go/services",
		"--go-grpc_opt=Maccount_balance_file.proto=github.com/hashgraph/hedera-protobufs-go/streams",
		"--go-grpc_opt=paths=source_relative",
		"-Iproto/streams",
		"-Iproto/services",
		"proto/streams/account_balance_file.proto",
	)

	cmd.Dir = dir

	mustRunCommand(cmd)
}

func mustRunCommand(cmd *exec.Cmd) {
	_, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			fmt.Print(string(exitErr.Stderr))
			os.Exit(exitErr.ExitCode())
		}

		panic(err)
	}
}

func removeAllWithExt(dir string, module string, ext string) {
	err := filepath.Walk(path.Join(dir, module), func(filename string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(filename, ext) {
			err := os.Remove(filename)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}

func renamePackageDeclGrpcFiles(dir string, oldPackage string, newPackage string) {
	err := filepath.Walk(path.Join(dir, newPackage), func(filename string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(filename, "_grpc.pb.go") {
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				return err
			}

			contents := string(data)
			contents = strings.Replace(contents,
				fmt.Sprintf("package %s", oldPackage),
				fmt.Sprintf("package %s", newPackage),
				1,
			)

			return ioutil.WriteFile(filename, []byte(contents), info.Mode())
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}
