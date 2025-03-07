import React from 'react';
import SlotButton from './SlotButton';
import BookingForm from './BookingForm';
import { translations } from './translation/translations';

const SlotsList = ({
  slots,
  loading,
  error,
  selectedSlot,
  activeService,
  onSlotClick,
  onBookingSubmit,
  onClose,
  language
}) => {
  if (loading) return <p className="loading">{translations[language].loading} </p>;
  if (error) return <p className="error">{error}</p>;

  return (
    <div>
      {selectedSlot ? (
        <BookingForm
          slot={selectedSlot}
          serviceId={activeService}
          onBookingSubmit={onBookingSubmit}
          onClose={onClose}
          language={language}
        />
      ) : (
        <div className="slotList">
          {slots.map((slot) => (
            <SlotButton key={slot.id} slot={slot} onSlotClick={onSlotClick} />
          ))}
        </div>
      )}
    </div>
  );
};

export default SlotsList;