package main

type Block1 struct {
	FunctionCode     uint8
	BlockLength      uint8
	BlockID          [4]byte
	FormatVersion    uint8
	BlockSize        uint8
	DeviceTypeID     uint8
	ModuleTypeID     uint32
	MainVerNumN      uint8
	AdditVerNumN     uint8
	MainVerNum       uint8
	AdditVerNum      uint8
	CRCPO            uint32
	AppPartVer       [20]byte
	DeviceFactoryID  [20]byte
	CurrentYear      [2]byte
	CurrentMonth     uint8
	CurrentDay       uint8
	CurrentHour      uint8
	CurrentMinute    uint8
	CurrentSecond    uint8
	Timezone         uint8
	InterfaceChannel uint8
	PowerMode        uint8
	AdditBlockId     uint32
	VersionFormatNum uint8
	BlockSize1       uint8
	UniqNumber       [8]byte
}

type Block2 struct {
	FunctionCode     uint8
	BlockLength      uint8
	ForF1            [2]byte
	StandardV        [4]byte
	WorkV            [4]byte
	RashodStandard   float32
	RashodWork       float32
	Temperature      [4]byte
	GazP             [4]byte
	ConnFlags        [4]byte
	WarningFlags     [4]byte
	NSFlags          [4]byte
	WorkTimer        [2]byte
	SessTimer        [2]byte
	ForF2            [2]byte
	ForF3            [2]byte
	CalibrationDate  [4]byte
	CoffInto         float32
	CoffTransfer     float32
	ConnCounter      [2]byte
	SuccsConnCounter [2]byte
	OpenKl           [4]byte
	CloseKl          [4]byte
	BatteryP         [2]byte
	ForF4            [4]byte
	ForF5            [4]byte
	MCC              [2]byte
	MNC              [2]byte
	LAC              [2]byte
	CI               [2]byte
	Charge           [2]byte
}

type Block3 struct {
	FunctionCode          uint8
	BlockLength           uint8
	ForF1                 [2]byte
	DataConfirm           [2]byte
	ConnClose             [2]byte
	ArchiveFlags          [2]byte
	ForF2                 [2]byte
	DaysContr             [2]byte
	DeviceDT              [8]byte
	HashPass              [4]byte
	ForF3                 [2]byte
	ForF4                 [4]byte
	ForF5                 [4]byte
	AccessLvl             [2]byte
	ContractHour          [2]byte
	AzotPerc              float32
	CO2Perc               float32
	PlotPerc              float32
	CoffIntoConst         float32
	CoffInto              [2]byte
	ServerIP              [4]byte
	ServerPort            [2]byte
	ServTimeout           [2]byte
	ServTimeoutAfterFirst [2]byte
	ServTimeoutBetween    [2]byte
	RetryConn             [4]byte
	ForF6                 [2]byte
	MaxRash               float32
	ForF7                 [2]byte
	ForF8                 [2]byte
	ForF9                 [2]byte
	ClapManual            [2]byte
	ForF10                [2]byte
	ForF11                float32
	ForF12                float32
	TempMin               float32
	TempMax               float32
	TimeCont              float32
	MaxP                  float32
	ContP                 float32
	MinP                  float32
}

type Block4 struct {
	FunctionCode     uint8
	BlockLength      uint8
	ForF1            [2]byte
	ForF2            float32
	ForF3            [2]byte
	ForF4            [2]byte
	ForF5            [2]byte
	ForF6            [2]byte
	ForF7            [2]byte
	ForF8            [2]byte
	ServerReservIP   [4]byte
	ServerReservPort [2]byte
	TimeoutUpdate    [2]byte
	TimeoutFirstTry  [2]byte
	TimeoutRetry     [2]byte
	ConnRetry        [8]byte
	ServerUpdateIp   [4]byte
	ServerUpdatePort [2]byte
	RetryCount       [2]byte
	RetryCountUpdate [2]byte
}

type Arch struct {
	FuncCode        uint8
	StartADDR       [2]byte
	AmountOfBytes   [2]byte
	DATE_TIME       [4]byte
	VPerHour        [4]byte
	VPerHourStandar [4]byte
	AvgTemp         [4]byte
	AvgPascal       [4]byte
	NSPriznak       [2]byte
	Srez            [4]byte
}

type SubArch struct {
	DATE_TIME       [4]byte
	VPerHour        [4]byte
	VPerHourStandar [4]byte
	AvgTemp         [4]byte
	AvgPascal       [4]byte
	NSPriznak       [2]byte
	Srez            [4]byte
}

type ArchNs struct {
	FuncCode      uint8
	StartADDR     [2]byte
	AmountOfBytes [2]byte
	DATE_TIME     [4]byte
	NSID          [2]byte
	Para1         [4]byte
	Para2         [4]byte
	Para3         [4]byte
}
