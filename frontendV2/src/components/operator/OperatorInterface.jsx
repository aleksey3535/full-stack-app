import React, { useState } from 'react';
import axios from 'axios';
import LoginForm from './LoginForm';
import UserList from './UserList';
import "./operator.css"

const OperatorInterface = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [users, setUsers] = useState([]); // Данные о пользователях
  const [serviceId, setServiceId] = useState(1); // Выбранная услуга

  const handleLogin = async (login, password, service) => {
    try {
      // Авторизация
      const authResponse = await axios.post(
        'http://localhost:8000/api/admin/login',
        {
          login,
          password,
        }
      );

      if (authResponse.status === 200) {
        setIsAuthenticated(true); // Успешная авторизация
        setServiceId(service); // Сохраняем выбранную услугу

        // Получаем данные о пользователях
        await fetchUsers(service);
      } else {
        throw new Error('Ошибка авторизации');
      }
    } catch (err) {
      throw new Error('Повторите попытку');
    }
  };

  const fetchUsers = async (service) => {
    const usersResponse = await axios.get(
      `http://localhost:8000/api/admin/${service}/status`
    );
    setUsers(usersResponse.data.data); // Сохраняем данные о пользователях
  };

  const handleUserDone = async (serviceId, userId) => {
    try {
      // Отправляем запрос на сервер для отметки пользователя как "обслуженного"
      await axios.get(
        `http://localhost:8000/api/admin/${serviceId}/done/${userId}`
      );

      // Обновляем список пользователей
      await fetchUsers(serviceId);
    } catch (err) {
      console.error('Ошибка при отметке пользователя:', err);
    }
  };

  return (
    <div>
      {isAuthenticated ? (
        <UserList
          users={users}
          serviceId={serviceId}
          onUserDone={handleUserDone}
        />
      ) : (
        <LoginForm onLogin={handleLogin} />
      )}
    </div>
  );
};

export default OperatorInterface;