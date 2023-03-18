package gonvml

const (
	// NO_UNVERSIONED_FUNC_DEFS as defined in go-nvml/<predefine>:24
	NO_UNVERSIONED_FUNC_DEFS = 1
	// API_VERSION as defined in nvml/nvml.h
	API_VERSION = 11
	// API_VERSION_STR as defined in nvml/nvml.h
	API_VERSION_STR = "11"
	// VALUE_NOT_AVAILABLE as defined in nvml/nvml.h
	VALUE_NOT_AVAILABLE = -1
	// DEVICE_PCI_BUS_ID_BUFFER_SIZE as defined in nvml/nvml.h
	DEVICE_PCI_BUS_ID_BUFFER_SIZE = 32
	// DEVICE_PCI_BUS_ID_BUFFER_V2_SIZE as defined in nvml/nvml.h
	DEVICE_PCI_BUS_ID_BUFFER_V2_SIZE = 16
	// DEVICE_PCI_BUS_ID_LEGACY_FMT as defined in nvml/nvml.h
	DEVICE_PCI_BUS_ID_LEGACY_FMT = "%04X:%02X:%02X.0"
	// DEVICE_PCI_BUS_ID_FMT as defined in nvml/nvml.h
	DEVICE_PCI_BUS_ID_FMT = "%08X:%02X:%02X.0"
	// NVLINK_MAX_LINKS as defined in nvml/nvml.h
	NVLINK_MAX_LINKS = 18
	// TOPOLOGY_CPU as defined in nvml/nvml.h
	TOPOLOGY_CPU = 0
	// MAX_PHYSICAL_BRIDGE as defined in nvml/nvml.h
	MAX_PHYSICAL_BRIDGE = 128
	// MAX_THERMAL_SENSORS_PER_GPU as defined in nvml/nvml.h
	MAX_THERMAL_SENSORS_PER_GPU = 3
	// FlagDefault as defined in nvml/nvml.h
	FlagDefault = 0
	// FlagForce as defined in nvml/nvml.h
	FlagForce = 1
	// SINGLE_BIT_ECC as defined in nvml/nvml.h
	SINGLE_BIT_ECC = 0
	// DOUBLE_BIT_ECC as defined in nvml/nvml.h
	DOUBLE_BIT_ECC = 0
	// MAX_GPU_PERF_PSTATES as defined in nvml/nvml.h
	MAX_GPU_PERF_PSTATES = 16
	// GRID_LICENSE_EXPIRY_NOT_AVAILABLE as defined in nvml/nvml.h
	GRID_LICENSE_EXPIRY_NOT_AVAILABLE = 0
	// GRID_LICENSE_EXPIRY_INVALID as defined in nvml/nvml.h
	GRID_LICENSE_EXPIRY_INVALID = 1
	// GRID_LICENSE_EXPIRY_VALID as defined in nvml/nvml.h
	GRID_LICENSE_EXPIRY_VALID = 2
	// GRID_LICENSE_EXPIRY_NOT_APPLICABLE as defined in nvml/nvml.h
	GRID_LICENSE_EXPIRY_NOT_APPLICABLE = 3
	// GRID_LICENSE_EXPIRY_PERMANENT as defined in nvml/nvml.h
	GRID_LICENSE_EXPIRY_PERMANENT = 4
	// GRID_LICENSE_BUFFER_SIZE as defined in nvml/nvml.h
	GRID_LICENSE_BUFFER_SIZE = 128
	// VGPU_NAME_BUFFER_SIZE as defined in nvml/nvml.h
	VGPU_NAME_BUFFER_SIZE = 64
	// GRID_LICENSE_FEATURE_MAX_COUNT as defined in nvml/nvml.h
	GRID_LICENSE_FEATURE_MAX_COUNT = 3
	// VGPU_SCHEDULER_POLICY_UNKNOWN as defined in nvml/nvml.h
	VGPU_SCHEDULER_POLICY_UNKNOWN = 0
	// VGPU_SCHEDULER_POLICY_BEST_EFFORT as defined in nvml/nvml.h
	VGPU_SCHEDULER_POLICY_BEST_EFFORT = 1
	// VGPU_SCHEDULER_POLICY_EQUAL_SHARE as defined in nvml/nvml.h
	VGPU_SCHEDULER_POLICY_EQUAL_SHARE = 2
	// VGPU_SCHEDULER_POLICY_FIXED_SHARE as defined in nvml/nvml.h
	VGPU_SCHEDULER_POLICY_FIXED_SHARE = 3
	// SUPPORTED_VGPU_SCHEDULER_POLICY_COUNT as defined in nvml/nvml.h
	SUPPORTED_VGPU_SCHEDULER_POLICY_COUNT = 3
	// SCHEDULER_SW_MAX_LOG_ENTRIES as defined in nvml/nvml.h
	SCHEDULER_SW_MAX_LOG_ENTRIES = 200
	// GRID_LICENSE_STATE_UNKNOWN as defined in nvml/nvml.h
	GRID_LICENSE_STATE_UNKNOWN = 0
	// GRID_LICENSE_STATE_UNINITIALIZED as defined in nvml/nvml.h
	GRID_LICENSE_STATE_UNINITIALIZED = 1
	// GRID_LICENSE_STATE_UNLICENSED_UNRESTRICTED as defined in nvml/nvml.h
	GRID_LICENSE_STATE_UNLICENSED_UNRESTRICTED = 2
	// GRID_LICENSE_STATE_UNLICENSED_RESTRICTED as defined in nvml/nvml.h
	GRID_LICENSE_STATE_UNLICENSED_RESTRICTED = 3
	// GRID_LICENSE_STATE_UNLICENSED as defined in nvml/nvml.h
	GRID_LICENSE_STATE_UNLICENSED = 4
	// GRID_LICENSE_STATE_LICENSED as defined in nvml/nvml.h
	GRID_LICENSE_STATE_LICENSED = 5
	// GSP_FIRMWARE_VERSION_BUF_SIZE as defined in nvml/nvml.h
	GSP_FIRMWARE_VERSION_BUF_SIZE = 64
	// DEVICE_ARCH_KEPLER as defined in nvml/nvml.h
	DEVICE_ARCH_KEPLER = 2
	// DEVICE_ARCH_MAXWELL as defined in nvml/nvml.h
	DEVICE_ARCH_MAXWELL = 3
	// DEVICE_ARCH_PASCAL as defined in nvml/nvml.h
	DEVICE_ARCH_PASCAL = 4
	// DEVICE_ARCH_VOLTA as defined in nvml/nvml.h
	DEVICE_ARCH_VOLTA = 5
	// DEVICE_ARCH_TURING as defined in nvml/nvml.h
	DEVICE_ARCH_TURING = 6
	// DEVICE_ARCH_AMPERE as defined in nvml/nvml.h
	DEVICE_ARCH_AMPERE = 7
	// DEVICE_ARCH_ADA as defined in nvml/nvml.h
	DEVICE_ARCH_ADA = 8
	// DEVICE_ARCH_HOPPER as defined in nvml/nvml.h
	DEVICE_ARCH_HOPPER = 9
	// DEVICE_ARCH_UNKNOWN as defined in nvml/nvml.h
	DEVICE_ARCH_UNKNOWN = 4294967295
	// BUS_TYPE_UNKNOWN as defined in nvml/nvml.h
	BUS_TYPE_UNKNOWN = 0
	// BUS_TYPE_PCI as defined in nvml/nvml.h
	BUS_TYPE_PCI = 1
	// BUS_TYPE_PCIE as defined in nvml/nvml.h
	BUS_TYPE_PCIE = 2
	// BUS_TYPE_FPCI as defined in nvml/nvml.h
	BUS_TYPE_FPCI = 3
	// BUS_TYPE_AGP as defined in nvml/nvml.h
	BUS_TYPE_AGP = 4
	// FAN_POLICY_TEMPERATURE_CONTINOUS_SW as defined in nvml/nvml.h
	FAN_POLICY_TEMPERATURE_CONTINOUS_SW = 0
	// FAN_POLICY_MANUAL as defined in nvml/nvml.h
	FAN_POLICY_MANUAL = 1
	// POWER_SOURCE_AC as defined in nvml/nvml.h
	POWER_SOURCE_AC = 0
	// POWER_SOURCE_BATTERY as defined in nvml/nvml.h
	POWER_SOURCE_BATTERY = 1
	// PCIE_LINK_MAX_SPEED_INVALID as defined in nvml/nvml.h
	PCIE_LINK_MAX_SPEED_INVALID = 0
	// PCIE_LINK_MAX_SPEED_2500MBPS as defined in nvml/nvml.h
	PCIE_LINK_MAX_SPEED_2500MBPS = 1
	// PCIE_LINK_MAX_SPEED_5000MBPS as defined in nvml/nvml.h
	PCIE_LINK_MAX_SPEED_5000MBPS = 2
	// PCIE_LINK_MAX_SPEED_8000MBPS as defined in nvml/nvml.h
	PCIE_LINK_MAX_SPEED_8000MBPS = 3
	// PCIE_LINK_MAX_SPEED_16000MBPS as defined in nvml/nvml.h
	PCIE_LINK_MAX_SPEED_16000MBPS = 4
	// PCIE_LINK_MAX_SPEED_32000MBPS as defined in nvml/nvml.h
	PCIE_LINK_MAX_SPEED_32000MBPS = 5
	// PCIE_LINK_MAX_SPEED_64000MBPS as defined in nvml/nvml.h
	PCIE_LINK_MAX_SPEED_64000MBPS = 6
	// ADAPTIVE_CLOCKING_INFO_STATUS_DISABLED as defined in nvml/nvml.h
	ADAPTIVE_CLOCKING_INFO_STATUS_DISABLED = 0
	// ADAPTIVE_CLOCKING_INFO_STATUS_ENABLED as defined in nvml/nvml.h
	ADAPTIVE_CLOCKING_INFO_STATUS_ENABLED = 1
	// MAX_GPU_UTILIZATIONS as defined in nvml/nvml.h
	MAX_GPU_UTILIZATIONS = 8
	// FI_DEV_ECC_CURRENT as defined in nvml/nvml.h
	FI_DEV_ECC_CURRENT = 1
	// FI_DEV_ECC_PENDING as defined in nvml/nvml.h
	FI_DEV_ECC_PENDING = 2
	// FI_DEV_ECC_SBE_VOL_TOTAL as defined in nvml/nvml.h
	FI_DEV_ECC_SBE_VOL_TOTAL = 3
	// FI_DEV_ECC_DBE_VOL_TOTAL as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_VOL_TOTAL = 4
	// FI_DEV_ECC_SBE_AGG_TOTAL as defined in nvml/nvml.h
	FI_DEV_ECC_SBE_AGG_TOTAL = 5
	// FI_DEV_ECC_DBE_AGG_TOTAL as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_AGG_TOTAL = 6
	// FI_DEV_ECC_SBE_VOL_L1 as defined in nvml/nvml.h
	FI_DEV_ECC_SBE_VOL_L1 = 7
	// FI_DEV_ECC_DBE_VOL_L1 as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_VOL_L1 = 8
	// FI_DEV_ECC_SBE_VOL_L2 as defined in nvml/nvml.h
	FI_DEV_ECC_SBE_VOL_L2 = 9
	// FI_DEV_ECC_DBE_VOL_L2 as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_VOL_L2 = 10
	// FI_DEV_ECC_SBE_VOL_DEV as defined in nvml/nvml.h
	FI_DEV_ECC_SBE_VOL_DEV = 11
	// FI_DEV_ECC_DBE_VOL_DEV as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_VOL_DEV = 12
	// FI_DEV_ECC_SBE_VOL_REG as defined in nvml/nvml.h
	FI_DEV_ECC_SBE_VOL_REG = 13
	// FI_DEV_ECC_DBE_VOL_REG as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_VOL_REG = 14
	// FI_DEV_ECC_SBE_VOL_TEX as defined in nvml/nvml.h
	FI_DEV_ECC_SBE_VOL_TEX = 15
	// FI_DEV_ECC_DBE_VOL_TEX as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_VOL_TEX = 16
	// FI_DEV_ECC_DBE_VOL_CBU as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_VOL_CBU = 17
	// FI_DEV_ECC_SBE_AGG_L1 as defined in nvml/nvml.h
	FI_DEV_ECC_SBE_AGG_L1 = 18
	// FI_DEV_ECC_DBE_AGG_L1 as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_AGG_L1 = 19
	// FI_DEV_ECC_SBE_AGG_L2 as defined in nvml/nvml.h
	FI_DEV_ECC_SBE_AGG_L2 = 20
	// FI_DEV_ECC_DBE_AGG_L2 as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_AGG_L2 = 21
	// FI_DEV_ECC_SBE_AGG_DEV as defined in nvml/nvml.h
	FI_DEV_ECC_SBE_AGG_DEV = 22
	// FI_DEV_ECC_DBE_AGG_DEV as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_AGG_DEV = 23
	// FI_DEV_ECC_SBE_AGG_REG as defined in nvml/nvml.h
	FI_DEV_ECC_SBE_AGG_REG = 24
	// FI_DEV_ECC_DBE_AGG_REG as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_AGG_REG = 25
	// FI_DEV_ECC_SBE_AGG_TEX as defined in nvml/nvml.h
	FI_DEV_ECC_SBE_AGG_TEX = 26
	// FI_DEV_ECC_DBE_AGG_TEX as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_AGG_TEX = 27
	// FI_DEV_ECC_DBE_AGG_CBU as defined in nvml/nvml.h
	FI_DEV_ECC_DBE_AGG_CBU = 28
	// FI_DEV_RETIRED_SBE as defined in nvml/nvml.h
	FI_DEV_RETIRED_SBE = 29
	// FI_DEV_RETIRED_DBE as defined in nvml/nvml.h
	FI_DEV_RETIRED_DBE = 30
	// FI_DEV_RETIRED_PENDING as defined in nvml/nvml.h
	FI_DEV_RETIRED_PENDING = 31
	// FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L0 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L0 = 32
	// FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L1 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L1 = 33
	// FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L2 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L2 = 34
	// FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L3 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L3 = 35
	// FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L4 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L4 = 36
	// FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L5 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L5 = 37
	// FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_TOTAL as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_TOTAL = 38
	// FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L0 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L0 = 39
	// FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L1 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L1 = 40
	// FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L2 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L2 = 41
	// FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L3 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L3 = 42
	// FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L4 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L4 = 43
	// FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L5 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L5 = 44
	// FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_TOTAL as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_TOTAL = 45
	// FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L0 as defined in nvml/nvml.h
	FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L0 = 46
	// FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L1 as defined in nvml/nvml.h
	FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L1 = 47
	// FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L2 as defined in nvml/nvml.h
	FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L2 = 48
	// FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L3 as defined in nvml/nvml.h
	FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L3 = 49
	// FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L4 as defined in nvml/nvml.h
	FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L4 = 50
	// FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L5 as defined in nvml/nvml.h
	FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L5 = 51
	// FI_DEV_NVLINK_REPLAY_ERROR_COUNT_TOTAL as defined in nvml/nvml.h
	FI_DEV_NVLINK_REPLAY_ERROR_COUNT_TOTAL = 52
	// FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L0 as defined in nvml/nvml.h
	FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L0 = 53
	// FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L1 as defined in nvml/nvml.h
	FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L1 = 54
	// FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L2 as defined in nvml/nvml.h
	FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L2 = 55
	// FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L3 as defined in nvml/nvml.h
	FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L3 = 56
	// FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L4 as defined in nvml/nvml.h
	FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L4 = 57
	// FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L5 as defined in nvml/nvml.h
	FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L5 = 58
	// FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_TOTAL as defined in nvml/nvml.h
	FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_TOTAL = 59
	// FI_DEV_NVLINK_BANDWIDTH_C0_L0 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C0_L0 = 60
	// FI_DEV_NVLINK_BANDWIDTH_C0_L1 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C0_L1 = 61
	// FI_DEV_NVLINK_BANDWIDTH_C0_L2 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C0_L2 = 62
	// FI_DEV_NVLINK_BANDWIDTH_C0_L3 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C0_L3 = 63
	// FI_DEV_NVLINK_BANDWIDTH_C0_L4 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C0_L4 = 64
	// FI_DEV_NVLINK_BANDWIDTH_C0_L5 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C0_L5 = 65
	// FI_DEV_NVLINK_BANDWIDTH_C0_TOTAL as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C0_TOTAL = 66
	// FI_DEV_NVLINK_BANDWIDTH_C1_L0 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C1_L0 = 67
	// FI_DEV_NVLINK_BANDWIDTH_C1_L1 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C1_L1 = 68
	// FI_DEV_NVLINK_BANDWIDTH_C1_L2 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C1_L2 = 69
	// FI_DEV_NVLINK_BANDWIDTH_C1_L3 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C1_L3 = 70
	// FI_DEV_NVLINK_BANDWIDTH_C1_L4 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C1_L4 = 71
	// FI_DEV_NVLINK_BANDWIDTH_C1_L5 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C1_L5 = 72
	// FI_DEV_NVLINK_BANDWIDTH_C1_TOTAL as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C1_TOTAL = 73
	// FI_DEV_PERF_POLICY_POWER as defined in nvml/nvml.h
	FI_DEV_PERF_POLICY_POWER = 74
	// FI_DEV_PERF_POLICY_THERMAL as defined in nvml/nvml.h
	FI_DEV_PERF_POLICY_THERMAL = 75
	// FI_DEV_PERF_POLICY_SYNC_BOOST as defined in nvml/nvml.h
	FI_DEV_PERF_POLICY_SYNC_BOOST = 76
	// FI_DEV_PERF_POLICY_BOARD_LIMIT as defined in nvml/nvml.h
	FI_DEV_PERF_POLICY_BOARD_LIMIT = 77
	// FI_DEV_PERF_POLICY_LOW_UTILIZATION as defined in nvml/nvml.h
	FI_DEV_PERF_POLICY_LOW_UTILIZATION = 78
	// FI_DEV_PERF_POLICY_RELIABILITY as defined in nvml/nvml.h
	FI_DEV_PERF_POLICY_RELIABILITY = 79
	// FI_DEV_PERF_POLICY_TOTAL_APP_CLOCKS as defined in nvml/nvml.h
	FI_DEV_PERF_POLICY_TOTAL_APP_CLOCKS = 80
	// FI_DEV_PERF_POLICY_TOTAL_BASE_CLOCKS as defined in nvml/nvml.h
	FI_DEV_PERF_POLICY_TOTAL_BASE_CLOCKS = 81
	// FI_DEV_MEMORY_TEMP as defined in nvml/nvml.h
	FI_DEV_MEMORY_TEMP = 82
	// FI_DEV_TOTAL_ENERGY_CONSUMPTION as defined in nvml/nvml.h
	FI_DEV_TOTAL_ENERGY_CONSUMPTION = 83
	// FI_DEV_NVLINK_SPEED_MBPS_L0 as defined in nvml/nvml.h
	FI_DEV_NVLINK_SPEED_MBPS_L0 = 84
	// FI_DEV_NVLINK_SPEED_MBPS_L1 as defined in nvml/nvml.h
	FI_DEV_NVLINK_SPEED_MBPS_L1 = 85
	// FI_DEV_NVLINK_SPEED_MBPS_L2 as defined in nvml/nvml.h
	FI_DEV_NVLINK_SPEED_MBPS_L2 = 86
	// FI_DEV_NVLINK_SPEED_MBPS_L3 as defined in nvml/nvml.h
	FI_DEV_NVLINK_SPEED_MBPS_L3 = 87
	// FI_DEV_NVLINK_SPEED_MBPS_L4 as defined in nvml/nvml.h
	FI_DEV_NVLINK_SPEED_MBPS_L4 = 88
	// FI_DEV_NVLINK_SPEED_MBPS_L5 as defined in nvml/nvml.h
	FI_DEV_NVLINK_SPEED_MBPS_L5 = 89
	// FI_DEV_NVLINK_SPEED_MBPS_COMMON as defined in nvml/nvml.h
	FI_DEV_NVLINK_SPEED_MBPS_COMMON = 90
	// FI_DEV_NVLINK_LINK_COUNT as defined in nvml/nvml.h
	FI_DEV_NVLINK_LINK_COUNT = 91
	// FI_DEV_RETIRED_PENDING_SBE as defined in nvml/nvml.h
	FI_DEV_RETIRED_PENDING_SBE = 92
	// FI_DEV_RETIRED_PENDING_DBE as defined in nvml/nvml.h
	FI_DEV_RETIRED_PENDING_DBE = 93
	// FI_DEV_PCIE_REPLAY_COUNTER as defined in nvml/nvml.h
	FI_DEV_PCIE_REPLAY_COUNTER = 94
	// FI_DEV_PCIE_REPLAY_ROLLOVER_COUNTER as defined in nvml/nvml.h
	FI_DEV_PCIE_REPLAY_ROLLOVER_COUNTER = 95
	// FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L6 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L6 = 96
	// FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L7 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L7 = 97
	// FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L8 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L8 = 98
	// FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L9 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L9 = 99
	// FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L10 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L10 = 100
	// FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L11 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_FLIT_ERROR_COUNT_L11 = 101
	// FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L6 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L6 = 102
	// FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L7 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L7 = 103
	// FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L8 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L8 = 104
	// FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L9 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L9 = 105
	// FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L10 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L10 = 106
	// FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L11 as defined in nvml/nvml.h
	FI_DEV_NVLINK_CRC_DATA_ERROR_COUNT_L11 = 107
	// FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L6 as defined in nvml/nvml.h
	FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L6 = 108
	// FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L7 as defined in nvml/nvml.h
	FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L7 = 109
	// FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L8 as defined in nvml/nvml.h
	FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L8 = 110
	// FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L9 as defined in nvml/nvml.h
	FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L9 = 111
	// FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L10 as defined in nvml/nvml.h
	FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L10 = 112
	// FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L11 as defined in nvml/nvml.h
	FI_DEV_NVLINK_REPLAY_ERROR_COUNT_L11 = 113
	// FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L6 as defined in nvml/nvml.h
	FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L6 = 114
	// FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L7 as defined in nvml/nvml.h
	FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L7 = 115
	// FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L8 as defined in nvml/nvml.h
	FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L8 = 116
	// FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L9 as defined in nvml/nvml.h
	FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L9 = 117
	// FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L10 as defined in nvml/nvml.h
	FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L10 = 118
	// FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L11 as defined in nvml/nvml.h
	FI_DEV_NVLINK_RECOVERY_ERROR_COUNT_L11 = 119
	// FI_DEV_NVLINK_BANDWIDTH_C0_L6 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C0_L6 = 120
	// FI_DEV_NVLINK_BANDWIDTH_C0_L7 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C0_L7 = 121
	// FI_DEV_NVLINK_BANDWIDTH_C0_L8 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C0_L8 = 122
	// FI_DEV_NVLINK_BANDWIDTH_C0_L9 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C0_L9 = 123
	// FI_DEV_NVLINK_BANDWIDTH_C0_L10 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C0_L10 = 124
	// FI_DEV_NVLINK_BANDWIDTH_C0_L11 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C0_L11 = 125
	// FI_DEV_NVLINK_BANDWIDTH_C1_L6 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C1_L6 = 126
	// FI_DEV_NVLINK_BANDWIDTH_C1_L7 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C1_L7 = 127
	// FI_DEV_NVLINK_BANDWIDTH_C1_L8 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C1_L8 = 128
	// FI_DEV_NVLINK_BANDWIDTH_C1_L9 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C1_L9 = 129
	// FI_DEV_NVLINK_BANDWIDTH_C1_L10 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C1_L10 = 130
	// FI_DEV_NVLINK_BANDWIDTH_C1_L11 as defined in nvml/nvml.h
	FI_DEV_NVLINK_BANDWIDTH_C1_L11 = 131
	// FI_DEV_NVLINK_SPEED_MBPS_L6 as defined in nvml/nvml.h
	FI_DEV_NVLINK_SPEED_MBPS_L6 = 132
	// FI_DEV_NVLINK_SPEED_MBPS_L7 as defined in nvml/nvml.h
	FI_DEV_NVLINK_SPEED_MBPS_L7 = 133
	// FI_DEV_NVLINK_SPEED_MBPS_L8 as defined in nvml/nvml.h
	FI_DEV_NVLINK_SPEED_MBPS_L8 = 134
	// FI_DEV_NVLINK_SPEED_MBPS_L9 as defined in nvml/nvml.h
	FI_DEV_NVLINK_SPEED_MBPS_L9 = 135
	// FI_DEV_NVLINK_SPEED_MBPS_L10 as defined in nvml/nvml.h
	FI_DEV_NVLINK_SPEED_MBPS_L10 = 136
	// FI_DEV_NVLINK_SPEED_MBPS_L11 as defined in nvml/nvml.h
	FI_DEV_NVLINK_SPEED_MBPS_L11 = 137
	// FI_DEV_NVLINK_THROUGHPUT_DATA_TX as defined in nvml/nvml.h
	FI_DEV_NVLINK_THROUGHPUT_DATA_TX = 138
	// FI_DEV_NVLINK_THROUGHPUT_DATA_RX as defined in nvml/nvml.h
	FI_DEV_NVLINK_THROUGHPUT_DATA_RX = 139
	// FI_DEV_NVLINK_THROUGHPUT_RAW_TX as defined in nvml/nvml.h
	FI_DEV_NVLINK_THROUGHPUT_RAW_TX = 140
	// FI_DEV_NVLINK_THROUGHPUT_RAW_RX as defined in nvml/nvml.h
	FI_DEV_NVLINK_THROUGHPUT_RAW_RX = 141
	// FI_DEV_REMAPPED_COR as defined in nvml/nvml.h
	FI_DEV_REMAPPED_COR = 142
	// FI_DEV_REMAPPED_UNC as defined in nvml/nvml.h
	FI_DEV_REMAPPED_UNC = 143
	// FI_DEV_REMAPPED_PENDING as defined in nvml/nvml.h
	FI_DEV_REMAPPED_PENDING = 144
	// FI_DEV_REMAPPED_FAILURE as defined in nvml/nvml.h
	FI_DEV_REMAPPED_FAILURE = 145
	// FI_DEV_NVLINK_REMOTE_NVLINK_ID as defined in nvml/nvml.h
	FI_DEV_NVLINK_REMOTE_NVLINK_ID = 146
	// FI_DEV_NVSWITCH_CONNECTED_LINK_COUNT as defined in nvml/nvml.h
	FI_DEV_NVSWITCH_CONNECTED_LINK_COUNT = 147
	// FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L0 as defined in nvml/nvml.h
	FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L0 = 148
	// FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L1 as defined in nvml/nvml.h
	FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L1 = 149
	// FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L2 as defined in nvml/nvml.h
	FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L2 = 150
	// FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L3 as defined in nvml/nvml.h
	FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L3 = 151
	// FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L4 as defined in nvml/nvml.h
	FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L4 = 152
	// FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L5 as defined in nvml/nvml.h
	FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L5 = 153
	// FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L6 as defined in nvml/nvml.h
	FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L6 = 154
	// FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L7 as defined in nvml/nvml.h
	FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L7 = 155
	// FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L8 as defined in nvml/nvml.h
	FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L8 = 156
	// FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L9 as defined in nvml/nvml.h
	FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L9 = 157
	// FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L10 as defined in nvml/nvml.h
	FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L10 = 158
	// FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L11 as defined in nvml/nvml.h
	FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_L11 = 159
	// FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_TOTAL as defined in nvml/nvml.h
	FI_DEV_NVLINK_ECC_DATA_ERROR_COUNT_TOTAL = 160
	// FI_DEV_NVLINK_ERROR_DL_REPLAY as defined in nvml/nvml.h
	FI_DEV_NVLINK_ERROR_DL_REPLAY = 161
	// FI_DEV_NVLINK_ERROR_DL_RECOVERY as defined in nvml/nvml.h
	FI_DEV_NVLINK_ERROR_DL_RECOVERY = 162
	// FI_DEV_NVLINK_ERROR_DL_CRC as defined in nvml/nvml.h
	FI_DEV_NVLINK_ERROR_DL_CRC = 163
	// FI_DEV_NVLINK_GET_SPEED as defined in nvml/nvml.h
	FI_DEV_NVLINK_GET_SPEED = 164
	// FI_DEV_NVLINK_GET_STATE as defined in nvml/nvml.h
	FI_DEV_NVLINK_GET_STATE = 165
	// FI_DEV_NVLINK_GET_VERSION as defined in nvml/nvml.h
	FI_DEV_NVLINK_GET_VERSION = 166
	// FI_DEV_NVLINK_GET_POWER_STATE as defined in nvml/nvml.h
	FI_DEV_NVLINK_GET_POWER_STATE = 167
	// FI_DEV_NVLINK_GET_POWER_THRESHOLD as defined in nvml/nvml.h
	FI_DEV_NVLINK_GET_POWER_THRESHOLD = 168
	// FI_DEV_PCIE_L0_TO_RECOVERY_COUNTER as defined in nvml/nvml.h
	FI_DEV_PCIE_L0_TO_RECOVERY_COUNTER = 169
	// FI_MAX as defined in nvml/nvml.h
	FI_MAX = 170
	// EventTypeSingleBitEccError as defined in nvml/nvml.h
	EventTypeSingleBitEccError = 1
	// EventTypeDoubleBitEccError as defined in nvml/nvml.h
	EventTypeDoubleBitEccError = 2
	// EventTypePState as defined in nvml/nvml.h
	EventTypePState = 4
	// EventTypeXidCriticalError as defined in nvml/nvml.h
	EventTypeXidCriticalError = 8
	// EventTypeClock as defined in nvml/nvml.h
	EventTypeClock = 16
	// EventTypePowerSourceChange as defined in nvml/nvml.h
	EventTypePowerSourceChange = 128
	// EventMigConfigChange as defined in nvml/nvml.h
	EventMigConfigChange = 256
	// EventTypeNone as defined in nvml/nvml.h
	EventTypeNone = 0
	// EventTypeAll as defined in nvml/nvml.h
	EventTypeAll = 415
	// ClocksThrottleReasonGpuIdle as defined in nvml/nvml.h
	ClocksThrottleReasonGpuIdle = 1
	// ClocksThrottleReasonApplicationsClocksSetting as defined in nvml/nvml.h
	ClocksThrottleReasonApplicationsClocksSetting = 2
	// ClocksThrottleReasonUserDefinedClocks as defined in nvml/nvml.h
	ClocksThrottleReasonUserDefinedClocks = 2
	// ClocksThrottleReasonSwPowerCap as defined in nvml/nvml.h
	ClocksThrottleReasonSwPowerCap = 4
	// ClocksThrottleReasonHwSlowdown as defined in nvml/nvml.h
	ClocksThrottleReasonHwSlowdown = 8
	// ClocksThrottleReasonSyncBoost as defined in nvml/nvml.h
	ClocksThrottleReasonSyncBoost = 16
	// ClocksThrottleReasonSwThermalSlowdown as defined in nvml/nvml.h
	ClocksThrottleReasonSwThermalSlowdown = 32
	// ClocksThrottleReasonHwThermalSlowdown as defined in nvml/nvml.h
	ClocksThrottleReasonHwThermalSlowdown = 64
	// ClocksThrottleReasonHwPowerBrakeSlowdown as defined in nvml/nvml.h
	ClocksThrottleReasonHwPowerBrakeSlowdown = 128
	// ClocksThrottleReasonDisplayClockSetting as defined in nvml/nvml.h
	ClocksThrottleReasonDisplayClockSetting = 256
	// ClocksThrottleReasonNone as defined in nvml/nvml.h
	ClocksThrottleReasonNone = 0
	// ClocksThrottleReasonAll as defined in nvml/nvml.h
	ClocksThrottleReasonAll = 511
	// NVFBC_SESSION_FLAG_DIFFMAP_ENABLED as defined in nvml/nvml.h
	NVFBC_SESSION_FLAG_DIFFMAP_ENABLED = 1
	// NVFBC_SESSION_FLAG_CLASSIFICATIONMAP_ENABLED as defined in nvml/nvml.h
	NVFBC_SESSION_FLAG_CLASSIFICATIONMAP_ENABLED = 2
	// NVFBC_SESSION_FLAG_CAPTURE_WITH_WAIT_NO_WAIT as defined in nvml/nvml.h
	NVFBC_SESSION_FLAG_CAPTURE_WITH_WAIT_NO_WAIT = 4
	// NVFBC_SESSION_FLAG_CAPTURE_WITH_WAIT_INFINITE as defined in nvml/nvml.h
	NVFBC_SESSION_FLAG_CAPTURE_WITH_WAIT_INFINITE = 8
	// NVFBC_SESSION_FLAG_CAPTURE_WITH_WAIT_TIMEOUT as defined in nvml/nvml.h
	NVFBC_SESSION_FLAG_CAPTURE_WITH_WAIT_TIMEOUT = 16
	// GPU_FABRIC_UUID_LEN as defined in nvml/nvml.h
	GPU_FABRIC_UUID_LEN = 16
	// GPU_FABRIC_STATE_NOT_SUPPORTED as defined in nvml/nvml.h
	GPU_FABRIC_STATE_NOT_SUPPORTED = 0
	// GPU_FABRIC_STATE_NOT_STARTED as defined in nvml/nvml.h
	GPU_FABRIC_STATE_NOT_STARTED = 1
	// GPU_FABRIC_STATE_IN_PROGRESS as defined in nvml/nvml.h
	GPU_FABRIC_STATE_IN_PROGRESS = 2
	// GPU_FABRIC_STATE_COMPLETED as defined in nvml/nvml.h
	GPU_FABRIC_STATE_COMPLETED = 3
	// INIT_FLAG_NO_GPUS as defined in nvml/nvml.h
	INIT_FLAG_NO_GPUS = 1
	// INIT_FLAG_NO_ATTACH as defined in nvml/nvml.h
	INIT_FLAG_NO_ATTACH = 2
	// DEVICE_INFOROM_VERSION_BUFFER_SIZE as defined in nvml/nvml.h
	DEVICE_INFOROM_VERSION_BUFFER_SIZE = 16
	// DEVICE_UUID_BUFFER_SIZE as defined in nvml/nvml.h
	DEVICE_UUID_BUFFER_SIZE = 80
	// DEVICE_UUID_V2_BUFFER_SIZE as defined in nvml/nvml.h
	DEVICE_UUID_V2_BUFFER_SIZE = 96
	// DEVICE_PART_NUMBER_BUFFER_SIZE as defined in nvml/nvml.h
	DEVICE_PART_NUMBER_BUFFER_SIZE = 80
	// SYSTEM_DRIVER_VERSION_BUFFER_SIZE as defined in nvml/nvml.h
	SYSTEM_DRIVER_VERSION_BUFFER_SIZE = 80
	// SYSTEM_NVML_VERSION_BUFFER_SIZE as defined in nvml/nvml.h
	SYSTEM_NVML_VERSION_BUFFER_SIZE = 80
	// DEVICE_NAME_BUFFER_SIZE as defined in nvml/nvml.h
	DEVICE_NAME_BUFFER_SIZE = 64
	// DEVICE_NAME_V2_BUFFER_SIZE as defined in nvml/nvml.h
	DEVICE_NAME_V2_BUFFER_SIZE = 96
	// DEVICE_SERIAL_BUFFER_SIZE as defined in nvml/nvml.h
	DEVICE_SERIAL_BUFFER_SIZE = 30
	// DEVICE_VBIOS_VERSION_BUFFER_SIZE as defined in nvml/nvml.h
	DEVICE_VBIOS_VERSION_BUFFER_SIZE = 32
	// AFFINITY_SCOPE_NODE as defined in nvml/nvml.h
	AFFINITY_SCOPE_NODE = 0
	// AFFINITY_SCOPE_SOCKET as defined in nvml/nvml.h
	AFFINITY_SCOPE_SOCKET = 1
	// DEVICE_MIG_DISABLE as defined in nvml/nvml.h
	DEVICE_MIG_DISABLE = 0
	// DEVICE_MIG_ENABLE as defined in nvml/nvml.h
	DEVICE_MIG_ENABLE = 1
	// GPU_INSTANCE_PROFILE_1_SLICE as defined in nvml/nvml.h
	GPU_INSTANCE_PROFILE_1_SLICE = 0
	// GPU_INSTANCE_PROFILE_2_SLICE as defined in nvml/nvml.h
	GPU_INSTANCE_PROFILE_2_SLICE = 1
	// GPU_INSTANCE_PROFILE_3_SLICE as defined in nvml/nvml.h
	GPU_INSTANCE_PROFILE_3_SLICE = 2
	// GPU_INSTANCE_PROFILE_4_SLICE as defined in nvml/nvml.h
	GPU_INSTANCE_PROFILE_4_SLICE = 3
	// GPU_INSTANCE_PROFILE_7_SLICE as defined in nvml/nvml.h
	GPU_INSTANCE_PROFILE_7_SLICE = 4
	// GPU_INSTANCE_PROFILE_8_SLICE as defined in nvml/nvml.h
	GPU_INSTANCE_PROFILE_8_SLICE = 5
	// GPU_INSTANCE_PROFILE_6_SLICE as defined in nvml/nvml.h
	GPU_INSTANCE_PROFILE_6_SLICE = 6
	// GPU_INSTANCE_PROFILE_1_SLICE_REV1 as defined in nvml/nvml.h
	GPU_INSTANCE_PROFILE_1_SLICE_REV1 = 7
	// GPU_INSTANCE_PROFILE_2_SLICE_REV1 as defined in nvml/nvml.h
	GPU_INSTANCE_PROFILE_2_SLICE_REV1 = 8
	// GPU_INSTANCE_PROFILE_1_SLICE_REV2 as defined in nvml/nvml.h
	GPU_INSTANCE_PROFILE_1_SLICE_REV2 = 9
	// GPU_INSTANCE_PROFILE_COUNT as defined in nvml/nvml.h
	GPU_INSTANCE_PROFILE_COUNT = 10
	// COMPUTE_INSTANCE_PROFILE_1_SLICE as defined in nvml/nvml.h
	COMPUTE_INSTANCE_PROFILE_1_SLICE = 0
	// COMPUTE_INSTANCE_PROFILE_2_SLICE as defined in nvml/nvml.h
	COMPUTE_INSTANCE_PROFILE_2_SLICE = 1
	// COMPUTE_INSTANCE_PROFILE_3_SLICE as defined in nvml/nvml.h
	COMPUTE_INSTANCE_PROFILE_3_SLICE = 2
	// COMPUTE_INSTANCE_PROFILE_4_SLICE as defined in nvml/nvml.h
	COMPUTE_INSTANCE_PROFILE_4_SLICE = 3
	// COMPUTE_INSTANCE_PROFILE_7_SLICE as defined in nvml/nvml.h
	COMPUTE_INSTANCE_PROFILE_7_SLICE = 4
	// COMPUTE_INSTANCE_PROFILE_8_SLICE as defined in nvml/nvml.h
	COMPUTE_INSTANCE_PROFILE_8_SLICE = 5
	// COMPUTE_INSTANCE_PROFILE_6_SLICE as defined in nvml/nvml.h
	COMPUTE_INSTANCE_PROFILE_6_SLICE = 6
	// COMPUTE_INSTANCE_PROFILE_1_SLICE_REV1 as defined in nvml/nvml.h
	COMPUTE_INSTANCE_PROFILE_1_SLICE_REV1 = 7
	// COMPUTE_INSTANCE_PROFILE_COUNT as defined in nvml/nvml.h
	COMPUTE_INSTANCE_PROFILE_COUNT = 8
	// COMPUTE_INSTANCE_ENGINE_PROFILE_SHARED as defined in nvml/nvml.h
	COMPUTE_INSTANCE_ENGINE_PROFILE_SHARED = 0
	// COMPUTE_INSTANCE_ENGINE_PROFILE_COUNT as defined in nvml/nvml.h
	COMPUTE_INSTANCE_ENGINE_PROFILE_COUNT = 1
	// GPM_METRICS_GET_VERSION as defined in nvml/nvml.h
	GPM_METRICS_GET_VERSION = 1
	// GPM_SUPPORT_VERSION as defined in nvml/nvml.h
	GPM_SUPPORT_VERSION = 1
	// COUNTER_COLLECTION_UNIT_STREAM_STATE_DISABLE as defined in nvml/nvml.h
	COUNTER_COLLECTION_UNIT_STREAM_STATE_DISABLE = 0
	// COUNTER_COLLECTION_UNIT_STREAM_STATE_ENABLE as defined in nvml/nvml.h
	COUNTER_COLLECTION_UNIT_STREAM_STATE_ENABLE = 1
	// NVLINK_POWER_STATE_HIGH_SPEED as defined in nvml/nvml.h
	NVLINK_POWER_STATE_HIGH_SPEED = 0
	// NVLINK_POWER_STATE_LOW as defined in nvml/nvml.h
	NVLINK_POWER_STATE_LOW = 1
	// NVLINK_LOW_POWER_THRESHOLD_MIN as defined in nvml/nvml.h
	NVLINK_LOW_POWER_THRESHOLD_MIN = 1
	// NVLINK_LOW_POWER_THRESHOLD_MAX as defined in nvml/nvml.h
	NVLINK_LOW_POWER_THRESHOLD_MAX = 8191
	// NVLINK_LOW_POWER_THRESHOLD_RESET as defined in nvml/nvml.h
	NVLINK_LOW_POWER_THRESHOLD_RESET = 4294967295
)

// BridgeChipType as declared in nvml/nvml.h
type BridgeChipType int32

// BridgeChipType enumeration from nvml/nvml.h
const (
	BRIDGE_CHIP_PLX  BridgeChipType = iota
	BRIDGE_CHIP_BRO4 BridgeChipType = 1
)

// NvLinkUtilizationCountUnits as declared in nvml/nvml.h
type NvLinkUtilizationCountUnits int32

// NvLinkUtilizationCountUnits enumeration from nvml/nvml.h
const (
	NVLINK_COUNTER_UNIT_CYCLES   NvLinkUtilizationCountUnits = iota
	NVLINK_COUNTER_UNIT_PACKETS  NvLinkUtilizationCountUnits = 1
	NVLINK_COUNTER_UNIT_BYTES    NvLinkUtilizationCountUnits = 2
	NVLINK_COUNTER_UNIT_RESERVED NvLinkUtilizationCountUnits = 3
	NVLINK_COUNTER_UNIT_COUNT    NvLinkUtilizationCountUnits = 4
)

// NvLinkUtilizationCountPktTypes as declared in nvml/nvml.h
type NvLinkUtilizationCountPktTypes int32

// NvLinkUtilizationCountPktTypes enumeration from nvml/nvml.h
const (
	NVLINK_COUNTER_PKTFILTER_NOP        NvLinkUtilizationCountPktTypes = 1
	NVLINK_COUNTER_PKTFILTER_READ       NvLinkUtilizationCountPktTypes = 2
	NVLINK_COUNTER_PKTFILTER_WRITE      NvLinkUtilizationCountPktTypes = 4
	NVLINK_COUNTER_PKTFILTER_RATOM      NvLinkUtilizationCountPktTypes = 8
	NVLINK_COUNTER_PKTFILTER_NRATOM     NvLinkUtilizationCountPktTypes = 16
	NVLINK_COUNTER_PKTFILTER_FLUSH      NvLinkUtilizationCountPktTypes = 32
	NVLINK_COUNTER_PKTFILTER_RESPDATA   NvLinkUtilizationCountPktTypes = 64
	NVLINK_COUNTER_PKTFILTER_RESPNODATA NvLinkUtilizationCountPktTypes = 128
	NVLINK_COUNTER_PKTFILTER_ALL        NvLinkUtilizationCountPktTypes = 255
)

// NvLinkCapability as declared in nvml/nvml.h
type NvLinkCapability int32

// NvLinkCapability enumeration from nvml/nvml.h
const (
	NVLINK_CAP_P2P_SUPPORTED  NvLinkCapability = iota
	NVLINK_CAP_SYSMEM_ACCESS  NvLinkCapability = 1
	NVLINK_CAP_P2P_ATOMICS    NvLinkCapability = 2
	NVLINK_CAP_SYSMEM_ATOMICS NvLinkCapability = 3
	NVLINK_CAP_SLI_BRIDGE     NvLinkCapability = 4
	NVLINK_CAP_VALID          NvLinkCapability = 5
	NVLINK_CAP_COUNT          NvLinkCapability = 6
)

// NvLinkErrorCounter as declared in nvml/nvml.h
type NvLinkErrorCounter int32

// NvLinkErrorCounter enumeration from nvml/nvml.h
const (
	NVLINK_ERROR_DL_REPLAY   NvLinkErrorCounter = iota
	NVLINK_ERROR_DL_RECOVERY NvLinkErrorCounter = 1
	NVLINK_ERROR_DL_CRC_FLIT NvLinkErrorCounter = 2
	NVLINK_ERROR_DL_CRC_DATA NvLinkErrorCounter = 3
	NVLINK_ERROR_DL_ECC_DATA NvLinkErrorCounter = 4
	NVLINK_ERROR_COUNT       NvLinkErrorCounter = 5
)

// IntNvLinkDeviceType as declared in nvml/nvml.h
type IntNvLinkDeviceType int32

// IntNvLinkDeviceType enumeration from nvml/nvml.h
const (
	NVLINK_DEVICE_TYPE_GPU     IntNvLinkDeviceType = iota
	NVLINK_DEVICE_TYPE_IBMNPU  IntNvLinkDeviceType = 1
	NVLINK_DEVICE_TYPE_SWITCH  IntNvLinkDeviceType = 2
	NVLINK_DEVICE_TYPE_UNKNOWN IntNvLinkDeviceType = 255
)

// GpuTopologyLevel as declared in nvml/nvml.h
type GpuTopologyLevel int32

// GpuTopologyLevel enumeration from nvml/nvml.h
const (
	TOPOLOGY_INTERNAL   GpuTopologyLevel = iota
	TOPOLOGY_SINGLE     GpuTopologyLevel = 10
	TOPOLOGY_MULTIPLE   GpuTopologyLevel = 20
	TOPOLOGY_HOSTBRIDGE GpuTopologyLevel = 30
	TOPOLOGY_NODE       GpuTopologyLevel = 40
	TOPOLOGY_SYSTEM     GpuTopologyLevel = 50
)

// GpuP2PStatus as declared in nvml/nvml.h
type GpuP2PStatus int32

// GpuP2PStatus enumeration from nvml/nvml.h
const (
	P2P_STATUS_OK                         GpuP2PStatus = iota
	P2P_STATUS_CHIPSET_NOT_SUPPORED       GpuP2PStatus = 1
	P2P_STATUS_GPU_NOT_SUPPORTED          GpuP2PStatus = 2
	P2P_STATUS_IOH_TOPOLOGY_NOT_SUPPORTED GpuP2PStatus = 3
	P2P_STATUS_DISABLED_BY_REGKEY         GpuP2PStatus = 4
	P2P_STATUS_NOT_SUPPORTED              GpuP2PStatus = 5
	P2P_STATUS_UNKNOWN                    GpuP2PStatus = 6
)

// GpuP2PCapsIndex as declared in nvml/nvml.h
type GpuP2PCapsIndex int32

// GpuP2PCapsIndex enumeration from nvml/nvml.h
const (
	P2P_CAPS_INDEX_READ    GpuP2PCapsIndex = iota
	P2P_CAPS_INDEX_WRITE   GpuP2PCapsIndex = 1
	P2P_CAPS_INDEX_NVLINK  GpuP2PCapsIndex = 2
	P2P_CAPS_INDEX_ATOMICS GpuP2PCapsIndex = 3
	P2P_CAPS_INDEX_PROP    GpuP2PCapsIndex = 4
	P2P_CAPS_INDEX_UNKNOWN GpuP2PCapsIndex = 5
)

// SamplingType as declared in nvml/nvml.h
type SamplingType int32

// SamplingType enumeration from nvml/nvml.h
const (
	TOTAL_POWER_SAMPLES        SamplingType = iota
	GPU_UTILIZATION_SAMPLES    SamplingType = 1
	MEMORY_UTILIZATION_SAMPLES SamplingType = 2
	ENC_UTILIZATION_SAMPLES    SamplingType = 3
	DEC_UTILIZATION_SAMPLES    SamplingType = 4
	PROCESSOR_CLK_SAMPLES      SamplingType = 5
	MEMORY_CLK_SAMPLES         SamplingType = 6
	SAMPLINGTYPE_COUNT         SamplingType = 7
)

// PcieUtilCounter as declared in nvml/nvml.h
type PcieUtilCounter int32

// PcieUtilCounter enumeration from nvml/nvml.h
const (
	PCIE_UTIL_TX_BYTES PcieUtilCounter = iota
	PCIE_UTIL_RX_BYTES PcieUtilCounter = 1
	PCIE_UTIL_COUNT    PcieUtilCounter = 2
)

// ValueType as declared in nvml/nvml.h
type ValueType int32

// ValueType enumeration from nvml/nvml.h
const (
	VALUE_TYPE_DOUBLE             ValueType = iota
	VALUE_TYPE_UNSIGNED_INT       ValueType = 1
	VALUE_TYPE_UNSIGNED_LONG      ValueType = 2
	VALUE_TYPE_UNSIGNED_LONG_LONG ValueType = 3
	VALUE_TYPE_SIGNED_LONG_LONG   ValueType = 4
	VALUE_TYPE_COUNT              ValueType = 5
)

// PerfPolicyType as declared in nvml/nvml.h
type PerfPolicyType int32

// PerfPolicyType enumeration from nvml/nvml.h
const (
	PERF_POLICY_POWER             PerfPolicyType = iota
	PERF_POLICY_THERMAL           PerfPolicyType = 1
	PERF_POLICY_SYNC_BOOST        PerfPolicyType = 2
	PERF_POLICY_BOARD_LIMIT       PerfPolicyType = 3
	PERF_POLICY_LOW_UTILIZATION   PerfPolicyType = 4
	PERF_POLICY_RELIABILITY       PerfPolicyType = 5
	PERF_POLICY_TOTAL_APP_CLOCKS  PerfPolicyType = 10
	PERF_POLICY_TOTAL_BASE_CLOCKS PerfPolicyType = 11
	PERF_POLICY_COUNT             PerfPolicyType = 12
)

// EnableState as declared in nvml/nvml.h
type EnableState int32

// EnableState enumeration from nvml/nvml.h
const (
	FEATURE_DISABLED EnableState = iota
	FEATURE_ENABLED  EnableState = 1
)

// BrandType as declared in nvml/nvml.h
type BrandType int32

// BrandType enumeration from nvml/nvml.h
const (
	BRAND_UNKNOWN             BrandType = iota
	BRAND_QUADRO              BrandType = 1
	BRAND_TESLA               BrandType = 2
	BRAND_NVS                 BrandType = 3
	BRAND_GRID                BrandType = 4
	BRAND_GEFORCE             BrandType = 5
	BRAND_TITAN               BrandType = 6
	BRAND_NVIDIA_VAPPS        BrandType = 7
	BRAND_NVIDIA_VPC          BrandType = 8
	BRAND_NVIDIA_VCS          BrandType = 9
	BRAND_NVIDIA_VWS          BrandType = 10
	BRAND_NVIDIA_CLOUD_GAMING BrandType = 11
	BRAND_NVIDIA_VGAMING      BrandType = 11
	BRAND_QUADRO_RTX          BrandType = 12
	BRAND_NVIDIA_RTX          BrandType = 13
	BRAND_NVIDIA              BrandType = 14
	BRAND_GEFORCE_RTX         BrandType = 15
	BRAND_TITAN_RTX           BrandType = 16
	BRAND_COUNT               BrandType = 17
)

// TemperatureThresholds as declared in nvml/nvml.h
type TemperatureThresholds int32

// TemperatureThresholds enumeration from nvml/nvml.h
const (
	TEMPERATURE_THRESHOLD_SHUTDOWN      TemperatureThresholds = iota
	TEMPERATURE_THRESHOLD_SLOWDOWN      TemperatureThresholds = 1
	TEMPERATURE_THRESHOLD_MEM_MAX       TemperatureThresholds = 2
	TEMPERATURE_THRESHOLD_GPU_MAX       TemperatureThresholds = 3
	TEMPERATURE_THRESHOLD_ACOUSTIC_MIN  TemperatureThresholds = 4
	TEMPERATURE_THRESHOLD_ACOUSTIC_CURR TemperatureThresholds = 5
	TEMPERATURE_THRESHOLD_ACOUSTIC_MAX  TemperatureThresholds = 6
	TEMPERATURE_THRESHOLD_COUNT         TemperatureThresholds = 7
)

// TemperatureSensors as declared in nvml/nvml.h
type TemperatureSensors int32

// TemperatureSensors enumeration from nvml/nvml.h
const (
	TEMPERATURE_GPU   TemperatureSensors = iota
	TEMPERATURE_COUNT TemperatureSensors = 1
)

// ComputeMode as declared in nvml/nvml.h
type ComputeMode int32

// ComputeMode enumeration from nvml/nvml.h
const (
	COMPUTEMODE_DEFAULT           ComputeMode = iota
	COMPUTEMODE_EXCLUSIVE_THREAD  ComputeMode = 1
	COMPUTEMODE_PROHIBITED        ComputeMode = 2
	COMPUTEMODE_EXCLUSIVE_PROCESS ComputeMode = 3
	COMPUTEMODE_COUNT             ComputeMode = 4
)

// MemoryErrorType as declared in nvml/nvml.h
type MemoryErrorType int32

// MemoryErrorType enumeration from nvml/nvml.h
const (
	MEMORY_ERROR_TYPE_CORRECTED   MemoryErrorType = iota
	MEMORY_ERROR_TYPE_UNCORRECTED MemoryErrorType = 1
	MEMORY_ERROR_TYPE_COUNT       MemoryErrorType = 2
)

// EccCounterType as declared in nvml/nvml.h
type EccCounterType int32

// EccCounterType enumeration from nvml/nvml.h
const (
	VOLATILE_ECC           EccCounterType = iota
	AGGREGATE_ECC          EccCounterType = 1
	ECC_COUNTER_TYPE_COUNT EccCounterType = 2
)

// ClockType as declared in nvml/nvml.h
type ClockType int32

// ClockType enumeration from nvml/nvml.h
const (
	CLOCK_GRAPHICS ClockType = iota
	CLOCK_SM       ClockType = 1
	CLOCK_MEM      ClockType = 2
	CLOCK_VIDEO    ClockType = 3
	CLOCK_COUNT    ClockType = 4
)

// ClockId as declared in nvml/nvml.h
type ClockId int32

// ClockId enumeration from nvml/nvml.h
const (
	CLOCK_ID_CURRENT            ClockId = iota
	CLOCK_ID_APP_CLOCK_TARGET   ClockId = 1
	CLOCK_ID_APP_CLOCK_DEFAULT  ClockId = 2
	CLOCK_ID_CUSTOMER_BOOST_MAX ClockId = 3
	CLOCK_ID_COUNT              ClockId = 4
)

// DriverModel as declared in nvml/nvml.h
type DriverModel int32

// DriverModel enumeration from nvml/nvml.h
const (
	DRIVER_WDDM DriverModel = iota
	DRIVER_WDM  DriverModel = 1
)

// Pstates as declared in nvml/nvml.h
type Pstates int32

// Pstates enumeration from nvml/nvml.h
const (
	PSTATE_0       Pstates = iota
	PSTATE_1       Pstates = 1
	PSTATE_2       Pstates = 2
	PSTATE_3       Pstates = 3
	PSTATE_4       Pstates = 4
	PSTATE_5       Pstates = 5
	PSTATE_6       Pstates = 6
	PSTATE_7       Pstates = 7
	PSTATE_8       Pstates = 8
	PSTATE_9       Pstates = 9
	PSTATE_10      Pstates = 10
	PSTATE_11      Pstates = 11
	PSTATE_12      Pstates = 12
	PSTATE_13      Pstates = 13
	PSTATE_14      Pstates = 14
	PSTATE_15      Pstates = 15
	PSTATE_UNKNOWN Pstates = 32
)

// GpuOperationMode as declared in nvml/nvml.h
type GpuOperationMode int32

// GpuOperationMode enumeration from nvml/nvml.h
const (
	GOM_ALL_ON  GpuOperationMode = iota
	GOM_COMPUTE GpuOperationMode = 1
	GOM_LOW_DP  GpuOperationMode = 2
)

// InforomObject as declared in nvml/nvml.h
type InforomObject int32

// InforomObject enumeration from nvml/nvml.h
const (
	INFOROM_OEM   InforomObject = iota
	INFOROM_ECC   InforomObject = 1
	INFOROM_POWER InforomObject = 2
	INFOROM_COUNT InforomObject = 3
)

// Return as declared in nvml/nvml.h
type Return int32

// Return enumeration from nvml/nvml.h
const (
	SUCCESS                         Return = iota
	ERROR_UNINITIALIZED             Return = 1
	ERROR_INVALID_ARGUMENT          Return = 2
	ERROR_NOT_SUPPORTED             Return = 3
	ERROR_NO_PERMISSION             Return = 4
	ERROR_ALREADY_INITIALIZED       Return = 5
	ERROR_NOT_FOUND                 Return = 6
	ERROR_INSUFFICIENT_SIZE         Return = 7
	ERROR_INSUFFICIENT_POWER        Return = 8
	ERROR_DRIVER_NOT_LOADED         Return = 9
	ERROR_TIMEOUT                   Return = 10
	ERROR_IRQ_ISSUE                 Return = 11
	ERROR_LIBRARY_NOT_FOUND         Return = 12
	ERROR_FUNCTION_NOT_FOUND        Return = 13
	ERROR_CORRUPTED_INFOROM         Return = 14
	ERROR_GPU_IS_LOST               Return = 15
	ERROR_RESET_REQUIRED            Return = 16
	ERROR_OPERATING_SYSTEM          Return = 17
	ERROR_LIB_RM_VERSION_MISMATCH   Return = 18
	ERROR_IN_USE                    Return = 19
	ERROR_MEMORY                    Return = 20
	ERROR_NO_DATA                   Return = 21
	ERROR_VGPU_ECC_NOT_SUPPORTED    Return = 22
	ERROR_INSUFFICIENT_RESOURCES    Return = 23
	ERROR_FREQ_NOT_SUPPORTED        Return = 24
	ERROR_ARGUMENT_VERSION_MISMATCH Return = 25
	ERROR_DEPRECATED                Return = 26
	ERROR_UNKNOWN                   Return = 999
)

// MemoryLocation as declared in nvml/nvml.h
type MemoryLocation int32

// MemoryLocation enumeration from nvml/nvml.h
const (
	MEMORY_LOCATION_L1_CACHE       MemoryLocation = iota
	MEMORY_LOCATION_L2_CACHE       MemoryLocation = 1
	MEMORY_LOCATION_DRAM           MemoryLocation = 2
	MEMORY_LOCATION_DEVICE_MEMORY  MemoryLocation = 2
	MEMORY_LOCATION_REGISTER_FILE  MemoryLocation = 3
	MEMORY_LOCATION_TEXTURE_MEMORY MemoryLocation = 4
	MEMORY_LOCATION_TEXTURE_SHM    MemoryLocation = 5
	MEMORY_LOCATION_CBU            MemoryLocation = 6
	MEMORY_LOCATION_SRAM           MemoryLocation = 7
	MEMORY_LOCATION_COUNT          MemoryLocation = 8
)

// PageRetirementCause as declared in nvml/nvml.h
type PageRetirementCause int32

// PageRetirementCause enumeration from nvml/nvml.h
const (
	PAGE_RETIREMENT_CAUSE_MULTIPLE_SINGLE_BIT_ECC_ERRORS PageRetirementCause = iota
	PAGE_RETIREMENT_CAUSE_DOUBLE_BIT_ECC_ERROR           PageRetirementCause = 1
	PAGE_RETIREMENT_CAUSE_COUNT                          PageRetirementCause = 2
)

// RestrictedAPI as declared in nvml/nvml.h
type RestrictedAPI int32

// RestrictedAPI enumeration from nvml/nvml.h
const (
	RESTRICTED_API_SET_APPLICATION_CLOCKS  RestrictedAPI = iota
	RESTRICTED_API_SET_AUTO_BOOSTED_CLOCKS RestrictedAPI = 1
	RESTRICTED_API_COUNT                   RestrictedAPI = 2
)

// GpuVirtualizationMode as declared in nvml/nvml.h
type GpuVirtualizationMode int32

// GpuVirtualizationMode enumeration from nvml/nvml.h
const (
	GPU_VIRTUALIZATION_MODE_NONE        GpuVirtualizationMode = iota
	GPU_VIRTUALIZATION_MODE_PASSTHROUGH GpuVirtualizationMode = 1
	GPU_VIRTUALIZATION_MODE_VGPU        GpuVirtualizationMode = 2
	GPU_VIRTUALIZATION_MODE_HOST_VGPU   GpuVirtualizationMode = 3
	GPU_VIRTUALIZATION_MODE_HOST_VSGA   GpuVirtualizationMode = 4
)

// HostVgpuMode as declared in nvml/nvml.h
type HostVgpuMode int32

// HostVgpuMode enumeration from nvml/nvml.h
const (
	HOST_VGPU_MODE_NON_SRIOV HostVgpuMode = iota
	HOST_VGPU_MODE_SRIOV     HostVgpuMode = 1
)

// VgpuVmIdType as declared in nvml/nvml.h
type VgpuVmIdType int32

// VgpuVmIdType enumeration from nvml/nvml.h
const (
	VGPU_VM_ID_DOMAIN_ID VgpuVmIdType = iota
	VGPU_VM_ID_UUID      VgpuVmIdType = 1
)

// VgpuGuestInfoState as declared in nvml/nvml.h
type VgpuGuestInfoState int32

// VgpuGuestInfoState enumeration from nvml/nvml.h
const (
	VGPU_INSTANCE_GUEST_INFO_STATE_UNINITIALIZED VgpuGuestInfoState = iota
	VGPU_INSTANCE_GUEST_INFO_STATE_INITIALIZED   VgpuGuestInfoState = 1
)

// VgpuCapability as declared in nvml/nvml.h
type VgpuCapability int32

// VgpuCapability enumeration from nvml/nvml.h
const (
	VGPU_CAP_NVLINK_P2P           VgpuCapability = iota
	VGPU_CAP_GPUDIRECT            VgpuCapability = 1
	VGPU_CAP_MULTI_VGPU_EXCLUSIVE VgpuCapability = 2
	VGPU_CAP_EXCLUSIVE_TYPE       VgpuCapability = 3
	VGPU_CAP_EXCLUSIVE_SIZE       VgpuCapability = 4
	VGPU_CAP_COUNT                VgpuCapability = 5
)

// VgpuDriverCapability as declared in nvml/nvml.h
type VgpuDriverCapability int32

// VgpuDriverCapability enumeration from nvml/nvml.h
const (
	VGPU_DRIVER_CAP_HETEROGENEOUS_MULTI_VGPU VgpuDriverCapability = iota
	VGPU_DRIVER_CAP_COUNT                    VgpuDriverCapability = 1
)

// DeviceVgpuCapability as declared in nvml/nvml.h
type DeviceVgpuCapability int32

// DeviceVgpuCapability enumeration from nvml/nvml.h
const (
	DEVICE_VGPU_CAP_FRACTIONAL_MULTI_VGPU            DeviceVgpuCapability = iota
	DEVICE_VGPU_CAP_HETEROGENEOUS_TIMESLICE_PROFILES DeviceVgpuCapability = 1
	DEVICE_VGPU_CAP_HETEROGENEOUS_TIMESLICE_SIZES    DeviceVgpuCapability = 2
	DEVICE_VGPU_CAP_COUNT                            DeviceVgpuCapability = 3
)

// GpuUtilizationDomainId as declared in nvml/nvml.h
type GpuUtilizationDomainId int32

// GpuUtilizationDomainId enumeration from nvml/nvml.h
const (
	GPU_UTILIZATION_DOMAIN_GPU GpuUtilizationDomainId = iota
	GPU_UTILIZATION_DOMAIN_FB  GpuUtilizationDomainId = 1
	GPU_UTILIZATION_DOMAIN_VID GpuUtilizationDomainId = 2
	GPU_UTILIZATION_DOMAIN_BUS GpuUtilizationDomainId = 3
)

// FanState as declared in nvml/nvml.h
type FanState int32

// FanState enumeration from nvml/nvml.h
const (
	FAN_NORMAL FanState = iota
	FAN_FAILED FanState = 1
)

// LedColor as declared in nvml/nvml.h
type LedColor int32

// LedColor enumeration from nvml/nvml.h
const (
	LED_COLOR_GREEN LedColor = iota
	LED_COLOR_AMBER LedColor = 1
)

// EncoderType as declared in nvml/nvml.h
type EncoderType int32

// EncoderType enumeration from nvml/nvml.h
const (
	ENCODER_QUERY_H264 EncoderType = iota
	ENCODER_QUERY_HEVC EncoderType = 1
)

// FBCSessionType as declared in nvml/nvml.h
type FBCSessionType int32

// FBCSessionType enumeration from nvml/nvml.h
const (
	FBC_SESSION_TYPE_UNKNOWN FBCSessionType = iota
	FBC_SESSION_TYPE_TOSYS   FBCSessionType = 1
	FBC_SESSION_TYPE_CUDA    FBCSessionType = 2
	FBC_SESSION_TYPE_VID     FBCSessionType = 3
	FBC_SESSION_TYPE_HWENC   FBCSessionType = 4
)

// DetachGpuState as declared in nvml/nvml.h
type DetachGpuState int32

// DetachGpuState enumeration from nvml/nvml.h
const (
	DETACH_GPU_KEEP   DetachGpuState = iota
	DETACH_GPU_REMOVE DetachGpuState = 1
)

// PcieLinkState as declared in nvml/nvml.h
type PcieLinkState int32

// PcieLinkState enumeration from nvml/nvml.h
const (
	PCIE_LINK_KEEP      PcieLinkState = iota
	PCIE_LINK_SHUT_DOWN PcieLinkState = 1
)

// ClockLimitId as declared in nvml/nvml.h
type ClockLimitId int32

// ClockLimitId enumeration from nvml/nvml.h
const (
	CLOCK_LIMIT_ID_RANGE_START ClockLimitId = -256
	CLOCK_LIMIT_ID_TDP         ClockLimitId = -255
	CLOCK_LIMIT_ID_UNLIMITED   ClockLimitId = -254
)

// VgpuVmCompatibility as declared in nvml/nvml.h
type VgpuVmCompatibility int32

// VgpuVmCompatibility enumeration from nvml/nvml.h
const (
	VGPU_VM_COMPATIBILITY_NONE      VgpuVmCompatibility = iota
	VGPU_VM_COMPATIBILITY_COLD      VgpuVmCompatibility = 1
	VGPU_VM_COMPATIBILITY_HIBERNATE VgpuVmCompatibility = 2
	VGPU_VM_COMPATIBILITY_SLEEP     VgpuVmCompatibility = 4
	VGPU_VM_COMPATIBILITY_LIVE      VgpuVmCompatibility = 8
)

// VgpuPgpuCompatibilityLimitCode as declared in nvml/nvml.h
type VgpuPgpuCompatibilityLimitCode int32

// VgpuPgpuCompatibilityLimitCode enumeration from nvml/nvml.h
const (
	VGPU_COMPATIBILITY_LIMIT_NONE         VgpuPgpuCompatibilityLimitCode = iota
	VGPU_COMPATIBILITY_LIMIT_HOST_DRIVER  VgpuPgpuCompatibilityLimitCode = 1
	VGPU_COMPATIBILITY_LIMIT_GUEST_DRIVER VgpuPgpuCompatibilityLimitCode = 2
	VGPU_COMPATIBILITY_LIMIT_GPU          VgpuPgpuCompatibilityLimitCode = 4
	VGPU_COMPATIBILITY_LIMIT_OTHER        VgpuPgpuCompatibilityLimitCode = -2147483648
)

// ThermalTarget as declared in nvml/nvml.h
type ThermalTarget int32

// ThermalTarget enumeration from nvml/nvml.h
const (
	THERMAL_TARGET_NONE         ThermalTarget = iota
	THERMAL_TARGET_GPU          ThermalTarget = 1
	THERMAL_TARGET_MEMORY       ThermalTarget = 2
	THERMAL_TARGET_POWER_SUPPLY ThermalTarget = 4
	THERMAL_TARGET_BOARD        ThermalTarget = 8
	THERMAL_TARGET_VCD_BOARD    ThermalTarget = 9
	THERMAL_TARGET_VCD_INLET    ThermalTarget = 10
	THERMAL_TARGET_VCD_OUTLET   ThermalTarget = 11
	THERMAL_TARGET_ALL          ThermalTarget = 15
	THERMAL_TARGET_UNKNOWN      ThermalTarget = -1
)

// ThermalController as declared in nvml/nvml.h
type ThermalController int32

// ThermalController enumeration from nvml/nvml.h
const (
	THERMAL_CONTROLLER_NONE            ThermalController = iota
	THERMAL_CONTROLLER_GPU_INTERNAL    ThermalController = 1
	THERMAL_CONTROLLER_ADM1032         ThermalController = 2
	THERMAL_CONTROLLER_ADT7461         ThermalController = 3
	THERMAL_CONTROLLER_MAX6649         ThermalController = 4
	THERMAL_CONTROLLER_MAX1617         ThermalController = 5
	THERMAL_CONTROLLER_LM99            ThermalController = 6
	THERMAL_CONTROLLER_LM89            ThermalController = 7
	THERMAL_CONTROLLER_LM64            ThermalController = 8
	THERMAL_CONTROLLER_G781            ThermalController = 9
	THERMAL_CONTROLLER_ADT7473         ThermalController = 10
	THERMAL_CONTROLLER_SBMAX6649       ThermalController = 11
	THERMAL_CONTROLLER_VBIOSEVT        ThermalController = 12
	THERMAL_CONTROLLER_OS              ThermalController = 13
	THERMAL_CONTROLLER_NVSYSCON_CANOAS ThermalController = 14
	THERMAL_CONTROLLER_NVSYSCON_E551   ThermalController = 15
	THERMAL_CONTROLLER_MAX6649R        ThermalController = 16
	THERMAL_CONTROLLER_ADT7473S        ThermalController = 17
	THERMAL_CONTROLLER_UNKNOWN         ThermalController = -1
)

// GridLicenseFeatureCode as declared in nvml/nvml.h
type GridLicenseFeatureCode int32

// GridLicenseFeatureCode enumeration from nvml/nvml.h
const (
	GRID_LICENSE_FEATURE_CODE_UNKNOWN      GridLicenseFeatureCode = iota
	GRID_LICENSE_FEATURE_CODE_VGPU         GridLicenseFeatureCode = 1
	GRID_LICENSE_FEATURE_CODE_NVIDIA_RTX   GridLicenseFeatureCode = 2
	GRID_LICENSE_FEATURE_CODE_VWORKSTATION GridLicenseFeatureCode = 2
	GRID_LICENSE_FEATURE_CODE_GAMING       GridLicenseFeatureCode = 3
	GRID_LICENSE_FEATURE_CODE_COMPUTE      GridLicenseFeatureCode = 4
)

// GpmMetricId as declared in nvml/nvml.h
type GpmMetricId int32

// GpmMetricId enumeration from nvml/nvml.h
const (
	GPM_METRIC_GRAPHICS_UTIL           GpmMetricId = 1
	GPM_METRIC_SM_UTIL                 GpmMetricId = 2
	GPM_METRIC_SM_OCCUPANCY            GpmMetricId = 3
	GPM_METRIC_INTEGER_UTIL            GpmMetricId = 4
	GPM_METRIC_ANY_TENSOR_UTIL         GpmMetricId = 5
	GPM_METRIC_DFMA_TENSOR_UTIL        GpmMetricId = 6
	GPM_METRIC_HMMA_TENSOR_UTIL        GpmMetricId = 7
	GPM_METRIC_IMMA_TENSOR_UTIL        GpmMetricId = 9
	GPM_METRIC_DRAM_BW_UTIL            GpmMetricId = 10
	GPM_METRIC_FP64_UTIL               GpmMetricId = 11
	GPM_METRIC_FP32_UTIL               GpmMetricId = 12
	GPM_METRIC_FP16_UTIL               GpmMetricId = 13
	GPM_METRIC_PCIE_TX_PER_SEC         GpmMetricId = 20
	GPM_METRIC_PCIE_RX_PER_SEC         GpmMetricId = 21
	GPM_METRIC_NVDEC_0_UTIL            GpmMetricId = 30
	GPM_METRIC_NVDEC_1_UTIL            GpmMetricId = 31
	GPM_METRIC_NVDEC_2_UTIL            GpmMetricId = 32
	GPM_METRIC_NVDEC_3_UTIL            GpmMetricId = 33
	GPM_METRIC_NVDEC_4_UTIL            GpmMetricId = 34
	GPM_METRIC_NVDEC_5_UTIL            GpmMetricId = 35
	GPM_METRIC_NVDEC_6_UTIL            GpmMetricId = 36
	GPM_METRIC_NVDEC_7_UTIL            GpmMetricId = 37
	GPM_METRIC_NVJPG_0_UTIL            GpmMetricId = 40
	GPM_METRIC_NVJPG_1_UTIL            GpmMetricId = 41
	GPM_METRIC_NVJPG_2_UTIL            GpmMetricId = 42
	GPM_METRIC_NVJPG_3_UTIL            GpmMetricId = 43
	GPM_METRIC_NVJPG_4_UTIL            GpmMetricId = 44
	GPM_METRIC_NVJPG_5_UTIL            GpmMetricId = 45
	GPM_METRIC_NVJPG_6_UTIL            GpmMetricId = 46
	GPM_METRIC_NVJPG_7_UTIL            GpmMetricId = 47
	GPM_METRIC_NVOFA_0_UTIL            GpmMetricId = 50
	GPM_METRIC_NVLINK_TOTAL_RX_PER_SEC GpmMetricId = 60
	GPM_METRIC_NVLINK_TOTAL_TX_PER_SEC GpmMetricId = 61
	GPM_METRIC_NVLINK_L0_RX_PER_SEC    GpmMetricId = 62
	GPM_METRIC_NVLINK_L0_TX_PER_SEC    GpmMetricId = 63
	GPM_METRIC_NVLINK_L1_RX_PER_SEC    GpmMetricId = 64
	GPM_METRIC_NVLINK_L1_TX_PER_SEC    GpmMetricId = 65
	GPM_METRIC_NVLINK_L2_RX_PER_SEC    GpmMetricId = 66
	GPM_METRIC_NVLINK_L2_TX_PER_SEC    GpmMetricId = 67
	GPM_METRIC_NVLINK_L3_RX_PER_SEC    GpmMetricId = 68
	GPM_METRIC_NVLINK_L3_TX_PER_SEC    GpmMetricId = 69
	GPM_METRIC_NVLINK_L4_RX_PER_SEC    GpmMetricId = 70
	GPM_METRIC_NVLINK_L4_TX_PER_SEC    GpmMetricId = 71
	GPM_METRIC_NVLINK_L5_RX_PER_SEC    GpmMetricId = 72
	GPM_METRIC_NVLINK_L5_TX_PER_SEC    GpmMetricId = 73
	GPM_METRIC_NVLINK_L6_RX_PER_SEC    GpmMetricId = 74
	GPM_METRIC_NVLINK_L6_TX_PER_SEC    GpmMetricId = 75
	GPM_METRIC_NVLINK_L7_RX_PER_SEC    GpmMetricId = 76
	GPM_METRIC_NVLINK_L7_TX_PER_SEC    GpmMetricId = 77
	GPM_METRIC_NVLINK_L8_RX_PER_SEC    GpmMetricId = 78
	GPM_METRIC_NVLINK_L8_TX_PER_SEC    GpmMetricId = 79
	GPM_METRIC_NVLINK_L9_RX_PER_SEC    GpmMetricId = 80
	GPM_METRIC_NVLINK_L9_TX_PER_SEC    GpmMetricId = 81
	GPM_METRIC_NVLINK_L10_RX_PER_SEC   GpmMetricId = 82
	GPM_METRIC_NVLINK_L10_TX_PER_SEC   GpmMetricId = 83
	GPM_METRIC_NVLINK_L11_RX_PER_SEC   GpmMetricId = 84
	GPM_METRIC_NVLINK_L11_TX_PER_SEC   GpmMetricId = 85
	GPM_METRIC_NVLINK_L12_RX_PER_SEC   GpmMetricId = 86
	GPM_METRIC_NVLINK_L12_TX_PER_SEC   GpmMetricId = 87
	GPM_METRIC_NVLINK_L13_RX_PER_SEC   GpmMetricId = 88
	GPM_METRIC_NVLINK_L13_TX_PER_SEC   GpmMetricId = 89
	GPM_METRIC_NVLINK_L14_RX_PER_SEC   GpmMetricId = 90
	GPM_METRIC_NVLINK_L14_TX_PER_SEC   GpmMetricId = 91
	GPM_METRIC_NVLINK_L15_RX_PER_SEC   GpmMetricId = 92
	GPM_METRIC_NVLINK_L15_TX_PER_SEC   GpmMetricId = 93
	GPM_METRIC_NVLINK_L16_RX_PER_SEC   GpmMetricId = 94
	GPM_METRIC_NVLINK_L16_TX_PER_SEC   GpmMetricId = 95
	GPM_METRIC_NVLINK_L17_RX_PER_SEC   GpmMetricId = 96
	GPM_METRIC_NVLINK_L17_TX_PER_SEC   GpmMetricId = 97
	GPM_METRIC_MAX                     GpmMetricId = 98
)
