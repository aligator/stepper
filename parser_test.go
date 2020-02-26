package parser

import (
	"testing"
)

const testHtml = `<html>
    <body>
        <h1>Title</h1>
        <div id="main" class="test">
            <p>Hello <em>world</em>!</p>
        </div>
    </body>
</html>`

func getBasicParser() BasicParser {
	return NewBasicParser(testHtml)
}

func TestPeek(t *testing.T) {
	p := getBasicParser()

	testPeek := func (pos int, expected rune) {
		p.position = pos
		r := p.Peek()
		if r != expected {
			t.Errorf("the rune at position %d should be a %c but it is a %c", pos, expected, r)
		}

		if p.position != pos {
			t.Errorf("the position should not change during a peek. Expected: %d Actual: %d", pos, p.position)
		}
	}

	testPeek(0, '<')
	testPeek(15, 'y')
	testPeek(len(testHtml) - 1, '>')
}

func TestNext(t *testing.T) {
	p := getBasicParser()

	testNext := func (pos int, expected1 rune, expected2 rune) {
		p.position = pos

		r := p.Next()
		if r != expected1 {
			t.Errorf("The rune at position %d should be a %c but it is a %c.", pos, expected1, r)
		}

		r = p.Next()
		if r != expected2 {
			t.Errorf("The rune at position %d should be a %c but it is a %c.", pos+1, expected2, r)
		}

		if p.position != pos+2 {
			t.Errorf("The position should change during a Next(). Expected: %d Actual: %d", pos+2, p.position)
		}
	}

	testNext(0, '<', 'h')
	testNext(12, 'b', 'o')
}

func TestWhile(t *testing.T) {
	p := getBasicParser()
	expected := "<html"
	actual := string(p.While(func(r rune) bool {
		return r != '>'
	}))

	if expected != actual {
		t.Errorf("While() did not read correctly. Expected: %s Actual: %s", expected, actual)
	}

	if len(expected) != p.position {
		t.Errorf("While() should increment the position correctly. Expected: %d Actual: %d", len(expected), p.position)
	}
}
