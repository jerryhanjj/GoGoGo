package basic

import (
	"tutorial/basic/export"
	"tutorial/basic/function"
	"tutorial/basic/imports"
	"tutorial/basic/multiresult"
	"tutorial/basic/namedresult"
	"tutorial/basic/packages"
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

}
