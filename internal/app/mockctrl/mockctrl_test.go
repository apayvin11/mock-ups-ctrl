package mockctrl

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
	"testing"
	"time"

	"github.com/goburrow/modbus"
	"github.com/stretchr/testify/assert"
	"github.com/tbrandon/mbserver"
)

const (
	slaveId        = 1
	tcpAddressPort = ":1502"
)

func TestMockCtrl(t *testing.T) {
	// Run Modbus server
	serv := mbserver.NewServer()
	InitMockModbusParams(serv)
	assert.NoError(t, serv.ListenTCP(tcpAddressPort))
	defer serv.Close()

	// Modbus TCP Client
	handler := modbus.NewTCPClientHandler(tcpAddressPort)
	handler.Timeout = time.Second
	handler.SlaveId = slaveId
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	assert.NoError(t, handler.Connect())
	defer handler.Close()

	client := modbus.NewClient(handler)
	testCases := []struct {
		name          string
		reg           uint16
		expectedValue float32
	}{
		{
			name:          "group 1 voltage",
			reg:           0x0000,
			expectedValue: 48,
		},
		{
			name:          "group 1 current",
			reg:           0x0002,
			expectedValue: 10,
		},
		{
			name:          "group 1 battery 1 voltage",
			reg:           0x0010,
			expectedValue: 12,
		},
		{
			name:          "group 1 battery 1 temp",
			reg:           0x0012,
			expectedValue: 25,
		},
		{
			name:          "group 1 battery 1 resist",
			reg:           0x0014,
			expectedValue: 5,
		},
		{
			name:          "group 1 battery 2 voltage",
			reg:           0x0020,
			expectedValue: 13,
		},
		{
			name:          "group 2 battery 1 voltage",
			reg:           0x0110,
			expectedValue: 14,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input_registers_bytes, err := client.ReadInputRegisters(tc.reg, 2)
			assert.NoError(t, err)
			val, err := float32FromBytes(input_registers_bytes)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedValue, val)
		})
	}
}

func float32FromBytes(b []byte) (float32, error) {
	if len(b) != 4 {
		return 0, fmt.Errorf("invalid len([]byte): %d, expected 4", len(b))
	}
	return math.Float32frombits(binary.BigEndian.Uint32(b)), nil
}
