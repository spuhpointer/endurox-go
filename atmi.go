package atmi

/*
#cgo LDFLAGS: -latmi -lrt -lm -lubf -lnstd -ldl

#include <ndebug.h>
#include <xatmi.h>
#include <string.h>
#include <stdlib.h>

// Wrapper for TPNIT
static int go_tpinit(void) {
	return tpinit(NULL);
}

//ATMI library error code
static int go_tperrno(void) {
	return tperrno;
}

//Standard library error code
static int go_Nerror(void) {
	return Nerror;
}


static void free_string(char* s) { free(s); }
static char * malloc_string(int size) { return malloc(size); }


static void go_param_to_tpqctl(
		TPQCTL *qc,
		long *ctl_flags,
		long *ctl_deq_time,
		long *ctl_priority,
		long *ctl_diagnostic,
		char *ctl_diagmsg,
		char *ctl_msgid,
		char *ctl_corrid,
		char *ctl_replyqueue,
		char *ctl_failurequeue,
		char *ctl_cltid,
		long *ctl_urcode,
		long *ctl_appkey,
		long *ctl_delivery_qos,
		long *ctl_reply_qos,
		long *ctl_exp_time)
{
	qc->flags = *ctl_flags;
	qc->deq_time = *ctl_deq_time;
	qc->priority = *ctl_priority;
	qc->diagnostic = *ctl_diagnostic;
	strcpy(qc->diagmsg, ctl_diagmsg);
	memcpy(qc->msgid, ctl_msgid, TMMSGIDLEN);
	memcpy(qc->corrid, ctl_corrid, TMCORRIDLEN);
	strcpy(qc->replyqueue, ctl_replyqueue);
	strcpy(qc->failurequeue, ctl_failurequeue);
	strcpy(qc->cltid.clientdata, ctl_cltid);
	qc->urcode = *ctl_urcode;
	qc->appkey = *ctl_appkey;
	qc->delivery_qos = *ctl_delivery_qos;
	qc->reply_qos = *ctl_reply_qos;
	qc->exp_time = *ctl_exp_time;
}

static void go_tpqctl_to_param(
		TPQCTL *qc,
		long *ctl_flags,
		long *ctl_deq_time,
		long *ctl_priority,
		long *ctl_diagnostic,
		char *ctl_diagmsg,
		char *ctl_msgid,
		char *ctl_corrid,
		char *ctl_replyqueue,
		char *ctl_failurequeue,
		char *ctl_cltid,
		long *ctl_urcode,
		long *ctl_appkey,
		long *ctl_delivery_qos,
		long *ctl_reply_qos,
		long *ctl_exp_time)
{
	*ctl_flags = qc->flags;
	*ctl_deq_time = qc->deq_time;
	*ctl_priority = qc->priority;
	*ctl_diagnostic = qc->diagnostic;
	strcpy(ctl_diagmsg, qc->diagmsg);
	memcpy(ctl_msgid, qc->msgid, TMMSGIDLEN);
	memcpy(ctl_corrid, qc->corrid, TMCORRIDLEN);
	strcpy(ctl_replyqueue, qc->replyqueue);
	strcpy(ctl_failurequeue, qc->failurequeue);
	strcpy(ctl_cltid, qc->cltid.clientdata);
	qc->urcode = *ctl_urcode;
	qc->appkey = *ctl_appkey;
	qc->delivery_qos = *ctl_delivery_qos;
	qc->reply_qos = *ctl_reply_qos;
	qc->exp_time = *ctl_exp_time;
}

static int go_tpenqueue (char *qspace, char *qname, char *data, long len, long flags,
		long *ctl_flags,
		long *ctl_deq_time,
		long *ctl_priority,
		long *ctl_diagnostic,
		char *ctl_diagmsg,
		char *ctl_msgid,
		char *ctl_corrid,
		char *ctl_replyqueue,
		char *ctl_failurequeue,
		char *ctl_cltid,
		long *ctl_urcode,
		long *ctl_appkey,
		long *ctl_delivery_qos,
		long *ctl_reply_qos,
		long *ctl_exp_time
)
{
	int ret;
	TPQCTL qc;
	memset(&qc, 0, sizeof(qc));

	go_param_to_tpqctl(&qc,
			ctl_flags,
			ctl_deq_time,
			ctl_priority,
			ctl_diagnostic,
			ctl_diagmsg,
			ctl_msgid,
			ctl_corrid,
			ctl_replyqueue,
			ctl_failurequeue,
			ctl_cltid,
			ctl_urcode,
			ctl_appkey,
			ctl_delivery_qos,
			ctl_reply_qos,
			ctl_exp_time);

	ret = tpenqueue (qspace, qname, &qc, data, len, flags);

	go_tpqctl_to_param(&qc,
			ctl_flags,
			ctl_deq_time,
			ctl_priority,
			ctl_diagnostic,
			ctl_diagmsg,
			ctl_msgid,
			ctl_corrid,
			ctl_replyqueue,
			ctl_failurequeue,
			ctl_cltid,
			ctl_urcode,
			ctl_appkey,
			ctl_delivery_qos,
			ctl_reply_qos,
			ctl_exp_time);

	return ret;
}

static int go_tpdequeue (char *qspace, char *qname, char **data, long *len, long flags,
		long *ctl_flags,
		long *ctl_deq_time,
		long *ctl_priority,
		long *ctl_diagnostic,
		char *ctl_diagmsg,
		char *ctl_msgid,
		char *ctl_corrid,
		char *ctl_replyqueue,
		char *ctl_failurequeue,
		char *ctl_cltid,
		long *ctl_urcode,
		long *ctl_appkey,
		long *ctl_delivery_qos,
		long *ctl_reply_qos,
		long *ctl_exp_time
)
{
	int ret;
	TPQCTL qc;
	memset(&qc, 0, sizeof(qc));

	go_param_to_tpqctl(&qc,
			ctl_flags,
			ctl_deq_time,
			ctl_priority,
			ctl_diagnostic,
			ctl_diagmsg,
			ctl_msgid,
			ctl_corrid,
			ctl_replyqueue,
			ctl_failurequeue,
			ctl_cltid,
			ctl_urcode,
			ctl_appkey,
			ctl_delivery_qos,
			ctl_reply_qos,
			ctl_exp_time);

	ret = tpdequeue (qspace, qname, &qc, data, len, flags);

	go_tpqctl_to_param(&qc,
			ctl_flags,
			ctl_deq_time,
			ctl_priority,
			ctl_diagnostic,
			ctl_diagmsg,
			ctl_msgid,
			ctl_corrid,
			ctl_replyqueue,
			ctl_failurequeue,
			ctl_cltid,
			ctl_urcode,
			ctl_appkey,
			ctl_delivery_qos,
			ctl_reply_qos,
			ctl_exp_time);

	return ret;
}

*/
import "C"
import "unsafe"
import "fmt"
import "runtime"

//TODO: Think about runtime.SetFinalizer - might be usable for ATMI buffer free
//      and for UBF expression dealloc

/*
 * SUCCEED/FAIL flags
 */
const (
	SUCCEED = 0
	FAIL    = -1
)

/*
 * List of ATMI Error codes
 */
const (
	TPMINVAL      = 0
	TPEABORT      = 1
	TPEBADDESC    = 2
	TPEBLOCK      = 3
	TPEINVAL      = 4
	TPELIMIT      = 5
	TPENOENT      = 6
	TPEOS         = 7
	TPEPERM       = 8
	TPEPROTO      = 9
	TPESVCERR     = 10
	TPESVCFAIL    = 11
	TPESYSTEM     = 12
	TPETIME       = 13
	TPETRAN       = 14
	TPGOTSIG      = 15
	TPERMERR      = 16
	TPEITYPE      = 17
	TPEOTYPE      = 18
	TPERELEASE    = 19
	TPEHAZARD     = 20
	TPEHEURISTIC  = 21
	TPEEVENT      = 22
	TPEMATCH      = 23
	TPEDIAGNOSTIC = 24
	TPEMIB        = 25
	TPINITFAIL    = 30
	TPMAXVAL      = 31
)

/*
 * flag bits for C language xatmi routines
 */
const (
	TPNOBLOCK     = 0x00000001
	TPSIGRSTRT    = 0x00000002
	TPNOREPLY     = 0x00000004
	TPNOTRAN      = 0x00000008
	TPTRAN        = 0x00000010
	TPNOTIME      = 0x00000020
	TPGETANY      = 0x00000080
	TPNOCHANGE    = 0x00000100
	TPCONV        = 0x00000400
	TPSENDONLY    = 0x00000800
	TPRECVONLY    = 0x00001000
	TPTRANSUSPEND = 0x00040000 /* Suspend current transaction */
)

/*
 * values for rval in tpreturn
 */
const (
	TPFAIL    = 0x0001
	TPSUCCESS = 0x0002
)

/*
 * Max message size (int bytes)
 */
const (
	ATMI_MSG_MAX_SIZE = 65536
)

/*
 * TPQCTL.flags flags
 */
const (
	TPNOFLAGS         = 0x00000
	TPQCORRID         = 0x00001  /* set/get correlation id */
	TPQFAILUREQ       = 0x00002  /* set/get failure queue */
	TPQBEFOREMSGID    = 0x00004  /* RFU, enqueue before message id */
	TPQGETBYMSGIDOLD  = 0x00008  /* RFU, deprecated */
	TPQMSGID          = 0x00010  /* get msgid of enq/deq message */
	TPQPRIORITY       = 0x00020  /* set/get message priority */
	TPQTOP            = 0x00040  /* RFU, enqueue at queue top */
	TPQWAIT           = 0x00080  /* RFU, wait for dequeuing */
	TPQREPLYQ         = 0x00100  /* set/get reply queue */
	TPQTIME_ABS       = 0x00200  /* RFU, set absolute time */
	TPQTIME_REL       = 0x00400  /* RFU, set absolute time */
	TPQGETBYCORRIDOLD = 0x00800  /* deprecated */
	TPQPEEK           = 0x01000  /* peek */
	TPQDELIVERYQOS    = 0x02000  /* RFU, delivery quality of service */
	TPQREPLYQOS       = 0x04000  /* RFU, reply message quality of service */
	TPQEXPTIME_ABS    = 0x08000  /* RFU, absolute expiration time */
	TPQEXPTIME_REL    = 0x10000  /* RFU, relative expiration time */
	TPQEXPTIME_NONE   = 0x20000  /* RFU, never expire */
	TPQGETBYMSGID     = 0x40008  /* dequeue by msgid */
	TPQGETBYCORRID    = 0x80800  /* dequeue by corrid */
	TPQASYNC          = 0x100000 /* Async complete */
)

/*
 * Values for TQPCTL.diagnostic
 */
const (
	MEINVAL     = -1
	MEBADRMID   = -2
	MENOTOPEN   = -3
	METRAN      = -4
	MEBADMSGID  = -5
	MESYSTEM    = -6
	MEOS        = -7
	MEABORTED   = -8
	MENOTA      = -8 /* QMEABORTED */
	MEPROTO     = -9
	MEBADQUEUE  = -10
	MENOMSG     = -11
	MEINUSE     = -12
	MENOSPACE   = -13
	MERELEASE   = -14
	MEINVHANDLE = -15
	MESHARE     = -16
)

/*
 * Q constants
 */
const (
	TMMSGIDLEN       = 32
	TMCORRIDLEN      = 32
	TMQNAMELEN       = 15
	NDRX_MAX_ID_SIZE = 96
)

/*
 * Log levels for TPLOG (corresponding to ndebug.h)
 */
const (
	LOG_ALWAYS = 1
	LOG_ERROR  = 2
	LOG_WARN   = 3
	LOG_INFO   = 4
	LOG_DEBUG  = 5
	LOG_DUMP   = 6
)

/*
 * Logging facilites
 */
const (
	LOG_FACILITY_NDRX       = 0x00001 /* settings for ATMI logging             */
	LOG_FACILITY_UBF        = 0x00002 /* settings for UBF logging              */
	LOG_FACILITY_TP         = 0x00004 /* settings for TP logging               */
	LOG_FACILITY_TP_THREAD  = 0x00008 /* settings for TP, thread based logging */
	LOG_FACILITY_TP_REQUEST = 0x00010 /* Request logging, thread based         */
)

/*
 * Transaction ID type
 */
type TPTRANID struct {
	c_tptranid C.TPTRANID
}

/*
 * Server context data (used for server's main thread
 * switching taks to worker thread)
 */
type TPSRVCTXDATA struct {
	c_ptr *C.char
}

/*
 * Event controll struct
 */
type TPEVCTL struct {
	flags int64
	name1 string
	name2 string
}

/*
 * Queue control structure
 */
type TPQCTL struct {
	flags        int64             /* indicates which of the values are set */
	deq_time     int64             /* absolute/relative  time for dequeuing */
	priority     int64             /* enqueue priority */
	diagnostic   int64             /* indicates reason for failure */
	diagmsg      string            /* diagnostic message */
	msgid        [TMMSGIDLEN]byte  /* id of message before which to queue */
	corrid       [TMCORRIDLEN]byte /* correlation id used to identify message */
	replyqueue   string            /* queue name for reply message */
	failurequeue string            /* queue name for failure message */
	cltid        string            /* client identifier for originating client */
	urcode       int64             /* application user-return code */
	appkey       int64             /* application authentication client key */
	delivery_qos int64             /* delivery quality of service  */
	reply_qos    int64             /* reply message quality of service  */
	exp_time     int64             /* expiration time  */
}

///////////////////////////////////////////////////////////////////////////////////
// ATMI Buffers section
///////////////////////////////////////////////////////////////////////////////////

//ATMI buffer
type ATMIBuf struct {
	C_ptr *C.char
	//We will need some API for length & buffer setting
	//Probably we need a wrapper for lenght function
	C_len C.long
}

//Base interface for typed buffer
type TypedBuffer interface {
	GetBuf() *ATMIBuf
}

//Have inteface to base ATMI buffer
func (u *ATMIBuf) GetBuf() *ATMIBuf {
	return u
}

///////////////////////////////////////////////////////////////////////////////////
// Error Handlers, ATMI level
///////////////////////////////////////////////////////////////////////////////////

//ATMI Error type
type atmiError struct {
	code    int
	message string
}

//ATMI error interface
type ATMIError interface {
	Error() string
	Code() int
	Message() string
}

//Generate ATMI error, read the codes
func NewAtmiError() ATMIError {
	var err atmiError
	err.code = int(C.go_tperrno())
	err.message = C.GoString(C.tpstrerror(C.go_tperrno()))
	return err
}

//Build a custom error
//@param err		Error buffer to build
//@param code	Error code to setup
//@param msg		Error message
func NewCustomAtmiError(code int, msg string) ATMIError {
	var err atmiError
	err.code = code
	err.message = msg
	return err
}

//Standard error interface
func (e atmiError) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.message)
}

//code getter
func (e atmiError) Code() int {
	return e.code
}

//message getter
func (e atmiError) Message() string {
	return e.message
}

///////////////////////////////////////////////////////////////////////////////////
// Error Handlers, NSTD - Enduro/X Standard library
///////////////////////////////////////////////////////////////////////////////////

//NSTD Error type
type nstdError struct {
	code    int
	message string
}

//NSTD error interface
type NSTDError interface {
	Error() string
	Code() int
	Message() string
}

//Generate NSTD error, read the codes
func NewNstdError() NSTDError {
	var err nstdError
	err.code = int(C.go_Nerror())
	err.message = C.GoString(C.Nstrerror(C.go_Nerror()))
	return err
}

//Build a custom error
//@param err		Error buffer to build
//@param code	Error code to setup
//@param msg		Error message
func NewCustomNstdError(code int, msg string) NSTDError {
	var err nstdError
	err.code = code
	err.message = msg
	return err
}

//Standard error interface
func (e nstdError) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.message)
}

//code getter
func (e nstdError) Code() int {
	return e.code
}

//message getter
func (e nstdError) Message() string {
	return e.message
}

///////////////////////////////////////////////////////////////////////////////////
// API Section
///////////////////////////////////////////////////////////////////////////////////

//TODO, maybe we need to use error deligates, so that user can override the error handling object?

//Allocate buffer
//Accepts the standard ATMI values
//We should add error handling here
//@param	 b_type 		Buffer type
//@param	 b_subtype 	Buffer sub-type
//@param	 size		Buffer size request
//@return 			ATMI Buffer, atmiError
func TpAlloc(b_type string, b_subtype string, size int64) (*ATMIBuf, ATMIError) {
	var buf ATMIBuf
	var err ATMIError

	c_type := C.CString(b_type)
	c_subtype := C.CString(b_subtype)

	size_l := C.long(size)

	buf.C_ptr = C.tpalloc(c_type, c_subtype, size_l)

	//Check the error
	if nil == buf.C_ptr {
		err = NewAtmiError()
	}

	C.free(unsafe.Pointer(c_type))
	C.free(unsafe.Pointer(c_subtype))

	runtime.SetFinalizer(&buf, TpFree)

	return &buf, err
}

//Reallocate the buffer
//@param buf		ATMI buffer
//@return 		ATMI Error
func (buf *ATMIBuf) TpRealloc(size int64) ATMIError {
	var err ATMIError

	buf.C_ptr = C.tprealloc(buf.C_ptr, C.long(size))

	if nil == buf.C_ptr {
		err = NewAtmiError()
	}

	return err
}

//Initialize client
//@return		ATMI Error
func TpInit() ATMIError {
	var err ATMIError

	if SUCCEED != C.go_tpinit() {
		err = NewAtmiError()
	}

	return err
}

// Do the service call, assume using the same buffer
// for return value.
// This works for self describing buffers. Otherwise we need a buffer size in
// ATMIBuf.
// @param svc	service name
// @param buf	ATMI buffer
// @param flags 	Flags to be used
// @return atmiError
func TpCall(svc string, tb TypedBuffer, flags int64) (int, ATMIError) {
	var err ATMIError
	c_svc := C.CString(svc)

	buf := tb.GetBuf()

	ret := C.tpcall(c_svc, buf.C_ptr, buf.C_len, &buf.C_ptr, &buf.C_len, C.long(flags))

	if SUCCEED != ret {
		err = NewAtmiError()
	}

	C.free(unsafe.Pointer(c_svc))

	return int(ret), err
}

//TP Async call
//@param svc		Service Name to call
//@param buf		ATMI buffer
//@param flags	Flags to be used for call (see flags section)
//@return		Call Descriptor (cd), ATMI Error
func TpACall(svc string, tb TypedBuffer, flags int64) (int, ATMIError) {
	var err ATMIError
	c_svc := C.CString(svc)

	buf := tb.GetBuf()

	ret := C.tpacall(c_svc, buf.C_ptr, buf.C_len, C.long(flags))

	if FAIL == ret {
		err = NewAtmiError()
	}

	C.free(unsafe.Pointer(c_svc))

	return int(ret), err
}

//Get async call reply
//@param cd	call
//@param buf	ATMI buffer
//@param flags call flags
func TpGetRply(cd *int, tb TypedBuffer, flags int64) (int, ATMIError) {
	var err ATMIError
	var c_cd C.int

	buf := tb.GetBuf()

	ret := C.tpgetrply(&c_cd, &buf.C_ptr, &buf.C_len, C.long(flags))

	if SUCCEED != ret {
		err = NewAtmiError()
	} else {
		*cd = int(c_cd)
	}

	return int(ret), err
}

//Cancel async call
//@param cd		Call descriptor
//@return ATMI error
func TpCancel(cd int) ATMIError {
	var err ATMIError

	ret := C.tpcancel(C.int(cd))

	if SUCCEED != ret {
		err = NewAtmiError()
	}

	return err
}

//Connect to service in conversational mode
//@param svc		Service name
//@param data	ATMI buffers
//@param flags	Flags
//@return		call descriptor (cd), ATMI error
func TpConnect(svc string, tb TypedBuffer, flags int64) (int, ATMIError) {
	var err ATMIError
	c_svc := C.CString(svc)

	data := tb.GetBuf()

	ret := C.tpconnect(c_svc, data.C_ptr, data.C_len, C.long(flags))

	if FAIL == ret {
		err = NewAtmiError()
	}

	C.free(unsafe.Pointer(c_svc))

	return int(ret), err
}

//Disconnect from conversation
//@param cd		Call Descriptor
//@return ATMI Error
func TpDiscon(cd int) ATMIError {
	var err ATMIError

	ret := C.tpdiscon(C.int(cd))

	if SUCCEED != ret {
		err = NewAtmiError()
	}

	return err
}

//Receive data from conversation
//@param cd			call descriptor
//@param	 data		ATMI buffer
//@param revent		Return Event
//@return			ATMI Error
func TpRecv(cd int, tb TypedBuffer, flags int64, revent *int64) ATMIError {
	var err ATMIError

	c_revent := C.long(*revent)

	data := tb.GetBuf()

	ret := C.tprecv(C.int(cd), &data.C_ptr, &data.C_len, C.long(flags), &c_revent)

	if FAIL == ret {
		err = NewAtmiError()
	}

	*revent = int64(c_revent)

	return err
}

//Receive data from conversation
//@param cd			call descriptor
//@param	 data		ATMI buffer
//@param revent		Return Event
//@return			ATMI Error
func TpSend(cd int, tb TypedBuffer, flags int64, revent *int64) ATMIError {
	var err ATMIError

	c_revent := C.long(*revent)

	data := tb.GetBuf()

	ret := C.tpsend(C.int(cd), data.C_ptr, data.C_len, C.long(flags), &c_revent)

	if SUCCEED != ret {
		err = NewAtmiError()
	}

	*revent = int64(c_revent)

	return err
}

//Free the ATMI buffer
//@param buf		ATMI buffer
func TpFree(buf *ATMIBuf) {
	C.tpfree(buf.C_ptr)
	buf.C_ptr = nil
}

//Commit global transaction
//@param	 flags		flags for abort operation
func TpCommit(flags int64) ATMIError {
	var err ATMIError

	ret := C.tpcommit(C.long(flags))

	if SUCCEED != ret {
		err = NewAtmiError()
	}

	return err
}

//Abort global transaction
//@param	 flags		flags for abort operation (must be 0)
//@return ATMI Error
func TpAbort(flags int64) ATMIError {
	var err ATMIError

	ret := C.tpabort(C.long(flags))

	if SUCCEED != ret {
		err = NewAtmiError()
	}

	return err
}

//Open XA Sub-system
//@return ATMI Error
func TpOpen() ATMIError {
	var err ATMIError

	ret := C.tpopen()

	if SUCCEED != ret {
		err = NewAtmiError()
	}

	return err
}

// Close XA Sub-system
//@return ATMI Error
func TpClose() ATMIError {
	var err ATMIError

	ret := C.tpclose()

	if SUCCEED != ret {
		err = NewAtmiError()
	}

	return err
}

//Check are we in globa transaction?
//@return 	0 - not in global Tx, 1 - in global Tx
func TpGetLev() int {

	ret := C.tpgetlev()

	return int(ret)
}

//Begin transaction
//@param timeout		Transaction Timeout
//@param flags		Transaction flags
//@return	ATMI Error
func TpBegin(timeout uint64, flags int64) ATMIError {

	var err ATMIError

	ret := C.tpbegin(C.ulong(timeout), C.long(flags))

	if SUCCEED != ret {
		err = NewAtmiError()
	}

	return err
}

//Suspend transaction
//@param tranid	Transaction Id reference
//@param flags	Flags for suspend (must be 0)
//@return 	ATMI Error
func TpSuspend(tranid *TPTRANID, flags int64) ATMIError {
	var err ATMIError

	ret := C.tpsuspend(&tranid.c_tptranid, C.long(flags))

	if SUCCEED != ret {
		err = NewAtmiError()
	}

	return err
}

//Resume transaction
//@param tranid	Transaction Id reference
//@param flags	Flags for tran resume (must be 0)
//@return 	ATMI Error
func TpResume(tranid *TPTRANID, flags int64) ATMIError {
	var err ATMIError

	ret := C.tpresume(&tranid.c_tptranid, C.long(flags))

	if SUCCEED != ret {
		err = NewAtmiError()
	}

	return err
}

//Get cluster node id
//@return		Node Id
func TpGetnodeId() int64 {
	ret := C.tpgetnodeid()
	return int64(ret)
}

//Post the event to subscribers
//@param eventname	Name of the event to post
//@param data		ATMI buffer
//@param flags		flags
//@return		Number Of events posted, ATMI error
func TpPost(eventname string, tb TypedBuffer, len int64, flags int64) (int, ATMIError) {
	var err ATMIError
	c_eventname := C.CString(eventname)

	data := tb.GetBuf()
	ret := C.tppost(c_eventname, data.C_ptr, data.C_len, C.long(flags))

	if FAIL == ret {
		err = NewAtmiError()
	}

	C.free(unsafe.Pointer(c_eventname))

	return int(ret), err
}

//Return ATMI buffer info
//@param ptr 	Pointer to ATMI buffer
//@param itype	ptr to string to return the buffer type  (can be nil)
//@param subtype ptr to string to return sub-type (can be nil)
func TpTypes(ptr *ATMIBuf, itype *string, subtype *string) (int64, ATMIError) {
	var err ATMIError

	/* we should allocat the fields there...  */

	var c_type *C.char
	var c_subtype *C.char

	c_type = C.malloc_string(16)
	c_subtype = C.malloc_string(16)

	ret := C.tptypes(ptr.C_ptr, c_type, c_subtype)

	if FAIL == ret {
		err = NewAtmiError()
	} else {
		if nil != itype && nil != c_type {
			*itype = C.GoString(c_type)
		}

		if nil != subtype && nil != c_subtype {
			*subtype = C.GoString(c_subtype)
		}
	}

	if nil != c_type {
		C.free_string(c_type)
	}

	if nil != c_subtype {
		C.free_string(c_subtype)
	}

	return int64(ret), err
}

//Terminate the client
//@return ATMI error
func TpTerm() ATMIError {
	ret := C.tpterm()
	if SUCCEED != ret {
		return NewAtmiError()
	}

	return nil
}

//Glue function for tpenqueue and tpdequeue
//@param qspace	Name of the event to post
//@param qname		ATMI buffer
//@param ctl		Control structure
//@param tb		Typed buffer
//@param flags		ATMI call flags
//@param is_enq		Is Enqueue? If not then dequeue
//@return		ATMI error
func tp_enq_deq(qspace string, qname string, ctl *TPQCTL, tb TypedBuffer, flags int64, is_enq bool) ATMIError {
	var err ATMIError

	c_qspace := C.CString(qspace)
	defer C.free(unsafe.Pointer(c_qspace))

	c_qname := C.CString(qname)
	defer C.free(unsafe.Pointer(c_qname))

	c_ctl_flags := C.long(ctl.flags)
	c_ctl_deq_time := C.long(ctl.deq_time)
	c_ctl_priority := C.long(ctl.priority)
	c_ctl_diagnostic := C.long(ctl.diagnostic)
	c_ctl_diagmsg := C.calloc(1, 256)
	c_ctl_diagmsg_ptr := (*C.char)(unsafe.Pointer(c_ctl_diagmsg))
	defer C.free(unsafe.Pointer(c_ctl_diagmsg))

	c_ctl_msgid := C.malloc(TMMSGIDLEN)
	c_ctl_msgid_ptr := (*C.char)(unsafe.Pointer(c_ctl_msgid))
	defer C.free(unsafe.Pointer(c_ctl_msgid))
	for i := 0; i < TMMSGIDLEN; i++ {
		*(*C.char)(unsafe.Pointer(uintptr(c_ctl_msgid) + uintptr(i))) = C.char(ctl.msgid[i])
	}

	c_ctl_corrid := C.malloc(TMCORRIDLEN)
	c_ctl_corrid_ptr := (*C.char)(unsafe.Pointer(c_ctl_corrid))
	defer C.free(unsafe.Pointer(c_ctl_corrid))
	for i := 0; i < TMCORRIDLEN; i++ {
		*(*C.char)(unsafe.Pointer(uintptr(c_ctl_corrid) + uintptr(i))) = C.char(ctl.corrid[i])
	}

	/* Allocate the buffer for reply q, because we might receive this on
	   dequeue.
	*/
	c_ctl_replyqueue_tmp := C.CString(ctl.replyqueue)
	defer C.free(unsafe.Pointer(c_ctl_replyqueue_tmp))
	c_ctl_replyqueue := C.malloc(TMQNAMELEN + 1)
	c_ctl_replyqueue_ptr := (*C.char)(unsafe.Pointer(c_ctl_corrid))
	defer C.free(unsafe.Pointer(c_ctl_replyqueue))

	if C.strlen(c_ctl_replyqueue_tmp) > TMQNAMELEN {
		return NewCustomAtmiError(TPEINVAL,
			fmt.Sprintf("Invalid reply queue len, max: %d", TMQNAMELEN))
	}
	C.strcpy(c_ctl_replyqueue_ptr, c_ctl_replyqueue_tmp)

	/* Allocate the buffer for failure q, because we might receive this on
	   dequeue.
	*/
	c_ctl_failurequeue_tmp := C.CString(ctl.failurequeue)
	defer C.free(unsafe.Pointer(c_ctl_failurequeue_tmp))
	c_ctl_failurequeue := C.malloc(TMQNAMELEN + 1)
	c_ctl_failurequeue_ptr := (*C.char)(unsafe.Pointer(c_ctl_corrid))
	defer C.free(unsafe.Pointer(c_ctl_failurequeue))

	if C.strlen(c_ctl_failurequeue_tmp) > TMQNAMELEN {
		return NewCustomAtmiError(TPEINVAL,
			fmt.Sprintf("Invalid failure queue len, max: %d", TMQNAMELEN))
	}
	C.strcpy(c_ctl_failurequeue_ptr, c_ctl_failurequeue_tmp)

	/* The same goes with client id... we might return it on dequeue */
	c_ctl_cltid_tmp := C.CString(ctl.cltid)
	defer C.free(unsafe.Pointer(c_ctl_cltid_tmp))
	c_ctl_cltid := C.malloc(TMQNAMELEN + 1)
	c_ctl_cltid_ptr := (*C.char)(unsafe.Pointer(c_ctl_corrid))
	defer C.free(unsafe.Pointer(c_ctl_cltid))

	if C.strlen(c_ctl_cltid_tmp) > NDRX_MAX_ID_SIZE {
		return NewCustomAtmiError(TPEINVAL,
			fmt.Sprintf("Invalid client id len, max: %d", TPEINVAL))
	}
	C.strcpy(c_ctl_cltid_ptr, c_ctl_cltid_tmp)

	c_ctl_urcode := C.long(ctl.urcode)
	c_ctl_appkey := C.long(ctl.appkey)
	c_ctl_delivery_qos := C.long(ctl.delivery_qos)
	c_ctl_reply_qos := C.long(ctl.reply_qos)
	c_ctl_exp_time := C.long(ctl.exp_time)

	buf := tb.GetBuf()
	var ret C.int
	if is_enq {
		ret = C.go_tpenqueue(c_qspace, c_qname, buf.C_ptr, buf.C_len, C.long(flags),
			&c_ctl_flags,
			&c_ctl_deq_time,
			&c_ctl_priority,
			&c_ctl_diagnostic,
			c_ctl_diagmsg_ptr,
			c_ctl_msgid_ptr,
			c_ctl_corrid_ptr,
			c_ctl_replyqueue_ptr,
			c_ctl_failurequeue_ptr,
			c_ctl_cltid_ptr,
			&c_ctl_urcode,
			&c_ctl_appkey,
			&c_ctl_delivery_qos,
			&c_ctl_reply_qos,
			&c_ctl_exp_time)
	} else {
		ret = C.go_tpdequeue(c_qspace, c_qname, &buf.C_ptr, &buf.C_len, C.long(flags),
			&c_ctl_flags,
			&c_ctl_deq_time,
			&c_ctl_priority,
			&c_ctl_diagnostic,
			c_ctl_diagmsg_ptr,
			c_ctl_msgid_ptr,
			c_ctl_corrid_ptr,
			c_ctl_replyqueue_ptr,
			c_ctl_failurequeue_ptr,
			c_ctl_cltid_ptr,
			&c_ctl_urcode,
			&c_ctl_appkey,
			&c_ctl_delivery_qos,
			&c_ctl_reply_qos,
			&c_ctl_exp_time)
	}

	/* transfer back to structure values we got... */
	ctl.flags = int64(c_ctl_flags)
	ctl.deq_time = int64(c_ctl_deq_time)
	ctl.priority = int64(c_ctl_priority)
	ctl.diagnostic = int64(c_ctl_diagnostic)

	ctl.diagmsg = C.GoString(c_ctl_diagmsg_ptr)

	for i := 0; i < TMMSGIDLEN; i++ {
		ctl.msgid[i] = byte(*(*C.char)(unsafe.Pointer(uintptr(c_ctl_msgid) + uintptr(i))))
	}

	for i := 0; i < TMCORRIDLEN; i++ {
		ctl.corrid[i] = byte(*(*C.char)(unsafe.Pointer(uintptr(c_ctl_corrid) + uintptr(i))))
	}

	ctl.replyqueue = C.GoString(c_ctl_replyqueue_ptr)
	ctl.failurequeue = C.GoString(c_ctl_failurequeue_ptr)
	ctl.cltid = C.GoString(c_ctl_cltid_ptr)

	ctl.urcode = int64(c_ctl_urcode)
	ctl.appkey = int64(c_ctl_appkey)
	ctl.delivery_qos = int64(c_ctl_delivery_qos)
	ctl.reply_qos = int64(c_ctl_reply_qos)
	ctl.exp_time = int64(c_ctl_exp_time)

	if FAIL == ret {
		err = NewAtmiError()
	}

	return err
}

//Enqueue message to Q
//@param qspace	Name of the event to post
//@param qname		ATMI buffer
//@param ctl		Control structure
//@param tb		Typed buffer
//@param flags		ATMI call flags
//@return		ATMI error
func TpEnqueue(qspace string, qname string, ctl *TPQCTL, tb TypedBuffer, flags int64) ATMIError {
	return tp_enq_deq(qspace, qname, ctl, tb, flags, true)
}

//Dequeue message from Q
//@param qspace	Name of the event to post
//@param qname		ATMI buffer
//@param ctl		Control structure
//@param tb		Typed buffer
//@param flags		ATMI call flags
//@return		ATMI error
func TpDequeue(qspace string, qname string, ctl *TPQCTL, tb TypedBuffer, flags int64) ATMIError {
	return tp_enq_deq(qspace, qname, ctl, tb, flags, false)
}
