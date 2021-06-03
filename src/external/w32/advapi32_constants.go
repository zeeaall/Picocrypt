package github.com/HACKERALERT/Picocrypt/src/external/w32

// Registry predefined keys
const (
	HKEY_CLASSES_ROOT     HKEY = 0x80000000
	HKEY_CURRENT_USER     HKEY = 0x80000001
	HKEY_LOCAL_MACHINE    HKEY = 0x80000002
	HKEY_USERS            HKEY = 0x80000003
	HKEY_PERFORMANCE_DATA HKEY = 0x80000004
	HKEY_CURRENT_CONFIG   HKEY = 0x80000005
	HKEY_DYN_DATA         HKEY = 0x80000006
)

// Registry Key Security and Access Rights
const (
	KEY_ALL_ACCESS         = 0xF003F
	KEY_CREATE_SUB_KEY     = 0x0004
	KEY_ENUMERATE_SUB_KEYS = 0x0008
	KEY_NOTIFY             = 0x0010
	KEY_QUERY_VALUE        = 0x0001
	KEY_SET_VALUE          = 0x0002
	KEY_READ               = 0x20019
	KEY_WRITE              = 0x20006
)

const (
	NFR_ANSI    = 1
	NFR_UNICODE = 2
	NF_QUERY    = 3
	NF_REQUERY  = 4
)

// Registry value types
const (
	RRF_RT_REG_NONE         = 0x00000001
	RRF_RT_REG_SZ           = 0x00000002
	RRF_RT_REG_EXPAND_SZ    = 0x00000004
	RRF_RT_REG_BINARY       = 0x00000008
	RRF_RT_REG_DWORD        = 0x00000010
	RRF_RT_REG_MULTI_SZ     = 0x00000020
	RRF_RT_REG_QWORD        = 0x00000040
	RRF_RT_DWORD            = (RRF_RT_REG_BINARY | RRF_RT_REG_DWORD)
	RRF_RT_QWORD            = (RRF_RT_REG_BINARY | RRF_RT_REG_QWORD)
	RRF_RT_ANY              = 0x0000ffff
	RRF_NOEXPAND            = 0x10000000
	RRF_ZEROONFAILURE       = 0x20000000
	REG_PROCESS_APPKEY      = 0x00000001
	REG_MUI_STRING_TRUNCATE = 0x00000001
)

// Service Control Manager object specific access types
const (
	SC_MANAGER_CONNECT            = 0x0001
	SC_MANAGER_CREATE_SERVICE     = 0x0002
	SC_MANAGER_ENUMERATE_SERVICE  = 0x0004
	SC_MANAGER_LOCK               = 0x0008
	SC_MANAGER_QUERY_LOCK_STATUS  = 0x0010
	SC_MANAGER_MODIFY_BOOT_CONFIG = 0x0020
	SC_MANAGER_ALL_ACCESS         = STANDARD_RIGHTS_REQUIRED | SC_MANAGER_CONNECT | SC_MANAGER_CREATE_SERVICE | SC_MANAGER_ENUMERATE_SERVICE | SC_MANAGER_LOCK | SC_MANAGER_QUERY_LOCK_STATUS | SC_MANAGER_MODIFY_BOOT_CONFIG
)

// Service Types (Bit Mask)
const (
	SERVICE_KERNEL_DRIVER       = 0x00000001
	SERVICE_FILE_SYSTEM_DRIVER  = 0x00000002
	SERVICE_ADAPTER             = 0x00000004
	SERVICE_RECOGNIZER_DRIVER   = 0x00000008
	SERVICE_DRIVER              = SERVICE_KERNEL_DRIVER | SERVICE_FILE_SYSTEM_DRIVER | SERVICE_RECOGNIZER_DRIVER
	SERVICE_WIN32_OWN_PROCESS   = 0x00000010
	SERVICE_WIN32_SHARE_PROCESS = 0x00000020
	SERVICE_WIN32               = SERVICE_WIN32_OWN_PROCESS | SERVICE_WIN32_SHARE_PROCESS
	SERVICE_INTERACTIVE_PROCESS = 0x00000100
	SERVICE_TYPE_ALL            = SERVICE_WIN32 | SERVICE_ADAPTER | SERVICE_DRIVER | SERVICE_INTERACTIVE_PROCESS
)

// Service State -- for CurrentState
const (
	SERVICE_STOPPED          = 0x00000001
	SERVICE_START_PENDING    = 0x00000002
	SERVICE_STOP_PENDING     = 0x00000003
	SERVICE_RUNNING          = 0x00000004
	SERVICE_CONTINUE_PENDING = 0x00000005
	SERVICE_PAUSE_PENDING    = 0x00000006
	SERVICE_PAUSED           = 0x00000007
)

// Controls Accepted  (Bit Mask)
const (
	SERVICE_ACCEPT_STOP                  = 0x00000001
	SERVICE_ACCEPT_PAUSE_CONTINUE        = 0x00000002
	SERVICE_ACCEPT_SHUTDOWN              = 0x00000004
	SERVICE_ACCEPT_PARAMCHANGE           = 0x00000008
	SERVICE_ACCEPT_NETBINDCHANGE         = 0x00000010
	SERVICE_ACCEPT_HARDWAREPROFILECHANGE = 0x00000020
	SERVICE_ACCEPT_POWEREVENT            = 0x00000040
	SERVICE_ACCEPT_SESSIONCHANGE         = 0x00000080
	SERVICE_ACCEPT_PRESHUTDOWN           = 0x00000100
	SERVICE_ACCEPT_TIMECHANGE            = 0x00000200
	SERVICE_ACCEPT_TRIGGEREVENT          = 0x00000400
)

// Service object specific access type
const (
	SERVICE_QUERY_CONFIG         = 0x0001
	SERVICE_CHANGE_CONFIG        = 0x0002
	SERVICE_QUERY_STATUS         = 0x0004
	SERVICE_ENUMERATE_DEPENDENTS = 0x0008
	SERVICE_START                = 0x0010
	SERVICE_STOP                 = 0x0020
	SERVICE_PAUSE_CONTINUE       = 0x0040
	SERVICE_INTERROGATE          = 0x0080
	SERVICE_USER_DEFINED_CONTROL = 0x0100

	SERVICE_ALL_ACCESS = STANDARD_RIGHTS_REQUIRED |
		SERVICE_QUERY_CONFIG |
		SERVICE_CHANGE_CONFIG |
		SERVICE_QUERY_STATUS |
		SERVICE_ENUMERATE_DEPENDENTS |
		SERVICE_START |
		SERVICE_STOP |
		SERVICE_PAUSE_CONTINUE |
		SERVICE_INTERROGATE |
		SERVICE_USER_DEFINED_CONTROL
)

const (
	KERNEL_LOGGER_NAME = "NT Kernel Logger"
)

// WNODE flags, for ETW (Event Tracing for Windows) / WMI
const (
	WNODE_FLAG_ALL_DATA              = 0x00000001
	WNODE_FLAG_SINGLE_INSTANCE       = 0x00000002
	WNODE_FLAG_SINGLE_ITEM           = 0x00000004
	WNODE_FLAG_EVENT_ITEM            = 0x00000008
	WNODE_FLAG_FIXED_INSTANCE_SIZE   = 0x00000010
	WNODE_FLAG_TOO_SMALL             = 0x00000020
	WNODE_FLAG_INSTANCES_SAME        = 0x00000040
	WNODE_FLAG_STATIC_INSTANCE_NAMES = 0x00000080
	WNODE_FLAG_INTERNAL              = 0x00000100
	WNODE_FLAG_USE_TIMESTAMP         = 0x00000200
	WNODE_FLAG_PERSIST_EVENT         = 0x00000400
	WNODE_FLAG_EVENT_REFERENCE       = 0x00002000
	WNODE_FLAG_ANSI_INSTANCENAMES    = 0x00004000
	WNODE_FLAG_METHOD_ITEM           = 0x00008000
	WNODE_FLAG_PDO_INSTANCE_NAMES    = 0x00010000
	WNODE_FLAG_TRACED_GUID           = 0x00020000
	WNODE_FLAG_LOG_WNODE             = 0x00040000
	WNODE_FLAG_USE_GUID_PTR          = 0x00080000
	WNODE_FLAG_USE_MOF_PTR           = 0x00100000
	WNODE_FLAG_NO_HEADER             = 0x00200000
	WNODE_FLAG_SEVERITY_MASK         = 0xff000000
)

// ETW flags and types etc
const (
	EVENT_TRACE_TYPE_INFO                  = 0x00
	EVENT_TRACE_TYPE_START                 = 0x01
	EVENT_TRACE_TYPE_END                   = 0x02
	EVENT_TRACE_TYPE_STOP                  = 0x02
	EVENT_TRACE_TYPE_DC_START              = 0x03
	EVENT_TRACE_TYPE_DC_END                = 0x04
	EVENT_TRACE_TYPE_EXTENSION             = 0x05
	EVENT_TRACE_TYPE_REPLY                 = 0x06
	EVENT_TRACE_TYPE_DEQUEUE               = 0x07
	EVENT_TRACE_TYPE_RESUME                = 0x07
	EVENT_TRACE_TYPE_CHECKPOINT            = 0x08
	EVENT_TRACE_TYPE_SUSPEND               = 0x08
	EVENT_TRACE_TYPE_WINEVT_SEND           = 0x09
	EVENT_TRACE_TYPE_WINEVT_RECEIVE        = 0XF0
	TRACE_LEVEL_NONE                       = 0
	TRACE_LEVEL_CRITICAL                   = 1
	TRACE_LEVEL_FATAL                      = 1
	TRACE_LEVEL_ERROR                      = 2
	TRACE_LEVEL_WARNING                    = 3
	TRACE_LEVEL_INFORMATION                = 4
	TRACE_LEVEL_VERBOSE                    = 5
	TRACE_LEVEL_RESERVED6                  = 6
	TRACE_LEVEL_RESERVED7                  = 7
	TRACE_LEVEL_RESERVED8                  = 8
	TRACE_LEVEL_RESERVED9                  = 9
	EVENT_TRACE_TYPE_LOAD                  = 0x0A
	EVENT_TRACE_TYPE_IO_READ               = 0x0A
	EVENT_TRACE_TYPE_IO_WRITE              = 0x0B
	EVENT_TRACE_TYPE_IO_READ_INIT          = 0x0C
	EVENT_TRACE_TYPE_IO_WRITE_INIT         = 0x0D
	EVENT_TRACE_TYPE_IO_FLUSH              = 0x0E
	EVENT_TRACE_TYPE_IO_FLUSH_INIT         = 0x0F
	EVENT_TRACE_TYPE_MM_TF                 = 0x0A
	EVENT_TRACE_TYPE_MM_DZF                = 0x0B
	EVENT_TRACE_TYPE_MM_COW                = 0x0C
	EVENT_TRACE_TYPE_MM_GPF                = 0x0D
	EVENT_TRACE_TYPE_MM_HPF                = 0x0E
	EVENT_TRACE_TYPE_MM_AV                 = 0x0F
	EVENT_TRACE_TYPE_SEND                  = 0x0A
	EVENT_TRACE_TYPE_RECEIVE               = 0x0B
	EVENT_TRACE_TYPE_CONNECT               = 0x0C
	EVENT_TRACE_TYPE_DISCONNECT            = 0x0D
	EVENT_TRACE_TYPE_RETRANSMIT            = 0x0E
	EVENT_TRACE_TYPE_ACCEPT                = 0x0F
	EVENT_TRACE_TYPE_RECONNECT             = 0x10
	EVENT_TRACE_TYPE_CONNFAIL              = 0x11
	EVENT_TRACE_TYPE_COPY_TCP              = 0x12
	EVENT_TRACE_TYPE_COPY_ARP              = 0x13
	EVENT_TRACE_TYPE_ACKFULL               = 0x14
	EVENT_TRACE_TYPE_ACKPART               = 0x15
	EVENT_TRACE_TYPE_ACKDUP                = 0x16
	EVENT_TRACE_TYPE_GUIDMAP               = 0x0A
	EVENT_TRACE_TYPE_CONFIG                = 0x0B
	EVENT_TRACE_TYPE_SIDINFO               = 0x0C
	EVENT_TRACE_TYPE_SECURITY              = 0x0D
	EVENT_TRACE_TYPE_REGCREATE             = 0x0A
	EVENT_TRACE_TYPE_REGOPEN               = 0x0B
	EVENT_TRACE_TYPE_REGDELETE             = 0x0C
	EVENT_TRACE_TYPE_REGQUERY              = 0x0D
	EVENT_TRACE_TYPE_REGSETVALUE           = 0x0E
	EVENT_TRACE_TYPE_REGDELETEVALUE        = 0x0F
	EVENT_TRACE_TYPE_REGQUERYVALUE         = 0x10
	EVENT_TRACE_TYPE_REGENUMERATEKEY       = 0x11
	EVENT_TRACE_TYPE_REGENUMERATEVALUEKEY  = 0x12
	EVENT_TRACE_TYPE_REGQUERYMULTIPLEVALUE = 0x13
	EVENT_TRACE_TYPE_REGSETINFORMATION     = 0x14
	EVENT_TRACE_TYPE_REGFLUSH              = 0x15
	EVENT_TRACE_TYPE_REGKCBCREATE          = 0x16
	EVENT_TRACE_TYPE_REGKCBDELETE          = 0x17
	EVENT_TRACE_TYPE_REGKCBRUNDOWNBEGIN    = 0x18
	EVENT_TRACE_TYPE_REGKCBRUNDOWNEND      = 0x19
	EVENT_TRACE_TYPE_REGVIRTUALIZE         = 0x1A
	EVENT_TRACE_TYPE_REGCLOSE              = 0x1B
	EVENT_TRACE_TYPE_REGSETSECURITY        = 0x1C
	EVENT_TRACE_TYPE_REGQUERYSECURITY      = 0x1D
	EVENT_TRACE_TYPE_REGCOMMIT             = 0x1E
	EVENT_TRACE_TYPE_REGPREPARE            = 0x1F
	EVENT_TRACE_TYPE_REGROLLBACK           = 0x20
	EVENT_TRACE_TYPE_REGMOUNTHIVE          = 0x21
	EVENT_TRACE_TYPE_CONFIG_CPU            = 0x0A
	EVENT_TRACE_TYPE_CONFIG_PHYSICALDISK   = 0x0B
	EVENT_TRACE_TYPE_CONFIG_LOGICALDISK    = 0x0C
	EVENT_TRACE_TYPE_CONFIG_NIC            = 0x0D
	EVENT_TRACE_TYPE_CONFIG_VIDEO          = 0x0E
	EVENT_TRACE_TYPE_CONFIG_SERVICES       = 0x0F
	EVENT_TRACE_TYPE_CONFIG_POWER          = 0x10
	EVENT_TRACE_TYPE_CONFIG_NETINFO        = 0x11
	EVENT_TRACE_TYPE_CONFIG_IRQ            = 0x15
	EVENT_TRACE_TYPE_CONFIG_PNP            = 0x16
	EVENT_TRACE_TYPE_CONFIG_IDECHANNEL     = 0x17
	EVENT_TRACE_TYPE_CONFIG_PLATFORM       = 0x19
	EVENT_TRACE_FLAG_PROCESS               = 0x00000001
	EVENT_TRACE_FLAG_THREAD                = 0x00000002
	EVENT_TRACE_FLAG_IMAGE_LOAD            = 0x00000004
	EVENT_TRACE_FLAG_DISK_IO               = 0x00000100
	EVENT_TRACE_FLAG_DISK_FILE_IO          = 0x00000200
	EVENT_TRACE_FLAG_MEMORY_PAGE_FAULTS    = 0x00001000
	EVENT_TRACE_FLAG_MEMORY_HARD_FAULTS    = 0x00002000
	EVENT_TRACE_FLAG_NETWORK_TCPIP         = 0x00010000
	EVENT_TRACE_FLAG_REGISTRY              = 0x00020000
	EVENT_TRACE_FLAG_DBGPRINT              = 0x00040000
	EVENT_TRACE_FLAG_PROCESS_COUNTERS      = 0x00000008
	EVENT_TRACE_FLAG_CSWITCH               = 0x00000010
	EVENT_TRACE_FLAG_DPC                   = 0x00000020
	EVENT_TRACE_FLAG_INTERRUPT             = 0x00000040
	EVENT_TRACE_FLAG_SYSTEMCALL            = 0x00000080
	EVENT_TRACE_FLAG_DISK_IO_INIT          = 0x00000400
	EVENT_TRACE_FLAG_ALPC                  = 0x00100000
	EVENT_TRACE_FLAG_SPLIT_IO              = 0x00200000
	EVENT_TRACE_FLAG_DRIVER                = 0x00800000
	EVENT_TRACE_FLAG_PROFILE               = 0x01000000
	EVENT_TRACE_FLAG_FILE_IO               = 0x02000000
	EVENT_TRACE_FLAG_FILE_IO_INIT          = 0x04000000
	EVENT_TRACE_FLAG_DISPATCHER            = 0x00000800
	EVENT_TRACE_FLAG_VIRTUAL_ALLOC         = 0x00004000
	EVENT_TRACE_FLAG_EXTENSION             = 0x80000000
	EVENT_TRACE_FLAG_FORWARD_WMI           = 0x40000000
	EVENT_TRACE_FLAG_ENABLE_RESERVE        = 0x20000000
	EVENT_TRACE_FILE_MODE_NONE             = 0x00000000
	EVENT_TRACE_FILE_MODE_SEQUENTIAL       = 0x00000001
	EVENT_TRACE_FILE_MODE_CIRCULAR         = 0x00000002
	EVENT_TRACE_FILE_MODE_APPEND           = 0x00000004
	EVENT_TRACE_REAL_TIME_MODE             = 0x00000100
	EVENT_TRACE_DELAY_OPEN_FILE_MODE       = 0x00000200
	EVENT_TRACE_BUFFERING_MODE             = 0x00000400
	EVENT_TRACE_PRIVATE_LOGGER_MODE        = 0x00000800
	EVENT_TRACE_ADD_HEADER_MODE            = 0x00001000
	EVENT_TRACE_USE_GLOBAL_SEQUENCE        = 0x00004000
	EVENT_TRACE_USE_LOCAL_SEQUENCE         = 0x00008000
	EVENT_TRACE_RELOG_MODE                 = 0x00010000
	EVENT_TRACE_USE_PAGED_MEMORY           = 0x01000000
	EVENT_TRACE_FILE_MODE_NEWFILE          = 0x00000008
	EVENT_TRACE_FILE_MODE_PREALLOCATE      = 0x00000020
	EVENT_TRACE_NONSTOPPABLE_MODE          = 0x00000040
	EVENT_TRACE_SECURE_MODE                = 0x00000080
	EVENT_TRACE_USE_KBYTES_FOR_SIZE        = 0x00002000
	EVENT_TRACE_PRIVATE_IN_PROC            = 0x00020000
	EVENT_TRACE_MODE_RESERVED              = 0x00100000
	EVENT_TRACE_NO_PER_PROCESSOR_BUFFERING = 0x10000000
	EVENT_TRACE_CONTROL_QUERY              = 0
	EVENT_TRACE_CONTROL_STOP               = 1
	EVENT_TRACE_CONTROL_UPDATE             = 2
	EVENT_TRACE_CONTROL_FLUSH              = 3
)
