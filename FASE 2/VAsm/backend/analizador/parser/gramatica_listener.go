// Code generated from backend/analizador/parser/gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // gramatica
import "github.com/antlr4-go/antlr/v4"

// gramaticaListener is a complete listener for a parse tree produced by gramaticaParser.
type gramaticaListener interface {
	antlr.ParseTreeListener

	// EnterInit is called when entering the init production.
	EnterInit(c *InitContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterPrintln is called when entering the println production.
	EnterPrintln(c *PrintlnContext)

	// EnterDeclaracionMultiple is called when entering the declaracionMultiple production.
	EnterDeclaracionMultiple(c *DeclaracionMultipleContext)

	// EnterDeclaracionMultipleSimple is called when entering the declaracionMultipleSimple production.
	EnterDeclaracionMultipleSimple(c *DeclaracionMultipleSimpleContext)

	// EnterDeclaracionMultipleSinTipo is called when entering the declaracionMultipleSinTipo production.
	EnterDeclaracionMultipleSinTipo(c *DeclaracionMultipleSinTipoContext)

	// EnterAsignacionMultiple is called when entering the asignacionMultiple production.
	EnterAsignacionMultiple(c *AsignacionMultipleContext)

	// EnterBloqueFuncion is called when entering the bloqueFuncion production.
	EnterBloqueFuncion(c *BloqueFuncionContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterLlamadaFuncionesSinParametro is called when entering the llamadaFuncionesSinParametro production.
	EnterLlamadaFuncionesSinParametro(c *LlamadaFuncionesSinParametroContext)

	// EnterLlamadaFuncionesConParametro is called when entering the llamadaFuncionesConParametro production.
	EnterLlamadaFuncionesConParametro(c *LlamadaFuncionesConParametroContext)

	// EnterFnSinParametro is called when entering the fnSinParametro production.
	EnterFnSinParametro(c *FnSinParametroContext)

	// EnterFnConParametro is called when entering the fnConParametro production.
	EnterFnConParametro(c *FnConParametroContext)

	// EnterSliceDef is called when entering the sliceDef production.
	EnterSliceDef(c *SliceDefContext)

	// EnterSliceLiteral is called when entering the sliceLiteral production.
	EnterSliceLiteral(c *SliceLiteralContext)

	// EnterAccesoElementoSlice is called when entering the accesoElementoSlice production.
	EnterAccesoElementoSlice(c *AccesoElementoSliceContext)

	// EnterModificacionElementoSlice is called when entering the modificacionElementoSlice production.
	EnterModificacionElementoSlice(c *ModificacionElementoSliceContext)

	// EnterFnIndexOf is called when entering the fnIndexOf production.
	EnterFnIndexOf(c *FnIndexOfContext)

	// EnterFnJoin is called when entering the fnJoin production.
	EnterFnJoin(c *FnJoinContext)

	// EnterFnLen is called when entering the fnLen production.
	EnterFnLen(c *FnLenContext)

	// EnterFnAppend is called when entering the fnAppend production.
	EnterFnAppend(c *FnAppendContext)

	// EnterDeclaracionSliceVacio is called when entering the declaracionSliceVacio production.
	EnterDeclaracionSliceVacio(c *DeclaracionSliceVacioContext)

	// EnterTipoRetorno is called when entering the tipoRetorno production.
	EnterTipoRetorno(c *TipoRetornoContext)

	// EnterRetorno is called when entering the retorno production.
	EnterRetorno(c *RetornoContext)

	// EnterFnTypeOf is called when entering the fnTypeOf production.
	EnterFnTypeOf(c *FnTypeOfContext)

	// EnterFnAtoi is called when entering the fnAtoi production.
	EnterFnAtoi(c *FnAtoiContext)

	// EnterFnParseToFloat is called when entering the fnParseToFloat production.
	EnterFnParseToFloat(c *FnParseToFloatContext)

	// EnterAsigIncre is called when entering the asigIncre production.
	EnterAsigIncre(c *AsigIncreContext)

	// EnterAsigDecre is called when entering the asigDecre production.
	EnterAsigDecre(c *AsigDecreContext)

	// EnterAsignacionNormal is called when entering the AsignacionNormal production.
	EnterAsignacionNormal(c *AsignacionNormalContext)

	// EnterAsignacionIncremento is called when entering the AsignacionIncremento production.
	EnterAsignacionIncremento(c *AsignacionIncrementoContext)

	// EnterAsignacionDecremento is called when entering the AsignacionDecremento production.
	EnterAsignacionDecremento(c *AsignacionDecrementoContext)

	// EnterListaIDS is called when entering the listaIDS production.
	EnterListaIDS(c *ListaIDSContext)

	// EnterListaPar is called when entering the listaPar production.
	EnterListaPar(c *ListaParContext)

	// EnterParametro is called when entering the parametro production.
	EnterParametro(c *ParametroContext)

	// EnterListaExpr is called when entering the listaExpr production.
	EnterListaExpr(c *ListaExprContext)

	// EnterListaExprList is called when entering the listaExprList production.
	EnterListaExprList(c *ListaExprListContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterSif is called when entering the sif production.
	EnterSif(c *SifContext)

	// EnterElseifPart is called when entering the elseifPart production.
	EnterElseifPart(c *ElseifPartContext)

	// EnterForCondicional is called when entering the ForCondicional production.
	EnterForCondicional(c *ForCondicionalContext)

	// EnterForClasico is called when entering the ForClasico production.
	EnterForClasico(c *ForClasicoContext)

	// EnterForRange is called when entering the ForRange production.
	EnterForRange(c *ForRangeContext)

	// EnterSSwitch is called when entering the sSwitch production.
	EnterSSwitch(c *SSwitchContext)

	// EnterBreak_ is called when entering the break_ production.
	EnterBreak_(c *Break_Context)

	// EnterContinue_ is called when entering the continue_ production.
	EnterContinue_(c *Continue_Context)

	// EnterCaseBlock is called when entering the caseBlock production.
	EnterCaseBlock(c *CaseBlockContext)

	// EnterDefaultBlock is called when entering the defaultBlock production.
	EnterDefaultBlock(c *DefaultBlockContext)

	// EnterTipos is called when entering the tipos production.
	EnterTipos(c *TiposContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitPrintln is called when exiting the println production.
	ExitPrintln(c *PrintlnContext)

	// ExitDeclaracionMultiple is called when exiting the declaracionMultiple production.
	ExitDeclaracionMultiple(c *DeclaracionMultipleContext)

	// ExitDeclaracionMultipleSimple is called when exiting the declaracionMultipleSimple production.
	ExitDeclaracionMultipleSimple(c *DeclaracionMultipleSimpleContext)

	// ExitDeclaracionMultipleSinTipo is called when exiting the declaracionMultipleSinTipo production.
	ExitDeclaracionMultipleSinTipo(c *DeclaracionMultipleSinTipoContext)

	// ExitAsignacionMultiple is called when exiting the asignacionMultiple production.
	ExitAsignacionMultiple(c *AsignacionMultipleContext)

	// ExitBloqueFuncion is called when exiting the bloqueFuncion production.
	ExitBloqueFuncion(c *BloqueFuncionContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitLlamadaFuncionesSinParametro is called when exiting the llamadaFuncionesSinParametro production.
	ExitLlamadaFuncionesSinParametro(c *LlamadaFuncionesSinParametroContext)

	// ExitLlamadaFuncionesConParametro is called when exiting the llamadaFuncionesConParametro production.
	ExitLlamadaFuncionesConParametro(c *LlamadaFuncionesConParametroContext)

	// ExitFnSinParametro is called when exiting the fnSinParametro production.
	ExitFnSinParametro(c *FnSinParametroContext)

	// ExitFnConParametro is called when exiting the fnConParametro production.
	ExitFnConParametro(c *FnConParametroContext)

	// ExitSliceDef is called when exiting the sliceDef production.
	ExitSliceDef(c *SliceDefContext)

	// ExitSliceLiteral is called when exiting the sliceLiteral production.
	ExitSliceLiteral(c *SliceLiteralContext)

	// ExitAccesoElementoSlice is called when exiting the accesoElementoSlice production.
	ExitAccesoElementoSlice(c *AccesoElementoSliceContext)

	// ExitModificacionElementoSlice is called when exiting the modificacionElementoSlice production.
	ExitModificacionElementoSlice(c *ModificacionElementoSliceContext)

	// ExitFnIndexOf is called when exiting the fnIndexOf production.
	ExitFnIndexOf(c *FnIndexOfContext)

	// ExitFnJoin is called when exiting the fnJoin production.
	ExitFnJoin(c *FnJoinContext)

	// ExitFnLen is called when exiting the fnLen production.
	ExitFnLen(c *FnLenContext)

	// ExitFnAppend is called when exiting the fnAppend production.
	ExitFnAppend(c *FnAppendContext)

	// ExitDeclaracionSliceVacio is called when exiting the declaracionSliceVacio production.
	ExitDeclaracionSliceVacio(c *DeclaracionSliceVacioContext)

	// ExitTipoRetorno is called when exiting the tipoRetorno production.
	ExitTipoRetorno(c *TipoRetornoContext)

	// ExitRetorno is called when exiting the retorno production.
	ExitRetorno(c *RetornoContext)

	// ExitFnTypeOf is called when exiting the fnTypeOf production.
	ExitFnTypeOf(c *FnTypeOfContext)

	// ExitFnAtoi is called when exiting the fnAtoi production.
	ExitFnAtoi(c *FnAtoiContext)

	// ExitFnParseToFloat is called when exiting the fnParseToFloat production.
	ExitFnParseToFloat(c *FnParseToFloatContext)

	// ExitAsigIncre is called when exiting the asigIncre production.
	ExitAsigIncre(c *AsigIncreContext)

	// ExitAsigDecre is called when exiting the asigDecre production.
	ExitAsigDecre(c *AsigDecreContext)

	// ExitAsignacionNormal is called when exiting the AsignacionNormal production.
	ExitAsignacionNormal(c *AsignacionNormalContext)

	// ExitAsignacionIncremento is called when exiting the AsignacionIncremento production.
	ExitAsignacionIncremento(c *AsignacionIncrementoContext)

	// ExitAsignacionDecremento is called when exiting the AsignacionDecremento production.
	ExitAsignacionDecremento(c *AsignacionDecrementoContext)

	// ExitListaIDS is called when exiting the listaIDS production.
	ExitListaIDS(c *ListaIDSContext)

	// ExitListaPar is called when exiting the listaPar production.
	ExitListaPar(c *ListaParContext)

	// ExitParametro is called when exiting the parametro production.
	ExitParametro(c *ParametroContext)

	// ExitListaExpr is called when exiting the listaExpr production.
	ExitListaExpr(c *ListaExprContext)

	// ExitListaExprList is called when exiting the listaExprList production.
	ExitListaExprList(c *ListaExprListContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitSif is called when exiting the sif production.
	ExitSif(c *SifContext)

	// ExitElseifPart is called when exiting the elseifPart production.
	ExitElseifPart(c *ElseifPartContext)

	// ExitForCondicional is called when exiting the ForCondicional production.
	ExitForCondicional(c *ForCondicionalContext)

	// ExitForClasico is called when exiting the ForClasico production.
	ExitForClasico(c *ForClasicoContext)

	// ExitForRange is called when exiting the ForRange production.
	ExitForRange(c *ForRangeContext)

	// ExitSSwitch is called when exiting the sSwitch production.
	ExitSSwitch(c *SSwitchContext)

	// ExitBreak_ is called when exiting the break_ production.
	ExitBreak_(c *Break_Context)

	// ExitContinue_ is called when exiting the continue_ production.
	ExitContinue_(c *Continue_Context)

	// ExitCaseBlock is called when exiting the caseBlock production.
	ExitCaseBlock(c *CaseBlockContext)

	// ExitDefaultBlock is called when exiting the defaultBlock production.
	ExitDefaultBlock(c *DefaultBlockContext)

	// ExitTipos is called when exiting the tipos production.
	ExitTipos(c *TiposContext)
}
