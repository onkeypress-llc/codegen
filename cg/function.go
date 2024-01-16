package cg

// import (
// 	"fmt"
// 	"strings"
// )

// type TemplateFunction struct {
// 	Imports      *Imports
// 	Comment      string
// 	FunctionName string
// 	ReturnTypes  []*Type
// 	Args         []*TemplateFunctionArg
// 	Body         *TemplateFunctionBody
// }

// type TemplateFunctionBody struct {
// 	Content string
// }

// type TemplateFunctionArg struct {
// 	Name string
// 	Type *Type
// }

// func (b *TemplateFunctionBody) String() string {
// 	return b.Content
// }

// func (f *TemplateFunction) CommentLine() string {
// 	if f.Comment == "" {
// 		return f.Comment
// 	}
// 	return fmt.Sprintf("//%s", f.Comment)
// }

// func (f *TemplateFunction) DeclarationArgString() string {
// 	values := make([]string, len(f.Args))
// 	for i := range f.Args {
// 		arg := f.Args[i]
// 		values[i] = fmt.Sprintf("%s %s", arg.Name, arg.Type.String())
// 	}
// 	return strings.Join(values, ", ")
// }

// func (f *TemplateFunction) DeclarationReturnTypeString() string {
// 	length := len(f.ReturnTypes)
// 	if length < 1 {
// 		return ""
// 	}
// 	values := make([]string, length)

// 	for i := range f.ReturnTypes {
// 		t := f.ReturnTypes[i]
// 		values[i] = t.String()
// 	}
// 	if length == 1 {
// 		return values[0]
// 	}
// 	return fmt.Sprintf("(%s)", strings.Join(values, ", "))
// }

// func (f *TemplateFunction) DeclarationLine() string {
// 	return fmt.Sprintf("func %s(%s) %s {", f.FunctionName, f.DeclarationArgString(), f.DeclarationReturnTypeString())
// }

// func (f *TemplateFunction) ClosingLine() string {
// 	return "}"
// }

// func (f *TemplateFunction) ImplementationString() string {
// 	lines := []string{
// 		f.CommentLine(),
// 		f.DeclarationLine(),

// 		f.Body.String(),
// 		f.ClosingLine(),
// 	}
// 	out := []string{}
// 	for i := range lines {
// 		line := lines[i]
// 		if line == "" {
// 			continue
// 		}
// 		out = append(out, line)
// 	}

// 	return strings.Join(out, "\n")
// }

// func newTemplateFunction(name string, returnTypes []*Type, args []*TemplateFunctionArg, body *TemplateFunctionBody) (*TemplateFunction, error) {
// 	if body == nil {
// 		body = &TemplateFunctionBody{Content: "panic(\"Not implemented\")"}
// 	}
// 	importSet, err := createImportSetForFunctionComponents(args, returnTypes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &TemplateFunction{
// 		FunctionName: name,
// 		ReturnTypes:  returnTypes,
// 		Args:         args,
// 		Body:         body,
// 		Imports:      importSet,
// 	}, nil
// }

// func createImportSetForFunction(f *TemplateFunction) (*Imports, error) {
// 	return createImportSetForFunctionComponents(f.Args, f.ReturnTypes)
// }

// func createImportSetForFunctionComponents(args []*TemplateFunctionArg, returnTypes []*Type) (*Imports, error) {
// 	imports := []*Import{}
// 	for i := range args {
// 		arg := args[i]
// 		imports = append(imports, arg.Type.Import)
// 	}
// 	for i := range returnTypes {
// 		returnType := returnTypes[i]
// 		imports = append(imports, returnType.Import)
// 	}
// 	return NewSetWithArray(imports)
// }
