package frontend

import (
	compiler "VAsm/backend/analizador/parser"
	"VAsm/backend/visitor"
	"VAsm/frontend/cst"
	"VAsm/frontend/errors"
	"VAsm/frontend/symbols"
	"bytes"
	"fmt"
	"image/color"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/antlr4-go/antlr/v4"
)

type Editor struct {
	Content        *widget.Entry
	Window         fyne.Window
	ConsoleLabel   *widget.Label
	currentFile    fyne.URI
	errorTable     *errors.ErrorTable
	errorData      []errors.ErrorTable
	errorTableUI   *widget.Table
	Visitor        *visitor.EvalVisitor
	simbolTable    *symbols.TablaSimbolos
	symbolTable    *widget.Table
	ImageContainer *fyne.Container
	ImageScroll    *container.Scroll
}

func BuildMainWindow(win fyne.Window) *Editor {
	editor := &Editor{
		Content:     widget.NewMultiLineEntry(),
		Window:      win,
		simbolTable: symbols.NewTablaSimbolos(),
		errorTable:  errors.NewErrorTable(),
	}

	editor.Visitor = visitor.NewEvalVisitor(editor.simbolTable)

	editor.Content.SetPlaceHolder("Escribe tu código aquí...")
	editor.Content.TextStyle = fyne.TextStyle{Monospace: true}

	editor.ConsoleLabel = widget.NewLabel("Resultado...")
	editor.ConsoleLabel.Wrapping = fyne.TextWrapWord
	editor.ConsoleLabel.TextStyle = fyne.TextStyle{Monospace: true}
	consoleScroll := container.NewVScroll(editor.ConsoleLabel)
	consoleContainer := container.NewStack(
		canvas.NewRectangle(color.NRGBA{R: 30, G: 30, B: 30, A: 255}),
		consoleScroll,
	)
	// tabla VISUAL de errores, la logica esta en la carpeta errors
	editor.errorTableUI = widget.NewTable(
		func() (int, int) {
			return len(editor.errorTable.Errors) + 1, 5
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			if id.Row == 0 {
				headers := []string{"No", "Descripción", "Línea", "Columna", "Tipo"}
				label.SetText(headers[id.Col])
				label.TextStyle.Bold = true
				label.Alignment = fyne.TextAlignCenter
			} else {
				sortedErrors := editor.errorTable.GetSortedErrors()
				if id.Row-1 < len(sortedErrors) {
					err := sortedErrors[id.Row-1]
					switch id.Col {
					case 0:
						label.SetText(fmt.Sprintf("%d", id.Row))
					case 1:
						label.SetText(err.Msg)
						label.Wrapping = fyne.TextWrapWord
					case 2:
						label.SetText(fmt.Sprintf("%d", err.Line))
					case 3:
						label.SetText(fmt.Sprintf("%d", err.Column))
					case 4:
						label.SetText(err.Type)
					}
				}
				label.Alignment = fyne.TextAlignLeading
			}
		},
	)
	editor.errorTableUI.SetColumnWidth(0, 60)
	editor.errorTableUI.SetColumnWidth(1, 800)
	editor.errorTableUI.SetColumnWidth(2, 60)
	editor.errorTableUI.SetColumnWidth(3, 70)
	editor.errorTableUI.SetColumnWidth(4, 200)

	go func() {
		time.Sleep(300 * time.Millisecond)
		editor.adjustDescriptionColumn()
	}()

	editor.symbolTable = widget.NewTable(
		func() (int, int) {
			return len(editor.simbolTable.GenerarReporte()) + 1, 7
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			if id.Row == 0 {
				headers := []string{"ID", "Tipo simbolo", "Tipo dato", "Ambito", "Valor", "Linea", "Columna"}
				label.SetText(headers[id.Col])
				label.TextStyle.Bold = true
				label.Alignment = fyne.TextAlignCenter
			} else {
				reporte := editor.simbolTable.GenerarReporte()
				if id.Row-1 < len(reporte) {
					simbolo := reporte[id.Row-1]
					switch id.Col {
					case 0:
						label.SetText(simbolo.ID)
					case 1:
						label.SetText(simbolo.TipoSimbolo)
					case 2:
						label.SetText(simbolo.TipoDato)
					case 3:
						label.SetText(simbolo.Ambito)
					case 4:
						label.SetText(fmt.Sprintf("%v", simbolo.Valor))
					case 5:
						label.SetText(fmt.Sprintf("%d", simbolo.Linea))
					case 6:
						label.SetText(fmt.Sprintf("%d", simbolo.Columna))
					}
					label.Alignment = fyne.TextAlignCenter
				}
			}
		},
	)

	editor.symbolTable.SetColumnWidth(0, 100)
	editor.symbolTable.SetColumnWidth(1, 150)
	editor.symbolTable.SetColumnWidth(2, 100)
	editor.symbolTable.SetColumnWidth(3, 300)
	editor.symbolTable.SetColumnWidth(4, 400)
	editor.symbolTable.SetColumnWidth(5, 60)
	editor.symbolTable.SetColumnWidth(6, 60)

	go func() {
		time.Sleep(300 * time.Millisecond)
		editor.adjustAmbitColumn()
	}()
	// Inicialización previa de ImageScroll e ImageContainer
	rect := canvas.NewRectangle(color.White)
	scroll := container.NewScroll(rect)
	stack := container.NewStack(
		canvas.NewRectangle(color.White),
		scroll,
	)

	editor.ImageScroll = scroll
	editor.ImageContainer = stack
	tabs := container.NewAppTabs(
		container.NewTabItem("Errores", editor.errorTableUI),
		container.NewTabItem("Símbolos", editor.symbolTable),
	)
	// Crear los contenedores principales
	bottomSplit := container.NewVSplit(consoleContainer, tabs)
	middleSplit := container.NewVSplit(editor.Content, bottomSplit)
	mainSplit := container.NewVSplit(middleSplit, editor.ImageContainer)
	mainSplit.SetOffset(0.3)

	// Barra de botones
	buttonBar := container.NewHBox(
		widget.NewButton("Guardar", editor.saveFile),
		widget.NewButton("Abrir", editor.openFile),
		widget.NewButton("Crear archivo", editor.createNewFile),
		// Botón Ejecutar con análisis y actualización de imagen SVG
		widget.NewButton("Ejecutar", func() {
			editor.errorTable = errors.NewErrorTable()
			editor.errorTableUI.Refresh()
			editor.simbolTable = symbols.NewTablaSimbolos()
			editor.Visitor = visitor.NewEvalVisitor(editor.simbolTable)

			codigo := editor.Content.Text
			if codigo == "" {
				editor.ConsoleLabel.SetText("Error: No hay código para analizar")
				return
			}

			editor.errorData = []errors.ErrorTable{}
			editor.ConsoleLabel.SetText("Analizando código...")
			editor.errorTableUI.Refresh()

			defer func() {
				if r := recover(); r != nil {
					errMsg := fmt.Sprintf("%v", r)
					editor.ConsoleLabel.SetText("Error: " + errMsg)

					line, column := 1, 1
					if err, ok := r.(*visitor.EvaluationError); ok {
						line, column = err.Line, err.Column
					}

					editor.errorTable.NewRuntimeError(line, column, errMsg)
					editor.errorTableUI.Refresh()
				}
			}()

			// Configuración del análisis léxico y sintáctico
			input := antlr.NewInputStream(codigo)

			lexer := compiler.NewgramaticaLexer(input)
			lexerErrorListener := errors.NewLexicalErrorListener()
			lexer.RemoveErrorListeners()
			lexer.AddErrorListener(lexerErrorListener)

			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
			parser := compiler.NewgramaticaParser(stream)

			parser.SetErrorHandler(errors.NewCustomErrorStrategy())

			parserErrorListener := errors.NewSyntaxErrorListener(editor.errorTable)
			parser.RemoveErrorListeners()
			parser.AddErrorListener(parserErrorListener)

			tree := parser.Init()

			symbolVisitor := visitor.NewSymbolVisitor(editor.simbolTable)
			if editor.simbolTable == nil {
				editor.simbolTable = symbols.NewTablaSimbolos()
			}

			symbolVisitor.Visit(tree)
			if len(lexerErrorListener.ErrorTable.Errors) > 0 {
				for _, err := range lexerErrorListener.ErrorTable.Errors {
					editor.errorTable.AddError(err.Line, err.Column, err.Msg, err.Type)
				}
			}

			editor.symbolTable.Refresh()
			editor.UpdatSymb()

			// Mostrar errores si hay
			if len(editor.errorTable.Errors) > 0 {
				editor.ConsoleLabel.SetText("Se encontraron errores durante el análisis")
				editor.UpdateErrors()
				return
			}

			result := editor.Visitor.Visit(tree)
			output := editor.Visitor.Console.String()
			output = strings.TrimSpace(output)

			tmpSVG, err := os.CreateTemp("", "cst_*.svg")
			if err != nil {
				log.Println("Error creando archivo temporal SVG:", err)
				return
			}
			defer os.Remove(tmpSVG.Name())
			svgRaw := cst.CstReport(codigo)
			_, err = tmpSVG.WriteString(svgRaw)
			if err != nil {
				log.Println("Error escribiendo SVG en archivo temporal:", err)
				return
			}
			tmpSVG.Close()

			// 2. Crear archivo temporal PNG
			tmpPNG, err := os.CreateTemp("", "cst_*.png")
			if err != nil {
				log.Println("Error creando archivo temporal PNG:", err)
				return
			}
			defer os.Remove(tmpPNG.Name())
			tmpPNG.Close()
			// descomentar aqui para calificacion

/*			// 3. Convertir SVG a PNG usando rsvg-convert (debes tenerlo instalado)
			cmd := exec.Command("rsvg-convert", "-o", tmpPNG.Name(), tmpSVG.Name())
			err = cmd.Run()
			if err != nil {
				log.Println("Error convirtiendo SVG a PNG:", err)
				return
			}

			// 4. Cargar imagen PNG y mostrarla en la interfaz
			pngImage := canvas.NewImageFromFile(tmpPNG.Name())
			pngImage.FillMode = canvas.ImageFillContain
			pngImage.SetMinSize(fyne.NewSize(400, 400))

			editor.ImageScroll.Content = pngImage
			editor.ImageScroll.Refresh()

			// cerrar comentario para calificacion arriba linea 295*/
			var texto string
			if resultados, ok := result.([]interface{}); ok {
				for _, v := range resultados {
					texto += fmt.Sprintf("%v ", v)
				}
				texto = strings.TrimSpace(texto)
				editor.ConsoleLabel.SetText(fmt.Sprintf("%s", output))
				err = os.MkdirAll("assembler", os.ModePerm)
				if err != nil {
					log.Println("Error creando carpeta 'assembler':", err)
					return
				}

				archivo := filepath.Join("assembler", "program.s")
				file, err := os.Create(archivo)
				if err != nil {
					log.Println("Error creando archivo program.s:", err)
					return
				}
				defer file.Close()

				_, err = file.WriteString(editor.ConsoleLabel.Text)
				if err != nil {
					log.Println("Error escribiendo en archivo program.s:", err)
					return
				}
			} else {
				editor.ConsoleLabel.SetText(fmt.Sprintf("%s", output))
			}
		}),
	)

	// Contenedor inferior y medio
	bottomSplit = container.NewVSplit(consoleContainer, tabs)
	middleSplit = container.NewVSplit(editor.Content, bottomSplit)

	// Contenedor principal con el split de la imagen
	mainSplit = container.NewVSplit(middleSplit, editor.ImageContainer)
	mainSplit.SetOffset(0.3)

	// Contenedor principal con barra de botones arriba
	mainContent := container.NewBorder(
		buttonBar, nil, nil, nil,
		mainSplit,
	)

	// Configuración ventana
	win.SetFullScreen(true)
	win.CenterOnScreen()
	win.Resize(fyne.NewSize(1000, 800))
	win.SetContent(mainContent)

	return editor
}

func ShowCstSVGAsPNG(w fyne.Window, svgContent string) {
	tmpSVG, err := os.CreateTemp("", "cst_*.svg")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpSVG.Name())
	defer tmpSVG.Close()

	_, err = tmpSVG.WriteString(svgContent)
	if err != nil {
		panic(err)
	}

	tmpPNG, err := os.CreateTemp("", "cst_*.png")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpPNG.Name())
	defer tmpPNG.Close()

	// Convierte SVG a PNG (requiere rsvg-convert)
	cmd := exec.Command("rsvg-convert", "-o", tmpPNG.Name(), tmpSVG.Name())
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	// Mostrar PNG en Fyne
	img := canvas.NewImageFromFile(tmpPNG.Name())
	img.FillMode = canvas.ImageFillContain
	w.SetContent(img)
}

func (e *Editor) createNewFile() {
	if e.Content.Text != "" {
		dialog.ShowConfirm("Nuevo archivo",
			"¿Deseas guardar los cambios antes de crear un nuevo archivo?",
			func(save bool) {
				if save {
					e.saveFile()
				}
				e.clearEditor()
			}, e.Window)
	} else {
		e.clearEditor()
	}
}

func (e *Editor) saveFile() {
	if e.currentFile == nil {
		e.saveFileAs()
		return
	}

	write, err := storage.Writer(e.currentFile)
	if err != nil {
		dialog.ShowError(err, e.Window)
		return
	}
	defer write.Close()

	_, err = io.WriteString(write, e.Content.Text)
	if err != nil {
		dialog.ShowError(err, e.Window)
		return
	}

	e.ConsoleLabel.SetText("Archivo guardado: " + e.currentFile.Name())
}

func (e *Editor) saveFileAs() {
	dialog.ShowFileSave(func(uri fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, e.Window)
			return
		}
		if uri == nil {
			return
		}
		defer uri.Close()

		if filepath.Ext(uri.URI().Name()) != ".vch" {
			dialog.ShowInformation("Extensión incorrecta",
				"Los archivos deben tener extensión .vch", e.Window)
			return
		}

		_, err = io.WriteString(uri, e.Content.Text)
		if err != nil {
			dialog.ShowError(err, e.Window)
			return
		}

		e.currentFile = uri.URI()
		e.ConsoleLabel.SetText("Archivo guardado como: " + e.currentFile.Name())
	}, e.Window)
}

func (e *Editor) openFile() {
	dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, e.Window)
			return
		}
		if uri == nil {
			return
		}
		defer uri.Close()

		if filepath.Ext(uri.URI().Name()) != ".vch" {
			dialog.ShowInformation("Extensión incorrecta",
				"Solo se pueden abrir archivos .vch", e.Window)
			return
		}

		var buf bytes.Buffer
		_, err = io.Copy(&buf, uri)
		if err != nil {
			dialog.ShowError(err, e.Window)
			return
		}

		e.Content.SetText(buf.String())
		e.currentFile = uri.URI()
		e.ConsoleLabel.SetText("Archivo abierto: " + e.currentFile.Name())
	}, e.Window)
}

func (e *Editor) clearEditor() {
	e.Content.SetText("")
	e.Content.Refresh()
	e.ConsoleLabel.SetText("Resultado...")
	e.ConsoleLabel.Refresh()
}

func (e *Editor) adjustDescriptionColumn() {
	maxWidth := float32(800)
	for _, err := range e.errorTable.Errors {
		textLen := float32(len(err.Msg)) * 7
		if textLen > maxWidth && textLen < 800 {
			maxWidth = textLen
		}
	}
	if e.errorTableUI != nil {
		e.errorTableUI.SetColumnWidth(1, maxWidth)
		e.errorTableUI.Refresh()
	}
}

func (e *Editor) UpdateErrors() {
	e.errorTableUI.Refresh()
	go e.adjustDescriptionColumn()
}
func (e *Editor) adjustAmbitColumn() {
	maxWidth := float32(300)
	for _, entorno := range e.simbolTable.EntornosFunciones {
		for _, simbolo := range entorno.Simbolos {
			textLen := float32(len(simbolo.Ambito)) * 7
			if textLen > maxWidth && textLen < 300 {
				maxWidth = textLen
			}
		}
	}
	if e.symbolTable != nil {
		e.symbolTable.SetColumnWidth(3, maxWidth)
		e.symbolTable.Refresh()
	}
}

func (e *Editor) UpdatSymb() {
	e.symbolTable.Refresh()
	go e.adjustAmbitColumn()
}
