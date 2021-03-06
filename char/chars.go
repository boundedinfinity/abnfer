package char

type Char byte

func (c Char) Byte() byte     { return byte(c) }
func (c Char) String() string { return string(c) }

const (
	NULL                      Char = 0x00
	START_OF_HEADER           Char = 0x01
	START_OF_TEXT             Char = 0x02
	END_OF_TEXT               Char = 0x03
	END_OF_TRANSMISSION       Char = 0x04
	ENQUIRY                   Char = 0x05
	ACKNOWLEDGE               Char = 0x06
	BELL                      Char = 0x07
	BACKSPACE                 Char = 0x08
	HORIZONTAL_TAB            Char = 0x09
	LINE_FEED                 Char = 0x0A
	VERTICAL_TAB              Char = 0x0B
	FORM_FEED                 Char = 0x0C
	CARRIAGE_RETURN           Char = 0x0D
	SHIFT_OUT                 Char = 0x0E
	SHIFT_IN                  Char = 0x0F
	DATA_LINK_ESCAPE          Char = 0x10
	DEVICE_CONTROL_1          Char = 0x11
	DEVICE_CONTROL_2          Char = 0x12
	DEVICE_CONTROL_3          Char = 0x13
	DEVICE_CONTROL_4          Char = 0x14
	NEGATIVE_ACKNOLEDGE       Char = 0x15
	SYNCHRONIZE               Char = 0x16
	END_OF_TRANSMISSION_BLOCK Char = 0x17
	CANCEL                    Char = 0x18
	END_OF_MEDIUM             Char = 0x19
	SUBSTITUE                 Char = 0x1A
	ESCAPE                    Char = 0x1B
	FILE_SEPARATOR            Char = 0x1C
	GROUP_SEPARATOR           Char = 0x1D
	RECORD_SEPARATOR          Char = 0x1E
	UNIT_SEPARATOR            Char = 0x1F
	SPACE                     Char = 0x20
	EXCLAMATION_MARK          Char = 0x21
	DOUBLE_QUOTE              Char = 0x22
	NUMBER                    Char = 0x23
	DOLLAR                    Char = 0x24
	PERCENT                   Char = 0x25
	AMPERSAND                 Char = 0x26
	SINGLE_QUOTE              Char = 0x27
	LEFT_PARENTHESIS          Char = 0x28
	RIGHT_PARENTHESIS         Char = 0x29
	ASTERISK                  Char = 0x2A
	PLUS                      Char = 0x2B
	COMMA                     Char = 0x2C
	MINUS                     Char = 0x2D
	PERIOD                    Char = 0x2E
	SLASH                     Char = 0x2F
	ZERO                      Char = 0x30
	ONE                       Char = 0x31
	TWO                       Char = 0x32
	THREE                     Char = 0x33
	FOUR                      Char = 0x34
	FIVE                      Char = 0x35
	SIX                       Char = 0x36
	SEVEN                     Char = 0x37
	EIGHT                     Char = 0x38
	NINE                      Char = 0x39
	COLON                     Char = 0x3A
	SEMICOLON                 Char = 0x3B
	LESS_THAN                 Char = 0x3C
	EQUAL_SIGN                Char = 0x3D
	GREATER_THAN              Char = 0x3E
	QUESTION_MARK             Char = 0x3F
	AT_SIGN                   Char = 0x40
	UPPERCASE_A               Char = 0x41
	UPPERCASE_B               Char = 0x42
	UPPERCASE_C               Char = 0x43
	UPPERCASE_D               Char = 0x44
	UPPERCASE_E               Char = 0x45
	UPPERCASE_F               Char = 0x46
	UPPERCASE_G               Char = 0x47
	UPPERCASE_H               Char = 0x48
	UPPERCASE_I               Char = 0x49
	UPPERCASE_J               Char = 0x4A
	UPPERCASE_K               Char = 0x4B
	UPPERCASE_L               Char = 0x4C
	UPPERCASE_M               Char = 0x4D
	UPPERCASE_N               Char = 0x4E
	UPPERCASE_O               Char = 0x4F
	UPPERCASE_P               Char = 0x50
	UPPERCASE_Q               Char = 0x51
	UPPERCASE_R               Char = 0x52
	UPPERCASE_S               Char = 0x53
	UPPERCASE_T               Char = 0x54
	UPPERCASE_U               Char = 0x55
	UPPERCASE_V               Char = 0x56
	UPPERCASE_W               Char = 0x57
	UPPERCASE_X               Char = 0x58
	UPPERCASE_Y               Char = 0x59
	UPPERCASE_Z               Char = 0x5A
	LEFT_SQUARE_BRACKET       Char = 0x5B
	BACKSLASH                 Char = 0x5C
	RIGHT_SQUARE_BRACKET      Char = 0x5D
	CARET                     Char = 0x5E
	UNDERSCORE                Char = 0x5F
	GRAVE                     Char = 0x60
	LOWERCASE_A               Char = 0x61
	LOWERCASE_B               Char = 0x62
	LOWERCASE_C               Char = 0x63
	LOWERCASE_D               Char = 0x64
	LOWERCASE_E               Char = 0x65
	LOWERCASE_F               Char = 0x66
	LOWERCASE_G               Char = 0x67
	LOWERCASE_H               Char = 0x68
	LOWERCASE_I               Char = 0x69
	LOWERCASE_J               Char = 0x6A
	LOWERCASE_K               Char = 0x6B
	LOWERCASE_L               Char = 0x6C
	LOWERCASE_M               Char = 0x6D
	LOWERCASE_N               Char = 0x6E
	LOWERCASE_O               Char = 0x6F
	LOWERCASE_p               Char = 0x70
	LOWERCASE_Q               Char = 0x71
	LOWERCASE_R               Char = 0x72
	LOWERCASE_S               Char = 0x73
	LOWERCASE_T               Char = 0x74
	LOWERCASE_U               Char = 0x75
	LOWERCASE_V               Char = 0x76
	LOWERCASE_W               Char = 0x77
	LOWERCASE_X               Char = 0x78
	LOWERCASE_Y               Char = 0x79
	LOWERCASE_Z               Char = 0x7A
	LEFT_CURLY_BRACKET        Char = 0x7B
	VERTICAL_BAR              Char = 0x7C
	RIGHT_CURLY_BRACKET       Char = 0x7D
	TILDE                     Char = 0x7E
	DELETE                    Char = 0x7F
)
