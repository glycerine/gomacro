/*
 * gomacro - A Go intepreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * globals.go
 *
 *  Created on: Feb 19, 2017
 *      Author: Massimiliano Ghilardi
 */

package interpreter

import (
	"go/ast"
	r "reflect"
)

type CallStack struct {
	Frames []CallFrame
}

type CallFrame struct {
	FuncEnv       *Env
	InnerEnv      *Env          // innermost Env
	CurrentCall   *ast.CallExpr // call currently in progress
	defers        []func()
	panick        interface{} // current panic
	panicking     bool
	runningDefers bool
}

type Builtin struct {
	Exec   func(env *Env, args []ast.Expr) (r.Value, []r.Value)
	ArgNum int // if negative, do not check
}

type Function struct {
	Exec   func(env *Env, args []r.Value) (r.Value, []r.Value)
	ArgNum int // if negative, do not check
}

type Macro struct {
	Closure func(args []r.Value) (results []r.Value)
	ArgNum  int
}

type TypedValue struct {
	Type  r.Type
	Value r.Value
}

/**
 * inside Methods, each string is the method name
 * and each TypedValue is {
 *   Type: the method signature, i.e. the type of a func() *without* the receiver (to allow comparison with Interface methods)
 *   Value: the method implementation, i.e. a func() whose first argument is the receiver,
 * }
 */
type Methods map[string]TypedValue

/**
 * Interface is the interpreted version of Golang interface values.
 * Each Interface contains {
 *   Type:  the interface type. returned by Env.evalInterface(), i.e. the type of a struct {  interface{}; functions... }
 *   Value: the datum implementing the interface. Value.Type() must be its concrete type, i.e. == r.TypeOf(Value.Interface())
 * }
 */
type Interface TypedValue
