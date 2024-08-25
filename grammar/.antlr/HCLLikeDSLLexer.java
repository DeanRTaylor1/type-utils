// Generated from /Users/deantaylor/Desktop/coding/deanrtaylor1/2024/type-utils/grammar/HCLLikeDSL.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class HCLLikeDSLLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, STRING=9, 
		NUMBER=10, BOOLEAN=11, IDENTIFIER=12, WS=13, COMMENT=14;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "STRING", 
			"NUMBER", "BOOLEAN", "IDENTIFIER", "WS", "COMMENT"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'import'", "'{'", "'}'", "'='", "'repeated'", "'['", "','", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, "STRING", "NUMBER", 
			"BOOLEAN", "IDENTIFIER", "WS", "COMMENT"
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


	public HCLLikeDSLLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "HCLLikeDSL.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\u000eq\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0001\u0000\u0001\u0000\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001"+
		"\u0001\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001"+
		"\u0007\u0001\u0007\u0001\b\u0001\b\u0005\b<\b\b\n\b\f\b?\t\b\u0001\b\u0001"+
		"\b\u0001\t\u0004\tD\b\t\u000b\t\f\tE\u0001\t\u0001\t\u0004\tJ\b\t\u000b"+
		"\t\f\tK\u0003\tN\b\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0003\nY\b\n\u0001\u000b\u0001\u000b\u0005\u000b]\b"+
		"\u000b\n\u000b\f\u000b`\t\u000b\u0001\f\u0004\fc\b\f\u000b\f\f\fd\u0001"+
		"\f\u0001\f\u0001\r\u0001\r\u0005\rk\b\r\n\r\f\rn\t\r\u0001\r\u0001\r\u0000"+
		"\u0000\u000e\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b"+
		"\u0006\r\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b\u0017\f\u0019\r\u001b"+
		"\u000e\u0001\u0000\u0006\u0003\u0000\n\n\r\r\"\"\u0001\u000009\u0003\u0000"+
		"AZ__az\u0004\u000009AZ__az\u0003\u0000\t\n\r\r  \u0002\u0000\n\n\r\rx"+
		"\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000"+
		"\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000"+
		"\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000"+
		"\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011"+
		"\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015"+
		"\u0001\u0000\u0000\u0000\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019"+
		"\u0001\u0000\u0000\u0000\u0000\u001b\u0001\u0000\u0000\u0000\u0001\u001d"+
		"\u0001\u0000\u0000\u0000\u0003$\u0001\u0000\u0000\u0000\u0005&\u0001\u0000"+
		"\u0000\u0000\u0007(\u0001\u0000\u0000\u0000\t*\u0001\u0000\u0000\u0000"+
		"\u000b3\u0001\u0000\u0000\u0000\r5\u0001\u0000\u0000\u0000\u000f7\u0001"+
		"\u0000\u0000\u0000\u00119\u0001\u0000\u0000\u0000\u0013C\u0001\u0000\u0000"+
		"\u0000\u0015X\u0001\u0000\u0000\u0000\u0017Z\u0001\u0000\u0000\u0000\u0019"+
		"b\u0001\u0000\u0000\u0000\u001bh\u0001\u0000\u0000\u0000\u001d\u001e\u0005"+
		"i\u0000\u0000\u001e\u001f\u0005m\u0000\u0000\u001f \u0005p\u0000\u0000"+
		" !\u0005o\u0000\u0000!\"\u0005r\u0000\u0000\"#\u0005t\u0000\u0000#\u0002"+
		"\u0001\u0000\u0000\u0000$%\u0005{\u0000\u0000%\u0004\u0001\u0000\u0000"+
		"\u0000&\'\u0005}\u0000\u0000\'\u0006\u0001\u0000\u0000\u0000()\u0005="+
		"\u0000\u0000)\b\u0001\u0000\u0000\u0000*+\u0005r\u0000\u0000+,\u0005e"+
		"\u0000\u0000,-\u0005p\u0000\u0000-.\u0005e\u0000\u0000./\u0005a\u0000"+
		"\u0000/0\u0005t\u0000\u000001\u0005e\u0000\u000012\u0005d\u0000\u0000"+
		"2\n\u0001\u0000\u0000\u000034\u0005[\u0000\u00004\f\u0001\u0000\u0000"+
		"\u000056\u0005,\u0000\u00006\u000e\u0001\u0000\u0000\u000078\u0005]\u0000"+
		"\u00008\u0010\u0001\u0000\u0000\u00009=\u0005\"\u0000\u0000:<\b\u0000"+
		"\u0000\u0000;:\u0001\u0000\u0000\u0000<?\u0001\u0000\u0000\u0000=;\u0001"+
		"\u0000\u0000\u0000=>\u0001\u0000\u0000\u0000>@\u0001\u0000\u0000\u0000"+
		"?=\u0001\u0000\u0000\u0000@A\u0005\"\u0000\u0000A\u0012\u0001\u0000\u0000"+
		"\u0000BD\u0007\u0001\u0000\u0000CB\u0001\u0000\u0000\u0000DE\u0001\u0000"+
		"\u0000\u0000EC\u0001\u0000\u0000\u0000EF\u0001\u0000\u0000\u0000FM\u0001"+
		"\u0000\u0000\u0000GI\u0005.\u0000\u0000HJ\u0007\u0001\u0000\u0000IH\u0001"+
		"\u0000\u0000\u0000JK\u0001\u0000\u0000\u0000KI\u0001\u0000\u0000\u0000"+
		"KL\u0001\u0000\u0000\u0000LN\u0001\u0000\u0000\u0000MG\u0001\u0000\u0000"+
		"\u0000MN\u0001\u0000\u0000\u0000N\u0014\u0001\u0000\u0000\u0000OP\u0005"+
		"t\u0000\u0000PQ\u0005r\u0000\u0000QR\u0005u\u0000\u0000RY\u0005e\u0000"+
		"\u0000ST\u0005f\u0000\u0000TU\u0005a\u0000\u0000UV\u0005l\u0000\u0000"+
		"VW\u0005s\u0000\u0000WY\u0005e\u0000\u0000XO\u0001\u0000\u0000\u0000X"+
		"S\u0001\u0000\u0000\u0000Y\u0016\u0001\u0000\u0000\u0000Z^\u0007\u0002"+
		"\u0000\u0000[]\u0007\u0003\u0000\u0000\\[\u0001\u0000\u0000\u0000]`\u0001"+
		"\u0000\u0000\u0000^\\\u0001\u0000\u0000\u0000^_\u0001\u0000\u0000\u0000"+
		"_\u0018\u0001\u0000\u0000\u0000`^\u0001\u0000\u0000\u0000ac\u0007\u0004"+
		"\u0000\u0000ba\u0001\u0000\u0000\u0000cd\u0001\u0000\u0000\u0000db\u0001"+
		"\u0000\u0000\u0000de\u0001\u0000\u0000\u0000ef\u0001\u0000\u0000\u0000"+
		"fg\u0006\f\u0000\u0000g\u001a\u0001\u0000\u0000\u0000hl\u0005#\u0000\u0000"+
		"ik\b\u0005\u0000\u0000ji\u0001\u0000\u0000\u0000kn\u0001\u0000\u0000\u0000"+
		"lj\u0001\u0000\u0000\u0000lm\u0001\u0000\u0000\u0000mo\u0001\u0000\u0000"+
		"\u0000nl\u0001\u0000\u0000\u0000op\u0006\r\u0000\u0000p\u001c\u0001\u0000"+
		"\u0000\u0000\t\u0000=EKMX^dl\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}