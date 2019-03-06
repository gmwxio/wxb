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

	// EnterTop_level_statement is called when entering the top_level_statement production.
	EnterTop_level_statement(c *Top_level_statementContext)

	// EnterOption_File is called when entering the Option_File production.
	EnterOption_File(c *Option_FileContext)

	// EnterOption_Msg is called when entering the Option_Msg production.
	EnterOption_Msg(c *Option_MsgContext)

	// EnterOption_Enum is called when entering the Option_Enum production.
	EnterOption_Enum(c *Option_EnumContext)

	// EnterOption_Service is called when entering the Option_Service production.
	EnterOption_Service(c *Option_ServiceContext)

	// EnterOption_Rpc is called when entering the Option_Rpc production.
	EnterOption_Rpc(c *Option_RpcContext)

	// EnterOptionName is called when entering the optionName production.
	EnterOptionName(c *OptionNameContext)

	// EnterExtend is called when entering the extend production.
	EnterExtend(c *ExtendContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterPackageStatement is called when entering the packageStatement production.
	EnterPackageStatement(c *PackageStatementContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// EnterMessageBody is called when entering the messageBody production.
	EnterMessageBody(c *MessageBodyContext)

	// EnterEnumDefinition is called when entering the enumDefinition production.
	EnterEnumDefinition(c *EnumDefinitionContext)

	// EnterEnumBody is called when entering the enumBody production.
	EnterEnumBody(c *EnumBodyContext)

	// EnterEnumField is called when entering the enumField production.
	EnterEnumField(c *EnumFieldContext)

	// EnterEnumValueOption is called when entering the enumValueOption production.
	EnterEnumValueOption(c *EnumValueOptionContext)

	// EnterService is called when entering the service production.
	EnterService(c *ServiceContext)

	// EnterServiceBody is called when entering the serviceBody production.
	EnterServiceBody(c *ServiceBodyContext)

	// EnterRpc is called when entering the rpc production.
	EnterRpc(c *RpcContext)

	// EnterRpcParam is called when entering the rpcParam production.
	EnterRpcParam(c *RpcParamContext)

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

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldOptions is called when entering the fieldOptions production.
	EnterFieldOptions(c *FieldOptionsContext)

	// EnterFieldOption is called when entering the fieldOption production.
	EnterFieldOption(c *FieldOptionContext)

	// EnterOneof is called when entering the oneof production.
	EnterOneof(c *OneofContext)

	// EnterOneofField is called when entering the oneofField production.
	EnterOneofField(c *OneofFieldContext)

	// EnterMapField is called when entering the mapField production.
	EnterMapField(c *MapFieldContext)

	// EnterKeyType is called when entering the keyType production.
	EnterKeyType(c *KeyTypeContext)

	// EnterMessageType is called when entering the messageType production.
	EnterMessageType(c *MessageTypeContext)

	// EnterMessageOrEnumType is called when entering the messageOrEnumType production.
	EnterMessageOrEnumType(c *MessageOrEnumTypeContext)

	// EnterBool_lit is called when entering the bool_lit production.
	EnterBool_lit(c *Bool_litContext)

	// EnterEmptyStatement is called when entering the emptyStatement production.
	EnterEmptyStatement(c *EmptyStatementContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterPron is called when entering the pron production.
	EnterPron(c *PronContext)

	// EnterPronSTR is called when entering the PronSTR production.
	EnterPronSTR(c *PronSTRContext)

	// EnterPronOBJ is called when entering the PronOBJ production.
	EnterPronOBJ(c *PronOBJContext)

	// ExitProto is called when exiting the proto production.
	ExitProto(c *ProtoContext)

	// ExitSyntax is called when exiting the syntax production.
	ExitSyntax(c *SyntaxContext)

	// ExitTop_level_statement is called when exiting the top_level_statement production.
	ExitTop_level_statement(c *Top_level_statementContext)

	// ExitOption_File is called when exiting the Option_File production.
	ExitOption_File(c *Option_FileContext)

	// ExitOption_Msg is called when exiting the Option_Msg production.
	ExitOption_Msg(c *Option_MsgContext)

	// ExitOption_Enum is called when exiting the Option_Enum production.
	ExitOption_Enum(c *Option_EnumContext)

	// ExitOption_Service is called when exiting the Option_Service production.
	ExitOption_Service(c *Option_ServiceContext)

	// ExitOption_Rpc is called when exiting the Option_Rpc production.
	ExitOption_Rpc(c *Option_RpcContext)

	// ExitOptionName is called when exiting the optionName production.
	ExitOptionName(c *OptionNameContext)

	// ExitExtend is called when exiting the extend production.
	ExitExtend(c *ExtendContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitPackageStatement is called when exiting the packageStatement production.
	ExitPackageStatement(c *PackageStatementContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)

	// ExitMessageBody is called when exiting the messageBody production.
	ExitMessageBody(c *MessageBodyContext)

	// ExitEnumDefinition is called when exiting the enumDefinition production.
	ExitEnumDefinition(c *EnumDefinitionContext)

	// ExitEnumBody is called when exiting the enumBody production.
	ExitEnumBody(c *EnumBodyContext)

	// ExitEnumField is called when exiting the enumField production.
	ExitEnumField(c *EnumFieldContext)

	// ExitEnumValueOption is called when exiting the enumValueOption production.
	ExitEnumValueOption(c *EnumValueOptionContext)

	// ExitService is called when exiting the service production.
	ExitService(c *ServiceContext)

	// ExitServiceBody is called when exiting the serviceBody production.
	ExitServiceBody(c *ServiceBodyContext)

	// ExitRpc is called when exiting the rpc production.
	ExitRpc(c *RpcContext)

	// ExitRpcParam is called when exiting the rpcParam production.
	ExitRpcParam(c *RpcParamContext)

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

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldOptions is called when exiting the fieldOptions production.
	ExitFieldOptions(c *FieldOptionsContext)

	// ExitFieldOption is called when exiting the fieldOption production.
	ExitFieldOption(c *FieldOptionContext)

	// ExitOneof is called when exiting the oneof production.
	ExitOneof(c *OneofContext)

	// ExitOneofField is called when exiting the oneofField production.
	ExitOneofField(c *OneofFieldContext)

	// ExitMapField is called when exiting the mapField production.
	ExitMapField(c *MapFieldContext)

	// ExitKeyType is called when exiting the keyType production.
	ExitKeyType(c *KeyTypeContext)

	// ExitMessageType is called when exiting the messageType production.
	ExitMessageType(c *MessageTypeContext)

	// ExitMessageOrEnumType is called when exiting the messageOrEnumType production.
	ExitMessageOrEnumType(c *MessageOrEnumTypeContext)

	// ExitBool_lit is called when exiting the bool_lit production.
	ExitBool_lit(c *Bool_litContext)

	// ExitEmptyStatement is called when exiting the emptyStatement production.
	ExitEmptyStatement(c *EmptyStatementContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitPron is called when exiting the pron production.
	ExitPron(c *PronContext)

	// ExitPronSTR is called when exiting the PronSTR production.
	ExitPronSTR(c *PronSTRContext)

	// ExitPronOBJ is called when exiting the PronOBJ production.
	ExitPronOBJ(c *PronOBJContext)
}
