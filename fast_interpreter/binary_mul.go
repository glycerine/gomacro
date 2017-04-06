/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
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
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * binary_mul.go
 *
 *  Created on Apr 02, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/token"
	r "reflect"
)

func (c *Comp) Mul(op token.Token, xe *Expr, ye *Expr) *Expr {
	xc, yc := xe.Const(), ye.Const()
	toSameFuncType(op, xe, ye)
	if !isCategory(xe.Type.Kind(), r.Int, r.Uint, r.Float64, r.Complex128) {
		return c.invalidBinaryExpr(op, xe, ye)
	}
	// if both x and y are constants, BinaryExpr will invoke EvalConst()
	// on our return value. no need to optimize that.
	var fun I
	if xc == yc {
		x, y := xe.Fun, ye.Fun
		switch x := x.(type) {
		case func(*Env) int:
			y := y.(func(*Env) int)
			fun = func(env *Env) int {
				return x(env) * y(env)
			}
		case func(*Env) int8:
			y := y.(func(*Env) int8)
			fun = func(env *Env) int8 {
				return x(env) * y(env)
			}
		case func(*Env) int16:
			y := y.(func(*Env) int16)
			fun = func(env *Env) int16 {
				return x(env) * y(env)
			}
		case func(*Env) int32:
			y := y.(func(*Env) int32)
			fun = func(env *Env) int32 {
				return x(env) * y(env)
			}
		case func(*Env) int64:
			y := y.(func(*Env) int64)
			fun = func(env *Env) int64 {
				return x(env) * y(env)
			}
		case func(*Env) uint:
			y := y.(func(*Env) uint)
			fun = func(env *Env) uint {
				return x(env) * y(env)
			}
		case func(*Env) uint8:
			y := y.(func(*Env) uint8)
			fun = func(env *Env) uint8 {
				return x(env) * y(env)
			}
		case func(*Env) uint16:
			y := y.(func(*Env) uint16)
			fun = func(env *Env) uint16 {
				return x(env) * y(env)
			}
		case func(*Env) uint32:
			y := y.(func(*Env) uint32)
			fun = func(env *Env) uint32 {
				return x(env) * y(env)
			}
		case func(*Env) uint64:
			y := y.(func(*Env) uint64)
			fun = func(env *Env) uint64 {
				return x(env) * y(env)
			}
		case func(*Env) uintptr:
			y := y.(func(*Env) uintptr)
			fun = func(env *Env) uintptr {
				return x(env) * y(env)
			}
		case func(*Env) float32:
			y := y.(func(*Env) float32)
			fun = func(env *Env) float32 {
				return x(env) * y(env)
			}
		case func(*Env) float64:
			y := y.(func(*Env) float64)
			fun = func(env *Env) float64 {
				return x(env) * y(env)
			}
		case func(*Env) complex64:
			y := y.(func(*Env) complex64)
			fun = func(env *Env) complex64 {
				return x(env) * y(env)
			}
		case func(*Env) complex128:
			y := y.(func(*Env) complex128)
			fun = func(env *Env) complex128 {
				return x(env) * y(env)
			}
		default:
			return c.invalidBinaryExpr(op, xe, ye)
		}
	} else if yc {
		x := xe.Fun
		y := ye.Value
		if isLiteralNumber(y, 1) {
			return xe
		}
		switch x := x.(type) {
		case func(*Env) int:
			y := y.(int)
			fun = func(env *Env) int {
				return x(env) * y
			}
		case func(*Env) int8:
			y := y.(int8)
			fun = func(env *Env) int8 {
				return x(env) * y
			}
		case func(*Env) int16:
			y := y.(int16)
			fun = func(env *Env) int16 {
				return x(env) * y
			}
		case func(*Env) int32:
			y := y.(int32)
			fun = func(env *Env) int32 {
				return x(env) * y
			}
		case func(*Env) int64:
			y := y.(int64)
			fun = func(env *Env) int64 {
				return x(env) * y
			}
		case func(*Env) uint:
			y := y.(uint)
			fun = func(env *Env) uint {
				return x(env) * y
			}
		case func(*Env) uint8:
			y := y.(uint8)
			fun = func(env *Env) uint8 {
				return x(env) * y
			}
		case func(*Env) uint16:
			y := y.(uint16)
			fun = func(env *Env) uint16 {
				return x(env) * y
			}
		case func(*Env) uint32:
			y := y.(uint32)
			fun = func(env *Env) uint32 {
				return x(env) * y
			}
		case func(*Env) uint64:
			y := y.(uint64)
			fun = func(env *Env) uint64 {
				return x(env) * y
			}
		case func(*Env) uintptr:
			y := y.(uintptr)
			fun = func(env *Env) uintptr {
				return x(env) * y
			}
		case func(*Env) float32:
			y := y.(float32)
			fun = func(env *Env) float32 {
				return x(env) * y
			}
		case func(*Env) float64:
			y := y.(float64)
			fun = func(env *Env) float64 {
				return x(env) * y
			}
		case func(*Env) complex64:
			y := y.(complex64)
			fun = func(env *Env) complex64 {
				return x(env) * y
			}
		case func(*Env) complex128:
			y := y.(complex128)
			fun = func(env *Env) complex128 {
				return x(env) * y
			}
		default:
			return c.invalidBinaryExpr(op, xe, ye)
		}
	} else {
		x := xe.Value
		y := ye.Fun
		if isLiteralNumber(x, 1) {
			return ye
		}
		switch y := y.(type) {
		case func(*Env) int:
			x := x.(int)
			fun = func(env *Env) int {
				return x * y(env)
			}
		case func(*Env) int8:
			x := x.(int8)
			fun = func(env *Env) int8 {
				return x * y(env)
			}
		case func(*Env) int16:
			x := x.(int16)
			fun = func(env *Env) int16 {
				return x * y(env)
			}
		case func(*Env) int32:
			x := x.(int32)
			fun = func(env *Env) int32 {
				return x * y(env)
			}
		case func(*Env) int64:
			x := x.(int64)
			fun = func(env *Env) int64 {
				return x * y(env)
			}
		case func(*Env) uint:
			x := x.(uint)
			fun = func(env *Env) uint {
				return x * y(env)
			}
		case func(*Env) uint8:
			x := x.(uint8)
			fun = func(env *Env) uint8 {
				return x * y(env)
			}
		case func(*Env) uint16:
			x := x.(uint16)
			fun = func(env *Env) uint16 {
				return x * y(env)
			}
		case func(*Env) uint32:
			x := x.(uint32)
			fun = func(env *Env) uint32 {
				return x * y(env)
			}
		case func(*Env) uint64:
			x := x.(uint64)
			fun = func(env *Env) uint64 {
				return x * y(env)
			}
		case func(*Env) uintptr:
			x := x.(uintptr)
			fun = func(env *Env) uintptr {
				return x * y(env)
			}
		case func(*Env) float32:
			x := x.(float32)
			fun = func(env *Env) float32 {
				return x * y(env)
			}
		case func(*Env) float64:
			x := x.(float64)
			fun = func(env *Env) float64 {
				return x * y(env)
			}
		case func(*Env) complex64:
			x := x.(complex64)
			fun = func(env *Env) complex64 {
				return x * y(env)
			}
		case func(*Env) complex128:
			x := x.(complex128)
			fun = func(env *Env) complex128 {
				return x * y(env)
			}
		default:
			return c.invalidBinaryExpr(op, xe, ye)
		}
	}
	return ExprFun(xe.Type, fun)
}
