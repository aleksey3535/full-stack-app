import React from 'react';

const ConfirmationModal = ({ isOpen, onConfirm, onCancel }) => {
  if (!isOpen) return null;

  return (
    <div className="modalOverlay">
      <div className="modalContent">
        <h3>Завершить эту запись?</h3>
        <div className="modalButtons">
          <button className="confirmButton" onClick={onConfirm}>
            Да
          </button>
          <button className="cancelButton" onClick={onCancel}>
            Отмена
          </button>
        </div>
      </div>
    </div>
  );
};

export default ConfirmationModal;