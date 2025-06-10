// Generated from /home/sebas/Desktop/Compiladores 2/OLC2_EVJUNIO2025/Clase2/compiler/parser/Vlang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class VlangParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, LEN=3, CAP=4, APPEND=5, MUT=6, BOOLEANO=7, ENTERO=8, FLOAT=9, 
		DECIMAL=10, CADENA=11, CARACTER=12, ID=13, IF_KW=14, ELSE_KW=15, WHILE_KW=16, 
		FOR_KW=17, IN_KW=18, RETURN_KW=19, BREAK_KW=20, CONTINUE_KW=21, PLUS=22, 
		MINUS=23, MUL=24, DIV=25, MOD=26, NOT=27, OR=28, AND=29, EQ=30, NEQ=31, 
		LT=32, LE=33, GT=34, GE=35, ASSIGN=36, INC=37, DEC=38, LPAREN=39, RPAREN=40, 
		LBRACK=41, LCOR=42, RBRACK=43, RCOR=44, DOT=45, COMMA=46, WS=47, LINE_COMMENT=48, 
		BLOCK_COMMENT=49;
	public static final int
		RULE_programa = 0, RULE_stmt = 1, RULE_if_stmt = 2, RULE_if_chain = 3, 
		RULE_else_stmt = 4, RULE_while_stmt = 5, RULE_for_stmt = 6, RULE_range = 7, 
		RULE_transfer_stmt = 8, RULE_assign_stmt = 9, RULE_decl_stmt = 10, RULE_id_pattern = 11, 
		RULE_expresion = 12, RULE_parametros = 13, RULE_valores = 14, RULE_valor = 15, 
		RULE_incredecre = 16;
	private static String[] makeRuleNames() {
		return new String[] {
			"programa", "stmt", "if_stmt", "if_chain", "else_stmt", "while_stmt", 
			"for_stmt", "range", "transfer_stmt", "assign_stmt", "decl_stmt", "id_pattern", 
			"expresion", "parametros", "valores", "valor", "incredecre"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'println('", "'print('", "'len'", "'cap'", "'append'", "'mut'", 
			null, null, null, null, null, null, null, "'if'", "'else'", "'while'", 
			"'for'", "'in'", "'return'", "'break'", "'continue'", "'+'", "'-'", "'*'", 
			"'/'", "'%'", "'!'", "'||'", "'&&'", "'=='", "'!='", "'<'", "'<='", "'>'", 
			"'>='", "'='", "'++'", "'--'", "'('", "')'", "'['", "'{'", "']'", "'}'", 
			"'.'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, "LEN", "CAP", "APPEND", "MUT", "BOOLEANO", "ENTERO", 
			"FLOAT", "DECIMAL", "CADENA", "CARACTER", "ID", "IF_KW", "ELSE_KW", "WHILE_KW", 
			"FOR_KW", "IN_KW", "RETURN_KW", "BREAK_KW", "CONTINUE_KW", "PLUS", "MINUS", 
			"MUL", "DIV", "MOD", "NOT", "OR", "AND", "EQ", "NEQ", "LT", "LE", "GT", 
			"GE", "ASSIGN", "INC", "DEC", "LPAREN", "RPAREN", "LBRACK", "LCOR", "RBRACK", 
			"RCOR", "DOT", "COMMA", "WS", "LINE_COMMENT", "BLOCK_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Vlang.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public VlangParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramaContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(VlangParser.EOF, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ProgramaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_programa; }
	}

	public final ProgramaContext programa() throws RecognitionException {
		ProgramaContext _localctx = new ProgramaContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_programa);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(35); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(34);
				stmt();
				}
				}
				setState(37); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 3891270L) != 0) );
			setState(39);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StmtContext extends ParserRuleContext {
		public Decl_stmtContext decl_stmt() {
			return getRuleContext(Decl_stmtContext.class,0);
		}
		public Assign_stmtContext assign_stmt() {
			return getRuleContext(Assign_stmtContext.class,0);
		}
		public If_stmtContext if_stmt() {
			return getRuleContext(If_stmtContext.class,0);
		}
		public While_stmtContext while_stmt() {
			return getRuleContext(While_stmtContext.class,0);
		}
		public For_stmtContext for_stmt() {
			return getRuleContext(For_stmtContext.class,0);
		}
		public Transfer_stmtContext transfer_stmt() {
			return getRuleContext(Transfer_stmtContext.class,0);
		}
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_stmt);
		try {
			setState(55);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MUT:
				enterOuterAlt(_localctx, 1);
				{
				setState(41);
				decl_stmt();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 2);
				{
				setState(42);
				assign_stmt();
				}
				break;
			case IF_KW:
				enterOuterAlt(_localctx, 3);
				{
				setState(43);
				if_stmt();
				}
				break;
			case WHILE_KW:
				enterOuterAlt(_localctx, 4);
				{
				setState(44);
				while_stmt();
				}
				break;
			case FOR_KW:
				enterOuterAlt(_localctx, 5);
				{
				setState(45);
				for_stmt();
				}
				break;
			case RETURN_KW:
			case BREAK_KW:
			case CONTINUE_KW:
				enterOuterAlt(_localctx, 6);
				{
				setState(46);
				transfer_stmt();
				}
				break;
			case T__0:
				enterOuterAlt(_localctx, 7);
				{
				setState(47);
				match(T__0);
				setState(48);
				expresion(0);
				setState(49);
				match(RPAREN);
				}
				break;
			case T__1:
				enterOuterAlt(_localctx, 8);
				{
				setState(51);
				match(T__1);
				setState(52);
				expresion(0);
				setState(53);
				match(RPAREN);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class If_stmtContext extends ParserRuleContext {
		public If_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_stmt; }
	 
		public If_stmtContext() { }
		public void copyFrom(If_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfStmtContext extends If_stmtContext {
		public List<If_chainContext> if_chain() {
			return getRuleContexts(If_chainContext.class);
		}
		public If_chainContext if_chain(int i) {
			return getRuleContext(If_chainContext.class,i);
		}
		public List<TerminalNode> ELSE_KW() { return getTokens(VlangParser.ELSE_KW); }
		public TerminalNode ELSE_KW(int i) {
			return getToken(VlangParser.ELSE_KW, i);
		}
		public Else_stmtContext else_stmt() {
			return getRuleContext(Else_stmtContext.class,0);
		}
		public IfStmtContext(If_stmtContext ctx) { copyFrom(ctx); }
	}

	public final If_stmtContext if_stmt() throws RecognitionException {
		If_stmtContext _localctx = new If_stmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_if_stmt);
		int _la;
		try {
			int _alt;
			_localctx = new IfStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(57);
			if_chain();
			setState(62);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(58);
					match(ELSE_KW);
					setState(59);
					if_chain();
					}
					} 
				}
				setState(64);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			setState(66);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE_KW) {
				{
				setState(65);
				else_stmt();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class If_chainContext extends ParserRuleContext {
		public If_chainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_chain; }
	 
		public If_chainContext() { }
		public void copyFrom(If_chainContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfChainContext extends If_chainContext {
		public TerminalNode IF_KW() { return getToken(VlangParser.IF_KW, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public IfChainContext(If_chainContext ctx) { copyFrom(ctx); }
	}

	public final If_chainContext if_chain() throws RecognitionException {
		If_chainContext _localctx = new If_chainContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_if_chain);
		int _la;
		try {
			_localctx = new IfChainContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(68);
			match(IF_KW);
			setState(69);
			expresion(0);
			setState(70);
			match(LCOR);
			setState(74);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3891270L) != 0)) {
				{
				{
				setState(71);
				stmt();
				}
				}
				setState(76);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(77);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Else_stmtContext extends ParserRuleContext {
		public Else_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_else_stmt; }
	 
		public Else_stmtContext() { }
		public void copyFrom(Else_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ElseStmtContext extends Else_stmtContext {
		public TerminalNode ELSE_KW() { return getToken(VlangParser.ELSE_KW, 0); }
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ElseStmtContext(Else_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Else_stmtContext else_stmt() throws RecognitionException {
		Else_stmtContext _localctx = new Else_stmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_else_stmt);
		int _la;
		try {
			_localctx = new ElseStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(79);
			match(ELSE_KW);
			setState(80);
			match(LCOR);
			setState(84);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3891270L) != 0)) {
				{
				{
				setState(81);
				stmt();
				}
				}
				setState(86);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(87);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class While_stmtContext extends ParserRuleContext {
		public While_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_stmt; }
	 
		public While_stmtContext() { }
		public void copyFrom(While_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class WhileStmtContext extends While_stmtContext {
		public TerminalNode WHILE_KW() { return getToken(VlangParser.WHILE_KW, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public WhileStmtContext(While_stmtContext ctx) { copyFrom(ctx); }
	}

	public final While_stmtContext while_stmt() throws RecognitionException {
		While_stmtContext _localctx = new While_stmtContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_while_stmt);
		int _la;
		try {
			_localctx = new WhileStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			match(WHILE_KW);
			setState(90);
			expresion(0);
			setState(91);
			match(LCOR);
			setState(95);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3891270L) != 0)) {
				{
				{
				setState(92);
				stmt();
				}
				}
				setState(97);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(98);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class For_stmtContext extends ParserRuleContext {
		public For_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_stmt; }
	 
		public For_stmtContext() { }
		public void copyFrom(For_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForStmtContext extends For_stmtContext {
		public TerminalNode FOR_KW() { return getToken(VlangParser.FOR_KW, 0); }
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode IN_KW() { return getToken(VlangParser.IN_KW, 0); }
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public RangeContext range() {
			return getRuleContext(RangeContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ForStmtContext(For_stmtContext ctx) { copyFrom(ctx); }
	}

	public final For_stmtContext for_stmt() throws RecognitionException {
		For_stmtContext _localctx = new For_stmtContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_for_stmt);
		int _la;
		try {
			_localctx = new ForStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(100);
			match(FOR_KW);
			setState(101);
			match(ID);
			setState(102);
			match(IN_KW);
			setState(105);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				{
				setState(103);
				expresion(0);
				}
				break;
			case 2:
				{
				setState(104);
				range();
				}
				break;
			}
			setState(107);
			match(LCOR);
			setState(111);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3891270L) != 0)) {
				{
				{
				setState(108);
				stmt();
				}
				}
				setState(113);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(114);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RangeContext extends ParserRuleContext {
		public RangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_range; }
	 
		public RangeContext() { }
		public void copyFrom(RangeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NumericRangeContext extends RangeContext {
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<TerminalNode> DOT() { return getTokens(VlangParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(VlangParser.DOT, i);
		}
		public NumericRangeContext(RangeContext ctx) { copyFrom(ctx); }
	}

	public final RangeContext range() throws RecognitionException {
		RangeContext _localctx = new RangeContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_range);
		try {
			_localctx = new NumericRangeContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(116);
			expresion(0);
			setState(117);
			match(DOT);
			setState(118);
			match(DOT);
			setState(119);
			match(DOT);
			setState(120);
			expresion(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Transfer_stmtContext extends ParserRuleContext {
		public Transfer_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_transfer_stmt; }
	 
		public Transfer_stmtContext() { }
		public void copyFrom(Transfer_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStmtContext extends Transfer_stmtContext {
		public TerminalNode CONTINUE_KW() { return getToken(VlangParser.CONTINUE_KW, 0); }
		public ContinueStmtContext(Transfer_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakStmtContext extends Transfer_stmtContext {
		public TerminalNode BREAK_KW() { return getToken(VlangParser.BREAK_KW, 0); }
		public BreakStmtContext(Transfer_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStmtContext extends Transfer_stmtContext {
		public TerminalNode RETURN_KW() { return getToken(VlangParser.RETURN_KW, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ReturnStmtContext(Transfer_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Transfer_stmtContext transfer_stmt() throws RecognitionException {
		Transfer_stmtContext _localctx = new Transfer_stmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_transfer_stmt);
		try {
			setState(128);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RETURN_KW:
				_localctx = new ReturnStmtContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(122);
				match(RETURN_KW);
				setState(124);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
				case 1:
					{
					setState(123);
					expresion(0);
					}
					break;
				}
				}
				break;
			case BREAK_KW:
				_localctx = new BreakStmtContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(126);
				match(BREAK_KW);
				}
				break;
			case CONTINUE_KW:
				_localctx = new ContinueStmtContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(127);
				match(CONTINUE_KW);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Assign_stmtContext extends ParserRuleContext {
		public Assign_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_stmt; }
	 
		public Assign_stmtContext() { }
		public void copyFrom(Assign_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DirectAssignContext extends Assign_stmtContext {
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public DirectAssignContext(Assign_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Assign_stmtContext assign_stmt() throws RecognitionException {
		Assign_stmtContext _localctx = new Assign_stmtContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_assign_stmt);
		try {
			_localctx = new DirectAssignContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			id_pattern();
			setState(131);
			match(ASSIGN);
			setState(132);
			expresion(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Decl_stmtContext extends ParserRuleContext {
		public Decl_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decl_stmt; }
	 
		public Decl_stmtContext() { }
		public void copyFrom(Decl_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclAssignContext extends Decl_stmtContext {
		public TerminalNode MUT() { return getToken(VlangParser.MUT, 0); }
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public DeclAssignContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclPatternContext extends Decl_stmtContext {
		public TerminalNode MUT() { return getToken(VlangParser.MUT, 0); }
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public DeclPatternContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclAssignPatternContext extends Decl_stmtContext {
		public TerminalNode MUT() { return getToken(VlangParser.MUT, 0); }
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public DeclAssignPatternContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Decl_stmtContext decl_stmt() throws RecognitionException {
		Decl_stmtContext _localctx = new Decl_stmtContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_decl_stmt);
		try {
			setState(145);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				_localctx = new DeclAssignContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(134);
				match(MUT);
				setState(135);
				match(ID);
				setState(136);
				match(ASSIGN);
				setState(137);
				expresion(0);
				}
				break;
			case 2:
				_localctx = new DeclAssignPatternContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(138);
				match(MUT);
				setState(139);
				id_pattern();
				setState(140);
				match(ASSIGN);
				setState(141);
				expresion(0);
				}
				break;
			case 3:
				_localctx = new DeclPatternContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(143);
				match(MUT);
				setState(144);
				id_pattern();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Id_patternContext extends ParserRuleContext {
		public Id_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_id_pattern; }
	 
		public Id_patternContext() { }
		public void copyFrom(Id_patternContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdPatternContext extends Id_patternContext {
		public List<TerminalNode> ID() { return getTokens(VlangParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(VlangParser.ID, i);
		}
		public List<TerminalNode> DOT() { return getTokens(VlangParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(VlangParser.DOT, i);
		}
		public IdPatternContext(Id_patternContext ctx) { copyFrom(ctx); }
	}

	public final Id_patternContext id_pattern() throws RecognitionException {
		Id_patternContext _localctx = new Id_patternContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_id_pattern);
		int _la;
		try {
			_localctx = new IdPatternContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
			match(ID);
			setState(152);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(148);
				match(DOT);
				setState(149);
				match(ID);
				}
				}
				setState(154);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpresionContext extends ParserRuleContext {
		public ExpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion; }
	 
		public ExpresionContext() { }
		public void copyFrom(ExpresionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CorchetesexpreContext extends ExpresionContext {
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public CorchetesexpreContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UnarioContext extends ExpresionContext {
		public Token op;
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode NOT() { return getToken(VlangParser.NOT, 0); }
		public TerminalNode MINUS() { return getToken(VlangParser.MINUS, 0); }
		public UnarioContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParentesisexpreContext extends ExpresionContext {
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public ParentesisexpreContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IncredecrContext extends ExpresionContext {
		public IncredecreContext incredecre() {
			return getRuleContext(IncredecreContext.class,0);
		}
		public IncredecrContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorexpresionContext extends ExpresionContext {
		public ValorContext valor() {
			return getRuleContext(ValorContext.class,0);
		}
		public ValorexpresionContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionforContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public AsignacionforContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public IdContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expdotexp1Context extends ExpresionContext {
		public List<TerminalNode> ID() { return getTokens(VlangParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(VlangParser.ID, i);
		}
		public TerminalNode DOT() { return getToken(VlangParser.DOT, 0); }
		public Expdotexp1Context(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpdotexpContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode DOT() { return getToken(VlangParser.DOT, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ExpdotexpContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BinaryExpContext extends ExpresionContext {
		public ExpresionContext left;
		public Token op;
		public ExpresionContext right;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode MUL() { return getToken(VlangParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(VlangParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(VlangParser.MOD, 0); }
		public TerminalNode PLUS() { return getToken(VlangParser.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(VlangParser.MINUS, 0); }
		public TerminalNode LT() { return getToken(VlangParser.LT, 0); }
		public TerminalNode LE() { return getToken(VlangParser.LE, 0); }
		public TerminalNode GT() { return getToken(VlangParser.GT, 0); }
		public TerminalNode GE() { return getToken(VlangParser.GE, 0); }
		public TerminalNode EQ() { return getToken(VlangParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(VlangParser.NEQ, 0); }
		public TerminalNode AND() { return getToken(VlangParser.AND, 0); }
		public TerminalNode OR() { return getToken(VlangParser.OR, 0); }
		public BinaryExpContext(ExpresionContext ctx) { copyFrom(ctx); }
	}

	public final ExpresionContext expresion() throws RecognitionException {
		return expresion(0);
	}

	private ExpresionContext expresion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpresionContext _localctx = new ExpresionContext(_ctx, _parentState);
		ExpresionContext _prevctx = _localctx;
		int _startState = 24;
		enterRecursionRule(_localctx, 24, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				{
				_localctx = new ValorexpresionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(156);
				valor();
				}
				break;
			case 2:
				{
				_localctx = new ParentesisexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(157);
				match(LPAREN);
				setState(158);
				expresion(0);
				setState(159);
				match(RPAREN);
				}
				break;
			case 3:
				{
				_localctx = new CorchetesexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(161);
				match(LBRACK);
				setState(162);
				expresion(0);
				setState(163);
				match(RBRACK);
				}
				break;
			case 4:
				{
				_localctx = new UnarioContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(165);
				((UnarioContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==MINUS || _la==NOT) ) {
					((UnarioContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(166);
				expresion(12);
				}
				break;
			case 5:
				{
				_localctx = new IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(167);
				match(ID);
				}
				break;
			case 6:
				{
				_localctx = new IncredecrContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(168);
				incredecre();
				}
				break;
			case 7:
				{
				_localctx = new Expdotexp1Context(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(169);
				match(ID);
				setState(170);
				match(DOT);
				setState(171);
				match(ID);
				}
				break;
			case 8:
				{
				_localctx = new ExpdotexpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(172);
				match(ID);
				setState(173);
				match(DOT);
				setState(174);
				expresion(2);
				}
				break;
			case 9:
				{
				_localctx = new AsignacionforContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(175);
				match(ID);
				setState(176);
				match(ASSIGN);
				setState(177);
				expresion(1);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(200);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(198);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
					case 1:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(180);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(181);
						((BinaryExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 117440512L) != 0)) ) {
							((BinaryExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(182);
						((BinaryExpContext)_localctx).right = expresion(12);
						}
						break;
					case 2:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(183);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(184);
						((BinaryExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==PLUS || _la==MINUS) ) {
							((BinaryExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(185);
						((BinaryExpContext)_localctx).right = expresion(11);
						}
						break;
					case 3:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(186);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(187);
						((BinaryExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 64424509440L) != 0)) ) {
							((BinaryExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(188);
						((BinaryExpContext)_localctx).right = expresion(10);
						}
						break;
					case 4:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(189);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(190);
						((BinaryExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==EQ || _la==NEQ) ) {
							((BinaryExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(191);
						((BinaryExpContext)_localctx).right = expresion(9);
						}
						break;
					case 5:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(192);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(193);
						((BinaryExpContext)_localctx).op = match(AND);
						setState(194);
						((BinaryExpContext)_localctx).right = expresion(8);
						}
						break;
					case 6:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(195);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(196);
						((BinaryExpContext)_localctx).op = match(OR);
						setState(197);
						((BinaryExpContext)_localctx).right = expresion(7);
						}
						break;
					}
					} 
				}
				setState(202);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParametrosContext extends ParserRuleContext {
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public ParametrosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametros; }
	}

	public final ParametrosContext parametros() throws RecognitionException {
		ParametrosContext _localctx = new ParametrosContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(203);
			expresion(0);
			setState(208);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(204);
				match(COMMA);
				setState(205);
				expresion(0);
				}
				}
				setState(210);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ValoresContext extends ParserRuleContext {
		public ValorContext valor() {
			return getRuleContext(ValorContext.class,0);
		}
		public ValoresContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valores; }
	}

	public final ValoresContext valores() throws RecognitionException {
		ValoresContext _localctx = new ValoresContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_valores);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(211);
			valor();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ValorContext extends ParserRuleContext {
		public ValorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valor; }
	 
		public ValorContext() { }
		public void copyFrom(ValorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorDecimalContext extends ValorContext {
		public TerminalNode DECIMAL() { return getToken(VlangParser.DECIMAL, 0); }
		public ValorDecimalContext(ValorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorEnteroContext extends ValorContext {
		public TerminalNode ENTERO() { return getToken(VlangParser.ENTERO, 0); }
		public ValorEnteroContext(ValorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorBooleanoContext extends ValorContext {
		public TerminalNode BOOLEANO() { return getToken(VlangParser.BOOLEANO, 0); }
		public ValorBooleanoContext(ValorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorCaracterContext extends ValorContext {
		public TerminalNode CARACTER() { return getToken(VlangParser.CARACTER, 0); }
		public ValorCaracterContext(ValorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorCadenaContext extends ValorContext {
		public TerminalNode CADENA() { return getToken(VlangParser.CADENA, 0); }
		public ValorCadenaContext(ValorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorFloatContext extends ValorContext {
		public TerminalNode FLOAT() { return getToken(VlangParser.FLOAT, 0); }
		public ValorFloatContext(ValorContext ctx) { copyFrom(ctx); }
	}

	public final ValorContext valor() throws RecognitionException {
		ValorContext _localctx = new ValorContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_valor);
		try {
			setState(219);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				_localctx = new ValorEnteroContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(213);
				match(ENTERO);
				}
				break;
			case DECIMAL:
				_localctx = new ValorDecimalContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(214);
				match(DECIMAL);
				}
				break;
			case CADENA:
				_localctx = new ValorCadenaContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(215);
				match(CADENA);
				}
				break;
			case BOOLEANO:
				_localctx = new ValorBooleanoContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(216);
				match(BOOLEANO);
				}
				break;
			case CARACTER:
				_localctx = new ValorCaracterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(217);
				match(CARACTER);
				}
				break;
			case FLOAT:
				_localctx = new ValorFloatContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(218);
				match(FLOAT);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IncredecreContext extends ParserRuleContext {
		public IncredecreContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_incredecre; }
	 
		public IncredecreContext() { }
		public void copyFrom(IncredecreContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IncrementoContext extends IncredecreContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode INC() { return getToken(VlangParser.INC, 0); }
		public IncrementoContext(IncredecreContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DecrementoContext extends IncredecreContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode DEC() { return getToken(VlangParser.DEC, 0); }
		public DecrementoContext(IncredecreContext ctx) { copyFrom(ctx); }
	}

	public final IncredecreContext incredecre() throws RecognitionException {
		IncredecreContext _localctx = new IncredecreContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_incredecre);
		try {
			setState(225);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				_localctx = new IncrementoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(221);
				match(ID);
				setState(222);
				match(INC);
				}
				break;
			case 2:
				_localctx = new DecrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(223);
				match(ID);
				setState(224);
				match(DEC);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 12:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 11);
		case 1:
			return precpred(_ctx, 10);
		case 2:
			return precpred(_ctx, 9);
		case 3:
			return precpred(_ctx, 8);
		case 4:
			return precpred(_ctx, 7);
		case 5:
			return precpred(_ctx, 6);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u00011\u00e4\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0001\u0000\u0004\u0000$\b\u0000\u000b\u0000"+
		"\f\u0000%\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001"+
		"8\b\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0005\u0002=\b\u0002\n\u0002"+
		"\f\u0002@\t\u0002\u0001\u0002\u0003\u0002C\b\u0002\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0005\u0003I\b\u0003\n\u0003\f\u0003L\t\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0005\u0004"+
		"S\b\u0004\n\u0004\f\u0004V\t\u0004\u0001\u0004\u0001\u0004\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005^\b\u0005\n\u0005\f\u0005"+
		"a\t\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0003\u0006j\b\u0006\u0001\u0006\u0001\u0006"+
		"\u0005\u0006n\b\u0006\n\u0006\f\u0006q\t\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\b\u0001\b\u0003\b}\b\b\u0001\b\u0001\b\u0003\b\u0081\b\b\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0003\n\u0092\b\n\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0005\u000b\u0097\b\u000b\n\u000b\f\u000b\u009a"+
		"\t\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003\f\u00b3\b\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0005"+
		"\f\u00c7\b\f\n\f\f\f\u00ca\t\f\u0001\r\u0001\r\u0001\r\u0005\r\u00cf\b"+
		"\r\n\r\f\r\u00d2\t\r\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u00dc\b\u000f\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0003\u0010\u00e2\b\u0010\u0001"+
		"\u0010\u0000\u0001\u0018\u0011\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010"+
		"\u0012\u0014\u0016\u0018\u001a\u001c\u001e \u0000\u0005\u0002\u0000\u0017"+
		"\u0017\u001b\u001b\u0001\u0000\u0018\u001a\u0001\u0000\u0016\u0017\u0001"+
		"\u0000 #\u0001\u0000\u001e\u001f\u00fc\u0000#\u0001\u0000\u0000\u0000"+
		"\u00027\u0001\u0000\u0000\u0000\u00049\u0001\u0000\u0000\u0000\u0006D"+
		"\u0001\u0000\u0000\u0000\bO\u0001\u0000\u0000\u0000\nY\u0001\u0000\u0000"+
		"\u0000\fd\u0001\u0000\u0000\u0000\u000et\u0001\u0000\u0000\u0000\u0010"+
		"\u0080\u0001\u0000\u0000\u0000\u0012\u0082\u0001\u0000\u0000\u0000\u0014"+
		"\u0091\u0001\u0000\u0000\u0000\u0016\u0093\u0001\u0000\u0000\u0000\u0018"+
		"\u00b2\u0001\u0000\u0000\u0000\u001a\u00cb\u0001\u0000\u0000\u0000\u001c"+
		"\u00d3\u0001\u0000\u0000\u0000\u001e\u00db\u0001\u0000\u0000\u0000 \u00e1"+
		"\u0001\u0000\u0000\u0000\"$\u0003\u0002\u0001\u0000#\"\u0001\u0000\u0000"+
		"\u0000$%\u0001\u0000\u0000\u0000%#\u0001\u0000\u0000\u0000%&\u0001\u0000"+
		"\u0000\u0000&\'\u0001\u0000\u0000\u0000\'(\u0005\u0000\u0000\u0001(\u0001"+
		"\u0001\u0000\u0000\u0000)8\u0003\u0014\n\u0000*8\u0003\u0012\t\u0000+"+
		"8\u0003\u0004\u0002\u0000,8\u0003\n\u0005\u0000-8\u0003\f\u0006\u0000"+
		".8\u0003\u0010\b\u0000/0\u0005\u0001\u0000\u000001\u0003\u0018\f\u0000"+
		"12\u0005(\u0000\u000028\u0001\u0000\u0000\u000034\u0005\u0002\u0000\u0000"+
		"45\u0003\u0018\f\u000056\u0005(\u0000\u000068\u0001\u0000\u0000\u0000"+
		"7)\u0001\u0000\u0000\u00007*\u0001\u0000\u0000\u00007+\u0001\u0000\u0000"+
		"\u00007,\u0001\u0000\u0000\u00007-\u0001\u0000\u0000\u00007.\u0001\u0000"+
		"\u0000\u00007/\u0001\u0000\u0000\u000073\u0001\u0000\u0000\u00008\u0003"+
		"\u0001\u0000\u0000\u00009>\u0003\u0006\u0003\u0000:;\u0005\u000f\u0000"+
		"\u0000;=\u0003\u0006\u0003\u0000<:\u0001\u0000\u0000\u0000=@\u0001\u0000"+
		"\u0000\u0000><\u0001\u0000\u0000\u0000>?\u0001\u0000\u0000\u0000?B\u0001"+
		"\u0000\u0000\u0000@>\u0001\u0000\u0000\u0000AC\u0003\b\u0004\u0000BA\u0001"+
		"\u0000\u0000\u0000BC\u0001\u0000\u0000\u0000C\u0005\u0001\u0000\u0000"+
		"\u0000DE\u0005\u000e\u0000\u0000EF\u0003\u0018\f\u0000FJ\u0005*\u0000"+
		"\u0000GI\u0003\u0002\u0001\u0000HG\u0001\u0000\u0000\u0000IL\u0001\u0000"+
		"\u0000\u0000JH\u0001\u0000\u0000\u0000JK\u0001\u0000\u0000\u0000KM\u0001"+
		"\u0000\u0000\u0000LJ\u0001\u0000\u0000\u0000MN\u0005,\u0000\u0000N\u0007"+
		"\u0001\u0000\u0000\u0000OP\u0005\u000f\u0000\u0000PT\u0005*\u0000\u0000"+
		"QS\u0003\u0002\u0001\u0000RQ\u0001\u0000\u0000\u0000SV\u0001\u0000\u0000"+
		"\u0000TR\u0001\u0000\u0000\u0000TU\u0001\u0000\u0000\u0000UW\u0001\u0000"+
		"\u0000\u0000VT\u0001\u0000\u0000\u0000WX\u0005,\u0000\u0000X\t\u0001\u0000"+
		"\u0000\u0000YZ\u0005\u0010\u0000\u0000Z[\u0003\u0018\f\u0000[_\u0005*"+
		"\u0000\u0000\\^\u0003\u0002\u0001\u0000]\\\u0001\u0000\u0000\u0000^a\u0001"+
		"\u0000\u0000\u0000_]\u0001\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000"+
		"`b\u0001\u0000\u0000\u0000a_\u0001\u0000\u0000\u0000bc\u0005,\u0000\u0000"+
		"c\u000b\u0001\u0000\u0000\u0000de\u0005\u0011\u0000\u0000ef\u0005\r\u0000"+
		"\u0000fi\u0005\u0012\u0000\u0000gj\u0003\u0018\f\u0000hj\u0003\u000e\u0007"+
		"\u0000ig\u0001\u0000\u0000\u0000ih\u0001\u0000\u0000\u0000jk\u0001\u0000"+
		"\u0000\u0000ko\u0005*\u0000\u0000ln\u0003\u0002\u0001\u0000ml\u0001\u0000"+
		"\u0000\u0000nq\u0001\u0000\u0000\u0000om\u0001\u0000\u0000\u0000op\u0001"+
		"\u0000\u0000\u0000pr\u0001\u0000\u0000\u0000qo\u0001\u0000\u0000\u0000"+
		"rs\u0005,\u0000\u0000s\r\u0001\u0000\u0000\u0000tu\u0003\u0018\f\u0000"+
		"uv\u0005-\u0000\u0000vw\u0005-\u0000\u0000wx\u0005-\u0000\u0000xy\u0003"+
		"\u0018\f\u0000y\u000f\u0001\u0000\u0000\u0000z|\u0005\u0013\u0000\u0000"+
		"{}\u0003\u0018\f\u0000|{\u0001\u0000\u0000\u0000|}\u0001\u0000\u0000\u0000"+
		"}\u0081\u0001\u0000\u0000\u0000~\u0081\u0005\u0014\u0000\u0000\u007f\u0081"+
		"\u0005\u0015\u0000\u0000\u0080z\u0001\u0000\u0000\u0000\u0080~\u0001\u0000"+
		"\u0000\u0000\u0080\u007f\u0001\u0000\u0000\u0000\u0081\u0011\u0001\u0000"+
		"\u0000\u0000\u0082\u0083\u0003\u0016\u000b\u0000\u0083\u0084\u0005$\u0000"+
		"\u0000\u0084\u0085\u0003\u0018\f\u0000\u0085\u0013\u0001\u0000\u0000\u0000"+
		"\u0086\u0087\u0005\u0006\u0000\u0000\u0087\u0088\u0005\r\u0000\u0000\u0088"+
		"\u0089\u0005$\u0000\u0000\u0089\u0092\u0003\u0018\f\u0000\u008a\u008b"+
		"\u0005\u0006\u0000\u0000\u008b\u008c\u0003\u0016\u000b\u0000\u008c\u008d"+
		"\u0005$\u0000\u0000\u008d\u008e\u0003\u0018\f\u0000\u008e\u0092\u0001"+
		"\u0000\u0000\u0000\u008f\u0090\u0005\u0006\u0000\u0000\u0090\u0092\u0003"+
		"\u0016\u000b\u0000\u0091\u0086\u0001\u0000\u0000\u0000\u0091\u008a\u0001"+
		"\u0000\u0000\u0000\u0091\u008f\u0001\u0000\u0000\u0000\u0092\u0015\u0001"+
		"\u0000\u0000\u0000\u0093\u0098\u0005\r\u0000\u0000\u0094\u0095\u0005-"+
		"\u0000\u0000\u0095\u0097\u0005\r\u0000\u0000\u0096\u0094\u0001\u0000\u0000"+
		"\u0000\u0097\u009a\u0001\u0000\u0000\u0000\u0098\u0096\u0001\u0000\u0000"+
		"\u0000\u0098\u0099\u0001\u0000\u0000\u0000\u0099\u0017\u0001\u0000\u0000"+
		"\u0000\u009a\u0098\u0001\u0000\u0000\u0000\u009b\u009c\u0006\f\uffff\uffff"+
		"\u0000\u009c\u00b3\u0003\u001e\u000f\u0000\u009d\u009e\u0005\'\u0000\u0000"+
		"\u009e\u009f\u0003\u0018\f\u0000\u009f\u00a0\u0005(\u0000\u0000\u00a0"+
		"\u00b3\u0001\u0000\u0000\u0000\u00a1\u00a2\u0005)\u0000\u0000\u00a2\u00a3"+
		"\u0003\u0018\f\u0000\u00a3\u00a4\u0005+\u0000\u0000\u00a4\u00b3\u0001"+
		"\u0000\u0000\u0000\u00a5\u00a6\u0007\u0000\u0000\u0000\u00a6\u00b3\u0003"+
		"\u0018\f\f\u00a7\u00b3\u0005\r\u0000\u0000\u00a8\u00b3\u0003 \u0010\u0000"+
		"\u00a9\u00aa\u0005\r\u0000\u0000\u00aa\u00ab\u0005-\u0000\u0000\u00ab"+
		"\u00b3\u0005\r\u0000\u0000\u00ac\u00ad\u0005\r\u0000\u0000\u00ad\u00ae"+
		"\u0005-\u0000\u0000\u00ae\u00b3\u0003\u0018\f\u0002\u00af\u00b0\u0005"+
		"\r\u0000\u0000\u00b0\u00b1\u0005$\u0000\u0000\u00b1\u00b3\u0003\u0018"+
		"\f\u0001\u00b2\u009b\u0001\u0000\u0000\u0000\u00b2\u009d\u0001\u0000\u0000"+
		"\u0000\u00b2\u00a1\u0001\u0000\u0000\u0000\u00b2\u00a5\u0001\u0000\u0000"+
		"\u0000\u00b2\u00a7\u0001\u0000\u0000\u0000\u00b2\u00a8\u0001\u0000\u0000"+
		"\u0000\u00b2\u00a9\u0001\u0000\u0000\u0000\u00b2\u00ac\u0001\u0000\u0000"+
		"\u0000\u00b2\u00af\u0001\u0000\u0000\u0000\u00b3\u00c8\u0001\u0000\u0000"+
		"\u0000\u00b4\u00b5\n\u000b\u0000\u0000\u00b5\u00b6\u0007\u0001\u0000\u0000"+
		"\u00b6\u00c7\u0003\u0018\f\f\u00b7\u00b8\n\n\u0000\u0000\u00b8\u00b9\u0007"+
		"\u0002\u0000\u0000\u00b9\u00c7\u0003\u0018\f\u000b\u00ba\u00bb\n\t\u0000"+
		"\u0000\u00bb\u00bc\u0007\u0003\u0000\u0000\u00bc\u00c7\u0003\u0018\f\n"+
		"\u00bd\u00be\n\b\u0000\u0000\u00be\u00bf\u0007\u0004\u0000\u0000\u00bf"+
		"\u00c7\u0003\u0018\f\t\u00c0\u00c1\n\u0007\u0000\u0000\u00c1\u00c2\u0005"+
		"\u001d\u0000\u0000\u00c2\u00c7\u0003\u0018\f\b\u00c3\u00c4\n\u0006\u0000"+
		"\u0000\u00c4\u00c5\u0005\u001c\u0000\u0000\u00c5\u00c7\u0003\u0018\f\u0007"+
		"\u00c6\u00b4\u0001\u0000\u0000\u0000\u00c6\u00b7\u0001\u0000\u0000\u0000"+
		"\u00c6\u00ba\u0001\u0000\u0000\u0000\u00c6\u00bd\u0001\u0000\u0000\u0000"+
		"\u00c6\u00c0\u0001\u0000\u0000\u0000\u00c6\u00c3\u0001\u0000\u0000\u0000"+
		"\u00c7\u00ca\u0001\u0000\u0000\u0000\u00c8\u00c6\u0001\u0000\u0000\u0000"+
		"\u00c8\u00c9\u0001\u0000\u0000\u0000\u00c9\u0019\u0001\u0000\u0000\u0000"+
		"\u00ca\u00c8\u0001\u0000\u0000\u0000\u00cb\u00d0\u0003\u0018\f\u0000\u00cc"+
		"\u00cd\u0005.\u0000\u0000\u00cd\u00cf\u0003\u0018\f\u0000\u00ce\u00cc"+
		"\u0001\u0000\u0000\u0000\u00cf\u00d2\u0001\u0000\u0000\u0000\u00d0\u00ce"+
		"\u0001\u0000\u0000\u0000\u00d0\u00d1\u0001\u0000\u0000\u0000\u00d1\u001b"+
		"\u0001\u0000\u0000\u0000\u00d2\u00d0\u0001\u0000\u0000\u0000\u00d3\u00d4"+
		"\u0003\u001e\u000f\u0000\u00d4\u001d\u0001\u0000\u0000\u0000\u00d5\u00dc"+
		"\u0005\b\u0000\u0000\u00d6\u00dc\u0005\n\u0000\u0000\u00d7\u00dc\u0005"+
		"\u000b\u0000\u0000\u00d8\u00dc\u0005\u0007\u0000\u0000\u00d9\u00dc\u0005"+
		"\f\u0000\u0000\u00da\u00dc\u0005\t\u0000\u0000\u00db\u00d5\u0001\u0000"+
		"\u0000\u0000\u00db\u00d6\u0001\u0000\u0000\u0000\u00db\u00d7\u0001\u0000"+
		"\u0000\u0000\u00db\u00d8\u0001\u0000\u0000\u0000\u00db\u00d9\u0001\u0000"+
		"\u0000\u0000\u00db\u00da\u0001\u0000\u0000\u0000\u00dc\u001f\u0001\u0000"+
		"\u0000\u0000\u00dd\u00de\u0005\r\u0000\u0000\u00de\u00e2\u0005%\u0000"+
		"\u0000\u00df\u00e0\u0005\r\u0000\u0000\u00e0\u00e2\u0005&\u0000\u0000"+
		"\u00e1\u00dd\u0001\u0000\u0000\u0000\u00e1\u00df\u0001\u0000\u0000\u0000"+
		"\u00e2!\u0001\u0000\u0000\u0000\u0013%7>BJT_io|\u0080\u0091\u0098\u00b2"+
		"\u00c6\u00c8\u00d0\u00db\u00e1";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}