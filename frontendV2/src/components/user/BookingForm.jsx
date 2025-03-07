import React, { useState } from 'react';
import axios from 'axios';
import { translations } from './translation/translations';

const BookingForm = ({ slot, serviceId, onBookingSubmit, onClose, language }) => {
  const [fullName, setFullName] = useState('');
  const [passportNumber, setPassportNumber] = useState('');
  const [bookingResult, setBookingResult] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    try {
      const response = await axios.post(
        `http://localhost:8000/api/user/register/${slot.id}`,
        {
          fullName,
          passportNumber,
        }
      );
      setBookingResult(response.data); // Сохраняем результат записи
    } catch (err) {
      setError(err.response.data.error);
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="bookingFormContainer">
      {bookingResult ? (
        <div className="bookingResult">
          <h3>{translations[language].success}</h3>
          <p>{translations[language].queueNumber}: {bookingResult.queueNumber}</p>
          <p>{translations[language].fullName}: {bookingResult.fullName}</p>
          <p>{translations[language].passportNumber}: {bookingResult.passportNumber}</p>
          <p>{translations[language].time}: {bookingResult.startTime}</p>
          <p>{translations[language].cabinet}: {bookingResult.cabinet}</p>
          <button onClick={onClose}>{translations[language].close} </button>
        </div>
      ) : (
        <form className="bookingForm" onSubmit={handleSubmit}>
          <p>{translations[language].selectedTime}: {slot.startTime}</p>
          <label>
            {translations[language].fullName}:
            <input
              type="text"
              value={fullName}
              onChange={(e) => setFullName(e.target.value)}
              required
            />
          </label>
          <label>
            {translations[language].passportNumber}
            <input
              type="text"
              value={passportNumber}
              onChange={(e) => setPassportNumber(e.target.value)}
              required
            />
          </label>
          <button type="submit" disabled={loading}>
            {loading ? translations[language].sending : translations[language].make}
          </button>
          {error && <p className="error">{error}</p>}
        </form>
      )}
    </div>
  );
};

export default BookingForm;