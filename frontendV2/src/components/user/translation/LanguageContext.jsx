import React, { createContext, useState } from 'react';

// Создаем контекст
export const LanguageContext = createContext();

// Провайдер для контекста
export const LanguageProvider = ({ children }) => {
  const [language, setLanguage] = useState('ru'); // По умолчанию русский язык

  const toggleLanguage = () => {
    setLanguage((prevLanguage) => (prevLanguage === 'ru' ? 'en' : 'ru'));
  };

  return (
    <LanguageContext.Provider value={{ language, toggleLanguage }}>
      {children}
    </LanguageContext.Provider>
  );
};