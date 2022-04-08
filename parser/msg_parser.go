// Code generated from MsgParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // MsgParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 54, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 5, 3, 22, 10, 3, 3, 3, 3, 3, 5,
	3, 26, 10, 3, 7, 3, 28, 10, 3, 12, 3, 14, 3, 31, 11, 3, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 6, 5, 38, 10, 5, 13, 5, 14, 5, 39, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 2, 2, 10, 2, 4,
	6, 8, 10, 12, 14, 16, 2, 3, 3, 2, 3, 4, 2, 49, 2, 18, 3, 2, 2, 2, 4, 21,
	3, 2, 2, 2, 6, 32, 3, 2, 2, 2, 8, 34, 3, 2, 2, 2, 10, 43, 3, 2, 2, 2, 12,
	45, 3, 2, 2, 2, 14, 49, 3, 2, 2, 2, 16, 51, 3, 2, 2, 2, 18, 19, 5, 4, 3,
	2, 19, 3, 3, 2, 2, 2, 20, 22, 5, 6, 4, 2, 21, 20, 3, 2, 2, 2, 21, 22, 3,
	2, 2, 2, 22, 29, 3, 2, 2, 2, 23, 25, 5, 8, 5, 2, 24, 26, 5, 6, 4, 2, 25,
	24, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 28, 3, 2, 2, 2, 27, 23, 3, 2, 2,
	2, 28, 31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 5, 3,
	2, 2, 2, 31, 29, 3, 2, 2, 2, 32, 33, 9, 2, 2, 2, 33, 7, 3, 2, 2, 2, 34,
	35, 7, 5, 2, 2, 35, 37, 5, 10, 6, 2, 36, 38, 5, 12, 7, 2, 37, 36, 3, 2,
	2, 2, 38, 39, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 41,
	3, 2, 2, 2, 41, 42, 7, 7, 2, 2, 42, 9, 3, 2, 2, 2, 43, 44, 7, 6, 2, 2,
	44, 11, 3, 2, 2, 2, 45, 46, 5, 14, 8, 2, 46, 47, 7, 10, 2, 2, 47, 48, 5,
	16, 9, 2, 48, 13, 3, 2, 2, 2, 49, 50, 7, 8, 2, 2, 50, 15, 3, 2, 2, 2, 51,
	52, 7, 9, 2, 2, 52, 17, 3, 2, 2, 2, 6, 21, 25, 29, 39,
}
var literalNames = []string{
	"", "", "", "'<x-'", "", "'/>'", "", "", "'='",
}
var symbolicNames = []string{
	"", "TEXT", "SEA_WS", "TAG_OPEN", "TAG_TYPE", "TAG_CLOSE", "ATTR_NAME",
	"ATTR_VALUE", "EQ", "WS",
}

var ruleNames = []string{
	"msg", "msgContent", "msgText", "msgTag", "tagType", "tagAttr", "attrName",
	"attrValue",
}

type MsgParser struct {
	*antlr.BaseParser
}

// NewMsgParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *MsgParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewMsgParser(input antlr.TokenStream) *MsgParser {
	this := new(MsgParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "MsgParser.g4"

	return this
}

// MsgParser tokens.
const (
	MsgParserEOF        = antlr.TokenEOF
	MsgParserTEXT       = 1
	MsgParserSEA_WS     = 2
	MsgParserTAG_OPEN   = 3
	MsgParserTAG_TYPE   = 4
	MsgParserTAG_CLOSE  = 5
	MsgParserATTR_NAME  = 6
	MsgParserATTR_VALUE = 7
	MsgParserEQ         = 8
	MsgParserWS         = 9
)

// MsgParser rules.
const (
	MsgParserRULE_msg        = 0
	MsgParserRULE_msgContent = 1
	MsgParserRULE_msgText    = 2
	MsgParserRULE_msgTag     = 3
	MsgParserRULE_tagType    = 4
	MsgParserRULE_tagAttr    = 5
	MsgParserRULE_attrName   = 6
	MsgParserRULE_attrValue  = 7
)

// IMsgContext is an interface to support dynamic dispatch.
type IMsgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMsgContext differentiates from other interfaces.
	IsMsgContext()
}

type MsgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsgContext() *MsgContext {
	var p = new(MsgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MsgParserRULE_msg
	return p
}

func (*MsgContext) IsMsgContext() {}

func NewMsgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsgContext {
	var p = new(MsgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsgParserRULE_msg

	return p
}

func (s *MsgContext) GetParser() antlr.Parser { return s.parser }

func (s *MsgContext) MsgContent() IMsgContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMsgContentContext)
}

func (s *MsgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.EnterMsg(s)
	}
}

func (s *MsgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.ExitMsg(s)
	}
}

func (p *MsgParser) Msg() (localctx IMsgContext) {
	this := p
	_ = this

	localctx = NewMsgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MsgParserRULE_msg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.MsgContent()
	}

	return localctx
}

// IMsgContentContext is an interface to support dynamic dispatch.
type IMsgContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMsgContentContext differentiates from other interfaces.
	IsMsgContentContext()
}

type MsgContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsgContentContext() *MsgContentContext {
	var p = new(MsgContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MsgParserRULE_msgContent
	return p
}

func (*MsgContentContext) IsMsgContentContext() {}

func NewMsgContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsgContentContext {
	var p = new(MsgContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsgParserRULE_msgContent

	return p
}

func (s *MsgContentContext) GetParser() antlr.Parser { return s.parser }

func (s *MsgContentContext) AllMsgText() []IMsgTextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMsgTextContext)(nil)).Elem())
	var tst = make([]IMsgTextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMsgTextContext)
		}
	}

	return tst
}

func (s *MsgContentContext) MsgText(i int) IMsgTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgTextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMsgTextContext)
}

func (s *MsgContentContext) AllMsgTag() []IMsgTagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMsgTagContext)(nil)).Elem())
	var tst = make([]IMsgTagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMsgTagContext)
		}
	}

	return tst
}

func (s *MsgContentContext) MsgTag(i int) IMsgTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgTagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMsgTagContext)
}

func (s *MsgContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsgContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.EnterMsgContent(s)
	}
}

func (s *MsgContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.ExitMsgContent(s)
	}
}

func (p *MsgParser) MsgContent() (localctx IMsgContentContext) {
	this := p
	_ = this

	localctx = NewMsgContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MsgParserRULE_msgContent)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MsgParserTEXT || _la == MsgParserSEA_WS {
		{
			p.SetState(18)
			p.MsgText()
		}

	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MsgParserTAG_OPEN {
		{
			p.SetState(21)
			p.MsgTag()
		}
		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MsgParserTEXT || _la == MsgParserSEA_WS {
			{
				p.SetState(22)
				p.MsgText()
			}

		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMsgTextContext is an interface to support dynamic dispatch.
type IMsgTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMsgTextContext differentiates from other interfaces.
	IsMsgTextContext()
}

type MsgTextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsgTextContext() *MsgTextContext {
	var p = new(MsgTextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MsgParserRULE_msgText
	return p
}

func (*MsgTextContext) IsMsgTextContext() {}

func NewMsgTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsgTextContext {
	var p = new(MsgTextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsgParserRULE_msgText

	return p
}

func (s *MsgTextContext) GetParser() antlr.Parser { return s.parser }

func (s *MsgTextContext) TEXT() antlr.TerminalNode {
	return s.GetToken(MsgParserTEXT, 0)
}

func (s *MsgTextContext) SEA_WS() antlr.TerminalNode {
	return s.GetToken(MsgParserSEA_WS, 0)
}

func (s *MsgTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsgTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.EnterMsgText(s)
	}
}

func (s *MsgTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.ExitMsgText(s)
	}
}

func (p *MsgParser) MsgText() (localctx IMsgTextContext) {
	this := p
	_ = this

	localctx = NewMsgTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MsgParserRULE_msgText)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MsgParserTEXT || _la == MsgParserSEA_WS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMsgTagContext is an interface to support dynamic dispatch.
type IMsgTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMsgTagContext differentiates from other interfaces.
	IsMsgTagContext()
}

type MsgTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsgTagContext() *MsgTagContext {
	var p = new(MsgTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MsgParserRULE_msgTag
	return p
}

func (*MsgTagContext) IsMsgTagContext() {}

func NewMsgTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsgTagContext {
	var p = new(MsgTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsgParserRULE_msgTag

	return p
}

func (s *MsgTagContext) GetParser() antlr.Parser { return s.parser }

func (s *MsgTagContext) TAG_OPEN() antlr.TerminalNode {
	return s.GetToken(MsgParserTAG_OPEN, 0)
}

func (s *MsgTagContext) TagType() ITagTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagTypeContext)
}

func (s *MsgTagContext) TAG_CLOSE() antlr.TerminalNode {
	return s.GetToken(MsgParserTAG_CLOSE, 0)
}

func (s *MsgTagContext) AllTagAttr() []ITagAttrContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagAttrContext)(nil)).Elem())
	var tst = make([]ITagAttrContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagAttrContext)
		}
	}

	return tst
}

func (s *MsgTagContext) TagAttr(i int) ITagAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagAttrContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagAttrContext)
}

func (s *MsgTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsgTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.EnterMsgTag(s)
	}
}

func (s *MsgTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.ExitMsgTag(s)
	}
}

func (p *MsgParser) MsgTag() (localctx IMsgTagContext) {
	this := p
	_ = this

	localctx = NewMsgTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MsgParserRULE_msgTag)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(MsgParserTAG_OPEN)
	}
	{
		p.SetState(33)
		p.TagType()
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MsgParserATTR_NAME {
		{
			p.SetState(34)
			p.TagAttr()
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(39)
		p.Match(MsgParserTAG_CLOSE)
	}

	return localctx
}

// ITagTypeContext is an interface to support dynamic dispatch.
type ITagTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagTypeContext differentiates from other interfaces.
	IsTagTypeContext()
}

type TagTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagTypeContext() *TagTypeContext {
	var p = new(TagTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MsgParserRULE_tagType
	return p
}

func (*TagTypeContext) IsTagTypeContext() {}

func NewTagTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagTypeContext {
	var p = new(TagTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsgParserRULE_tagType

	return p
}

func (s *TagTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TagTypeContext) TAG_TYPE() antlr.TerminalNode {
	return s.GetToken(MsgParserTAG_TYPE, 0)
}

func (s *TagTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.EnterTagType(s)
	}
}

func (s *TagTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.ExitTagType(s)
	}
}

func (p *MsgParser) TagType() (localctx ITagTypeContext) {
	this := p
	_ = this

	localctx = NewTagTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MsgParserRULE_tagType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(MsgParserTAG_TYPE)
	}

	return localctx
}

// ITagAttrContext is an interface to support dynamic dispatch.
type ITagAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagAttrContext differentiates from other interfaces.
	IsTagAttrContext()
}

type TagAttrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagAttrContext() *TagAttrContext {
	var p = new(TagAttrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MsgParserRULE_tagAttr
	return p
}

func (*TagAttrContext) IsTagAttrContext() {}

func NewTagAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagAttrContext {
	var p = new(TagAttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsgParserRULE_tagAttr

	return p
}

func (s *TagAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *TagAttrContext) AttrName() IAttrNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrNameContext)
}

func (s *TagAttrContext) EQ() antlr.TerminalNode {
	return s.GetToken(MsgParserEQ, 0)
}

func (s *TagAttrContext) AttrValue() IAttrValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrValueContext)
}

func (s *TagAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.EnterTagAttr(s)
	}
}

func (s *TagAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.ExitTagAttr(s)
	}
}

func (p *MsgParser) TagAttr() (localctx ITagAttrContext) {
	this := p
	_ = this

	localctx = NewTagAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MsgParserRULE_tagAttr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.AttrName()
	}
	{
		p.SetState(44)
		p.Match(MsgParserEQ)
	}
	{
		p.SetState(45)
		p.AttrValue()
	}

	return localctx
}

// IAttrNameContext is an interface to support dynamic dispatch.
type IAttrNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrNameContext differentiates from other interfaces.
	IsAttrNameContext()
}

type AttrNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrNameContext() *AttrNameContext {
	var p = new(AttrNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MsgParserRULE_attrName
	return p
}

func (*AttrNameContext) IsAttrNameContext() {}

func NewAttrNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrNameContext {
	var p = new(AttrNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsgParserRULE_attrName

	return p
}

func (s *AttrNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrNameContext) ATTR_NAME() antlr.TerminalNode {
	return s.GetToken(MsgParserATTR_NAME, 0)
}

func (s *AttrNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.EnterAttrName(s)
	}
}

func (s *AttrNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.ExitAttrName(s)
	}
}

func (p *MsgParser) AttrName() (localctx IAttrNameContext) {
	this := p
	_ = this

	localctx = NewAttrNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MsgParserRULE_attrName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(MsgParserATTR_NAME)
	}

	return localctx
}

// IAttrValueContext is an interface to support dynamic dispatch.
type IAttrValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrValueContext differentiates from other interfaces.
	IsAttrValueContext()
}

type AttrValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrValueContext() *AttrValueContext {
	var p = new(AttrValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MsgParserRULE_attrValue
	return p
}

func (*AttrValueContext) IsAttrValueContext() {}

func NewAttrValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrValueContext {
	var p = new(AttrValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MsgParserRULE_attrValue

	return p
}

func (s *AttrValueContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrValueContext) ATTR_VALUE() antlr.TerminalNode {
	return s.GetToken(MsgParserATTR_VALUE, 0)
}

func (s *AttrValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.EnterAttrValue(s)
	}
}

func (s *AttrValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MsgParserListener); ok {
		listenerT.ExitAttrValue(s)
	}
}

func (p *MsgParser) AttrValue() (localctx IAttrValueContext) {
	this := p
	_ = this

	localctx = NewAttrValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MsgParserRULE_attrValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(MsgParserATTR_VALUE)
	}

	return localctx
}
