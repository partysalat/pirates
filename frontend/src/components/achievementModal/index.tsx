import React, { PropsWithChildren, useState } from 'react';
import { CardMedia, Dialog, DialogContent, DialogTitle } from '@mui/material';

import styles from './achievementModal.module.css';
import { DialogComponentProps } from '../accounting/accountingButton/AccountingButtonWithDrinkDialog';
import { Achievement } from '../../contexts/usersContext';
import { Achievement as NewsAchievement } from '../../contexts/newsContext';

interface Props extends DialogComponentProps {
  achievement: Achievement | NewsAchievement;
}
export const ClickableAchievementModal: React.FC<
  PropsWithChildren<Omit<Props, 'onClose' | 'open'>>
> = ({ achievement, children }) => {
  const [isOpen, setOpen] = useState(false);
  return (
    <>
      <div
        className={styles['clickable-modal']}
        onClick={() => {
          setOpen(true);
        }}
      >
        {children}
      </div>
      <AchievementModal
        achievement={achievement}
        open={isOpen}
        onClose={() => setOpen(false)}
      />
    </>
  );
};
const AchievementModal: React.FC<Props> = ({ achievement, open, onClose }) => {
  return (
    <Dialog
      fullWidth
      maxWidth="sm"
      scroll="paper"
      open={open}
      onClose={onClose}
    >
      <DialogTitle>
        {achievement.name} - {achievement.description}
      </DialogTitle>
      <DialogContent className={styles['dialog-content']}>
        <CardMedia
          component="img"
          image={achievement.image}
          alt="green iguana"
        />
      </DialogContent>
    </Dialog>
  );
};
