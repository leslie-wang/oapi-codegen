package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/leslie-wang/oapi-codegen/pkg/codegen"
	"github.com/leslie-wang/oapi-codegen/pkg/util"
	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
)

func errExit(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
func main() {
	swagger, err := util.LoadSwagger(os.Args[1])
	if err != nil {
		errExit("error loading swagger spec\n: %s", err)
	}
	_, err = generate(swagger, "experiment")
	if err != nil {
		errExit(err.Error())
	}
	//fmt.Println(code)
}

func generate(swagger *openapi3.Swagger, packageName string) (string, error) {
	// This creates the golang templates text package
	t := template.New("oapi-codegen").Funcs(codegen.TemplateFunctions)
	// This parses all of our own template files into the template object
	// above
	t, err := Parse(t)
	if err != nil {
		return "", errors.Wrap(err, "error parsing oapi-codegen templates")
	}

	ops, err := codegen.OperationDefinitions(swagger)
	if err != nil {
		return "", errors.Wrap(err, "error creating operation definitions")
	}

	if err := generateAPI("apigen", "apigen/apigen.go", t, swagger, ops); err != nil {
		return "", errors.Wrap(err, "error creating API file")
	}
	if err := generateImpl("impl", "impl/impl.go", t, swagger, ops,
		[]string{`"github.com/leslie-wang/oapi-codegen/cmd/skeleton/apigen"`}); err != nil {
		return "", errors.Wrap(err, "error creating API file")
	}
	if err := generateRouter("main", "apigen-main/router.go", t, swagger, ops,
		[]string{`"github.com/leslie-wang/oapi-codegen/cmd/skeleton/apigen"`}); err != nil {
		return "", errors.Wrap(err, "error creating API file")
	}
	if err := generateMain("main", "apigen-main/main.go", t, swagger, ops,
		[]string{`echomiddleware "github.com/labstack/echo/v4/middleware"`,
			`"github.com/leslie-wang/oapi-codegen/cmd/skeleton/apigen"`,
			`"github.com/leslie-wang/oapi-codegen/cmd/skeleton/impl"`}); err != nil {
		return "", errors.Wrap(err, "error creating API file")
	}
	return "", nil
}

func writeFile(filename, codes string) error {
	dir, _ := filepath.Split(filename)
	if dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	outBytes, err := imports.Process(filename, []byte(codes), nil)
	if err != nil {
		fmt.Println(codes)
		return errors.Wrap(err, "error formatting Go code")
	}
	return ioutil.WriteFile(filename, []byte(outBytes), 0755)
}

func generateAPI(packageName, filename string, t *template.Template, swagger *openapi3.Swagger,
	ops []codegen.OperationDefinition) error {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	importsOut, err := codegen.GenerateImports(t, nil, packageName)
	if err != nil {
		return errors.Wrap(err, "error generating imports")
	}

	_, err = w.WriteString(importsOut)
	if err != nil {
		return errors.Wrap(err, "error writing imports")
	}

	typeDefinitions, err := codegen.GenerateTypeDefinitions(t, swagger, ops, nil)
	if err != nil {
		return errors.Wrap(err, "error generating type definitions")
	}

	_, err = w.WriteString(typeDefinitions)
	if err != nil {
		return errors.Wrap(err, "error writing type definitions")
	}

	err = t.ExecuteTemplate(w, "responses-type.tmpl", ops)
	if err != nil {
		return errors.Wrap(err, "Error generating responses types")
	}

	serverIntf, err := codegen.GenerateServerInterface(t, ops)
	if err != nil {
		return errors.Wrap(err, "Error generating server types and interface: %s")
	}

	_, err = w.WriteString(serverIntf)
	if err != nil {
		return errors.Wrap(err, "error writing type definitions")
	}

	err = w.Flush()
	if err != nil {
		return errors.Wrap(err, "error flushing output buffer")
	}

	// remove any byte-order-marks which break Go-Code
	return writeFile(filename, codegen.SanitizeCode(buf.String()))
}

func generateImpl(packageName, filename string, t *template.Template, swagger *openapi3.Swagger,
	ops []codegen.OperationDefinition, externalImports []string) error {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	importsOut, err := codegen.GenerateImports(t, externalImports, packageName)
	if err != nil {
		return errors.Wrap(err, "error generating imports")
	}

	_, err = w.WriteString(importsOut)
	if err != nil {
		return errors.Wrap(err, "error writing imports")
	}

	err = t.ExecuteTemplate(w, "impl.tmpl", ops)
	if err != nil {
		return errors.Wrap(err, "error generating impl skeleton")
	}

	err = w.Flush()
	if err != nil {
		return errors.Wrap(err, "error flushing output buffer")
	}

	// remove any byte-order-marks which break Go-Code
	return writeFile(filename, codegen.SanitizeCode(buf.String()))
}

func generateRouter(packageName, filename string, t *template.Template, swagger *openapi3.Swagger,
	ops []codegen.OperationDefinition, externalImports []string) error {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	importsOut, err := codegen.GenerateImports(t, externalImports, packageName)
	if err != nil {
		return errors.Wrap(err, "error generating imports")
	}

	_, err = w.WriteString(importsOut)
	if err != nil {
		return errors.Wrap(err, "error writing imports")
	}

	wrappers, err := codegen.GenerateWrappers(t, ops)
	if err != nil {
		return fmt.Errorf("Error generating handler wrappers: %s", err)
	}

	_, err = w.WriteString(wrappers)
	if err != nil {
		return errors.Wrap(err, "error writing wrappers")
	}

	register, err := codegen.GenerateRegistration(t, ops)
	if err != nil {
		return fmt.Errorf("Error generating handler registration: %s", err)
	}

	_, err = w.WriteString(register)
	if err != nil {
		return errors.Wrap(err, "error writing register")
	}

	err = w.Flush()
	if err != nil {
		return errors.Wrap(err, "error flushing output buffer")
	}

	// remove any byte-order-marks which break Go-Code
	return writeFile(filename, codegen.SanitizeCode(buf.String()))
}

func generateMain(packageName, filename string, t *template.Template, swagger *openapi3.Swagger,
	ops []codegen.OperationDefinition, externalImports []string) error {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	importsOut, err := codegen.GenerateImports(t, externalImports, packageName)
	if err != nil {
		return errors.Wrap(err, "error generating imports")
	}

	_, err = w.WriteString(importsOut)
	if err != nil {
		return errors.Wrap(err, "error writing imports")
	}

	err = t.ExecuteTemplate(w, "main.tmpl", ops)
	if err != nil {
		return errors.Wrap(err, "error generating impl skeleton")
	}
	err = w.Flush()
	if err != nil {
		return errors.Wrap(err, "error flushing output buffer")
	}

	// remove any byte-order-marks which break Go-Code
	return writeFile(filename, codegen.SanitizeCode(buf.String()))
}
