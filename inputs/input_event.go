package inputs

type Action int

var (
	MouseButtonPressed  = Action(0)
	MouseButtonReleased = Action(1)
	KeyPressed          = Action(2)
	KeyHeld             = Action(3)
)

func NewInputMouseEvent(event MouseEvent) InputEvent {
	return InputEvent{eventType: InputEventTypeMouse, mouse: event}
}

func NewInputKeyEvent(event KeyEvent) InputEvent {
	return InputEvent{eventType: InputEventTypeKey, key: event}
}

type InputEventType int

var (
	InputEventTypeMouse = InputEventType(0)
	InputEventTypeKey   = InputEventType(1)
)

type ProducersInput interface {
	Update() []InputEvent
}
type InputEvent struct {
	eventType InputEventType
	mouse     MouseEvent
	key       KeyEvent
	clientID  string
}

func (event InputEvent) IsMouse() bool {
	return event.eventType == InputEventTypeMouse
}

func (event InputEvent) GetClientID() string {
	return event.clientID
}

func (event InputEvent) GetMouse() MouseEvent {
	return event.mouse
}

func (event InputEvent) IsKey() bool {
	return event.eventType == InputEventTypeKey
}

func (event InputEvent) GetKey() KeyEvent {
	return event.key
}

type InputsProducer struct {
	producers []ProducersInput
}

func NewInputsProducer(procuders ...ProducersInput) *InputsProducer {
	return &InputsProducer{producers: procuders}
}

func (producer *InputsProducer) Update() []InputEvent {
	var events []InputEvent
	for _, p := range producer.producers {
		events = append(events, p.Update()...)
	}

	return events
}
