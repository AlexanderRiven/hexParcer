package main

type Block1 struct {
	FunctionCode     uint8    //Код функции
	BlockLength      uint8    //Длина блока
	BlockID          [4]byte  //Идентификатор блока данных
	FormatVersion    uint8    //Номер версии формата
	BlockSize        uint8    //Размер блока
	DeviceTypeID     uint8    //Идентификатор типа устройства в целом
	ModuleTypeID     uint32   //Идентификатор типа модуля устройства
	MainVerNumN      uint8    //Номер главной версии МНЗЧ ПО
	AdditVerNumN     uint8    //Номер дополнительной версии МНЗЧ ПО
	MainVerNum       uint8    //Номер главной версии МЗЧ ПО
	AdditVerNum      uint8    //Номер дополнительной версии МЗЧ ПО
	CRCPO            uint32   //CRC МЗЧ ПО
	AppPartVer       [20]byte // Версия аппаратной части
	DeviceFactoryID  [20]byte //Заводской номер устройства
	CurrentYear      [2]byte  //Текущий год по прибору
	CurrentMonth     uint8    //Текущий месяц по прибору
	CurrentDay       uint8    //Текущий день по прибору
	CurrentHour      uint8    //Текущий час по прибору
	CurrentMinute    uint8    //Текущая минута по прибору
	CurrentSecond    uint8    //Текущая секунда по прибору
	Timezone         uint8    //Часовой пояс
	InterfaceChannel uint8    //Канал интерфейса связи
	PowerMode        uint8    //Режим работы: включен
	AdditBlockId     uint32   // Идентификатор блока дополнительных данных
	VersionFormatNum uint8    //Номер версии формата
	BlockSize1       uint8    //Размер блока
	UniqNumber       [8]byte  //Уникальный серийный номер МК
}

type Block2 struct {
	FunctionCode     uint8
	BlockLength      uint8
	ForF1            [2]byte //Для завода-изготовителя
	StandardV        [4]byte //Накопленный объем, приведенный к стандартным условиям
	WorkV            [4]byte //Рабочий накопленный объем
	RashodStandard   float32 //Мгновенный расход по стандартным условиям
	RashodWork       float32 //Мгновенный рабочий расход
	Temperature      [4]byte //Температура газа
	GazP             [4]byte //Абсолютное давление газа
	ConnFlags        [4]byte //Флаги причин выхода на связь
	WarningFlags     [4]byte //Флаги предупреждений
	NSFlags          [4]byte //Флаги нештатных ситуаций
	WorkTimer        [2]byte //Счетчик времени работы модема
	SessTimer        [2]byte //Счетчик успешных сессий
	ForF2            [2]byte //Для завода-изготовителя
	ForF3            [2]byte //Для завода-изготовителя
	CalibrationDate  [4]byte //Дата калибровки
	CoffInto         float32 //К-т сжимаемости
	CoffTransfer     float32 //К-т перевода
	ConnCounter      [2]byte //Количество подключений
	SuccsConnCounter [2]byte //Количество успешных подключений
	OpenKl           [4]byte //Количество открытий внешнего клапана
	CloseKl          [4]byte //Количество закрытий внешнего клапана
	BatteryP         [2]byte //Напряжение батареи под нагрузкой (мВ)
	ForF4            [4]byte //Для завода-изготовителя
	ForF5            [4]byte //Для завода-изготовителя
	MCC              [2]byte
	MNC              [2]byte
	LAC              [2]byte
	CI               [2]byte
	Charge           [2]byte //Текущий заряд батареи (в сотых процентов)
}

type Block3 struct {
	FunctionCode          uint8
	BlockLength           uint8
	ForF1                 [2]byte //Для завода-изготовителя
	DataConfirm           [2]byte //Регистр подтверждения приема данных сервером
	ConnClose             [2]byte //Закрытие соединения
	ArchiveFlags          [2]byte //Флаги архива
	ForF2                 [2]byte //Для завода-изготовителя
	DaysContr             [2]byte //Контрактные сутки
	DeviceDT              [8]byte //Дата/время прибора
	HashPass              [4]byte //Текущий HASH пароля
	ForF3                 [2]byte //Для завода-изготовителя
	ForF4                 [4]byte //Для завода-изготовителя
	ForF5                 [4]byte //Для завода-изготовителя
	AccessLvl             [2]byte //Текущий уровень доступа
	ContractHour          [2]byte //Текущий уровень доступа
	AzotPerc              float32 //Состав газа: содержание азота
	CO2Perc               float32 //Состав газа: содержание СО2
	PlotPerc              float32 //Состав газа: плотность
	CoffIntoConst         float32 //Коэффициент сжимаемости: константа
	CoffInto              [2]byte //Коэффициент сжимаемости:
	ServerIP              [4]byte //IP–адрес сервера
	ServerPort            [2]byte //Порт сервера
	ServTimeout           [2]byte //Таймаут соединения с сервером, сек
	ServTimeoutAfterFirst [2]byte //Таймаут после первой попытки подключения, сек
	ServTimeoutBetween    [2]byte //Таймаут между попытками подключения (последующие) , мин
	RetryConn             [4]byte //Повтор выхода на связь
	ForF6                 [2]byte //Для завода-изготовителя
	MaxRash               float32 //Максимальный расход
	ForF7                 [2]byte //Для завода-изготовителя
	ForF8                 [2]byte //Для завода-изготовителя
	ForF9                 [2]byte //Для завода-изготовителя
	ClapManual            [2]byte //Управление внешним клапаном
	ForF10                [2]byte //Для завода-изготовителя
	ForF11                float32 //Для завода-изготовителя
	ForF12                float32 //Для завода-изготовителя
	TempMin               float32 //Минимальная температура
	TempMax               float32 //Максимальная температура
	TimeCont              float32 //Договорная температура
	MaxP                  float32 //Максимальное давление
	ContP                 float32 //Договорное давление
	MinP                  float32 //Минимальное давление
}

type Block4 struct {
	FunctionCode     uint8
	BlockLength      uint8
	ForF1            [2]byte //Для завода-изготовителя
	ForF2            float32 //Для завода-изготовителя
	ForF3            [2]byte //Для завода-изготовителя
	ForF4            [2]byte //Для завода-изготовителя
	ForF5            [2]byte //Для завода-изготовителя
	ForF6            [2]byte //Для завода-изготовителя
	ForF7            [2]byte //Для завода-изготовителя
	ForF8            [2]byte //Для завода-изготовителя
	ServerReservIP   [4]byte //IP–адрес резервного сервера
	ServerReservPort [2]byte //Порт резервного сервера
	TimeoutUpdate    [2]byte //Таймаут соединения с сервером обновлений
	TimeoutFirstTry  [2]byte //Таймаут после первой попытки подключения к серверу обновлений
	TimeoutRetry     [2]byte //Таймаут между попытками подключения к серверу обновлений
	ConnRetry        [8]byte //Повтор выхода на связь с сервером обновлений
	ServerUpdateIp   [4]byte //IP–адрес сервера обновлений
	ServerUpdatePort [2]byte //Порт сервера обновлений
	RetryCount       [2]byte //Кол-во попыток повтора дозвона
	RetryCountUpdate [2]byte //Кол-во попыток повтора дозвона на сервер обновления
}

type Arch struct {
	FuncCode        uint8
	StartADDR       [2]byte
	AmountOfBytes   [2]byte //Общая длина архивных записей
	DATE_TIME       [4]byte //Дата и время записи
	VPerHour        [4]byte //Объем газа за час в рабочих условиях, литры
	VPerHourStandar [4]byte //Объем газа за час, приведенный к стандартным условиям, литры
	AvgTemp         [4]byte //Среднечасовое значение температуры
	AvgPascal       [4]byte //Среднечасовое значение абсолютного давления
	NSPriznak       [2]byte //Признак возникновения НС
	Srez            [4]byte //Срез значения стандартного объема газа нарастающим итогом на конец суток
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
	NSID          [2]byte //Идентификатор события
	Para1         [4]byte
	Para2         [4]byte
	Para3         [4]byte
}
