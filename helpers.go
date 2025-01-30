package AstUtils

import (
	"errors"
	"go/ast"
	"go/token"
	"regexp"
	"strings"
	"unicode"
)

func ToSnakeCase(s string) string {
	match := regexp.MustCompilePOSIX("([a-z])([A-Z]|[0-9])|[0-9][A-Z]")
	return match.ReplaceAllString(s, "${1}_${2}")
}

func SetUnexported(name string) string {
	r := []rune(name)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func SetExported(name string) string {
	r := []rune(name)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func AddModifiedComment(file *ast.File, comment string) {
	file.Comments[0].List = append(file.Comments[0].List, &ast.Comment{
		Text: "// " + comment,
	})
}

func PreviouslyModified(file *ast.File, searchString string) bool {
	for _, comment := range file.Comments {
		for _, c := range comment.List {
			if strings.Contains(c.Text, searchString) {
				return true
			}
		}
	}
	return false
}

func AddMissingImports(file *ast.File, imports []string) {
	requiredImports := map[string]bool{}
	for _, imp := range imports {
		requiredImports[imp] = true
	}

	var importSpecs []*FoundNodes
	var completed bool
	SearchNodes(file, &importSpecs, []*ast.Node{}, func(n *ast.Node, parents []*ast.Node, completed *bool) bool {
		if _, ok := (*n).(*ast.ImportSpec); ok {
			return true
		}
		return false
	}, &completed)

	for i, spec := range importSpecs {
		for i2, parent := range spec.Parents {
			if _, ok := (*parent).(*ast.GenDecl); ok {
				for i3, _ := range file.Decls {
					if (*importSpecs[i].Parents[i2]).(*ast.GenDecl) == file.Decls[i3] {
						file.Decls[i3] = file.Decls[len(file.Decls)-1]
						file.Decls = file.Decls[:len(file.Decls)-1]
					}
				}
			}
		}
		if _, ok := requiredImports[strings.ReplaceAll((*spec.Node).(*ast.ImportSpec).Path.Value, "\"", "")]; ok {
			delete(requiredImports, strings.ReplaceAll((*spec.Node).(*ast.ImportSpec).Path.Value, "\"", ""))
		}
	}
	var specs []ast.Spec
	for importString, _ := range requiredImports {
		specs = append(specs, &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"" + importString + "\"",
			},
		})
	}
	var _ = &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"encoding/json\"",
				},
			},
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"errors\"",
				},
			},
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"fmt\"",
				},
			},
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"github.com/google/uuid\"",
				},
			},
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"time\"",
				},
			},
		},
	}
	file.Decls = append([]ast.Decl{&ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: specs,
	}}, file.Decls...)
}

func ReplaceImports(file *ast.File, imports []string) {
	for i, decl := range file.Decls {
		if GenDecl, ok := decl.(*ast.GenDecl); ok && GenDecl.Tok == token.IMPORT {
			file.Decls[i].(*ast.GenDecl).Specs = []ast.Spec{}
			for _, imp := range imports {
				file.Decls[i].(*ast.GenDecl).Specs = append(file.Decls[i].(*ast.GenDecl).Specs, &ast.ImportSpec{
					Path: &ast.BasicLit{
						Kind:  token.STRING,
						Value: "\"" + imp + "\"",
					},
				},
				)
			}
		}
	}
}

func IsBasicField(field *ast.Field) bool {
	basicTypes := []string{"string", "bool", "int8", "uint8", "byte", "int16", "uint16", "int32", "rune", "uint32", "int64", "uint64", "int", "uint", "uintptr", "float32", "float64", "complex64", "complex128"}
	switch t := field.Type.(type) {
	case *ast.Ident:
		if t == nil {
			return false
		}
		for _, basicType := range basicTypes {
			if t.Name == basicType {
				return true
			}
		}
	case *ast.StarExpr:
		if t == nil {
			return false
		}
		if ident, ok := t.X.(*ast.Ident); ok {
			if ident == nil {
				return false
			}
			for _, basicType := range basicTypes {
				if ident.Name == basicType {
					return true
				}
			}
		}
	default:
		return false
	}
	return false
}

func ExtractTagsByKey(tag *ast.BasicLit, valueMap ...map[string][]string) map[string][]string {
	var found map[string][]string
	if valueMap == nil || len(valueMap) == 0 || valueMap[0] == nil {
		found = make(map[string][]string)
	} else {
		found = valueMap[0]
	}
	if tag == nil {
		return found
	}
	tags := strings.Split(strings.ReplaceAll(tag.Value, "`", ""), " ")

	for _, s := range tags {
		v := strings.SplitN(strings.ReplaceAll(s, "\"", ""), ":", 2)
		if len(v) == 1 {
			continue
		}
		if _, ok := found[v[0]]; !ok {
			found[v[0]] = []string{v[1]}
		} else {
			found[v[0]] = append(found[v[0]], v[1])
		}
	}
	return found
}

func GetTagValue(tag string, tagKey string) string {
	tag = strings.ReplaceAll(tag, "`", "")
	tags := strings.Split(tag, " ")
	for _, s := range tags {
		s = strings.ReplaceAll(s, "\"", "")
		ss := strings.Split(s, ":")
		if len(ss) > 1 && ss[0] == tagKey {
			return ss[1]
		}
	}
	return ""
}

func GetJsonTagValue(tag string) string {
	return GetTagValue(tag, "json")
}

func GetJsonTagName(tag *ast.BasicLit) (string, error) {
	keys := ExtractTagsByKey(tag)
	if v, ok := keys["json"]; ok {
		return strings.ReplaceAll(v[0], ",omitempty", ""), nil
	}
	return "", errors.New("json tag not found in tag")
}

type TagCombiner interface {
	Combine(values []string) (string, error)
}

// CombineTags Combines two tags. Uses the first seen tag, as combined tag,
// unless a TagCombiner for the tag key is present.
func CombineTags(tag1, tag2 *ast.BasicLit, combiners map[string]TagCombiner) (*ast.BasicLit, error) {
	var tagString string
	var err error

	tags := make(map[string][]string)
	tags = ExtractTagsByKey(tag1, tags)
	tags = ExtractTagsByKey(tag2, tags)

	for key, values := range tags {
		var combinedTags string
		if combiner, ok := combiners[key]; ok {
			combinedTags, err = combiner.Combine(values)
			if err != nil {
				return nil, err
			}
		} else {
			combinedTags = values[0]
		}
		tagString += key + ":\"" + combinedTags + "\" "
	}
	if tagString != "" {
		return &ast.BasicLit{
			ValuePos: 0,
			Kind:     token.STRING,
			Value:    "`" + tagString + "`",
		}, nil
	}
	return &ast.BasicLit{}, nil
}

func RemoveTag(key string, lit *ast.BasicLit) *ast.BasicLit {
	tags := ExtractTagsByKey(lit, nil)
	delete(tags, key)
	var TagString string
	for k, v := range tags {
		TagString += k + ":\"" + v[0] + "\" "
	}
	if TagString != "" {
		return &ast.BasicLit{
			ValuePos: 0,
			Kind:     token.STRING,
			Value:    "`" + TagString + "`",
		}
	}
	return &ast.BasicLit{}
}
