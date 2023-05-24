package main

import (
	"flag"
	"fmt"
	"github.com/catalystsquad/protoc-gen-go-gorm/plugin"
	"github.com/golang/glog"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	enumsAsInts = flag.Bool("enums_as_ints", false, "render enums as integers as opposed to strings")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gp *protogen.Plugin) error {
		gp.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		opts := plugin.PluginOptions{
			EnumsAsInts: *enumsAsInts,
		}

		for _, name := range gp.Request.FileToGenerate {
			f := gp.FilesByPath[name]

			if len(f.Messages) == 0 {
				glog.Infof("Skipping %s, no messages", name)
				continue
			}

			glog.Infof("Processing %s", name)
			glog.Infof("Generating %s\n", fmt.Sprintf("%s.pb.gorm.go", f.GeneratedFilenamePrefix))

			gf := gp.NewGeneratedFile(fmt.Sprintf("%s.pb.gorm.go", f.GeneratedFilenamePrefix), f.GoImportPath)

			err := plugin.ApplyTemplate(gf, f, opts)
			if err != nil {
				gf.Skip()
				gp.Error(err)
				continue
			}
		}

		return nil
	})
}
