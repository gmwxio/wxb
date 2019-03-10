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

	// EnterAssociaton is called when entering the associaton production.
	EnterAssociaton(c *AssociatonContext)

	// EnterEnumLeft is called when entering the EnumLeft production.
	EnterEnumLeft(c *EnumLeftContext)

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

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterConstantObj is called when entering the constantObj production.
	EnterConstantObj(c *ConstantObjContext)

	// EnterPronSTR is called when entering the PronSTR production.
	EnterPronSTR(c *PronSTRContext)

	// EnterPronOBJ is called when entering the PronOBJ production.
	EnterPronOBJ(c *PronOBJContext)

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

	// ExitAssociaton is called when exiting the associaton production.
	ExitAssociaton(c *AssociatonContext)

	// ExitEnumLeft is called when exiting the EnumLeft production.
	ExitEnumLeft(c *EnumLeftContext)

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

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitConstantObj is called when exiting the constantObj production.
	ExitConstantObj(c *ConstantObjContext)

	// ExitPronSTR is called when exiting the PronSTR production.
	ExitPronSTR(c *PronSTRContext)

	// ExitPronOBJ is called when exiting the PronOBJ production.
	ExitPronOBJ(c *PronOBJContext)
}
