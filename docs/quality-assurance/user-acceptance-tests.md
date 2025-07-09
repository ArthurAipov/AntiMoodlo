# User Acceptance Tests (UAT) with Acceptance Criteria

---

## UAT 1: Login as a Teacher – Passed

### Acceptance Criteria:
- The login page is accessible.
- The system accepts valid teacher credentials.
- The system successfully authenticates the user.
- The user is redirected to the teacher's dashboard after login.

### User Acceptance Test:
**GIVEN:**
- The system is available.
- A valid teacher account exists with the provided login and password.

**WHEN:**
- The user enters the provided login and password on the login page.
- The user clicks the "Login" button.

**THEN:**
- The system logs the user in.
- The system redirects the user to the teacher's dashboard.

**Task for Customer:**  
Please try to log in as a teacher using the provided login and password.

---

## UAT 2: Create a Quiz via GUI – Passed

### Acceptance Criteria:
- The quiz creation page is accessible to logged-in teachers.
- The user can input a quiz title.
- The user can set start and end dates.
- The user can add at least one question.
- The user can select the question type.
- The quiz is successfully saved.
- The newly created quiz appears in the list of available quizzes.

### User Acceptance Test:
**GIVEN:**
- The user is logged in as a teacher.
- The GUI system for creating quizzes is accessible.

**WHEN:**
- The user opens the quiz creation page.
- The user enters a Quiz Title.
- The user sets a Start Date and End Date.
- The user adds at least one question.
- The user selects a question type (e.g. Multiple Choice, Matching, Open-Ended).
- The user saves the quiz.

**THEN:**
- The quiz is successfully created and appears in the list of available quizzes.

**Task for Customer:**  
Please try to create a quiz using the GUI system by adding a title, start and end dates, at least one question, and selecting a question type.

---

## UAT 3: Preview the Created Quiz – Passed

### Acceptance Criteria:
- The preview function is available for created quizzes.
- The system displays the quiz title.
- The system displays all added questions.
- The system displays all answer options correctly.

### User Acceptance Test:
**GIVEN:**
- The user is logged in as a teacher.
- There is at least one created quiz.

**WHEN:**
- The user selects the quiz.
- The user clicks the "Preview" button.

**THEN:**
- The system correctly displays the quiz preview with all questions and options.

**Task for Customer:**  
Please try to preview the quiz you have just created and verify that all questions and options are correctly displayed.

---

## UAT 4: Upload URL of Picture for Question – Passed

### Acceptance Criteria:
- The question creation page includes a field to enter a picture URL.
- The user can successfully input and save the picture URL.
- The saved picture URL is correctly stored and visible in the question's details.

### User Acceptance Test:
**GIVEN:**
- The user is logged in as a teacher.
- The question creation page is accessible.

**WHEN:**
- The user creates a new question.
- The user enters the question text.
- The user enters a picture URL in the designated field.
- The user saves the question.

**THEN:**
- The system saves the question with the provided picture URL.
- The picture URL is visible in the question’s details.

**Task for Customer:**  
Please try to create a new question by adding a picture URL and verify that the picture URL is correctly saved and displayed in the question details.

---

## UAT 5: Preview Question with Picture – Passed

### Acceptance Criteria:
- The preview function correctly displays the picture attached to the question via URL.
- The system shows both the question text and the picture.

### User Acceptance Test:
**GIVEN:**
- The user is logged in as a teacher.
- There is a question with an attached picture URL.

**WHEN:**
- The user selects the question.
- The user clicks the "Preview" button.

**THEN:**
- The system correctly displays the question text and the picture.

**Task for Customer:**  
Please try to preview a question with a picture and verify that both the question text and the picture are correctly displayed.

---

# Results of User Testing

All tests were successfully passed.

During **UAT 4: Upload URL of Picture for Question**, the customer requested an additional feature: the ability to upload pictures for answer options as well.

"""