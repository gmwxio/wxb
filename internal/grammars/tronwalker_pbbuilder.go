package tron

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	antlr "github.com/wxio/goantlr"
	"github.com/wxio/wxb/internal/ctree"
	"github.com/wxio/wxb/internal/grammars/walker"
)

type walkTronPB struct {
	File string `type:"arg" help:"proto file" predict:"files"`
}

func (et *walkTronPB) Run() error {
	by, err := ioutil.ReadFile(et.File)
	if err != nil {
		return err
	}
	tr, _, _, err := BuildTronAST(string(by))
	if err != nil {
		return err
	}
	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walker.NewTronWalker(tts)
	p.SetTokenStream(tts)
	// p.RemoveErrorListeners()
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tbl := &WalkPBListener{}
	tbl.debug = true
	ret := p.Proto()
	antlr.ParseTreeWalkerDefault.Walk(tbl, ret)
	fmt.Printf("FDP '%v'\n", tbl.protoFile)
	return tbl.err
}

type WalkPBListener struct {
	*walker.BaseTronWalkerListener
	//
	protoDS   *descriptor.FileDescriptorSet
	protoFile *descriptor.FileDescriptorProto
	indent    string
	debug     bool
	warning   string
	err       error
}

func (tbl *WalkPBListener) FileDescriptorProto() *descriptor.FileDescriptorProto {
	return tbl.protoFile
}

func (tbl *WalkPBListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if tbl.debug {
		fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
	}
	if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
		tbl.warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
	} else {
		tbl.err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
	}
	// panic("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

func (tbl *WalkPBListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAmbiguity")
	if tbl.debug {
		fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
	}
}

func (tbl *WalkPBListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAttemptingFullContext")
	if tbl.debug {
		fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
	}
}

func (tbl *WalkPBListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	panic("ReportContextSensitivity")
}

// VisitTerminal is called when a terminal node is visited.
func (s *WalkPBListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *WalkPBListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *WalkPBListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *WalkPBListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProto is called when production proto is entered.
func (s *WalkPBListener) EnterProto(ctx *walker.ProtoContext) {
	s.protoDS = &descriptor.FileDescriptorSet{}
	s.protoFile = &descriptor.FileDescriptorProto{}
}

// ExitProto is called when production proto is exited.
func (s *WalkPBListener) ExitProto(ctx *walker.ProtoContext) {}

// EnterSyntaxNode is called when production SyntaxNode is entered.
func (s *WalkPBListener) EnterSyntaxNode(ctx *walker.SyntaxNodeContext) {}

// ExitSyntaxNode is called when production SyntaxNode is exited.
func (s *WalkPBListener) ExitSyntaxNode(ctx *walker.SyntaxNodeContext) {}

// EnterPkgNode is called when production PkgNode is entered.
func (s *WalkPBListener) EnterPkgNode(ctx *walker.PkgNodeContext) {
	p := strings.Join(ctx.GetStart().(*MessageNode).Name, ".")
	s.protoFile.Package = &p
}

// ExitPkgNode is called when production PkgNode is exited.
func (s *WalkPBListener) ExitPkgNode(ctx *walker.PkgNodeContext) {}

// EnterImportNode is called when production ImportNode is entered.
func (s *WalkPBListener) EnterImportNode(ctx *walker.ImportNodeContext) {
	imp := ctx.GetStart().(*ImportNode)
	s.protoFile.Dependency = append(s.protoFile.Dependency, imp.Path)
}

// ExitImportNode is called when production ImportNode is exited.
func (s *WalkPBListener) ExitImportNode(ctx *walker.ImportNodeContext) {}

// EnterMsgNode is called when production MsgNode is entered.
func (s *WalkPBListener) EnterMsgNode(ctx *walker.MsgNodeContext) {}

// ExitMsgNode is called when production MsgNode is exited.
func (s *WalkPBListener) ExitMsgNode(ctx *walker.MsgNodeContext) {}

// EnterEnumNode is called when production EnumNode is entered.
func (s *WalkPBListener) EnterEnumNode(ctx *walker.EnumNodeContext) {}

// ExitEnumNode is called when production EnumNode is exited.
func (s *WalkPBListener) ExitEnumNode(ctx *walker.EnumNodeContext) {}

// EnterSvcNode is called when production SvcNode is entered.
func (s *WalkPBListener) EnterSvcNode(ctx *walker.SvcNodeContext) {}

// ExitSvcNode is called when production SvcNode is exited.
func (s *WalkPBListener) ExitSvcNode(ctx *walker.SvcNodeContext) {}

// EnterExtNode is called when production ExtNode is entered.
func (s *WalkPBListener) EnterExtNode(ctx *walker.ExtNodeContext) {}

// ExitExtNode is called when production ExtNode is exited.
func (s *WalkPBListener) ExitExtNode(ctx *walker.ExtNodeContext) {}

// EnterMsgOptNode is called when production MsgOptNode is entered.
func (s *WalkPBListener) EnterMsgOptNode(ctx *walker.MsgOptNodeContext) {}

// ExitMsgOptNode is called when production MsgOptNode is exited.
func (s *WalkPBListener) ExitMsgOptNode(ctx *walker.MsgOptNodeContext) {}

// EnterMsgFldNode is called when production MsgFldNode is entered.
func (s *WalkPBListener) EnterMsgFldNode(ctx *walker.MsgFldNodeContext) {}

// ExitMsgFldNode is called when production MsgFldNode is exited.
func (s *WalkPBListener) ExitMsgFldNode(ctx *walker.MsgFldNodeContext) {}

// EnterMsgMsgNode is called when production MsgMsgNode is entered.
func (s *WalkPBListener) EnterMsgMsgNode(ctx *walker.MsgMsgNodeContext) {}

// ExitMsgMsgNode is called when production MsgMsgNode is exited.
func (s *WalkPBListener) ExitMsgMsgNode(ctx *walker.MsgMsgNodeContext) {}

// EnterMsgEnumNode is called when production MsgEnumNode is entered.
func (s *WalkPBListener) EnterMsgEnumNode(ctx *walker.MsgEnumNodeContext) {}

// ExitMsgEnumNode is called when production MsgEnumNode is exited.
func (s *WalkPBListener) ExitMsgEnumNode(ctx *walker.MsgEnumNodeContext) {}

// EnterMsgOneofNode is called when production MsgOneofNode is entered.
func (s *WalkPBListener) EnterMsgOneofNode(ctx *walker.MsgOneofNodeContext) {}

// ExitMsgOneofNode is called when production MsgOneofNode is exited.
func (s *WalkPBListener) ExitMsgOneofNode(ctx *walker.MsgOneofNodeContext) {}

// EnterOneofOptNode is called when production OneofOptNode is entered.
func (s *WalkPBListener) EnterOneofOptNode(ctx *walker.OneofOptNodeContext) {}

// ExitOneofOptNode is called when production OneofOptNode is exited.
func (s *WalkPBListener) ExitOneofOptNode(ctx *walker.OneofOptNodeContext) {}

// EnterOneofFldNode is called when production OneofFldNode is entered.
func (s *WalkPBListener) EnterOneofFldNode(ctx *walker.OneofFldNodeContext) {}

// ExitOneofFldNode is called when production OneofFldNode is exited.
func (s *WalkPBListener) ExitOneofFldNode(ctx *walker.OneofFldNodeContext) {}

// EnterEnumOptNode is called when production EnumOptNode is entered.
func (s *WalkPBListener) EnterEnumOptNode(ctx *walker.EnumOptNodeContext) {}

// ExitEnumOptNode is called when production EnumOptNode is exited.
func (s *WalkPBListener) ExitEnumOptNode(ctx *walker.EnumOptNodeContext) {}

// EnterEnumValNode is called when production EnumValNode is entered.
func (s *WalkPBListener) EnterEnumValNode(ctx *walker.EnumValNodeContext) {}

// ExitEnumValNode is called when production EnumValNode is exited.
func (s *WalkPBListener) ExitEnumValNode(ctx *walker.EnumValNodeContext) {}

// EnterSvcOptNode is called when production SvcOptNode is entered.
func (s *WalkPBListener) EnterSvcOptNode(ctx *walker.SvcOptNodeContext) {}

// ExitSvcOptNode is called when production SvcOptNode is exited.
func (s *WalkPBListener) ExitSvcOptNode(ctx *walker.SvcOptNodeContext) {}

// EnterSvcRpcNode is called when production SvcRpcNode is entered.
func (s *WalkPBListener) EnterSvcRpcNode(ctx *walker.SvcRpcNodeContext) {}

// ExitSvcRpcNode is called when production SvcRpcNode is exited.
func (s *WalkPBListener) ExitSvcRpcNode(ctx *walker.SvcRpcNodeContext) {}

// EnterExtOptNode is called when production ExtOptNode is entered.
func (s *WalkPBListener) EnterExtOptNode(ctx *walker.ExtOptNodeContext) {}

// ExitExtOptNode is called when production ExtOptNode is exited.
func (s *WalkPBListener) ExitExtOptNode(ctx *walker.ExtOptNodeContext) {}

// EnterExtFldNode is called when production ExtFldNode is entered.
func (s *WalkPBListener) EnterExtFldNode(ctx *walker.ExtFldNodeContext) {}

// ExitExtFldNode is called when production ExtFldNode is exited.
func (s *WalkPBListener) ExitExtFldNode(ctx *walker.ExtFldNodeContext) {}
