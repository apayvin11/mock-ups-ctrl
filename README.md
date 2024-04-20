# MOCK-UPS-CTRL
## Description
mock ups controller, modbus server
# Modbus register map
float32 big endian byte order
## InputRegisters:   

| Category            | Register | Parameter                     | Type     | Unit |
| ------------------- | -------- | ----------------------------- | -------- | ----------- |
| **Battery group 1** | 0x0000   | battery group 1 voltage       | float32  | Volt
|                     | 0x0002   | battery group 1 current       | float32  | Ampere
|                     | 0x0010   | battery 1 voltage             | float32  | Volt
|                     | 0x0012   | battery 1 temp                | float32  | °C
|                     | 0x0014   | battery 1 internal resistance | float32  | mΩ
|                     | 0x0020   | battery 2 voltage             | float32  | Volt
|                     | 0x0022   | battery 2 temp                | float32  | °C
|                     | 0x0024   | battery 2 internal resistance | float32  | mΩ
|                     | 0x0030   | battery 3 voltage             | float32  | Volt
|                     | 0x0032   | battery 3 temp                | float32  | °C
|                     | 0x0034   | battery 3 internal resistance | float32  | mΩ
||...
|                     | 0x00n0   | battery n voltage             | float32  | Volt
|                     | 0x00n2   | battery n temp                | float32  | °C
|                     | 0x00n4   | battery n internal resistance | float32  | mΩ
| **Battery group 2** | 0x0100   | battery group 2 voltage       | float32  | Volt
|                     | 0x0102   | battery group 2 current       | float32  | Ampere
|                     | 0x0110   | battery 1 voltage             | float32  | Volt
|                     | 0x0112   | battery 1 temp                | float32  | °C
|                     | 0x0114   | battery 1 internal resistance | float32  | mΩ
|                     | 0x0120   | battery 2 voltage             | float32  | Volt
|                     | 0x0122   | battery 2 temp                | float32  | °C
|                     | 0x0124   | battery 2 internal resistance | float32  | mΩ
|                     | 0x0130   | battery 3 voltage             | float32  | Volt
|                     | 0x0132   | battery 3 temp                | float32  | °C
|                     | 0x0134   | battery 3 internal resistance | float32  | mΩ
||...
|                     | 0x01n0   | battery n voltage             | float32  | Volt
|                     | 0x01n2   | battery n temp                | float32  | °C
|                     | 0x01n4   | battery n internal resistance | float32  | mΩ