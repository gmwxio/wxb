// Code generated from TronParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // TronParser

import "github.com/wxio/goantlr"

// TronParserListener is a complete listener for a parse tree produced by TronParser.
type TronParserListener interface {
	antlr.ParseTreeListener

	// EnterProto is called when entering the proto production.
	EnterProto(c *ProtoContext)

	// EnterSyntax is called when entering the syntax production.
	EnterSyntax(c *SyntaxContext)

	// EnterImportStatement is called when entering the ImportStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterPackageStatement is called when entering the PackageStatement production.
	EnterPackageStatement(c *PackageStatementContext)

	// EnterOptionFileDef is called when entering the OptionFileDef production.
	EnterOptionFileDef(c *OptionFileDefContext)

	// EnterExt_Msg_Enum_Svc is called when entering the Ext_Msg_Enum_Svc production.
	EnterExt_Msg_Enum_Svc(c *Ext_Msg_Enum_SvcContext)

	// EnterEmptyStm is called when entering the EmptyStm production.
	EnterEmptyStm(c *EmptyStmContext)

	// EnterAssoc is called when entering the Assoc production.
	EnterAssoc(c *AssocContext)

	// EnterEmptyTopLvl is called when entering the EmptyTopLvl production.
	EnterEmptyTopLvl(c *EmptyTopLvlContext)

	// EnterMsgEnumSvcExt is called when entering the MsgEnumSvcExt production.
	EnterMsgEnumSvcExt(c *MsgEnumSvcExtContext)

	// EnterEmptyTopLvlStm is called when entering the EmptyTopLvlStm production.
	EnterEmptyTopLvlStm(c *EmptyTopLvlStmContext)

	// EnterRange is called when entering the Range production.
	EnterRange(c *RangeContext)

	// EnterTLIOption is called when entering the TLIOption production.
	EnterTLIOption(c *TLIOptionContext)

	// EnterRPCSig is called when entering the RPCSig production.
	EnterRPCSig(c *RPCSigContext)

	// EnterEmptyStmStm is called when entering the EmptyStmStm production.
	EnterEmptyStmStm(c *EmptyStmStmContext)

	// EnterEnumLeft is called when entering the EnumLeft production.
	EnterEnumLeft(c *EnumLeftContext)

	// EnterMsgSvcExt is called when entering the MsgSvcExt production.
	EnterMsgSvcExt(c *MsgSvcExtContext)

	// EnterOpt_Single is called when entering the Opt_Single production.
	EnterOpt_Single(c *Opt_SingleContext)

	// EnterOpt is called when entering the Opt production.
	EnterOpt(c *OptContext)

	// EnterSingleFull_RepLocal is called when entering the SingleFull_RepLocal production.
	EnterSingleFull_RepLocal(c *SingleFull_RepLocalContext)

	// EnterSingleLocal is called when entering the SingleLocal production.
	EnterSingleLocal(c *SingleLocalContext)

	// EnterRepeated is called when entering the Repeated production.
	EnterRepeated(c *RepeatedContext)

	// EnterMapLeft is called when entering the MapLeft production.
	EnterMapLeft(c *MapLeftContext)

	// EnterMapLocalLeft is called when entering the MapLocalLeft production.
	EnterMapLocalLeft(c *MapLocalLeftContext)

	// EnterRight_assoc is called when entering the right_assoc production.
	EnterRight_assoc(c *Right_assocContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterFullname is called when entering the fullname production.
	EnterFullname(c *FullnameContext)

	// EnterMessageType is called when entering the messageType production.
	EnterMessageType(c *MessageTypeContext)

	// EnterRpcDelim is called when entering the rpcDelim production.
	EnterRpcDelim(c *RpcDelimContext)

	// EnterRanges is called when entering the ranges production.
	EnterRanges(c *RangesContext)

	// EnterRangee is called when entering the rangee production.
	EnterRangee(c *RangeeContext)

	// EnterFieldOptions is called when entering the fieldOptions production.
	EnterFieldOptions(c *FieldOptionsContext)

	// EnterFieldOption is called when entering the fieldOption production.
	EnterFieldOption(c *FieldOptionContext)

	// EnterOptionName is called when entering the optionName production.
	EnterOptionName(c *OptionNameContext)

	// EnterTronObj is called when entering the TronObj production.
	EnterTronObj(c *TronObjContext)

	// EnterTronObjs is called when entering the TronObjs production.
	EnterTronObjs(c *TronObjsContext)

	// EnterPronSTR is called when entering the PronSTR production.
	EnterPronSTR(c *PronSTRContext)

	// EnterPronARRAY is called when entering the PronARRAY production.
	EnterPronARRAY(c *PronARRAYContext)

	// EnterPronARRAYOFOBJ is called when entering the PronARRAYOFOBJ production.
	EnterPronARRAYOFOBJ(c *PronARRAYOFOBJContext)

	// ExitProto is called when exiting the proto production.
	ExitProto(c *ProtoContext)

	// ExitSyntax is called when exiting the syntax production.
	ExitSyntax(c *SyntaxContext)

	// ExitImportStatement is called when exiting the ImportStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitPackageStatement is called when exiting the PackageStatement production.
	ExitPackageStatement(c *PackageStatementContext)

	// ExitOptionFileDef is called when exiting the OptionFileDef production.
	ExitOptionFileDef(c *OptionFileDefContext)

	// ExitExt_Msg_Enum_Svc is called when exiting the Ext_Msg_Enum_Svc production.
	ExitExt_Msg_Enum_Svc(c *Ext_Msg_Enum_SvcContext)

	// ExitEmptyStm is called when exiting the EmptyStm production.
	ExitEmptyStm(c *EmptyStmContext)

	// ExitAssoc is called when exiting the Assoc production.
	ExitAssoc(c *AssocContext)

	// ExitEmptyTopLvl is called when exiting the EmptyTopLvl production.
	ExitEmptyTopLvl(c *EmptyTopLvlContext)

	// ExitMsgEnumSvcExt is called when exiting the MsgEnumSvcExt production.
	ExitMsgEnumSvcExt(c *MsgEnumSvcExtContext)

	// ExitEmptyTopLvlStm is called when exiting the EmptyTopLvlStm production.
	ExitEmptyTopLvlStm(c *EmptyTopLvlStmContext)

	// ExitRange is called when exiting the Range production.
	ExitRange(c *RangeContext)

	// ExitTLIOption is called when exiting the TLIOption production.
	ExitTLIOption(c *TLIOptionContext)

	// ExitRPCSig is called when exiting the RPCSig production.
	ExitRPCSig(c *RPCSigContext)

	// ExitEmptyStmStm is called when exiting the EmptyStmStm production.
	ExitEmptyStmStm(c *EmptyStmStmContext)

	// ExitEnumLeft is called when exiting the EnumLeft production.
	ExitEnumLeft(c *EnumLeftContext)

	// ExitMsgSvcExt is called when exiting the MsgSvcExt production.
	ExitMsgSvcExt(c *MsgSvcExtContext)

	// ExitOpt_Single is called when exiting the Opt_Single production.
	ExitOpt_Single(c *Opt_SingleContext)

	// ExitOpt is called when exiting the Opt production.
	ExitOpt(c *OptContext)

	// ExitSingleFull_RepLocal is called when exiting the SingleFull_RepLocal production.
	ExitSingleFull_RepLocal(c *SingleFull_RepLocalContext)

	// ExitSingleLocal is called when exiting the SingleLocal production.
	ExitSingleLocal(c *SingleLocalContext)

	// ExitRepeated is called when exiting the Repeated production.
	ExitRepeated(c *RepeatedContext)

	// ExitMapLeft is called when exiting the MapLeft production.
	ExitMapLeft(c *MapLeftContext)

	// ExitMapLocalLeft is called when exiting the MapLocalLeft production.
	ExitMapLocalLeft(c *MapLocalLeftContext)

	// ExitRight_assoc is called when exiting the right_assoc production.
	ExitRight_assoc(c *Right_assocContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitFullname is called when exiting the fullname production.
	ExitFullname(c *FullnameContext)

	// ExitMessageType is called when exiting the messageType production.
	ExitMessageType(c *MessageTypeContext)

	// ExitRpcDelim is called when exiting the rpcDelim production.
	ExitRpcDelim(c *RpcDelimContext)

	// ExitRanges is called when exiting the ranges production.
	ExitRanges(c *RangesContext)

	// ExitRangee is called when exiting the rangee production.
	ExitRangee(c *RangeeContext)

	// ExitFieldOptions is called when exiting the fieldOptions production.
	ExitFieldOptions(c *FieldOptionsContext)

	// ExitFieldOption is called when exiting the fieldOption production.
	ExitFieldOption(c *FieldOptionContext)

	// ExitOptionName is called when exiting the optionName production.
	ExitOptionName(c *OptionNameContext)

	// ExitTronObj is called when exiting the TronObj production.
	ExitTronObj(c *TronObjContext)

	// ExitTronObjs is called when exiting the TronObjs production.
	ExitTronObjs(c *TronObjsContext)

	// ExitPronSTR is called when exiting the PronSTR production.
	ExitPronSTR(c *PronSTRContext)

	// ExitPronARRAY is called when exiting the PronARRAY production.
	ExitPronARRAY(c *PronARRAYContext)

	// ExitPronARRAYOFOBJ is called when exiting the PronARRAYOFOBJ production.
	ExitPronARRAYOFOBJ(c *PronARRAYOFOBJContext)
}
