package comms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommsStruct(t *testing.T) {
	datastream_buffer := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	datastream := Datastream{DatastreamBuffer: datastream_buffer}
	assert.Equal(t, datastream_buffer, datastream.DatastreamBuffer)
}

func TestFirstExample(t *testing.T) {
	datastream_buffer := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	datastream := Datastream{DatastreamBuffer: datastream_buffer}
	assert.Equal(t, 7, datastream.StartOfPacketMarkerLocation())
}

func TestSecondExample(t *testing.T) {
	datastream_buffer := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	datastream := Datastream{DatastreamBuffer: datastream_buffer}
	assert.Equal(t, 5, datastream.StartOfPacketMarkerLocation())
}

func TestThirdExample(t *testing.T) {
	datastream_buffer := "nppdvjthqldpwncqszvftbrmjlhg"
	datastream := Datastream{DatastreamBuffer: datastream_buffer}
	assert.Equal(t, 6, datastream.StartOfPacketMarkerLocation())
}

func TestFourthExample(t *testing.T) {
	datastream_buffer := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	datastream := Datastream{DatastreamBuffer: datastream_buffer}
	assert.Equal(t, 10, datastream.StartOfPacketMarkerLocation())
}

func TestFifthExample(t *testing.T) {
	datastream_buffer := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	datastream := Datastream{DatastreamBuffer: datastream_buffer}
	assert.Equal(t, 11, datastream.StartOfPacketMarkerLocation())
}

func TestSixthExample(t *testing.T) {
	datastream_buffer := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	datastream := Datastream{DatastreamBuffer: datastream_buffer}
	assert.Equal(t, 19, datastream.StartOfMessageMarkerLocation())
}

func TestSeventhExample(t *testing.T) {
	datastream_buffer := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	datastream := Datastream{DatastreamBuffer: datastream_buffer}
	assert.Equal(t, 23, datastream.StartOfMessageMarkerLocation())
}

func TestEightExample(t *testing.T) {
	datastream_buffer := "nppdvjthqldpwncqszvftbrmjlhg"
	datastream := Datastream{DatastreamBuffer: datastream_buffer}
	assert.Equal(t, 23, datastream.StartOfMessageMarkerLocation())
}

func TestNinthExample(t *testing.T) {
	datastream_buffer := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	datastream := Datastream{DatastreamBuffer: datastream_buffer}
	assert.Equal(t, 29, datastream.StartOfMessageMarkerLocation())
}

func TestTenthExample(t *testing.T) {
	datastream_buffer := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	datastream := Datastream{DatastreamBuffer: datastream_buffer}
	assert.Equal(t, 26, datastream.StartOfMessageMarkerLocation())
}
