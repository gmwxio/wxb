// Code generated from tron_parser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // tron_parser

import "github.com/wxio/goantlr"

// Basetron_parserListener is a complete listener for a parse tree produced by tron_parser.
type Basetron_parserListener struct{}

var _ tron_parserListener = &Basetron_parserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Basetron_parserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Basetron_parserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Basetron_parserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Basetron_parserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProto is called when production proto is entered.
func (s *Basetron_parserListener) EnterProto(ctx *ProtoContext) {}

// ExitProto is called when production proto is exited.
func (s *Basetron_parserListener) ExitProto(ctx *ProtoContext) {}

// EnterTop_level_statement is called when production top_level_statement is entered.
func (s *Basetron_parserListener) EnterTop_level_statement(ctx *Top_level_statementContext) {}

// ExitTop_level_statement is called when production top_level_statement is exited.
func (s *Basetron_parserListener) ExitTop_level_statement(ctx *Top_level_statementContext) {}

// EnterSyntax is called when production syntax is entered.
func (s *Basetron_parserListener) EnterSyntax(ctx *SyntaxContext) {}

// ExitSyntax is called when production syntax is exited.
func (s *Basetron_parserListener) ExitSyntax(ctx *SyntaxContext) {}

// EnterExtend is called when production extend is entered.
func (s *Basetron_parserListener) EnterExtend(ctx *ExtendContext) {}

// ExitExtend is called when production extend is exited.
func (s *Basetron_parserListener) ExitExtend(ctx *ExtendContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *Basetron_parserListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *Basetron_parserListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterPackageStatement is called when production packageStatement is entered.
func (s *Basetron_parserListener) EnterPackageStatement(ctx *PackageStatementContext) {}

// ExitPackageStatement is called when production packageStatement is exited.
func (s *Basetron_parserListener) ExitPackageStatement(ctx *PackageStatementContext) {}

// EnterOption is called when production option is entered.
func (s *Basetron_parserListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *Basetron_parserListener) ExitOption(ctx *OptionContext) {}

// EnterOptionName is called when production optionName is entered.
func (s *Basetron_parserListener) EnterOptionName(ctx *OptionNameContext) {}

// ExitOptionName is called when production optionName is exited.
func (s *Basetron_parserListener) ExitOptionName(ctx *OptionNameContext) {}

// EnterMessage is called when production message is entered.
func (s *Basetron_parserListener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *Basetron_parserListener) ExitMessage(ctx *MessageContext) {}

// EnterMessageBody is called when production messageBody is entered.
func (s *Basetron_parserListener) EnterMessageBody(ctx *MessageBodyContext) {}

// ExitMessageBody is called when production messageBody is exited.
func (s *Basetron_parserListener) ExitMessageBody(ctx *MessageBodyContext) {}

// EnterEnumDefinition is called when production enumDefinition is entered.
func (s *Basetron_parserListener) EnterEnumDefinition(ctx *EnumDefinitionContext) {}

// ExitEnumDefinition is called when production enumDefinition is exited.
func (s *Basetron_parserListener) ExitEnumDefinition(ctx *EnumDefinitionContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *Basetron_parserListener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *Basetron_parserListener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterEnumField is called when production enumField is entered.
func (s *Basetron_parserListener) EnterEnumField(ctx *EnumFieldContext) {}

// ExitEnumField is called when production enumField is exited.
func (s *Basetron_parserListener) ExitEnumField(ctx *EnumFieldContext) {}

// EnterEnumValueOption is called when production enumValueOption is entered.
func (s *Basetron_parserListener) EnterEnumValueOption(ctx *EnumValueOptionContext) {}

// ExitEnumValueOption is called when production enumValueOption is exited.
func (s *Basetron_parserListener) ExitEnumValueOption(ctx *EnumValueOptionContext) {}

// EnterService is called when production service is entered.
func (s *Basetron_parserListener) EnterService(ctx *ServiceContext) {}

// ExitService is called when production service is exited.
func (s *Basetron_parserListener) ExitService(ctx *ServiceContext) {}

// EnterServiceBody is called when production serviceBody is entered.
func (s *Basetron_parserListener) EnterServiceBody(ctx *ServiceBodyContext) {}

// ExitServiceBody is called when production serviceBody is exited.
func (s *Basetron_parserListener) ExitServiceBody(ctx *ServiceBodyContext) {}

// EnterRpc is called when production rpc is entered.
func (s *Basetron_parserListener) EnterRpc(ctx *RpcContext) {}

// ExitRpc is called when production rpc is exited.
func (s *Basetron_parserListener) ExitRpc(ctx *RpcContext) {}

// EnterRpcParam is called when production rpcParam is entered.
func (s *Basetron_parserListener) EnterRpcParam(ctx *RpcParamContext) {}

// ExitRpcParam is called when production rpcParam is exited.
func (s *Basetron_parserListener) ExitRpcParam(ctx *RpcParamContext) {}

// EnterReserved is called when production reserved is entered.
func (s *Basetron_parserListener) EnterReserved(ctx *ReservedContext) {}

// ExitReserved is called when production reserved is exited.
func (s *Basetron_parserListener) ExitReserved(ctx *ReservedContext) {}

// EnterRanges is called when production ranges is entered.
func (s *Basetron_parserListener) EnterRanges(ctx *RangesContext) {}

// ExitRanges is called when production ranges is exited.
func (s *Basetron_parserListener) ExitRanges(ctx *RangesContext) {}

// EnterRangee is called when production rangee is entered.
func (s *Basetron_parserListener) EnterRangee(ctx *RangeeContext) {}

// ExitRangee is called when production rangee is exited.
func (s *Basetron_parserListener) ExitRangee(ctx *RangeeContext) {}

// EnterFieldNames is called when production fieldNames is entered.
func (s *Basetron_parserListener) EnterFieldNames(ctx *FieldNamesContext) {}

// ExitFieldNames is called when production fieldNames is exited.
func (s *Basetron_parserListener) ExitFieldNames(ctx *FieldNamesContext) {}

// EnterTyper is called when production typer is entered.
func (s *Basetron_parserListener) EnterTyper(ctx *TyperContext) {}

// ExitTyper is called when production typer is exited.
func (s *Basetron_parserListener) ExitTyper(ctx *TyperContext) {}

// EnterTypes is called when production types is entered.
func (s *Basetron_parserListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *Basetron_parserListener) ExitTypes(ctx *TypesContext) {}

// EnterField is called when production field is entered.
func (s *Basetron_parserListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *Basetron_parserListener) ExitField(ctx *FieldContext) {}

// EnterFieldOptions is called when production fieldOptions is entered.
func (s *Basetron_parserListener) EnterFieldOptions(ctx *FieldOptionsContext) {}

// ExitFieldOptions is called when production fieldOptions is exited.
func (s *Basetron_parserListener) ExitFieldOptions(ctx *FieldOptionsContext) {}

// EnterFieldOption is called when production fieldOption is entered.
func (s *Basetron_parserListener) EnterFieldOption(ctx *FieldOptionContext) {}

// ExitFieldOption is called when production fieldOption is exited.
func (s *Basetron_parserListener) ExitFieldOption(ctx *FieldOptionContext) {}

// EnterOneof is called when production oneof is entered.
func (s *Basetron_parserListener) EnterOneof(ctx *OneofContext) {}

// ExitOneof is called when production oneof is exited.
func (s *Basetron_parserListener) ExitOneof(ctx *OneofContext) {}

// EnterOneofField is called when production oneofField is entered.
func (s *Basetron_parserListener) EnterOneofField(ctx *OneofFieldContext) {}

// ExitOneofField is called when production oneofField is exited.
func (s *Basetron_parserListener) ExitOneofField(ctx *OneofFieldContext) {}

// EnterMapField is called when production mapField is entered.
func (s *Basetron_parserListener) EnterMapField(ctx *MapFieldContext) {}

// ExitMapField is called when production mapField is exited.
func (s *Basetron_parserListener) ExitMapField(ctx *MapFieldContext) {}

// EnterKeyType is called when production keyType is entered.
func (s *Basetron_parserListener) EnterKeyType(ctx *KeyTypeContext) {}

// ExitKeyType is called when production keyType is exited.
func (s *Basetron_parserListener) ExitKeyType(ctx *KeyTypeContext) {}

// EnterFullIdent is called when production fullIdent is entered.
func (s *Basetron_parserListener) EnterFullIdent(ctx *FullIdentContext) {}

// ExitFullIdent is called when production fullIdent is exited.
func (s *Basetron_parserListener) ExitFullIdent(ctx *FullIdentContext) {}

// EnterMessageName is called when production messageName is entered.
func (s *Basetron_parserListener) EnterMessageName(ctx *MessageNameContext) {}

// ExitMessageName is called when production messageName is exited.
func (s *Basetron_parserListener) ExitMessageName(ctx *MessageNameContext) {}

// EnterEnumName is called when production enumName is entered.
func (s *Basetron_parserListener) EnterEnumName(ctx *EnumNameContext) {}

// ExitEnumName is called when production enumName is exited.
func (s *Basetron_parserListener) ExitEnumName(ctx *EnumNameContext) {}

// EnterMessageOrEnumName is called when production messageOrEnumName is entered.
func (s *Basetron_parserListener) EnterMessageOrEnumName(ctx *MessageOrEnumNameContext) {}

// ExitMessageOrEnumName is called when production messageOrEnumName is exited.
func (s *Basetron_parserListener) ExitMessageOrEnumName(ctx *MessageOrEnumNameContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *Basetron_parserListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *Basetron_parserListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterOneofName is called when production oneofName is entered.
func (s *Basetron_parserListener) EnterOneofName(ctx *OneofNameContext) {}

// ExitOneofName is called when production oneofName is exited.
func (s *Basetron_parserListener) ExitOneofName(ctx *OneofNameContext) {}

// EnterMapName is called when production mapName is entered.
func (s *Basetron_parserListener) EnterMapName(ctx *MapNameContext) {}

// ExitMapName is called when production mapName is exited.
func (s *Basetron_parserListener) ExitMapName(ctx *MapNameContext) {}

// EnterServiceName is called when production serviceName is entered.
func (s *Basetron_parserListener) EnterServiceName(ctx *ServiceNameContext) {}

// ExitServiceName is called when production serviceName is exited.
func (s *Basetron_parserListener) ExitServiceName(ctx *ServiceNameContext) {}

// EnterRpcName is called when production rpcName is entered.
func (s *Basetron_parserListener) EnterRpcName(ctx *RpcNameContext) {}

// ExitRpcName is called when production rpcName is exited.
func (s *Basetron_parserListener) ExitRpcName(ctx *RpcNameContext) {}

// EnterMessageType is called when production messageType is entered.
func (s *Basetron_parserListener) EnterMessageType(ctx *MessageTypeContext) {}

// ExitMessageType is called when production messageType is exited.
func (s *Basetron_parserListener) ExitMessageType(ctx *MessageTypeContext) {}

// EnterMessageOrEnumType is called when production messageOrEnumType is entered.
func (s *Basetron_parserListener) EnterMessageOrEnumType(ctx *MessageOrEnumTypeContext) {}

// ExitMessageOrEnumType is called when production messageOrEnumType is exited.
func (s *Basetron_parserListener) ExitMessageOrEnumType(ctx *MessageOrEnumTypeContext) {}

// EnterBool_lit is called when production bool_lit is entered.
func (s *Basetron_parserListener) EnterBool_lit(ctx *Bool_litContext) {}

// ExitBool_lit is called when production bool_lit is exited.
func (s *Basetron_parserListener) ExitBool_lit(ctx *Bool_litContext) {}

// EnterEmptyStatement is called when production emptyStatement is entered.
func (s *Basetron_parserListener) EnterEmptyStatement(ctx *EmptyStatementContext) {}

// ExitEmptyStatement is called when production emptyStatement is exited.
func (s *Basetron_parserListener) ExitEmptyStatement(ctx *EmptyStatementContext) {}

// EnterConstant is called when production constant is entered.
func (s *Basetron_parserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *Basetron_parserListener) ExitConstant(ctx *ConstantContext) {}

// EnterPron is called when production pron is entered.
func (s *Basetron_parserListener) EnterPron(ctx *PronContext) {}

// ExitPron is called when production pron is exited.
func (s *Basetron_parserListener) ExitPron(ctx *PronContext) {}

// EnterPronSTR is called when production PronSTR is entered.
func (s *Basetron_parserListener) EnterPronSTR(ctx *PronSTRContext) {}

// ExitPronSTR is called when production PronSTR is exited.
func (s *Basetron_parserListener) ExitPronSTR(ctx *PronSTRContext) {}

// EnterPronOBJ is called when production PronOBJ is entered.
func (s *Basetron_parserListener) EnterPronOBJ(ctx *PronOBJContext) {}

// ExitPronOBJ is called when production PronOBJ is exited.
func (s *Basetron_parserListener) ExitPronOBJ(ctx *PronOBJContext) {}
