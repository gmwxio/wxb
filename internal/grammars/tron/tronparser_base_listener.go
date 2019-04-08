// Code generated from TronParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // TronParser

import "github.com/wxio/goantlr"

// BaseTronParserListener is a complete listener for a parse tree produced by TronParser.
type BaseTronParserListener struct{}

var _ TronParserListener = &BaseTronParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTronParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTronParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTronParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTronParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProto is called when production proto is entered.
func (s *BaseTronParserListener) EnterProto(ctx *ProtoContext) {}

// ExitProto is called when production proto is exited.
func (s *BaseTronParserListener) ExitProto(ctx *ProtoContext) {}

// EnterSyntax is called when production syntax is entered.
func (s *BaseTronParserListener) EnterSyntax(ctx *SyntaxContext) {}

// ExitSyntax is called when production syntax is exited.
func (s *BaseTronParserListener) ExitSyntax(ctx *SyntaxContext) {}

// EnterImportStatement is called when production ImportStatement is entered.
func (s *BaseTronParserListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production ImportStatement is exited.
func (s *BaseTronParserListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterPackageStatement is called when production PackageStatement is entered.
func (s *BaseTronParserListener) EnterPackageStatement(ctx *PackageStatementContext) {}

// ExitPackageStatement is called when production PackageStatement is exited.
func (s *BaseTronParserListener) ExitPackageStatement(ctx *PackageStatementContext) {}

// EnterOptionFileDef is called when production OptionFileDef is entered.
func (s *BaseTronParserListener) EnterOptionFileDef(ctx *OptionFileDefContext) {}

// ExitOptionFileDef is called when production OptionFileDef is exited.
func (s *BaseTronParserListener) ExitOptionFileDef(ctx *OptionFileDefContext) {}

// EnterExt_Msg_Enum_Svc is called when production Ext_Msg_Enum_Svc is entered.
func (s *BaseTronParserListener) EnterExt_Msg_Enum_Svc(ctx *Ext_Msg_Enum_SvcContext) {}

// ExitExt_Msg_Enum_Svc is called when production Ext_Msg_Enum_Svc is exited.
func (s *BaseTronParserListener) ExitExt_Msg_Enum_Svc(ctx *Ext_Msg_Enum_SvcContext) {}

// EnterEmptyStm is called when production EmptyStm is entered.
func (s *BaseTronParserListener) EnterEmptyStm(ctx *EmptyStmContext) {}

// ExitEmptyStm is called when production EmptyStm is exited.
func (s *BaseTronParserListener) ExitEmptyStm(ctx *EmptyStmContext) {}

// EnterAssoc is called when production Assoc is entered.
func (s *BaseTronParserListener) EnterAssoc(ctx *AssocContext) {}

// ExitAssoc is called when production Assoc is exited.
func (s *BaseTronParserListener) ExitAssoc(ctx *AssocContext) {}

// EnterEmptyTopLvl is called when production EmptyTopLvl is entered.
func (s *BaseTronParserListener) EnterEmptyTopLvl(ctx *EmptyTopLvlContext) {}

// ExitEmptyTopLvl is called when production EmptyTopLvl is exited.
func (s *BaseTronParserListener) ExitEmptyTopLvl(ctx *EmptyTopLvlContext) {}

// EnterMsgEnumSvcExt is called when production MsgEnumSvcExt is entered.
func (s *BaseTronParserListener) EnterMsgEnumSvcExt(ctx *MsgEnumSvcExtContext) {}

// ExitMsgEnumSvcExt is called when production MsgEnumSvcExt is exited.
func (s *BaseTronParserListener) ExitMsgEnumSvcExt(ctx *MsgEnumSvcExtContext) {}

// EnterEmptyTopLvlStm is called when production EmptyTopLvlStm is entered.
func (s *BaseTronParserListener) EnterEmptyTopLvlStm(ctx *EmptyTopLvlStmContext) {}

// ExitEmptyTopLvlStm is called when production EmptyTopLvlStm is exited.
func (s *BaseTronParserListener) ExitEmptyTopLvlStm(ctx *EmptyTopLvlStmContext) {}

// EnterRange is called when production Range is entered.
func (s *BaseTronParserListener) EnterRange(ctx *RangeContext) {}

// ExitRange is called when production Range is exited.
func (s *BaseTronParserListener) ExitRange(ctx *RangeContext) {}

// EnterTLIOption is called when production TLIOption is entered.
func (s *BaseTronParserListener) EnterTLIOption(ctx *TLIOptionContext) {}

// ExitTLIOption is called when production TLIOption is exited.
func (s *BaseTronParserListener) ExitTLIOption(ctx *TLIOptionContext) {}

// EnterRPCSig is called when production RPCSig is entered.
func (s *BaseTronParserListener) EnterRPCSig(ctx *RPCSigContext) {}

// ExitRPCSig is called when production RPCSig is exited.
func (s *BaseTronParserListener) ExitRPCSig(ctx *RPCSigContext) {}

// EnterEmptyStmStm is called when production EmptyStmStm is entered.
func (s *BaseTronParserListener) EnterEmptyStmStm(ctx *EmptyStmStmContext) {}

// ExitEmptyStmStm is called when production EmptyStmStm is exited.
func (s *BaseTronParserListener) ExitEmptyStmStm(ctx *EmptyStmStmContext) {}

// EnterEnumLeft is called when production EnumLeft is entered.
func (s *BaseTronParserListener) EnterEnumLeft(ctx *EnumLeftContext) {}

// ExitEnumLeft is called when production EnumLeft is exited.
func (s *BaseTronParserListener) ExitEnumLeft(ctx *EnumLeftContext) {}

// EnterMsgSvcExt is called when production MsgSvcExt is entered.
func (s *BaseTronParserListener) EnterMsgSvcExt(ctx *MsgSvcExtContext) {}

// ExitMsgSvcExt is called when production MsgSvcExt is exited.
func (s *BaseTronParserListener) ExitMsgSvcExt(ctx *MsgSvcExtContext) {}

// EnterOpt_Single is called when production Opt_Single is entered.
func (s *BaseTronParserListener) EnterOpt_Single(ctx *Opt_SingleContext) {}

// ExitOpt_Single is called when production Opt_Single is exited.
func (s *BaseTronParserListener) ExitOpt_Single(ctx *Opt_SingleContext) {}

// EnterOpt is called when production Opt is entered.
func (s *BaseTronParserListener) EnterOpt(ctx *OptContext) {}

// ExitOpt is called when production Opt is exited.
func (s *BaseTronParserListener) ExitOpt(ctx *OptContext) {}

// EnterSingleFull_RepLocal is called when production SingleFull_RepLocal is entered.
func (s *BaseTronParserListener) EnterSingleFull_RepLocal(ctx *SingleFull_RepLocalContext) {}

// ExitSingleFull_RepLocal is called when production SingleFull_RepLocal is exited.
func (s *BaseTronParserListener) ExitSingleFull_RepLocal(ctx *SingleFull_RepLocalContext) {}

// EnterSingleLocal is called when production SingleLocal is entered.
func (s *BaseTronParserListener) EnterSingleLocal(ctx *SingleLocalContext) {}

// ExitSingleLocal is called when production SingleLocal is exited.
func (s *BaseTronParserListener) ExitSingleLocal(ctx *SingleLocalContext) {}

// EnterRepeated is called when production Repeated is entered.
func (s *BaseTronParserListener) EnterRepeated(ctx *RepeatedContext) {}

// ExitRepeated is called when production Repeated is exited.
func (s *BaseTronParserListener) ExitRepeated(ctx *RepeatedContext) {}

// EnterMapLeft is called when production MapLeft is entered.
func (s *BaseTronParserListener) EnterMapLeft(ctx *MapLeftContext) {}

// ExitMapLeft is called when production MapLeft is exited.
func (s *BaseTronParserListener) ExitMapLeft(ctx *MapLeftContext) {}

// EnterMapLocalLeft is called when production MapLocalLeft is entered.
func (s *BaseTronParserListener) EnterMapLocalLeft(ctx *MapLocalLeftContext) {}

// ExitMapLocalLeft is called when production MapLocalLeft is exited.
func (s *BaseTronParserListener) ExitMapLocalLeft(ctx *MapLocalLeftContext) {}

// EnterRight_assoc is called when production right_assoc is entered.
func (s *BaseTronParserListener) EnterRight_assoc(ctx *Right_assocContext) {}

// ExitRight_assoc is called when production right_assoc is exited.
func (s *BaseTronParserListener) ExitRight_assoc(ctx *Right_assocContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseTronParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseTronParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterFullname is called when production fullname is entered.
func (s *BaseTronParserListener) EnterFullname(ctx *FullnameContext) {}

// ExitFullname is called when production fullname is exited.
func (s *BaseTronParserListener) ExitFullname(ctx *FullnameContext) {}

// EnterMessageType is called when production messageType is entered.
func (s *BaseTronParserListener) EnterMessageType(ctx *MessageTypeContext) {}

// ExitMessageType is called when production messageType is exited.
func (s *BaseTronParserListener) ExitMessageType(ctx *MessageTypeContext) {}

// EnterRpcDelim is called when production rpcDelim is entered.
func (s *BaseTronParserListener) EnterRpcDelim(ctx *RpcDelimContext) {}

// ExitRpcDelim is called when production rpcDelim is exited.
func (s *BaseTronParserListener) ExitRpcDelim(ctx *RpcDelimContext) {}

// EnterRanges is called when production ranges is entered.
func (s *BaseTronParserListener) EnterRanges(ctx *RangesContext) {}

// ExitRanges is called when production ranges is exited.
func (s *BaseTronParserListener) ExitRanges(ctx *RangesContext) {}

// EnterRangee is called when production rangee is entered.
func (s *BaseTronParserListener) EnterRangee(ctx *RangeeContext) {}

// ExitRangee is called when production rangee is exited.
func (s *BaseTronParserListener) ExitRangee(ctx *RangeeContext) {}

// EnterFieldOptions is called when production fieldOptions is entered.
func (s *BaseTronParserListener) EnterFieldOptions(ctx *FieldOptionsContext) {}

// ExitFieldOptions is called when production fieldOptions is exited.
func (s *BaseTronParserListener) ExitFieldOptions(ctx *FieldOptionsContext) {}

// EnterFieldOption is called when production fieldOption is entered.
func (s *BaseTronParserListener) EnterFieldOption(ctx *FieldOptionContext) {}

// ExitFieldOption is called when production fieldOption is exited.
func (s *BaseTronParserListener) ExitFieldOption(ctx *FieldOptionContext) {}

// EnterOptionName is called when production optionName is entered.
func (s *BaseTronParserListener) EnterOptionName(ctx *OptionNameContext) {}

// ExitOptionName is called when production optionName is exited.
func (s *BaseTronParserListener) ExitOptionName(ctx *OptionNameContext) {}

// EnterTronObj is called when production TronObj is entered.
func (s *BaseTronParserListener) EnterTronObj(ctx *TronObjContext) {}

// ExitTronObj is called when production TronObj is exited.
func (s *BaseTronParserListener) ExitTronObj(ctx *TronObjContext) {}

// EnterTronObjs is called when production TronObjs is entered.
func (s *BaseTronParserListener) EnterTronObjs(ctx *TronObjsContext) {}

// ExitTronObjs is called when production TronObjs is exited.
func (s *BaseTronParserListener) ExitTronObjs(ctx *TronObjsContext) {}

// EnterPronSTR is called when production PronSTR is entered.
func (s *BaseTronParserListener) EnterPronSTR(ctx *PronSTRContext) {}

// ExitPronSTR is called when production PronSTR is exited.
func (s *BaseTronParserListener) ExitPronSTR(ctx *PronSTRContext) {}

// EnterPronARRAY is called when production PronARRAY is entered.
func (s *BaseTronParserListener) EnterPronARRAY(ctx *PronARRAYContext) {}

// ExitPronARRAY is called when production PronARRAY is exited.
func (s *BaseTronParserListener) ExitPronARRAY(ctx *PronARRAYContext) {}

// EnterPronARRAYOFOBJ is called when production PronARRAYOFOBJ is entered.
func (s *BaseTronParserListener) EnterPronARRAYOFOBJ(ctx *PronARRAYOFOBJContext) {}

// ExitPronARRAYOFOBJ is called when production PronARRAYOFOBJ is exited.
func (s *BaseTronParserListener) ExitPronARRAYOFOBJ(ctx *PronARRAYOFOBJContext) {}
