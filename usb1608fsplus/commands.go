package usb1608fsplus

const timeout = 2000

type command byte

// Log level enumeration
const (
	// Digital I/O commands
	commandDigitalTristate command = 0x00
	commandDigitalPort     command = 0x01
	commandDigitalLatch    command = 0x02
	// Analog input commands
	commandAnalogInput       command = 0x10
	commandAnalogStartScan   command = 0x11
	commandAnalogStopScan    command = 0x12
	commandAnalogConfig      command = 0x14
	commandAnalogClearBuffer command = 0x15
	// Counter/timer commands
	commandEventCounter command = 0x20
	// Memory commands
	commandCalibrationMemory command = 0x30
	commandUserMemory        command = 0x31
	commandMBDMemory         command = 0x32
	// Miscellaneous commands
	commandBlinkLED        command = 0x41
	commandReset           command = 0x42
	commandGetStatus       command = 0x44
	commandSerialNum       command = 0x48
	commandUpgradeFirmware command = 0x50
	// Message-Based DAQ (MBD) Protocal commands
	commandTextMBD command = 0x80
	commandRawMBD  command = 0x81
)

var commands = map[command]string{
	commandDigitalTristate:   "Read/write tri-state register",
	commandDigitalPort:       "Read digital port pins",
	commandDigitalLatch:      "Read/write digital port output latch register",
	commandAnalogInput:       "Read analog input channel",
	commandAnalogStartScan:   "Start analog input scan",
	commandAnalogStopScan:    "Stop analog input scan",
	commandAnalogConfig:      "Configure the analog input channel",
	commandAnalogClearBuffer: "Clear the analog input scan FIFO buffer",
	commandEventCounter:      "Read/reset event counter",
	commandCalibrationMemory: "Read/write calibration memory",
	commandUserMemory:        "Read/write user memory",
	commandMBDMemory:         "Read/write Message-Based DAQ (MBD) memory",
	commandBlinkLED:          "Blink LED",
	commandReset:             "Reset device",
	commandGetStatus:         "Read device status",
	commandSerialNum:         "Read/write serial number",
	commandUpgradeFirmware:   "Enter device firmware upgrade (DFU) mode",
	commandTextMBD:           "Text-based MBD command/response",
	commandRawMBD:            "Raw MBD response",
}

func (c command) String() string {
	return commands[c]
}

type scanOption byte

// Analog input scan options
const (
	scanBlockTransferMode     scanOption = 0x0
	scanImmediateTransferMode scanOption = 0x1
	scanInternalPacerOff      scanOption = 0x0
	scanInternalPcerOn        scanOption = 0x2
	scanNoTrigger             scanOption = 0x0
	scanTriggerRisingEdge     scanOption = 0x1 << 2
	scanTriggerFallingEdge    scanOption = 0x2 << 2
	scanTriggerHighLevel      scanOption = 0x3 << 2
	scanTriggerLowLevel       scanOption = 0x4 << 2
	scanDebugMode             scanOption = 0x10
	scanStallOnOverrun        scanOption = 0x0
	scanInhibitStall                     = 0x1 << 7
)

type analogInput byte

// Analog input setup
const (
	singleEnded  analogInput = 0
	differential analogInput = 1
	calibration  analogInput = 3
)

const (
	lastChannel               byte = 0x80
	maxBulkTransferPacketSize byte = 64
)

type voltageRange byte

// Ranges
const (
	range10V    voltageRange = 0x0 // ±10V
	range5V     voltageRange = 0x1 // ±5V
	range2500mV voltageRange = 0x2 // ±2.5V
	range2000mV voltageRange = 0x3 // ±2V
	range1250mV voltageRange = 0x4 // ±1.25V
	range1000mV voltageRange = 0x5 // ±1V
	range625mV  voltageRange = 0x6 // ±0.625V
	range312mV  voltageRange = 0x7 // ±0.3125V
)

type statusBit byte

// Status bit values
const (
	scanRunning statusBit = 0x1 << 1
	scanOverrun statusBit = 0x1 << 2
)

const (
	maxNumADChannels = 8  // max number of A/D channels in device
	maxNumGainLevels = 8  // max number of gain levels in device
	maxPacketSize    = 64 // max packet size for FS device
)
