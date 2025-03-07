import React from 'react';


const SlotButton = ({ slot, onSlotClick }) => {
  return (
    <button className="slotButton" onClick={() => onSlotClick(slot)}>
      {slot.startTime}
    </button>
  );
};

export default SlotButton;