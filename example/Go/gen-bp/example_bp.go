// Code generated by bitproto. DO NOT EDIT.

// Proto drone describes the structure of the drone.
package drone

import (
	"strconv"
	"encoding/json"

	bp "github.com/hit9/bitproto/lib/go"
)

type Timestamp int64 // 64bit

func (m Timestamp) BpProcessor() bp.Processor {
	return bp.NewAliasProcessor(bp.NewInt(64))
}

type TernaryInt32 [3]int32 // 96bit

func (m TernaryInt32) BpProcessor() bp.Processor {
	return bp.NewAliasProcessor(bp.NewArray(false, 3, bp.NewInt(32)))
}

type DroneStatus uint8 // 3bit

const (
	DRONE_STATUS_UNKNOWN DroneStatus = 0
	DRONE_STATUS_STANDBY = 1
	DRONE_STATUS_RISING = 2
	DRONE_STATUS_LANDING = 3
	DRONE_STATUS_FLYING = 4
)

func (m DroneStatus) BpProcessor() bp.Processor {
	return bp.NewEnumProcessor(bp.NewUint(3))
}

// String returns the name of this enum item.
func (v DroneStatus) String() string {
	switch v {
	case 0:
		return "DRONE_STATUS_UNKNOWN"
	case 1:
		return "DRONE_STATUS_STANDBY"
	case 2:
		return "DRONE_STATUS_RISING"
	case 3:
		return "DRONE_STATUS_LANDING"
	case 4:
		return "DRONE_STATUS_FLYING"
	default:
		return "DroneStatus(" + strconv.FormatInt(int64(v), 10) + ")"
	}
}

type PropellerStatus uint8 // 2bit

const (
	PROPELLER_STATUS_UNKNOWN PropellerStatus = 0
	PROPELLER_STATUS_IDLE = 1
	PROPELLER_STATUS_ROTATING = 2
)

func (m PropellerStatus) BpProcessor() bp.Processor {
	return bp.NewEnumProcessor(bp.NewUint(2))
}

// String returns the name of this enum item.
func (v PropellerStatus) String() string {
	switch v {
	case 0:
		return "PROPELLER_STATUS_UNKNOWN"
	case 1:
		return "PROPELLER_STATUS_IDLE"
	case 2:
		return "PROPELLER_STATUS_ROTATING"
	default:
		return "PropellerStatus(" + strconv.FormatInt(int64(v), 10) + ")"
	}
}

type RotatingDirection uint8 // 2bit

const (
	ROTATING_DIRECTION_UNKNOWN RotatingDirection = 0
	ROTATING_DIRECTION_CLOCK_WISE = 1
	ROTATING_DIRECTION_ANTI_CLOCK_WISE = 2
)

func (m RotatingDirection) BpProcessor() bp.Processor {
	return bp.NewEnumProcessor(bp.NewUint(2))
}

// String returns the name of this enum item.
func (v RotatingDirection) String() string {
	switch v {
	case 0:
		return "ROTATING_DIRECTION_UNKNOWN"
	case 1:
		return "ROTATING_DIRECTION_CLOCK_WISE"
	case 2:
		return "ROTATING_DIRECTION_ANTI_CLOCK_WISE"
	default:
		return "RotatingDirection(" + strconv.FormatInt(int64(v), 10) + ")"
	}
}

type PowerStatus uint8 // 2bit

const (
	POWER_STATUS_UNKNOWN PowerStatus = 0
	POWER_STATUS_OFF = 1
	POWER_STATUS_ON = 2
)

func (m PowerStatus) BpProcessor() bp.Processor {
	return bp.NewEnumProcessor(bp.NewUint(2))
}

// String returns the name of this enum item.
func (v PowerStatus) String() string {
	switch v {
	case 0:
		return "POWER_STATUS_UNKNOWN"
	case 1:
		return "POWER_STATUS_OFF"
	case 2:
		return "POWER_STATUS_ON"
	default:
		return "PowerStatus(" + strconv.FormatInt(int64(v), 10) + ")"
	}
}

type LandingGearStatus uint8 // 2bit

const (
	LANDING_GEAR_STATUS_UNKNOWN LandingGearStatus = 0
	LANDING_GEAR_STATUS_UNFOLDED = 1
	LANDING_GEAR_STATUS_FOLDED = 2
)

func (m LandingGearStatus) BpProcessor() bp.Processor {
	return bp.NewEnumProcessor(bp.NewUint(2))
}

// String returns the name of this enum item.
func (v LandingGearStatus) String() string {
	switch v {
	case 0:
		return "LANDING_GEAR_STATUS_UNKNOWN"
	case 1:
		return "LANDING_GEAR_STATUS_UNFOLDED"
	case 2:
		return "LANDING_GEAR_STATUS_FOLDED"
	default:
		return "LandingGearStatus(" + strconv.FormatInt(int64(v), 10) + ")"
	}
}

type Propeller struct {
	Id uint8 `json:"id"` // 8bit
	Status PropellerStatus `json:"status"` // 2bit
	Direction RotatingDirection `json:"direction"` // 2bit
}

// Number of bytes to serialize struct Propeller
const BYTES_LENGTH_PROPELLER uint32 = 2

func (m *Propeller) Size() uint32 { return 2 }

// Returns string representation for struct Propeller.
func (m *Propeller) String() string {
	v, _ := json.Marshal(m)
	return string(v)
}

// Encode struct Propeller to bytes buffer.
func (m *Propeller) Encode() []byte {
	ctx := bp.NewEncodeContext(int(m.Size()))
	m.BpProcessor().Process(ctx, nil, m)
	return ctx.Buffer()
}

func (m *Propeller) Decode(s []byte) {
	ctx := bp.NewDecodeContext(s)
	m.BpProcessor().Process(ctx, nil, m)
}

func (m *Propeller) BpProcessor() bp.Processor {
	fieldDescriptors := []*bp.MessageFieldProcessor{
		bp.NewMessageFieldProcessor(1, bp.NewUint(8)),
		bp.NewMessageFieldProcessor(2, (PropellerStatus(0)).BpProcessor()),
		bp.NewMessageFieldProcessor(3, (RotatingDirection(0)).BpProcessor()),
	}
	return bp.NewMessageProcessor(false, 12, fieldDescriptors)
}

func (m *Propeller) BpGetAccessor(di *bp.DataIndexer) bp.Accessor {
	switch di.F() {
	default:
		return nil  // Won't reached
	}
}

func (m *Propeller) BpSetByte(di *bp.DataIndexer, lshift int, b byte) {
	switch di.F() {
		case 1:
			m.Id |= (uint8(b) << lshift)
		case 2:
			m.Status |= (PropellerStatus(b) << lshift)
		case 3:
			m.Direction |= (RotatingDirection(b) << lshift)
		default:
			return
	}
}

func (m *Propeller) BpGetByte(di *bp.DataIndexer, rshift int) byte {
	switch di.F() {
		case 1:
			return byte(m.Id >> rshift)
		case 2:
			return byte(m.Status >> rshift)
		case 3:
			return byte(m.Direction >> rshift)
		default:
			return byte(0) // Won't reached
	}
}

type Power struct {
	Battery uint8 `json:"battery"` // 8bit
	Status PowerStatus `json:"status"` // 2bit
	IsCharging bool `json:"is_charging"` // 1bit
}

// Number of bytes to serialize struct Power
const BYTES_LENGTH_POWER uint32 = 2

func (m *Power) Size() uint32 { return 2 }

// Returns string representation for struct Power.
func (m *Power) String() string {
	v, _ := json.Marshal(m)
	return string(v)
}

// Encode struct Power to bytes buffer.
func (m *Power) Encode() []byte {
	ctx := bp.NewEncodeContext(int(m.Size()))
	m.BpProcessor().Process(ctx, nil, m)
	return ctx.Buffer()
}

func (m *Power) Decode(s []byte) {
	ctx := bp.NewDecodeContext(s)
	m.BpProcessor().Process(ctx, nil, m)
}

func (m *Power) BpProcessor() bp.Processor {
	fieldDescriptors := []*bp.MessageFieldProcessor{
		bp.NewMessageFieldProcessor(1, bp.NewUint(8)),
		bp.NewMessageFieldProcessor(2, (PowerStatus(0)).BpProcessor()),
		bp.NewMessageFieldProcessor(3, bp.NewBool()),
	}
	return bp.NewMessageProcessor(false, 11, fieldDescriptors)
}

func (m *Power) BpGetAccessor(di *bp.DataIndexer) bp.Accessor {
	switch di.F() {
	default:
		return nil  // Won't reached
	}
}

func (m *Power) BpSetByte(di *bp.DataIndexer, lshift int, b byte) {
	switch di.F() {
		case 1:
			m.Battery |= (uint8(b) << lshift)
		case 2:
			m.Status |= (PowerStatus(b) << lshift)
		case 3:
			m.IsCharging = bp.Byte2bool(b)
		default:
			return
	}
}

func (m *Power) BpGetByte(di *bp.DataIndexer, rshift int) byte {
	switch di.F() {
		case 1:
			return byte(m.Battery >> rshift)
		case 2:
			return byte(m.Status >> rshift)
		case 3:
			return bp.Bool2byte(m.IsCharging) >> rshift
		default:
			return byte(0) // Won't reached
	}
}

type Network struct {
	// Degree of signal, between 1~10.
	Signal uint8 `json:"signal"` // 4bit
	// The timestamp of the last time received heartbeat packet.
	HeartbeatAt Timestamp `json:"heartbeat_at"` // 64bit
}

// Number of bytes to serialize struct Network
const BYTES_LENGTH_NETWORK uint32 = 9

func (m *Network) Size() uint32 { return 9 }

// Returns string representation for struct Network.
func (m *Network) String() string {
	v, _ := json.Marshal(m)
	return string(v)
}

// Encode struct Network to bytes buffer.
func (m *Network) Encode() []byte {
	ctx := bp.NewEncodeContext(int(m.Size()))
	m.BpProcessor().Process(ctx, nil, m)
	return ctx.Buffer()
}

func (m *Network) Decode(s []byte) {
	ctx := bp.NewDecodeContext(s)
	m.BpProcessor().Process(ctx, nil, m)
}

func (m *Network) BpProcessor() bp.Processor {
	fieldDescriptors := []*bp.MessageFieldProcessor{
		bp.NewMessageFieldProcessor(1, bp.NewUint(4)),
		bp.NewMessageFieldProcessor(2, (Timestamp(0)).BpProcessor()),
	}
	return bp.NewMessageProcessor(false, 68, fieldDescriptors)
}

func (m *Network) BpGetAccessor(di *bp.DataIndexer) bp.Accessor {
	switch di.F() {
	default:
		return nil  // Won't reached
	}
}

func (m *Network) BpSetByte(di *bp.DataIndexer, lshift int, b byte) {
	switch di.F() {
		case 1:
			m.Signal |= (uint8(b) << lshift)
		case 2:
			m.HeartbeatAt |= (Timestamp(b) << lshift)
		default:
			return
	}
}

func (m *Network) BpGetByte(di *bp.DataIndexer, rshift int) byte {
	switch di.F() {
		case 1:
			return byte(m.Signal >> rshift)
		case 2:
			return byte(m.HeartbeatAt >> rshift)
		default:
			return byte(0) // Won't reached
	}
}

type LandingGear struct {
	Status LandingGearStatus `json:"status"` // 2bit
}

// Number of bytes to serialize struct LandingGear
const BYTES_LENGTH_LANDING_GEAR uint32 = 1

func (m *LandingGear) Size() uint32 { return 1 }

// Returns string representation for struct LandingGear.
func (m *LandingGear) String() string {
	v, _ := json.Marshal(m)
	return string(v)
}

// Encode struct LandingGear to bytes buffer.
func (m *LandingGear) Encode() []byte {
	ctx := bp.NewEncodeContext(int(m.Size()))
	m.BpProcessor().Process(ctx, nil, m)
	return ctx.Buffer()
}

func (m *LandingGear) Decode(s []byte) {
	ctx := bp.NewDecodeContext(s)
	m.BpProcessor().Process(ctx, nil, m)
}

func (m *LandingGear) BpProcessor() bp.Processor {
	fieldDescriptors := []*bp.MessageFieldProcessor{
		bp.NewMessageFieldProcessor(1, (LandingGearStatus(0)).BpProcessor()),
	}
	return bp.NewMessageProcessor(false, 2, fieldDescriptors)
}

func (m *LandingGear) BpGetAccessor(di *bp.DataIndexer) bp.Accessor {
	switch di.F() {
	default:
		return nil  // Won't reached
	}
}

func (m *LandingGear) BpSetByte(di *bp.DataIndexer, lshift int, b byte) {
	switch di.F() {
		case 1:
			m.Status |= (LandingGearStatus(b) << lshift)
		default:
			return
	}
}

func (m *LandingGear) BpGetByte(di *bp.DataIndexer, rshift int) byte {
	switch di.F() {
		case 1:
			return byte(m.Status >> rshift)
		default:
			return byte(0) // Won't reached
	}
}

type Position struct {
	Latitude uint32 `json:"latitude"` // 32bit
	Longitude uint32 `json:"longitude"` // 32bit
	Altitude uint32 `json:"altitude"` // 32bit
}

// Number of bytes to serialize struct Position
const BYTES_LENGTH_POSITION uint32 = 12

func (m *Position) Size() uint32 { return 12 }

// Returns string representation for struct Position.
func (m *Position) String() string {
	v, _ := json.Marshal(m)
	return string(v)
}

// Encode struct Position to bytes buffer.
func (m *Position) Encode() []byte {
	ctx := bp.NewEncodeContext(int(m.Size()))
	m.BpProcessor().Process(ctx, nil, m)
	return ctx.Buffer()
}

func (m *Position) Decode(s []byte) {
	ctx := bp.NewDecodeContext(s)
	m.BpProcessor().Process(ctx, nil, m)
}

func (m *Position) BpProcessor() bp.Processor {
	fieldDescriptors := []*bp.MessageFieldProcessor{
		bp.NewMessageFieldProcessor(1, bp.NewUint(32)),
		bp.NewMessageFieldProcessor(2, bp.NewUint(32)),
		bp.NewMessageFieldProcessor(3, bp.NewUint(32)),
	}
	return bp.NewMessageProcessor(false, 96, fieldDescriptors)
}

func (m *Position) BpGetAccessor(di *bp.DataIndexer) bp.Accessor {
	switch di.F() {
	default:
		return nil  // Won't reached
	}
}

func (m *Position) BpSetByte(di *bp.DataIndexer, lshift int, b byte) {
	switch di.F() {
		case 1:
			m.Latitude |= (uint32(b) << lshift)
		case 2:
			m.Longitude |= (uint32(b) << lshift)
		case 3:
			m.Altitude |= (uint32(b) << lshift)
		default:
			return
	}
}

func (m *Position) BpGetByte(di *bp.DataIndexer, rshift int) byte {
	switch di.F() {
		case 1:
			return byte(m.Latitude >> rshift)
		case 2:
			return byte(m.Longitude >> rshift)
		case 3:
			return byte(m.Altitude >> rshift)
		default:
			return byte(0) // Won't reached
	}
}

// Pose in flight. https://en.wikipedia.org/wiki/Aircraft_principal_axes
type Pose struct {
	Yaw int32 `json:"yaw"` // 32bit
	Pitch int32 `json:"pitch"` // 32bit
	Roll int32 `json:"roll"` // 32bit
}

// Number of bytes to serialize struct Pose
const BYTES_LENGTH_POSE uint32 = 12

func (m *Pose) Size() uint32 { return 12 }

// Returns string representation for struct Pose.
func (m *Pose) String() string {
	v, _ := json.Marshal(m)
	return string(v)
}

// Encode struct Pose to bytes buffer.
func (m *Pose) Encode() []byte {
	ctx := bp.NewEncodeContext(int(m.Size()))
	m.BpProcessor().Process(ctx, nil, m)
	return ctx.Buffer()
}

func (m *Pose) Decode(s []byte) {
	ctx := bp.NewDecodeContext(s)
	m.BpProcessor().Process(ctx, nil, m)
}

func (m *Pose) BpProcessor() bp.Processor {
	fieldDescriptors := []*bp.MessageFieldProcessor{
		bp.NewMessageFieldProcessor(1, bp.NewInt(32)),
		bp.NewMessageFieldProcessor(2, bp.NewInt(32)),
		bp.NewMessageFieldProcessor(3, bp.NewInt(32)),
	}
	return bp.NewMessageProcessor(false, 96, fieldDescriptors)
}

func (m *Pose) BpGetAccessor(di *bp.DataIndexer) bp.Accessor {
	switch di.F() {
	default:
		return nil  // Won't reached
	}
}

func (m *Pose) BpSetByte(di *bp.DataIndexer, lshift int, b byte) {
	switch di.F() {
		case 1:
			m.Yaw |= (int32(b) << lshift)
		case 2:
			m.Pitch |= (int32(b) << lshift)
		case 3:
			m.Roll |= (int32(b) << lshift)
		default:
			return
	}
}

func (m *Pose) BpGetByte(di *bp.DataIndexer, rshift int) byte {
	switch di.F() {
		case 1:
			return byte(m.Yaw >> rshift)
		case 2:
			return byte(m.Pitch >> rshift)
		case 3:
			return byte(m.Roll >> rshift)
		default:
			return byte(0) // Won't reached
	}
}

type Flight struct {
	Pose Pose `json:"pose"` // 96bit
	// Velocity at X, Y, Z axis.
	Velocity TernaryInt32 `json:"velocity"` // 96bit
	// Acceleration at X, Y, Z axis.
	Acceleration TernaryInt32 `json:"acceleration"` // 96bit
}

// Number of bytes to serialize struct Flight
const BYTES_LENGTH_FLIGHT uint32 = 36

func (m *Flight) Size() uint32 { return 36 }

// Returns string representation for struct Flight.
func (m *Flight) String() string {
	v, _ := json.Marshal(m)
	return string(v)
}

// Encode struct Flight to bytes buffer.
func (m *Flight) Encode() []byte {
	ctx := bp.NewEncodeContext(int(m.Size()))
	m.BpProcessor().Process(ctx, nil, m)
	return ctx.Buffer()
}

func (m *Flight) Decode(s []byte) {
	ctx := bp.NewDecodeContext(s)
	m.BpProcessor().Process(ctx, nil, m)
}

func (m *Flight) BpProcessor() bp.Processor {
	fieldDescriptors := []*bp.MessageFieldProcessor{
		bp.NewMessageFieldProcessor(1, (&Pose{}).BpProcessor()),
		bp.NewMessageFieldProcessor(2, (TernaryInt32{}).BpProcessor()),
		bp.NewMessageFieldProcessor(3, (TernaryInt32{}).BpProcessor()),
	}
	return bp.NewMessageProcessor(false, 288, fieldDescriptors)
}

func (m *Flight) BpGetAccessor(di *bp.DataIndexer) bp.Accessor {
	switch di.F() {
	case 1:
		return &(m.Pose)
	default:
		return nil  // Won't reached
	}
}

func (m *Flight) BpSetByte(di *bp.DataIndexer, lshift int, b byte) {
	switch di.F() {
		case 2:
			m.Velocity[di.I(0)] |= (int32(b) << lshift)
		case 3:
			m.Acceleration[di.I(0)] |= (int32(b) << lshift)
		default:
			return
	}
}

func (m *Flight) BpGetByte(di *bp.DataIndexer, rshift int) byte {
	switch di.F() {
		case 2:
			return byte(m.Velocity[di.I(0)] >> rshift)
		case 3:
			return byte(m.Acceleration[di.I(0)] >> rshift)
		default:
			return byte(0) // Won't reached
	}
}

type Drone struct {
	Status DroneStatus `json:"status"` // 3bit
	Position Position `json:"position"` // 96bit
	Flight Flight `json:"flight"` // 288bit
	Propellers [4]Propeller `json:"propellers"` // 48bit
	Power Power `json:"power"` // 11bit
	Network Network `json:"network"` // 68bit
	LandingGear LandingGear `json:"landing_gear"` // 2bit
}

// Number of bytes to serialize struct Drone
const BYTES_LENGTH_DRONE uint32 = 65

func (m *Drone) Size() uint32 { return 65 }

// Returns string representation for struct Drone.
func (m *Drone) String() string {
	v, _ := json.Marshal(m)
	return string(v)
}

// Encode struct Drone to bytes buffer.
func (m *Drone) Encode() []byte {
	ctx := bp.NewEncodeContext(int(m.Size()))
	m.BpProcessor().Process(ctx, nil, m)
	return ctx.Buffer()
}

func (m *Drone) Decode(s []byte) {
	ctx := bp.NewDecodeContext(s)
	m.BpProcessor().Process(ctx, nil, m)
}

func (m *Drone) BpProcessor() bp.Processor {
	fieldDescriptors := []*bp.MessageFieldProcessor{
		bp.NewMessageFieldProcessor(1, (DroneStatus(0)).BpProcessor()),
		bp.NewMessageFieldProcessor(2, (&Position{}).BpProcessor()),
		bp.NewMessageFieldProcessor(3, (&Flight{}).BpProcessor()),
		bp.NewMessageFieldProcessor(4, bp.NewArray(false, 4, (&Propeller{}).BpProcessor())),
		bp.NewMessageFieldProcessor(5, (&Power{}).BpProcessor()),
		bp.NewMessageFieldProcessor(6, (&Network{}).BpProcessor()),
		bp.NewMessageFieldProcessor(7, (&LandingGear{}).BpProcessor()),
	}
	return bp.NewMessageProcessor(false, 516, fieldDescriptors)
}

func (m *Drone) BpGetAccessor(di *bp.DataIndexer) bp.Accessor {
	switch di.F() {
	case 2:
		return &(m.Position)
	case 3:
		return &(m.Flight)
	case 4:
		return &(m.Propellers[di.I(0)])
	case 5:
		return &(m.Power)
	case 6:
		return &(m.Network)
	case 7:
		return &(m.LandingGear)
	default:
		return nil  // Won't reached
	}
}

func (m *Drone) BpSetByte(di *bp.DataIndexer, lshift int, b byte) {
	switch di.F() {
		case 1:
			m.Status |= (DroneStatus(b) << lshift)
		default:
			return
	}
}

func (m *Drone) BpGetByte(di *bp.DataIndexer, rshift int) byte {
	switch di.F() {
		case 1:
			return byte(m.Status >> rshift)
		default:
			return byte(0) // Won't reached
	}
}