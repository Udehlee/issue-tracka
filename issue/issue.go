package issue

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type CreatedIssue struct {
	ID        string     `json:"issue_id"`
	Title     string     `json:"title"`
	Text      string     `json:"text"`
	CreatedAt time.Time  `json:"created_at"`
	Comments  []*Comment `json:"comment"`
}

type Memory struct {
	FilePath string
}

func NewMemory(filePath ...string) *Memory {
	m := &Memory{
		FilePath: "issue/issue.json",
	}

	if len(filePath) > 0 {
		m.FilePath = filePath[0]
	}
	return m
}

func (m *Memory) Create(issueTitle, issueText string) (CreatedIssue, error) {
	id := GenerateUUID()

	i := CreatedIssue{
		ID:        id,
		Title:     issueTitle,
		Text:      issueText,
		CreatedAt: time.Now(),
		Comments:  []*Comment{},
	}

	return i, nil
}

// Save saves created issue to the generated id
func (m *Memory) Save(issue CreatedIssue, filePath string) error {
	i := []CreatedIssue{}

	i = append(i, issue)
	data := m.JSON(i)

	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}

// List lists all issue in the memory
func (m *Memory) List() ([]CreatedIssue, error) {
	i := []CreatedIssue{}

	file, err := m.ReadFromFile()
	if err != nil {
		return i, fmt.Errorf("error reading file: %w", err)
	}

	if len(file) == 0 {
		return i, fmt.Errorf("no issues found")
	}

	for _, v := range file {
		issue := CreatedIssue{
			ID:    v.ID,
			Title: v.Title,
		}
		i = append(i, issue)
	}

	return i, nil
}

// Open a issue with a specified id
func (m *Memory) Open(Id string) (CreatedIssue, error) {
	i := CreatedIssue{}

	file, err := m.ReadFromFile()
	if err != nil {
		return i, fmt.Errorf("failed to read file: %v", err)
	}

	for i := range file {
		if file[i].ID == Id {
			return file[i], nil
		}
	}

	return i, fmt.Errorf("issue with ID %s not found", Id)
}

// AddComment adds new comment to existing issue
func (m *Memory) AddComment(Id, name, text string) error {
	issueID, err := m.FindIssueByID(Id)
	if err != nil {
		return err
	}

	newComment := Comment{
		Name: name,
		Text: text,
	}

	issueID.Comments = append(issueID.Comments, &newComment)

	fmt.Println("Comment added successfully")
	return nil
}

func (m *Memory) FindIssueByID(Id string) (*CreatedIssue, error) {
	file, err := m.ReadFromFile()
	if err != nil {
		return nil, fmt.Errorf("failed reading file: %v", err)
	}

	for i := range file {
		if file[i].ID == Id {
			return &file[i], nil
		}
	}
	return nil, fmt.Errorf("id not found")
}

func (m *Memory) ReadFromFile() ([]CreatedIssue, error) {
	data, err := os.ReadFile(m.FilePath)
	if err != nil {
		return nil, err
	}

	i := []CreatedIssue{}

	err = json.Unmarshal(data, &i)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return i, nil
}

func (m *Memory) JSON(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return ""
	}
	return string(jsonData)
}

func GenerateUUID() string {
	Id := uuid.New().String()

	if len(Id) > 6 {
		Id = Id[:6]
	}

	return Id
}
