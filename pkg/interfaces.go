package pkg

// SomeService is an example inteface to use in initial tests
type SomeService interface {
	SomeOperation() bool
}

// Article describes what the Article should look like
type Article struct {
	ID           string `json:"id"`
	AmbientVideo string `json:"ambientVideo"`
	StoryType    string `json:"storyType"`
	Intro        string `json:"intro"`
}
