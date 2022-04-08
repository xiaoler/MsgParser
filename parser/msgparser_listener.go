// Code generated from MsgParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // MsgParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MsgParserListener is a complete listener for a parse tree produced by MsgParser.
type MsgParserListener interface {
	antlr.ParseTreeListener

	// EnterMsg is called when entering the msg production.
	EnterMsg(c *MsgContext)

	// EnterMsgContent is called when entering the msgContent production.
	EnterMsgContent(c *MsgContentContext)

	// EnterMsgText is called when entering the msgText production.
	EnterMsgText(c *MsgTextContext)

	// EnterMsgTag is called when entering the msgTag production.
	EnterMsgTag(c *MsgTagContext)

	// EnterTagType is called when entering the tagType production.
	EnterTagType(c *TagTypeContext)

	// EnterTagAttr is called when entering the tagAttr production.
	EnterTagAttr(c *TagAttrContext)

	// EnterAttrName is called when entering the attrName production.
	EnterAttrName(c *AttrNameContext)

	// EnterAttrValue is called when entering the attrValue production.
	EnterAttrValue(c *AttrValueContext)

	// ExitMsg is called when exiting the msg production.
	ExitMsg(c *MsgContext)

	// ExitMsgContent is called when exiting the msgContent production.
	ExitMsgContent(c *MsgContentContext)

	// ExitMsgText is called when exiting the msgText production.
	ExitMsgText(c *MsgTextContext)

	// ExitMsgTag is called when exiting the msgTag production.
	ExitMsgTag(c *MsgTagContext)

	// ExitTagType is called when exiting the tagType production.
	ExitTagType(c *TagTypeContext)

	// ExitTagAttr is called when exiting the tagAttr production.
	ExitTagAttr(c *TagAttrContext)

	// ExitAttrName is called when exiting the attrName production.
	ExitAttrName(c *AttrNameContext)

	// ExitAttrValue is called when exiting the attrValue production.
	ExitAttrValue(c *AttrValueContext)
}
