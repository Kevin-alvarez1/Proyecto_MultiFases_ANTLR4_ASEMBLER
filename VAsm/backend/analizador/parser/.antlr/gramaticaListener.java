// Generated from /home/vicente/Compi2Capi/OLC2_PROYECTO2_202203038/VAsm/backend/analizador/parser/gramatica.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link gramaticaParser}.
 */
public interface gramaticaListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#init}.
	 * @param ctx the parse tree
	 */
	void enterInit(gramaticaParser.InitContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#init}.
	 * @param ctx the parse tree
	 */
	void exitInit(gramaticaParser.InitContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterInstrucciones(gramaticaParser.InstruccionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitInstrucciones(gramaticaParser.InstruccionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#instruccion}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion(gramaticaParser.InstruccionContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#instruccion}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion(gramaticaParser.InstruccionContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#print}.
	 * @param ctx the parse tree
	 */
	void enterPrint(gramaticaParser.PrintContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#print}.
	 * @param ctx the parse tree
	 */
	void exitPrint(gramaticaParser.PrintContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#println}.
	 * @param ctx the parse tree
	 */
	void enterPrintln(gramaticaParser.PrintlnContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#println}.
	 * @param ctx the parse tree
	 */
	void exitPrintln(gramaticaParser.PrintlnContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#declaracionMultiple}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracionMultiple(gramaticaParser.DeclaracionMultipleContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#declaracionMultiple}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracionMultiple(gramaticaParser.DeclaracionMultipleContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#declaracionMultipleSimple}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracionMultipleSimple(gramaticaParser.DeclaracionMultipleSimpleContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#declaracionMultipleSimple}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracionMultipleSimple(gramaticaParser.DeclaracionMultipleSimpleContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#declaracionMultipleSinTipo}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracionMultipleSinTipo(gramaticaParser.DeclaracionMultipleSinTipoContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#declaracionMultipleSinTipo}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracionMultipleSinTipo(gramaticaParser.DeclaracionMultipleSinTipoContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#asignacionMultiple}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionMultiple(gramaticaParser.AsignacionMultipleContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#asignacionMultiple}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionMultiple(gramaticaParser.AsignacionMultipleContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#bloqueFuncion}.
	 * @param ctx the parse tree
	 */
	void enterBloqueFuncion(gramaticaParser.BloqueFuncionContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#bloqueFuncion}.
	 * @param ctx the parse tree
	 */
	void exitBloqueFuncion(gramaticaParser.BloqueFuncionContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#bloque}.
	 * @param ctx the parse tree
	 */
	void enterBloque(gramaticaParser.BloqueContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#bloque}.
	 * @param ctx the parse tree
	 */
	void exitBloque(gramaticaParser.BloqueContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#llamadaFuncionesSinParametro}.
	 * @param ctx the parse tree
	 */
	void enterLlamadaFuncionesSinParametro(gramaticaParser.LlamadaFuncionesSinParametroContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#llamadaFuncionesSinParametro}.
	 * @param ctx the parse tree
	 */
	void exitLlamadaFuncionesSinParametro(gramaticaParser.LlamadaFuncionesSinParametroContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#llamadaFuncionesConParametro}.
	 * @param ctx the parse tree
	 */
	void enterLlamadaFuncionesConParametro(gramaticaParser.LlamadaFuncionesConParametroContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#llamadaFuncionesConParametro}.
	 * @param ctx the parse tree
	 */
	void exitLlamadaFuncionesConParametro(gramaticaParser.LlamadaFuncionesConParametroContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#fnSinParametro}.
	 * @param ctx the parse tree
	 */
	void enterFnSinParametro(gramaticaParser.FnSinParametroContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#fnSinParametro}.
	 * @param ctx the parse tree
	 */
	void exitFnSinParametro(gramaticaParser.FnSinParametroContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#fnConParametro}.
	 * @param ctx the parse tree
	 */
	void enterFnConParametro(gramaticaParser.FnConParametroContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#fnConParametro}.
	 * @param ctx the parse tree
	 */
	void exitFnConParametro(gramaticaParser.FnConParametroContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#sliceDef}.
	 * @param ctx the parse tree
	 */
	void enterSliceDef(gramaticaParser.SliceDefContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#sliceDef}.
	 * @param ctx the parse tree
	 */
	void exitSliceDef(gramaticaParser.SliceDefContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#sliceLiteral}.
	 * @param ctx the parse tree
	 */
	void enterSliceLiteral(gramaticaParser.SliceLiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#sliceLiteral}.
	 * @param ctx the parse tree
	 */
	void exitSliceLiteral(gramaticaParser.SliceLiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#accesoElementoSlice}.
	 * @param ctx the parse tree
	 */
	void enterAccesoElementoSlice(gramaticaParser.AccesoElementoSliceContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#accesoElementoSlice}.
	 * @param ctx the parse tree
	 */
	void exitAccesoElementoSlice(gramaticaParser.AccesoElementoSliceContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#modificacionElementoSlice}.
	 * @param ctx the parse tree
	 */
	void enterModificacionElementoSlice(gramaticaParser.ModificacionElementoSliceContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#modificacionElementoSlice}.
	 * @param ctx the parse tree
	 */
	void exitModificacionElementoSlice(gramaticaParser.ModificacionElementoSliceContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#fnIndexOf}.
	 * @param ctx the parse tree
	 */
	void enterFnIndexOf(gramaticaParser.FnIndexOfContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#fnIndexOf}.
	 * @param ctx the parse tree
	 */
	void exitFnIndexOf(gramaticaParser.FnIndexOfContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#fnJoin}.
	 * @param ctx the parse tree
	 */
	void enterFnJoin(gramaticaParser.FnJoinContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#fnJoin}.
	 * @param ctx the parse tree
	 */
	void exitFnJoin(gramaticaParser.FnJoinContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#fnLen}.
	 * @param ctx the parse tree
	 */
	void enterFnLen(gramaticaParser.FnLenContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#fnLen}.
	 * @param ctx the parse tree
	 */
	void exitFnLen(gramaticaParser.FnLenContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#fnAppend}.
	 * @param ctx the parse tree
	 */
	void enterFnAppend(gramaticaParser.FnAppendContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#fnAppend}.
	 * @param ctx the parse tree
	 */
	void exitFnAppend(gramaticaParser.FnAppendContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#declaracionSliceVacio}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracionSliceVacio(gramaticaParser.DeclaracionSliceVacioContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#declaracionSliceVacio}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracionSliceVacio(gramaticaParser.DeclaracionSliceVacioContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#tipoRetorno}.
	 * @param ctx the parse tree
	 */
	void enterTipoRetorno(gramaticaParser.TipoRetornoContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#tipoRetorno}.
	 * @param ctx the parse tree
	 */
	void exitTipoRetorno(gramaticaParser.TipoRetornoContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#retorno}.
	 * @param ctx the parse tree
	 */
	void enterRetorno(gramaticaParser.RetornoContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#retorno}.
	 * @param ctx the parse tree
	 */
	void exitRetorno(gramaticaParser.RetornoContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#fnTypeOf}.
	 * @param ctx the parse tree
	 */
	void enterFnTypeOf(gramaticaParser.FnTypeOfContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#fnTypeOf}.
	 * @param ctx the parse tree
	 */
	void exitFnTypeOf(gramaticaParser.FnTypeOfContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#fnAtoi}.
	 * @param ctx the parse tree
	 */
	void enterFnAtoi(gramaticaParser.FnAtoiContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#fnAtoi}.
	 * @param ctx the parse tree
	 */
	void exitFnAtoi(gramaticaParser.FnAtoiContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#fnParseToFloat}.
	 * @param ctx the parse tree
	 */
	void enterFnParseToFloat(gramaticaParser.FnParseToFloatContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#fnParseToFloat}.
	 * @param ctx the parse tree
	 */
	void exitFnParseToFloat(gramaticaParser.FnParseToFloatContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#asigIncre}.
	 * @param ctx the parse tree
	 */
	void enterAsigIncre(gramaticaParser.AsigIncreContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#asigIncre}.
	 * @param ctx the parse tree
	 */
	void exitAsigIncre(gramaticaParser.AsigIncreContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#asigDecre}.
	 * @param ctx the parse tree
	 */
	void enterAsigDecre(gramaticaParser.AsigDecreContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#asigDecre}.
	 * @param ctx the parse tree
	 */
	void exitAsigDecre(gramaticaParser.AsigDecreContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AsignacionNormal}
	 * labeled alternative in {@link gramaticaParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionNormal(gramaticaParser.AsignacionNormalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AsignacionNormal}
	 * labeled alternative in {@link gramaticaParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionNormal(gramaticaParser.AsignacionNormalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AsignacionIncremento}
	 * labeled alternative in {@link gramaticaParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionIncremento(gramaticaParser.AsignacionIncrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AsignacionIncremento}
	 * labeled alternative in {@link gramaticaParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionIncremento(gramaticaParser.AsignacionIncrementoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AsignacionDecremento}
	 * labeled alternative in {@link gramaticaParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionDecremento(gramaticaParser.AsignacionDecrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AsignacionDecremento}
	 * labeled alternative in {@link gramaticaParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionDecremento(gramaticaParser.AsignacionDecrementoContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#listaIDS}.
	 * @param ctx the parse tree
	 */
	void enterListaIDS(gramaticaParser.ListaIDSContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#listaIDS}.
	 * @param ctx the parse tree
	 */
	void exitListaIDS(gramaticaParser.ListaIDSContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#listaPar}.
	 * @param ctx the parse tree
	 */
	void enterListaPar(gramaticaParser.ListaParContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#listaPar}.
	 * @param ctx the parse tree
	 */
	void exitListaPar(gramaticaParser.ListaParContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#parametro}.
	 * @param ctx the parse tree
	 */
	void enterParametro(gramaticaParser.ParametroContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#parametro}.
	 * @param ctx the parse tree
	 */
	void exitParametro(gramaticaParser.ParametroContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#listaExpr}.
	 * @param ctx the parse tree
	 */
	void enterListaExpr(gramaticaParser.ListaExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#listaExpr}.
	 * @param ctx the parse tree
	 */
	void exitListaExpr(gramaticaParser.ListaExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#listaExprList}.
	 * @param ctx the parse tree
	 */
	void enterListaExprList(gramaticaParser.ListaExprListContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#listaExprList}.
	 * @param ctx the parse tree
	 */
	void exitListaExprList(gramaticaParser.ListaExprListContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterExpresion(gramaticaParser.ExpresionContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitExpresion(gramaticaParser.ExpresionContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#sif}.
	 * @param ctx the parse tree
	 */
	void enterSif(gramaticaParser.SifContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#sif}.
	 * @param ctx the parse tree
	 */
	void exitSif(gramaticaParser.SifContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#elseifPart}.
	 * @param ctx the parse tree
	 */
	void enterElseifPart(gramaticaParser.ElseifPartContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#elseifPart}.
	 * @param ctx the parse tree
	 */
	void exitElseifPart(gramaticaParser.ElseifPartContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForCondicional}
	 * labeled alternative in {@link gramaticaParser#sfor}.
	 * @param ctx the parse tree
	 */
	void enterForCondicional(gramaticaParser.ForCondicionalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForCondicional}
	 * labeled alternative in {@link gramaticaParser#sfor}.
	 * @param ctx the parse tree
	 */
	void exitForCondicional(gramaticaParser.ForCondicionalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForClasico}
	 * labeled alternative in {@link gramaticaParser#sfor}.
	 * @param ctx the parse tree
	 */
	void enterForClasico(gramaticaParser.ForClasicoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForClasico}
	 * labeled alternative in {@link gramaticaParser#sfor}.
	 * @param ctx the parse tree
	 */
	void exitForClasico(gramaticaParser.ForClasicoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForRange}
	 * labeled alternative in {@link gramaticaParser#sfor}.
	 * @param ctx the parse tree
	 */
	void enterForRange(gramaticaParser.ForRangeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForRange}
	 * labeled alternative in {@link gramaticaParser#sfor}.
	 * @param ctx the parse tree
	 */
	void exitForRange(gramaticaParser.ForRangeContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#sSwitch}.
	 * @param ctx the parse tree
	 */
	void enterSSwitch(gramaticaParser.SSwitchContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#sSwitch}.
	 * @param ctx the parse tree
	 */
	void exitSSwitch(gramaticaParser.SSwitchContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#break_}.
	 * @param ctx the parse tree
	 */
	void enterBreak_(gramaticaParser.Break_Context ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#break_}.
	 * @param ctx the parse tree
	 */
	void exitBreak_(gramaticaParser.Break_Context ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#continue_}.
	 * @param ctx the parse tree
	 */
	void enterContinue_(gramaticaParser.Continue_Context ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#continue_}.
	 * @param ctx the parse tree
	 */
	void exitContinue_(gramaticaParser.Continue_Context ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#caseBlock}.
	 * @param ctx the parse tree
	 */
	void enterCaseBlock(gramaticaParser.CaseBlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#caseBlock}.
	 * @param ctx the parse tree
	 */
	void exitCaseBlock(gramaticaParser.CaseBlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#defaultBlock}.
	 * @param ctx the parse tree
	 */
	void enterDefaultBlock(gramaticaParser.DefaultBlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#defaultBlock}.
	 * @param ctx the parse tree
	 */
	void exitDefaultBlock(gramaticaParser.DefaultBlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#tipos}.
	 * @param ctx the parse tree
	 */
	void enterTipos(gramaticaParser.TiposContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#tipos}.
	 * @param ctx the parse tree
	 */
	void exitTipos(gramaticaParser.TiposContext ctx);
}