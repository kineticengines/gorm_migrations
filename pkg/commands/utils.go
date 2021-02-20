package commands

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kineticengines/gorm-migrations/pkg/definitions"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
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
		return nil, definitions.ErrFailedToLoadGormgxFile
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

// readModelsFromPath read models defined in the path defined.
// construct type info tho assert whether a model is eligible for migration procedure
func readModelsFromPath(path string) (*types.Package, error) {
	return readFileSet(path)
}

func readIntentModels(modelsPkgs *[]*types.Package, paths []string, verbose bool) error {
	printVerbose(verbose, log.InfoLevel, "Reading intent models")
	for _, path := range paths {
		pkg, err := readModelsFromPath(path)
		if err != nil {
			return err
		}
		*modelsPkgs = append(*modelsPkgs, pkg)
	}
	return nil
}

func readFileSet(path string) (*types.Package, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("parse error %w", err)
	}
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("types", fset, []*ast.File{f}, nil)
	if err != nil {
		return nil, fmt.Errorf("type check  error %w", err)
	}

	return pkg, nil
}

func readInterfaceFile() []*types.Named {
	// read interface definition file. Will be used to assert if a model satisfies it
	pkgI, err := readFileSet("pkg/definitions/interface.go")
	if err != nil {
		return nil
	}
	var allNamedInteraface []*types.Named
	for _, name := range pkgI.Scope().Names() {
		if obj, ok := pkgI.Scope().Lookup(name).(*types.TypeName); ok {
			allNamedInteraface = append(allNamedInteraface, obj.Type().(*types.Named))
		}
	}
	if !types.IsInterface(allNamedInteraface[0]) {
		return nil
	}
	return allNamedInteraface
}

func analyzePkg(pkg *types.Package, verbose bool) error {
	printVerbose(verbose, log.InfoLevel, "Analyzing package scopes")
	scope := pkg.Scope()
	printVerbose(verbose, log.InfoLevel, fmt.Sprintf("Analyzing package scopes : scope size: %v", scope.Len()))

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
		}
	}

	typeMap := make(map[*types.Named]map[string]definitions.FieldMeta)
	for _, v := range validObjects {
		t := extractFieldsFromStruct(v)
		typeMap[v] = t
	}
	fmt.Println(typeMap)
	return nil
}

func extractFieldsFromStruct(v *types.Named) map[string]definitions.FieldMeta {
	fieldsMap := make(map[string]definitions.FieldMeta)
	u := v.Underlying().(*types.Struct)
	for i := 0; i < u.NumFields(); i++ {
		fieldName := u.Field(i).Name()
		meta := definitions.FieldMeta{}
		meta.Tag = u.Tag(i)
		ft := computeBasicType(u.Field(i).Type().Underlying())
		// if ft == definitions.Nil {
		// 	log.Panicf("%v; got %T", definitions.ErrNilType)
		// }
		meta.FieldType = ft
		fieldsMap[fieldName] = meta
	}
	return fieldsMap
}

func computeBasicType(u types.Type) definitions.BasicType {
	switch x := u.(type) {
	case *types.Struct:
		if x.Field(0).Name() == "wall" || x.Field(0).Name() == "ext" || x.Field(0).Name() == "loc" {
			// todo
		}

	case *types.Pointer:
		elem := x.Underlying().(*types.Pointer).Elem()
		return computeBasicType(elem)
	case *types.Basic:
		switch x.Kind() {
		case types.String:
			return definitions.String
		case types.Bool:
			return definitions.Bool
		}
	default:
		log.Infoln(x)
	}
	return definitions.Nil
}
