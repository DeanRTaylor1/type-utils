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
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		STRING=10, NUMBER=11, BOOLEAN=12, IDENTIFIER=13, WS=14, COMMENT=15;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"STRING", "NUMBER", "BOOLEAN", "IDENTIFIER", "WS", "COMMENT"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'HCLCONFIG'", "'{'", "'}'", "'='", "'import'", "'repeated'", "'['", 
			"','", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, "STRING", 
			"NUMBER", "BOOLEAN", "IDENTIFIER", "WS", "COMMENT"
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
		"\u0004\u0000\u000f}\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001"+
		"\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001"+
		"\b\u0001\b\u0001\t\u0001\t\u0005\tH\b\t\n\t\f\tK\t\t\u0001\t\u0001\t\u0001"+
		"\n\u0004\nP\b\n\u000b\n\f\nQ\u0001\n\u0001\n\u0004\nV\b\n\u000b\n\f\n"+
		"W\u0003\nZ\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0003\u000be\b\u000b"+
		"\u0001\f\u0001\f\u0005\fi\b\f\n\f\f\fl\t\f\u0001\r\u0004\ro\b\r\u000b"+
		"\r\f\rp\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0005\u000ew\b\u000e\n"+
		"\u000e\f\u000ez\t\u000e\u0001\u000e\u0001\u000e\u0000\u0000\u000f\u0001"+
		"\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007"+
		"\u000f\b\u0011\t\u0013\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e\u001d"+
		"\u000f\u0001\u0000\u0006\u0003\u0000\n\n\r\r\"\"\u0001\u000009\u0003\u0000"+
		"AZ__az\u0004\u000009AZ__az\u0003\u0000\t\n\r\r  \u0002\u0000\n\n\r\r\u0084"+
		"\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000"+
		"\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000"+
		"\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000"+
		"\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011"+
		"\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015"+
		"\u0001\u0000\u0000\u0000\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019"+
		"\u0001\u0000\u0000\u0000\u0000\u001b\u0001\u0000\u0000\u0000\u0000\u001d"+
		"\u0001\u0000\u0000\u0000\u0001\u001f\u0001\u0000\u0000\u0000\u0003)\u0001"+
		"\u0000\u0000\u0000\u0005+\u0001\u0000\u0000\u0000\u0007-\u0001\u0000\u0000"+
		"\u0000\t/\u0001\u0000\u0000\u0000\u000b6\u0001\u0000\u0000\u0000\r?\u0001"+
		"\u0000\u0000\u0000\u000fA\u0001\u0000\u0000\u0000\u0011C\u0001\u0000\u0000"+
		"\u0000\u0013E\u0001\u0000\u0000\u0000\u0015O\u0001\u0000\u0000\u0000\u0017"+
		"d\u0001\u0000\u0000\u0000\u0019f\u0001\u0000\u0000\u0000\u001bn\u0001"+
		"\u0000\u0000\u0000\u001dt\u0001\u0000\u0000\u0000\u001f \u0005H\u0000"+
		"\u0000 !\u0005C\u0000\u0000!\"\u0005L\u0000\u0000\"#\u0005C\u0000\u0000"+
		"#$\u0005O\u0000\u0000$%\u0005N\u0000\u0000%&\u0005F\u0000\u0000&\'\u0005"+
		"I\u0000\u0000\'(\u0005G\u0000\u0000(\u0002\u0001\u0000\u0000\u0000)*\u0005"+
		"{\u0000\u0000*\u0004\u0001\u0000\u0000\u0000+,\u0005}\u0000\u0000,\u0006"+
		"\u0001\u0000\u0000\u0000-.\u0005=\u0000\u0000.\b\u0001\u0000\u0000\u0000"+
		"/0\u0005i\u0000\u000001\u0005m\u0000\u000012\u0005p\u0000\u000023\u0005"+
		"o\u0000\u000034\u0005r\u0000\u000045\u0005t\u0000\u00005\n\u0001\u0000"+
		"\u0000\u000067\u0005r\u0000\u000078\u0005e\u0000\u000089\u0005p\u0000"+
		"\u00009:\u0005e\u0000\u0000:;\u0005a\u0000\u0000;<\u0005t\u0000\u0000"+
		"<=\u0005e\u0000\u0000=>\u0005d\u0000\u0000>\f\u0001\u0000\u0000\u0000"+
		"?@\u0005[\u0000\u0000@\u000e\u0001\u0000\u0000\u0000AB\u0005,\u0000\u0000"+
		"B\u0010\u0001\u0000\u0000\u0000CD\u0005]\u0000\u0000D\u0012\u0001\u0000"+
		"\u0000\u0000EI\u0005\"\u0000\u0000FH\b\u0000\u0000\u0000GF\u0001\u0000"+
		"\u0000\u0000HK\u0001\u0000\u0000\u0000IG\u0001\u0000\u0000\u0000IJ\u0001"+
		"\u0000\u0000\u0000JL\u0001\u0000\u0000\u0000KI\u0001\u0000\u0000\u0000"+
		"LM\u0005\"\u0000\u0000M\u0014\u0001\u0000\u0000\u0000NP\u0007\u0001\u0000"+
		"\u0000ON\u0001\u0000\u0000\u0000PQ\u0001\u0000\u0000\u0000QO\u0001\u0000"+
		"\u0000\u0000QR\u0001\u0000\u0000\u0000RY\u0001\u0000\u0000\u0000SU\u0005"+
		".\u0000\u0000TV\u0007\u0001\u0000\u0000UT\u0001\u0000\u0000\u0000VW\u0001"+
		"\u0000\u0000\u0000WU\u0001\u0000\u0000\u0000WX\u0001\u0000\u0000\u0000"+
		"XZ\u0001\u0000\u0000\u0000YS\u0001\u0000\u0000\u0000YZ\u0001\u0000\u0000"+
		"\u0000Z\u0016\u0001\u0000\u0000\u0000[\\\u0005t\u0000\u0000\\]\u0005r"+
		"\u0000\u0000]^\u0005u\u0000\u0000^e\u0005e\u0000\u0000_`\u0005f\u0000"+
		"\u0000`a\u0005a\u0000\u0000ab\u0005l\u0000\u0000bc\u0005s\u0000\u0000"+
		"ce\u0005e\u0000\u0000d[\u0001\u0000\u0000\u0000d_\u0001\u0000\u0000\u0000"+
		"e\u0018\u0001\u0000\u0000\u0000fj\u0007\u0002\u0000\u0000gi\u0007\u0003"+
		"\u0000\u0000hg\u0001\u0000\u0000\u0000il\u0001\u0000\u0000\u0000jh\u0001"+
		"\u0000\u0000\u0000jk\u0001\u0000\u0000\u0000k\u001a\u0001\u0000\u0000"+
		"\u0000lj\u0001\u0000\u0000\u0000mo\u0007\u0004\u0000\u0000nm\u0001\u0000"+
		"\u0000\u0000op\u0001\u0000\u0000\u0000pn\u0001\u0000\u0000\u0000pq\u0001"+
		"\u0000\u0000\u0000qr\u0001\u0000\u0000\u0000rs\u0006\r\u0000\u0000s\u001c"+
		"\u0001\u0000\u0000\u0000tx\u0005#\u0000\u0000uw\b\u0005\u0000\u0000vu"+
		"\u0001\u0000\u0000\u0000wz\u0001\u0000\u0000\u0000xv\u0001\u0000\u0000"+
		"\u0000xy\u0001\u0000\u0000\u0000y{\u0001\u0000\u0000\u0000zx\u0001\u0000"+
		"\u0000\u0000{|\u0006\u000e\u0000\u0000|\u001e\u0001\u0000\u0000\u0000"+
		"\t\u0000IQWYdjpx\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}