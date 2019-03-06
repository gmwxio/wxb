// Code generated from tronParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // tronParser

import "github.com/wxio/goantlr"

// tronParserListener is a complete listener for a parse tree produced by tronParser.
type tronParserListener interface {
	antlr.ParseTreeListener

	// EnterProto is called when entering the proto production.
	EnterProto(c *ProtoContext)

	// EnterSyntax is called when entering the syntax production.
	EnterSyntax(c *SyntaxContext)

	// EnterImportStatement is called when entering the ImportStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterPackageStatement is called when entering the PackageStatement production.
	EnterPackageStatement(c *PackageStatementContext)

	// EnterExtend is called when entering the Extend production.
	EnterExtend(c *ExtendContext)

	// EnterOptionDef is called when entering the OptionDef production.
	EnterOptionDef(c *OptionDefContext)

	// EnterMessage is called when entering the Message production.
	EnterMessage(c *MessageContext)

	// EnterEnumDefinition is called when entering the EnumDefinition production.
	EnterEnumDefinition(c *EnumDefinitionContext)

	// EnterService is called when entering the Service production.
	EnterService(c *ServiceContext)

	// EnterEmptyStm is called when entering the EmptyStm production.
	EnterEmptyStm(c *EmptyStmContext)

	// EnterMessageBody is called when entering the messageBody production.
	EnterMessageBody(c *MessageBodyContext)

	// EnterEnumBody is called when entering the enumBody production.
	EnterEnumBody(c *EnumBodyContext)

	// EnterEnumField is called when entering the enumField production.
	EnterEnumField(c *EnumFieldContext)

	// EnterServiceBody is called when entering the serviceBody production.
	EnterServiceBody(c *ServiceBodyContext)

	// EnterRpc is called when entering the rpc production.
	EnterRpc(c *RpcContext)

	// EnterRpcParam is called when entering the rpcParam production.
	EnterRpcParam(c *RpcParamContext)

	// EnterRpcStream is called when entering the rpcStream production.
	EnterRpcStream(c *RpcStreamContext)

	// EnterReserved is called when entering the reserved production.
	EnterReserved(c *ReservedContext)

	// EnterRanges is called when entering the ranges production.
	EnterRanges(c *RangesContext)

	// EnterRangee is called when entering the rangee production.
	EnterRangee(c *RangeeContext)

	// EnterFieldNames is called when entering the fieldNames production.
	EnterFieldNames(c *FieldNamesContext)

	// EnterTyper is called when entering the typer production.
	EnterTyper(c *TyperContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldRepeat is called when entering the fieldRepeat production.
	EnterFieldRepeat(c *FieldRepeatContext)

	// EnterOneof is called when entering the oneof production.
	EnterOneof(c *OneofContext)

	// EnterOneofField is called when entering the oneofField production.
	EnterOneofField(c *OneofFieldContext)

	// EnterMapField is called when entering the mapField production.
	EnterMapField(c *MapFieldContext)

	// EnterMessageType is called when entering the messageType production.
	EnterMessageType(c *MessageTypeContext)

	// EnterMessageOrEnumType is called when entering the messageOrEnumType production.
	EnterMessageOrEnumType(c *MessageOrEnumTypeContext)

	// EnterEmptyStatement is called when entering the emptyStatement production.
	EnterEmptyStatement(c *EmptyStatementContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

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

	// ExitExtend is called when exiting the Extend production.
	ExitExtend(c *ExtendContext)

	// ExitOptionDef is called when exiting the OptionDef production.
	ExitOptionDef(c *OptionDefContext)

	// ExitMessage is called when exiting the Message production.
	ExitMessage(c *MessageContext)

	// ExitEnumDefinition is called when exiting the EnumDefinition production.
	ExitEnumDefinition(c *EnumDefinitionContext)

	// ExitService is called when exiting the Service production.
	ExitService(c *ServiceContext)

	// ExitEmptyStm is called when exiting the EmptyStm production.
	ExitEmptyStm(c *EmptyStmContext)

	// ExitMessageBody is called when exiting the messageBody production.
	ExitMessageBody(c *MessageBodyContext)

	// ExitEnumBody is called when exiting the enumBody production.
	ExitEnumBody(c *EnumBodyContext)

	// ExitEnumField is called when exiting the enumField production.
	ExitEnumField(c *EnumFieldContext)

	// ExitServiceBody is called when exiting the serviceBody production.
	ExitServiceBody(c *ServiceBodyContext)

	// ExitRpc is called when exiting the rpc production.
	ExitRpc(c *RpcContext)

	// ExitRpcParam is called when exiting the rpcParam production.
	ExitRpcParam(c *RpcParamContext)

	// ExitRpcStream is called when exiting the rpcStream production.
	ExitRpcStream(c *RpcStreamContext)

	// ExitReserved is called when exiting the reserved production.
	ExitReserved(c *ReservedContext)

	// ExitRanges is called when exiting the ranges production.
	ExitRanges(c *RangesContext)

	// ExitRangee is called when exiting the rangee production.
	ExitRangee(c *RangeeContext)

	// ExitFieldNames is called when exiting the fieldNames production.
	ExitFieldNames(c *FieldNamesContext)

	// ExitTyper is called when exiting the typer production.
	ExitTyper(c *TyperContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldRepeat is called when exiting the fieldRepeat production.
	ExitFieldRepeat(c *FieldRepeatContext)

	// ExitOneof is called when exiting the oneof production.
	ExitOneof(c *OneofContext)

	// ExitOneofField is called when exiting the oneofField production.
	ExitOneofField(c *OneofFieldContext)

	// ExitMapField is called when exiting the mapField production.
	ExitMapField(c *MapFieldContext)

	// ExitMessageType is called when exiting the messageType production.
	ExitMessageType(c *MessageTypeContext)

	// ExitMessageOrEnumType is called when exiting the messageOrEnumType production.
	ExitMessageOrEnumType(c *MessageOrEnumTypeContext)

	// ExitEmptyStatement is called when exiting the emptyStatement production.
	ExitEmptyStatement(c *EmptyStatementContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

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
