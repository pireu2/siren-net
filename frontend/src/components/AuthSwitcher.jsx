import React, { useState } from "react";
import { motion } from "framer-motion";
import { LoginForm } from "./login-form";
import { Register } from "./Register";
import Popup from "./popUp/pop-up";
import { useForm } from "react-hook-form";


const AuthShuffle = () => {

  const loginForm = useForm({
    defaultValues: {
    username: '',
    password: ''
  }});

  const initialCards = [
    { id: 1, component: <LoginForm form={loginForm}/>, name: "Login" }, // aici dau formu din componenta superioara strict pentru ca atunci cand dau register,
    { id: 2, component: <Register loginForm={loginForm}/>, name: "Register" },                    //sa pot sa dau shuffle cu informatia de acolo, daca nu, nu pot sa dau shuffle
  ];

  const [cards, setCards] = useState(initialCards);

  const handleDragEnd = (event, info, cardId) => {
    if (Math.abs(info.offset.x) > 100) {
      setCards((prevCards) => {
        const newCards = prevCards.filter((c) => c.id !== cardId);
        return [...newCards, prevCards.find((c) => c.id === cardId)];
      });
    }
  };

  const handleShuffle = (cardId) => {
    setCards((prevCards) => {
      const newCards = prevCards.filter((c) => c.id !== cardId);
      return [...newCards, prevCards.find((c) => c.id === cardId)];
    });
  };

  return (
    <div className="flex items-center justify-center min-h-screen">
      <div className="relative w-[18.75rem] h-[37.5rem]">
        {cards.map((card, index) => (
          <motion.div
            key={card.id}
            drag
            dragConstraints={{ left: 0, right: 0, top: 0, bottom: 0 }}
            onDragEnd={(event, info) => handleDragEnd(event, info, card.id)}
            initial={{ rotate: Math.random() * 10 - 5 }}
            animate={{ x: 0, y: 0, rotate: Math.random() * 10 - 5 }}
            exit={{ scale: 0.9, opacity: 0 }}
            transition={{ type: "spring", stiffness: 300, damping: 20 }}
            className="absolute w-full h-full bg-white rounded-lg shadow-lg border-4 border-white flex items-center justify-center cursor-grab"
            style={{ zIndex: cards.length - index }}
          >
            {React.cloneElement(card.component, { onShuffle: () => handleShuffle(card.id)} )}
          </motion.div>
        ))}
      </div>
    </div>
  );
};

export default AuthShuffle;
