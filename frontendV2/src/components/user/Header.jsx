import React from 'react';
import { translations } from './translation/translations';

const Header = ({ activeService, onServiceClick, language }) => {
  return (
    <div className="headerContainer">
      <h1 className="headerTitle">{translations[language].title}</h1>
      <div className="buttonContainer">
        <button
          className={`button ${activeService === 1 ? 'active' : ''}`}
          onClick={() => onServiceClick(1)}
        >
          
          {translations[language].service1}
        </button>
        <button
          className={`button ${activeService === 2 ? 'active' : ''}`}
          onClick={() => onServiceClick(2)}
        >
          {translations[language].service2}
        </button>
        <button
          className={`button ${activeService === 3 ? 'active' : ''}`}
          onClick={() => onServiceClick(3)}
        >
          {translations[language].service3}
        </button>
      </div>
    </div>
  );
};

export default Header;