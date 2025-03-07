import React, { useState } from 'react';

const LoginForm = ({ onLogin }) => {
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const [service, setService] = useState(1); // По умолчанию выбрана услуга 1
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');

    try {
      await onLogin(login, password, service);
    } catch (err) {
      setError('Повторите попытку');
    }
  };

  return (
    <div className="operatorLoginContainer">
      <h2>Авторизация оператора</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Логин:
          <input
            type="text"
            value={login}
            onChange={(e) => setLogin(e.target.value)}
            required
          />
        </label>
        <label>
          Пароль:
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </label>
        <label>
          Услуга:
          <select value={service} onChange={(e) => setService(Number(e.target.value))}>
            <option value={1}>Продление визы</option>
            <option value={2}>Продление миграционного учета</option>
            <option value={3}>Подача документов на приглашение</option>
          </select>
        </label>
        <button type="submit">Войти</button>
      </form>
      {error && <p className="error">{error}</p>}
    </div>
  );
};

export default LoginForm;