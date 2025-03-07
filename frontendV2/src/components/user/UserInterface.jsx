import React, { useContext, useState } from 'react';
import axios from 'axios';
import Header from './Header';
import SlotsList from './SlotsList';
import "./user.css"
import { LanguageContext } from './translation/LanguageContext';
import { translations } from './translation/translations';

const UserInterface = () => {
  const [slots, setSlots] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [activeService, setActiveService] = useState(null);
  const [selectedSlot, setSelectedSlot] = useState(null);
  const {language, toggleLanguage} = useContext(LanguageContext)

  const handleServiceClick = async (serviceId) => {
    setSelectedSlot(null); // Сбрасываем выбранный слот
    setSlots([]); // Очищаем список слотов
    setActiveService(serviceId); // Устанавливаем активную услугу
    setLoading(true);
    setError(null);

    try {
      const response = await axios.get(`http://localhost:8000/api/user/${serviceId}/status`);
      setSlots(response.data.Data);
    } catch (err) {
      setError('Ошибка при загрузке данных');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleSlotClick = (slot) => {
    setSelectedSlot(slot);
  };

  const handleBookingSubmit = async (slotId, fullName, passportNumber) => {
    try {
      const response = await axios.post(
        `http://localhost:8000/api/user/register/${slotId}`,
        {
          fullName,
          passportNumber,
        }
      );
      alert('Запись успешно оформлена!');
      console.log('Ответ сервера:', response.data);
    } catch (err) {
      setError('Ошибка при отправке данных');
      console.error(err);
    }
  };

  const handleClose = () => {
    setSelectedSlot(null); // Сбрасываем выбранный слот
    setActiveService(null);
    setSlots([]);
  };

  return (
    <div className="container">
      <button className='languageButton' onClick={toggleLanguage}>{translations[language].changeLanguage} </button>
      <Header activeService={activeService} onServiceClick={handleServiceClick} language={language} />
      <SlotsList
        slots={slots}
        loading={loading}
        error={error}
        selectedSlot={selectedSlot}
        activeService={activeService}
        onSlotClick={handleSlotClick}
        onBookingSubmit={handleBookingSubmit}
        onClose={handleClose}
        language={language}
      />
    </div>
  );
};

export default UserInterface;