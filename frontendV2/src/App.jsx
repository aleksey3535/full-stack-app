import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import { LanguageProvider } from './components/user/translation/LanguageContext';
import UserInterface from './components/user/UserInterface';
import OperatorInterface from './components/operator/OperatorInterface';

const App = () => {
  return (
    <LanguageProvider>
      <Router>
        <Routes>
          <Route path="/" element={<UserInterface />} /> {/* Основной интерфейс */}
          <Route path="/operator" element={<OperatorInterface />} /> {/* Интерфейс оператора */}
        </Routes>
      </Router>
    </LanguageProvider>
  );
};

export default App;