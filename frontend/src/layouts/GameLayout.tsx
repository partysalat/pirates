import React from 'react';
import { Game } from '../components/game';

export const GameLayout: React.FC = () => {
  return (
    <div
      style={{
        width: '100%',
        minHeight: '100vh',
        background: 'black',
      }}
    >
      <Game />
    </div>
  );
};
