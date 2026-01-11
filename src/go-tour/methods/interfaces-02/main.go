package main

import "fmt"

// 1. The Contract (The Mask)
// "I don't care who you are, as long as you can Save()."
type Saver interface {
	Save(username string)
}

// 2. Real Implementation (Production)
type PostgresSaver struct {
	ConnectionString string
}

// The logic lives here effectively ONCE in memory
func (p PostgresSaver) Save(name string) {
	fmt.Println("Connecting to", p.ConnectionString, "-> Saved:", name)
}

// 3. Mock Implementation (Testing)
type MockSaver struct {
	StoredUsers []string // Specific field for testing only!
}

func (m *MockSaver) Save(name string) {
	fmt.Println("Mock: Adding", name, "to array")
	m.StoredUsers = append(m.StoredUsers, name)
}

// 4. The Business Logic
// It accepts the INTERFACE, not the specific struct.
func CreateUser(s Saver, name string) {
	s.Save(name)
}

func main() {
	// Usage 1: Real
	db := PostgresSaver{ConnectionString: "postgres://..."}
	CreateUser(db, "Emilia")

	// Usage 2: Test
	mock := &MockSaver{}

	CreateUser(mock, "Rem")

	fmt.Println("Mock stored users:", mock.StoredUsers)

	// Safety: You CANNOT pass an empty struct that doesn't work.
	// CreateUser(Vertex{}, "Crash") // COMPILE ERROR! Won't build.
}
