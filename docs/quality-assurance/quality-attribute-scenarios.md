# Quality Attribute Scenarios

## Usability

### Appropriateness recognizability

This attribute refers to how easy it is for users (especially students and teachers) to understand what the system does and how to use it.

It is important because we have different types of users (students, teachers, admins), and everyone needs to quickly navigate the interface: students need to understand how to start a quiz; teachers need to understand how to add a picture to a question or check grades.

#### Scenario 1: Understanding the interface by the student
- **Source**: User (student)
- **Stimulus**: Logs into the system and searches for a quiz to take
- **Environment**: Any device, browser
- **Artifact**: User interface
- **Response**: Find the quiz you need in less than 10 seconds without any help
- **Response measure**: 90% of students don't ask questions about how to start a quiz

**How to test**: Conduct a usability test - give students the task of "finding and completing a quiz", measure the time and questions.

---

## Reliability

### Maturity

This means the stability of the system under normal load - so that the system does not break when adding quizzes, passing, loading images, etc.

This is important because losing quiz data or scores will cause users to mistrust you.

#### Scenario 2: Stable work when creating a quiz
- **Source**: User (teacher)
- **Stimulus**: Creates a quiz with 10 questions and images
- **Environment**: Production
- **Artifact**: Backend and database
- **Response**: The quiz is saved correctly, without errors
- **Response measure**: 99% success rate of quiz creation without failures

**How to test**: Write unit tests and run load tests - automatically create 10+ quizzes with images.

---

## Maintainability

### Modifiability

#### Scenario 3: adding a new question type (eg matching or coding)

- **Source**: Developer
- **Stimulus**: It is necessary to implement a new type of question into the system (for example, “matching” or “open answered”)
- **Environment**: During development
- **Artifact**: Backend и frontend
- **Response**: Changes are implemented without the need to rewrite existing modules (the architecture supports extensibility)
- **Response measure**:Adding a new type takes no more than 2 days, and all existing tests continue to pass without errors.

**How to test**:  
The functionality of adding and displaying a new type of question has already been tested, for example, through:
- `MatchingQuestions.integration.test.js` — checking the rendering of the corresponding elements
- `EditorForQuestion.test.js` — testing the question editor, including with different types
- `quiz_test.go` — integration tests for creating and receiving quizzes on the backend

All these tests pass successfully after adding a new question type, which confirms the modularity and ease of extending the architecture.