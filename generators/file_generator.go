package generators

import (
	"bytes"
	"fmt"
	"github.com/bitquery/protobuf-sql/compiler/protogen"
	"github.com/bitquery/protobuf-sql/reflect/protoreflect"
	"github.com/stoewer/go-strcase"
	"html/template"
	"strings"
)

const sqlField = "%s\t%s"
const arrayType = "array(%s)"
const columnWidth = 45

type fileGenerator struct {
	gen      *protogen.Plugin
	message  *protogen.Message
	gf       *protogen.GeneratedFile
	fieldsSl []field
	template *template.Template

	logs *bytes.Buffer
}

type field struct {
	Name string
	Type string
}

type templateData struct {
	Database string
	Table    string
	Fields   string
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

	fg.generateFile()
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

	var fieldsBuilder strings.Builder
	for i, f := range g.fieldsSl {
		paddedName := fmt.Sprintf("%-*s ", columnWidth, f.Name)
		fRow := fmt.Sprintf(sqlField, paddedName, f.Type)
		if i != len(g.fieldsSl)-1 {
			fRow += ",\n\t"
		}
		fieldsBuilder.WriteString(fRow)
	}

	td := templateData{
		Database: g.gen.DBName,
		Table:    tableName(g.message.GoIdent.GoName, g.gen.MessageSuffix),
		Fields:   fieldsBuilder.String(),
	}

	err := g.template.Execute(g.gf, td)
	if err != nil {
		fmt.Println("error while executing template: ", err)
		return
	}
}

func (g *fileGenerator) generateFields(message *protogen.Message, prefix string, wrapCount int) {
	currLevelWrapCount := wrapCount
	nextLevelWrapCount := currLevelWrapCount
	for _, f := range message.Fields {
		fieldName := prefix + f.GoName

		if f.Desc.IsList() {
			nextLevelWrapCount = currLevelWrapCount + 1
		}

		if f.Desc.Kind() == protoreflect.MessageKind {
			g.generateFields(f.Message, fieldName+"_", nextLevelWrapCount)
			continue
		}

		fType := protoToClickhouse[f.Desc.Kind().String()]
		for i := 0; i < nextLevelWrapCount; i++ {
			fType = fmt.Sprintf(arrayType, fType)
		}

		g.fieldsSl = append(g.fieldsSl, field{
			fieldName, fType,
		})
	}
}

func (g *fileGenerator) setupTemplate() {
	if g.template == nil {
		g.template = template.Must(template.ParseFiles(g.gen.TemplatePath))
	}
}

// genFileName removes suffix "Message" from message name, convert to snake case and add ".sql" suffix
func fileName(messageName, suffix string) string {
	return tableName(messageName, suffix) + ".sql"
}

func tableName(messageName, suffix string) string {
	return strcase.SnakeCase(strings.TrimSuffix(messageName, suffix)) + "s"
}
