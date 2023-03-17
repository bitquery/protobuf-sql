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

const plainField = "%s %s"
const arrayField = "%s array(%s)"

var templ *template.Template

type fileGenerator struct {
	gen           *protogen.Plugin
	message       *protogen.Message
	gf            *protogen.GeneratedFile
	fieldsBuilder strings.Builder
	template      *template.Template

	logs *bytes.Buffer
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

	g.generateFields(g.message, "")

	td := templateData{
		Database: g.gen.DBName,
		Table:    tableName(g.message.GoIdent.GoName, g.gen.MessageSuffix),
		Fields:   g.fieldsBuilder.String(),
	}

	err := g.template.Execute(g.gf, td)
	if err != nil {
		fmt.Println("error while executing template: ", err)
		return
	}
}

func (g *fileGenerator) generateFields(message *protogen.Message, prefix string) {
	for _, f := range message.Fields {
		if f.Desc.Kind() == protoreflect.MessageKind {
			// if f.Desc.IsList() {
			// 	g.gen.Error(fmt.Errorf("list of messages is not supported"))
			// }
			g.generateFields(f.Message, prefix+f.GoName+"_")
			// fmt.Fprint(fg.logs, f.Message)
			// fmt.Fprint(fg.logs, f.Extendee)
			// fmt.Fprint(fg.logs, f.Parent)
			// fmt.Fprint(fg.logs, f.Desc.Message())
			// fmt.Println(fg.logs.String())

			continue
		}

		if g.fieldsBuilder.Len() > 0 {
			g.fieldsBuilder.WriteString(",\n\t")
		}

		if f.Desc.IsList() {
			g.fieldsBuilder.WriteString(
				fmt.Sprintf(arrayField, prefix+f.GoName, f.Desc.Kind().String()),
			)
			continue
		}

		g.fieldsBuilder.WriteString(
			fmt.Sprintf(plainField, prefix+f.GoName, f.Desc.Kind().String()),
		)

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
