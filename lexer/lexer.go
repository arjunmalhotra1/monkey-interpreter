package lexer

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char).
	readPosition int  // current reading position in input (after current char).
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.position >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	// Here l.position will point to the position we last read from & l.readPosition will point to the position we will next read from.
	l.position = l.readPosition
	l.readPosition += 1
}
