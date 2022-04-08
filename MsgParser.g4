parser grammar MsgParser;

options {
	tokenVocab = MsgLexer;
}

msg: msgContent;

msgContent: msgText? (msgTag msgText?)*;

msgText: TEXT | SEA_WS;

msgTag: TAG_OPEN tagType tagAttr+ TAG_CLOSE;

tagType: TAG_TYPE;

tagAttr: attrName EQ attrValue;

attrName: ATTR_NAME;

attrValue: ATTR_VALUE;
