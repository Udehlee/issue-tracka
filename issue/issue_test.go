package issue

import (
	"os"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	is := NewMemory()

	issueTitle := "title"
	issueText := "text"

	createdissue, err := is.Create(issueTitle, issueText)
	if err != nil {
		t.Errorf("got %v", err)
	}

	if createdissue.Title != issueTitle {
		t.Errorf("Expected %q, got %q", issueTitle, createdissue.Title)
	}

	if createdissue.Text != issueText {
		t.Errorf("Expected %q, got %q", issueText, createdissue.Text)
	}
}

func TestSave(t *testing.T) {
	is := NewMemory()

	i := CreatedIssue{
		ID:        "9bq34s",
		Title:     "Title",
		Text:      "Text",
		CreatedAt: time.Now(),
	}

	// Create a temporary file in a temporary directory
	tempFile, err := os.CreateTemp("", "issue_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	err = is.Save(i, tempFile.Name())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read temp file: %v", err)
	}

	expectedData := is.JSON([]CreatedIssue{i})
	if string(data) != expectedData {
		t.Errorf("Expected %s, got %s", expectedData, string(data))
	}
}

func TestList(t *testing.T) {

	i := []CreatedIssue{
		{
			ID:        "9bq34s",
			Title:     "Title",
			Text:      "Text",
			CreatedAt: time.Now(),
		},
		{
			ID:        "9bq34t",
			Title:     "Another Title",
			Text:      "Another Text",
			CreatedAt: time.Now(),
		},
	}

	// Create a temporary file in a temporary directory
	tempFile, err := os.CreateTemp("", "issue_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	is := NewMemory(tempFile.Name())

	data := is.JSON(i)
	err = os.WriteFile(tempFile.Name(), []byte(data), 0644)
	if err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}

	// retrieve issues from file
	issues, err := is.List()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(issues) == 0 {
		t.Errorf("Expected at least 1 issue, got %d", len(issues))
	}

	if issues[0].ID != "9bq34s" || issues[0].Title != "Title" {
		t.Errorf("expected values doesnt match: got ID %s and Title %s", issues[0].ID, issues[0].Title)
	}

}

func TestOpen(t *testing.T) {

	i := CreatedIssue{
		ID:        "9bq34s",
		Title:     "Title",
		Text:      "Text",
		CreatedAt: time.Now(),
	}

	// Create a temporary file
	tempFile, err := os.CreateTemp("", "issue_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	is := NewMemory(tempFile.Name())

	data := is.JSON([]CreatedIssue{i})
	err = os.WriteFile(tempFile.Name(), []byte(data), 0644)
	if err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}

	issue, err := is.Open("9bq34s")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if issue.ID != i.ID {
		t.Errorf("Expected ID '%s', got ID '%s'", i.ID, issue.ID)
	}

}
