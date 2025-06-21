package abiutil

import (
	"bytes"
	_ "embed"

	"fmt"
	"log"
	"text/template"

	"github.com/NethermindEth/starknet.go/utils"
)

//go:embed templates/decoders.tmpl
var decoderTmpl string

func GenerateDecoderFunctions(events []CairoAbiEntry, pkg string) string {
	ctx := TemplateContext{Package: pkg}
	for _, ev := range events {
		if ev.Kind != "struct" {
			continue
		}
		ctx.Events = append(ctx.Events, EventData{
			Name:       ev.Name,
			StructName: simplifyTypeName(ev.Name),
			Kind:       ev.Kind,
			Selector:   utils.GetSelectorFromNameFelt(simplifyTypeName(ev.Name)).String(),
			KeyFields:  buildFields(ev.Members, "key"),
			DataFields: buildFields(ev.Members, "data"),
		})
	}

	tmpl := template.Must(template.
		New("decoders.tmpl").
		Funcs(template.FuncMap{
			"countFields": func(fields []Field) int {
				total := 0
				for _, f := range fields {
					switch f.GoType {
					case "types.U256", "types.U128":
						total += 2
					case "types.I257":
						total += 3
					default:
						total++
					}
				}
				return total
			},
			"assign": func(sliceName string, f Field) string {
				switch f.GoType {
				case "types.U256", "types.U128":
					return fmt.Sprintf("\t\t\t%s: %s{\n\t\t\t\tLow: %s[%d].String(),\n\t\t\t\tHigh: %s[%d].String(),\n\t\t\t},",
						f.Name, f.GoType, sliceName, f.IndexOff, sliceName, f.IndexOff+1)
				case "types.I257":
					return fmt.Sprintf("\t\t\t%s: types.I257{\n\t\t\t\tAbs: types.U256{\n\t\t\t\t\tLow: %s[%d].String(),\n\t\t\t\t\tHigh: %s[%d].String(),\n\t\t\t\t},\n\t\t\t\tIsNegative: %s[%d].Cmp(&felt.Zero) != 0,\n\t\t\t},",
						f.Name, sliceName, f.IndexOff, sliceName, f.IndexOff+1, sliceName, f.IndexOff+2)
				case "uint64":
					return fmt.Sprintf("\t\t\t%s: %s[%d].Uint64(),", f.Name, sliceName, f.IndexOff)
				case "bool":
					return fmt.Sprintf("\t\t\t%s: %s[%d].Cmp(&felt.Zero) != 0,", f.Name, sliceName, f.IndexOff)
				case "types.TextFelt":
					return fmt.Sprintf("\t\t\t%s: types.NewTextFelt(%s[%d]),", f.Name, sliceName, f.IndexOff)
				case "types.Address":
					return fmt.Sprintf("\t\t\t%s: types.NewAddress(%s[%d]),", f.Name, sliceName, f.IndexOff)
				default:
					return fmt.Sprintf("\t\t\t%s: %s[%d].String(),", f.Name, sliceName, f.IndexOff)
				}
			},
		}).
		Parse(decoderTmpl))

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, ctx); err != nil {
		log.Fatalf("template execution failed: %v", err)
	}
	return buf.String()
}
