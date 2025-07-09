import React, { useState, useEffect } from "react";
import yaml from "js-yaml";
import { FiLogOut, FiUpload, FiEye, FiX, FiCheck, FiPlus, FiTrash2, FiSave, FiCode, FiRotateCcw, FiRotateCw} from "react-icons/fi";
import { useAuth } from "../context/AuthContext";
import api from "../services/Api";
import Timer from "../components/Timer";
import "../styles/ConfigUploadPage.css";
import { v4 as uuidv4 } from 'uuid';
import Sidebar from "../components/Sidebar";
import Quiz from "../components/Quiz";
import EditorForQuestion from "../components/EditorForQuestion";


//штука чтобы startdate и enddate определялись и работали корректно
function parseMoscow(iso) {
  //преобразует строку в Date с учетом московского времени
  if (iso instanceof Date || typeof iso === "number") {
    return new Date(iso);
  }
  if (typeof iso !== "string") {
    iso = String(iso);
  }
  if (/[+\-]\d{2}:\d{2}$|Z$/.test(iso)) {
    return new Date(iso);
  }
  const [date, time] = iso.split("T");
  const [Y, M, D] = date.split("-").map(Number);
  const [h, m, s = 0] = time.split(":").map(Number);
  return new Date(Date.UTC(Y, M - 1, D, h - 3, m, s));
}

function getNowMoscow() {
  const nowLocal = new Date();
  const offsetLocalMin = nowLocal.getTimezoneOffset();
  const offsetMoscowMin = 3 * 60;
  const deltaMs = (offsetMoscowMin + offsetLocalMin) * 60_000;
  return new Date(nowLocal.getTime() + deltaMs);
}

//вроде как окно, открывающееся при просмотре превью квиза
function QuizModal({ visible, onClose, quizConfig }) {
  if (!visible) return null; //чтобы окно не высвечивалось когда не надо 

  return (
    <div className="modal-overlay">
      <div className="modal-content"> 
        <button
          onClick={onClose}
          aria-label="Close modal"
          className="close-modal"
        >
          <FiX />
        </button>
        {/*строка ниже нужна чтобы рендерился сам квиз*/}
        <Quiz quizConfig={quizConfig} />
      </div>
    </div>
  );
}

//просто декоративная штука, не трогаем
function CustomCheckbox({ checked, isRadio }) {
  return (
    <span
      className={`custom-checkbox ${isRadio ? 'custom-radio' : ''} ${checked ? 'checked' : ''}`}
    >
      {checked && (
        isRadio ? (
          <span className="radio-dot" />
        ) : (
          <FiCheck className="check-icon" />
        )
      )}
    </span>
  );
}

/*useState(...) хранит всё — от заголовка до YAML-текста

useEffect(...) при монтировании вызывает getQuizzes() и сохраняет результат

handleFileUpload(...) — читает YAML-файл, парсит его и устанавливает текст в yamlText. (Сейчас он не делает setQuizConfig — это надо будет поправить для полной загрузки данных)
 */
function ConfigUploadPage() {
  const [quizStart, setQuizStart] = useState("");
  const [quizEnd, setQuizEnd] = useState("");
  const [quizDurationInput, setQuizDurationInput] = useState(60);
  const [quizConfig, setQuizConfig] = useState(null);
  const [error, setError] = useState(null);
  const [showModal, setShowModal] = useState(false);
  const [sidebarWidth, setSidebarWidth] = useState(60);
  const [activeTab, setActiveTab] = useState('upload');
  const [savedQuizzes, setSavedQuizzes] = useState([]);
  const [selectedQuizResults, setSelectedQuizResults] = useState(null);
  const [selectedQuizTitle, setSelectedQuizTitle] = useState("");
  const [selectedResultsQuiz, setSelectedResultsQuiz] = useState(null);
  const [grades, setGrades] = useState([]);
  const [gradesLoading, setGradesLoading] = useState(false);

  const [isLoading, setIsLoading] = useState(false);
  const [quizTitle, setQuizTitle] = useState("");
  const [quizDescription, setQuizDescription] = useState("");
  const [questions, setQuestions] = useState([]);
  const [yamlText, setYamlText] = useState(() => {
    return localStorage.getItem("autosaveYaml") || "";
  });
  useEffect(() => {
  const timeout = setTimeout(() => {
    if (yamlText.trim()) {
      localStorage.setItem("autosaveYaml", yamlText);
    }
  }, 1000); // сохранить через 1 сек после изменений

  return () => clearTimeout(timeout);
}, [yamlText]);

  const [yamlHistory, setYamlHistory]   = useState([""]);
  const [historyIndex, setHistoryIndex] = useState(0);
  const [yamlError, setYamlError] = useState(null);
  const [quizId,    setQuizId]    = useState(null);
  const [currentQuestion, setCurrentQuestion] = useState({
    id: 1,
    text: "",
    type: "single",
    points: 1,
    options: ["", ""],
    imageurl: "", // картинка самого вопроса
    correctOption: null,
    correctOptions: [],
    left_items: [""],
    right_items: [""],
    correct_matches: {},
    correctAnswerText: ""  
  });

  const [expandedQuestionIds, setExpandedQuestionIds] = useState([]);
const toggleExpand = id => {
  setExpandedQuestionIds(prev =>
    prev.includes(id)
      ? prev.filter(x => x !== id)
      : [...prev, id]
  );
};

  useEffect(() => {
    const loadQuizzes = async () => {
      try {
        setIsLoading(true);
        const response = await api.getQuizzes();
        const allQuizzes = response.data;

        if (user.userrole === 3) {
          // админ видит все квизы
          setSavedQuizzes(allQuizzes);
        } else if (user.userrole === 2) {
          // преподаватель видит только свои
          const myQuizzes = allQuizzes.filter(q => q.userid === user.userid);
          setSavedQuizzes(myQuizzes);
        }
      } catch (err) {
        console.error("Failed to load quizzes:", err);
        const saved = localStorage.getItem('savedQuizzes');
        if (saved) {
          setSavedQuizzes(JSON.parse(saved));
        }
      } finally {
        setIsLoading(false);
      }
    };

    loadQuizzes();
  }, []);

  const handleYamlChange = (e) => {
    const value = e.target.value;
    if (value === yamlText) return;
    setYamlText(value);
    setYamlHistory((prev) => {
      const next = prev.slice(0, historyIndex + 1);
      next.push(value);
      setHistoryIndex(next.length - 1);
      return next;
    });
  };
  
  const handleUndo = () => {
    if (historyIndex === 0) return;
    const idx = historyIndex - 1;
    setHistoryIndex(idx);
    setYamlText(yamlHistory[idx]);
  };
  
  const handleRedo = () => {
    if (historyIndex >= yamlHistory.length - 1) return;
    const idx = historyIndex + 1;
    setHistoryIndex(idx);
    setYamlText(yamlHistory[idx]);
  };
  
  const handleFileUpload = async (event) => {
    const file = event.target.files[0];
    if (!file) return;

    try {
      setIsLoading(true);
      const reader = new FileReader();
      reader.onload = async (e) => {
        const text = e.target.result;
        setYamlText(text);
        setYamlHistory([text]); // сбрасываем стек истории на только что прочитанный файл
        setHistoryIndex(0);
        try {
    const text = await file.text();
    const parsed = yaml.load(text);
  } catch (err) {
    console.error("❌ Ошибка при разборе YAML:", err);
    alert("Не удалось прочитать YAML-файл");
  } 
      };
      reader.readAsText(file);
    } catch (err) {
      setError("Ошибка загрузки файла: " + err.message);
      setQuizConfig(null);
    } finally {
      setIsLoading(false);
    }
  };

  //Экспортирует YAML-файл на основе текущего текста yamlText.
  const exportYaml = () => {
  let filename = "quiz_config";

  try {
    const parsed = yaml.load(yamlText);

    if (parsed?.quiz?.title) {
      filename = parsed.quiz.title
        .trim()
        .replace(/\s+/g, "_")       // заменяем пробелы на "_"
        .replace(/[^\w\-]/g, "")    // убираем все кроме букв, цифр, _
        .slice(0, 50);              // ограничим длину имени файла
    }
  } catch (e) {
    console.warn("Could not parse YAML for filename:", e.message);
  }

  const blob = new Blob([yamlText], { type: "text/yaml" });
  const url = URL.createObjectURL(blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = `${filename || "quiz_config"}.yaml`;
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
  URL.revokeObjectURL(url);
};



//открывает модальное окно предпросмотра (превью) квиза, если YAML корректный
  const openPreviewFromYaml = () => {
    try {
      const parsed = yaml.load(yamlText);
      setQuizConfig(parsed);
      setShowModal(true);
      setYamlError(null);
    } catch (e) {
      setYamlError(`YAML parsing error: ${e.message}`);
    }
  };

//загружает YAML в GUI-редактор (в поля формы)
  const openInGuiCreator = () => {
    try {
      const parsed = yaml.load(yamlText);
      if (!parsed || !parsed.quiz) {
        throw new Error("Invalid YAML structure: missing 'quiz' root");
      }
      
      setQuizTitle(parsed.quiz.title || "");
      setQuizDescription(parsed.quiz.description || "");
      setQuestions(parsed.quiz.questions || []);
      
      if (parsed.quiz.duration) {
        setQuizDurationInput(parsed.quiz.duration);
      }
      
      if (parsed.quiz.start) {
        const date = new Date(parsed.quiz.start);
        const offsetDate = new Date(date.getTime() - date.getTimezoneOffset() * 60000);
        setQuizStart(offsetDate.toISOString().slice(0, 16));
      }

      if (parsed.quiz.end) {
        const date = new Date(parsed.quiz.end);
        const offsetDate = new Date(date.getTime() - date.getTimezoneOffset() * 60000);
        setQuizEnd(offsetDate.toISOString().slice(0, 16));
      }

      setActiveTab('create');
      setQuizId(null);
      setYamlError(null);
    } catch (e) {
      setYamlError(`YAML parsing error: ${e.message}`);
    }
  };

//достаем юзер айди
const { user } = useAuth();
const userId = user?.userid;

//проверка, можно ли добавить текущий вопрос (все ли поля заполнены корректно)
const isDisabled = () => {
  const textEmpty = !currentQuestion.text.trim();

  if (currentQuestion.type === "single") {
    return (
      textEmpty ||
      currentQuestion.correctOption === null ||
      currentQuestion.options.every(o => {
        const t = typeof o === "string" ? o : o.text;
        return !t.trim();
      })
    );
  }
  
  if (currentQuestion.type === "multiple") {
    return (
      textEmpty ||
      currentQuestion.correctOptions.length === 0 ||
      currentQuestion.options.every(o => {
        const t = typeof o === "string" ? o : o.text;
        return !t.trim();
      })
    );
  }

  if (currentQuestion.type === "matching") {
    const leftFilled = currentQuestion.left_items?.filter(item => item.trim()).length;
    const rightFilled = currentQuestion.right_items?.filter(item => item.trim()).length;
    const matches = currentQuestion.correct_matches || {};
    const allMatched = leftFilled > 0 && rightFilled > 0 && leftFilled === Object.keys(matches).length;
    return textEmpty || !allMatched;
  }
  
    if (currentQuestion.type === "open") {
    return textEmpty || !currentQuestion.correctAnswerText.trim();
  }
  return true;
};
//сохраняет квиз в бд и обновляет список в интерфейсе
const saveQuiz = async () => {
  // валидация
  if (!quizTitle.trim() || questions.length === 0) {
    alert("Нужно задать заголовок и хотя бы один вопрос");
    return;
  }

  try {
    setIsLoading(true);
    console.log("🚀 Questions перед сохранением:", questions);
    await api.saveQuizToServer({
      quizTitle,                    // строка
      description: quizDescription, // строка (может быть пустой)
      duration: quizDurationInput,  // число
      userId: user?.userid,         // из AuthContext
      questions                     // массив вопросов, как у тебя сейчас
    });

    alert("Квиз успешно сохранён!");
    // (по желанию) сбросить состояния:
    // setQuizTitle("");
    // setQuizDescription("");
    // setQuestions([]);
  } catch (err) {
    console.error("Ошибка при сохранении квиза:", err);
    alert("Ошибка при сохранении квиза");
  } finally {
    setIsLoading(false);
  }
};

// отрисовалась кнопка updateQuiz — добавляем саму логику
const updateQuiz = async () => {
  if (!quizTitle.trim() || questions.length === 0) {
    alert("Нужно задать заголовок и хотя бы один вопрос");
    return;
  }
  try {
    setIsLoading(true);

    // В ConfigUploadPage.js, внутри функции updateQuiz перед вызовом API:

const isoStart = quizStart
  ? new Date(quizStart).toISOString()
  : null;
const isoEnd = quizEnd
  ? new Date(quizEnd).toISOString()
  : null;

// И затем:
await api.updateQuizMeta(quizId, {
  id:           quizId,
  courseid:     1,
  title:        quizTitle,
  description:  quizDescription,
  duration:     quizDurationInput,
  startdate:    isoStart,
  enddate:      isoEnd,
  maxgrade:     100,
  stateid:      1,
  submiteddate: new Date().toISOString(),
  userid:       userId
});


    // 2) Пробегаем по каждому вопросу и обновляем его и ответы
    const typeMap = { single:1, multiple:2, open:3, matching:4 };

    for (const q of questions) {
      // 2.1) обновляем вопрос
      await api.updateQuestion(q.id, {
        id:             q.id,
        questiontext:   q.question,
        questiontypeid: typeMap[q.type],
        quizid:         quizId,
        imageurl:       q.imageurl || "",
        points:         q.points
      });

      // 2.2) забираем текущие ответы с сервера
      const ans = await api.getAllAnswers(q.id);
      const { options, correctAnswers, matchPairs, openAnswers } = ans.data;

      if (q.type === "single" || q.type === "multiple") {
        // — обновляем опции
        for (let i = 0; i < q.options.length; i++) {
          const text = q.options[i];
          const existing = options[i];
          if (existing) {
            await api.updateOption(existing.id, {
              id:         existing.id,
              optiontext: text,
              questionid: q.id
            });
          } else {
            await api.createOption(q.id, {
              optiontext: text,
              questionid: q.id
            });
          }
        }
        // — сбрасываем все старые правильные ответы
        for (const ca of correctAnswers) {
          await api.deleteCorrectAnswer(ca.id);
        }
        // — добавляем новые
        const correctIdxList = q.type === "single"
          ? [q.correct_option_index]
          : q.correct_option_indexes;
        for (const idx of correctIdxList) {
          await api.addCorrectAnswer(q.id, {
            optionid:   options[idx].id,
            questionid: q.id
          });
        }
      }
      else if (q.type === "matching") {
        // — обновляем пары
        for (const mp of matchPairs) {
          const newR = q.correct_matches[mp.lefttext];
          if (newR) {
            await api.updateMatchPair(mp.id, {
              id:         mp.id,
              lefttext:   mp.lefttext,
              righttext:  newR,
              questionid: q.id
            });
          } else {
            await api.deleteMatchPair(mp.id);
          }
        }
        // — создаём новыe пары, которых не было
        for (const [L, R] of Object.entries(q.correct_matches)) {
          if (!matchPairs.find(x => x.lefttext===L && x.righttext===R)) {
            await api.createMatchPair({ lefttext:L, righttext:R, questionid:q.id });
          }
        }
      }
      else if (q.type === "open") {
        const { openAnswers } = ans.data;
        // если openAnswers не пустой – обновляем
        if (openAnswers && openAnswers.length > 0) {
          const existing = openAnswers[0];
          await api.updateOpenAnswer(existing.id, {
            answertext: q.correct_answer_text,    // или q.correctAnswerText, в зависимости от того, как вы храните это поле
            questionid: q.id
          });
        }
        // иначе – создаём новый
        else {
          await api.addOpenAnswer(q.id, {
            answertext: q.correct_answer_text,
            questionid: q.id
          });
        }
      }
    }

    alert("Квиз успешно обновлён!");
  }
  catch (err) {
    console.error(err);
    alert("Ошибка при обновлении квиза");
  }
  finally {
    setIsLoading(false);
  }
};



//загружает квиз с сервера и устанавливает все поля формы + вопросы
//я не уверена рабочая ли это часть, но лучше не трогать
 const loadSavedQuiz = async (quizMeta) => {
  try {
    setIsLoading(true);
    const { data: quizData } = await api.getQuiz(quizMeta.id);
    setQuizId(quizMeta.id);
    setQuizTitle(quizData.title || "");
    setQuizDescription(quizData.description || "");
    // для datetime-local формат YYYY-MM-DDTHH:MM
    setQuizStart(quizData.startdate?.slice(0, 16) || "");
    setQuizEnd(  quizData.enddate?.slice(0, 16)   || "");
    setQuizDurationInput(quizData.duration || 60);

    const { data: questionsList } = await api.getQuestions(quizMeta.id);

    const detailed = await Promise.all(
      questionsList.map(async (q) => {
        const { data: ans } = await api.getAllAnswers(q.id);
        const { options: rawOptions, correctAnswers, matchPairs, openAnswers } = ans;

        // 1) сначала сконвертируем rawOptions в чистый массив строк:
        const opts = rawOptions.map(o => o.optiontext);

        // 2) вычислим индексы правильных ответов
        const correctIdx = correctAnswers
          .map(ca => rawOptions.findIndex(o => o.id === ca.optionid))
          .filter(i => i >= 0);

        // 3) определим тип
        let type;
        switch (q.questiontypeid) {
          case 1: type = "single";   break;
          case 2: type = "multiple"; break;
          case 3: type = "open";     break;
          case 4: type = "matching"; break;
          default: type = "single";
        }

        // 4) Собираем итоговый объект
        const base = {
          id:       q.id,
          question: q.questiontext,    // название вопроса
          type,
          points:   q.points,
          imageurl: q.imageurl || ""
        };

        if (type === "single" || type === "multiple") {
          return {
            ...base,
            options: opts,                        // <- вот здесь теперь просто строки
            // правильный индекс (single) или список индексов (multiple)
            correct_option_index:   type === "single"   ? correctIdx[0]      : undefined,
            correct_option_indexes: type === "multiple" ? correctIdx         : undefined
          };
        }

        if (type === "open") {
          return {
            ...base,
            correct_answer_text: openAnswers[0]?.answertext || ""
          };
        }

        if (type === "matching") {
          const left  = rawOptions
            .filter(o => !matchPairs.some(p => p.righttext === o.optiontext))
            .map(o => o.optiontext);
          const right = rawOptions
            .filter(o =>  matchPairs.some(p => p.righttext === o.optiontext))
            .map(o => o.optiontext);
          const correct = Object.fromEntries(
            matchPairs.map(p => [p.lefttext, p.righttext])
          );
          return {
            ...base,
            left_items:      left,
            right_items:     right,
            correct_matches: correct
          };
        }

        return base;
      })
    );

    setQuestions(detailed);
    setActiveTab("create");
  }
  finally {
    setIsLoading(false);
  }
};




//удаляет квиз с сервера и из локального списка
  const deleteSavedQuiz = async (id, e) => {
    e.stopPropagation();
    if (!window.confirm('Вы уверены, что хотите удалить этот квиз?')) return;

    try {
      setIsLoading(true);
      await api.deleteQuizFromServer(id);
      const updatedQuizzes = savedQuizzes.filter(q => q.id !== id);
      setSavedQuizzes(updatedQuizzes);
      localStorage.setItem('savedQuizzes', JSON.stringify(updatedQuizzes));
    } catch (error) {
      console.error("Error deleting quiz:", error);
      alert('Ошибка при удалении квиза');
    } finally {
      setIsLoading(false);
    }
  };

//добавляет текущий вопрос в список questions, очищает форму
  const addQuestion = () => {
  const base = {
    id: uuidv4(),
    question: currentQuestion.text,
    type: currentQuestion.type
  };

  let newQuestion;

  if (currentQuestion.type === "matching") {
    const allOptions = [
      ...currentQuestion.left_items.filter(Boolean),
      ...currentQuestion.right_items.filter(Boolean)
    ];
    newQuestion = {
      ...base,
      imageurl: currentQuestion.imageurl || undefined,
      left_items: currentQuestion.left_items.slice(),
      right_items: currentQuestion.right_items.slice(),
      correct_matches: currentQuestion.correct_matches,
      options: allOptions,  // 👈 добавляем options для совместимости
      points: currentQuestion.points
    };
  }

 else if (currentQuestion.type === "single") {
    newQuestion = {
      ...base,
      imageurl: currentQuestion.imageurl || undefined,
      options: currentQuestion.options.filter(t => t.trim() !== ""),
      correct_option_index: currentQuestion.correctOption,
      points: currentQuestion.points 
    }
  } else if (currentQuestion.type === "open") {
    newQuestion = {
    id: questions.length + 1,
    question: currentQuestion.text,
    imageurl: currentQuestion.imageurl || undefined,
    type: "open",
    correct_answer_text: currentQuestion.correctAnswerText,
    points: currentQuestion.points  
    };
  } else {
    newQuestion = {
      ...base,
      imageurl: currentQuestion.imageurl || undefined,
      options: currentQuestion.options.filter(t => t.trim() !== ""),
      correct_option_indexes: currentQuestion.correctOptions,
      points: currentQuestion.points 
    };
  }

  setQuestions([...questions, newQuestion]);

  setCurrentQuestion({
    id: questions.length + 2,
    text: "",
    type: "single",
    points: 1,
    options: ["", ""],
    imageurl: "", // картинка самого вопроса
    correctOption: null,
    correctOptions: [],
    left_items: [""],
    right_items: [""],
    correct_matches: {}
  });
};

//обновляет конкретный вариант ответа в массиве options
  const updateOption = (index, value) => {
    setCurrentQuestion((q) => ({
      ...q,
      options: q.options.map((o, i) => (i === index ? value : o)),
    }));
  };

//добавляет окно для ввода нового варианта ответа
  const addOption = () => {
    setCurrentQuestion({
      ...currentQuestion,
      options: [...currentQuestion.options, ""]
    });
  };

//удаляет один из вариантов ответа. Также корректирует индекс правильного ответа
  const removeOption = (index) => {
    const newOptions = currentQuestion.options.filter((_, i) => i !== index);
    setCurrentQuestion({
      ...currentQuestion,
      options: newOptions,
      correctOption: currentQuestion.correctOption === index ? null : 
                   currentQuestion.correctOption > index ? currentQuestion.correctOption - 1 : 
                   currentQuestion.correctOption,
      correctOptions: currentQuestion.correctOptions
        .map(opt => opt > index ? opt - 1 : opt)
        .filter(opt => opt !== index)
    });
  };
//обрабатывает установку/снятие правильных ответов 
  const handleCorrectAnswerChange = (index, isChecked) => {
    if (currentQuestion.type === "single") {
      setCurrentQuestion({...currentQuestion, correctOption: index});
    } else {
      const newCorrect = isChecked 
        ? [...currentQuestion.correctOptions, index]
        : currentQuestion.correctOptions.filter(i => i !== index);
      setCurrentQuestion({...currentQuestion, correctOptions: newCorrect});
    }
  };

//собирает YAML-файл из текущего состояния квиза и запускает его скачивание
  const generateYAML = () => {
    const quiz = {
      quiz: {
        title: quizTitle,
        description: quizDescription,
        duration: quizDurationInput,
        start: quizStart,
        end: quizEnd,
        questions: questions
      }
    };
    
    const yamlStr = yaml.dump(quiz);
    const blob = new Blob([yamlStr], { type: "text/yaml" });
    const url = URL.createObjectURL(blob);
    
    const a = document.createElement("a");
    a.href = url;
    let filename = quizTitle.trim()
  .replace(/\s+/g, "_")
  .replace(/[^\w\-]/g, "")
  .slice(0, 50); // чтобы файл не был слишком длинным

    if (!filename) {
      filename = "quiz_config";
    }

  
    a.download = `${filename}.yaml`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  };
//Подготавливает и показывает модальное окно предпросмотра на основе текущих данных в GUI
  const previewQuiz = () => {
    setQuizConfig({
      quiz: {
        title: quizTitle,
        description: quizDescription,
        start: quizStart,
        end: quizEnd,
        duration: quizDurationInput,
        questions: questions
      }
    });
    if (quizStart && quizEnd && new Date(quizStart) >= new Date(quizEnd)) {
      alert("Дата окончания должна быть позже даты начала");
      return;
    }
    setShowModal(true);
  };
  const handleShowResults = async (quizId, quizTitle) => {
  try {
    setIsLoading(true);
    const { data } = await api.getGrades();

    // фильтруем по quizId
    const filtered = data.filter(g => g.quizid === quizId);

    // сохраняем в state
    setSelectedQuizResults(filtered);
    setSelectedQuizTitle(quizTitle);
  } catch (err) {
    console.error("Ошибка при загрузке результатов:", err);
    alert("Не удалось загрузить результаты.");
  } finally {
    setIsLoading(false);
  }
};



  const renderGUICreator = () => (
    <div className="editor-container">
      <h2 className="editor-title">Create new quiz</h2>
      
      <div className="editor-field">
        <label className="editor-label">Quiz Title:</label>
        <input
          type="text"
          value={quizTitle}
          onChange={(e) => setQuizTitle(e.target.value)}
          className="editor-input"
          placeholder="Enter quiz title"
        />
      </div>

      <div className="editor-field">
        <label className="editor-label">Description:</label>
        <textarea
          value={quizDescription}
          onChange={(e) => setQuizDescription(e.target.value)}
          className="editor-textarea"
          placeholder="Enter quiz description"
        />
      </div>

      <div className="editor-field duration-field">
        <label className="editor-label">Duration (сек):</label>
        <input
          type="number"
          min="1"
          value={quizDurationInput}
          onChange={e => setQuizDurationInput(Number(e.target.value))}
          className="duration-input"
        />
      </div>

      <div className="editor-field"> {/*открытие квиза поле ввода*/}
        <label className="editor-label">Start Date & Time (МСК):</label>
        <input
          type="datetime-local"
          value={quizStart}
          onChange={(e) => setQuizStart(e.target.value)}
          className="editor-input"
        />
      </div>

      <div className="editor-field"> {/*закрытие квиза поле ввода*/}
        <label className="editor-label">End Date & Time (МСК):</label>
        <input
          type="datetime-local"
          value={quizEnd}
          onChange={(e) => setQuizEnd(e.target.value)}
          className="editor-input"
        />
      </div>
      <div className="question-container">
        <h3 className="section-title">Add Question</h3>
        <div className="editor-field">
        <label className="editor-label">Question Image URL:</label>
        <input
          type="text"
          value={currentQuestion.imageurl || ""}
          onChange={e => setCurrentQuestion({ ...currentQuestion, imageurl: e.target.value.trim() })}
          className="editor-input"
          placeholder="https://example.com/pic.png"
        />
        {currentQuestion.imageurl && (
          <img src={currentQuestion.imageurl}
              alt=""
              style={{ maxHeight: 80, marginTop: 6 }}/>
        )}
      </div>
        <div className="editor-field">
          <label className="editor-label">Question Text:</label>
          <input
            type="text"
            value={currentQuestion.text}
            onChange={(e) => setCurrentQuestion({...currentQuestion, text: e.target.value})}
            className="editor-input"
            placeholder="Enter your question"
          />
        </div>

        <div className="editor-field">
          <label className="editor-label">Points:</label>
          <input
            type="number"
            min="1"
            value={currentQuestion.points}
            onChange={e => setCurrentQuestion({
              ...currentQuestion,
              points: Math.max(1, Number(e.target.value))
            })}
            className="editor-input"
          />
        </div>

        <div className="editor-field">
          <label className="editor-label">Question Type:</label>
          <select
            value={currentQuestion.type}
            onChange={(e) => setCurrentQuestion({
              ...currentQuestion, 
              type: e.target.value,
              correctOption: null,
              correctOptions: []
            })}
            className="editor-select"
          >
            <option value="single">Single correct answer</option>
            <option value="multiple">Multiple correct answers</option>
            <option value="matching">Matching correct answer</option>
            <option value="open">Open answer</option>
          </select>
        </div>
      {/* если че удалить, хз тут ли опен ансвер должны быть */}
        {currentQuestion.type === "open" && (
  <div className="editor-field">
    <label className="editor-label">Expected Answer:</label>
    <input
      type="text"
      value={currentQuestion.correctAnswerText}
      onChange={e => setCurrentQuestion({
        ...currentQuestion,
        correctAnswerText: e.target.value
      })}
      className="editor-input"
      placeholder="Enter correct answer text"
    />
  </div>
)}

    {currentQuestion.type === "matching" && (
  <div>
    {/* Left column */}
    <div className="editor-field">
      <label className="editor-label">Left Column:</label>
      {currentQuestion.left_items.map((item, index) => (
        <div key={index} style={{ display: "flex", alignItems: "center", gap: 8, marginBottom: 4 }}>
          <input
            type="text"
            value={item}
            onChange={e => {
              const newLeft = [...currentQuestion.left_items];
              newLeft[index] = e.target.value;
              setCurrentQuestion({...currentQuestion, left_items: newLeft});
            }}
            placeholder={`Left ${index + 1}`}
            className="option-input"
          />
        </div>
      ))}
      <button
        onClick={() => {
          if (currentQuestion.left_items.length >= currentQuestion.right_items.length) return;
          setCurrentQuestion({...currentQuestion, left_items: [...currentQuestion.left_items, ""]});
        }}
        className="add-option-btn"
        disabled={currentQuestion.left_items.length >= currentQuestion.right_items.length}
      >
        <FiPlus size={16} /> Add Left
      </button>
      <button
        onClick={() => {
          const newLeft = [...currentQuestion.left_items];
          newLeft.pop();
          setCurrentQuestion({
            ...currentQuestion,
            left_items: newLeft
          });
        }}
        className="remove-option-btn"
        disabled={currentQuestion.left_items.length === 0}
        style={{ marginLeft: 8 }}
      >
        <FiTrash2 /> Remove Left
      </button>
    </div>

    {/* Right column */}
    <div className="editor-field">
      <label className="editor-label">Right Column:</label>
      {currentQuestion.right_items.map((item, index) => (
        <div key={index} style={{ display: "flex", alignItems: "center", gap: 8, marginBottom: 4 }}>
          <input
            type="text"
            value={item}
            onChange={e => {
              const newRight = [...currentQuestion.right_items];
              newRight[index] = e.target.value;
              setCurrentQuestion({...currentQuestion, right_items: newRight});
            }}
            placeholder={`Right ${index + 1}`}
            className="option-input"
          />
        </div>
      ))}
      <button
        onClick={() => {
          const diff = currentQuestion.right_items.length - currentQuestion.left_items.length;
          if (diff >= 5) return;
          setCurrentQuestion({...currentQuestion, right_items: [...currentQuestion.right_items, ""]});
        }}
        className="add-option-btn"
        disabled={currentQuestion.right_items.length - currentQuestion.left_items.length >= 5}
      >
        <FiPlus size={16} /> Add Right
      </button>
      <button
        onClick={() => {
          let newRight = [...currentQuestion.right_items];
          let newLeft = [...currentQuestion.left_items];
          if (newRight.length === newLeft.length) {
            newLeft.pop();
          }
          newRight.pop();
          setCurrentQuestion({
            ...currentQuestion,
            right_items: newRight,
            left_items: newLeft
          });
        }}
        className="remove-option-btn"
        disabled={currentQuestion.right_items.length === 0}
        style={{ marginLeft: 8 }}
      >
        <FiTrash2 /> Remove Right
      </button>
    </div>

    {/* Correct matches */}
    <div className="editor-field">
      <label className="editor-label">Correct Matches:</label>
      {currentQuestion.left_items.map((left, i) => (
        <div key={i} style={{ display: "flex", alignItems: "center", marginBottom: 4 }}>
          <span style={{ width: 80 }}>{left || `Left ${i + 1}`}</span>
          <select
            value={currentQuestion.correct_matches[left] || ""}
            onChange={e => {
              setCurrentQuestion({
                ...currentQuestion,
                correct_matches: {
                  ...currentQuestion.correct_matches,
                  [left]: e.target.value
                }
              });
            }}
            className="editor-select"
          >
            <option value="">Select match</option>
            {currentQuestion.right_items.map((right, j) => (
              <option key={j} value={right}>{right}</option>
            ))}
          </select>
        </div>
      ))}
    </div>
  </div>
)}


{(currentQuestion.type === "single" || currentQuestion.type === "multiple") && (
  <div className="editor-field">
    <label className="editor-label">Options:</label>

    {currentQuestion.options.map((opt, index) => (
      <div key={index} className="option-row">
        {/* radio / checkbox */}
        <input
          type={currentQuestion.type === "single" ? "radio" : "checkbox"}
          name={`correct-${currentQuestion.id}`}
          checked={
            currentQuestion.type === "single"
              ? currentQuestion.correctOption === index
              : currentQuestion.correctOptions.includes(index)
          }
          onChange={(e) => handleCorrectAnswerChange(index, e.target.checked)}
          className="correct-answer-input"
        />

        {/* текст варианта */}
        <input
          type="text"
          value={opt}
          onChange={(e) => updateOption(index, e.target.value)}
          className="option-input"
          placeholder={`Option ${index + 1}`}
        />

        {/* удалить */}
        <button
          onClick={() => removeOption(index)}
          className="remove-option-btn"
        >
          <FiTrash2 />
        </button>
      </div>
    ))}

    <button onClick={addOption} className="add-option-btn">
      <FiPlus size={16} /> Add Option
    </button>
  </div>
)}


        <button
          onClick={addQuestion}
          disabled={isDisabled()}
          className="add-question-btn"
        >
          Add Question
        </button>
      </div>

      {questions.length > 0 && (
        <div>
          <h3 className="section-title">Added Questions ({questions.length})</h3>
          <div className="questions-list">
            {questions.map((q, i) => (
  <EditorForQuestion
    key={q.id}
    question={q}
    onUpdate={(updatedQ) =>
      setQuestions((prev) =>
        prev.map((x) => (x.id === updatedQ.id ? updatedQ : x))
      )
    }
    onDelete={() =>
      setQuestions((prev) => prev.filter((x) => x.id !== q.id))
    }
  />
))}
          </div>

          <div className="actions-container">
            <button
              onClick={previewQuiz}
              className="action-btn preview-btn"
            >
              <FiEye /> Preview Quiz
            </button>
            <button
              onClick={generateYAML}
              disabled={!quizTitle || questions.length === 0}
              className="action-btn export-btn"
            >
              <FiUpload /> Export as YAML
            </button>
            <button
              onClick={quizId ? updateQuiz : saveQuiz}
              disabled={!quizTitle.trim() || questions.length === 0 || isLoading}
              className="action-btn save-btn"
            >
              {isLoading ? (quizId ? "Updating…" : "Saving…") : <><FiSave /> {quizId ? "Update Quiz" : "Save Quiz"}</>}
            </button>

          </div>
        </div>
      )}
    </div>
  );

  // 1) Функция для загрузки полного превью из БД
const handlePreviewSavedQuiz = async (quizMeta) => {
  setIsLoading(true);
  try {
    // a) метаданные квиза
    const { data: quizData } = await api.getQuiz(quizMeta.id);

    // b) список вопросов
    const { data: questionsList } = await api.getQuestions(quizMeta.id);

    // c) получение деталей по каждому вопросу
    const questions = await Promise.all(
      questionsList.map(async (q) => {
        const { data: ans } = await api.getAllAnswers(q.id);

        let type;
        if      (q.questiontypeid === 1) type = "single";
        else if (q.questiontypeid === 2) type = "multiple";
        else if (q.questiontypeid === 3) type = "open";
        else if (q.questiontypeid === 4) type = "matching";
        else                             type = "single";

        if (type === "single") {
          const options = ans.options.map((o) => o.optiontext);
          const correctOid = ans.correctAnswers[0]?.optionid;
          const correctIndex = options.findIndex((_, i) => ans.options[i].id === correctOid);
          return {
            id: q.id,
            question: q.questiontext,
            type,
            options,
            correct_option_index: correctIndex >= 0 ? correctIndex : undefined,
            imageurl: q.imageurl,
            points: q.points
          };
        }

        if (type === "multiple") {
          const options = ans.options.map((o) => o.optiontext);
          const correctIndexes = ans.correctAnswers
            .map((ca) => ca.optionid)
            .map((oid) => ans.options.findIndex((o) => o.id === oid))
            .filter((i) => i >= 0)
            .sort();
          return {
            id: q.id,
            question: q.questiontext,
            type,
            options,
            correct_option_indexes: correctIndexes,
            imageurl: q.imageurl,
            points: q.points
          };
        }

        if (type === "open") {
          return {
            id: q.id,
            question: q.questiontext,
            type,
            correct_answer_text: ans.openAnswers[0]?.answertext || "",
            imageurl: q.imageurl,
            points: q.points
          };
        }

        if (type === "matching") {
          const left_items = ans.options
            .filter((o) => !ans.matchPairs.some((mp) => mp.righttext === o.optiontext))
            .map((o) => o.optiontext);

          const right_items = ans.options
            .filter((o) => ans.matchPairs.some((mp) => mp.righttext === o.optiontext))
            .map((o) => o.optiontext);

          const correct_matches = Object.fromEntries(
            ans.matchPairs.map((mp) => [mp.lefttext, mp.righttext])
          );

          return {
            id: q.id,
            question: q.questiontext,
            type,
            left_items,
            right_items,
            correct_matches,
            imageurl: q.imageurl,
            points: q.points
          };
        }

        return {
          id: q.id,
          question: q.questiontext,
          type: "single",
          options: ans.options.map((o) => o.optiontext),
          correct_option_index: undefined,
          imageurl: q.imageurl,
          points: q.points
        };
      })
    );

    // d) Открываем превью
    setQuizConfig({
      quiz: {
        title: quizData.title,
        description: quizData.description,
        start: quizData.startdate,
        end: quizData.enddate,
        duration: quizData.duration,
        questions
      }
    });
    setShowModal(true);
  } catch (err) {
    console.error("Ошибка при загрузке превью квиза:", err);
    alert("Не удалось загрузить превью теста из базы");
  } finally {
    setIsLoading(false);
  }
};


const yamlExample = `QUIZ EXAMPLE:
quiz:
  title: Sample
  description: This quiz includes all question types with point values.
  duration: 300  # секунды
  start: 2025-06-25T10:00
  end: 2025-06-30T23:59
  questions:
    - id: 1
      type: single
      question: What is the capital of France?
      options:
        - Berlin
        - Paris
        - Madrid
      correct_option_index: 1
      points: 2

    - id: 2
      type: multiple
      question: Select all prime numbers
      options:
        - "4"
        - "7"
        - "11"
        - "9"
      correct_option_indexes:
        - 1
        - 2
      points: 3

    - id: 3
      type: matching
      question: Match the country to its flag color
      left_items:
        - Japan
        - Germany
        - Italy
      right_items:
        - Red circle
        - Black, red, yellow
        - Green, white, red
      correct_matches:
        Japan: Red circle
        Germany: Black, red, yellow
        Italy: Green, white, red
      points: 4

    - id: 4
      type: open
      question: Who wrote 'War and Peace'?
      correct_answer_text: Leo Tolstoy
      points: 5

`;


  const renderSavedQuizzes = () => {
  // Если выбраны результаты — показываем только их
  if (selectedResultsQuiz) {
    return (
      <div className="results-container">
        <h2 className="section-title">Results for: {selectedResultsQuiz.title}</h2>
        {gradesLoading ? (
          <p className="loading-text">Loading results...</p>
        ) : grades.length === 0 ? (
          <p className="no-results-text">There are no results for this test.</p>
        ) : (
          <table className="results-table">
            <thead>
              <tr>
                <th>Name</th>
                <th>Points</th>
              </tr>
            </thead>
            <tbody>
              {grades.map((g) => (
                <tr key={g.id}>
                  <td>{g.login}</td>
                  <td>{g.points}</td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
        <button
          onClick={() => setSelectedResultsQuiz(null)}
          className="action-btn"
          style={{ marginTop: "20px" }}
        >
          Back to Saved Quizzes
        </button>
      </div>
    );
  }

  // Если результаты не выбраны — показываем список квизов
  return (
    <div className="saved-quizzes-container">
      <h2 className="section-title">Saved Quizzes</h2>

      {isLoading ? (
        <p className="loading-text">Loading quizzes...</p>
      ) : savedQuizzes.length === 0 ? (
        <p className="no-quizzes-text">No saved quizzes yet</p>
      ) : (
        <div className="quizzes-list">
          {savedQuizzes.map((quiz) => (
             <div
                key={quiz.id}
                className="quiz-item"
                style={{ cursor: 'pointer' }}
                onClick={() => {
                  // 1) Переключаем вкладку на “Create”
                  setActiveTab('create');
                  // 2) Загружаем квиз в форму
                  loadSavedQuiz(quiz);
                }}
              >
              <div className="quiz-header">
                <div>
                  <h3 className="quiz-title">{quiz.title}</h3>
                  <p className="quiz-description">{quiz.description}</p>
                  <p className="quiz-date">
                    {new Date(quiz.createdAt).toLocaleString()}
                  </p>
                </div>
                <div style={{ display: "flex", gap: "10px", alignItems: "center" }}>
                  <button
                    onClick={(e) => {
                      e.stopPropagation();
                      handlePreviewSavedQuiz(quiz)}}
                    className="preview-quiz-btn"
                  >
                    <FiEye />
                  </button>
                  <button
                    onClick={async (e) => {
                      e.stopPropagation();
                      setSelectedResultsQuiz(quiz);
                      setGradesLoading(true);
                      try {
                        // 1) Загружаем оценки
                        const resGrades = await api.getGrades();
                        const filtered = resGrades.data.filter((g) => g.quizid === quiz.id);

                        // 2) Загружаем пользователей
                        const resUsers = await api.getUsers();
                        const users = resUsers.data;

                        // 3) Делаем карту userid -> login
                        const userMap = new Map(users.map(u => [u.userid, u.userlogin]));

                        // 4) В каждый grade добавляем поле login
                        const gradesWithLogin = filtered.map(g => ({
                          ...g,
                          login: userMap.get(g.userid) || `User #${g.userid}`
                        }));

                        // 5) Сохраняем
                        setGrades(gradesWithLogin);
                      } catch (err) {
                        console.error(err);
                        setGrades([]);
                      } finally {
                        setGradesLoading(false);
                      }
                    }}
                    className="results-quiz-btn"
                  >
                    Results
                  </button>
                   <button
                    onClick={(e) => {
                      e.stopPropagation();
                      deleteSavedQuiz(quiz.id, e)}}
                    className="delete-quiz-btn"
                  >
                    <FiTrash2 />
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};




  const renderUploadTab = () => (
    <div className="upload-container">
      <h2 className="section-title">Upload Quiz Config</h2>
      <label
        htmlFor="config-upload"
        className="upload-label"
      >
        <FiUpload style={{ marginRight: "8px" }} /> Select YAML File
        <input
          type="file"
          id="config-upload"
          style={{ display: "none" }}
          accept=".yaml,.yml"
          onChange={handleFileUpload}
          disabled={isLoading}
        />
      </label>

      {isLoading && <p className="loading-text">Uploading file...</p>}

      {error && (
        <div className="error-message">
          {error}
        </div>
      )}

      <div className="yaml-editor-container">
      <div
        className="yaml-editor-header"
        style={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center"
        }}
      >
        <h3 className="editor-titleg" style={{ margin: 0 }}>YAML Editor</h3>
        <div>
          <button
            className="icon-btn"
            onClick={handleUndo}
            disabled={historyIndex === 0}
            title="Undo (Ctrl+Z)"
          >
            <FiRotateCcw size={18} />
          </button>
          <button
            className="icon-btn"
            onClick={handleRedo}
            disabled={historyIndex === yamlHistory.length - 1}
            title="Redo (Ctrl+Y)"
          >
            <FiRotateCw size={18} />
          </button>
        </div>
      </div>
        <textarea
          value={yamlText}
          onChange={handleYamlChange}
          className="yaml-textarea" //штука чтоб высвечивалась в конструкторе yaml файла
          placeholder={yamlExample}
          rows={20}
        />
        
        {yamlError && (
          <div className="error-message">
            {yamlError}
          </div>
        )}
        
        <div className="editor-buttons">
          <button
            onClick={exportYaml}
            className="action-btn"
            disabled={!yamlText.trim()}
          >
            <FiUpload /> Export YAML
          </button>
          <button
            onClick={openPreviewFromYaml}
            className="action-btn"
            disabled={!yamlText.trim()}
          >
            <FiEye /> Open Preview
          </button>
          <button
            onClick={openInGuiCreator}
            className="action-btn"
            disabled={!yamlText.trim()}
          >
            <FiCode /> Open in GUI-Creator
          </button>
        </div>
      </div>
    </div>
  );

  return (
    <div className="quiz-creator-container">
      <Sidebar width={sidebarWidth} setWidth={setSidebarWidth} onNavigate={setActiveTab}/>
      
      <div className="main-content" style={{ marginLeft: sidebarWidth }}>
        <div className="header-container">
          <h1 className="page-title">Quiz Creator</h1>
        </div>

        <div className="tab-container">
          <button
            className={`tab-button ${activeTab === 'upload' ? 'active' : 'inactive'}`}
            onClick={() => setActiveTab('upload')}
          >
            <FiUpload /> Upload Config
          </button>
          
          <button
            className={`tab-button ${activeTab === 'create' ? 'active' : 'inactive'}`}
            onClick={() => setActiveTab('create')}
          >
            <FiPlus /> Create New
          </button>
          
          <button
            className={`tab-button ${activeTab === 'saved' ? 'active' : 'inactive'}`}
            onClick={() => setActiveTab('saved')}
          >
            <FiSave /> Saved Quizzes
          </button>
        </div>

        {activeTab === 'upload' ? renderUploadTab() : 
         activeTab === 'create' ? renderGUICreator() : 
         renderSavedQuizzes()}

        <QuizModal
          visible={showModal}
          onClose={() => setShowModal(false)}
          quizConfig={quizConfig}
        />
      </div>
    </div>
  );
}

export default ConfigUploadPage;