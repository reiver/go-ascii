package ascii


const (
	Null                    byte =   0
	StartOfHeading          byte =   1
	StartOfText             byte =   2
	EndOfText               byte =   3
	EndOfTransmission       byte =   4
	Enquiry                 byte =   5
	Acknowledgment          byte =   6
	Bell                    byte =   7
	BackSpace               byte =   8
	HorizontalTab           byte =   9
	LineFeed                byte =  10
	VerticalTab             byte =  11
	FormFeed                byte =  12
	CarriageReturn          byte =  13
	ShiftOut                byte =  14
	ShiftIn                 byte =  15
	DataLinkEscape          byte =  16
	DeviceControl1          byte =  17
	DeviceControl2          byte =  18
	DeviceControl3          byte =  19
	DeviceControl4          byte =  20
	NegativeAcknowledgement byte =  21
	SynchronousIdle         byte =  22
	EndOfTransmitBlock      byte =  23
	Cancel                  byte =  24
	EndOfMedium             byte =  25
	Substitute              byte =  26
	Escape                  byte =  27
	FileSeparator           byte =  28
	GroupSeparator          byte =  29
	RecordSeparator         byte =  30
	UnitSeparator           byte =  31
	
	Delete                  byte = 127
)
