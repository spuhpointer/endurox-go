package atmi
/*
** JSON IPC Buffer support
**
** @file typed_json.go
**
** -----------------------------------------------------------------------------
** Enduro/X Middleware Platform for Distributed Transaction Processing
** Copyright (C) 2015, ATR Baltic, SIA. All Rights Reserved.
** This software is released under one of the following licenses:
** GPL or ATR Baltic's license for commercial use.
** -----------------------------------------------------------------------------
** GPL license:
**
** This program is free software; you can redistribute it and/or modify it under
** the terms of the GNU General Public License as published by the Free Software
** Foundation; either version 2 of the License, or (at your option) any later
** version.
**
** This program is distributed in the hope that it will be useful, but WITHOUT ANY
** WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
** PARTICULAR PURPOSE. See the GNU General Public License for more details.
**
** You should have received a copy of the GNU General Public License along with
** this program; if not, write to the Free Software Foundation, Inc., 59 Temple
** Place, Suite 330, Boston, MA 02111-1307 USA
**
** -----------------------------------------------------------------------------
** A commercial use license is available from ATR Baltic, SIA
** contact@atrbaltic.com
** -----------------------------------------------------------------------------
*/

/*
#cgo pkg-config: atmisrvinteg

#include <xatmi.h>
#include <string.h>
#include <stdlib.h>
#include <ubf.h>

*/
import "C"
import "unsafe"

//UBF Buffer
type TypedJSON struct {
	Buf *ATMIBuf
}

//Return The ATMI buffer to caller
func (u *TypedJSON) GetBuf() *ATMIBuf {
	return u.Buf
}

//Allocate new JSON buffer
//@param s - source string
func (ac *ATMICtx) NewJSON(b []byte) (*TypedJSON, ATMIError) {
	var buf TypedJSON

	c_val := C.CString(string(b))
	defer C.free(unsafe.Pointer(c_val))

	size := int64(C.strlen(c_val) + 1) /* 1 for EOS. */

	if ptr, err := ac.TpAlloc("JSON", "", size); nil != err {
		return nil, err
	} else {
		buf.Buf = ptr
		C.strcpy(buf.Buf.C_ptr, c_val)

		buf.Buf.TpSetCtxt(ac)

		return &buf, nil
	}
}

//Get the JSON Handler from ATMI Buffer
func (ac *ATMICtx) CastToJSON(abuf *ATMIBuf) (*TypedJSON, ATMIError) {
	var buf TypedJSON

	buf.Buf = abuf

	return &buf, nil
}

//Get the string value out from buffer
//@return JSON value
func (j *TypedJSON) GetJSONText() string {
	return C.GoString(j.Buf.C_ptr)
}

//Get JSON bytes..
func (j *TypedJSON) GetJSON() []byte {
	return []byte(C.GoString(j.Buf.C_ptr))
}

//Set JSON bytes
func (j *TypedJSON) SetJSON(b []byte) ATMIError {
	return j.SetJSONText(string(b))
}

//Set the string to the buffer
//@param str 	JSON value
func (j *TypedJSON) SetJSONText(gs string) ATMIError {
	c_val := C.CString(gs)
	defer C.free(unsafe.Pointer(c_val))

	new_size := int64(C.strlen(c_val) + 1) /* 1 for EOS. */

	if cur_size, err := j.Buf.Ctx.TpTypes(j.Buf, nil, nil); nil != err {
		return err
	} else {
		if cur_size >= new_size {
			C.strcpy(j.Buf.C_ptr, c_val)
		} else if err := j.Buf.TpRealloc(new_size); nil != err {
			return err
		} else {
			C.strcpy(j.Buf.C_ptr, c_val)
		}
	}

	return nil
}

///////////////////////////////////////////////////////////////////////////////////
// Wrappers for memory management
///////////////////////////////////////////////////////////////////////////////////

func (u *TypedJSON) TpRealloc(size int64) ATMIError {
	return u.Buf.TpRealloc(size)
}
