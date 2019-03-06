// Code generated from tronParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // tronParser

import "github.com/wxio/goantlr"

// BasetronParserListener is a complete listener for a parse tree produced by tronParser.
type BasetronParserListener struct{}

var _ tronParserListener = &BasetronParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetronParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetronParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetronParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetronParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProto is called when production proto is entered.
func (s *BasetronParserListener) EnterProto(ctx *ProtoContext) {}

// ExitProto is called when production proto is exited.
func (s *BasetronParserListener) ExitProto(ctx *ProtoContext) {}

// EnterSyntax is called when production syntax is entered.
func (s *BasetronParserListener) EnterSyntax(ctx *SyntaxContext) {}

// ExitSyntax is called when production syntax is exited.
func (s *BasetronParserListener) ExitSyntax(ctx *SyntaxContext) {}

// EnterTop_level_statement is called when production top_level_statement is entered.
func (s *BasetronParserListener) EnterTop_level_statement(ctx *Top_level_statementContext) {}

// ExitTop_level_statement is called when production top_level_statement is exited.
func (s *BasetronParserListener) ExitTop_level_statement(ctx *Top_level_statementContext) {}

// EnterOption_File is called when production Option_File is entered.
func (s *BasetronParserListener) EnterOption_File(ctx *Option_FileContext) {}

// ExitOption_File is called when production Option_File is exited.
func (s *BasetronParserListener) ExitOption_File(ctx *Option_FileContext) {}

// EnterOption_Msg is called when production Option_Msg is entered.
func (s *BasetronParserListener) EnterOption_Msg(ctx *Option_MsgContext) {}

// ExitOption_Msg is called when production Option_Msg is exited.
func (s *BasetronParserListener) ExitOption_Msg(ctx *Option_MsgContext) {}

// EnterOption_Enum is called when production Option_Enum is entered.
func (s *BasetronParserListener) EnterOption_Enum(ctx *Option_EnumContext) {}

// ExitOption_Enum is called when production Option_Enum is exited.
func (s *BasetronParserListener) ExitOption_Enum(ctx *Option_EnumContext) {}

// EnterOption_Service is called when production Option_Service is entered.
func (s *BasetronParserListener) EnterOption_Service(ctx *Option_ServiceContext) {}

// ExitOption_Service is called when production Option_Service is exited.
func (s *BasetronParserListener) ExitOption_Service(ctx *Option_ServiceContext) {}

// EnterOption_Rpc is called when production Option_Rpc is entered.
func (s *BasetronParserListener) EnterOption_Rpc(ctx *Option_RpcContext) {}

// ExitOption_Rpc is called when production Option_Rpc is exited.
func (s *BasetronParserListener) ExitOption_Rpc(ctx *Option_RpcContext) {}

// EnterOptionName is called when production optionName is entered.
func (s *BasetronParserListener) EnterOptionName(ctx *OptionNameContext) {}

// ExitOptionName is called when production optionName is exited.
func (s *BasetronParserListener) ExitOptionName(ctx *OptionNameContext) {}

// EnterExtend is called when production extend is entered.
func (s *BasetronParserListener) EnterExtend(ctx *ExtendContext) {}

// ExitExtend is called when production extend is exited.
func (s *BasetronParserListener) ExitExtend(ctx *ExtendContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BasetronParserListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BasetronParserListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterPackageStatement is called when production packageStatement is entered.
func (s *BasetronParserListener) EnterPackageStatement(ctx *PackageStatementContext) {}

// ExitPackageStatement is called when production packageStatement is exited.
func (s *BasetronParserListener) ExitPackageStatement(ctx *PackageStatementContext) {}

// EnterMessage is called when production message is entered.
func (s *BasetronParserListener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *BasetronParserListener) ExitMessage(ctx *MessageContext) {}

// EnterMessageBody is called when production messageBody is entered.
func (s *BasetronParserListener) EnterMessageBody(ctx *MessageBodyContext) {}

// ExitMessageBody is called when production messageBody is exited.
func (s *BasetronParserListener) ExitMessageBody(ctx *MessageBodyContext) {}

// EnterEnumDefinition is called when production enumDefinition is entered.
func (s *BasetronParserListener) EnterEnumDefinition(ctx *EnumDefinitionContext) {}

// ExitEnumDefinition is called when production enumDefinition is exited.
func (s *BasetronParserListener) ExitEnumDefinition(ctx *EnumDefinitionContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BasetronParserListener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BasetronParserListener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterEnumField is called when production enumField is entered.
func (s *BasetronParserListener) EnterEnumField(ctx *EnumFieldContext) {}

// ExitEnumField is called when production enumField is exited.
func (s *BasetronParserListener) ExitEnumField(ctx *EnumFieldContext) {}

// EnterEnumValueOption is called when production enumValueOption is entered.
func (s *BasetronParserListener) EnterEnumValueOption(ctx *EnumValueOptionContext) {}

// ExitEnumValueOption is called when production enumValueOption is exited.
func (s *BasetronParserListener) ExitEnumValueOption(ctx *EnumValueOptionContext) {}

// EnterService is called when production service is entered.
func (s *BasetronParserListener) EnterService(ctx *ServiceContext) {}

// ExitService is called when production service is exited.
func (s *BasetronParserListener) ExitService(ctx *ServiceContext) {}

// EnterServiceBody is called when production serviceBody is entered.
func (s *BasetronParserListener) EnterServiceBody(ctx *ServiceBodyContext) {}

// ExitServiceBody is called when production serviceBody is exited.
func (s *BasetronParserListener) ExitServiceBody(ctx *ServiceBodyContext) {}

// EnterRpc is called when production rpc is entered.
func (s *BasetronParserListener) EnterRpc(ctx *RpcContext) {}

// ExitRpc is called when production rpc is exited.
func (s *BasetronParserListener) ExitRpc(ctx *RpcContext) {}

// EnterRpcParam is called when production rpcParam is entered.
func (s *BasetronParserListener) EnterRpcParam(ctx *RpcParamContext) {}

// ExitRpcParam is called when production rpcParam is exited.
func (s *BasetronParserListener) ExitRpcParam(ctx *RpcParamContext) {}

// EnterReserved is called when production reserved is entered.
func (s *BasetronParserListener) EnterReserved(ctx *ReservedContext) {}

// ExitReserved is called when production reserved is exited.
func (s *BasetronParserListener) ExitReserved(ctx *ReservedContext) {}

// EnterRanges is called when production ranges is entered.
func (s *BasetronParserListener) EnterRanges(ctx *RangesContext) {}

// ExitRanges is called when production ranges is exited.
func (s *BasetronParserListener) ExitRanges(ctx *RangesContext) {}

// EnterRangee is called when production rangee is entered.
func (s *BasetronParserListener) EnterRangee(ctx *RangeeContext) {}

// ExitRangee is called when production rangee is exited.
func (s *BasetronParserListener) ExitRangee(ctx *RangeeContext) {}

// EnterFieldNames is called when production fieldNames is entered.
func (s *BasetronParserListener) EnterFieldNames(ctx *FieldNamesContext) {}

// ExitFieldNames is called when production fieldNames is exited.
func (s *BasetronParserListener) ExitFieldNames(ctx *FieldNamesContext) {}

// EnterTyper is called when production typer is entered.
func (s *BasetronParserListener) EnterTyper(ctx *TyperContext) {}

// ExitTyper is called when production typer is exited.
func (s *BasetronParserListener) ExitTyper(ctx *TyperContext) {}

// EnterTypes is called when production types is entered.
func (s *BasetronParserListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BasetronParserListener) ExitTypes(ctx *TypesContext) {}

// EnterField is called when production field is entered.
func (s *BasetronParserListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BasetronParserListener) ExitField(ctx *FieldContext) {}

// EnterFieldOptions is called when production fieldOptions is entered.
func (s *BasetronParserListener) EnterFieldOptions(ctx *FieldOptionsContext) {}

// ExitFieldOptions is called when production fieldOptions is exited.
func (s *BasetronParserListener) ExitFieldOptions(ctx *FieldOptionsContext) {}

// EnterFieldOption is called when production fieldOption is entered.
func (s *BasetronParserListener) EnterFieldOption(ctx *FieldOptionContext) {}

// ExitFieldOption is called when production fieldOption is exited.
func (s *BasetronParserListener) ExitFieldOption(ctx *FieldOptionContext) {}

// EnterOneof is called when production oneof is entered.
func (s *BasetronParserListener) EnterOneof(ctx *OneofContext) {}

// ExitOneof is called when production oneof is exited.
func (s *BasetronParserListener) ExitOneof(ctx *OneofContext) {}

// EnterOneofField is called when production oneofField is entered.
func (s *BasetronParserListener) EnterOneofField(ctx *OneofFieldContext) {}

// ExitOneofField is called when production oneofField is exited.
func (s *BasetronParserListener) ExitOneofField(ctx *OneofFieldContext) {}

// EnterMapField is called when production mapField is entered.
func (s *BasetronParserListener) EnterMapField(ctx *MapFieldContext) {}

// ExitMapField is called when production mapField is exited.
func (s *BasetronParserListener) ExitMapField(ctx *MapFieldContext) {}

// EnterKeyType is called when production keyType is entered.
func (s *BasetronParserListener) EnterKeyType(ctx *KeyTypeContext) {}

// ExitKeyType is called when production keyType is exited.
func (s *BasetronParserListener) ExitKeyType(ctx *KeyTypeContext) {}

// EnterMessageType is called when production messageType is entered.
func (s *BasetronParserListener) EnterMessageType(ctx *MessageTypeContext) {}

// ExitMessageType is called when production messageType is exited.
func (s *BasetronParserListener) ExitMessageType(ctx *MessageTypeContext) {}

// EnterMessageOrEnumType is called when production messageOrEnumType is entered.
func (s *BasetronParserListener) EnterMessageOrEnumType(ctx *MessageOrEnumTypeContext) {}

// ExitMessageOrEnumType is called when production messageOrEnumType is exited.
func (s *BasetronParserListener) ExitMessageOrEnumType(ctx *MessageOrEnumTypeContext) {}

// EnterBool_lit is called when production bool_lit is entered.
func (s *BasetronParserListener) EnterBool_lit(ctx *Bool_litContext) {}

// ExitBool_lit is called when production bool_lit is exited.
func (s *BasetronParserListener) ExitBool_lit(ctx *Bool_litContext) {}

// EnterEmptyStatement is called when production emptyStatement is entered.
func (s *BasetronParserListener) EnterEmptyStatement(ctx *EmptyStatementContext) {}

// ExitEmptyStatement is called when production emptyStatement is exited.
func (s *BasetronParserListener) ExitEmptyStatement(ctx *EmptyStatementContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasetronParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasetronParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterPron is called when production pron is entered.
func (s *BasetronParserListener) EnterPron(ctx *PronContext) {}

// ExitPron is called when production pron is exited.
func (s *BasetronParserListener) ExitPron(ctx *PronContext) {}

// EnterPronSTR is called when production PronSTR is entered.
func (s *BasetronParserListener) EnterPronSTR(ctx *PronSTRContext) {}

// ExitPronSTR is called when production PronSTR is exited.
func (s *BasetronParserListener) ExitPronSTR(ctx *PronSTRContext) {}

// EnterPronOBJ is called when production PronOBJ is entered.
func (s *BasetronParserListener) EnterPronOBJ(ctx *PronOBJContext) {}

// ExitPronOBJ is called when production PronOBJ is exited.
func (s *BasetronParserListener) ExitPronOBJ(ctx *PronOBJContext) {}
