# MOCK-UPS-CTRL
## Description
mock ups controller, modbus server
modbus over tcp only, default port 1502 (see config)
# Modbus register map
float32 big endian byte order
## InputRegisters:   

| Register | Parameter                     | Type     | Unit   
| -------- | ----------------------------- | -------- | ------ 
| 0x0000   | input AC voltage              | float32  | Volt
| 0x0002   | input AC current              | float32  | Ampere
| 0x0004   | battery group voltage         | float32  | Volt
| 0x0006   | battery group current         | float32  | Ampere
| 0x0010   | battery 1 voltage             | float32  | Volt
| 0x0012   | battery 1 temp                | float32  | °C
| 0x0014   | battery 1 internal resistance | float32  | mΩ
| 0x0020   | battery 2 voltage             | float32  | Volt
| 0x0022   | battery 2 temp                | float32  | °C
| 0x0024   | battery 2 internal resistance | float32  | mΩ
| 0x0030   | battery 3 voltage             | float32  | Volt
| 0x0032   | battery 3 temp                | float32  | °C
| 0x0034   | battery 3 internal resistance | float32  | mΩ
| 0x0040   | battery 4 voltage             | float32  | Volt
| 0x0042   | battery 4 temp                | float32  | °C
| 0x0044   | battery 4 internal resistance | float32  | mΩ

## Discrete Inputs
### Alarms
| Register | Parameter            
| -------- | -------------------- 
| 0x0000   | UPS in Battery Mode  
| 0x0001   | Low battery          
| 0x0002   | Overload             
