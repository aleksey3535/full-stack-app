import React from 'react';
import UserItem from './UserItem';

const UserList = ({ users, serviceId, onUserDone }) => {
  return (
    <div className="operatorDashboard">
      <h2>Текущая очередь  ( 
        {serviceId == 1 ? "Продление визы" : ""}
        {serviceId == 2 ? "Продление миграционного учета": ""}
        {serviceId == 3 ? "Подача документов на приглашение": ""})
        </h2>
      <ul className="userList">
        {!users.length && <h2>Список пуст</h2>}
        {users.map((user) => (
          <UserItem
            key={user.id}
            user={user}
            serviceId={serviceId}
            onUserDone={onUserDone}
          />
        ))}
      </ul>
    </div>
  );
};

export default UserList;