// Code generated from backend/analizador/parser/gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // gramatica
import "github.com/antlr4-go/antlr/v4"

// BasegramaticaListener is a complete listener for a parse tree produced by gramaticaParser.
type BasegramaticaListener struct{}

var _ gramaticaListener = &BasegramaticaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegramaticaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegramaticaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegramaticaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegramaticaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInit is called when production init is entered.
func (s *BasegramaticaListener) EnterInit(ctx *InitContext) {}

// ExitInit is called when production init is exited.
func (s *BasegramaticaListener) ExitInit(ctx *InitContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BasegramaticaListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BasegramaticaListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BasegramaticaListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BasegramaticaListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterPrint is called when production print is entered.
func (s *BasegramaticaListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BasegramaticaListener) ExitPrint(ctx *PrintContext) {}

// EnterPrintln is called when production println is entered.
func (s *BasegramaticaListener) EnterPrintln(ctx *PrintlnContext) {}

// ExitPrintln is called when production println is exited.
func (s *BasegramaticaListener) ExitPrintln(ctx *PrintlnContext) {}

// EnterDeclaracionMultiple is called when production declaracionMultiple is entered.
func (s *BasegramaticaListener) EnterDeclaracionMultiple(ctx *DeclaracionMultipleContext) {}

// ExitDeclaracionMultiple is called when production declaracionMultiple is exited.
func (s *BasegramaticaListener) ExitDeclaracionMultiple(ctx *DeclaracionMultipleContext) {}

// EnterDeclaracionMultipleSimple is called when production declaracionMultipleSimple is entered.
func (s *BasegramaticaListener) EnterDeclaracionMultipleSimple(ctx *DeclaracionMultipleSimpleContext) {
}

// ExitDeclaracionMultipleSimple is called when production declaracionMultipleSimple is exited.
func (s *BasegramaticaListener) ExitDeclaracionMultipleSimple(ctx *DeclaracionMultipleSimpleContext) {
}

// EnterDeclaracionMultipleSinTipo is called when production declaracionMultipleSinTipo is entered.
func (s *BasegramaticaListener) EnterDeclaracionMultipleSinTipo(ctx *DeclaracionMultipleSinTipoContext) {
}

// ExitDeclaracionMultipleSinTipo is called when production declaracionMultipleSinTipo is exited.
func (s *BasegramaticaListener) ExitDeclaracionMultipleSinTipo(ctx *DeclaracionMultipleSinTipoContext) {
}

// EnterAsignacionMultiple is called when production asignacionMultiple is entered.
func (s *BasegramaticaListener) EnterAsignacionMultiple(ctx *AsignacionMultipleContext) {}

// ExitAsignacionMultiple is called when production asignacionMultiple is exited.
func (s *BasegramaticaListener) ExitAsignacionMultiple(ctx *AsignacionMultipleContext) {}

// EnterBloqueFuncion is called when production bloqueFuncion is entered.
func (s *BasegramaticaListener) EnterBloqueFuncion(ctx *BloqueFuncionContext) {}

// ExitBloqueFuncion is called when production bloqueFuncion is exited.
func (s *BasegramaticaListener) ExitBloqueFuncion(ctx *BloqueFuncionContext) {}

// EnterLlamadaFuncionesSinParametro is called when production llamadaFuncionesSinParametro is entered.
func (s *BasegramaticaListener) EnterLlamadaFuncionesSinParametro(ctx *LlamadaFuncionesSinParametroContext) {
}

// ExitLlamadaFuncionesSinParametro is called when production llamadaFuncionesSinParametro is exited.
func (s *BasegramaticaListener) ExitLlamadaFuncionesSinParametro(ctx *LlamadaFuncionesSinParametroContext) {
}

// EnterLlamadaFuncionesConParametro is called when production llamadaFuncionesConParametro is entered.
func (s *BasegramaticaListener) EnterLlamadaFuncionesConParametro(ctx *LlamadaFuncionesConParametroContext) {
}

// ExitLlamadaFuncionesConParametro is called when production llamadaFuncionesConParametro is exited.
func (s *BasegramaticaListener) ExitLlamadaFuncionesConParametro(ctx *LlamadaFuncionesConParametroContext) {
}

// EnterFnSinParametro is called when production fnSinParametro is entered.
func (s *BasegramaticaListener) EnterFnSinParametro(ctx *FnSinParametroContext) {}

// ExitFnSinParametro is called when production fnSinParametro is exited.
func (s *BasegramaticaListener) ExitFnSinParametro(ctx *FnSinParametroContext) {}

// EnterFnConParametro is called when production fnConParametro is entered.
func (s *BasegramaticaListener) EnterFnConParametro(ctx *FnConParametroContext) {}

// ExitFnConParametro is called when production fnConParametro is exited.
func (s *BasegramaticaListener) ExitFnConParametro(ctx *FnConParametroContext) {}

// EnterSliceDef is called when production sliceDef is entered.
func (s *BasegramaticaListener) EnterSliceDef(ctx *SliceDefContext) {}

// ExitSliceDef is called when production sliceDef is exited.
func (s *BasegramaticaListener) ExitSliceDef(ctx *SliceDefContext) {}

// EnterSliceLiteral is called when production sliceLiteral is entered.
func (s *BasegramaticaListener) EnterSliceLiteral(ctx *SliceLiteralContext) {}

// ExitSliceLiteral is called when production sliceLiteral is exited.
func (s *BasegramaticaListener) ExitSliceLiteral(ctx *SliceLiteralContext) {}

// EnterAccesoElementoSlice is called when production accesoElementoSlice is entered.
func (s *BasegramaticaListener) EnterAccesoElementoSlice(ctx *AccesoElementoSliceContext) {}

// ExitAccesoElementoSlice is called when production accesoElementoSlice is exited.
func (s *BasegramaticaListener) ExitAccesoElementoSlice(ctx *AccesoElementoSliceContext) {}

// EnterModificacionElementoSlice is called when production modificacionElementoSlice is entered.
func (s *BasegramaticaListener) EnterModificacionElementoSlice(ctx *ModificacionElementoSliceContext) {
}

// ExitModificacionElementoSlice is called when production modificacionElementoSlice is exited.
func (s *BasegramaticaListener) ExitModificacionElementoSlice(ctx *ModificacionElementoSliceContext) {
}

// EnterFnIndexOf is called when production fnIndexOf is entered.
func (s *BasegramaticaListener) EnterFnIndexOf(ctx *FnIndexOfContext) {}

// ExitFnIndexOf is called when production fnIndexOf is exited.
func (s *BasegramaticaListener) ExitFnIndexOf(ctx *FnIndexOfContext) {}

// EnterFnJoin is called when production fnJoin is entered.
func (s *BasegramaticaListener) EnterFnJoin(ctx *FnJoinContext) {}

// ExitFnJoin is called when production fnJoin is exited.
func (s *BasegramaticaListener) ExitFnJoin(ctx *FnJoinContext) {}

// EnterFnLen is called when production fnLen is entered.
func (s *BasegramaticaListener) EnterFnLen(ctx *FnLenContext) {}

// ExitFnLen is called when production fnLen is exited.
func (s *BasegramaticaListener) ExitFnLen(ctx *FnLenContext) {}

// EnterFnAppend is called when production fnAppend is entered.
func (s *BasegramaticaListener) EnterFnAppend(ctx *FnAppendContext) {}

// ExitFnAppend is called when production fnAppend is exited.
func (s *BasegramaticaListener) ExitFnAppend(ctx *FnAppendContext) {}

// EnterDeclaracionSliceVacio is called when production declaracionSliceVacio is entered.
func (s *BasegramaticaListener) EnterDeclaracionSliceVacio(ctx *DeclaracionSliceVacioContext) {}

// ExitDeclaracionSliceVacio is called when production declaracionSliceVacio is exited.
func (s *BasegramaticaListener) ExitDeclaracionSliceVacio(ctx *DeclaracionSliceVacioContext) {}

// EnterTipoRetorno is called when production tipoRetorno is entered.
func (s *BasegramaticaListener) EnterTipoRetorno(ctx *TipoRetornoContext) {}

// ExitTipoRetorno is called when production tipoRetorno is exited.
func (s *BasegramaticaListener) ExitTipoRetorno(ctx *TipoRetornoContext) {}

// EnterRetorno is called when production retorno is entered.
func (s *BasegramaticaListener) EnterRetorno(ctx *RetornoContext) {}

// ExitRetorno is called when production retorno is exited.
func (s *BasegramaticaListener) ExitRetorno(ctx *RetornoContext) {}

// EnterFnTypeOf is called when production fnTypeOf is entered.
func (s *BasegramaticaListener) EnterFnTypeOf(ctx *FnTypeOfContext) {}

// ExitFnTypeOf is called when production fnTypeOf is exited.
func (s *BasegramaticaListener) ExitFnTypeOf(ctx *FnTypeOfContext) {}

// EnterFnAtoi is called when production fnAtoi is entered.
func (s *BasegramaticaListener) EnterFnAtoi(ctx *FnAtoiContext) {}

// ExitFnAtoi is called when production fnAtoi is exited.
func (s *BasegramaticaListener) ExitFnAtoi(ctx *FnAtoiContext) {}

// EnterFnParseToFloat is called when production fnParseToFloat is entered.
func (s *BasegramaticaListener) EnterFnParseToFloat(ctx *FnParseToFloatContext) {}

// ExitFnParseToFloat is called when production fnParseToFloat is exited.
func (s *BasegramaticaListener) ExitFnParseToFloat(ctx *FnParseToFloatContext) {}

// EnterAsigIncre is called when production asigIncre is entered.
func (s *BasegramaticaListener) EnterAsigIncre(ctx *AsigIncreContext) {}

// ExitAsigIncre is called when production asigIncre is exited.
func (s *BasegramaticaListener) ExitAsigIncre(ctx *AsigIncreContext) {}

// EnterAsigDecre is called when production asigDecre is entered.
func (s *BasegramaticaListener) EnterAsigDecre(ctx *AsigDecreContext) {}

// ExitAsigDecre is called when production asigDecre is exited.
func (s *BasegramaticaListener) ExitAsigDecre(ctx *AsigDecreContext) {}

// EnterAsignacionNormal is called when production AsignacionNormal is entered.
func (s *BasegramaticaListener) EnterAsignacionNormal(ctx *AsignacionNormalContext) {}

// ExitAsignacionNormal is called when production AsignacionNormal is exited.
func (s *BasegramaticaListener) ExitAsignacionNormal(ctx *AsignacionNormalContext) {}

// EnterAsignacionIncremento is called when production AsignacionIncremento is entered.
func (s *BasegramaticaListener) EnterAsignacionIncremento(ctx *AsignacionIncrementoContext) {}

// ExitAsignacionIncremento is called when production AsignacionIncremento is exited.
func (s *BasegramaticaListener) ExitAsignacionIncremento(ctx *AsignacionIncrementoContext) {}

// EnterAsignacionDecremento is called when production AsignacionDecremento is entered.
func (s *BasegramaticaListener) EnterAsignacionDecremento(ctx *AsignacionDecrementoContext) {}

// ExitAsignacionDecremento is called when production AsignacionDecremento is exited.
func (s *BasegramaticaListener) ExitAsignacionDecremento(ctx *AsignacionDecrementoContext) {}

// EnterListaIDS is called when production listaIDS is entered.
func (s *BasegramaticaListener) EnterListaIDS(ctx *ListaIDSContext) {}

// ExitListaIDS is called when production listaIDS is exited.
func (s *BasegramaticaListener) ExitListaIDS(ctx *ListaIDSContext) {}

// EnterListaPar is called when production listaPar is entered.
func (s *BasegramaticaListener) EnterListaPar(ctx *ListaParContext) {}

// ExitListaPar is called when production listaPar is exited.
func (s *BasegramaticaListener) ExitListaPar(ctx *ListaParContext) {}

// EnterParametro is called when production parametro is entered.
func (s *BasegramaticaListener) EnterParametro(ctx *ParametroContext) {}

// ExitParametro is called when production parametro is exited.
func (s *BasegramaticaListener) ExitParametro(ctx *ParametroContext) {}

// EnterListaExpr is called when production listaExpr is entered.
func (s *BasegramaticaListener) EnterListaExpr(ctx *ListaExprContext) {}

// ExitListaExpr is called when production listaExpr is exited.
func (s *BasegramaticaListener) ExitListaExpr(ctx *ListaExprContext) {}

// EnterListaExprList is called when production listaExprList is entered.
func (s *BasegramaticaListener) EnterListaExprList(ctx *ListaExprListContext) {}

// ExitListaExprList is called when production listaExprList is exited.
func (s *BasegramaticaListener) ExitListaExprList(ctx *ListaExprListContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BasegramaticaListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BasegramaticaListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterSif is called when production sif is entered.
func (s *BasegramaticaListener) EnterSif(ctx *SifContext) {}

// ExitSif is called when production sif is exited.
func (s *BasegramaticaListener) ExitSif(ctx *SifContext) {}

// EnterElseifPart is called when production elseifPart is entered.
func (s *BasegramaticaListener) EnterElseifPart(ctx *ElseifPartContext) {}

// ExitElseifPart is called when production elseifPart is exited.
func (s *BasegramaticaListener) ExitElseifPart(ctx *ElseifPartContext) {}

// EnterBloque is called when production bloque is entered.
func (s *BasegramaticaListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *BasegramaticaListener) ExitBloque(ctx *BloqueContext) {}

// EnterForCondicional is called when production ForCondicional is entered.
func (s *BasegramaticaListener) EnterForCondicional(ctx *ForCondicionalContext) {}

// ExitForCondicional is called when production ForCondicional is exited.
func (s *BasegramaticaListener) ExitForCondicional(ctx *ForCondicionalContext) {}

// EnterForClasico is called when production ForClasico is entered.
func (s *BasegramaticaListener) EnterForClasico(ctx *ForClasicoContext) {}

// ExitForClasico is called when production ForClasico is exited.
func (s *BasegramaticaListener) ExitForClasico(ctx *ForClasicoContext) {}

// EnterForRange is called when production ForRange is entered.
func (s *BasegramaticaListener) EnterForRange(ctx *ForRangeContext) {}

// ExitForRange is called when production ForRange is exited.
func (s *BasegramaticaListener) ExitForRange(ctx *ForRangeContext) {}

// EnterSSwitch is called when production sSwitch is entered.
func (s *BasegramaticaListener) EnterSSwitch(ctx *SSwitchContext) {}

// ExitSSwitch is called when production sSwitch is exited.
func (s *BasegramaticaListener) ExitSSwitch(ctx *SSwitchContext) {}

// EnterBreak_ is called when production break_ is entered.
func (s *BasegramaticaListener) EnterBreak_(ctx *Break_Context) {}

// ExitBreak_ is called when production break_ is exited.
func (s *BasegramaticaListener) ExitBreak_(ctx *Break_Context) {}

// EnterContinue_ is called when production continue_ is entered.
func (s *BasegramaticaListener) EnterContinue_(ctx *Continue_Context) {}

// ExitContinue_ is called when production continue_ is exited.
func (s *BasegramaticaListener) ExitContinue_(ctx *Continue_Context) {}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BasegramaticaListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BasegramaticaListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterDefaultBlock is called when production defaultBlock is entered.
func (s *BasegramaticaListener) EnterDefaultBlock(ctx *DefaultBlockContext) {}

// ExitDefaultBlock is called when production defaultBlock is exited.
func (s *BasegramaticaListener) ExitDefaultBlock(ctx *DefaultBlockContext) {}

// EnterTipos is called when production tipos is entered.
func (s *BasegramaticaListener) EnterTipos(ctx *TiposContext) {}

// ExitTipos is called when production tipos is exited.
func (s *BasegramaticaListener) ExitTipos(ctx *TiposContext) {}
