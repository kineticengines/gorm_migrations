package commands

import (
	"fmt"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kineticengines/gorm-migrations/pkg/definitions"
	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"

	"golang.org/x/tools/go/packages"
)

// gormgxFilePath ...
func gormgxFilePath() (*string, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, definitions.ErrUnableToGetWorkingDirectory
	}
	file := filepath.Join(path, definitions.GormgxYamlFileName)
	return &file, nil
}

func readYamlToconfig() (*definitions.Config, error) {
	yamlPath, err := gormgxFilePath()
	if err != nil {
		return nil, definitions.ErrFailedToFetchGormgxPath
	}
	content, err := ioutil.ReadFile(*yamlPath)
	if err != nil {
		return nil, definitions.ErrFailedToReadGormgxFile

	}
	var cfg definitions.Config
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return nil, definitions.ErrFailedToUnmarshalGormgxFile

	}
	return &cfg, nil
}

func checkIntialMIgrationExists() bool {
	path, err := os.Getwd()
	if err != nil {
		return false
	}
	initMigrationPath := filepath.Join(path, definitions.DefaultMIgrationsPath, definitions.IntialMigrationFileName)
	if _, err := os.Stat(initMigrationPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func printVerbose(verbose bool, logLevel log.Level, message interface{}) {
	if verbose {
		switch logLevel {
		case log.PanicLevel:
			log.Panicf("%v", message)
		case log.FatalLevel:
			log.Fatalf("%v", message)

		case log.ErrorLevel:
			log.Errorf("%v", message)

		case log.WarnLevel:
			log.Warnf("%v", message)

		case log.InfoLevel:
			log.Infof("%v", message)

		case log.DebugLevel:
			log.Debugf("%v", message)

		case log.TraceLevel:
			log.Tracef("%v", message)

		}
	}
}

func readIntentModels(modelsPkgs *[]*types.Package, paths []string, verbose bool) error {
	printVerbose(verbose, log.InfoLevel, "Reading intent models")
	for _, path := range paths {
		pkg, err := ReadModelsFromPath(path)
		if err != nil {
			return err
		}
		*modelsPkgs = append(*modelsPkgs, pkg)
	}
	return nil
}

// ReadModelsFromPath read models defined in the path defined.
// construct type info tho assert whether a model is eligible for migration procedure
func ReadModelsFromPath(path string) (*types.Package, error) {
	return readFileSet(path)
}

func readFileSet(path string) (*types.Package, error) {
	cfg := &packages.Config{Mode: packages.NeedName | packages.NeedTypesInfo | packages.NeedTypes}
	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		log.Fatalf("package load error: %v", err)
	}
	pkg := pkgs[0].Types
	return pkg, nil
}

func readInterfaceFile() []*types.Named {
	// read interface definition file. Will be used to assert if a model satisfies it
	pkgI, err := readFileSet("pkg/definitions/interface.go")
	if err != nil {
		return nil
	}
	var allNamedInterface []*types.Named
	for _, name := range pkgI.Scope().Names() {
		if obj, ok := pkgI.Scope().Lookup(name).(*types.TypeName); ok {
			allNamedInterface = append(allNamedInterface, obj.Type().(*types.Named))
		}
	}
	if !types.IsInterface(allNamedInterface[0]) {
		return nil
	}
	return allNamedInterface
}

func analyzePkg(pkg *types.Package, verbose bool) error {
	printVerbose(verbose, log.InfoLevel, "Analyzing package scopes")

	// Find all named types at package level.
	var allNamed []*types.Named
	for _, name := range pkg.Scope().Names() {
		if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok {
			allNamed = append(allNamed, obj.Type().(*types.Named))
		}
	}

	validObjects := []*types.Named{}
	allNamedInteraface := readInterfaceFile()
	for _, T := range allNamed {
		if types.AssignableTo(types.NewPointer(T), allNamedInteraface[0]) {
			validObjects = append(validObjects, T)
		} else {
			printVerbose(verbose, log.WarnLevel, fmt.Sprintf("object %v does not satify interface GormModel", T))
		}
	}

	typeMap := make(map[*types.Named]*TableTree)
	for _, v := range validObjects {
		t := nameTypeFieldsMeta(v)
		typeMap[v] = t
	}
	log.Printf("type map %v", typeMap)
	return nil
}

func nameTypeFieldsMeta(v *types.Named) *TableTree {
	u := v.Underlying().(*types.Struct)
	tree := new(TableTree)
	tree.AddNodes(u)
	return tree
}

// func extractFieldsFromStruct(u *types.Struct) map[string]definitions.FieldMeta {
// 	fieldsMap := make(map[string]definitions.FieldMeta)
// 	for i := 0; i < u.NumFields(); i++ {
// 		fieldName := u.Field(i).Name()
// 		meta := definitions.FieldMeta{}
// 		meta.Tag = u.Tag(i)
// 		ft := computeBasicType(u.Field(i).Type().Underlying())
// 		// if ft == definitions.Nil {
// 		// 	log.Panicf("%v; got %T", definitions.ErrNilType)
// 		// }
// 		meta.FieldType = ft
// 		fieldsMap[fieldName] = meta
// 	}
// 	return fieldsMap
// }

// func computeBasicType(u types.Type) definitions.BasicType {
// 	switch x := u.(type) {
// 	case *types.Struct:
// 		// todo
// 		log.Infof("type struct %v", x)
// 	case *types.Pointer:
// 		elem := x.Underlying().(*types.Pointer).Elem()
// 		return computeBasicType(elem)
// 	case *types.Basic:
// 		switch x.Kind() {
// 		case types.Int:
// 			return definitions.Int
// 		case types.Int8:
// 			return definitions.Int8
// 		case types.Int16:
// 			return definitions.Int16
// 		case types.Int32:
// 			return definitions.Int32
// 		case types.Int64:
// 			return definitions.Int64
// 		case types.Uint:
// 			return definitions.Uint
// 		case types.Uint8:
// 			return definitions.Uint8
// 		case types.Uint16:
// 			return definitions.Uint16
// 		case types.Uint32:
// 			return definitions.Uint32
// 		case types.Uint64:
// 			return definitions.Uint64
// 		case types.Float32:
// 			return definitions.Float32
// 		case types.Float64:
// 			return definitions.Float64
// 		case types.Complex64:
// 			return definitions.Complex64
// 		case types.Complex128:
// 			return definitions.Complex128
// 		case types.String:
// 			return definitions.String
// 		case types.Bool:
// 			return definitions.Bool
// 		}
// 	default:
// 		log.Infof("basiv type %v", x)
// 	}
// 	return definitions.Nil
// }
