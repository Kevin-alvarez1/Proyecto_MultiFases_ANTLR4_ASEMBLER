// Code generated from backend/analizador/parser/gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // gramatica
import "github.com/antlr4-go/antlr/v4"

type BasegramaticaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasegramaticaVisitor) VisitInit(ctx *InitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitInstrucciones(ctx *InstruccionesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitInstruccion(ctx *InstruccionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitPrint(ctx *PrintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitPrintln(ctx *PrintlnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDeclaracionMultiple(ctx *DeclaracionMultipleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDeclaracionMultipleSimple(ctx *DeclaracionMultipleSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDeclaracionMultipleSinTipo(ctx *DeclaracionMultipleSinTipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAsignacionMultiple(ctx *AsignacionMultipleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitBloqueFuncion(ctx *BloqueFuncionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitBloque(ctx *BloqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLlamadaFuncionesSinParametro(ctx *LlamadaFuncionesSinParametroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitLlamadaFuncionesConParametro(ctx *LlamadaFuncionesConParametroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFnSinParametro(ctx *FnSinParametroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFnConParametro(ctx *FnConParametroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSliceDef(ctx *SliceDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSliceLiteral(ctx *SliceLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAccesoElementoSlice(ctx *AccesoElementoSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitModificacionElementoSlice(ctx *ModificacionElementoSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFnIndexOf(ctx *FnIndexOfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFnJoin(ctx *FnJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFnLen(ctx *FnLenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFnAppend(ctx *FnAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDeclaracionSliceVacio(ctx *DeclaracionSliceVacioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitTipoRetorno(ctx *TipoRetornoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitRetorno(ctx *RetornoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFnTypeOf(ctx *FnTypeOfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFnAtoi(ctx *FnAtoiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitFnParseToFloat(ctx *FnParseToFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAsigIncre(ctx *AsigIncreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAsigDecre(ctx *AsigDecreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAsignacionNormal(ctx *AsignacionNormalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAsignacionIncremento(ctx *AsignacionIncrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitAsignacionDecremento(ctx *AsignacionDecrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitListaIDS(ctx *ListaIDSContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitListaPar(ctx *ListaParContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitParametro(ctx *ParametroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitListaExpr(ctx *ListaExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitListaExprList(ctx *ListaExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitExpresion(ctx *ExpresionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSif(ctx *SifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitElseifPart(ctx *ElseifPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitForCondicional(ctx *ForCondicionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitForClasico(ctx *ForClasicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitForRange(ctx *ForRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitSSwitch(ctx *SSwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitBreak_(ctx *Break_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitContinue_(ctx *Continue_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitCaseBlock(ctx *CaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitDefaultBlock(ctx *DefaultBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasegramaticaVisitor) VisitTipos(ctx *TiposContext) interface{} {
	return v.VisitChildren(ctx)
}
