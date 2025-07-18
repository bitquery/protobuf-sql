package generators

import (
	"bytes"
	"fmt"
	"github.com/bitquery/protobuf-sql/compiler/protogen"
	"github.com/bitquery/protobuf-sql/reflect/protoreflect"
	"github.com/gertd/go-pluralize"
	"github.com/stoewer/go-strcase"
	"html/template"
	"strings"
)

const sqlField = "%s\t%s"
const arrayTypeTemplate = arrayType + "(%s)"
const columnWidth = 45

type fileGenerator struct {
	gen               *protogen.Plugin
	message           *protogen.Message
	gf                *protogen.GeneratedFile
	fieldsToIgnoreMap map[string]struct{}
	fieldsSl          []field
	template          *template.Template

	logs *bytes.Buffer
}

type field struct {
	Name string
	Type string
}

type templateData struct {
	Database     string
	Table        string
	FieldsWTypes string
	Fields       string
	Suffix       string
	IgnoreFields string
}

func GenerateFileForMessage(gen *protogen.Plugin, m *protogen.Message) {
	fg := &fileGenerator{
		gen:     gen,
		message: m,

		logs: bytes.NewBuffer(nil),
	}

	if fg.FilterMessage() {
		return
	}

	fg.setupTemplate()
	fg.parseIgnoreFields()
	fg.generateFile()
	//log.Fatal(fg.logs.String())
}

func (g *fileGenerator) FilterMessage() bool {
	if strings.HasSuffix(g.message.GoIdent.GoName, g.gen.MessageSuffix) {
		return false
	}
	return true
}

func (g *fileGenerator) generateFile() {
	fileName := fileName(g.message.GoIdent.GoName, g.gen.MessageSuffix)
	g.gf = g.gen.NewGeneratedFile(fileName, "")

	g.generateFields(g.message, "", 0)

	var fieldsWTypesBuilder strings.Builder
	for i, f := range g.fieldsSl {
		paddedName := fmt.Sprintf("%-*s ", columnWidth, f.Name)
		fRow := fmt.Sprintf(sqlField, paddedName, f.Type)
		if i != len(g.fieldsSl)-1 {
			fRow += ",\n\t"
		}
		fieldsWTypesBuilder.WriteString(fRow)
	}

	var fieldsBuilder strings.Builder
	for i, f := range g.fieldsSl {
		fRow := f.Name
		if i != len(g.fieldsSl)-1 {
			fRow += ",\n\t"
		}
		fieldsBuilder.WriteString(fRow)
	}

	td := templateData{
		Database:     g.gen.DBName,
		Table:        tableName(g.message.GoIdent.GoName, g.gen.MessageSuffix),
		FieldsWTypes: fieldsWTypesBuilder.String(),
		Fields:       fieldsBuilder.String(),
		Suffix:       g.gen.TemplateSuffix,
	}

	err := g.template.Execute(g.gf, td)
	if err != nil {
		fmt.Println("error while executing template: ", err)
		return
	}
}

func (g *fileGenerator) generateFields(message *protogen.Message, prefix string, wrapCount int) {
	currLevelWrapCount := wrapCount
	for _, f := range message.Fields {
		fieldName := prefix + f.GoName

		if f.Desc.Kind() == protoreflect.MessageKind {
			if f.Desc.IsList() {
				g.generateFields(f.Message, fieldName+"_", currLevelWrapCount+1)
				continue
			}
			g.generateFields(f.Message, fieldName+"_", currLevelWrapCount)
			continue
		}

		wc := currLevelWrapCount
		if f.Desc.IsList() {
			wc++
		}

		fType := protoToClickhouse[f.Desc.Kind().String()]
		for i := 0; i < wc; i++ {
			fType = fmt.Sprintf(arrayTypeTemplate, fType)
		}

		if _, ok := g.fieldsToIgnoreMap[fieldName]; !ok {
			// remove trailing '_' if exists
			if len(fieldName) > 0 && fieldName[len(fieldName)-1] == '_' {
				fieldName = fieldName[:len(fieldName)-1]
			}
			g.fieldsSl = append(g.fieldsSl, field{
				fieldName, fType,
			})
		}
	}
}

func (g *fileGenerator) setupTemplate() {
	if g.template == nil {
		g.template = template.Must(template.ParseFiles(g.gen.TemplatePath))
	}
}

func (g *fileGenerator) parseIgnoreFields() {
	if g.gen.IgnoreFields != "" {
		g.logs.WriteString(fmt.Sprintf("Ignoring fields: %s\n", g.gen.IgnoreFields))

		fieldsToIgnore := strings.Split(g.gen.IgnoreFields, ";")
		g.fieldsToIgnoreMap = make(map[string]struct{}, len(fieldsToIgnore))
		for _, f := range fieldsToIgnore {
			g.logs.WriteString(fmt.Sprintf("Ignoring field: %s\n", f))
			g.fieldsToIgnoreMap[f] = struct{}{}
		}
	}
}

// genFileName removes suffix "Message" from message name, convert to snake case and add ".sql" suffix
func fileName(messageName, suffix string) string {
	return tableName(messageName, suffix) + ".sql"
}

func tableName(messageName, suffix string) string {
	singular := strcase.SnakeCase(strings.TrimSuffix(messageName, suffix))
	return pluralize.NewClient().Plural(singular)
}
