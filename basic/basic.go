package basic

import (
	"tutorial/basic/arrayslice"
	closure "tutorial/basic/closures"
	"tutorial/basic/deferfun"
	"tutorial/basic/export"
	formodule "tutorial/basic/for"
	"tutorial/basic/function"
	"tutorial/basic/funcvalue"
	"tutorial/basic/gostruct"
	ifmodule "tutorial/basic/if"
	"tutorial/basic/imports"
	loopfunction "tutorial/basic/loopFunction"
	mapgo "tutorial/basic/map"
	"tutorial/basic/multiresult"
	"tutorial/basic/namedresult"
	"tutorial/basic/packages"
	"tutorial/basic/pointer"
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
	pointer.PointerExample()
	gostruct.RunStructs()
	arrayslice.RunSliceBounds()
	arrayslice.RunMakeSlice()
	arrayslice.RunSliceSlice()
	arrayslice.SliceAppend()
	arrayslice.RunSliceRange()
	arrayslice.ShowPic()
	arrayslice.ShowPicSimple()
	mapgo.RunMapGo()
	mapgo.TestWordCount()
	funcvalue.RunFuncValue()
	closure.RunClosure()
	closure.RunFibonacci()
}
