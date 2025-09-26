package basic

import (
	"tutorial/basic/deferfun"
	"tutorial/basic/export"
	formodule "tutorial/basic/for"
	"tutorial/basic/function"
	ifmodule "tutorial/basic/if"
	"tutorial/basic/imports"
	loopfunction "tutorial/basic/loopFunction"
	"tutorial/basic/multiresult"
	"tutorial/basic/namedresult"
	"tutorial/basic/packages"
	"tutorial/basic/switchcase"
	"tutorial/basic/types"
	"tutorial/basic/variables"
)

func RunBasic() {
	// Run all the functions from the tutorial of basic part.

	// Packages, variables, and functions.
	imports.FnImportModule()
	export.FnExport()
	packages.Fn_packages()
	function.Fn_function()
	multiresult.Fn_multiresult()
	namedresult.Fn_namedresult()
	variables.Fn_variables()
	variables.FnVariablesWithInit()
	variables.FnShortVar()
	types.FnTypes()
	types.FnZeroValue()
	types.TypeConver()
	types.TypeInference()
	types.TypeConst()
	types.NumericConst()

	// Flow control statements: for, if, switch and defer.
	formodule.FnFor()
	formodule.ForClear()
	ifmodule.FnIf()
	ifmodule.IfPow()
	loopfunction.LoopFunction()
	loopfunction.TestDemo()
	switchcase.RunSwitchCase()
	deferfun.DeferExample()
}
