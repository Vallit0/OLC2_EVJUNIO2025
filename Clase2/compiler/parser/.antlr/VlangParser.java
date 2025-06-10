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
		T__0=1, LEN=2, CAP=3, APPEND=4, MUT=5, BOOLEANO=6, ENTERO=7, FLOAT=8, 
		DECIMAL=9, CADENA=10, CARACTER=11, ID=12, IF_KW=13, ELSE_KW=14, WHILE_KW=15, 
		FOR_KW=16, IN_KW=17, RETURN_KW=18, BREAK_KW=19, CONTINUE_KW=20, PLUS=21, 
		MINUS=22, MUL=23, DIV=24, MOD=25, NOT=26, OR=27, AND=28, EQ=29, NEQ=30, 
		LT=31, LE=32, GT=33, GE=34, ASSIGN=35, INC=36, DEC=37, LPAREN=38, RPAREN=39, 
		LBRACK=40, LCOR=41, RBRACK=42, RCOR=43, DOT=44, COMMA=45, WS=46, LINE_COMMENT=47, 
		BLOCK_COMMENT=48;
	public static final int
		RULE_programa = 0, RULE_stmt = 1, RULE_if_stmt = 2, RULE_if_chain = 3, 
		RULE_else_stmt = 4, RULE_while_stmt = 5, RULE_for_stmt = 6, RULE_range = 7, 
		RULE_transfer_stmt = 8, RULE_assign_stmt = 9, RULE_decl_stmt = 10, RULE_var_type = 11, 
		RULE_id_pattern = 12, RULE_expresion = 13, RULE_parametros = 14, RULE_valores = 15, 
		RULE_valor = 16, RULE_incredecre = 17;
	private static String[] makeRuleNames() {
		return new String[] {
			"programa", "stmt", "if_stmt", "if_chain", "else_stmt", "while_stmt", 
			"for_stmt", "range", "transfer_stmt", "assign_stmt", "decl_stmt", "var_type", 
			"id_pattern", "expresion", "parametros", "valores", "valor", "incredecre"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'println('", "'len'", "'cap'", "'append'", "'mut'", null, null, 
			null, null, null, null, null, "'if'", "'else'", "'while'", "'for'", "'in'", 
			"'return'", "'break'", "'continue'", "'+'", "'-'", "'*'", "'/'", "'%'", 
			"'!'", "'||'", "'&&'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", 
			"'='", "'++'", "'--'", "'('", "')'", "'['", "'{'", "']'", "'}'", "'.'", 
			"','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, "LEN", "CAP", "APPEND", "MUT", "BOOLEANO", "ENTERO", "FLOAT", 
			"DECIMAL", "CADENA", "CARACTER", "ID", "IF_KW", "ELSE_KW", "WHILE_KW", 
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
			setState(37); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(36);
				stmt();
				}
				}
				setState(39); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 1945634L) != 0) );
			setState(41);
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
			setState(53);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MUT:
				enterOuterAlt(_localctx, 1);
				{
				setState(43);
				decl_stmt();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 2);
				{
				setState(44);
				assign_stmt();
				}
				break;
			case IF_KW:
				enterOuterAlt(_localctx, 3);
				{
				setState(45);
				if_stmt();
				}
				break;
			case WHILE_KW:
				enterOuterAlt(_localctx, 4);
				{
				setState(46);
				while_stmt();
				}
				break;
			case FOR_KW:
				enterOuterAlt(_localctx, 5);
				{
				setState(47);
				for_stmt();
				}
				break;
			case RETURN_KW:
			case BREAK_KW:
			case CONTINUE_KW:
				enterOuterAlt(_localctx, 6);
				{
				setState(48);
				transfer_stmt();
				}
				break;
			case T__0:
				enterOuterAlt(_localctx, 7);
				{
				setState(49);
				match(T__0);
				setState(50);
				expresion(0);
				setState(51);
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
			setState(55);
			if_chain();
			setState(60);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(56);
					match(ELSE_KW);
					setState(57);
					if_chain();
					}
					} 
				}
				setState(62);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			setState(64);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE_KW) {
				{
				setState(63);
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
			setState(66);
			match(IF_KW);
			setState(67);
			expresion(0);
			setState(68);
			match(LCOR);
			setState(72);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1945634L) != 0)) {
				{
				{
				setState(69);
				stmt();
				}
				}
				setState(74);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(75);
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
			setState(77);
			match(ELSE_KW);
			setState(78);
			match(LCOR);
			setState(82);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1945634L) != 0)) {
				{
				{
				setState(79);
				stmt();
				}
				}
				setState(84);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(85);
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
			setState(87);
			match(WHILE_KW);
			setState(88);
			expresion(0);
			setState(89);
			match(LCOR);
			setState(93);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1945634L) != 0)) {
				{
				{
				setState(90);
				stmt();
				}
				}
				setState(95);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(96);
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
			setState(98);
			match(FOR_KW);
			setState(99);
			match(ID);
			setState(100);
			match(IN_KW);
			setState(103);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				{
				setState(101);
				expresion(0);
				}
				break;
			case 2:
				{
				setState(102);
				range();
				}
				break;
			}
			setState(105);
			match(LCOR);
			setState(109);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1945634L) != 0)) {
				{
				{
				setState(106);
				stmt();
				}
				}
				setState(111);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(112);
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
			setState(114);
			expresion(0);
			setState(115);
			match(DOT);
			setState(116);
			match(DOT);
			setState(117);
			match(DOT);
			setState(118);
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
			setState(126);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RETURN_KW:
				_localctx = new ReturnStmtContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(120);
				match(RETURN_KW);
				setState(122);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
				case 1:
					{
					setState(121);
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
				setState(124);
				match(BREAK_KW);
				}
				break;
			case CONTINUE_KW:
				_localctx = new ContinueStmtContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(125);
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
			setState(128);
			id_pattern();
			setState(129);
			match(ASSIGN);
			setState(130);
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
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public DeclAssignContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclPatternContext extends Decl_stmtContext {
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public DeclPatternContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclAssignPatternContext extends Decl_stmtContext {
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
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
				setState(132);
				var_type();
				setState(133);
				match(ID);
				setState(134);
				match(ASSIGN);
				setState(135);
				expresion(0);
				}
				break;
			case 2:
				_localctx = new DeclAssignPatternContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(137);
				var_type();
				setState(138);
				id_pattern();
				setState(139);
				match(ASSIGN);
				setState(140);
				expresion(0);
				}
				break;
			case 3:
				_localctx = new DeclPatternContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(142);
				var_type();
				setState(143);
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
	public static class Var_typeContext extends ParserRuleContext {
		public Var_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_type; }
	 
		public Var_typeContext() { }
		public void copyFrom(Var_typeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MutTypeContext extends Var_typeContext {
		public TerminalNode MUT() { return getToken(VlangParser.MUT, 0); }
		public MutTypeContext(Var_typeContext ctx) { copyFrom(ctx); }
	}

	public final Var_typeContext var_type() throws RecognitionException {
		Var_typeContext _localctx = new Var_typeContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_var_type);
		try {
			_localctx = new MutTypeContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
			match(MUT);
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
		enterRule(_localctx, 24, RULE_id_pattern);
		int _la;
		try {
			_localctx = new IdPatternContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(149);
			match(ID);
			setState(154);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(150);
				match(DOT);
				setState(151);
				match(ID);
				}
				}
				setState(156);
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
		int _startState = 26;
		enterRecursionRule(_localctx, 26, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(180);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				{
				_localctx = new ValorexpresionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(158);
				valor();
				}
				break;
			case 2:
				{
				_localctx = new ParentesisexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(159);
				match(LPAREN);
				setState(160);
				expresion(0);
				setState(161);
				match(RPAREN);
				}
				break;
			case 3:
				{
				_localctx = new CorchetesexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(163);
				match(LBRACK);
				setState(164);
				expresion(0);
				setState(165);
				match(RBRACK);
				}
				break;
			case 4:
				{
				_localctx = new UnarioContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(167);
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
				setState(168);
				expresion(12);
				}
				break;
			case 5:
				{
				_localctx = new IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(169);
				match(ID);
				}
				break;
			case 6:
				{
				_localctx = new IncredecrContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(170);
				incredecre();
				}
				break;
			case 7:
				{
				_localctx = new Expdotexp1Context(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(171);
				match(ID);
				setState(172);
				match(DOT);
				setState(173);
				match(ID);
				}
				break;
			case 8:
				{
				_localctx = new ExpdotexpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(174);
				match(ID);
				setState(175);
				match(DOT);
				setState(176);
				expresion(2);
				}
				break;
			case 9:
				{
				_localctx = new AsignacionforContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(177);
				match(ID);
				setState(178);
				match(ASSIGN);
				setState(179);
				expresion(1);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(202);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(200);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
					case 1:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(182);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(183);
						((BinaryExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 58720256L) != 0)) ) {
							((BinaryExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(184);
						((BinaryExpContext)_localctx).right = expresion(12);
						}
						break;
					case 2:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(185);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(186);
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
						setState(187);
						((BinaryExpContext)_localctx).right = expresion(11);
						}
						break;
					case 3:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(188);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(189);
						((BinaryExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 32212254720L) != 0)) ) {
							((BinaryExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(190);
						((BinaryExpContext)_localctx).right = expresion(10);
						}
						break;
					case 4:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(191);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(192);
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
						setState(193);
						((BinaryExpContext)_localctx).right = expresion(9);
						}
						break;
					case 5:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(194);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(195);
						((BinaryExpContext)_localctx).op = match(AND);
						setState(196);
						((BinaryExpContext)_localctx).right = expresion(8);
						}
						break;
					case 6:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(197);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(198);
						((BinaryExpContext)_localctx).op = match(OR);
						setState(199);
						((BinaryExpContext)_localctx).right = expresion(7);
						}
						break;
					}
					} 
				}
				setState(204);
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
		enterRule(_localctx, 28, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(205);
			expresion(0);
			setState(210);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(206);
				match(COMMA);
				setState(207);
				expresion(0);
				}
				}
				setState(212);
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
		enterRule(_localctx, 30, RULE_valores);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(213);
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
		enterRule(_localctx, 32, RULE_valor);
		try {
			setState(221);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				_localctx = new ValorEnteroContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(215);
				match(ENTERO);
				}
				break;
			case DECIMAL:
				_localctx = new ValorDecimalContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(216);
				match(DECIMAL);
				}
				break;
			case CADENA:
				_localctx = new ValorCadenaContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(217);
				match(CADENA);
				}
				break;
			case BOOLEANO:
				_localctx = new ValorBooleanoContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(218);
				match(BOOLEANO);
				}
				break;
			case CARACTER:
				_localctx = new ValorCaracterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(219);
				match(CARACTER);
				}
				break;
			case FLOAT:
				_localctx = new ValorFloatContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(220);
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
		enterRule(_localctx, 34, RULE_incredecre);
		try {
			setState(227);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				_localctx = new IncrementoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(223);
				match(ID);
				setState(224);
				match(INC);
				}
				break;
			case 2:
				_localctx = new DecrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(225);
				match(ID);
				setState(226);
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
		case 13:
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
		"\u0004\u00010\u00e6\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0001\u0000\u0004\u0000"+
		"&\b\u0000\u000b\u0000\f\u0000\'\u0001\u0000\u0001\u0000\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0003\u00016\b\u0001\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0005\u0002;\b\u0002\n\u0002\f\u0002>\t\u0002\u0001"+
		"\u0002\u0003\u0002A\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0005\u0003G\b\u0003\n\u0003\f\u0003J\t\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0005\u0004Q\b\u0004\n\u0004"+
		"\f\u0004T\t\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0005\u0005\\\b\u0005\n\u0005\f\u0005_\t\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0003\u0006h\b\u0006\u0001\u0006\u0001\u0006\u0005\u0006l\b\u0006"+
		"\n\u0006\f\u0006o\t\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0003"+
		"\b{\b\b\u0001\b\u0001\b\u0003\b\u007f\b\b\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0003\n\u0092\b\n\u0001\u000b\u0001"+
		"\u000b\u0001\f\u0001\f\u0001\f\u0005\f\u0099\b\f\n\f\f\f\u009c\t\f\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0003\r\u00b5\b\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0005\r\u00c9\b\r\n"+
		"\r\f\r\u00cc\t\r\u0001\u000e\u0001\u000e\u0001\u000e\u0005\u000e\u00d1"+
		"\b\u000e\n\u000e\f\u000e\u00d4\t\u000e\u0001\u000f\u0001\u000f\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0003\u0010"+
		"\u00de\b\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0003\u0011"+
		"\u00e4\b\u0011\u0001\u0011\u0000\u0001\u001a\u0012\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"\u0000"+
		"\u0005\u0002\u0000\u0016\u0016\u001a\u001a\u0001\u0000\u0017\u0019\u0001"+
		"\u0000\u0015\u0016\u0001\u0000\u001f\"\u0001\u0000\u001d\u001e\u00fc\u0000"+
		"%\u0001\u0000\u0000\u0000\u00025\u0001\u0000\u0000\u0000\u00047\u0001"+
		"\u0000\u0000\u0000\u0006B\u0001\u0000\u0000\u0000\bM\u0001\u0000\u0000"+
		"\u0000\nW\u0001\u0000\u0000\u0000\fb\u0001\u0000\u0000\u0000\u000er\u0001"+
		"\u0000\u0000\u0000\u0010~\u0001\u0000\u0000\u0000\u0012\u0080\u0001\u0000"+
		"\u0000\u0000\u0014\u0091\u0001\u0000\u0000\u0000\u0016\u0093\u0001\u0000"+
		"\u0000\u0000\u0018\u0095\u0001\u0000\u0000\u0000\u001a\u00b4\u0001\u0000"+
		"\u0000\u0000\u001c\u00cd\u0001\u0000\u0000\u0000\u001e\u00d5\u0001\u0000"+
		"\u0000\u0000 \u00dd\u0001\u0000\u0000\u0000\"\u00e3\u0001\u0000\u0000"+
		"\u0000$&\u0003\u0002\u0001\u0000%$\u0001\u0000\u0000\u0000&\'\u0001\u0000"+
		"\u0000\u0000\'%\u0001\u0000\u0000\u0000\'(\u0001\u0000\u0000\u0000()\u0001"+
		"\u0000\u0000\u0000)*\u0005\u0000\u0000\u0001*\u0001\u0001\u0000\u0000"+
		"\u0000+6\u0003\u0014\n\u0000,6\u0003\u0012\t\u0000-6\u0003\u0004\u0002"+
		"\u0000.6\u0003\n\u0005\u0000/6\u0003\f\u0006\u000006\u0003\u0010\b\u0000"+
		"12\u0005\u0001\u0000\u000023\u0003\u001a\r\u000034\u0005\'\u0000\u0000"+
		"46\u0001\u0000\u0000\u00005+\u0001\u0000\u0000\u00005,\u0001\u0000\u0000"+
		"\u00005-\u0001\u0000\u0000\u00005.\u0001\u0000\u0000\u00005/\u0001\u0000"+
		"\u0000\u000050\u0001\u0000\u0000\u000051\u0001\u0000\u0000\u00006\u0003"+
		"\u0001\u0000\u0000\u00007<\u0003\u0006\u0003\u000089\u0005\u000e\u0000"+
		"\u00009;\u0003\u0006\u0003\u0000:8\u0001\u0000\u0000\u0000;>\u0001\u0000"+
		"\u0000\u0000<:\u0001\u0000\u0000\u0000<=\u0001\u0000\u0000\u0000=@\u0001"+
		"\u0000\u0000\u0000><\u0001\u0000\u0000\u0000?A\u0003\b\u0004\u0000@?\u0001"+
		"\u0000\u0000\u0000@A\u0001\u0000\u0000\u0000A\u0005\u0001\u0000\u0000"+
		"\u0000BC\u0005\r\u0000\u0000CD\u0003\u001a\r\u0000DH\u0005)\u0000\u0000"+
		"EG\u0003\u0002\u0001\u0000FE\u0001\u0000\u0000\u0000GJ\u0001\u0000\u0000"+
		"\u0000HF\u0001\u0000\u0000\u0000HI\u0001\u0000\u0000\u0000IK\u0001\u0000"+
		"\u0000\u0000JH\u0001\u0000\u0000\u0000KL\u0005+\u0000\u0000L\u0007\u0001"+
		"\u0000\u0000\u0000MN\u0005\u000e\u0000\u0000NR\u0005)\u0000\u0000OQ\u0003"+
		"\u0002\u0001\u0000PO\u0001\u0000\u0000\u0000QT\u0001\u0000\u0000\u0000"+
		"RP\u0001\u0000\u0000\u0000RS\u0001\u0000\u0000\u0000SU\u0001\u0000\u0000"+
		"\u0000TR\u0001\u0000\u0000\u0000UV\u0005+\u0000\u0000V\t\u0001\u0000\u0000"+
		"\u0000WX\u0005\u000f\u0000\u0000XY\u0003\u001a\r\u0000Y]\u0005)\u0000"+
		"\u0000Z\\\u0003\u0002\u0001\u0000[Z\u0001\u0000\u0000\u0000\\_\u0001\u0000"+
		"\u0000\u0000][\u0001\u0000\u0000\u0000]^\u0001\u0000\u0000\u0000^`\u0001"+
		"\u0000\u0000\u0000_]\u0001\u0000\u0000\u0000`a\u0005+\u0000\u0000a\u000b"+
		"\u0001\u0000\u0000\u0000bc\u0005\u0010\u0000\u0000cd\u0005\f\u0000\u0000"+
		"dg\u0005\u0011\u0000\u0000eh\u0003\u001a\r\u0000fh\u0003\u000e\u0007\u0000"+
		"ge\u0001\u0000\u0000\u0000gf\u0001\u0000\u0000\u0000hi\u0001\u0000\u0000"+
		"\u0000im\u0005)\u0000\u0000jl\u0003\u0002\u0001\u0000kj\u0001\u0000\u0000"+
		"\u0000lo\u0001\u0000\u0000\u0000mk\u0001\u0000\u0000\u0000mn\u0001\u0000"+
		"\u0000\u0000np\u0001\u0000\u0000\u0000om\u0001\u0000\u0000\u0000pq\u0005"+
		"+\u0000\u0000q\r\u0001\u0000\u0000\u0000rs\u0003\u001a\r\u0000st\u0005"+
		",\u0000\u0000tu\u0005,\u0000\u0000uv\u0005,\u0000\u0000vw\u0003\u001a"+
		"\r\u0000w\u000f\u0001\u0000\u0000\u0000xz\u0005\u0012\u0000\u0000y{\u0003"+
		"\u001a\r\u0000zy\u0001\u0000\u0000\u0000z{\u0001\u0000\u0000\u0000{\u007f"+
		"\u0001\u0000\u0000\u0000|\u007f\u0005\u0013\u0000\u0000}\u007f\u0005\u0014"+
		"\u0000\u0000~x\u0001\u0000\u0000\u0000~|\u0001\u0000\u0000\u0000~}\u0001"+
		"\u0000\u0000\u0000\u007f\u0011\u0001\u0000\u0000\u0000\u0080\u0081\u0003"+
		"\u0018\f\u0000\u0081\u0082\u0005#\u0000\u0000\u0082\u0083\u0003\u001a"+
		"\r\u0000\u0083\u0013\u0001\u0000\u0000\u0000\u0084\u0085\u0003\u0016\u000b"+
		"\u0000\u0085\u0086\u0005\f\u0000\u0000\u0086\u0087\u0005#\u0000\u0000"+
		"\u0087\u0088\u0003\u001a\r\u0000\u0088\u0092\u0001\u0000\u0000\u0000\u0089"+
		"\u008a\u0003\u0016\u000b\u0000\u008a\u008b\u0003\u0018\f\u0000\u008b\u008c"+
		"\u0005#\u0000\u0000\u008c\u008d\u0003\u001a\r\u0000\u008d\u0092\u0001"+
		"\u0000\u0000\u0000\u008e\u008f\u0003\u0016\u000b\u0000\u008f\u0090\u0003"+
		"\u0018\f\u0000\u0090\u0092\u0001\u0000\u0000\u0000\u0091\u0084\u0001\u0000"+
		"\u0000\u0000\u0091\u0089\u0001\u0000\u0000\u0000\u0091\u008e\u0001\u0000"+
		"\u0000\u0000\u0092\u0015\u0001\u0000\u0000\u0000\u0093\u0094\u0005\u0005"+
		"\u0000\u0000\u0094\u0017\u0001\u0000\u0000\u0000\u0095\u009a\u0005\f\u0000"+
		"\u0000\u0096\u0097\u0005,\u0000\u0000\u0097\u0099\u0005\f\u0000\u0000"+
		"\u0098\u0096\u0001\u0000\u0000\u0000\u0099\u009c\u0001\u0000\u0000\u0000"+
		"\u009a\u0098\u0001\u0000\u0000\u0000\u009a\u009b\u0001\u0000\u0000\u0000"+
		"\u009b\u0019\u0001\u0000\u0000\u0000\u009c\u009a\u0001\u0000\u0000\u0000"+
		"\u009d\u009e\u0006\r\uffff\uffff\u0000\u009e\u00b5\u0003 \u0010\u0000"+
		"\u009f\u00a0\u0005&\u0000\u0000\u00a0\u00a1\u0003\u001a\r\u0000\u00a1"+
		"\u00a2\u0005\'\u0000\u0000\u00a2\u00b5\u0001\u0000\u0000\u0000\u00a3\u00a4"+
		"\u0005(\u0000\u0000\u00a4\u00a5\u0003\u001a\r\u0000\u00a5\u00a6\u0005"+
		"*\u0000\u0000\u00a6\u00b5\u0001\u0000\u0000\u0000\u00a7\u00a8\u0007\u0000"+
		"\u0000\u0000\u00a8\u00b5\u0003\u001a\r\f\u00a9\u00b5\u0005\f\u0000\u0000"+
		"\u00aa\u00b5\u0003\"\u0011\u0000\u00ab\u00ac\u0005\f\u0000\u0000\u00ac"+
		"\u00ad\u0005,\u0000\u0000\u00ad\u00b5\u0005\f\u0000\u0000\u00ae\u00af"+
		"\u0005\f\u0000\u0000\u00af\u00b0\u0005,\u0000\u0000\u00b0\u00b5\u0003"+
		"\u001a\r\u0002\u00b1\u00b2\u0005\f\u0000\u0000\u00b2\u00b3\u0005#\u0000"+
		"\u0000\u00b3\u00b5\u0003\u001a\r\u0001\u00b4\u009d\u0001\u0000\u0000\u0000"+
		"\u00b4\u009f\u0001\u0000\u0000\u0000\u00b4\u00a3\u0001\u0000\u0000\u0000"+
		"\u00b4\u00a7\u0001\u0000\u0000\u0000\u00b4\u00a9\u0001\u0000\u0000\u0000"+
		"\u00b4\u00aa\u0001\u0000\u0000\u0000\u00b4\u00ab\u0001\u0000\u0000\u0000"+
		"\u00b4\u00ae\u0001\u0000\u0000\u0000\u00b4\u00b1\u0001\u0000\u0000\u0000"+
		"\u00b5\u00ca\u0001\u0000\u0000\u0000\u00b6\u00b7\n\u000b\u0000\u0000\u00b7"+
		"\u00b8\u0007\u0001\u0000\u0000\u00b8\u00c9\u0003\u001a\r\f\u00b9\u00ba"+
		"\n\n\u0000\u0000\u00ba\u00bb\u0007\u0002\u0000\u0000\u00bb\u00c9\u0003"+
		"\u001a\r\u000b\u00bc\u00bd\n\t\u0000\u0000\u00bd\u00be\u0007\u0003\u0000"+
		"\u0000\u00be\u00c9\u0003\u001a\r\n\u00bf\u00c0\n\b\u0000\u0000\u00c0\u00c1"+
		"\u0007\u0004\u0000\u0000\u00c1\u00c9\u0003\u001a\r\t\u00c2\u00c3\n\u0007"+
		"\u0000\u0000\u00c3\u00c4\u0005\u001c\u0000\u0000\u00c4\u00c9\u0003\u001a"+
		"\r\b\u00c5\u00c6\n\u0006\u0000\u0000\u00c6\u00c7\u0005\u001b\u0000\u0000"+
		"\u00c7\u00c9\u0003\u001a\r\u0007\u00c8\u00b6\u0001\u0000\u0000\u0000\u00c8"+
		"\u00b9\u0001\u0000\u0000\u0000\u00c8\u00bc\u0001\u0000\u0000\u0000\u00c8"+
		"\u00bf\u0001\u0000\u0000\u0000\u00c8\u00c2\u0001\u0000\u0000\u0000\u00c8"+
		"\u00c5\u0001\u0000\u0000\u0000\u00c9\u00cc\u0001\u0000\u0000\u0000\u00ca"+
		"\u00c8\u0001\u0000\u0000\u0000\u00ca\u00cb\u0001\u0000\u0000\u0000\u00cb"+
		"\u001b\u0001\u0000\u0000\u0000\u00cc\u00ca\u0001\u0000\u0000\u0000\u00cd"+
		"\u00d2\u0003\u001a\r\u0000\u00ce\u00cf\u0005-\u0000\u0000\u00cf\u00d1"+
		"\u0003\u001a\r\u0000\u00d0\u00ce\u0001\u0000\u0000\u0000\u00d1\u00d4\u0001"+
		"\u0000\u0000\u0000\u00d2\u00d0\u0001\u0000\u0000\u0000\u00d2\u00d3\u0001"+
		"\u0000\u0000\u0000\u00d3\u001d\u0001\u0000\u0000\u0000\u00d4\u00d2\u0001"+
		"\u0000\u0000\u0000\u00d5\u00d6\u0003 \u0010\u0000\u00d6\u001f\u0001\u0000"+
		"\u0000\u0000\u00d7\u00de\u0005\u0007\u0000\u0000\u00d8\u00de\u0005\t\u0000"+
		"\u0000\u00d9\u00de\u0005\n\u0000\u0000\u00da\u00de\u0005\u0006\u0000\u0000"+
		"\u00db\u00de\u0005\u000b\u0000\u0000\u00dc\u00de\u0005\b\u0000\u0000\u00dd"+
		"\u00d7\u0001\u0000\u0000\u0000\u00dd\u00d8\u0001\u0000\u0000\u0000\u00dd"+
		"\u00d9\u0001\u0000\u0000\u0000\u00dd\u00da\u0001\u0000\u0000\u0000\u00dd"+
		"\u00db\u0001\u0000\u0000\u0000\u00dd\u00dc\u0001\u0000\u0000\u0000\u00de"+
		"!\u0001\u0000\u0000\u0000\u00df\u00e0\u0005\f\u0000\u0000\u00e0\u00e4"+
		"\u0005$\u0000\u0000\u00e1\u00e2\u0005\f\u0000\u0000\u00e2\u00e4\u0005"+
		"%\u0000\u0000\u00e3\u00df\u0001\u0000\u0000\u0000\u00e3\u00e1\u0001\u0000"+
		"\u0000\u0000\u00e4#\u0001\u0000\u0000\u0000\u0013\'5<@HR]gmz~\u0091\u009a"+
		"\u00b4\u00c8\u00ca\u00d2\u00dd\u00e3";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}