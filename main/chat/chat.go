package chat

import (
	"fmt"
	"time"

	"github.com/spettacolo/chat/main/message"

	"github.com/gdamore/tcell/v2"
)

type Chat struct {
	screen      tcell.Screen
	messages    []message.Message
	inputBuffer string
	scrollPos   int
	style       tcell.Style
}

func NewChat() (*Chat, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := screen.Init(); err != nil {
		return nil, err
	}
	screen.EnableMouse()
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	return &Chat{
		screen:      screen,
		messages:    make([]message.Message, 0),
		inputBuffer: "",
		scrollPos:   0,
		style:       style,
	}, nil
}

func (c *Chat) Run() { // Fixed receiver syntax
	for {
		c.draw()
		event := c.screen.PollEvent()
		switch ev := event.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				c.screen.Fini()
				return
			case tcell.KeyEnter:
				if len(c.inputBuffer) > 0 {
					c.addMessage(c.inputBuffer)
					c.inputBuffer = ""
					// Forza lo scroll in basso quando viene inviato un messaggio
					c.scrollToBottom()
				}
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				if len(c.inputBuffer) > 0 {
					c.inputBuffer = c.inputBuffer[:len(c.inputBuffer)-1]
				}
			case tcell.KeyPgUp:
				c.scroll(-5)
			case tcell.KeyPgDn:
				c.scroll(5)
			case tcell.KeyUp:
				c.scroll(-1)
			case tcell.KeyDown:
				c.scroll(1)
			default:
				if ev.Rune() != 0 {
					c.inputBuffer += string(ev.Rune())
				}
			}
		case *tcell.EventMouse:
			switch ev.Buttons() {
			case tcell.WheelUp:
				c.scroll(-3)
			case tcell.WheelDown:
				c.scroll(3)
			}
		case *tcell.EventResize:
			c.screen.Sync()
			// Mantieni la vista in basso durante il ridimensionamento
			c.scrollToBottom()
		}
	}
}

func (c *Chat) scrollToBottom() { // Fixed receiver syntax
	_, height := c.screen.Size()
	messageArea := height - 3
	c.scrollPos = max(0, len(c.messages)-messageArea)
}

func (c *Chat) scroll(delta int) { // Fixed receiver syntax
	_, height := c.screen.Size()
	messageArea := height - 3
	maxScroll := max(0, len(c.messages)-messageArea)
	newPos := c.scrollPos + delta
	if newPos < 0 {
		c.scrollPos = 0
	} else if newPos > maxScroll {
		c.scrollPos = maxScroll
	} else {
		c.scrollPos = newPos
	}
}

func (c *Chat) draw() { // Fixed receiver syntax
	c.screen.Clear()
	width, height := c.screen.Size()
	messageArea := height - 3
	availableMessages := len(c.messages)
	startIdx := max(0, min(c.scrollPos, availableMessages-messageArea))
	// Disegna i messaggi
	for i := 0; i < messageArea && i < availableMessages; i++ {
		msgIdx := startIdx + i
		if msgIdx < availableMessages {
			msg := c.messages[msgIdx]
			timeStr := msg.Time.Format("15:04")
			content := fmt.Sprintf("[%s] %s", timeStr, msg.Content)
			drawText(c.screen, 0, i, width, content, c.style)
		}
	}
	// Disegna la linea di separazione
	for x := 0; x < width; x++ {
		c.screen.SetContent(x, height-3, '-', nil, c.style)
	}
	// Disegna l'input
	drawText(c.screen, 0, height-2, width, "> "+c.inputBuffer, c.style)
	// Indicatore di scroll
	if len(c.messages) > messageArea {
		scrollPercent := float64(c.scrollPos) / float64(len(c.messages)-messageArea)
		scrollBarPos := int(float64(messageArea-1) * scrollPercent)
		c.screen.SetContent(width-1, scrollBarPos, '█', nil, c.style)
	}
	c.screen.Show()
}

func (c *Chat) addMessage(content string) { // Fixed receiver syntax
	c.messages = append(c.messages, message.Message{
		Sender:      "exampleUser",
		RoomID:      "exampleRoom",
		Content:     content,
		MessageType: "group",
		Time:        time.Now(),
	})
}

func drawText(screen tcell.Screen, x, y, maxWidth int, text string, style tcell.Style) {
	for i, r := range []rune(text) {
		if i >= maxWidth {
			break
		}
		screen.SetContent(x+i, y, r, nil, style)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
