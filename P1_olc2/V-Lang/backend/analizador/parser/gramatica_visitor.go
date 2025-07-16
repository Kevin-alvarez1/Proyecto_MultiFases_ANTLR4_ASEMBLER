// Code generated from backend/analizador/parser/gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // gramatica
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by gramaticaParser.
type gramaticaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by gramaticaParser#init.
	VisitInit(ctx *InitContext) interface{}

	// Visit a parse tree produced by gramaticaParser#instrucciones.
	VisitInstrucciones(ctx *InstruccionesContext) interface{}

	// Visit a parse tree produced by gramaticaParser#instruccion.
	VisitInstruccion(ctx *InstruccionContext) interface{}

	// Visit a parse tree produced by gramaticaParser#print.
	VisitPrint(ctx *PrintContext) interface{}

	// Visit a parse tree produced by gramaticaParser#println.
	VisitPrintln(ctx *PrintlnContext) interface{}

	// Visit a parse tree produced by gramaticaParser#declaracionMultiple.
	VisitDeclaracionMultiple(ctx *DeclaracionMultipleContext) interface{}

	// Visit a parse tree produced by gramaticaParser#declaracionMultipleSimple.
	VisitDeclaracionMultipleSimple(ctx *DeclaracionMultipleSimpleContext) interface{}

	// Visit a parse tree produced by gramaticaParser#declaracionMultipleSinTipo.
	VisitDeclaracionMultipleSinTipo(ctx *DeclaracionMultipleSinTipoContext) interface{}

	// Visit a parse tree produced by gramaticaParser#asignacionMultiple.
	VisitAsignacionMultiple(ctx *AsignacionMultipleContext) interface{}

	// Visit a parse tree produced by gramaticaParser#structDef.
	VisitStructDef(ctx *StructDefContext) interface{}

	// Visit a parse tree produced by gramaticaParser#atributos.
	VisitAtributos(ctx *AtributosContext) interface{}

	// Visit a parse tree produced by gramaticaParser#structInst.
	VisitStructInst(ctx *StructInstContext) interface{}

	// Visit a parse tree produced by gramaticaParser#structInit.
	VisitStructInit(ctx *StructInitContext) interface{}

	// Visit a parse tree produced by gramaticaParser#listaStructs.
	VisitListaStructs(ctx *ListaStructsContext) interface{}

	// Visit a parse tree produced by gramaticaParser#accesoStruct.
	VisitAccesoStruct(ctx *AccesoStructContext) interface{}

	// Visit a parse tree produced by gramaticaParser#asigStruct.
	VisitAsigStruct(ctx *AsigStructContext) interface{}

	// Visit a parse tree produced by gramaticaParser#funStruct.
	VisitFunStruct(ctx *FunStructContext) interface{}

	// Visit a parse tree produced by gramaticaParser#caja.
	VisitCaja(ctx *CajaContext) interface{}

	// Visit a parse tree produced by gramaticaParser#valores.
	VisitValores(ctx *ValoresContext) interface{}

	// Visit a parse tree produced by gramaticaParser#actStruct.
	VisitActStruct(ctx *ActStructContext) interface{}

	// Visit a parse tree produced by gramaticaParser#bloqueFuncion.
	VisitBloqueFuncion(ctx *BloqueFuncionContext) interface{}

	// Visit a parse tree produced by gramaticaParser#llamadaFuncionesSinParametro.
	VisitLlamadaFuncionesSinParametro(ctx *LlamadaFuncionesSinParametroContext) interface{}

	// Visit a parse tree produced by gramaticaParser#llamadaFuncionesConParametro.
	VisitLlamadaFuncionesConParametro(ctx *LlamadaFuncionesConParametroContext) interface{}

	// Visit a parse tree produced by gramaticaParser#fnSinParametro.
	VisitFnSinParametro(ctx *FnSinParametroContext) interface{}

	// Visit a parse tree produced by gramaticaParser#fnConParametro.
	VisitFnConParametro(ctx *FnConParametroContext) interface{}

	// Visit a parse tree produced by gramaticaParser#sliceDef.
	VisitSliceDef(ctx *SliceDefContext) interface{}

	// Visit a parse tree produced by gramaticaParser#sliceLiteral.
	VisitSliceLiteral(ctx *SliceLiteralContext) interface{}

	// Visit a parse tree produced by gramaticaParser#accesoElementoSlice.
	VisitAccesoElementoSlice(ctx *AccesoElementoSliceContext) interface{}

	// Visit a parse tree produced by gramaticaParser#modificacionElementoSlice.
	VisitModificacionElementoSlice(ctx *ModificacionElementoSliceContext) interface{}

	// Visit a parse tree produced by gramaticaParser#fnIndexOf.
	VisitFnIndexOf(ctx *FnIndexOfContext) interface{}

	// Visit a parse tree produced by gramaticaParser#fnJoin.
	VisitFnJoin(ctx *FnJoinContext) interface{}

	// Visit a parse tree produced by gramaticaParser#fnLen.
	VisitFnLen(ctx *FnLenContext) interface{}

	// Visit a parse tree produced by gramaticaParser#fnAppend.
	VisitFnAppend(ctx *FnAppendContext) interface{}

	// Visit a parse tree produced by gramaticaParser#declaracionSliceVacio.
	VisitDeclaracionSliceVacio(ctx *DeclaracionSliceVacioContext) interface{}

	// Visit a parse tree produced by gramaticaParser#tipoRetorno.
	VisitTipoRetorno(ctx *TipoRetornoContext) interface{}

	// Visit a parse tree produced by gramaticaParser#retorno.
	VisitRetorno(ctx *RetornoContext) interface{}

	// Visit a parse tree produced by gramaticaParser#fnTypeOf.
	VisitFnTypeOf(ctx *FnTypeOfContext) interface{}

	// Visit a parse tree produced by gramaticaParser#fnAtoi.
	VisitFnAtoi(ctx *FnAtoiContext) interface{}

	// Visit a parse tree produced by gramaticaParser#fnParseToFloat.
	VisitFnParseToFloat(ctx *FnParseToFloatContext) interface{}

	// Visit a parse tree produced by gramaticaParser#asigIncre.
	VisitAsigIncre(ctx *AsigIncreContext) interface{}

	// Visit a parse tree produced by gramaticaParser#asigDecre.
	VisitAsigDecre(ctx *AsigDecreContext) interface{}

	// Visit a parse tree produced by gramaticaParser#AsignacionNormal.
	VisitAsignacionNormal(ctx *AsignacionNormalContext) interface{}

	// Visit a parse tree produced by gramaticaParser#AsignacionIncremento.
	VisitAsignacionIncremento(ctx *AsignacionIncrementoContext) interface{}

	// Visit a parse tree produced by gramaticaParser#AsignacionDecremento.
	VisitAsignacionDecremento(ctx *AsignacionDecrementoContext) interface{}

	// Visit a parse tree produced by gramaticaParser#listaIDS.
	VisitListaIDS(ctx *ListaIDSContext) interface{}

	// Visit a parse tree produced by gramaticaParser#listaPar.
	VisitListaPar(ctx *ListaParContext) interface{}

	// Visit a parse tree produced by gramaticaParser#parametro.
	VisitParametro(ctx *ParametroContext) interface{}

	// Visit a parse tree produced by gramaticaParser#listaExpr.
	VisitListaExpr(ctx *ListaExprContext) interface{}

	// Visit a parse tree produced by gramaticaParser#listaExprList.
	VisitListaExprList(ctx *ListaExprListContext) interface{}

	// Visit a parse tree produced by gramaticaParser#expresion.
	VisitExpresion(ctx *ExpresionContext) interface{}

	// Visit a parse tree produced by gramaticaParser#sif.
	VisitSif(ctx *SifContext) interface{}

	// Visit a parse tree produced by gramaticaParser#elseifPart.
	VisitElseifPart(ctx *ElseifPartContext) interface{}

	// Visit a parse tree produced by gramaticaParser#bloque.
	VisitBloque(ctx *BloqueContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ForCondicional.
	VisitForCondicional(ctx *ForCondicionalContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ForClasico.
	VisitForClasico(ctx *ForClasicoContext) interface{}

	// Visit a parse tree produced by gramaticaParser#ForRange.
	VisitForRange(ctx *ForRangeContext) interface{}

	// Visit a parse tree produced by gramaticaParser#sSwitch.
	VisitSSwitch(ctx *SSwitchContext) interface{}

	// Visit a parse tree produced by gramaticaParser#break_.
	VisitBreak_(ctx *Break_Context) interface{}

	// Visit a parse tree produced by gramaticaParser#continue_.
	VisitContinue_(ctx *Continue_Context) interface{}

	// Visit a parse tree produced by gramaticaParser#caseBlock.
	VisitCaseBlock(ctx *CaseBlockContext) interface{}

	// Visit a parse tree produced by gramaticaParser#defaultBlock.
	VisitDefaultBlock(ctx *DefaultBlockContext) interface{}

	// Visit a parse tree produced by gramaticaParser#tipos.
	VisitTipos(ctx *TiposContext) interface{}
}
