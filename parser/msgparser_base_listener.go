// Code generated from MsgParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // MsgParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMsgParserListener is a complete listener for a parse tree produced by MsgParser.
type BaseMsgParserListener struct{}

var _ MsgParserListener = &BaseMsgParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMsgParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMsgParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMsgParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMsgParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMsg is called when production msg is entered.
func (s *BaseMsgParserListener) EnterMsg(ctx *MsgContext) {}

// ExitMsg is called when production msg is exited.
func (s *BaseMsgParserListener) ExitMsg(ctx *MsgContext) {}

// EnterMsgContent is called when production msgContent is entered.
func (s *BaseMsgParserListener) EnterMsgContent(ctx *MsgContentContext) {}

// ExitMsgContent is called when production msgContent is exited.
func (s *BaseMsgParserListener) ExitMsgContent(ctx *MsgContentContext) {}

// EnterMsgText is called when production msgText is entered.
func (s *BaseMsgParserListener) EnterMsgText(ctx *MsgTextContext) {}

// ExitMsgText is called when production msgText is exited.
func (s *BaseMsgParserListener) ExitMsgText(ctx *MsgTextContext) {}

// EnterMsgTag is called when production msgTag is entered.
func (s *BaseMsgParserListener) EnterMsgTag(ctx *MsgTagContext) {}

// ExitMsgTag is called when production msgTag is exited.
func (s *BaseMsgParserListener) ExitMsgTag(ctx *MsgTagContext) {}

// EnterTagType is called when production tagType is entered.
func (s *BaseMsgParserListener) EnterTagType(ctx *TagTypeContext) {}

// ExitTagType is called when production tagType is exited.
func (s *BaseMsgParserListener) ExitTagType(ctx *TagTypeContext) {}

// EnterTagAttr is called when production tagAttr is entered.
func (s *BaseMsgParserListener) EnterTagAttr(ctx *TagAttrContext) {}

// ExitTagAttr is called when production tagAttr is exited.
func (s *BaseMsgParserListener) ExitTagAttr(ctx *TagAttrContext) {}

// EnterAttrName is called when production attrName is entered.
func (s *BaseMsgParserListener) EnterAttrName(ctx *AttrNameContext) {}

// ExitAttrName is called when production attrName is exited.
func (s *BaseMsgParserListener) ExitAttrName(ctx *AttrNameContext) {}

// EnterAttrValue is called when production attrValue is entered.
func (s *BaseMsgParserListener) EnterAttrValue(ctx *AttrValueContext) {}

// ExitAttrValue is called when production attrValue is exited.
func (s *BaseMsgParserListener) ExitAttrValue(ctx *AttrValueContext) {}
